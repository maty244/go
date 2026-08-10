package main

import "fmt"

// =========================================================
// go test: EL FRAMEWORK DE TESTING INTEGRADO EN GO
// =========================================================
// Go trae testing incorporado en la herramienta `go`, sin
// necesitar librerías externas. La convención es:
//
//   1. El archivo termina en "_test.go". No necesita llamarse igual
//      que el archivo de producción: la unidad que se prueba es el paquete.
//      En esta carpeta usamos main_test.go por comodidad.
//   2. Cada función de test empieza con "Test" + Mayúscula, y
//      recibe un *testing.T: func TestAlgo(t *testing.T)
//   3. Se ejecuta con: go test ./...
//
// Este archivo (main.go) tiene la función que vamos a testear.
// Es EXACTAMENTE como los archivos anteriores: package main,
// función normal, corre con `go run .`. Lo nuevo es main_test.go.

// Sumar es la función que vamos a poner a prueba.
func Sumar(a, b int) int {
	return a + b
}

// EsPar dice si un número es par.
func EsPar(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println("=== Funciones normales, listas para testear ===")
	fmt.Println("Sumar(2, 3) =", Sumar(2, 3))
	fmt.Println("EsPar(4) =", EsPar(4))
	fmt.Println("EsPar(7) =", EsPar(7))

	fmt.Println("\n=== Cómo correr los tests ===")
	fmt.Println("  go test ./10_testing/01_go_test_basico/     → corre los tests de esta carpeta")
	fmt.Println("  go test ./10_testing/01_go_test_basico/ -v  → modo verboso (ver cada test)")
	fmt.Println("  go test ./...                             → corre TODOS los tests del módulo")

	fmt.Println("\n(mirá main_test.go en esta misma carpeta para ver los tests)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  archivo_test.go        → convención de nombre para archivos de test")
	fmt.Println("  func TestX(t *testing.T) → una función de test")
	fmt.Println("  t.Errorf(...)          → marca el test como fallido, pero SIGUE corriendo")
	fmt.Println("  t.Fatalf(...)          → marca como fallido y CORTA el test ahí mismo")
	fmt.Println("  go test ./...          → ejecuta todos los tests del proyecto")
}
