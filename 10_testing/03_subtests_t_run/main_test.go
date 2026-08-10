package main

import "testing"

// =========================================================
// TABLA + t.Run: EL COMBO COMPLETO
// =========================================================
// Cada caso de la tabla se ejecuta con t.Run(nombre, func...).
// Ventajas sobre el tema 02:
//   1. Si el caso "nombre vacío" falla, el output de `go test -v`
//      dice EXACTAMENTE "--- FAIL: TestNormalizarNombre/nombre_vacio",
//      no solo "TestNormalizarNombre falló".
//   2. Podés correr un subtest puntual con -run, sin ejecutar
//      toda la tabla.
//   3. Un panic en un subtest NO frena a los demás subtests.

func TestNormalizarNombre(t *testing.T) {
	casos := []struct {
		nombre   string // nombre del CASO (no del cliente)
		entrada  string
		esperado string
	}{
		{"nombre simple en minúsculas", "matias", "Matias"},
		{"nombre todo en mayúsculas", "ANA", "Ana"},
		{"varias palabras", "carlos GOMEZ", "Carlos Gomez"},
		{"espacios de más", "  matias  ", "Matias"},
		{"unicode no se corta", "ÁLVARO núñez", "Álvaro Núñez"},
		{"nombre vacio", "  ", ""},
	}

	for _, c := range casos {
		// t.Run crea un subtest con el nombre del caso. El "t"
		// que recibe la función interna es propio de ESE subtest.
		t.Run(c.nombre, func(t *testing.T) {
			resultado := NormalizarNombre(c.entrada)
			if resultado != c.esperado {
				t.Errorf("NormalizarNombre(%q) = %q; esperado %q",
					c.entrada, resultado, c.esperado)
			}
		})
	}
}

// ─────────────────────────────────────────────────────────
// ASÍ SE VE EL OUTPUT CON -v (a modo de referencia)
// ─────────────────────────────────────────────────────────
// === RUN   TestNormalizarNombre
// === RUN   TestNormalizarNombre/nombre_simple_en_minúsculas
// === RUN   TestNormalizarNombre/nombre_todo_en_mayúsculas
// === RUN   TestNormalizarNombre/varias_palabras
// === RUN   TestNormalizarNombre/espacios_de_más
// === RUN   TestNormalizarNombre/nombre_vacio
// --- PASS: TestNormalizarNombre (0.00s)
//     --- PASS: TestNormalizarNombre/nombre_simple_en_minúsculas (0.00s)
//     --- PASS: TestNormalizarNombre/nombre_todo_en_mayúsculas (0.00s)
//     ... (uno por cada subtest)
