package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// =========================================================
// http.Client: CONSUMIR APIs DE TERCEROS
// =========================================================
// Tu API también va a ser CLIENTE de otras: cotizaciones, pagos
// (MercadoPago), envíos, clima... Las reglas de oro:
//
//   1. TIMEOUT SIEMPRE. El http.DefaultClient NO tiene timeout:
//      si la otra API se cuelga, tu programa espera PARA SIEMPRE
//      (y si esto pasa dentro de tu servidor, se te cuelgan las
//      goroutines una tras otra hasta tirarlo abajo).
//   2. Cerrar SIEMPRE el body (defer resp.Body.Close()).
//   3. Chequear el STATUS antes de usar el body: que haya
//      respuesta no significa que haya éxito.
//   4. Decodificar a un struct con SOLO los campos que usás.
//
// Para la demo, el "tercero" es un servidor de cotizaciones que
// levantamos en el mismo programa (simula la API de un banco),
// incluyendo una ruta lenta para ver el timeout en acción.

// cotizacionDTO: SOLO los campos que nos interesan de la respuesta.
// Si la API manda 30 campos más, json.Decode los ignora sin drama.
type cotizacionDTO struct {
	Moneda string  `json:"moneda"`
	Venta  float64 `json:"venta"`
}

func main() {
	fmt.Println("=== http.Client: consumir una API de terceros ===")

	// ─────────────────────────────────────────────────────────
	// EL "TERCERO": una API de cotizaciones simulada
	// ─────────────────────────────────────────────────────────
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotizacion/{moneda}", func(w http.ResponseWriter, r *http.Request) {
		valores := map[string]float64{"dolar": 1350.50, "euro": 1480.75}
		valor, existe := valores[r.PathValue("moneda")]
		if !existe {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"error":"moneda desconocida"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotizacionDTO{Moneda: r.PathValue("moneda"), Venta: valor})
	})
	// Ruta que simula una API colgada (tarda 5 segundos)
	mux.HandleFunc("GET /lenta", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		fmt.Fprint(w, "al fin respondo")
	})

	go http.ListenAndServe(":8080", mux)
	esperarServidor()

	// ─────────────────────────────────────────────────────────
	// EL CLIENTE BIEN HECHO: con timeout
	// ─────────────────────────────────────────────────────────
	// MAL:  resp, _ := http.Get(url)            ← sin timeout, nunca
	// BIEN: cliente propio con Timeout definido:
	cliente := &http.Client{Timeout: 2 * time.Second}

	// ─────────────────────────────────────────────────────────
	// CASO FELIZ: pedir, chequear status, decodificar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso feliz: cotización del dólar ===")
	cotizacion, err := obtenerCotizacion(cliente, "dolar")
	if err != nil {
		fmt.Println("  error:", err)
	} else {
		fmt.Printf("  1 %s = $%.2f pesos\n", cotizacion.Moneda, cotizacion.Venta)
	}

	// ─────────────────────────────────────────────────────────
	// STATUS != 200: la respuesta llegó, pero NO es éxito
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== La API responde 404 (moneda inexistente) ===")
	_, err = obtenerCotizacion(cliente, "bitcoin")
	fmt.Println("  error:", err)

	// ─────────────────────────────────────────────────────────
	// TIMEOUT EN ACCIÓN: la API colgada no nos cuelga a nosotros
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== La API se cuelga: nuestro timeout nos salva ===")
	fmt.Println("  (la ruta tarda 5s, nuestro timeout es 2s...)")
	inicio := time.Now()
	_, err = cliente.Get("http://localhost:8080/lenta")
	fmt.Printf("  error tras %.1fs: %v\n", time.Since(inicio).Seconds(), err != nil)
	fmt.Println("  Sin timeout, esta línea se hubiera impreso en 5s (o nunca)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  &http.Client{Timeout: ...}   → SIEMPRE; el default NO tiene")
	fmt.Println("  defer resp.Body.Close()      → SIEMPRE; libera la conexión")
	fmt.Println("  Chequear resp.StatusCode     → respuesta ≠ éxito")
	fmt.Println("  Struct con campos que usás   → el resto del JSON se ignora solo")
	fmt.Println("  El error del client          → envolvelo con contexto (%w)")
}

// obtenerCotizacion encapsula el patrón completo de consumir una API:
// request → chequear status → decodificar → devolver struct o error.
func obtenerCotizacion(cliente *http.Client, moneda string) (*cotizacionDTO, error) {
	resp, err := cliente.Get("http://localhost:8080/cotizacion/" + moneda)
	if err != nil {
		return nil, fmt.Errorf("consultando cotización de %s: %w", moneda, err)
	}
	defer resp.Body.Close() // SIEMPRE, incluso si el status es malo

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cotización de %s: la API respondió %s", moneda, resp.Status)
	}

	var c cotizacionDTO
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, fmt.Errorf("decodificando cotización: %w", err)
	}
	return &c, nil
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
