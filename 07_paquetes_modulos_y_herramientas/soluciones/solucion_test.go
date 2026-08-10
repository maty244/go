package solucion

import "testing"

func TestRutaImport(t *testing.T) {
	ruta, err := RutaImport("gocito/", "/catalogo")
	if err != nil || ruta != "gocito/catalogo" {
		t.Fatalf("ruta=%q err=%v", ruta, err)
	}
}
