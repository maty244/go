package main

import "testing"

// =========================================================
// LA TABLA: un slice de casos
// =========================================================
// Cada caso es un struct anónimo con lo que necesitamos: la
// entrada y la salida esperada. Le damos un nombre a cada caso
// (lo usamos a fondo en el tema 03 con t.Run).

func TestClasificar(t *testing.T) {
	casos := []struct {
		nombre   string
		precio   float64
		esperado string
	}{
		{"precio negativo", -10, "inválido"},
		{"precio cero", 0, "inválido"},
		{"precio económico", 500, "económico"},
		{"límite económico/medio", 999.99, "económico"},
		{"precio medio", 15000, "medio"},
		{"límite medio/premium", 49999.99, "medio"},
		{"precio premium", 500000, "premium"},
	}

	// UN SOLO loop valida TODOS los casos de la tabla.
	for _, c := range casos {
		resultado := Clasificar(c.precio)
		if resultado != c.esperado {
			t.Errorf("%s: Clasificar(%.2f) = %q; esperado %q",
				c.nombre, c.precio, resultado, c.esperado)
		}
	}
}

// ─────────────────────────────────────────────────────────
// AGREGAR UN CASO NUEVO ES SOLO AGREGAR UNA LÍNEA
// ─────────────────────────────────────────────────────────
// Si mañana Clasificar cambia sus reglas, agregás una fila a la
// tabla de arriba. No hace falta escribir una función de test
// nueva ni repetir la lógica de comparación.

// ─────────────────────────────────────────────────────────
// COMPARAR CON LA ALTERNATIVA: un test por caso (MÁS repetitivo)
// ─────────────────────────────────────────────────────────
// func TestClasificarNegativo(t *testing.T) {
// 	if Clasificar(-10) != "inválido" { t.Error("...") }
// }
// func TestClasificarCero(t *testing.T) {
// 	if Clasificar(0) != "inválido" { t.Error("...") }
// }
// ... y así por cada caso. Con la tabla, esto se evita por completo.
