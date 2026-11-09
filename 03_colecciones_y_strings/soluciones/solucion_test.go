package solucion

import "testing"

func TestClonar(t *testing.T) {
	original := []string{"Go"}
	copia := Clonar(original)
	copia[0] = "otro"
	if original[0] != "Go" || Clonar(nil) != nil {
		t.Fatal("clon incorrecto")
	}
	if copia := Clonar([]string{}); copia == nil || len(copia) != 0 {
		t.Fatal("el clon no preservó vacío no-nil")
	}
}
