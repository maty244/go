package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// =========================================================
// RESPONDER JSON: EL IDIOMA DE LAS APIs
// =========================================================
// Una API REST casi siempre habla JSON. Responder JSON bien
// significa TRES cosas, siempre juntas:
//
//   1. Header "Content-Type: application/json"
//      (le dice al cliente CÓMO interpretar el body)
//   2. El status code correcto (200, 201, 404...)
//   3. El body codificado con encoding/json (visto en JSON/)
//
// Como esto se repite en CADA handler, el patrón universal es
// escribir UNA función helper y usarla en todos lados.

// ─────────────────────────────────────────────────────────
// EL STRUCT QUE SE CONVIERTE EN JSON (tags, visto en JSON/)
// ─────────────────────────────────────────────────────────

type Producto struct {
	ID     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

// respuestaError: TODAS las respuestas de error de la API tienen
// la MISMA forma. Consistencia = clientes felices.
type respuestaError struct {
	Error string `json:"error"`
}

// ─────────────────────────────────────────────────────────
// EL HELPER: responderJSON (lo vas a escribir en cada API)
// ─────────────────────────────────────────────────────────
// ORDEN IMPORTANTE: primero headers, después WriteHeader(status),
// después el body. Si escribís el body primero, Go manda un 200
// automático y tu status se pierde.

func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Ya mandamos el status, poco se puede hacer: solo loguear
		fmt.Println("error codificando JSON:", err)
	}
}

func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, respuestaError{Error: mensaje})
}

func main() {
	fmt.Println("=== Responder JSON ===")

	catalogo := map[string]Producto{
		"1": {ID: 1, Nombre: "Alfajor", Precio: 800, Stock: 15},
		"2": {ID: 2, Nombre: "Gaseosa", Precio: 1200, Stock: 3},
	}

	mux := http.NewServeMux()

	// Lista completa: un slice de structs → array JSON
	mux.HandleFunc("GET /productos", func(w http.ResponseWriter, r *http.Request) {
		lista := []Producto{catalogo["1"], catalogo["2"]}
		responderJSON(w, http.StatusOK, lista)
	})

	// Detalle: struct → objeto JSON, o error 404 CON FORMA JSON
	mux.HandleFunc("GET /productos/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, existe := catalogo[r.PathValue("id")]
		if !existe {
			// MAL: http.Error(w, "no existe", 404) → responde texto
			// plano; el cliente que espera JSON se rompe.
			// BIEN: error con la MISMA forma JSON que todo lo demás:
			responderError(w, http.StatusNotFound, "producto no encontrado")
			return
		}
		responderJSON(w, http.StatusOK, p)
	})

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor("http://localhost:8080/productos")

	// ─────────────────────────────────────────────────────────
	// PROBANDO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== GET /productos (lista → array JSON) ===")
	pedir("http://localhost:8080/productos")

	fmt.Println("\n=== GET /productos/1 (detalle → objeto JSON) ===")
	pedir("http://localhost:8080/productos/1")

	fmt.Println("\n=== GET /productos/99 (error → JSON consistente, no texto) ===")
	pedir("http://localhost:8080/productos/99")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  Content-Type: application/json  → SIEMPRE antes de escribir`)
	fmt.Println(`  w.WriteHeader(status)           → después de los headers,`)
	fmt.Println(`                                     antes del body`)
	fmt.Println(`  json.NewEncoder(w).Encode(x)    → struct → JSON directo a w`)
	fmt.Println(`  helper responderJSON            → una vez, usado en todos lados`)
	fmt.Println(`  Errores también en JSON         → {"error": "..."}, misma forma siempre`)
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
	fmt.Printf("  Status: %s | Content-Type: %s\n", resp.Status, resp.Header.Get("Content-Type"))
	fmt.Printf("  Body:   %s", string(body))
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
