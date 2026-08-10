package main

import (
	"fmt"
	"strconv"
)

// =========================================================
// strconv: CONVERTIR STRING ↔ NÚMERO
// =========================================================
// En 01_fundamentos/07_conversiones_de_tipo vimos conversiones de tipo: float64(x),
// int(y)... Eso funciona ENTRE NÚMEROS. Pero con strings:
//
//   string(65)      → "A" (¡el carácter 65, no el texto "65"!)
//   int("123")      → ❌ no compila
//
// Para texto ↔ número existe el paquete strconv
// ("string conversion"). Las cuatro estrellas:
//
//   strconv.Itoa(123)       → "123"          (Int TO Ascii)
//   strconv.Atoi("123")     → 123, error     (Ascii TO Int)
//   strconv.ParseFloat(...) → float64, error
//   strconv.ParseBool(...)  → bool, error
//
// ¿Por qué Atoi devuelve error y Itoa no?
//   - Número → texto: SIEMPRE se puede. Sin error.
//   - Texto → número: puede fallar ("hola" no es un número).
//     Por eso devuelven (valor, error), como vimos en
//     06_interfaces_y_errores/12_patron_de_errores.

func main() {
	// ─────────────────────────────────────────────────────────
	// LA TRAMPA: string(número)
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== La trampa de string(número) ===")
	fmt.Println("string(rune(65)):", string(rune(65))) // "A" — código Unicode 65
	fmt.Println("strconv.Itoa(65):", strconv.Itoa(65)) // "65" — lo que querías

	// ─────────────────────────────────────────────────────────
	// Itoa: int → string (nunca falla)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Itoa: número → texto ===")
	edad := 28
	mensaje := "tengo " + strconv.Itoa(edad) + " años"
	fmt.Println(mensaje)

	// Alternativa: fmt.Sprintf("%d", edad) hace lo mismo;
	// Itoa es más directo cuando es solo el número.

	// ─────────────────────────────────────────────────────────
	// Atoi: string → int (puede fallar → manejar el error)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Atoi: texto → número ===")

	// Caso feliz
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("número:", n, "| el doble:", n*2)
	}

	// Caso que falla: SIEMPRE chequear el error
	malo, err := strconv.Atoi("42abc")
	if err != nil {
		fmt.Println("falló como esperábamos:", err)
		fmt.Println("(el valor devuelto es el zero value:", malo, ")")
	}

	// Uso típico: procesar números que llegan como texto
	// (de un archivo, de la consola, de una web... SIEMPRE llegan como texto)
	precios := []string{"1500", "2300", "abc", "800"}
	total := 0
	fmt.Println("\nsumando precios:", precios)
	for _, p := range precios {
		valor, err := strconv.Atoi(p)
		if err != nil {
			fmt.Printf("  ignorando %q: no es un número\n", p)
			continue
		}
		total += valor
	}
	fmt.Println("total:", total) // 4600

	// ─────────────────────────────────────────────────────────
	// ParseFloat: string → float64
	// ─────────────────────────────────────────────────────────
	// El segundo argumento es la precisión en bits: usá 64.

	fmt.Println("\n=== ParseFloat ===")
	precio, err := strconv.ParseFloat("1999.99", 64)
	if err == nil {
		conIVA := precio * 1.21
		fmt.Printf("precio: %.2f | con IVA: %.2f\n", precio, conIVA)
	}

	// OJO: usa punto decimal, no coma (formato inglés)
	_, err = strconv.ParseFloat("1999,99", 64)
	fmt.Println("con coma falla:", err)

	// ─────────────────────────────────────────────────────────
	// ParseBool: string → bool
	// ─────────────────────────────────────────────────────────
	// Acepta: "true", "false", "1", "0", "t", "f", "T", "F",
	//         "TRUE", "FALSE", "True", "False"
	// Típico para leer configuraciones.

	fmt.Println("\n=== ParseBool ===")
	for _, s := range []string{"true", "1", "F", "si"} {
		b, err := strconv.ParseBool(s)
		if err != nil {
			fmt.Printf("  %q → error (no lo reconoce)\n", s)
		} else {
			fmt.Printf("  %q → %v\n", s, b)
		}
	}

	// ─────────────────────────────────────────────────────────
	// ParseInt: control fino (bases y tamaño)
	// ─────────────────────────────────────────────────────────
	// ParseInt(s, base, bits) → int64, error
	//   base: 10 normal, 2 binario, 16 hexadecimal, 0 = autodetectar
	//   bits: 64 (después podés convertir con int())
	// Atoi("x") es un atajo de ParseInt("x", 10, 0).

	fmt.Println("\n=== ParseInt con bases ===")
	bin, _ := strconv.ParseInt("1010", 2, 64)
	hex, _ := strconv.ParseInt("ff", 16, 64)
	fmt.Println(`"1010" en base 2  →`, bin) // 10
	fmt.Println(`"ff" en base 16 →`, hex)   // 255

	// ─────────────────────────────────────────────────────────
	// FormatFloat: float → string con control
	// ─────────────────────────────────────────────────────────
	// FormatFloat(valor, formato, decimales, bits)
	//   formato 'f' = decimal normal | decimales -1 = los necesarios

	fmt.Println("\n=== FormatFloat ===")
	pi := 3.14159265
	fmt.Println(strconv.FormatFloat(pi, 'f', 2, 64))  // "3.14"
	fmt.Println(strconv.FormatFloat(pi, 'f', -1, 64)) // todos los dígitos
	// (fmt.Sprintf("%.2f", pi) es equivalente y más común)

	// ─────────────────────────────────────────────────────────
	// RESUMEN: tabla de conversiones texto ↔ número
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  QUIERO                       USO                         ¿ERROR?")
	fmt.Println("  ──────────────────────────   ─────────────────────────   ───────")
	fmt.Println("  int    → string              strconv.Itoa(n)             no")
	fmt.Println("  string → int                 strconv.Atoi(s)             sí")
	fmt.Println("  string → float64             strconv.ParseFloat(s, 64)   sí")
	fmt.Println("  string → bool                strconv.ParseBool(s)        sí")
	fmt.Println("  string → int (base 2/16)     strconv.ParseInt(s, b, 64)  sí")
	fmt.Println("  float  → string              strconv.FormatFloat(...)    no")
	fmt.Printf("  lo que sea → string          fmt.Sprintf(%q, x)        no\n", "%v")
	fmt.Println()
	fmt.Println("  Regla: texto → número SIEMPRE puede fallar → manejá el error.")
}
