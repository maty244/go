package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// =========================================================
// EJERCICIO INTEGRADOR: SIMULANDO UN REQUEST CON TIMEOUT
// =========================================================
// Esto es EXACTAMENTE lo que va a pasar en 13_http_rest/ más adelante en
// la hoja de ruta: cada request que llega a un servidor Go trae
// su propio context, y ese context se propaga a TODO lo que el
// handler necesite hacer (consultar la base de datos, llamar a
// otra API, etc.). Si el cliente se desconecta o se cumple un
// timeout, el context avisa, y las operaciones en curso pueden
// frenar en vez de seguir trabajando en vano.

// ErrTimeout es el error que devolvemos cuando una operación no
// llega a tiempo.
var ErrTimeout = errors.New("la operación no completó a tiempo")

// consultarProducto simula una consulta a un servicio externo o DB.
// Acá "tarda" según cuánto le digamos, pero respeta el ctx.
func consultarProducto(ctx context.Context, nombre string, latencia time.Duration) (string, error) {
	select {
	case <-time.After(latencia):
		return fmt.Sprintf("%s: en stock, $%.2f", nombre, 999.99), nil
	case <-ctx.Done():
		// Envolvemos tanto nuestro error de dominio (ErrTimeout) como la causa
		// del context (ctx.Err()) usando múltiples %w (soportado desde Go 1.20+).
		return "", fmt.Errorf("consultarProducto(%s): %w (%w)", nombre, ErrTimeout, ctx.Err())
	}
}

// manejarRequest simula un handler que atiende un pedido de
// consulta con un timeout total: no importa cuántos productos
// consulte, TODO el request tiene un límite de tiempo.
func manejarRequest(productos []string, latenciaPorProducto time.Duration, timeoutTotal time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutTotal)
	defer cancel()

	fmt.Printf("Request con timeout de %v para %d producto(s)\n", timeoutTotal, len(productos))

	for _, producto := range productos {
		resultado, err := consultarProducto(ctx, producto, latenciaPorProducto)
		if err != nil {
			if errors.Is(err, ErrTimeout) && errors.Is(err, context.DeadlineExceeded) {
				fmt.Printf("  ⏱  %s: se cortó por timeout\n", producto)
			} else {
				fmt.Printf("  ✗  %s: error inesperado: %v\n", producto, err)
			}
			continue
		}
		fmt.Println("  ✓ ", resultado)
	}
}

func main() {
	fmt.Println("=== EJERCICIO INTEGRADOR: requests con timeout ===")

	// ─────────────────────────────────────────────────────────
	// CASO 1: todo responde a tiempo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Caso 1: consultas rápidas, timeout generoso ---")
	manejarRequest(
		[]string{"Notebook", "Mouse"},
		50*time.Millisecond,  // cada consulta tarda 50ms
		500*time.Millisecond, // timeout total: 500ms
	)

	// ─────────────────────────────────────────────────────────
	// CASO 2: el timeout es más corto que la latencia
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Caso 2: consultas lentas, timeout corto ---")
	manejarRequest(
		[]string{"Teclado", "Monitor"},
		300*time.Millisecond, // cada consulta tarda 300ms
		150*time.Millisecond, // timeout total: solo 150ms
	)

	// ─────────────────────────────────────────────────────────
	// POR QUÉ ESTO IMPORTA EN UN SERVIDOR REAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Por qué esto importa ===")
	fmt.Println("  Sin timeout: un cliente lento podría dejar un request")
	fmt.Println("  'colgado' para siempre, consumiendo recursos del servidor.")
	fmt.Println("  Con context.WithTimeout, el servidor SIEMPRE sabe cuándo")
	fmt.Println("  cortar, sin importar qué tan lenta sea la operación.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 09_biblioteca_estandar/05_contexto/ ===")
	fmt.Println("  context.Background()   → punto de partida de cada request")
	fmt.Println("  context.WithTimeout    → límite de tiempo para TODO el request")
	fmt.Println("  ctx propagado          → viaja a cada función que hace trabajo real")
	fmt.Println("  select + ctx.Done()    → cada operación respeta el límite")
	fmt.Println("  errors.Is + ctx.Err()  → distinguir timeout de otros errores")
}
