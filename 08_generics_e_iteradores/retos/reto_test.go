//go:build reto

package reto

import (
	"slices"
	"testing"
)

func TestFiltrar(t *testing.T) {
	resultado := slices.Collect(Filtrar(slices.Values([]int{1, 2, 3, 4}), func(n int) bool { return n%2 == 0 }))
	if !slices.Equal(resultado, []int{2, 4}) {
		t.Fatalf("resultado=%v", resultado)
	}
}
