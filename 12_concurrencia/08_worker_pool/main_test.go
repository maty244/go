package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Más trabajos que los buffers detectan el bloqueo que un ejemplo de ocho
// trabajos no revela. No exigimos orden: las goroutines pueden intercalarse.
func TestPoolProcesaYCierra(t *testing.T) {
	for _, cantidad := range []int{0, 1, 20} {
		t.Run(fmt.Sprint(cantidad), func(t *testing.T) {
			entrada := make([]Trabajo, cantidad)
			for i := range entrada {
				entrada[i] = Trabajo{ID: i + 1, Producto: "Go"}
			}
			ctx, cancelar := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancelar()
			salida := procesarTrabajos(entrada)
			vistos := make(map[int]bool)
			for {
				select {
				case <-ctx.Done():
					t.Fatal("el pool no terminó: posible bloqueo")
				case resultado, abierto := <-salida:
					if !abierto {
						if len(vistos) != cantidad {
							t.Fatalf("recibidos=%d; quiero=%d", len(vistos), cantidad)
						}
						return
					}
					if resultado.TrabajoID < 1 || resultado.TrabajoID > cantidad || vistos[resultado.TrabajoID] {
						t.Fatalf("ID inesperado o duplicado: %d", resultado.TrabajoID)
					}
					vistos[resultado.TrabajoID] = true
				}
			}
		})
	}
}
