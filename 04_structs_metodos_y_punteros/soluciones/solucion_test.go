package solucion

import "testing"

func TestRetirar(t *testing.T) {
	cuenta := Cuenta{SaldoCentavos: 100}
	if err := cuenta.Retirar(40); err != nil || cuenta.SaldoCentavos != 60 {
		t.Fatalf("cuenta=%+v err=%v", cuenta, err)
	}
}
