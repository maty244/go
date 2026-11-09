package main

import "fmt"

// =========================================================
// FUZZING, BENCHMARKS Y EJEMPLOS EJECUTABLES
// =========================================================
// Además de `func TestX(t *testing.T)`, el paquete `testing` de Go
// soporta otros tres tipos de funciones en archivos `*_test.go`:
//
//   1. Fuzzing:     func FuzzX(f *testing.F)
//      Genera entradas aleatorias/mutadas automáticamente para
//      descubrir casos borde (UTF-8 raro, strings vacíos, etc.)
//      que no se te ocurrirían a mano.
//   2. Benchmarks:  func BenchmarkX(b *testing.B)
//      Mide cuánto tiempo (ns/op) y memoria tarda una función
//      usando `for b.Loop()` (incorporado en Go 1.24+).
//   3. Examples:    func ExampleX()
//      Sirve como documentación viva: `go test` ejecuta la función
//      y compara lo que imprime por stdout contra el comentario
//      final `// Output:`.

func Invertir(texto string) string {
	runas := []rune(texto)
	for izquierda, derecha := 0, len(runas)-1; izquierda < derecha; izquierda, derecha = izquierda+1, derecha-1 {
		runas[izquierda], runas[derecha] = runas[derecha], runas[izquierda]
	}
	return string(runas)
}

func main() {
	fmt.Println("=== Invertir con soporte UTF-8 (runes) ===")
	fmt.Println(Invertir("Go y UTF-8: ñ"))
	fmt.Println("\nProbá correr desde la raíz:")
	fmt.Println("  go test -v ./10_testing/08_fuzzing_benchmarks_ejemplos")
	fmt.Println("  go test -fuzz=Fuzz -fuzztime=3s ./10_testing/08_fuzzing_benchmarks_ejemplos")
	fmt.Println("  go test -bench=. ./10_testing/08_fuzzing_benchmarks_ejemplos")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  func FuzzX(f *testing.F)      → prueba propiedades con entradas generadas")
	fmt.Println("  func BenchmarkX(b *testing.B) → mide rendimiento con `for b.Loop()`")
	fmt.Println("  func ExampleX()               → ejemplo documentado verificado con // Output:")
}
