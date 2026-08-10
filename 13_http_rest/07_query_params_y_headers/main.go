package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"
)

// =========================================================
// QUERY PARAMS Y HEADERS
// =========================================================
// Además de la ruta y el body, un request lleva información en
// dos lugares más:
//
//   QUERY PARAMS: lo que va después del "?" en la URL
//     /productos?q=alfajor&limite=10&pagina=2
//     → para FILTRAR, BUSCAR, PAGINAR (siempre en GETs)
//
//   HEADERS: metadatos clave-valor del request/response
//     Authorization: Bearer abc123     (quién sos)
//     Content-Type: application/json   (formato del body)
//     X-Request-ID: 550e8400...        (headers propios: prefijo X-)
//
// REGLA para elegir dónde va cada cosa:
//   - Identifica el recurso            → RUTA      /productos/42
//   - Modifica QUÉ se devuelve         → QUERY     ?limite=10
//   - Metadatos (auth, formato, trace) → HEADER
//   - Los datos en sí (crear/editar)   → BODY

type Producto struct {
	ID             int    `json:"id"`
	Nombre         string `json:"nombre"`
	PrecioCentavos int64  `json:"precio_centavos"`
}

func main() {
	fmt.Println("=== Query params y headers ===")

	catalogo := []Producto{
		{1, "Alfajor", 800}, {2, "Gaseosa", 1200}, {3, "Mouse", 8500},
		{4, "Teclado", 12000}, {5, "Monitor", 95000}, {6, "Notebook", 450000},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /productos", func(w http.ResponseWriter, r *http.Request) {
		// ─────────────────────────────────────────────────────
		// LEER QUERY PARAMS: r.URL.Query()
		// ─────────────────────────────────────────────────────
		// .Get() devuelve "" si el parámetro no vino → siempre
		// definí un default y validá (¡llegan como STRING!)

		q := r.URL.Query()

		// Paginación con defaults: ?limite=2&pagina=1
		limite, err := strconv.Atoi(q.Get("limite"))
		if err != nil || limite <= 0 || limite > 100 {
			limite = 3 // default razonable + tope máximo
		}
		pagina, err := strconv.Atoi(q.Get("pagina"))
		if err != nil || pagina <= 0 {
			pagina = 1
		}

		// Filtro de precio máximo en centavos: ?max=10000
		resultado := catalogo
		if maxStr := q.Get("max"); maxStr != "" {
			if maximo, err := strconv.ParseInt(maxStr, 10, 64); err == nil {
				filtrado := []Producto{}
				for _, p := range resultado {
					if p.PrecioCentavos <= maximo {
						filtrado = append(filtrado, p)
					}
				}
				resultado = filtrado
			}
		}

		// Aplicar la paginación sobre el resultado filtrado
		desde := (pagina - 1) * limite
		hasta := min(desde+limite, len(resultado))
		if desde > len(resultado) {
			desde, hasta = 0, 0
		}
		paginado := resultado[desde:hasta]

		// ─────────────────────────────────────────────────────
		// ESCRIBIR HEADERS DE RESPUESTA: metadatos útiles
		// ─────────────────────────────────────────────────────
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Total", strconv.Itoa(len(resultado))) // total sin paginar
		w.Header().Set("X-Pagina", strconv.Itoa(pagina))
		json.NewEncoder(w).Encode(paginado)
	})

	// ─────────────────────────────────────────────────────────
	// LEER HEADERS DEL REQUEST: r.Header.Get()
	// ─────────────────────────────────────────────────────────
	mux.HandleFunc("GET /quien-soy", func(w http.ResponseWriter, r *http.Request) {
		idioma := r.Header.Get("Accept-Language")
		if idioma == "" {
			idioma = "(no enviado)"
		}
		fmt.Fprintf(w, "User-Agent: %s | Accept-Language: %s",
			r.Header.Get("User-Agent"), idioma)
	})

	servidor := httptest.NewServer(mux)
	defer servidor.Close()

	// ─────────────────────────────────────────────────────────
	// PROBANDO QUERY PARAMS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Sin parámetros (usa defaults: limite=3, pagina=1) ===")
	pedir(servidor.URL + "/productos")

	fmt.Println("\n=== Página 2 ===")
	pedir(servidor.URL + "/productos?limite=3&pagina=2")

	fmt.Println("\n=== Filtro: solo hasta $10000 ===")
	pedir(servidor.URL + "/productos?max=10000")

	fmt.Println("\n=== Combinando: filtro + paginación ===")
	pedir(servidor.URL + "/productos?max=100000&limite=2&pagina=2")

	// ─────────────────────────────────────────────────────────
	// PROBANDO HEADERS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Mandando un header custom ===")
	cliente := &http.Client{Timeout: 3 * time.Second}
	req, _ := http.NewRequest("GET", servidor.URL+"/quien-soy", nil)
	req.Header.Set("Accept-Language", "es-AR") // así se setea un header
	resp, err := cliente.Do(req)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(" ", string(body))
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  r.URL.Query().Get("q")   → leer query param (llega como string!)`)
	fmt.Println(`  strconv.Atoi + default    → convertir y validar SIEMPRE`)
	fmt.Println(`  r.Header.Get("X")         → leer header del request`)
	fmt.Println(`  w.Header().Set("X", v)    → header de respuesta (antes del body)`)
	fmt.Println(`  Ruta=quién | Query=cómo | Header=metadatos | Body=datos`)
}

func pedir(url string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	resp, err := cliente.Get(url)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("  GET %s\n", url)
	fmt.Printf("  → X-Total: %s | X-Pagina: %s\n", resp.Header.Get("X-Total"), resp.Header.Get("X-Pagina"))
	fmt.Printf("  → %s", string(body))
}

func esperarServidor() {
	cliente := &http.Client{Timeout: 200 * time.Millisecond}
	for i := 0; i < 50; i++ {
		resp, err := cliente.Get("http://localhost:8080/")
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
