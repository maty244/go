package almacenjson

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"gocito/11_proyecto_cli/internal/tareas"
)

func TestGuardarYCargar(t *testing.T) {
	t.Parallel()

	ruta := filepath.Join(t.TempDir(), "datos", "tareas.json")
	almacen := Almacen{Ruta: ruta}
	estado := tareas.Estado{
		SiguienteID: 2,
		Tareas: []tareas.Tarea{{
			ID: 1, Titulo: "Aprender Go", CreadaEn: time.Date(2026, 8, 9, 0, 0, 0, 0, time.UTC),
		}},
	}

	if err := almacen.Guardar(estado); err != nil {
		t.Fatalf("Guardar: %v", err)
	}
	recuperado, err := almacen.Cargar()
	if err != nil {
		t.Fatalf("Cargar: %v", err)
	}
	if len(recuperado.Tareas) != 1 || recuperado.Tareas[0].Titulo != "Aprender Go" {
		t.Fatalf("estado recuperado inesperado: %+v", recuperado)
	}

	info, err := os.Stat(ruta)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("permisos = %o; quiero 600", info.Mode().Perm())
	}
	temporales, err := filepath.Glob(filepath.Join(filepath.Dir(ruta), ".tareas-*.tmp"))
	if err != nil {
		t.Fatal(err)
	}
	if len(temporales) != 0 {
		t.Fatalf("quedaron temporales: %v", temporales)
	}
}

func TestCargarArchivoInexistente(t *testing.T) {
	t.Parallel()

	almacen := Almacen{Ruta: filepath.Join(t.TempDir(), "no-existe.json")}
	estado, err := almacen.Cargar()
	if err != nil {
		t.Fatalf("Cargar: %v", err)
	}
	if len(estado.Tareas) != 0 {
		t.Fatalf("estado inicial inesperado: %+v", estado)
	}
}
