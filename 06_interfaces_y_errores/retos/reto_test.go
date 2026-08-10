//go:build reto

package reto

import (
	"errors"
	"strings"
	"testing"
)

func TestCargar(t *testing.T) {
	err := Cargar(42)
	if !errors.Is(err, ErrNoEncontrado) || !strings.Contains(err.Error(), "42") {
		t.Fatalf("error=%v", err)
	}
}
