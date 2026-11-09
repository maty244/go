package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// =========================================================
// 1. FUZZING: PROBAR UNA PROPIEDAD CON MILES DE ENTRADAS
// =========================================================
// En un test normal vos elegís la entrada y la salida esperada.
// En Fuzzing, vos definís una PROPIEDAD que siempre debe cumplirse
// (acá: invertir dos veces un texto UTF-8 válido devuelve el texto
// original) y Go inventa entradas para intentar romperla.
// Con `go test` normal corre solo las semillas de `f.Add(...)`;
// con `go test -fuzz=Fuzz` muta las semillas buscando fallos.
func FuzzInvertir(f *testing.F) {
	for _, semilla := range []string{"", "Go", "mañana", "🙂界"} {
		f.Add(semilla)
	}
	f.Fuzz(func(t *testing.T, entrada string) {
		if !utf8.ValidString(entrada) {
			t.Skip() // la propiedad está definida para texto UTF-8 válido
		}
		if obtenido := Invertir(Invertir(entrada)); obtenido != entrada {
			t.Fatalf("invertir dos veces %q produjo %q", entrada, obtenido)
		}
	})
}

// =========================================================
// 2. BENCHMARK: MEDIR RENDIMIENTO CON b.Loop() (Go 1.24+)
// =========================================================
// `for b.Loop()` (introducido en Go 1.24) reemplaza al clásico
// `for i := 0; i < b.N; i++`. Tiene dos grandes ventajas:
//   - No mide el tiempo de inicialización previo al loop ni la limpieza posterior.
//   - Evita que el compilador optimice y elimine el cuerpo del benchmark.
func BenchmarkInvertir(b *testing.B) {
	entrada := "Go mide este texto con UTF-8: pingüino 🙂"
	for b.Loop() { // Go 1.24+: reemplaza al patrón manual for i < b.N
		Invertir(entrada)
	}
}

// =========================================================
// 3. EXAMPLE: DOCUMENTACIÓN QUE SE EJECUTA COMO TEST
// =========================================================
// `go test` ejecuta ExampleInvertir y verifica que lo impreso en
// stdout coincida exactamente con lo que sigue a `// Output:`.
func ExampleInvertir() {
	fmt.Println(Invertir("Go🙂"))
	// Output: 🙂oG
}
