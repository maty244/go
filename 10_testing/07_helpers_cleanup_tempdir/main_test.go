package main

import (
	"os"
	"path/filepath"
	"testing"
)

func leerArchivo(t *testing.T, ruta string) string {
	t.Helper() // si falla, testing señala a quien llamó al helper (línea 38), no acá
	datos, err := os.ReadFile(ruta)
	if err != nil {
		t.Fatalf("ReadFile(%q): %v", ruta, err)
	}
	return string(datos)
}

func TestGuardarTexto(t *testing.T) {
	directorio := t.TempDir()        // testing lo borra aunque el test falle
	t.Setenv("GOCITO_MODO", "test")  // restaura el entorno automáticamente
	ctx := t.Context()               // Go 1.24+: se cancela solo al terminar el test
	dirArtefactos := t.ArtifactDir() // Go 1.26: carpeta de artefactos para este test

	if ctx.Err() != nil {
		t.Fatal("el contexto del test debería estar activo mientras corre")
	}
	if dirArtefactos == "" {
		t.Fatal("ArtifactDir debería devolver una ruta válida")
	}

	limpiezaEjecutada := false
	t.Cleanup(func() { limpiezaEjecutada = true })
	// Cleanup corre al terminar el test. No afirmamos aquí porque todavía no
	// ocurrió; se usa para cerrar servidores, conexiones o restaurar estado.
	_ = limpiezaEjecutada

	if err := GuardarTexto(directorio, "saludo.txt", "hola"); err != nil {
		t.Fatal(err)
	}
	obtenido := leerArchivo(t, filepath.Join(directorio, "saludo.txt"))
	if obtenido != "hola" {
		t.Fatalf("contenido = %q; quiero hola", obtenido)
	}
}
