package solucion

import (
	"errors"
	"testing"
)

func TestCargar(t *testing.T) {
	if err := Cargar(1); !errors.Is(err, ErrNoEncontrado) {
		t.Fatalf("error=%v", err)
	}
}
