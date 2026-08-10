package main

import "fmt"

// =========================================================
// INDEXAR E ITERAR STRINGS
// =========================================================
// Acá se juntan las dos formas de ver un string:
//
//   s[i]          → te da el BYTE en la posición i (uint8)
//   for range s   → te da las RUNES (letras), una por una
//
// Con texto ASCII puro las dos coinciden. Con tildes, eñes
// o emojis, indexar por byte puede "cortar" una letra a la
// mitad y darte basura. Regla práctica:
//
//   ¿Querés recorrer letras? → for range
//   ¿Querés la letra N?      → []rune(s)[N]
//   ¿Sabés que es ASCII?     → s[i] está bien

func main() {
	// ─────────────────────────────────────────────────────────
	// INDEXAR: s[i] devuelve un BYTE
	// ─────────────────────────────────────────────────────────
	s := "gol"
	fmt.Println("=== Indexar (ASCII, todo bien) ===")
	fmt.Println("s[0] =", s[0])              // 103 (el byte, un número)
	fmt.Printf("s[0] como char: %c\n", s[0]) // g
	fmt.Printf("s[2] como char: %c\n", s[2]) // l

	// ─────────────────────────────────────────────────────────
	// LA TRAMPA: indexar texto con tildes
	// ─────────────────────────────────────────────────────────
	// "café": la é ocupa los bytes 3 y 4.
	// Pedir s[3] te da MEDIA letra → un byte que solo no significa nada.

	cafe := "café"
	fmt.Println("\n=== Trampa con tildes ===")
	fmt.Println("string:", cafe, "| len:", len(cafe)) // len=5, pero 4 letras
	fmt.Printf("cafe[3] = %d (¡medio carácter!)\n", cafe[3])
	fmt.Printf("cafe[3] como char: %c (basura: %q)\n", cafe[3], string(cafe[3]))

	// La forma correcta de pedir "la cuarta letra":
	letras := []rune(cafe)
	fmt.Printf("[]rune(cafe)[3] = %c (correcto)\n", letras[3]) // é

	// ─────────────────────────────────────────────────────────
	// FOR CLÁSICO: itera por BYTES
	// ─────────────────────────────────────────────────────────
	// Con for i := 0; i < len(s); i++ recorrés byte a byte.
	// Mirá cómo la ñ aparece como dos números sueltos.

	palabra := "niño"
	fmt.Println("\n=== for clásico: bytes ===")
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("  índice %d → byte %3d → %q\n", i, palabra[i], string(palabra[i]))
	}
	// la ñ se parte en los bytes 195 y 177

	// ─────────────────────────────────────────────────────────
	// FOR RANGE: itera por RUNES (la forma correcta)
	// ─────────────────────────────────────────────────────────
	// for range sobre un string decodifica UTF-8 automáticamente:
	//   i → índice del BYTE donde empieza la letra (¡puede saltar!)
	//   r → la rune completa
	// Fijate que el índice salta de 2 a 4: la ñ ocupó 2 bytes.

	fmt.Println("\n=== for range: runes ===")
	for i, r := range palabra {
		fmt.Printf("  byte-índice %d → rune %c (código %d)\n", i, r, r)
	}

	// Si solo querés las letras, ignorá el índice con _
	fmt.Print("letras de 'añejo': ")
	for _, r := range "añejo" {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// ACCESO POR POSICIÓN DE LETRA: []rune
	// ─────────────────────────────────────────────────────────
	// Si necesitás "la letra en la posición N" o recorrer al revés,
	// convertí una vez a []rune y trabajá con el slice.

	fmt.Println("\n=== []rune para acceso por letra ===")
	texto := "camión"
	rs := []rune(texto)
	fmt.Printf("primera letra: %c\n", rs[0])
	fmt.Printf("última letra:  %c\n", rs[len(rs)-1])

	// Recorrer al revés (clásico ejercicio)
	fmt.Print("al revés: ")
	for i := len(rs) - 1; i >= 0; i-- {
		fmt.Printf("%c", rs[i])
	}
	fmt.Println() // nóimac

	// ─────────────────────────────────────────────────────────
	// SUB-STRINGS CON SLICING: también corta por BYTES
	// ─────────────────────────────────────────────────────────
	// s[a:b] funciona como en slices, pero sobre BYTES.
	// Con ASCII es seguro; con tildes podés cortar mal.

	fmt.Println("\n=== Slicing de strings ===")
	frase := "hola mundo"
	fmt.Println("frase[0:4]:", frase[0:4]) // "hola" (ASCII, seguro)
	fmt.Println("frase[5:]:", frase[5:])   // "mundo"

	// MAL: cortar "café" en el byte 4 parte la é
	fmt.Printf("MAL  cafe[:4] = %q (é cortada)\n", cafe[:4])
	// BIEN: cortar sobre []rune y volver a string
	fmt.Printf("BIEN string([]rune(cafe)[:4]) = %q\n", string([]rune(cafe)[:4]))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  s[i]            → BYTE (peligroso con tildes/eñes)")
	fmt.Println("  s[a:b]          → sub-string por BYTES (ídem)")
	fmt.Println("  for range s     → letra por letra (decodifica UTF-8)")
	fmt.Println("  []rune(s)[n]    → letra N de verdad")
	fmt.Println("  Regla: para recorrer texto, usá for range.")
}
