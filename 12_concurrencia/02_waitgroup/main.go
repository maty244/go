package main

import (
	"fmt"
	"sync"
	"time"
)

// =========================================================
// sync.WaitGroup: ESPERAR A QUE VARIAS GOROUTINES TERMINEN
// =========================================================
// En el tema 01 usamos time.Sleep como "parche" para esperar a
// las goroutines. sync.WaitGroup es la forma CORRECTA: un
// contador que sabe exactamente cuándo todo el trabajo terminó,
// sin adivinar cuánto tiempo esperar.
//
// Desde Go 1.25, el patrón preferido es WaitGroup.Go:
//
//   var wg sync.WaitGroup
//   wg.Go(func() { trabajo() })
//   wg.Wait()        // bloquea hasta que el contador llegue a 0
//
// En código anterior vas a encontrar Add + go + defer Done. Sigue siendo
// válido y se muestra abajo, pero es más fácil equivocarse al contar.

func procesarPedido(id int) {
	time.Sleep(50 * time.Millisecond) // simula trabajo
	fmt.Printf("  Pedido #%d procesado\n", id)
}

func main() {
	fmt.Println("=== sync.WaitGroup: esperar goroutines correctamente ===")

	var wg sync.WaitGroup

	pedidos := []int{101, 102, 103, 104}

	for _, id := range pedidos {
		wg.Go(func() { procesarPedido(id) })
	}

	fmt.Println("Lanzamos todos los pedidos, esperando a que terminen...")
	wg.Wait() // bloquea EXACTAMENTE hasta que las 4 llamen Done()
	fmt.Println("Todos los pedidos terminaron. Ahora sí seguimos con seguridad.")

	// ─────────────────────────────────────────────────────────
	// COMPATIBILIDAD: Add + Done
	// ─────────────────────────────────────────────────────────
	// WaitGroup no debe copiarse después de usarse (si lo pasás a una función,
	// pasalo siempre por puntero *sync.WaitGroup).
	// En la forma antigua, Add(1) SIEMPRE debe ocurrir ANTES de lanzar `go func()`
	// (si pusieras Add(1) adentro de la goroutine, wg.Wait() podría ejecutarse
	// antes de que la goroutine llegue a sumar 1). `wg.Go` hace ese Add(1) previo
	// y el Done() final automáticamente por vos.
	fmt.Println("\n=== Forma anterior: Add + Done ===")

	var wg2 sync.WaitGroup
	tareas := 5
	wg2.Add(tareas) // equivalente a llamar Add(1) cinco veces ANTES de lanzar las goroutines

	for i := 1; i <= tareas; i++ {
		go func(n int) {
			defer wg2.Done()
			fmt.Printf("  Tarea %d completada\n", n)
		}(i) // pasamos "i" como argumento: estilo explícito y claro,
		// aunque en Go 1.26 ya no haría falta para evitar el bug
		// clásico del for (ver 05_funciones_y_closures/10_trampa_closure_en_for)
	}

	wg2.Wait()
	fmt.Println("Las 5 tareas terminaron")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  wg.Go(f)          → inicia y cuenta la tarea automáticamente")
	fmt.Println("  f no debe panic   → requisito documentado de WaitGroup.Go")
	fmt.Println("  Add/Done          → forma válida que verás en código anterior")
	fmt.Println("  wg.Wait()         → bloquea hasta que el contador llegue a 0")
	fmt.Println("  No copiar         → un WaitGroup no se copia después de usarse")
	fmt.Println("  Reemplaza a Sleep  → espera EXACTA, no una adivinanza de tiempo")
}
