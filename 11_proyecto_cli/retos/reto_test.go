//go:build reto

package reto

import (
	"testing"

	"gocito/11_proyecto_cli/internal/tareas"
)

func TestBuscar(t *testing.T) {
	lista := []tareas.Tarea{{ID: 1, Titulo: "Estudiar Go"}, {ID: 2, Titulo: "Dormir"}}
	resultado := Buscar(lista, "GO")
	if len(resultado) != 1 || resultado[0].ID != 1 {
		t.Fatalf("resultado=%v", resultado)
	}
}
