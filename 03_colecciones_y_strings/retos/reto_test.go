//go:build reto

package reto

import "testing"

func TestClonar(t *testing.T) {
	if Clonar(nil) != nil {
		t.Fatal("nil dejó de ser nil")
	}
	if copia := Clonar([]string{}); copia == nil || len(copia) != 0 {
		t.Fatal("un slice vacío no-nil dejó de ser vacío no-nil")
	}
	original := []string{"ñ", "Go"}
	copia := Clonar(original)
	copia[0] = "otro"
	if original[0] != "ñ" {
		t.Fatal("la copia mantiene aliasing")
	}
}
