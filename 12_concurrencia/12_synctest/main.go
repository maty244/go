package main

import (
	"context"
	"fmt"
	"time"
)

// =========================================================
// testing/synctest: TESTS DE CONCURRENCIA SIN ESPERAS REALES (Go 1.25+)
// =========================================================
// Testear código concurrente que usa time.Sleep, time.Timer o
// context.WithTimeout solía tener dos problemas:
//   1. Tests lentos (si probás un timeout de 5 segundos, el test tarda 5s).
//   2. Tests "flaky" (inestables: fallan a veces si la máquina está lenta).
//
// El paquete estándar "testing/synctest" (estabilizado en Go 1.25)
// resuelve ambos problemas ejecutando tu test dentro de una "burbuja"
// aislada con un RELOJ VIRTUAL:
//   - synctest.Test(t, func(t *testing.T) { ... }): crea la burbuja.
//     Cuando todas las goroutines adentro de la burbuja están
//     bloqueadas esperando tiempo o canales, el reloj virtual avanza
//     INSTANTÁNEAMENTE hasta el próximo evento (una hora pasa en microsegundos).
//   - synctest.Wait(): espera a que todas las demás goroutines de la
//     burbuja queden bloqueadas de forma durable antes de continuar.

func esperarOCancelar(ctx context.Context, demora time.Duration) error {
	temporizador := time.NewTimer(demora)
	defer temporizador.Stop()
	select {
	case <-temporizador.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	fmt.Println("=== testing/synctest (Go 1.25+) ===")
	fmt.Println("testing/synctest crea una burbuja con reloj falso para tests deterministas.")
	fmt.Println("Abrí main_test.go: prueba un timeout de una hora sin esperar una hora real.")
	fmt.Println("Corré: go test -v ./12_concurrencia/12_synctest")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  synctest.Test(t, fn) → ejecuta fn en una burbuja con reloj virtual")
	fmt.Println("  Reloj instantáneo    → avanza solo cuando las goroutines se bloquean")
	fmt.Println("  synctest.Wait()      → espera a que las goroutines de la burbuja se estabilicen")
}
