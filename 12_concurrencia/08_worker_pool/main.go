package main

import (
	"fmt"
	"sync"
	"time"
)

// =========================================================
// WORKER POOL: UN NÚMERO FIJO DE GOROUTINES PROCESANDO TRABAJO
// =========================================================
// Lanzar una goroutine por tarea (como en temas anteriores) está
// bien para pocas tareas. Pero si tenés 100.000 tareas, lanzar
// 100.000 goroutines a la vez puede saturar recursos (conexiones
// a la base de datos, memoria, CPU). El patrón "worker pool"
// resuelve esto: un número FIJO de goroutines ("workers") que
// van tomando tareas de un channel compartido, de a una por vez.
//
// Piezas del patrón:
//   1. Un channel de "trabajos" (jobs) de entrada
//   2. N goroutines "worker" que reciben de ese channel y trabajan
//   3. Un channel de "resultados" de salida (opcional)
//   4. sync.WaitGroup para saber cuándo todos los workers terminaron

type Trabajo struct {
	ID       int
	Producto string
}

type Resultado struct {
	TrabajoID int
	Mensaje   string
}

// worker es la función que corre cada goroutine del pool: recibe
// trabajos del channel "trabajos" HASTA que se cierra, y manda
// cada resultado al channel "resultados".
func worker(id int, trabajos <-chan Trabajo, resultados chan<- Resultado) {
	for t := range trabajos { // termina solo cuando "trabajos" se cierra
		time.Sleep(30 * time.Millisecond) // simula procesamiento
		resultados <- Resultado{
			TrabajoID: t.ID,
			Mensaje:   fmt.Sprintf("worker %d procesó pedido de %s", id, t.Producto),
		}
	}
}

func main() {
	fmt.Println("=== Worker pool: 3 workers procesando 8 trabajos ===")

	const cantidadWorkers = 3

	trabajos := make(chan Trabajo, 8)
	resultados := make(chan Resultado, 8)

	var wg sync.WaitGroup

	// Lanzamos un número FIJO de workers, sin importar cuántos
	// trabajos haya (podrían ser 8 o 8 millones).
	for w := 1; w <= cantidadWorkers; w++ {
		wg.Go(func() { worker(w, trabajos, resultados) })
	}

	// Cargamos los trabajos en el channel
	productos := []string{"Alfajor", "Gaseosa", "Notebook", "Mouse", "Teclado", "Monitor", "Auriculares", "Cargador"}
	for i, producto := range productos {
		trabajos <- Trabajo{ID: i + 1, Producto: producto}
	}
	close(trabajos) // avisamos: no hay más trabajos (los workers van a terminar su range)

	// Cerramos "resultados" cuando TODOS los workers terminaron,
	// en una goroutine aparte (para no bloquear la lectura de abajo).
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Recibimos los resultados a medida que llegan (for range
	// termina solo cuando close(resultados) se ejecuta arriba).
	for r := range resultados {
		fmt.Printf("  [trabajo %d] %s\n", r.TrabajoID, r.Mensaje)
	}

	fmt.Println("Todos los trabajos fueron procesados")

	// ─────────────────────────────────────────────────────────
	// POR QUÉ NO SIMPLEMENTE "go worker() POR CADA TRABAJO"
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Por qué un número FIJO de workers ===")
	fmt.Println("  Con 3 workers, como máximo 3 trabajos se procesan concurrentemente")
	fmt.Println("  Esto limita el uso de recursos (conexiones DB, memoria, CPU)")
	fmt.Println("  Aunque haya 1000 trabajos en la cola, nunca se disparan 1000 goroutines")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  N workers fijos       → leen de un channel de trabajos compartido")
	fmt.Println("  close(trabajos)       → avisa a los workers que no hay más para tomar")
	fmt.Println("  wg.Wait() + goroutine → cierra 'resultados' cuando todos terminaron")
	fmt.Println("  for range resultados  → recibe todo hasta que se cierra")
	fmt.Println("  Uso típico            → procesar colas grandes con recursos limitados")
}
