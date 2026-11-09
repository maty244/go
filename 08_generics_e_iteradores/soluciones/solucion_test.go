package solucion

import (
	"slices"
	"testing"
)

func TestFiltrar(t *testing.T) {
	resultado := slices.Collect(Filtrar(slices.Values([]int{1, 2, 3, 4}), func(n int) bool { return n%2 == 0 }))
	if !slices.Equal(resultado, []int{2, 4}) {
		t.Fatalf("resultado=%v", resultado)
	}

	evaluados := 0
	seq := Filtrar(slices.Values([]int{2, 4, 6}), func(n int) bool {
		evaluados++
		return n%2 == 0
	})
	for range seq {
		break
	}
	if evaluados != 1 {
		t.Fatalf("no respetó el break de yield: evaluados=%d, quería 1", evaluados)
	}
}
