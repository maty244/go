//go:build reto

package reto

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
