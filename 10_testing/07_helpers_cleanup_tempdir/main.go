package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GuardarTexto(directorio, nombre, contenido string) error {
	ruta := filepath.Join(directorio, nombre)
	if err := os.WriteFile(ruta, []byte(contenido), 0o600); err != nil {
		return fmt.Errorf("guardar %s: %w", nombre, err)
	}
	return nil
}

func main() {
	fmt.Println("Los recursos temporales se prueban con t.TempDir y t.Cleanup.")
	fmt.Println("Abrí main_test.go: el helper usa t.Helper para informar la línea útil.")
}
