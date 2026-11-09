package main

import (
	"strings"
	"testing"
)

func TestConstruccionDeTexto(t *testing.T) {
	for _, partes := range [][]string{nil, {}, {"Go"}, {"ñ", " ", "🙂"}, {"", "Go", ""}} {
		// Join es la referencia del contrato, independiente de los dos loops.
		esperado := strings.Join(partes, "")
		if a, b := concatenar(partes), conBuilder(partes); a != esperado || b != esperado {
			t.Fatalf("partes=%q; += %q; Builder %q; quiero %q", partes, a, b, esperado)
		}
	}
}

func BenchmarkConcatenar(b *testing.B) {
	partes := strings.Split(strings.Repeat("Go ñ 🙂 ", 100), " ")
	b.ReportAllocs() // B/op y allocs/op: costo de memoria por operación
	for b.Loop() {   // el tiempo de preparación queda fuera de la medición
		concatenar(partes)
	}
}

func BenchmarkBuilder(b *testing.B) {
	partes := strings.Split(strings.Repeat("Go ñ 🙂 ", 100), " ")
	b.ReportAllocs()
	for b.Loop() {
		conBuilder(partes)
	}
}
