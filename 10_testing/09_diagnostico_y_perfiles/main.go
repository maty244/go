package main

import (
	"fmt"
	"strings"
)

// =========================================================
// DIAGNÓSTICO, BENCHMARKS DE MEMORIA Y PERFILES (pprof)
// =========================================================
// Regla de oro del rendimiento en Go:
//   1. Primero asegurate de que sea CORRECTO (con un Test).
//   2. Después MEDÍ cuánto cuesta (con un Benchmark y b.ReportAllocs()).
//   3. Si es lento o gasta mucha memoria, mirá DÓNDE se va con pprof.
//
// Acá comparamos dos formas de unir un slice de strings:
//   - concatenar: usa `texto += parte` en un loop. Como los strings
//     en Go son inmutables, cada `+=` reserva un string nuevo y
//     copia todo lo anterior (crecimiento cuadrático en memoria).
//   - conBuilder: usa `strings.Builder`, que acumula los bytes en un
//     buffer interno y recién construye el string al final.

// Ambas funciones devuelven el mismo texto. El experimento mide el costo de
// construirlo. Las cifras dependen del compilador, la máquina y la entrada.
func concatenar(partes []string) string {
	var texto string
	for _, parte := range partes {
		texto += parte // cada paso puede crear otro string con todo lo acumulado
	}
	return texto
}

func conBuilder(partes []string) string {
	var texto strings.Builder
	for _, parte := range partes {
		texto.WriteString(parte) // acumula bytes en un buffer interno
	}
	return texto.String()
}

func main() {
	partes := []string{"Go", " ", "se", " ", "puede", " ", "medir"}
	fmt.Println(concatenar(partes)) // Go se puede medir
	fmt.Println(conBuilder(partes)) // misma salida
	fmt.Println("Tests: ¿es correcto? Benchmarks: ¿cuánto cuesta? Perfiles: ¿dónde cuesta?")
	fmt.Println("Abrí PRACTICA.md: no hace falta un servidor para medir este programa.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  b.ReportAllocs()           → reporta B/op (bytes) y allocs/op (reservas)")
	fmt.Println("  texto += parte en loop     → copia todo en cada paso (muchas allocs)")
	fmt.Println("  strings.Builder            → acumula en un buffer interno (pocas allocs)")
	fmt.Println("  -cpuprofile / -memprofile  → genera perfiles para inspeccionar con `go tool pprof`")
}
