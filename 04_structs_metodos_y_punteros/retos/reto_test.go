//go:build reto

package reto

import (
	"errors"
	"testing"
)

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{SaldoCentavos: 1_000}
	if err := cuenta.Retirar(250); err != nil || cuenta.SaldoCentavos != 750 {
		t.Fatalf("cuenta=%+v err=%v", cuenta, err)
	}
	if err := cuenta.Retirar(-1); !errors.Is(err, ErrRetiroInvalido) {
		t.Fatalf("retiro negativo=%v", err)
	}
}
