package main

import (
	"fmt"
	"strings"
)

// =========================================================
// PAQUETE strings (1): BUSCAR Y PREGUNTAR
// =========================================================
// El paquete strings de la biblioteca estándar trae todas
// las herramientas para trabajar con texto. En este tema:
// las funciones que PREGUNTAN cosas sobre un string.
//
//   strings.Contains(s, sub)    → ¿s contiene sub?         (bool)
//   strings.HasPrefix(s, pre)   → ¿s empieza con pre?      (bool)
//   strings.HasSuffix(s, suf)   → ¿s termina con suf?      (bool)
//   strings.Index(s, sub)       → ¿dónde está sub?         (int, -1 si no está)
//   strings.LastIndex(s, sub)   → ¿dónde está la última?   (int)
//   strings.Count(s, sub)       → ¿cuántas veces aparece?  (int)
//   strings.EqualFold(a, b)     → ¿iguales sin importar mayúsculas? (bool)
//
// Todas distinguen mayúsculas de minúsculas ("Go" ≠ "go"),
// salvo EqualFold, que existe justamente para ignorarlas.

func main() {
	frase := "el que quiera celeste, que le cueste"

	// ─────────────────────────────────────────────────────────
	// Contains: ¿contiene este pedazo?
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== strings.Contains ===")
	fmt.Println("frase:", frase)
	fmt.Println(`¿contiene "celeste"?`, strings.Contains(frase, "celeste")) // true
	fmt.Println(`¿contiene "verde"?  `, strings.Contains(frase, "verde"))   // false
	fmt.Println(`¿contiene "CELESTE"?`, strings.Contains(frase, "CELESTE")) // false (¡mayúsculas!)

	// Uso típico: filtrar
	emails := []string{"matias@gmail.com", "ana@hotmail.com", "luis@gmail.com"}
	fmt.Println("\nemails de gmail:")
	for _, e := range emails {
		if strings.Contains(e, "@gmail.") {
			fmt.Println("  -", e)
		}
	}

	// ─────────────────────────────────────────────────────────
	// HasPrefix / HasSuffix: ¿empieza/termina con...?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== HasPrefix / HasSuffix ===")

	url := "https://go.dev/doc"
	archivo := "informe_ventas.pdf"

	fmt.Println(`¿url empieza con "https://"?`, strings.HasPrefix(url, "https://"))  // true
	fmt.Println(`¿archivo termina en ".pdf"?  `, strings.HasSuffix(archivo, ".pdf")) // true
	fmt.Println(`¿archivo termina en ".go"?   `, strings.HasSuffix(archivo, ".go"))  // false

	// Uso típico: clasificar archivos
	archivos := []string{"main.go", "notas.txt", "utils.go", "foto.png"}
	fmt.Println("\narchivos de Go:")
	for _, a := range archivos {
		if strings.HasSuffix(a, ".go") {
			fmt.Println("  -", a)
		}
	}

	// ─────────────────────────────────────────────────────────
	// Index / LastIndex: ¿en qué posición está?
	// ─────────────────────────────────────────────────────────
	// Devuelven el índice del BYTE donde empieza la primera
	// (o última) aparición. Si no está, devuelven -1.
	// Ese -1 se usa como "no encontrado" (no hay error ni nil).

	fmt.Println("\n=== Index / LastIndex ===")
	s := "go es go y siempre será go"
	fmt.Println("s:", s)
	fmt.Println(`Index(s, "go")     =`, strings.Index(s, "go"))     // 0 (primera)
	fmt.Println(`LastIndex(s, "go") =`, strings.LastIndex(s, "go")) // última
	fmt.Println(`Index(s, "rust")   =`, strings.Index(s, "rust"))   // -1 (no está)

	// Patrón típico: chequear -1 antes de usar el índice
	if i := strings.Index(s, "siempre"); i != -1 {
		fmt.Printf(`"siempre" encontrado en el byte %d`+"\n", i)
	}

	// Uso práctico: separar usuario y dominio de un email
	email := "matias@gmail.com"
	arroba := strings.Index(email, "@")
	fmt.Println("usuario:", email[:arroba])   // matias
	fmt.Println("dominio:", email[arroba+1:]) // gmail.com

	// ─────────────────────────────────────────────────────────
	// Count: ¿cuántas veces aparece?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Count ===")
	texto := "banana"
	fmt.Println(`Count("banana", "a") =`, strings.Count(texto, "a"))   // 3
	fmt.Println(`Count("banana", "na") =`, strings.Count(texto, "na")) // 2
	// Caso especial: contar string vacío da runes+1 (curiosidad, no se usa)
	fmt.Println(`Count("banana", "") =`, strings.Count(texto, "")) // 7

	// ─────────────────────────────────────────────────────────
	// EqualFold: comparar SIN distinguir mayúsculas
	// ─────────────────────────────────────────────────────────
	// MAL:  s1 == s2 falla si difieren las mayúsculas
	// BIEN: EqualFold compara "doblado" (case-insensitive)
	// Es la forma correcta de comparar input del usuario.

	fmt.Println("\n=== EqualFold ===")
	respuesta := "SI"
	fmt.Println(`respuesta == "si"        →`, respuesta == "si")                   // false
	fmt.Println(`EqualFold(respuesta,"si") →`, strings.EqualFold(respuesta, "si")) // true
	fmt.Println(`EqualFold("Go", "GO")     →`, strings.EqualFold("Go", "GO"))      // true

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Contains(s, sub)   → bool: ¿está adentro?")
	fmt.Println("  HasPrefix(s, pre)  → bool: ¿empieza con...?")
	fmt.Println("  HasSuffix(s, suf)  → bool: ¿termina con...?")
	fmt.Println("  Index(s, sub)      → int: posición o -1")
	fmt.Println("  LastIndex(s, sub)  → int: última posición o -1")
	fmt.Println("  Count(s, sub)      → int: cantidad de apariciones")
	fmt.Println("  EqualFold(a, b)    → bool: iguales ignorando mayúsculas")
	fmt.Println("  Todas case-sensitive, salvo EqualFold.")
}
