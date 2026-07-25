package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// =========================================================
// EJERCICIO INTEGRADOR: LA API DEL KIOSCO DIGITAL
// =========================================================
// Tu primera API REST completa, con TODO lo visto en la sección:
//
//   GET    /productos       → listar (público)
//   GET    /productos/{id}  → detalle (público)
//   POST   /productos       → crear (requiere token)
//   PUT    /productos/{id}  → actualizar (requiere token)
//   DELETE /productos/{id}  → borrar (requiere token)
//
// Piezas que se combinan:
//   - Routing moderno (tema 03) + JSON in/out (04, 05)
//   - Status codes correctos (06)
//   - Middleware de logging + recover (08) y auth (09)
//   - Almacén en memoria protegido con sync.Mutex (Concurrencia/):
//     el servidor atiende cada request en su PROPIA goroutine,
//     así que el map compartido NECESITA protección
//
// ─────────────────────────────────────────────────────────
// DOS FORMAS DE CORRERLO
// ─────────────────────────────────────────────────────────
//   go run .              → demo automática (cliente incluido)
//   SERVIDOR=1 go run .   → queda vivo escuchando en :8080 para
//                            que lo pruebes vos con curl o browser:
//
//   curl http://localhost:8080/productos
//   curl http://localhost:8080/productos/1
//   curl -X POST http://localhost:8080/productos \
//        -H "Authorization: Bearer token-secreto-del-kiosco" \
//        -H "Content-Type: application/json" \
//        -d '{"nombre":"Chicle","precio":300,"stock":50}'
//   (se corta con Ctrl+C)

// ─────────────────────────────────────────────────────────
// MODELO Y DTOs
// ─────────────────────────────────────────────────────────

type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

// productoDTO es lo que el cliente puede mandar (sin ID: lo
// asigna el servidor). Sirve para POST y PUT.
type productoDTO struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

func (dto productoDTO) validar() error {
	var problemas []string
	if strings.TrimSpace(dto.Nombre) == "" {
		problemas = append(problemas, "nombre es obligatorio")
	}
	if dto.Precio <= 0 {
		problemas = append(problemas, "precio debe ser mayor a 0")
	}
	if dto.Stock < 0 {
		problemas = append(problemas, "stock no puede ser negativo")
	}
	if len(problemas) > 0 {
		return fmt.Errorf("%s", strings.Join(problemas, "; "))
	}
	return nil
}

// ─────────────────────────────────────────────────────────
// EL ALMACÉN: map + Mutex (visto en Concurrencia/)
// ─────────────────────────────────────────────────────────
// Cada request corre en su propia goroutine. Sin el Mutex, dos
// POST simultáneos podrían pisarse (race condition). Probalo:
// go run -race . también funciona limpio.

type Almacen struct {
	mu        sync.Mutex
	productos map[int]Producto
	proximoID int
}

func NuevoAlmacen() *Almacen {
	return &Almacen{productos: map[int]Producto{}, proximoID: 1}
}

func (a *Almacen) Listar() []Producto {
	a.mu.Lock()
	defer a.mu.Unlock()
	lista := make([]Producto, 0, len(a.productos))
	for id := 1; id < a.proximoID; id++ { // orden estable por ID
		if p, ok := a.productos[id]; ok {
			lista = append(lista, p)
		}
	}
	return lista
}

func (a *Almacen) Buscar(id int) (Producto, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	p, ok := a.productos[id]
	return p, ok
}

func (a *Almacen) Crear(dto productoDTO) Producto {
	a.mu.Lock()
	defer a.mu.Unlock()
	p := Producto{ID: a.proximoID, Nombre: dto.Nombre, Precio: dto.Precio, Stock: dto.Stock}
	a.productos[p.ID] = p
	a.proximoID++
	return p
}

func (a *Almacen) Actualizar(id int, dto productoDTO) (Producto, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.productos[id]; !ok {
		return Producto{}, false
	}
	p := Producto{ID: id, Nombre: dto.Nombre, Precio: dto.Precio, Stock: dto.Stock}
	a.productos[id] = p
	return p, true
}

func (a *Almacen) Borrar(id int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.productos[id]; !ok {
		return false
	}
	delete(a.productos, id)
	return true
}

// ─────────────────────────────────────────────────────────
// HELPERS JSON (tema 04)
// ─────────────────────────────────────────────────────────

type respuestaError struct {
	Error string `json:"error"`
}

func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, respuestaError{Error: mensaje})
}

// leerDTO decodifica y valida el body (tema 05).
func leerDTO(w http.ResponseWriter, r *http.Request) (productoDTO, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var dto productoDTO
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dto); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return dto, false
	}
	if err := dto.validar(); err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return dto, false
	}
	return dto, true
}

// leerID parsea el {id} de la ruta (tema 03).
func leerID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		responderError(w, http.StatusBadRequest, "el id debe ser un número positivo")
		return 0, false
	}
	return id, true
}

// ─────────────────────────────────────────────────────────
// HANDLERS: finos, porque los helpers hacen el trabajo sucio
// ─────────────────────────────────────────────────────────

func rutasAPI(almacen *Almacen) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /productos", func(w http.ResponseWriter, r *http.Request) {
		responderJSON(w, http.StatusOK, almacen.Listar())
	})

	mux.HandleFunc("GET /productos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, ok := leerID(w, r)
		if !ok {
			return
		}
		p, existe := almacen.Buscar(id)
		if !existe {
			responderError(w, http.StatusNotFound, "producto no encontrado")
			return
		}
		responderJSON(w, http.StatusOK, p)
	})

	mux.Handle("POST /productos", conAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			dto, ok := leerDTO(w, r)
			if !ok {
				return
			}
			responderJSON(w, http.StatusCreated, almacen.Crear(dto))
		})))

	mux.Handle("PUT /productos/{id}", conAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, ok := leerID(w, r)
			if !ok {
				return
			}
			dto, ok := leerDTO(w, r)
			if !ok {
				return
			}
			p, existe := almacen.Actualizar(id, dto)
			if !existe {
				responderError(w, http.StatusNotFound, "producto no encontrado")
				return
			}
			responderJSON(w, http.StatusOK, p)
		})))

	mux.Handle("DELETE /productos/{id}", conAuth(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id, ok := leerID(w, r)
			if !ok {
				return
			}
			if !almacen.Borrar(id) {
				responderError(w, http.StatusNotFound, "producto no encontrado")
				return
			}
			w.WriteHeader(http.StatusNoContent)
		})))

	return mux
}

// ─────────────────────────────────────────────────────────
// MIDDLEWARES (temas 08 y 09)
// ─────────────────────────────────────────────────────────

const tokenValido = "token-secreto-del-kiosco" // en la vida real: variable de entorno

func conAuth(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !ok || token != tokenValido {
			responderError(w, http.StatusUnauthorized, "token faltante o inválido")
			return
		}
		siguiente.ServeHTTP(w, r)
	})
}

func conLogging(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		siguiente.ServeHTTP(w, r)
		slog.Info("request", "metodo", r.Method, "ruta", r.URL.Path,
			"duracion", time.Since(inicio).Round(time.Microsecond))
	})
}

func conRecover(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				slog.Error("panic atrapado", "ruta", r.URL.Path, "panic", p)
				responderError(w, http.StatusInternalServerError, "error interno")
			}
		}()
		siguiente.ServeHTTP(w, r)
	})
}

// ─────────────────────────────────────────────────────────
// MAIN: armar la API y (según el modo) demo o servidor real
// ─────────────────────────────────────────────────────────

func main() {
	almacen := NuevoAlmacen()
	almacen.Crear(productoDTO{Nombre: "Alfajor", Precio: 800, Stock: 15})
	almacen.Crear(productoDTO{Nombre: "Gaseosa", Precio: 1200, Stock: 3})

	api := conRecover(conLogging(rutasAPI(almacen)))

	// MODO SERVIDOR: SERVIDOR=1 go run . → queda vivo para curl/browser
	if os.Getenv("SERVIDOR") == "1" {
		fmt.Println("Kiosco Digital API escuchando en http://localhost:8080")
		fmt.Println("Probá: curl http://localhost:8080/productos  (Ctrl+C para salir)")
		if err := http.ListenAndServe(":8080", api); err != nil {
			fmt.Println("servidor:", err)
		}
		return
	}

	// MODO DEMO: servidor en goroutine + cliente que lo ejercita
	fmt.Println("=== API KIOSCO DIGITAL: demo completa ===")
	go http.ListenAndServe(":8080", api)
	esperarServidor()

	base := "http://localhost:8080/productos"
	auth := "Bearer " + tokenValido

	fmt.Println("\n─── Lecturas públicas ───")
	pedir("GET", base, "", "")
	pedir("GET", base+"/1", "", "")
	pedir("GET", base+"/99", "", "")  // 404
	pedir("GET", base+"/abc", "", "") // 400: id no numérico

	fmt.Println("\n─── Escrituras protegidas ───")
	pedir("POST", base, `{"nombre":"Mouse","precio":8500,"stock":20}`, "") // 401 sin token
	pedir("POST", base, `{"nombre":"Mouse","precio":8500,"stock":20}`, auth)
	pedir("POST", base, `{"nombre":"","precio":-1,"stock":-5}`, auth) // 400 validación
	pedir("PUT", base+"/2", `{"nombre":"Gaseosa 2L","precio":1500,"stock":10}`, auth)
	pedir("DELETE", base+"/1", "", auth) // 204
	pedir("DELETE", base+"/1", "", auth) // 404: ya no existe

	fmt.Println("\n─── Estado final del catálogo ───")
	pedir("GET", base, "", "")

	fmt.Println("\n=== RESUMEN: tu primera API REST completa ===")
	fmt.Println("  Routing 1.22+   → GET/POST/PUT/DELETE con {id} nativo")
	fmt.Println("  JSON in/out     → DTOs validados, respuestas consistentes")
	fmt.Println("  Status codes    → 200/201/204/400/401/404 según corresponde")
	fmt.Println("  Middlewares     → recover(logging(auth solo en escrituras))")
	fmt.Println("  sync.Mutex      → el almacén aguanta requests concurrentes")
	fmt.Println("  Próximo paso    → cambiar el Almacen por una base de datos")
	fmt.Println("                     real (misma interfaz, otro adentro 😉)")
}

func pedir(metodo, url, bodyJSON, authHeader string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	var body io.Reader
	if bodyJSON != "" {
		body = bytes.NewBufferString(bodyJSON)
	}
	req, _ := http.NewRequest(metodo, url, body)
	if bodyJSON != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	texto := strings.TrimSpace(string(respBody))
	if texto == "" {
		texto = "(sin body)"
	}
	fmt.Printf("  %-6s %-22s → %s | %s\n",
		metodo, strings.TrimPrefix(url, "http://localhost:8080"), resp.Status, texto)
}

func esperarServidor() {
	cliente := &http.Client{Timeout: 200 * time.Millisecond}
	for i := 0; i < 50; i++ {
		resp, err := cliente.Get("http://localhost:8080/productos")
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
