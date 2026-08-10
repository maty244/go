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

		err := esperarOCancelar(ctx, 2*time.Hour)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("error = %v; quiero DeadlineExceeded", err)
		}
		if transcurrido := time.Since(inicio); transcurrido != time.Hour {
			t.Fatalf("tiempo falso = %v; quiero 1h", transcurrido)
		}
	})
}
