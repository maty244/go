package main

import (
	"fmt"
	"strconv"
)

func main() {
	// =========================================================
	// CONVERSIONES DE TIPO EXPLÍCITAS
	// =========================================================
	// Go es un lenguaje de TIPADO ESTÁTICO Y FUERTE.
	// A diferencia de JavaScript o Python, Go NO hace conversiones
	// automáticas entre tipos. Si querés convertir, debés hacerlo
	// EXPLÍCITAMENTE. Esto evita bugs silenciosos.
	//
	// Sintaxis básica: tipoDestino(valor)

	// ─────────────────────────────────────────────────────────
	// CONVERSIONES ENTRE TIPOS NUMÉRICOS
	// ─────────────────────────────────────────────────────────

	var entero int = 42
	var flotante float64 = 3.99
	var pequeño int8 = 10

	fmt.Println("=== Conversiones entre números ===")

	// int → float64
	enteroAFloat := float64(entero)
	fmt.Printf("int(%d) → float64: %f\n", entero, enteroAFloat)

	// float64 → int (TRUNCA, no redondea. Cuidado!)
	floatAEntero := int(flotante)
	fmt.Printf("float64(%.2f) → int: %d  (se trunca, no se redondea)\n", flotante, floatAEntero)

	// int8 → int (ensanchamiento, siempre seguro)
	pequeñoAGrande := int(pequeño)
	fmt.Printf("int8(%d) → int: %d\n", pequeño, pequeñoAGrande)

	// int → int8 (reducción, puede perder datos si el valor es muy grande!)
	grande := 300
	grandeAInt8 := int8(grande) // 300 no cabe en int8 (max 127), overflow!
	fmt.Printf("int(%d) → int8: %d  ⚠️ overflow silencioso!\n", grande, grandeAInt8)

	// int → int64
	var i int = 1000
	var i64 int64 = int64(i)
	fmt.Printf("int(%d) → int64(%d)\n", i, i64)

	// ─────────────────────────────────────────────────────────
	// CONVERSIONES ENTRE NUMÉRICOS Y STRING
	// ─────────────────────────────────────────────────────────
	// Para convertir entre string y números NUNCA usamos Tipo(valor)
	// porque eso interpreta el número como un código Unicode.
	// En cambio, usamos el paquete "strconv" (string conversion).

	fmt.Println("\n=== strconv: números ↔ strings ===")

	// int → string (usando strconv.Itoa, "Integer to ASCII")
	numero := 42
	textoNumero := strconv.Itoa(numero) // Itoa = Int to ASCII
	fmt.Printf("int(%d) → string: '%s'\n", numero, textoNumero)
	fmt.Printf("Tipo: %T\n", textoNumero)

	// string → int (puede fallar si el string no es un número)
	texto := "123"
	textoAInt, err := strconv.Atoi(texto) // Atoi = ASCII to Int
	if err != nil {
		fmt.Println("Error al convertir:", err)
	} else {
		fmt.Printf("string('%s') → int: %d\n", texto, textoAInt)
	}

	// ¿Qué pasa si el string no es un número válido?
	textoInvalido := "abc"
	_, err2 := strconv.Atoi(textoInvalido)
	if err2 != nil {
		fmt.Printf("\nstring('%s') no es un número válido: %v\n", textoInvalido, err2)
	}

	// float64 → string
	precio := 99.95
	textoPrecio := strconv.FormatFloat(precio, 'f', 2, 64)
	// 'f' = formato decimal, 2 = 2 decimales, 64 = float64
	fmt.Printf("\nfloat64(%.2f) → string: '%s'\n", precio, textoPrecio)

	// string → float64
	textoDecimal := "3.14159"
	decimalFloat, err3 := strconv.ParseFloat(textoDecimal, 64)
	if err3 == nil {
		fmt.Printf("string('%s') → float64: %f\n", textoDecimal, decimalFloat)
	}

	// bool → string
	valor := true
	textoBool := strconv.FormatBool(valor)
	fmt.Printf("\nbool(%v) → string: '%s'\n", valor, textoBool)

	// string → bool
	textoBool2 := "true"
	boolVal, _ := strconv.ParseBool(textoBool2)
	fmt.Printf("string('%s') → bool: %v\n", textoBool2, boolVal)

	// ─────────────────────────────────────────────────────────
	// TRAMPA CLÁSICA: convertir entero a string sin strconv (¡NO hacer esto!)
	// ─────────────────────────────────────────────────────────
	// En Go, convertir un entero/rune directamente con string(x) NO convierte
	// el número a sus dígitos ("65"), sino que busca el carácter Unicode cuyo
	// código sea 65 ('A').
	// De hecho, "go vet" bloquea string(int) directamente para evitar este error.
	// Si hacés string(rune(65)), obtenés "A"; para obtener "65", usá strconv.Itoa(65).

	fmt.Println("\n=== ⚠️ Trampa: string(rune(int)) vs strconv.Itoa ===")
	n := 65
	comoCaracter := string(rune(n)) // convierte el punto Unicode 65 al carácter "A"
	comoDigitos := strconv.Itoa(n)  // convierte el número 65 al texto "65"

	fmt.Printf("string(rune(%d)) = '%s'  ← interpreta %d como código Unicode!\n", n, comoCaracter, n)
	fmt.Printf("strconv.Itoa(%d) = '%s'  ← convierte el número a sus dígitos en texto\n", n, comoDigitos)

	// ─────────────────────────────────────────────────────────
	// CONVERSIONES CON BYTE Y RUNE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== byte, rune ↔ string ===")

	// string → []byte (slice de bytes)
	str := "Hola"
	bytes := []byte(str)
	fmt.Printf("string('%s') → []byte: %v\n", str, bytes)

	// []byte → string
	bytes[0] = 'h' // minúscula
	strDeBytes := string(bytes)
	fmt.Printf("[]byte(%v) → string: '%s'\n", bytes, strDeBytes)

	// string → []rune (slice de runes, para texto Unicode)
	textoUTF8 := "Hola Ñoño"
	runes := []rune(textoUTF8)
	fmt.Printf("\nstring('%s') → []rune (len=%d): %v\n", textoUTF8, len(runes), runes)
	fmt.Printf("Como caracteres:\n")
	for i, r := range runes {
		fmt.Printf("  rune[%d] = '%c' (Unicode: %d)\n", i, r, r)
	}

	// []rune → string
	strDeRunes := string(runes)
	fmt.Printf("[]rune → string: '%s'\n", strDeRunes)

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Leer datos de entrada (siempre llegan como string)
	// ─────────────────────────────────────────────────────────
	// Cuando un usuario ingresa datos por consola o los leemos de
	// un archivo, siempre llegan como string. Hay que convertirlos.

	fmt.Println("\n=== Caso real: procesar datos de usuario ===")

	edadTexto := "25"       // llegó como string
	precioTexto := "199.99" // llegó como string

	edadInt, _ := strconv.Atoi(edadTexto)
	precioFloat, _ := strconv.ParseFloat(precioTexto, 64)

	fmt.Printf("Edad: %d años (tipo: %T)\n", edadInt, edadInt)
	fmt.Printf("Precio: $%.2f (tipo: %T)\n", precioFloat, precioFloat)

	// Ahora podemos operar con ellos matemáticamente
	descuento := 0.15
	precioFinal := precioFloat * (1 - descuento)
	fmt.Printf("Con 15%% descuento: $%.2f\n", precioFinal)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Número ↔ Número:  usá TipoDestino(valor)  → float64(x), int(x)")
	fmt.Println("int    → string:  strconv.Itoa(n)          → '42'")
	fmt.Println("string → int:     strconv.Atoi(s)          → 42")
	fmt.Println("float  → string:  strconv.FormatFloat(...)  → '3.14'")
	fmt.Println("string → float:   strconv.ParseFloat(s, 64) → 3.14")
	fmt.Println("string ↔ []byte:  []byte(s) y string(b)    ")
	fmt.Println("string ↔ []rune:  []rune(s) y string(r)    (para Unicode)")
}
