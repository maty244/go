package solucion

import "testing"

func TestAplicar(t *testing.T) {
	resultado := Aplicar([]int{2}, func(n int) int { return n * n })
	if len(resultado) != 1 || resultado[0] != 4 {
		t.Fatalf("resultado=%v", resultado)
	}
}
