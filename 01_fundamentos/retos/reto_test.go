//go:build reto

package reto

import (
	"errors"
	"testing"
)

func TestCalcularTotal(t *testing.T) {
	if total, err := CalcularTotal(250, 3); err != nil || total != 750 {
		t.Fatalf("total=%d err=%v", total, err)
	}
	if _, err := CalcularTotal(250, -1); !errors.Is(err, ErrValorNegativo) {
		t.Fatalf("error cantidad negativa=%v", err)
	}
	if _, err := CalcularTotal(-250, 3); !errors.Is(err, ErrValorNegativo) {
		t.Fatalf("error precio negativo=%v", err)
	}
}
