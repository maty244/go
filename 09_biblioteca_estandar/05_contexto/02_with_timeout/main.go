package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// =========================================================
// context.WithTimeout: CANCELAR AUTOMÁTICAMENTE DESPUÉS DE X TIEMPO
// =========================================================
// El caso de uso más común de context: "hacé esto, pero si tarda
// más de N segundos, avisame que ya no hace falta seguir".
//
//   ctx, cancel := context.WithTimeout(padre, duracion)
//   defer cancel() // SIEMPRE, para liberar recursos internos
//
// ctx.Done() devuelve un canal (como time.After, visto en
// 09_biblioteca_estandar/04_tiempo/05_timers_y_sleep) que se cierra cuando: (a) se cumple el timeout, o
// (b) alguien llama cancel() manualmente. Leer de ese canal
// (<-ctx.Done()) te avisa "andá terminando".
//
// DOS PREGUNTAS CLAVE:
// 1. ¿Cancelar un context mata automáticamente una goroutine?
//    NO. La cancelación en Go es COOPERATIVA: el context solo cierra
//    su canal Done() y setea Err(). El código que hace el trabajo
//    debe consultar ctx.Done() / ctx.Err() (o pasarlo a funciones de
//    red/DB que lo consulten) para frenar ordenadamente.
// 2. ¿Por qué llamamos a defer cancel() incluso si el timeout vence solo?
//    Porque lo habitual es que la operación termine ANTES del timeout;
//    llamar a cancel() libera de inmediato el timer interno del context
//    en lugar de dejarlo vivo en memoria hasta que venza el plazo.

func main() {
	fmt.Println("=== context.WithTimeout ===")

	// Creamos un context que se cancela solo a los 100ms
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // buena práctica: SIEMPRE, incluso si el timeout ya pasó

	fmt.Println("Esperando la cancelación del context...")
	<-ctx.Done() // se bloquea hasta que se cumple el timeout

	fmt.Println("El context se canceló. Motivo:", ctx.Err())

	// ─────────────────────────────────────────────────────────
	// ctx.Err() DICE POR QUÉ SE CANCELÓ
	// ─────────────────────────────────────────────────────────
	// context.DeadlineExceeded → se venció el timeout
	// context.Canceled         → alguien llamó cancel() a mano (tema 03)

	fmt.Println("\n=== Distinguir el motivo de cancelación ===")
	fmt.Println("¿Es DeadlineExceeded?", errors.Is(ctx.Err(), context.DeadlineExceeded))

	// ─────────────────────────────────────────────────────────
	// SIMULAR UNA OPERACIÓN "LENTA" QUE RESPETA EL TIMEOUT
	// ─────────────────────────────────────────────────────────
	// El patrón real: una función hace ALGO (una consulta, un
	// cálculo) mientras VIGILA si el context se canceló, para
	// frenar antes en vez de seguir trabajando en vano.

	fmt.Println("\n=== Operación que respeta el timeout ===")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	resultado, err := consultarConTimeout(ctx2, 500*time.Millisecond) // tarda MÁS que el timeout
	if err != nil {
		fmt.Println("Falló:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// ─────────────────────────────────────────────────────────
	// CUANDO LA OPERACIÓN TERMINA A TIEMPO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Operación que SÍ llega a tiempo ===")

	ctx3, cancel3 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel3()

	resultado, err = consultarConTimeout(ctx3, 50*time.Millisecond) // más rápida que el timeout
	if err != nil {
		fmt.Println("Falló:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  context.WithTimeout(padre, d)  → ctx que se cancela solo tras d")
	fmt.Println("  defer cancel()                  → SIEMPRE, libera recursos internos")
	fmt.Println("  <-ctx.Done()                    → bloquea hasta cancelación/timeout")
	fmt.Println("  ctx.Err()                       → DeadlineExceeded o Canceled")
	fmt.Println("  Patrón típico                    → select entre el trabajo y ctx.Done()")
}

// consultarConTimeout simula una operación que tarda "duracionTrabajo",
// pero respeta el timeout del context: si el context se cancela
// antes de que el trabajo termine, retorna un error en vez de
// esperar indefinidamente.
//
// select { case <-a: ...; case <-b: ... } espera a que CUALQUIERA
// de los canales esté listo, y ejecuta ese caso. Lo vas a ver a
// fondo en 12_concurrencia/; por ahora alcanza con leerlo como
// "lo que pase primero, gana".
func consultarConTimeout(ctx context.Context, duracionTrabajo time.Duration) (string, error) {
	select {
	case <-time.After(duracionTrabajo): // el "trabajo" termina
		return "datos obtenidos correctamente", nil
	case <-ctx.Done(): // el timeout se cumplió antes
		return "", fmt.Errorf("consultarConTimeout: %w", ctx.Err())
	}
}
