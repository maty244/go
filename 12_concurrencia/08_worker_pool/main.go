package main

import (
	"fmt"
	"sync"
	"time"
)

// =========================================================
// WORKER POOL, FAN-OUT, FAN-IN Y PIPELINES
// =========================================================
// Lanzar una goroutine por tarea (como en temas anteriores) está
// bien para pocas tareas. Pero si tenés 100.000 tareas, lanzar
// 100.000 goroutines a la vez puede saturar recursos (conexiones
// a la base de datos, memoria, CPU). El patrón "worker pool"
// resuelve esto: un número FIJO de goroutines ("workers") que
// van tomando tareas de un channel compartido, de a una por vez.
//
// Este ejemplo reúne los 3 patrones clásicos de concurrencia en Go:
//   - FAN-OUT (abrir en abanico): varias goroutines (los N workers)
//     leen de un mismo channel de entrada para repartirse el trabajo.
//   - FAN-IN (unificar en un embudo): las salidas de esas N goroutines
//     convergen en un único channel `resultados`, que un coordinador
//     cierra después de `wg.Wait()`.
//   - PIPELINE (tubería por etapas): una función que recibe un flujo
//     (slice o `<-chan A`), procesa en una goroutine y devuelve un
//     `<-chan B` listo para que la siguiente etapa lo consuma a medida
//     que sale (como `procesarTrabajos` acá abajo o `Cuadrados` en `retos/`).
//
// Piezas del patrón:
//   1. Un channel de "trabajos" (jobs) de entrada
//   2. N goroutines "worker" que reciben de ese channel y trabajan (Fan-Out)
//   3. Un channel de "resultados" de salida donde todas escriben (Fan-In)
//   4. sync.WaitGroup para saber cuándo todos los workers terminaron y cerrar la salida

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

// procesarTrabajos entrega resultados mientras se producen. El consumidor
// debe recorrer toda la salida; la versión cancelable se estudia en 09 y retos/.
// No modificamos entrada y el llamador no debe modificarla mientras se procesa.
func procesarTrabajos(entrada []Trabajo) <-chan Resultado {
	const cantidadWorkers = 3

	trabajos := make(chan Trabajo, cantidadWorkers)
	resultados := make(chan Resultado)

	var wg sync.WaitGroup

	// Lanzamos un número FIJO de workers, sin importar cuántos
	// trabajos haya (podrían ser 8 o 8 millones).
	for w := 1; w <= cantidadWorkers; w++ {
		wg.Go(func() { worker(w, trabajos, resultados) })
	}

	// Producir en otra goroutine permite al llamador consumir resultados YA.
	// Si enviáramos todo antes de retornar, una cola llena podría esperar a
	// workers bloqueados al enviar resultados: nadie llegaría a recibirlos.
	go func() {
		defer close(trabajos) // el productor es dueño del cierre de esta entrada
		for _, trabajo := range entrada {
			trabajos <- trabajo
		}
	}()

	// Cerramos "resultados" cuando TODOS los workers terminaron,
	// en una goroutine aparte (para no bloquear la lectura de abajo).
	go func() {
		wg.Wait()
		close(resultados)
	}()
	return resultados // <-chan permite recibir; el consumidor no puede enviar/cerrar
}

func main() {
	fmt.Println("=== Worker pool: 3 workers procesando 8 trabajos ===")
	productos := []string{"Alfajor", "Gaseosa", "Notebook", "Mouse", "Teclado", "Monitor", "Auriculares", "Cargador"}
	trabajos := make([]Trabajo, 0, len(productos))
	for i, producto := range productos {
		trabajos = append(trabajos, Trabajo{ID: i + 1, Producto: producto})
	}

	// Recibimos los resultados a medida que llegan (for range
	// termina solo cuando close(resultados) se ejecuta arriba).
	for r := range procesarTrabajos(trabajos) {
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
	fmt.Println("  Productor en goroutine → permite consumir resultados mientras se envían trabajos")
	fmt.Println("  close(trabajos)       → avisa a los workers que no hay más para tomar")
	fmt.Println("  wg.Wait() + goroutine → cierra 'resultados' cuando todos terminaron")
	fmt.Println("  for range resultados  → recibe todo hasta que se cierra")
	fmt.Println("  Uso típico            → procesar colas grandes con recursos limitados")
	fmt.Println("  Orden de resultados   → puede variar; identificá cada resultado por su TrabajoID")
}
