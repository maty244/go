package main

import (
	"fmt"
	"unicode/utf8"
)

// =========================================================
// UTF-8 Y RUNES
// =========================================================
// Las computadoras solo guardan números. Para guardar texto,
// hace falta una tabla que diga "este número = esta letra".
//
//   - UNICODE: la tabla universal. A cada carácter del mundo
//     le asigna un número llamado "code point".
//     Ej: 'a' = 97, 'ñ' = 241, '🎉' = 127881.
//
//   - UTF-8: la forma de GUARDAR esos números en bytes.
//     Es de tamaño variable:
//       1 byte  → letras ASCII (a-z, A-Z, 0-9, símbolos comunes)
//       2 bytes → tildes, eñes, griego, cirílico (ñ, é, á...)
//       3 bytes → chino, japonés, símbolos (中, €...)
//       4 bytes → emojis (🎉, 🚀...)
//
//   - RUNE: el tipo de Go para UN code point Unicode.
//     Es un alias de int32 (igual que byte es alias de uint8).
//     Se escribe con comillas SIMPLES: 'a', 'ñ', '🎉'.
//
// Go fue creado por los mismos que inventaron UTF-8
// (Ken Thompson y Rob Pike), por eso el soporte es nativo.

func main() {
	// ─────────────────────────────────────────────────────────
	// RUNE: un carácter, con comillas simples
	// ─────────────────────────────────────────────────────────
	// Una rune ES un número (int32). Si la imprimís directo,
	// ves el código. Con %c ves el carácter.

	var letra rune = 'a'
	var enie rune = 'ñ'
	var emoji rune = '🎉'

	fmt.Println("=== Runes son números ===")
	fmt.Println("'a' como número:", letra) // 97
	fmt.Println("'ñ' como número:", enie)  // 241
	fmt.Println("'🎉' como número:", emoji) // 127881
	fmt.Printf("Como caracteres: %c %c %c\n", letra, enie, emoji)
	fmt.Printf("El tipo de 'a' es: %T\n", letra) // int32 (rune es su alias)

	// ─────────────────────────────────────────────────────────
	// CUÁNTOS BYTES OCUPA CADA CARÁCTER
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Bytes por carácter (UTF-8) ===")
	fmt.Println(`len("a") =`, len("a")) // 1 byte
	fmt.Println(`len("ñ") =`, len("ñ")) // 2 bytes
	fmt.Println(`len("中") =`, len("中")) // 3 bytes
	fmt.Println(`len("🎉") =`, len("🎉")) // 4 bytes

	// ─────────────────────────────────────────────────────────
	// LEN vs RuneCountInString: bytes vs puntos de código
	// ─────────────────────────────────────────────────────────
	// len(s)                        → cantidad de BYTES
	// utf8.RuneCountInString(s)     → cantidad de RUNES, no de letras visibles
	// Un carácter visible (grafema) puede combinar varios puntos de código.

	palabra := "año"
	fmt.Println("\n=== Bytes vs letras ===")
	fmt.Println("palabra:", palabra)
	fmt.Println("len (bytes):", len(palabra))                                  // 4
	fmt.Println("RuneCountInString (runes):", utf8.RuneCountInString(palabra)) // 3

	frase := "el niño tomó café ☕"
	fmt.Println("\nfrase:", frase)
	fmt.Println("bytes: ", len(frase))                    // más que las letras
	fmt.Println("runes: ", utf8.RuneCountInString(frase)) // puntos de código

	// Se ven parecidas, pero Go no normaliza automáticamente el texto.
	// "é" es U+00E9; "e\u0301" combina e con un acento U+0301.
	compuesta, combinada := "é", "e\u0301"
	fmt.Println("\n=== Un carácter visible puede tener varias runes ===")
	fmt.Printf("é: bytes=%d runes=%d\n", len(compuesta), utf8.RuneCountInString(compuesta))          // 2, 1
	fmt.Printf("e + acento: bytes=%d runes=%d\n", len(combinada), utf8.RuneCountInString(combinada)) // 3, 2
	fmt.Println("¿Strings iguales?", compuesta == combinada)                                         // false: sus bytes difieren

	// Un string no valida sus bytes al crearse. range decodifica UTF-8 y
	// devuelve RuneError (U+FFFD) por cada byte de una codificación inválida.
	invalido := string([]byte{0xff, 'A'})
	fmt.Println("\n=== Bytes que no forman UTF-8 válido ===")
	fmt.Println("UTF-8 válido:", utf8.ValidString(invalido)) // false
	for i, r := range invalido {
		fmt.Printf("byte %d → %U\n", i, r) // U+FFFD, U+0041
	}

	// ─────────────────────────────────────────────────────────
	// VER LOS BYTES REALES DE UN STRING
	// ─────────────────────────────────────────────────────────
	// "ñ" en UTF-8 se guarda como DOS bytes: 195 y 177.
	// Por eso len("ñ") da 2.

	fmt.Println("\n=== Los bytes de 'niño' ===")
	s := "niño"
	for i := 0; i < len(s); i++ {
		fmt.Printf("  byte %d: %d\n", i, s[i])
	}
	// n=110, i=105, ñ=195+177 (¡dos bytes!), o=111

	// ─────────────────────────────────────────────────────────
	// CONVERSIONES rune ↔ string
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Conversiones ===")

	// rune → string: string(rune) SÍ funciona como esperás
	r := 'ñ'
	fmt.Println("string('ñ'):", string(r)) // "ñ"

	// string → []rune: cada letra es un elemento
	runes := []rune("café")
	fmt.Println(`[]rune("café"):`, runes)      // [99 97 102 233]
	fmt.Printf("última letra: %c\n", runes[3]) // é (acceso por LETRA)

	// []rune → string: vuelta atrás
	fmt.Println("de vuelta a string:", string(runes)) // "café"

	// ─────────────────────────────────────────────────────────
	// CUIDADO: string(número) NO convierte a texto
	// ─────────────────────────────────────────────────────────
	// string(65) NO da "65": da el carácter con código 65 → "A".
	// Para convertir números a texto está strconv (14_strconv_conversiones).

	fmt.Println("\n=== Trampa: string(número) ===")
	fmt.Println("string(65):", string(rune(65)))   // "A" (¡no "65"!)
	fmt.Println("string(241):", string(rune(241))) // "ñ"

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Unicode     → tabla: cada carácter tiene un número (code point)")
	fmt.Println("  UTF-8       → cómo se guardan esos números (1 a 4 bytes)")
	fmt.Println("  rune        → un code point; alias de int32; comillas simples")
	fmt.Println("  byte        → un byte crudo; alias de uint8")
	fmt.Println("  len(s)                    → BYTES")
	fmt.Println("  utf8.RuneCountInString(s) → PUNTOS DE CÓDIGO, no grafemas")
	fmt.Println("  []rune(s)   → para trabajar punto de código por punto de código")
	fmt.Println("  utf8.ValidString(s)      → verifica si los bytes son UTF-8 válido")
	fmt.Println("  string(65)  → \"A\", NO \"65\" (para eso, strconv)")
}
