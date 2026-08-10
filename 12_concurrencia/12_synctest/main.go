package main

import (
	"context"
	"fmt"
	"time"
)

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
	fmt.Println("testing/synctest crea una burbuja con reloj falso para tests deterministas.")
	fmt.Println("Abrí main_test.go: prueba un timeout de una hora sin esperar una hora real.")
}
