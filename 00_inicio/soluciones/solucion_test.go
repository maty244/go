package solucion

import "testing"

func TestSaludo(t *testing.T) {
	if Saludo("Mati") != "Hola, Mati" {
		t.Fatal("saludo incorrecto")
	}
}
