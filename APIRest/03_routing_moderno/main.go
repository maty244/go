package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// =========================================================
// ROUTING MODERNO: EL ServeMux DE GO 1.22+
// =========================================================
// Un "router" (o mux) decide QUÉ handler atiende cada request.
// Desde Go 1.22, el ServeMux de la librería estándar soporta lo
// que antes obligaba a usar librerías externas:
//
//   1. MÉTODO en el patrón:    "GET /productos"
//      (un POST a esa ruta da 405 Method Not Allowed AUTOMÁTICO)
//   2. PARÁMETROS en la ruta:  "GET /productos/{id}"
//      (se leen con r.PathValue("id"))
//   3. "{$}" para matchear EXACTAMENTE la raíz:  "GET /{$}"
//      (soluciona el problema del tema 02 donde "/" atrapaba todo)
//
// También conviene crear TU PROPIO mux con http.NewServeMux() en
// vez de usar el global (nil): más ordenado y testeable.

// ─────────────────────────────────────────────────────────
// MAL: cómo se ruteaba ANTES de Go 1.22 (lo vas a ver en
// tutoriales viejos)
// ─────────────────────────────────────────────────────────
// Un solo handler por ruta, y adentro un switch por método,
// parseando el ID a mano desde la URL:
//
//   http.HandleFunc("/productos/", func(w, r) {
//       id := strings.TrimPrefix(r.URL.Path, "/productos/")
//       switch r.Method {
//       case http.MethodGet:  ...
//       case http.MethodPost: ...
//       default: http.Error(w, "método no permitido", 405)
//       }
//   })
//
// Funciona, pero es verboso y fácil de romper. HOY no hace falta.

func main() {
	fmt.Println("=== Routing moderno (Go 1.22+) ===")

	// ─────────────────────────────────────────────────────────
	// BIEN: método + ruta + parámetros en el patrón
	// ─────────────────────────────────────────────────────────
	mux := http.NewServeMux()

	// Solo la raíz EXACTA (gracias a {$})
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Kiosco Digital — API v1")
	})

	// GET /productos → listar
	mux.HandleFunc("GET /productos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "lista de productos: Alfajor, Gaseosa, Mouse")
	})

	// GET /productos/{id} → detalle. El {id} se lee con PathValue.
	mux.HandleFunc("GET /productos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "detalle del producto %s", id)
	})

	// POST /productos → crear. MISMA ruta que "GET /productos",
	// DISTINTO método, DISTINTO handler. El mux separa solo.
	mux.HandleFunc("POST /productos", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprint(w, "producto creado")
	})

	// Parámetros anidados: dos variables en la misma ruta
	mux.HandleFunc("GET /clientes/{cliente}/pedidos/{pedido}",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "pedido %s del cliente %s",
				r.PathValue("pedido"), r.PathValue("cliente"))
		})

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor("http://localhost:8080/")

	// ─────────────────────────────────────────────────────────
	// PROBANDO EL ROUTING
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Rutas y parámetros ===")
	pedir("GET", "http://localhost:8080/")
	pedir("GET", "http://localhost:8080/productos")
	pedir("GET", "http://localhost:8080/productos/42")
	pedir("POST", "http://localhost:8080/productos")
	pedir("GET", "http://localhost:8080/clientes/7/pedidos/1001")

	// ─────────────────────────────────────────────────────────
	// LO QUE EL MUX RESUELVE GRATIS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== 404 y 405 automáticos ===")

	// Ruta que no existe → 404 Not Found (ya no la atrapa "/",
	// porque registramos "GET /{$}" que solo matchea la raíz exacta)
	pedir("GET", "http://localhost:8080/no-existe")

	// Método no registrado para una ruta existente → 405
	pedir("DELETE", "http://localhost:8080/productos")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  mux := http.NewServeMux()      → tu propio router`)
	fmt.Println(`  "GET /productos"               → método + ruta en el patrón`)
	fmt.Println(`  "GET /productos/{id}"          → parámetro en la ruta`)
	fmt.Println(`  r.PathValue("id")              → leer el parámetro`)
	fmt.Println(`  "GET /{$}"                     → SOLO la raíz exacta`)
	fmt.Println(`  404 / 405                      → los maneja el mux automáticamente`)
}

func pedir(metodo, url string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	req, _ := http.NewRequest(metodo, url, nil)
	resp, err := cliente.Do(req)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	texto := string(body)
	if len(texto) > 48 {
		texto = texto[:48] + "..."
	}
	fmt.Printf("  %-6s %-46s → %s | %s\n", metodo, url, resp.Status, texto)
}

func esperarServidor(url string) {
	cliente := &http.Client{Timeout: 200 * time.Millisecond}
	for i := 0; i < 50; i++ {
		resp, err := cliente.Get(url)
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
