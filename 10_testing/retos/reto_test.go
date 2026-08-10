//go:build reto

package reto

import "testing"

func TestEsPalindromo(t *testing.T) {
	for _, texto := range []string{"Anita lava la tina", "Añá"} {
		if !EsPalindromo(texto) {
			t.Errorf("%q debería ser palíndromo", texto)
		}
	}
	if EsPalindromo("Golang") {
		t.Fatal("Golang no es palíndromo")
	}
}
