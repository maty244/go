package main

import "fmt"

// =========================================================
// CHANNELS nil VS CHANNELS CERRADOS: TABLA DEFINITIVA
// =========================================================
// Entender cómo responde un channel según su estado evita los dos
// bugs más comunes en concurrencia con Go (bloqueos infinitos y
// panics al cerrar/enviar):
//
//   Operación     | Channel nil       | Channel abierto       | Channel cerrado
//   ──────────────┼───────────────────┼───────────────────────┼─────────────────────────
//   Recibir <-ch  | Bloquea siempre*  | Recibe o bloquea      | Drena buffer; luego devuelve (zero, false) SIN bloquear
//   Enviar ch<-v  | Bloquea siempre*  | Envía o bloquea       | PANIC ("send on closed channel")
//   close(ch)     | PANIC             | Lo cierra             | PANIC ("close of closed channel")
//
//   * En un `select`, un caso sobre un channel `nil` nunca está
//     listo, por lo que queda DESHABILITADO (muy útil para "apagar"
//     una rama del select cuando un canal se cierra).
//
// REGLA DE PROPIEDAD DEL CIERRE:
// - Quien ENVÍA (el productor) es quien cierra el canal.
// - Si hay VARIOS productores escribiendo al mismo canal de salida,
//   ninguno debe cerrarlo por su cuenta (porque otro productor haría
//   panic al enviar): se los espera con un sync.WaitGroup y una sola
//   goroutine coordinadora hace `wg.Wait()` y luego `close(salida)`
//   (exactamente como hicimos en 08_worker_pool).

func main() {
	fmt.Println("=== Channels nil vs cerrados ===")

	// El zero value de un channel es nil. Enviar o recibir en él bloquea
	// indefinidamente. En un select, ese caso queda deshabilitado.
	var ausente chan int
	select {
	case <-ausente:
		fmt.Println("no se ejecuta")
	default:
		fmt.Println("channel nil: caso deshabilitado")
	}

	// close indica que no habrá más envíos; no elimina los valores del buffer.
	cola := make(chan int, 2)
	cola <- 10
	cola <- 20
	close(cola) // el productor termina; no puede volver a enviar ni cerrar
	for i := 0; i < 3; i++ {
		valor, abierto := <-cola
		fmt.Println("recibir:", valor, abierto) // 10 true; 20 true; 0 false
	}

	// Un channel cerrado y drenado sigue listo para recibir el zero value.
	// En un select repetido puede crear un loop que consume CPU. Convertir
	// la variable LOCAL a nil deshabilita el caso; no modifica el channel.
	cola = nil
	select {
	case <-cola:
		fmt.Println("no se ejecuta")
	default:
		fmt.Println("channel deshabilitado después del cierre")
	}

	// No ejecutamos operaciones que bloquearían o provocarían panic.
	// Experimento: con un channel cerrado, quitá la asignación cola = nil.
	// Predecí qué rama del último select se ejecuta y qué valor recibe.

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Enviar/recibir en nil   → bloquea para siempre (deshabilita el case en select)")
	fmt.Println("  Recibir de cerrado      → entrega lo que quede y luego (zero, false) al instante")
	fmt.Println("  Enviar/cerrar cerrado   → PANIC inmediato")
	fmt.Println("  ch = nil tras !abierto  → apaga ese canal dentro de un select en loop")
}
