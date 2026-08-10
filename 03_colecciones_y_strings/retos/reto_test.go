//go:build reto

package reto

import "testing"

func TestClonar(t *testing.T) {
	if Clonar[int](nil) != nil {
		t.Fatal("nil dejó de ser nil")
	}
	original := []string{"ñ", "Go"}
	copia := Clonar(original)
	copia[0] = "otro"
	if original[0] != "ñ" {
		t.Fatal("la copia mantiene aliasing")
	}
}
