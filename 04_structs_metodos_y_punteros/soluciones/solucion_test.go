package solucion

import (
	"errors"
	"testing"
)

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{SaldoCentavos: 100}
	if err := cuenta.Retirar(40); err != nil || cuenta.SaldoCentavos != 60 {
		t.Fatalf("cuenta=%+v err=%v", cuenta, err)
	}
	if err := cuenta.Retirar(-1); !errors.Is(err, ErrRetiroInvalido) || cuenta.SaldoCentavos != 60 {
		t.Fatalf("retiro negativo=%v saldo=%d", err, cuenta.SaldoCentavos)
	}
	if err := cuenta.Retirar(200); !errors.Is(err, ErrRetiroInvalido) || cuenta.SaldoCentavos != 60 {
		t.Fatalf("sobregiro=%v saldo=%d", err, cuenta.SaldoCentavos)
	}
}
