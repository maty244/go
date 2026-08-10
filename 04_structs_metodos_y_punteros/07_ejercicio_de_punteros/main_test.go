package main

import "testing"

func TestRetirarRechazaMontoNegativoSinAlterarSaldo(t *testing.T) {
	cuenta := &Cuenta{Saldo: 10_000, Activa: true}
	if err := retirar(cuenta, -100); err == nil {
		t.Fatal("retirar un monto negativo debería fallar")
	}
	if cuenta.Saldo != 10_000 {
		t.Fatalf("saldo alterado: %d", cuenta.Saldo)
	}
}
