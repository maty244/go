package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// =========================================================
// GRACEFUL SHUTDOWN: APAGAR SIN CORTAR A NADIE
// =========================================================
// Cuando desplegás una versión nueva de tu API (o el servidor se
// reinicia), el proceso viejo recibe una señal para terminar.
// Si simplemente muere (os.Exit, o dejar que el proceso caiga),
// TODOS los requests que estaban a la mitad se cortan: clientes
// con errores, pagos a medias, datos corruptos.
//
// "Graceful shutdown" (apagado elegante) es:
//   1. Dejar de aceptar requests NUEVOS
//   2. Esperar a que terminen los que están EN CURSO
//   3. Recién ahí, salir (con un tope de tiempo por las dudas)
//
// Las piezas (todas ya las conocés):
//   - http.Server (el struct, no la función suelta) tiene el
//     método Shutdown(ctx) que hace exactamente los pasos 1 y 2
//   - signal.NotifyContext (Contexto/) crea un ctx que se cancela
//     cuando llega Ctrl+C o la señal de terminación del sistema
//   - context.WithTimeout le pone tope a la espera del paso 2
//
// En un programa real, el esqueleto es:
//
//   ctx, stop := signal.NotifyContext(context.Background(),
//       os.Interrupt, syscall.SIGTERM)
//   defer stop()
//
//   go srv.ListenAndServe()      // el servidor atiende
//   <-ctx.Done()                 // bloquea hasta Ctrl+C / SIGTERM
//
//   ctxApagado, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//   defer cancel()
//   srv.Shutdown(ctxApagado)     // apagado elegante con tope
//
// Para la demo, en vez de esperar a que VOS aprietes Ctrl+C,
// disparamos el shutdown desde el propio programa mientras hay
// un request lento en curso, para que VEAS que se respeta.

func main() {
	fmt.Println("=== Graceful shutdown ===")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /rapida", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "listo")
	})
	mux.HandleFunc("GET /pago-lento", func(w http.ResponseWriter, r *http.Request) {
		// Simula un pago que tarda 2 segundos en procesarse
		fmt.Println("  [SERVIDOR] procesando un pago... (tarda 2s)")
		time.Sleep(2 * time.Second)
		fmt.Println("  [SERVIDOR] pago completado ✓")
		fmt.Fprint(w, "pago procesado OK")
	})

	// ─────────────────────────────────────────────────────────
	// http.Server: EL STRUCT (más control que ListenAndServe suelto)
	// ─────────────────────────────────────────────────────────
	// Además de Shutdown, te deja configurar timeouts del servidor
	// (¡importantes en producción para clientes lentos/maliciosos!)
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,  // tope para leer headers
		WriteTimeout:      30 * time.Second, // tope para escribir respuesta
	}

	go func() {
		// Cuando Shutdown termina, ListenAndServe devuelve
		// http.ErrServerClosed: es lo ESPERADO, no un error real.
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor()

	// ─────────────────────────────────────────────────────────
	// LA DEMO: un pago en curso + shutdown en el medio
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n[CLIENTE] Disparo un pago lento (2s) en segundo plano...")

	pagoTerminado := make(chan string)
	go func() {
		cliente := &http.Client{Timeout: 10 * time.Second}
		resp, err := cliente.Get("http://localhost:8080/pago-lento")
		if err != nil {
			pagoTerminado <- "ERROR: " + err.Error()
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		pagoTerminado <- fmt.Sprintf("%s | %s", resp.Status, string(body))
	}()

	time.Sleep(300 * time.Millisecond) // el pago ya está en curso...

	fmt.Println("[PROGRAMA] ¡Llega la orden de apagar! (simula Ctrl+C / deploy)")
	fmt.Println("[PROGRAMA] Llamando a srv.Shutdown(ctx) con tope de 10s...")

	ctxApagado, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inicio := time.Now()
	if err := srv.Shutdown(ctxApagado); err != nil {
		fmt.Println("shutdown forzado:", err)
	}
	fmt.Printf("[PROGRAMA] Shutdown completado en %.1fs (esperó al pago!)\n",
		time.Since(inicio).Seconds())

	// El pago EN CURSO terminó bien a pesar del shutdown:
	fmt.Println("[CLIENTE] Resultado del pago:", <-pagoTerminado)

	// Y los requests NUEVOS ya no entran (el servidor está apagado):
	fmt.Println("\n[CLIENTE] Intento un request nuevo post-shutdown...")
	cliente := &http.Client{Timeout: 1 * time.Second}
	_, err := cliente.Get("http://localhost:8080/rapida")
	fmt.Println("  ¿Falló como corresponde?", err != nil)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  http.Server{...}          → el struct: timeouts + Shutdown")
	fmt.Println("  srv.Shutdown(ctx)         → no acepta nuevos, espera los en curso")
	fmt.Println("  signal.NotifyContext      → ctx que se cancela con Ctrl+C/SIGTERM")
	fmt.Println("  context.WithTimeout       → tope a la espera (no esperar infinito)")
	fmt.Println("  ErrServerClosed           → resultado NORMAL de un shutdown, no error")
	fmt.Println("  Por qué importa           → deploys sin cortar pagos a la mitad")
}

func esperarServidor() {
	cliente := &http.Client{Timeout: 200 * time.Millisecond}
	for i := 0; i < 50; i++ {
		resp, err := cliente.Get("http://localhost:8080/rapida")
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
