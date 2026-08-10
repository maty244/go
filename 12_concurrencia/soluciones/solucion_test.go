package solucion

import (
	"context"
	"testing"
)

func TestCuadrados(t *testing.T) {
	entrada := make(chan int, 1)
	entrada <- 4
	close(entrada)
	if obtenido := <-Cuadrados(context.Background(), entrada); obtenido != 16 {
		t.Fatalf("obtenido=%d", obtenido)
	}
}
