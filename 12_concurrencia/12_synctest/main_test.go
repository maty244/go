package main

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

func TestCancelacionDeterminista(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		inicio := time.Now()
		ctx, cancel := context.WithTimeout(t.Context(), time.Hour)
		defer cancel()

		// Aunque los plazos son de 1 hora y 2 horas, dentro de synctest.Test
		// el reloj virtual salta 1 hora en microsegundos apenas esperarOCancelar
		// se bloquea en el select.
		err := esperarOCancelar(ctx, 2*time.Hour)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("error = %v; quiero DeadlineExceeded", err)
		}
		if transcurrido := time.Since(inicio); transcurrido != time.Hour {
			t.Fatalf("tiempo falso = %v; quiero 1h", transcurrido)
		}

		// synctest.Wait() espera a que cualquier goroutine lanzada en la
		// burbuja llegue a un punto de bloqueo estable, sin usar time.Sleep:
		listo := false
		go func() {
			listo = true
		}()
		synctest.Wait()
		if !listo {
			t.Fatal("la goroutine debería haber ejecutado antes de que Wait retorne")
		}
	})
}
