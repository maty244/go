package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

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

func BenchmarkInvertir(b *testing.B) {
	entrada := "Go mide este texto con UTF-8: pingüino 🙂"
	for b.Loop() { // Go 1.24+: reemplaza al patrón manual for i < b.N
		Invertir(entrada)
	}
}

func ExampleInvertir() {
	fmt.Println(Invertir("Go🙂"))
	// Output: 🙂oG
}
