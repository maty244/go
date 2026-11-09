package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// =========================================================
// HELPERS, CLEANUP, TEMPDIR Y ENTORNO DE TEST
// =========================================================
// Cuando un test necesita interactuar con archivos, variables de
// entorno o recursos que hay que limpiar al terminar, *testing.T
// trae herramientas pensadas para que ningún test ensucie a otro:
//
//   t.Helper()      → marca una función auxiliar para que, si falla,
//                     Go reporte la línea del test que la llamó (y no
//                     la línea interna del helper).
//   t.TempDir()     → crea una carpeta temporal única para el test y la
//                     borra sola al terminar (incluso si el test falla).
//   t.Setenv(k, v)  → define una variable de entorno durante el test y
//                     restaura su valor original automáticamente al final.
//   t.Cleanup(fn)   → registra una función de limpieza que corre al
//                     terminar el test (como defer, pero funciona aunque
//                     lo registres adentro de un helper).
//   t.Context()     → (Go 1.24+) devuelve un context.Context que se
//                     cancela automáticamente justo antes de los Cleanup.
//   t.ArtifactDir() → (Go 1.26) devuelve un directorio para guardar
//                     artefactos de salida del test (si pasás -artifacts,
//                     los conserva; si no, actúa como temporal).

func GuardarTexto(directorio, nombre, contenido string) error {
	ruta := filepath.Join(directorio, nombre)
	if err := os.WriteFile(ruta, []byte(contenido), 0o600); err != nil {
		return fmt.Errorf("guardar %s: %w", nombre, err)
	}
	return nil
}

func main() {
	fmt.Println("=== Helpers, TempDir, Cleanup y Context en testing ===")
	fmt.Println("Los recursos temporales se prueban con t.TempDir y t.Cleanup.")
	fmt.Println("Abrí main_test.go: el helper usa t.Helper para informar la línea útil.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  t.Helper()        → reporta la línea del llamador cuando falla el helper")
	fmt.Println("  t.TempDir()       → carpeta temporal que se borra sola al terminar")
	fmt.Println("  t.Setenv(k, v)    → cambia env var y la restaura al final del test")
	fmt.Println("  t.Cleanup(func()) → limpieza garantizada al finalizar el test")
	fmt.Println("  t.Context()       → (Go 1.24+) context ligado a la vida del test")
	fmt.Println("  t.ArtifactDir()   → (Go 1.26) carpeta para artefactos de test (-artifacts)")
}
