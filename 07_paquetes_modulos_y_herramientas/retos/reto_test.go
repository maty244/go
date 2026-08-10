//go:build reto

package reto

import "testing"

func TestRutaImport(t *testing.T) {
	ruta, err := RutaImport("gocito/", "/catalogo")
	if err != nil || ruta != "gocito/catalogo" {
		t.Fatalf("ruta=%q err=%v", ruta, err)
	}
	if _, err := RutaImport("", "catalogo"); err == nil {
		t.Fatal("faltó rechazar módulo vacío")
	}
}
