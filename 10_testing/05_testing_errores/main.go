package main

import (
	"errors"
	"fmt"
)

// =========================================================
// TESTEAR FUNCIONES QUE DEVUELVEN error
// =========================================================
// La mayoría de tus funciones reales devuelven (T, error), como
// vimos en 06_interfaces_y_errores/. Testearlas bien significa chequear
// TANTO el caso feliz (sin error) COMO los casos de error
// (el tipo correcto de error, en el momento correcto).

var ErrDivisionPorCero = errors.New("no se puede dividir por cero")
var ErrMontoNegativo = errors.New("el monto no puede ser negativo")

// Dividir devuelve un error si b es 0.
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionPorCero
	}
	return a / b, nil
}

// ValidarMonto devuelve un error si el monto es negativo.
func ValidarMonto(monto float64) error {
	if monto < 0 {
		return fmt.Errorf("ValidarMonto(%.2f): %w", monto, ErrMontoNegativo)
	}
	return nil
}

// DividirEstricto es una versión que PANIQUEA en vez de devolver
// error (a propósito, para mostrar cómo se testea un panic
// esperado). En código real, casi siempre preferís error a panic
// (06_interfaces_y_errores/19_panic_recover), pero a veces heredás/necesitás testear
// una función así.
func DividirEstricto(a, b float64) float64 {
	if b == 0 {
		panic("DividirEstricto: división por cero")
	}
	return a / b
}

func main() {
	fmt.Println("=== Funciones que devuelven error ===")

	resultado, err := Dividir(10, 2)
	fmt.Printf("Dividir(10, 2) = %.2f, err=%v\n", resultado, err)

	_, err = Dividir(10, 0)
	fmt.Printf("Dividir(10, 0): err=%v\n", err)

	fmt.Println("\n(mirá main_test.go: cómo testear el caso feliz Y los errores)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Testear el caso feliz    → valor devuelto Y err == nil")
	fmt.Println("  Testear el caso de error → err != nil Y errors.Is del error esperado")
	fmt.Println("  errors.Is en el test     → mismo patrón que en el código real")
	fmt.Println("  recover en un test       → capturar panics esperados (ver ejemplo abajo)")
}
