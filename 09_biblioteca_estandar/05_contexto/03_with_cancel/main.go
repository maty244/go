package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// =========================================================
// context.WithCancel: CANCELAR MANUALMENTE, CUANDO VOS DECIDAS
// =========================================================
// A diferencia de WithTimeout (tema 02), acá NO hay un tiempo
// límite automático: vos controlás CUÁNDO se cancela, llamando a
// la función cancel() que te devuelve.
//
//   ctx, cancel := context.WithCancel(padre)
//   defer cancel() // siempre, aunque canceles antes a mano
//
// Es útil cuando la razón para parar no es "pasó tiempo", sino
// un EVENTO: el usuario apretó "cancelar", ya encontraste lo que
// buscabas y no hace falta seguir, otra parte del programa falló.

func main() {
	fmt.Println("=== context.WithCancel: cancelar manualmente ===")

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("¿Ya se canceló?", ctx.Err())

	cancel() // cancelamos ahora, a propósito

	fmt.Println("Después de llamar cancel():", ctx.Err())
	fmt.Println("¿Es context.Canceled?", errors.Is(ctx.Err(), context.Canceled))

	// ─────────────────────────────────────────────────────────
	// CASO REAL: cancelar el resto de un trabajo apenas algo falla
	// ─────────────────────────────────────────────────────────
	// Ejemplo: estás procesando varios pasos en cadena, y si el
	// paso 2 falla, no tiene sentido seguir con el 3, el 4, etc.
	// El context cancelado se propaga a TODO lo que lo esté usando.

	fmt.Println("\n=== Caso real: frenar el resto del trabajo ante un fallo ===")

	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	pasos := []string{"validar", "cobrar", "actualizar_stock", "notificar"}

	for i, paso := range pasos {
		if ctx2.Err() != nil {
			fmt.Printf("  Paso %d (%s): SALTEADO, el context ya se canceló\n", i+1, paso)
			continue
		}

		fmt.Printf("  Paso %d (%s): ejecutando...\n", i+1, paso)

		if paso == "cobrar" {
			fmt.Println("    → el pago fue rechazado, cancelamos lo que sigue")
			cancel2() // cancelamos: no tiene sentido actualizar stock ni notificar
		}
	}

	// ─────────────────────────────────────────────────────────
	// COMBINANDO WithCancel Y select (como en el tema 02)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Esperar cancelación con select ===")

	ctx3, cancel3 := context.WithCancel(context.Background())

	// Simulamos que "algo" cancela el contexto en 80ms, en paralelo
	// (esto usa una goroutine, que vas a estudiar a fondo en
	// 12_concurrencia/; por ahora, solo mirá el resultado).
	go func() {
		time.Sleep(80 * time.Millisecond)
		fmt.Println("  (evento externo) cancelando ahora...")
		cancel3()
	}()

	select {
	case <-ctx3.Done():
		fmt.Println("El trabajo se detuvo:", ctx3.Err())
	case <-time.After(1 * time.Second):
		fmt.Println("Esto no debería pasar en este ejemplo")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  context.WithCancel(padre)  → vos decidís CUÁNDO cancelar")
	fmt.Println("  cancel()                   → dispara la cancelación")
	fmt.Println("  ctx.Err() == context.Canceled → se canceló a mano (no por timeout)")
	fmt.Println("  Uso típico                  → frenar trabajo restante ante un fallo/evento")
	fmt.Println("  defer cancel()              → siempre, para liberar recursos internos")
}
