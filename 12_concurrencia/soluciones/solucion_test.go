package solucion

import (
	"context"
	"slices"
	"testing"
)

func TestCuadrados(t *testing.T) {
	entrada := make(chan int, 3)
	entrada <- 2
	entrada <- 3
	close(entrada)
	var resultado []int
	for n := range Cuadrados(context.Background(), entrada) {
		resultado = append(resultado, n)
	}
	if !slices.Equal(resultado, []int{4, 9}) {
		t.Fatalf("resultado=%v", resultado)
	}
}

func TestCuadradosRespetaCancelacion(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	entrada := make(chan int, 2)
	entrada <- 3
	entrada <- 4 // el consumidor no va a leer este segundo valor

	salida := Cuadrados(ctx, entrada)
	if primero := <-salida; primero != 9 {
		t.Fatalf("primero=%d; quiero 9", primero)
	}

	// El consumidor decide cortar antes de leer el segundo valor:
	cancel()

	// Como Cuadrados respeta ctx.Done() tanto al recibir como al enviar,
	// debe salir de su loop y cerrar `salida`.
	for range salida {
	}
}
