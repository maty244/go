package main

import "testing"

// =========================================================
// TESTS INCOMPLETOS A PROPÓSITO
// =========================================================
// Testeamos solo 2 de las 4 ramas de AplicarDescuento (VIP y
// default). Corré:
//
//   go test -cover ./10_testing/04_coverage/
//
// y vas a ver que el coverage NO es 100%: las ramas "frecuente"
// y "empleado" quedan sin ejercitar por ningún test. Ese es
// justamente el valor de -cover: te muestra el hueco.

func TestAplicarDescuentoVIP(t *testing.T) {
	resultado := AplicarDescuento(1000, "VIP")
	esperado := Dinero(800)

	if resultado != esperado {
		t.Errorf("AplicarDescuento(1000, VIP) = %d; esperado %d", resultado, esperado)
	}
}

func TestAplicarDescuentoSinDescuento(t *testing.T) {
	resultado := AplicarDescuento(1000, "regular")
	esperado := Dinero(1000)

	if resultado != esperado {
		t.Errorf("AplicarDescuento(1000, regular) = %d; esperado %d", resultado, esperado)
	}
}

// ─────────────────────────────────────────────────────────
// PARA LLEGAR A 100% EN AplicarDescuento, DESCOMENTÁ ESTO
// ─────────────────────────────────────────────────────────
// func TestAplicarDescuentoFrecuente(t *testing.T) {
// 	resultado := AplicarDescuento(1000, "frecuente")
// 	esperado := Dinero(900)
// 	if resultado != esperado {
// 		t.Errorf("AplicarDescuento(1000, frecuente) = %d; esperado %d", resultado, esperado)
// 	}
// }
//
// func TestAplicarDescuentoEmpleado(t *testing.T) {
// 	resultado := AplicarDescuento(1000, "empleado")
// 	esperado := Dinero(500)
// 	if resultado != esperado {
// 		t.Errorf("AplicarDescuento(1000, empleado) = %d; esperado %d", resultado, esperado)
// 	}
// }
