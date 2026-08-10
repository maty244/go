package solucion

import "testing"

func TestClonar(t *testing.T) {
	original := []int{1}
	copia := Clonar(original)
	copia[0] = 2
	if original[0] != 1 || Clonar[int](nil) != nil {
		t.Fatal("clon incorrecto")
	}
}
