package main

import "testing"

func TestSliceNilVacioYAliasing(t *testing.T) {
	var nilSlice []int
	vacio := []int{}
	if nilSlice != nil || vacio == nil || len(nilSlice) != len(vacio) {
		t.Fatalf("nil=%v vacío=%v", nilSlice, vacio)
	}

	original := []int{1, 2}
	alias := original
	alias[0] = 9
	if original[0] != 9 {
		t.Fatal("copiar un slice debería compartir el array subyacente")
	}

	copia := append([]int(nil), original...)
	copia[0] = 7
	if original[0] == 7 {
		t.Fatal("la copia independiente mantuvo aliasing")
	}
}
