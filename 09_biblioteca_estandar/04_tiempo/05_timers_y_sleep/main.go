package main

import (
	"fmt"
	"time"
)

// =========================================================
// time.Sleep Y TIMERS BÁSICOS
// =========================================================
// Hasta ahora vimos tiempo como DATO (fechas, duraciones). Este
// tema es sobre tiempo como ACCIÓN: pausar la ejecución, o
// programar algo para "dentro de X tiempo".

func main() {
	// ─────────────────────────────────────────────────────────
	// time.Sleep: PAUSAR LA EJECUCIÓN
	// ─────────────────────────────────────────────────────────
	// Bloquea el programa (o, más adelante en 12_concurrencia/, la
	// goroutine actual) durante la Duration indicada.

	fmt.Println("=== time.Sleep ===")
	fmt.Println("Esperando 200ms...")

	inicio := time.Now()
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Listo. Pasaron %.0fms\n", time.Since(inicio).Seconds()*1000)

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR Sleep (y cuándo NO)
	// ─────────────────────────────────────────────────────────
	// Sleep es útil para: reintentos con espera, simular demoras
	// en pruebas, throttling simple. NO lo uses para "esperar a
	// que algo termine" cuando ese algo corre en otra goroutine:
	// para eso existen WaitGroup y channels (12_concurrencia/, más
	// adelante), que son mucho más precisos que "dormir y confiar".

	// ─────────────────────────────────────────────────────────
	// time.After: UNA "ALARMA" QUE AVISA POR UN CANAL
	// ─────────────────────────────────────────────────────────
	// time.After(duracion) devuelve un CANAL (los vas a ver a
	// fondo en 12_concurrencia/) que recibe un valor automáticamente
	// cuando pasa ese tiempo. Por ahora, alcanza con saber leerlo
	// con "<-canal", que BLOQUEA hasta que llega el aviso.

	fmt.Println("\n=== time.After: esperar con un canal ===")
	fmt.Println("Esperando la señal de time.After(150ms)...")

	<-time.After(150 * time.Millisecond)
	fmt.Println("¡Llegó la señal!")

	// ─────────────────────────────────────────────────────────
	// time.Timer: COMO time.After, PERO SE PUEDE CANCELAR
	// ─────────────────────────────────────────────────────────
	// Un Timer es más flexible: podés detenerlo ANTES de que
	// dispare, con Stop(). Muy útil para timeouts que a veces se
	// resuelven antes de tiempo (por ejemplo: "esperá la respuesta
	// del servidor, pero si tarda más de 2 segundos, cancelá").

	fmt.Println("\n=== time.Timer: se puede cancelar ===")

	timer := time.NewTimer(300 * time.Millisecond)

	tareaTerminoAntes := true // simulamos que la tarea terminó rápido
	if tareaTerminoAntes {
		if timer.Stop() {
			fmt.Println("La tarea terminó a tiempo, cancelamos el timer")
		}
	} else {
		<-timer.C // esperar a que el timer dispare
		fmt.Println("Se acabó el tiempo")
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: reintentar con espera creciente (backoff simple)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: reintentos con espera ===")

	intentos := 3
	espera := 100 * time.Millisecond

	for i := 1; i <= intentos; i++ {
		fmt.Printf("Intento %d...\n", i)
		exito := i == intentos // simulamos que recién el último intento funciona

		if exito {
			fmt.Println("  ✓ Éxito")
			break
		}
		fmt.Printf("  ✗ Falló, esperando %v antes de reintentar\n", espera)
		time.Sleep(espera)
		espera *= 2 // cada reintento espera el doble
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  time.Sleep(d)       → pausa la ejecución esa duración")
	fmt.Println("  time.After(d)       → canal que avisa una sola vez, tras d")
	fmt.Println("  time.NewTimer(d)    → como After, pero se puede Stop() antes")
	fmt.Println("  Uso típico          → reintentos, timeouts, backoff")
	fmt.Println("  Con goroutines      → mejor WaitGroup/channels que Sleep (12_concurrencia/)")
}
