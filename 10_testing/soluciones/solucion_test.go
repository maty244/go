package solucion

import "testing"

func TestEsPalindromo(t *testing.T) {
	if !EsPalindromo("Anita lava la tina") || EsPalindromo("Go") {
		t.Fatal("resultado incorrecto")
	}
}
