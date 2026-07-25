package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// =========================================================
// ¿QUÉ ES HTTP? (la base de toda API REST)
// =========================================================
// HTTP es el protocolo con el que se comunican un CLIENTE (el
// browser, una app, otro servidor) y un SERVIDOR. Es texto con
// un formato acordado: el cliente manda un REQUEST, el servidor
// contesta con un RESPONSE. Nada más que eso.
//
// UN REQUEST tiene:
//   - MÉTODO:  qué querés hacer (GET, POST, PUT, DELETE...)
//   - RUTA:    sobre qué recurso (/productos, /clientes/42)
//   - HEADERS: metadatos (quién sos, qué formato aceptás...)
//   - BODY:    datos que mandás (opcional; típico en POST/PUT)
//
// UN RESPONSE tiene:
//   - STATUS CODE: cómo salió (200 OK, 404 no existe, 500 explotó)
//   - HEADERS:     metadatos de la respuesta
//   - BODY:        el contenido (JSON, HTML, lo que sea)
//
// ¿Y "REST"? Es una CONVENCIÓN para diseñar APIs sobre HTTP:
//   - Cada cosa de tu negocio es un RECURSO con su ruta
//     (/productos, /productos/42)
//   - El MÉTODO dice la acción: GET lee, POST crea, PUT
//     actualiza, DELETE borra
//   - El servidor no recuerda nada entre requests (sin estado):
//     cada request trae TODO lo necesario
//
// ─────────────────────────────────────────────────────────
// CÓMO FUNCIONA ESTE ARCHIVO (y toda la sección APIRest)
// ─────────────────────────────────────────────────────────
// Este programa hace de CLIENTE y SERVIDOR a la vez:
//   1. Levanta un servidor HTTP en el puerto 8080 (en una goroutine)
//   2. Se hace requests A SÍ MISMO con http.Client
//   3. Muestra en consola todo lo que viaja, y termina solo
// Así ves la película completa con un simple: go run .

func main() {
	fmt.Println("=== ¿QUÉ ES HTTP? Un request y un response por dentro ===")

	// ─────────────────────────────────────────────────────────
	// EL SERVIDOR: un handler que "radiografía" cada request
	// ─────────────────────────────────────────────────────────
	// (La sintaxis exacta la vemos en el tema 02; hoy el foco es
	// VER qué contiene un request y un response)

	http.HandleFunc("/radiografia", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\n  [SERVIDOR] Me llegó un request:")
		fmt.Println("    Método :", r.Method)
		fmt.Println("    Ruta   :", r.URL.Path)
		fmt.Println("    Header 'User-Agent':", r.Header.Get("User-Agent"))

		if r.Body != nil {
			body, _ := io.ReadAll(r.Body)
			if len(body) > 0 {
				fmt.Println("    Body   :", string(body))
			}
		}

		// El servidor arma su RESPONSE: status + headers + body
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "Recibido tu %s a %s, todo OK", r.Method, r.URL.Path)
	})

	// Ruta "silenciosa" que usamos solo para saber que el servidor
	// ya está listo (sin ensuciar la consola).
	http.HandleFunc("/salud", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Levantamos el servidor en segundo plano (goroutine, visto
	// en Concurrencia/). ListenAndServe BLOQUEA, por eso va aparte.
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor("http://localhost:8080/salud")

	// ─────────────────────────────────────────────────────────
	// EL CLIENTE: un GET (leer) y un POST (enviar datos)
	// ─────────────────────────────────────────────────────────
	cliente := &http.Client{Timeout: 3 * time.Second}

	fmt.Println("\n[CLIENTE] Mando un GET (pedir información)...")
	resp, err := cliente.Get("http://localhost:8080/radiografia")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	mostrarResponse(resp)

	fmt.Println("\n[CLIENTE] Mando un POST (enviar datos)...")
	resp, err = cliente.Post("http://localhost:8080/radiografia",
		"text/plain", nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	mostrarResponse(resp)

	// ─────────────────────────────────────────────────────────
	// LOS MÉTODOS Y SU SIGNIFICADO REST
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Los métodos HTTP en una API REST ===")
	fmt.Println("  GET    /productos     → listar productos (leer, no modifica nada)")
	fmt.Println("  GET    /productos/42  → ver el producto 42")
	fmt.Println("  POST   /productos     → crear un producto nuevo (datos en el body)")
	fmt.Println("  PUT    /productos/42  → actualizar el producto 42")
	fmt.Println("  DELETE /productos/42  → borrar el producto 42")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  HTTP        → request (método+ruta+headers+body) ↔")
	fmt.Println("                 response (status+headers+body)")
	fmt.Println("  REST        → recursos en rutas + métodos como acciones")
	fmt.Println("  GET         → leer | POST → crear | PUT → actualizar | DELETE → borrar")
	fmt.Println("  Sin estado  → cada request lleva TODO lo necesario")
}

// mostrarResponse imprime las partes de un response HTTP.
func mostrarResponse(resp *http.Response) {
	defer resp.Body.Close() // SIEMPRE cerrar el body (libera la conexión)

	fmt.Println("  [CLIENTE] Me respondió:")
	fmt.Println("    Status :", resp.Status) // "200 OK"
	fmt.Println("    Header 'Content-Type':", resp.Header.Get("Content-Type"))

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("    Body   :", string(body))
}

// esperarServidor hace polling hasta que el servidor esté listo
// (la goroutine tarda unos milisegundos en arrancar).
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
