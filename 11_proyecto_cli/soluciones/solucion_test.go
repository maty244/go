package solucion

import (
	"testing"

	"gocito/11_proyecto_cli/internal/tareas"
)

func TestBuscar(t *testing.T) {
	resultado := Buscar([]tareas.Tarea{{Titulo: "Go"}}, "go")
	if len(resultado) != 1 {
		t.Fatal("búsqueda incorrecta")
	}
}
