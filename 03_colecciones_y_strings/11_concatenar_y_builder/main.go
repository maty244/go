package main

import (
	"fmt"
	"strings"
)

// =========================================================
// CONCATENAR STRINGS Y strings.Builder
// =========================================================
// Hay varias formas de juntar strings. Las cuatro que importan:
//
//   1. +                → simple, para POCOS pedazos
//   2. fmt.Sprintf      → cuando mezclás texto con números/formato
//   3. strings.Join     → cuando ya tenés un []string
//   4. strings.Builder  → cuando armás un string EN UN LOOP
//
// ¿Por qué importa? Porque los strings son INMUTABLES (tema 04):
// cada + crea un string NUEVO copiando todo lo anterior.
// En un loop largo eso es copiar, y copiar, y copiar...

func main() {
	// ─────────────────────────────────────────────────────────
	// 1. OPERADOR +: perfecto para pocos pedazos
	// ─────────────────────────────────────────────────────────
	nombre := "Matias"
	saludo := "Hola, " + nombre + "! Bienvenido."
	fmt.Println("=== Operador + ===")
	fmt.Println(saludo)

	// += también funciona (reasigna con un string nuevo)
	mensaje := "Go"
	mensaje += " es"
	mensaje += " genial"
	fmt.Println(mensaje)

	// OJO: + solo une strings. Para números hace falta convertir
	// (con Sprintf acá, o strconv en el tema 08):
	// "edad: " + 25 // ❌ no compila: mismatched types

	// ─────────────────────────────────────────────────────────
	// 2. fmt.Sprintf: texto + valores con formato
	// ─────────────────────────────────────────────────────────
	// Como Printf, pero en vez de imprimir, RETORNA el string.

	producto := "Yerba"
	precio := 3500.50
	stock := 12

	etiqueta := fmt.Sprintf("%s — $%.2f (%d unidades)", producto, precio, stock)
	fmt.Println("\n=== fmt.Sprintf ===")
	fmt.Println(etiqueta)

	// ─────────────────────────────────────────────────────────
	// 3. strings.Join: unir un slice con separador
	// ─────────────────────────────────────────────────────────
	// Join es la forma idiomática (y eficiente) de unir un []string.
	// Calcula el tamaño final UNA vez y copia todo UNA vez.

	frutas := []string{"manzana", "banana", "uva"}
	fmt.Println("\n=== strings.Join ===")
	fmt.Println(strings.Join(frutas, ", "))  // manzana, banana, uva
	fmt.Println(strings.Join(frutas, " | ")) // manzana | banana | uva
	fmt.Println(strings.Join(frutas, ""))    // manzanabananauva

	// ─────────────────────────────────────────────────────────
	// EL PROBLEMA: concatenar con + DENTRO de un loop
	// ─────────────────────────────────────────────────────────
	// Cada vuelta crea un string nuevo y COPIA todo lo acumulado:
	//   vuelta 1: copia 1 elemento
	//   vuelta 2: copia 2 elementos
	//   vuelta 3: copia 3 elementos...
	// Con n vueltas terminás copiando ~n²/2 datos. Para 10 ítems
	// no se nota; para 100.000 sí (y mucho).

	fmt.Println("\n=== MAL: + en un loop ===")
	items := []string{"pan", "leche", "huevos", "yerba", "azúcar"}

	// MAL (funciona, pero escala pésimo):
	lista := ""
	for i, item := range items {
		if i > 0 {
			lista += ", " // cada += copia TODO lo anterior
		}
		lista += item
	}
	fmt.Println("resultado:", lista)

	// ─────────────────────────────────────────────────────────
	// 4. strings.Builder: la solución idiomática
	// ─────────────────────────────────────────────────────────
	// Un Builder acumula en un buffer interno que crece de a
	// saltos (como append en un slice). Recién al final, con
	// .String(), se construye el string definitivo. Una sola vez.
	//
	// Métodos principales:
	//   b.WriteString("texto") → agrega un string
	//   b.WriteRune('ñ')       → agrega una rune
	//   b.WriteByte('a')       → agrega un byte
	//   b.Len()                → bytes acumulados hasta ahora
	//   b.String()             → el string final

	fmt.Println("\n=== BIEN: strings.Builder ===")

	var b strings.Builder // listo para usar, sin inicialización especial
	for i, item := range items {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(item)
	}
	fmt.Println("resultado:", b.String())
	fmt.Println("largo:", b.Len(), "bytes")

	// Ejemplo típico: armar un reporte línea por línea
	var reporte strings.Builder
	reporte.WriteString("REPORTE DE STOCK\n")
	for _, item := range items {
		reporte.WriteString(fmt.Sprintf("  - %s\n", item))
	}
	reporte.WriteString("fin del reporte")
	fmt.Println("\n" + reporte.String())

	// ─────────────────────────────────────────────────────────
	// strings.Repeat: repetir un patrón
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== strings.Repeat ===")
	fmt.Println(strings.Repeat("=", 30)) // línea separadora
	fmt.Println(strings.Repeat("na ", 4) + "¡Batman!")

	// ─────────────────────────────────────────────────────────
	// RESUMEN: ¿cuál uso?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  SITUACIÓN                          USAR")
	fmt.Println("  ─────────────────────────────────  ─────────────────")
	fmt.Println("  2-4 pedazos sueltos                +")
	fmt.Println("  texto con números/formato          fmt.Sprintf")
	fmt.Println("  ya tengo un []string               strings.Join")
	fmt.Println("  acumular en un loop                strings.Builder")
	fmt.Println("  repetir un patrón                  strings.Repeat")
}
