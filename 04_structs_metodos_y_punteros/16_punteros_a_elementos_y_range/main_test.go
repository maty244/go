package main

import (
	"slices"
	"testing"
)

func TestRangeCopiaVsElementoReal(t *testing.T) {
	original := []int{10, 20, 30}
	copias := direccionesDeCopias(original)
	*copias[0] = 99
	if !slices.Equal(original, []int{10, 20, 30}) || *copias[1] != 20 || copias[0] == copias[1] {
		t.Fatal("las variables declaradas por range deben ser copias independientes de ints")
	}
	elementos := direccionesDeElementos(original)
	*elementos[0] = 77
	if original[0] != 77 || elementos[0] != &original[0] {
		t.Fatal("el puntero por índice debe identificar al elemento real")
	}
}

func TestRangeConAsignacionComparteUnaVariable(t *testing.T) {
	punteros := direccionesDeVariableReutilizada([]int{10, 20, 30})
	for _, p := range punteros {
		if p != punteros[0] || *p != 30 {
			t.Fatal("con = todos deben identificar la misma variable exterior")
		}
	}
	*punteros[0] = 88
	if *punteros[2] != 88 {
		t.Fatal("la mutación del destino compartido debe verse desde cualquier alias")
	}
}
