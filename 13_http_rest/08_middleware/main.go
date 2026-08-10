package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"time"
)

// =========================================================
// MIDDLEWARE: FUNCIONES QUE ENVUELVEN HANDLERS
// =========================================================
// Hay cosas que TODA ruta necesita: loguear cada request, atrapar
// panics, medir tiempos, verificar autenticación... Copiar ese
// código en cada handler sería un desastre.
//
// Un MIDDLEWARE es una función que recibe un handler y devuelve
// OTRO handler que hace algo antes/después de llamar al original:
//
//   func miMiddleware(siguiente http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // ... ANTES del handler ...
//           siguiente.ServeHTTP(w, r)   // llama al handler real
//           // ... DESPUÉS del handler ...
//       })
//   }
//
// ¿Te suena? ¡ES EL DECORADOR DE 05_funciones_y_closures/11_funciones_que_devuelven_funciones! La misma idea de
// conLog(), aplicada a http.Handler. Los middlewares son closures
// + interfaces, dos cosas que ya dominás.
//
// http.HandlerFunc(f) convierte una función común en algo que
// cumple la interfaz http.Handler (adaptador, visto en 06_interfaces_y_errores/).

// ─────────────────────────────────────────────────────────
// MIDDLEWARE 1: LOGGING con slog (visto en 09_biblioteca_estandar/06_logging/)
// ─────────────────────────────────────────────────────────
// Loguea método, ruta, y cuánto tardó cada request. En producción,
// esto es lo que te salva cuando algo anda lento o falla.

func conLogging(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()

		siguiente.ServeHTTP(w, r) // ejecutar el handler real

		slog.Info("request atendido",
			"metodo", r.Method,
			"ruta", r.URL.Path,
			"duracion", time.Since(inicio).Round(time.Microsecond),
		)
	})
}

// ─────────────────────────────────────────────────────────
// MIDDLEWARE 2: RECOVER (visto en 06_interfaces_y_errores/19_panic_recover)
// ─────────────────────────────────────────────────────────
// Si un handler entra en panic, SIN esto el panic mata la
// conexión (y en versiones simples, puede tumbar todo). Con esto,
// el panic se transforma en un 500 y el servidor sigue vivo.
// Es EXACTAMENTE el "caso real" que anticipamos en panic/recover.

func conRecover(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				slog.Error("panic atrapado", "ruta", r.URL.Path, "panic", p)
				http.Error(w, "error interno", http.StatusInternalServerError)
			}
		}()
		siguiente.ServeHTTP(w, r)
	})
}

// ─────────────────────────────────────────────────────────
// ENCADENAR MIDDLEWARES
// ─────────────────────────────────────────────────────────
// Se aplican como capas de cebolla: el último en envolver es el
// primero en ejecutarse.
//   conRecover(conLogging(mux))
//   → request → Recover → Logging → mux → handler
// (Recover va AFUERA de todo: si Logging también fallara, lo atrapa)

func main() {
	fmt.Println("=== Middleware: logging + recover ===")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "todo bien por acá")
	})

	mux.HandleFunc("GET /lento", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(80 * time.Millisecond) // simula una consulta lenta
		fmt.Fprint(w, "tardé un poco, perdón")
	})

	mux.HandleFunc("GET /bug", func(w http.ResponseWriter, r *http.Request) {
		var productos []string
		fmt.Fprint(w, productos[5]) // ¡PANIC! índice fuera de rango
	})

	// Envolvemos el mux ENTERO: los middlewares aplican a TODAS
	// las rutas de una sola vez.
	handlerFinal := conRecover(conLogging(mux))

	servidor := httptest.NewServer(handlerFinal)
	defer servidor.Close()

	// ─────────────────────────────────────────────────────────
	// PROBANDO: mirá los logs del servidor entre las respuestas
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Request normal (mirá el log de duración) ===")
	pedir(servidor.URL + "/ok")

	fmt.Println("\n=== Request lento (la duración lo delata) ===")
	pedir(servidor.URL + "/lento")

	fmt.Println("\n=== Handler con panic: el servidor NO muere ===")
	pedir(servidor.URL + "/bug")

	fmt.Println("\n=== El servidor sigue vivo después del panic ===")
	pedir(servidor.URL + "/ok")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Middleware               → func(http.Handler) http.Handler")
	fmt.Println("  Es el patrón decorador   → de 05_funciones_y_closures/11_funciones_que_devuelven_funciones, aplicado a HTTP")
	fmt.Println("  http.HandlerFunc(f)      → adapta tu función a la interfaz Handler")
	fmt.Println("  Se encadenan             → conRecover(conLogging(mux))")
	fmt.Println("  Recover SIEMPRE afuera   → atrapa panics de todo lo demás")
	fmt.Println("  Aplican a todas las rutas→ envolviendo el mux una sola vez")
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
	fmt.Printf("  → %s | %s\n", resp.Status, string(body))
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
