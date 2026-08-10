package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// =========================================================
// context + GOROUTINES: CANCELAR TRABAJO EN CURSO CORRECTAMENTE
// =========================================================
// En 09_biblioteca_estandar/05_contexto/ vimos context.Context con funciones simples,
// sincrónicas. Ahora lo combinamos con goroutines: el uso REAL
// más común de context es avisarle a VARIAS goroutines en
// paralelo "paren, ya no hace falta seguir".
//
// Sin esto, corrés el riesgo de un "goroutine leak": goroutines
// que quedan corriendo para siempre (consumiendo memoria) porque
// nadie les avisó que ya no hacen falta.

// procesarLote simula el procesamiento de un producto, respetando
// la cancelación del context (igual que consultarConTimeout en
// 09_biblioteca_estandar/05_contexto/02_with_timeout, pero ahora dentro de una goroutine del pool).
func procesarLote(ctx context.Context, id int, resultados chan<- string) {
	duracion := time.Duration(id) * 40 * time.Millisecond
	var mensaje string
	select {
	case <-time.After(duracion): // cada producto tarda algo distinto
		mensaje = fmt.Sprintf("producto %d procesado", id)
	case <-ctx.Done():
		mensaje = fmt.Sprintf("producto %d CANCELADO: %v", id, ctx.Err())
	}

	// El envío también respeta la cancelación: si nadie recibe y el canal se
	// bloquea, la goroutine no queda filtrada para siempre.
	select {
	case resultados <- mensaje:
	case <-ctx.Done():
	}
}

func main() {
	fmt.Println("=== Cancelar varias goroutines a la vez con un mismo context ===")

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Millisecond)
	defer cancel()

	resultados := make(chan string, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Go(func() { procesarLote(ctx, i, resultados) })
	}

	wg.Wait()
	close(resultados)

	for r := range resultados {
		fmt.Println(" -", r)
	}

	// ─────────────────────────────────────────────────────────
	// CANCELAR TODO EL TRABAJO RESTANTE ANTE UN ERROR
	// ─────────────────────────────────────────────────────────
	// Patrón real: si UNA goroutine encuentra un error grave,
	// cancela el context, y TODAS las demás (que comparten el
	// mismo ctx) se enteran y pueden frenar.

	fmt.Println("\n=== Una goroutine cancela el resto ante un error ===")

	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	resultados2 := make(chan string, 5)
	var wg2 sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg2.Go(func() {
			if i == 3 {
				// Este "worker" encuentra un problema grave y cancela
				// TODO el trabajo restante, no solo el suyo.
				time.Sleep(20 * time.Millisecond)
				resultados2 <- fmt.Sprintf("producto %d: ERROR FATAL, cancelando todo", i)
				cancel2()
				return
			}

			procesarLote(ctx2, i, resultados2)
		})
	}

	wg2.Wait()
	close(resultados2)

	for r := range resultados2 {
		fmt.Println(" -", r)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Un mismo ctx compartido  → cancela TODAS las goroutines que lo usan")
	fmt.Println("  WithTimeout              → límite de tiempo para todo el lote")
	fmt.Println("  WithCancel + cancel()    → una goroutine puede frenar a las demás")
	fmt.Println("  Sin esto                 → riesgo de 'goroutine leaks' (quedan colgadas)")
	fmt.Println("  Patrón real              → context.Context como primer parámetro, siempre")
}
