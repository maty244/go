package solucion

import (
	"slices"
	"testing"
)

func TestFiltrar(t *testing.T) {
	resultado := slices.Collect(Filtrar(slices.Values([]int{1, 2}), func(n int) bool { return n > 1 }))
	if !slices.Equal(resultado, []int{2}) {
		t.Fatalf("resultado=%v", resultado)
	}
}
