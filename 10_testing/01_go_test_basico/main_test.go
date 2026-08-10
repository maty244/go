package main

import "testing"

// =========================================================
// TU PRIMER TEST
// =========================================================
// t.Errorf(...) marca el test como FALLIDO, pero deja que el
// resto de la función siga ejecutando (útil para reportar varios
// problemas de un mismo test). t.Fatalf(...) también lo marca
// como fallido, pero CORTA la ejecución ahí mismo (útil cuando
// seguir no tiene sentido, por ejemplo si algo previo ya rompió
// una precondición).

func TestSumar(t *testing.T) {
	resultado := Sumar(2, 3)
	esperado := 5

	if resultado != esperado {
		// t.Errorf: reporta el fallo pero NO corta la función.
		t.Errorf("Sumar(2, 3) = %d; esperado %d", resultado, esperado)
	}
}

func TestSumarConNegativos(t *testing.T) {
	resultado := Sumar(-5, 3)
	esperado := -2

	if resultado != esperado {
		t.Errorf("Sumar(-5, 3) = %d; esperado %d", resultado, esperado)
	}
}

func TestEsPar(t *testing.T) {
	if !EsPar(4) {
		t.Error("EsPar(4) debería ser true")
	}
	if EsPar(7) {
		t.Error("EsPar(7) debería ser false")
	}
}

// ─────────────────────────────────────────────────────────
// UN TEST QUE INTENCIONALMENTE FALLA (para ver cómo se ve)
// ─────────────────────────────────────────────────────────
// Dejamos este comentado a propósito: si lo descomentás y corrés
// `go test -v ./10_testing/01_go_test_basico/`, vas a ver el output
// de un test fallido, con el mensaje exacto que le pusimos.
//
// func TestSumarQueFalla(t *testing.T) {
// 	resultado := Sumar(2, 2)
// 	if resultado != 5 { // a propósito, para que falle
// 		t.Errorf("Sumar(2, 2) = %d; esperado %d", resultado, 5)
// 	}
// }

// ─────────────────────────────────────────────────────────
// t.Fatalf: CORTA el test si algo previo no tiene sentido seguir
// ─────────────────────────────────────────────────────────

func TestConFatal(t *testing.T) {
	numeros := []int{2, 4, 6, 8}

	if len(numeros) == 0 {
		// Si la lista estuviera vacía, seguir el test no tendría
		// sentido (numeros[0] rompería el test con un panic real).
		t.Fatal("la lista no debería estar vacía")
	}

	if !EsPar(numeros[0]) {
		t.Errorf("se esperaba que %d fuera par", numeros[0])
	}
}
