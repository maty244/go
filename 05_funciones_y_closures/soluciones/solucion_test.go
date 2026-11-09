package solucion

import (
	"slices"
	"testing"
)

func TestAplicar(t *testing.T) {
	original := []int{1, 2}
	resultado := Aplicar(original, func(n int) int { return n + 1 }, func(n int) int { return n * 2 })
	if !slices.Equal(resultado, []int{4, 6}) || !slices.Equal(original, []int{1, 2}) {
		t.Fatalf("resultado=%v original=%v", resultado, original)
	}
}
