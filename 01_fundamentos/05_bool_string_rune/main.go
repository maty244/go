package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// =========================================================
	// TIPOS PRIMITIVOS - PARTE 2: string, bool, byte, rune
	// =========================================================

	// ─────────────────────────────────────────────────────────
	// BOOL (booleano)
	// ─────────────────────────────────────────────────────────
	// Solo puede tener dos valores: true o false.
	// Es el resultado de comparaciones y condiciones.
	// En Go el zero value de bool es false.

	var estaActivo bool = true
	var tieneDescuento bool = false
	esMayorDeEdad := true

	fmt.Println("=== bool ===")
	fmt.Println("¿Está activo?", estaActivo)
	fmt.Println("¿Tiene descuento?", tieneDescuento)
	fmt.Println("¿Es mayor de edad?", esMayorDeEdad)

	// Los booleanos son el resultado de comparaciones
	edad := 20
	esMayor := edad >= 18 // esto produce un bool
	fmt.Printf("edad=%d, ¿es mayor de 18? → %v\n", edad, esMayor)

	// Operaciones lógicas con bool
	tieneID := true
	puedeEntrar := esMayor && tieneID // ambas deben ser true
	fmt.Println("¿Puede entrar al boliche?", puedeEntrar)

	// ─────────────────────────────────────────────────────────
	// STRING
	// ─────────────────────────────────────────────────────────
	// Un string es una secuencia de bytes (texto).
	// En Go los strings son INMUTABLES: una vez creados no se modifican.
	// Se declaran con comillas dobles "".
	// Son UTF-8 por defecto (soportan emojis, acentos, etc.)

	nombre := "Matias"
	apellido := "García"
	frase := "Hola, ¿cómo estás?"
	conEmoji := "Go es genial 🚀"

	fmt.Println("\n=== string ===")
	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(frase)
	fmt.Println(conEmoji)

	// Concatenación con +
	nombreCompleto := nombre + " " + apellido
	fmt.Println("Nombre completo:", nombreCompleto)

	// Longitud de un string
	// IMPORTANTE: len() retorna BYTES, no caracteres.
	// Para texto ASCII, bytes == caracteres.
	// Para UTF-8 con acentos/emojis, bytes > caracteres.
	fmt.Println("\n=== Longitud de strings ===")
	fmt.Printf("len(\"%s\") = %d bytes\n", nombre, len(nombre))
	fmt.Printf("len(\"%s\") = %d bytes (la é ocupa 2 bytes en UTF-8)\n", apellido, len(apellido))
	fmt.Printf("Caracteres reales en \"%s\": %d\n", apellido, utf8.RuneCountInString(apellido))

	// Acceder a caracteres por índice
	// CUIDADO: el índice da un BYTE, no un caracter (rune).
	// Para texto ASCII está bien, para UTF-8 puede sorprender.
	saludo := "Hola"
	fmt.Println("\n=== Acceso por índice ===")
	fmt.Printf("saludo[0] = %d (valor del byte) = '%c' (como caracter)\n", saludo[0], saludo[0])
	fmt.Printf("saludo[1] = %d = '%c'\n", saludo[1], saludo[1])

	// Substrings (slices de string)
	texto := "Buenos Aires, Argentina"
	ciudad := texto[0:12] // bytes 0 al 11
	pais := texto[14:]    // desde byte 14 hasta el final
	primeros := texto[:6] // los primeros 6 bytes

	fmt.Println("\n=== Substrings ===")
	fmt.Println("Ciudad:", ciudad)
	fmt.Println("País:", pais)
	fmt.Println("Primeros:", primeros)

	// Strings multilínea con backticks ``
	// Con backticks el string es "raw": los \n son literales, no saltos de línea.
	json := `{
  "nombre": "Matias",
  "edad": 25,
  "ciudad": "Buenos Aires"
}`
	fmt.Println("\n=== String multilínea (raw string) ===")
	fmt.Println(json)

	// Funciones útiles del paquete strings
	frase2 := "  hola mundo  "
	fmt.Println("\n=== Funciones de strings ===")
	fmt.Println("Original:              '" + frase2 + "'")
	fmt.Println("ToUpper:               '" + strings.ToUpper(frase2) + "'")
	fmt.Println("ToLower:               '" + strings.ToLower(frase2) + "'")
	fmt.Println("TrimSpace:             '" + strings.TrimSpace(frase2) + "'")
	fmt.Println("Contains 'mundo':      ", strings.Contains(frase2, "mundo"))
	fmt.Println("Replace:               '" + strings.Replace(frase2, "mundo", "Go", 1) + "'")
	fmt.Println("Split:                 ", strings.Split("a,b,c", ","))
	fmt.Println("Join:                  ", strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println("HasPrefix 'hola':      ", strings.HasPrefix(strings.TrimSpace(frase2), "hola"))
	fmt.Println("HasSuffix 'mundo':     ", strings.HasSuffix(strings.TrimSpace(frase2), "mundo"))
	fmt.Println("Count 'o':             ", strings.Count(frase2, "o"))

	// ─────────────────────────────────────────────────────────
	// BYTE
	// ─────────────────────────────────────────────────────────
	// byte es un alias para uint8 (número de 0 a 255).
	// Representa un byte de datos, un caracter ASCII, etc.
	// Se usa para trabajar con datos binarios, archivos, redes.

	var b byte = 65 // 65 en ASCII es la letra 'A'
	fmt.Println("\n=== byte ===")
	fmt.Printf("byte 65 como número: %d\n", b)
	fmt.Printf("byte 65 como caracter: %c\n", b)

	// Iterar sobre los bytes de un string
	palabra := "Hola"
	fmt.Println("\nBytes de 'Hola':")
	for i, b := range []byte(palabra) {
		fmt.Printf("  índice %d → byte %d → caracter '%c'\n", i, b, b)
	}

	// ─────────────────────────────────────────────────────────
	// RUNE
	// ─────────────────────────────────────────────────────────
	// rune es un alias para int32.
	// Representa un CÓDIGO UNICODE (un caracter real, no un byte).
	// Esto es lo que necesitás usar cuando trabajás con texto UTF-8
	// que incluye acentos, caracteres especiales, emojis, etc.

	var r rune = 'A'  // comilla simple para rune literal
	var r2 rune = 'Ñ' // caracter especial español
	var r3 rune = '🚀' // emoji (ocupa 4 bytes pero es 1 rune)

	fmt.Println("\n=== rune ===")
	fmt.Printf("rune 'A':  valor int32=%d, caracter=%c\n", r, r)
	fmt.Printf("rune 'Ñ':  valor int32=%d, caracter=%c\n", r2, r2)
	fmt.Printf("rune '🚀': valor int32=%d, caracter=%c\n", r3, r3)

	// Diferencia crítica: bytes vs runes en strings con UTF-8
	textoEspanol := "Ñoño"
	fmt.Println("\n=== bytes vs runes ===")
	fmt.Printf("String: '%s'\n", textoEspanol)
	fmt.Printf("len() en bytes: %d\n", len(textoEspanol))
	fmt.Printf("Caracteres reales (runes): %d\n", utf8.RuneCountInString(textoEspanol))

	// Iterar correctamente sobre caracteres con range (da runes, no bytes)
	fmt.Println("\nIteración correcta con range (runes):")
	for i, r := range textoEspanol {
		fmt.Printf("  índice byte %d → rune %d → caracter '%c'\n", i, r, r)
	}

	// ─────────────────────────────────────────────────────────
	// STRING BUILDER - para construcción eficiente de strings
	// ─────────────────────────────────────────────────────────
	// Concatenar strings con + en un bucle es ineficiente porque
	// crea un string nuevo en memoria en cada iteración.
	// strings.Builder es la solución eficiente.

	fmt.Println("\n=== strings.Builder (construcción eficiente) ===")
	var sb strings.Builder
	palabras := []string{"Go", "es", "rápido", "y", "simple"}
	for i, p := range palabras {
		sb.WriteString(p)
		if i < len(palabras)-1 {
			sb.WriteString(" ")
		}
	}
	resultado := sb.String()
	fmt.Println(resultado)

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE TIPOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("bool   → true/false, para condiciones y flags")
	fmt.Println("string → texto inmutable, UTF-8 nativo")
	fmt.Println("byte   → alias de uint8, para datos binarios y ASCII")
	fmt.Println("rune   → alias de int32, para caracteres Unicode reales")
}
