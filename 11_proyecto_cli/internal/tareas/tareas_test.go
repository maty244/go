package tareas

import (
	"errors"
	"testing"
	"time"
)

func TestGestorFlujoCompleto(t *testing.T) {
	t.Parallel()

	ahora := time.Date(2026, time.August, 9, 10, 0, 0, 0, time.UTC)
	g := Nuevo(Estado{})

	creada, err := g.Agregar("  Estudiar slices  ", ahora)
	if err != nil {
		t.Fatalf("Agregar: %v", err)
	}
	if creada.ID != 1 || creada.Titulo != "Estudiar slices" {
		t.Fatalf("tarea inesperada: %+v", creada)
	}

	completada, err := g.Completar(creada.ID, ahora.Add(time.Hour))
	if err != nil {
		t.Fatalf("Completar: %v", err)
	}
	if !completada.Completada || completada.CompletadaEn.IsZero() {
		t.Fatalf("la tarea no quedó completa: %+v", completada)
	}

	// La copia evita que un consumidor modifique el slice interno.
	lista := g.Listar()
	lista[0].Titulo = "mutado desde afuera"
	if g.Listar()[0].Titulo != "Estudiar slices" {
		t.Fatal("Listar expuso el estado interno")
	}
}

func TestGestorErrores(t *testing.T) {
	t.Parallel()

	g := Nuevo(Estado{})
	if _, err := g.Agregar("  ", time.Now()); !errors.Is(err, ErrTituloVacio) {
		t.Fatalf("Agregar vacío: %v", err)
	}
	if _, err := g.Completar(99, time.Now()); !errors.Is(err, ErrNoEncontrada) {
		t.Fatalf("Completar inexistente: %v", err)
	}
}
