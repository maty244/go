package main

import "fmt"

// =========================================================
// ¿QUÉ ES UN STRING EN GO?
// =========================================================
// Un string es una secuencia INMUTABLE de bytes.
// Dos cosas importantes de esa definición:
//
//   1. BYTES, no "letras": Go guarda el texto codificado en
//      UTF-8. Una letra como "a" ocupa 1 byte, pero "ñ" ocupa 2.
//      (Esto lo vemos a fondo en el tema 02).
//
//   2. INMUTABLE: una vez creado, un string NO se puede modificar.
//      Toda "modificación" en realidad crea un string nuevo.
//      (Esto lo vemos a fondo en el tema 04).
//
// Internamente un string es muy parecido a un slice:
// un puntero a los bytes + una longitud. Por eso copiar
// un string es barato: se copia el "encabezado", no el texto.

func main() {
	// ─────────────────────────────────────────────────────────
	// DECLARACIÓN: comillas dobles
	// ─────────────────────────────────────────────────────────
	// Los strings se escriben con comillas DOBLES.
	// Las comillas simples son para runes (un solo carácter),
	// NO para strings: 'hola' no compila, 'h' sí.

	nombre := "Matias"
	var ciudad string = "Buenos Aires"
	var vacio string // zero value: "" (string vacío, no nil)

	fmt.Println("=== Declaración ===")
	fmt.Println("nombre:", nombre)
	fmt.Println("ciudad:", ciudad)
	fmt.Printf("vacio: %q (len=%d)\n", vacio, len(vacio))

	// ─────────────────────────────────────────────────────────
	// LEN(): cuenta BYTES, no letras
	// ─────────────────────────────────────────────────────────
	// Esta es LA trampa número uno con strings en Go.
	// Con texto en inglés no se nota (1 letra = 1 byte),
	// pero con tildes y eñes los números no coinciden.

	fmt.Println("\n=== len() cuenta bytes ===")
	fmt.Println(`len("hola") =`, len("hola")) // 4 → 4 letras, 4 bytes
	fmt.Println(`len("niño") =`, len("niño")) // 5 → 4 letras, pero ñ ocupa 2 bytes
	fmt.Println(`len("café") =`, len("café")) // 5 → 4 letras, é ocupa 2 bytes

	// ─────────────────────────────────────────────────────────
	// CARACTERES DE ESCAPE
	// ─────────────────────────────────────────────────────────
	// Dentro de comillas dobles, la barra invertida \ tiene
	// significado especial:
	//   \n  salto de línea
	//   \t  tabulación
	//   \"  comilla doble (para no cerrar el string)
	//   \\  barra invertida literal

	fmt.Println("\n=== Caracteres de escape ===")
	fmt.Println("línea 1\nlínea 2")
	fmt.Println("columna1\tcolumna2")
	fmt.Println("ella dijo \"hola\"")
	fmt.Println("ruta: C:\\Users\\Matias")

	// ─────────────────────────────────────────────────────────
	// RAW STRINGS: backticks ` `
	// ─────────────────────────────────────────────────────────
	// Con backticks, NADA se interpreta: ni \n, ni \t, ni comillas.
	// Lo que escribís es exactamente lo que obtenés.
	// Útil para rutas de Windows, JSON, expresiones regulares,
	// o texto de varias líneas.

	fmt.Println("\n=== Raw strings (backticks) ===")

	ruta := `C:\Users\Matias\go` // no hace falta escapar las \
	fmt.Println("ruta:", ruta)

	crudo := `esto NO es un salto: \n  y esto no es tab: \t`
	fmt.Println(crudo)

	// Raw string multilínea: los saltos de línea reales SÍ cuentan
	menu := `MENÚ DEL DÍA
  1. Milanesa con puré
  2. Empanadas
  3. Asado`
	fmt.Println(menu)

	// ─────────────────────────────────────────────────────────
	// STRING vs []BYTE
	// ─────────────────────────────────────────────────────────
	// Un string se puede convertir a []byte y viceversa.
	// La diferencia clave:
	//   string  → inmutable
	//   []byte  → mutable (es un slice común)
	// La conversión COPIA los datos (por eso el string original
	// queda protegido aunque modifiques el slice).

	fmt.Println("\n=== string vs []byte ===")

	s := "gol"
	b := []byte(s)                           // copia los bytes a un slice mutable
	fmt.Println("string:", s, "| bytes:", b) // [103 111 108] (códigos UTF-8)

	b[0] = 'G'                                  // el slice SÍ se puede modificar
	fmt.Println("slice modificado:", string(b)) // "Gol"
	fmt.Println("string original intacto:", s)  // "gol" (no cambió)

	// ─────────────────────────────────────────────────────────
	// CONCATENACIÓN BÁSICA con +
	// ─────────────────────────────────────────────────────────
	// (La forma eficiente para muchos pedazos la vemos en el tema 05)

	saludo := "Hola, " + nombre + "!"
	fmt.Println("\n=== Concatenación ===")
	fmt.Println(saludo)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  string         → secuencia INMUTABLE de bytes (UTF-8)`)
	fmt.Println(`  "texto"        → string normal (interpreta \n, \t, etc.)`)
	fmt.Println("  `texto`        → raw string (todo literal, sirve multilínea)")
	fmt.Println(`  'a'            → rune (un carácter), NO un string`)
	fmt.Println(`  len(s)         → cantidad de BYTES, no de letras`)
	fmt.Println(`  []byte(s)      → copia mutable de los bytes`)
	fmt.Println(`  zero value     → "" (string vacío)`)
}
