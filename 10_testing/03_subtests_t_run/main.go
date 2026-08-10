package main

import (
	"fmt"
	"strings"
	"unicode"
)

// =========================================================
// SUBTESTS CON t.Run
// =========================================================
// En el tema 02 armamos una tabla de casos, pero si UNO falla, el
// mensaje de error no siempre deja claro CUÁL. t.Run() resuelve
// esto: convierte cada caso de la tabla en un "subtest" con
// nombre propio, que aparece identificado en el output de
// `go test -v`, y que se puede correr INDIVIDUALMENTE.

// NormalizarNombre limpia un nombre de cliente: sin espacios de
// más, con la primera letra de cada palabra en mayúscula.
func NormalizarNombre(nombre string) string {
	nombre = strings.TrimSpace(nombre)
	if nombre == "" {
		return ""
	}
	palabras := strings.Fields(nombre) // separa por espacios, ignora espacios extra
	for i, p := range palabras {
		// Convertimos a runes antes de tomar la primera letra. p[:1] toma un
		// byte y rompe letras UTF-8 como Á, Ñ o É.
		letras := []rune(strings.ToLower(p))
		letras[0] = unicode.ToUpper(letras[0])
		palabras[i] = string(letras)
	}
	return strings.Join(palabras, " ")
}

func main() {
	fmt.Println("=== NormalizarNombre ===")

	ejemplos := []string{"  matias  ", "ANA", "carlos GOMEZ", "  "}
	for _, e := range ejemplos {
		fmt.Printf("  %q → %q\n", e, NormalizarNombre(e))
	}

	fmt.Println("\n(mirá main_test.go: los mismos casos, ahora con t.Run)")
	fmt.Println("Probá: go test -v ./10_testing/03_subtests_t_run/")
	fmt.Println("       go test -run TestNormalizarNombre/nombre_vacio ./10_testing/03_subtests_t_run/")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  t.Run(nombre, func(t *testing.T) {...})  → un subtest con nombre")
	fmt.Println("  go test -v                                → muestra cada subtest por separado")
	fmt.Println("  go test -run Test/nombre_del_subtest      → corre SOLO ese subtest")
	fmt.Println("  Si un subtest falla, los demás             → siguen corriendo igual")
}
