package main

import "fmt"

// =========================================================
// TABLE-DRIVEN TESTS: EL PATRÓN DE TESTING DE GO
// =========================================================
// Cuando una función tiene VARIOS casos para probar (varios
// valores de entrada y su salida esperada), escribir un TestX
// por cada caso es repetitivo. El patrón idiomático de Go es
// armar una TABLA (un slice de casos) y recorrerla con un loop,
// llamando a t.Run() para cada caso (eso lo vemos a fondo en el
// tema 03; acá nos enfocamos en la TABLA en sí).
//
// Es TAN común que, si alguna vez ves código Go real, este
// patrón va a aparecer en la gran mayoría de los tests.

// Clasificar devuelve una categoría de precio, útil para
// mostrar distinto en la interfaz según el rango.
func Clasificar(precio float64) string {
	switch {
	case precio <= 0:
		return "inválido"
	case precio < 1000:
		return "económico"
	case precio < 50000:
		return "medio"
	default:
		return "premium"
	}
}

func main() {
	fmt.Println("=== Función con VARIOS casos a testear ===")

	precios := []float64{-10, 0, 500, 15000, 500000}
	for _, p := range precios {
		fmt.Printf("  Clasificar(%.2f) = %s\n", p, Clasificar(p))
	}

	fmt.Println("\n(mirá main_test.go: un test-driven test para Clasificar)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Table-driven test  → un slice de struct con {entrada, esperado}")
	fmt.Println("  Un solo loop        → recorre la tabla y valida cada caso")
	fmt.Println("  Ventaja             → agregar un caso nuevo es agregar UNA línea")
	fmt.Println("  Es EL patrón        → casi todo el código Go real lo usa")
}
