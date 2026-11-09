package main

import "testing"

func TestClonarPreservaNilYVacio(t *testing.T) {
	if Clonar[int](nil) != nil {
		t.Fatal("nil no se preservó")
	}
	if copia := Clonar([]string{}); copia == nil || len(copia) != 0 {
		t.Fatal("vacío no-nil no se preservó")
	}
}

func TestClonarSeparaArrayExterior(t *testing.T) {
	original := []int{1}
	copia := Clonar(original)
	copia[0] = 9
	if original[0] != 1 {
		t.Fatal("el array sigue compartido")
	}
	anidado := [][]int{{1}}
	clon := Clonar(anidado)
	clon[0][0] = 7
	if anidado[0][0] != 7 {
		t.Fatal("el clon debería copiar superficialmente sus elementos")
	}
	clon[0] = []int{8}
	if anidado[0][0] != 7 {
		t.Fatal("reemplazar un elemento del clon cambió el array exterior original")
	}
}
