//go:build reto

package reto

import "testing"

func TestSaludo(t *testing.T) {
	if obtenido := Saludo("Mati"); obtenido != "Hola, Mati" {
		t.Fatalf("Saludo = %q", obtenido)
	}
}
