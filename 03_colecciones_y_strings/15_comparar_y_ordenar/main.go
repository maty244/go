package main

import (
	"fmt"
	"sort"
	"strings"
)

// =========================================================
// COMPARAR Y ORDENAR STRINGS
// =========================================================
// Los strings se comparan con los operadores de siempre:
//
//   ==  !=   → ¿son iguales?
//   <  >  <= >=  → orden "lexicográfico" (de diccionario)
//
// El orden compara BYTE a BYTE según el valor numérico.
// Eso trae dos sorpresas para hispanohablantes:
//
//   1. MAYÚSCULAS primero: "Z" (90) < "a" (97).
//      Todo el abecedario en mayúscula vale menos que
//      cualquier minúscula.
//   2. TILDES al final: "á" (2 bytes, >195) vale más que
//      cualquier letra sin tilde. "árbol" queda después
//      de "zorro" al ordenar.
//
// Para programas serios con español existe el paquete
// golang.org/x/text/collate (avanzado, fuera de esta biblia).
// Para esta etapa: conocer la trampa y normalizar cuando haga falta.

func main() {
	// ─────────────────────────────────────────────────────────
	// IGUALDAD: == compara contenido (no punteros)
	// ─────────────────────────────────────────────────────────
	// A diferencia de otros lenguajes (Java y su .equals()),
	// en Go == compara el TEXTO. Simple y directo.

	fmt.Println("=== Igualdad ===")
	a := "hola"
	b := "ho" + "la"                                 // construido distinto, mismo contenido
	fmt.Println(`"hola" == "ho"+"la" →`, a == b)     // true
	fmt.Println(`"hola" == "Hola"   →`, a == "Hola") // false (mayúscula)
	fmt.Println(`"hola" != "chau"   →`, a != "chau") // true

	// Para ignorar mayúsculas: EqualFold (visto en el tema 06)
	fmt.Println(`EqualFold("hola","HOLA") →`, strings.EqualFold("hola", "HOLA")) // true

	// ─────────────────────────────────────────────────────────
	// ORDEN LEXICOGRÁFICO: < y >
	// ─────────────────────────────────────────────────────────
	// Compara letra a letra desde la izquierda, como en el
	// diccionario. La primera diferencia decide.
	// Si uno es prefijo del otro, el más corto va primero.

	fmt.Println("\n=== Orden lexicográfico ===")
	fmt.Println(`"ana" < "berta"  →`, "ana" < "berta")  // true (a < b)
	fmt.Println(`"auto" < "avion" →`, "auto" < "avion") // true (u < v)
	fmt.Println(`"sol" < "soles"  →`, "sol" < "soles")  // true (prefijo)
	fmt.Println(`"10" < "9"       →`, "10" < "9")       // true (¡compara texto, no números!)

	// OJO con la última: "10" < "9" porque compara el carácter
	// '1' contra '9'. Para comparar números, primero convertir
	// con strconv.Atoi (tema 08).

	// ─────────────────────────────────────────────────────────
	// LAS DOS TRAMPAS: mayúsculas y tildes
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Trampas del orden por bytes ===")
	fmt.Println(`"Zorro" < "ana"  →`, "Zorro" < "ana")    // true: Z=90 < a=97
	fmt.Println(`"zorro" < "ángel" →`, "zorro" < "ángel") // true: á pesa más que z

	// Solución casera: comparar versiones normalizadas
	x, y := "Zorro", "ana"
	fmt.Println("normalizado:", strings.ToLower(x) < strings.ToLower(y)) // false: zorro > ana

	// ─────────────────────────────────────────────────────────
	// strings.Compare: -1, 0, 1
	// ─────────────────────────────────────────────────────────
	// Mismo orden que <, pero devuelve un int (estilo C).
	// La propia documentación de Go recomienda usar los
	// operadores ==, <, > — Compare existe por simetría.

	fmt.Println("\n=== strings.Compare ===")
	fmt.Println(`Compare("a","b") →`, strings.Compare("a", "b")) // -1 (a viene antes)
	fmt.Println(`Compare("b","b") →`, strings.Compare("b", "b")) //  0 (iguales)
	fmt.Println(`Compare("c","b") →`, strings.Compare("c", "b")) //  1 (c viene después)

	// ─────────────────────────────────────────────────────────
	// ORDENAR UN []string
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Ordenar slices de strings ===")
	nombres := []string{"matias", "ana", "luis", "berta", "carlos"}
	sort.Strings(nombres)
	fmt.Println("ordenado:", nombres)

	// Mirá cómo las trampas afectan el orden real:
	mezcla := []string{"zorro", "Ana", "árbol", "banana", "Único"}
	sort.Strings(mezcla)
	fmt.Println("con mayúsculas y tildes:", mezcla)
	// → mayúsculas primero, tildes al final

	// Orden personalizado con sort.Slice: ignorando mayúsculas/caso
	sort.Slice(mezcla, func(i, j int) bool {
		return strings.ToLower(mezcla[i]) < strings.ToLower(mezcla[j])
	})
	fmt.Println("ignorando mayúsculas:   ", mezcla)
	// (las tildes igual quedan al final: eso ya pide x/text/collate)

	// Ordenar por LONGITUD (otro criterio custom)
	palabras := []string{"sol", "estrella", "luna", "mar"}
	sort.Slice(palabras, func(i, j int) bool {
		return len(palabras[i]) < len(palabras[j])
	})
	fmt.Println("por longitud:", palabras)

	// ─────────────────────────────────────────────────────────
	// BÚSQUEDA EN SLICE ORDENADO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Búsqueda binaria ===")
	sort.Strings(nombres) // requisito: estar ordenado
	pos := sort.SearchStrings(nombres, "luis")
	fmt.Printf("'luis' está en el índice %d de %v\n", pos, nombres)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  ==, !=, <, >          → comparan contenido, byte a byte")
	fmt.Println("  EqualFold(a, b)       → igualdad sin mayúsculas")
	fmt.Println("  Trampa 1              → \"Z\" < \"a\" (mayúsculas valen menos)")
	fmt.Println("  Trampa 2              → tildes pesan más (\"á\" > \"z\")")
	fmt.Println("  Trampa 3              → \"10\" < \"9\" (texto, no números)")
	fmt.Println("  sort.Strings(s)       → ordena un []string")
	fmt.Println("  sort.Slice(s, func)   → orden con criterio propio")
}
