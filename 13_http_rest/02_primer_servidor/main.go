package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

// =========================================================
// TU PRIMER SERVIDOR HTTP
// =========================================================
// Un servidor HTTP en Go tiene DOS piezas:
//
//   1. HANDLERS: funciones que atienden requests. La firma SIEMPRE es:
//        func(w http.ResponseWriter, r *http.Request)
//      - r  → el REQUEST que llegó (método, ruta, headers, body)
//      - w  → donde ESCRIBÍS la respuesta (status, headers, body)
//
//   2. EL SERVIDOR en sí: http.ListenAndServe(":8080", nil)
//      escucha el puerto y despacha cada request al handler
//      que corresponda según la ruta.
//
// CONEXIÓN CON INTERFACES/: por debajo, todo es la interfaz
//   type Handler interface { ServeHTTP(ResponseWriter, *Request) }
// http.HandleFunc convierte tu función en un Handler. O sea:
// las APIs en Go son interfaces + funciones, cosas que YA sabés.

// ─────────────────────────────────────────────────────────
// HANDLER 1: saludo fijo
// ─────────────────────────────────────────────────────────

func manejarInicio(w http.ResponseWriter, r *http.Request) {
	// Escribir en w es escribir el BODY de la respuesta.
	// Si no seteás nada más, Go manda status 200 automáticamente.
	fmt.Fprint(w, "¡Bienvenido al Kiosco Digital!")
}

// ─────────────────────────────────────────────────────────
// HANDLER 2: usa datos del request
// ─────────────────────────────────────────────────────────

func manejarHora(w http.ResponseWriter, r *http.Request) {
	ahora := time.Now().Format("15:04:05")
	fmt.Fprintf(w, "Hola %s, son las %s", r.URL.Query().Get("nombre"), ahora)
}

func main() {
	fmt.Println("=== Primer servidor HTTP ===")

	// ─────────────────────────────────────────────────────────
	// REGISTRAR RUTAS: http.HandleFunc(ruta, handler)
	// ─────────────────────────────────────────────────────────
	// "Cuando llegue un request a ESTA ruta, llamá a ESTA función"

	http.HandleFunc("/", manejarInicio)
	http.HandleFunc("/hora", manejarHora)

	// ─────────────────────────────────────────────────────────
	// ARRANCAR EL SERVIDOR
	// ─────────────────────────────────────────────────────────
	// http.ListenAndServe BLOQUEA para siempre (queda escuchando).
	// En un programa real, esta suele ser LA última línea de main:
	//
	//   err := http.ListenAndServe(":8080", nil)  // bloquea acá
	//
	// Para esta demo lo corremos en una goroutine, así el programa
	// puede seguir y hacerse requests a sí mismo.
	//
	// 👉 PARA PROBARLO VOS: comentá todo desde "PARTE CLIENTE" hacia
	// abajo, reemplazalo por http.ListenAndServe(":8080", nil) sin
	// goroutine, corré `go run .` y abrí en el browser:
	//   http://localhost:8080/
	//   http://localhost:8080/hora?nombre=Mati
	// (se corta con Ctrl+C)

	servidor := httptest.NewServer(http.DefaultServeMux)
	defer servidor.Close()

	// ─────────────────────────────────────────────────────────
	// PARTE CLIENTE: probamos nuestras dos rutas
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Probando las rutas ===")

	pedir(servidor.URL + "/")
	pedir(servidor.URL + "/hora?nombre=Mati")

	// ─────────────────────────────────────────────────────────
	// DETALLE IMPORTANTE: "/" MATCHEA TODO
	// ─────────────────────────────────────────────────────────
	// La ruta "/" es especial: atrapa CUALQUIER ruta que no tenga
	// un handler más específico. Por eso esto NO da 404:

	fmt.Println("\n=== '/' atrapa cualquier ruta no registrada ===")
	pedir(servidor.URL + "/lo-que-sea") // cae en manejarInicio!
	fmt.Println("  (en el tema 03 vemos cómo controlar esto bien)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Handler                → func(w http.ResponseWriter, r *http.Request)")
	fmt.Println("  r                      → lo que LLEGÓ (método, ruta, query, body)")
	fmt.Println("  w                      → lo que RESPONDÉS (status, headers, body)")
	fmt.Println("  http.HandleFunc(ruta)  → conecta ruta → handler")
	fmt.Println("  http.ListenAndServe    → arranca el servidor (bloquea)")
	fmt.Println("  Por debajo             → la interfaz http.Handler (06_interfaces_y_errores/)")
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
	fmt.Printf("  GET %-40s → %s | %s\n", url, resp.Status, string(body))
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
