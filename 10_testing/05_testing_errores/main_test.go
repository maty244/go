package main

import (
	"errors"
	"testing"
)

// =========================================================
// CASO FELIZ: el valor es correcto Y no hay error
// =========================================================

func TestDividirCasoFeliz(t *testing.T) {
	resultado, err := Dividir(10, 2)

	if err != nil {
		t.Fatalf("no se esperaba error, se obtuvo: %v", err)
	}
	if resultado != 5 {
		t.Errorf("Dividir(10, 2) = %.2f; esperado 5", resultado)
	}
}

// =========================================================
// CASO DE ERROR: usamos errors.Is, igual que en código real
// =========================================================

func TestDividirPorCero(t *testing.T) {
	_, err := Dividir(10, 0)

	if err == nil {
		t.Fatal("se esperaba un error, pero err es nil")
	}
	if !errors.Is(err, ErrDivisionPorCero) {
		t.Errorf("se esperaba ErrDivisionPorCero, se obtuvo: %v", err)
	}
}

// ─────────────────────────────────────────────────────────
// TESTEAR errors.Is CON WRAPPING (ver 06_interfaces_y_errores/14_wrapping_con_porcentaje_w)
// ─────────────────────────────────────────────────────────

func TestValidarMontoNegativo(t *testing.T) {
	err := ValidarMonto(-100)

	if !errors.Is(err, ErrMontoNegativo) {
		t.Errorf("se esperaba ErrMontoNegativo (posiblemente envuelto), se obtuvo: %v", err)
	}
}

func TestValidarMontoValido(t *testing.T) {
	err := ValidarMonto(100)

	if err != nil {
		t.Errorf("no se esperaba error para un monto válido, se obtuvo: %v", err)
	}
}

// =========================================================
// TESTEAR UN PANIC ESPERADO: recover() dentro del test
// =========================================================
// El patrón: usamos defer + recover() para "atrapar" el panic
// (igual que en 06_interfaces_y_errores/19_panic_recover) y verificar que ocurrió.
// Si NO ocurre el panic esperado, marcamos el test como fallido.

func TestDividirEstrictoPanica(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("se esperaba un panic, pero la función no paniqueó")
		}
	}()

	DividirEstricto(10, 0) // esto DEBERÍA paniquear
	t.Error("esta línea no debería alcanzarse")
}

func TestDividirEstrictoCasoNormal(t *testing.T) {
	resultado := DividirEstricto(10, 2)
	if resultado != 5 {
		t.Errorf("DividirEstricto(10, 2) = %.2f; esperado 5", resultado)
	}
}
