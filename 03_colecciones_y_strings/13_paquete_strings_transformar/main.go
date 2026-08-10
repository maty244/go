package main

import (
	"fmt"
	"strings"
)

// =========================================================
// PAQUETE strings (2): TRANSFORMAR Y PARTIR
// =========================================================
// La segunda mitad del paquete strings: funciones que toman
// un string y devuelven uno NUEVO transformado (recordá:
// los strings son inmutables, nada se modifica "in place").
//
//   ToUpper / ToLower        → cambiar mayúsculas/minúsculas
//   Split / Fields           → partir en pedazos ([]string)
//   Join                     → unir pedazos (la inversa de Split)
//   TrimSpace / Trim / ...   → sacar basura de los bordes
//   Replace / ReplaceAll     → reemplazar apariciones
//
// Combo clásico del mundo real: limpiar input del usuario con
// TrimSpace + ToLower antes de comparar o guardar.

func main() {
	// ─────────────────────────────────────────────────────────
	// ToUpper / ToLower
	// ─────────────────────────────────────────────────────────
	// Manejan bien los caracteres del español (á→Á, ñ→Ñ).

	fmt.Println("=== ToUpper / ToLower ===")
	grito := "¡peligro! camión sin frenos"
	fmt.Println(strings.ToUpper(grito)) // ¡PELIGRO! CAMIÓN SIN FRENOS
	fmt.Println(strings.ToLower("DEJÁ DE GRITAR ÑOQUI"))

	// ─────────────────────────────────────────────────────────
	// Split: partir por un separador → []string
	// ─────────────────────────────────────────────────────────
	// Devuelve un slice con los pedazos. El separador desaparece.

	fmt.Println("\n=== Split ===")
	csv := "matias,28,buenos aires"
	campos := strings.Split(csv, ",")
	fmt.Printf("%q\n", campos) // ["matias" "28" "buenos aires"]
	fmt.Println("cantidad de campos:", len(campos))
	fmt.Println("segundo campo:", campos[1]) // 28

	fecha := "10/06/2026"
	partes := strings.Split(fecha, "/")
	fmt.Printf("día=%s mes=%s año=%s\n", partes[0], partes[1], partes[2])

	// Detalle: si el separador no está, devuelve 1 pedazo (todo el string)
	fmt.Printf("sin separador: %q\n", strings.Split("hola", ","))

	// ─────────────────────────────────────────────────────────
	// Fields: partir por espacios (la forma robusta)
	// ─────────────────────────────────────────────────────────
	// Split(s, " ") se rompe con dobles espacios; Fields parte
	// por CUALQUIER cantidad de espacios/tabs y descarta vacíos.

	fmt.Println("\n=== Fields vs Split(\" \") ===")
	sucio := "  hola   mundo  cruel "
	fmt.Printf("Split:  %q\n", strings.Split(sucio, " ")) // con pedazos vacíos ""
	fmt.Printf("Fields: %q\n", strings.Fields(sucio))     // ["hola" "mundo" "cruel"]

	// Uso típico: contar palabras
	oracion := "el perro de san roque no tiene rabo"
	fmt.Println("palabras:", len(strings.Fields(oracion))) // 8

	// ─────────────────────────────────────────────────────────
	// Join: la inversa de Split
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Split + Join (ida y vuelta) ===")
	ruta := "home/matias/proyectos/golang"
	dirs := strings.Split(ruta, "/")
	fmt.Printf("partes: %q\n", dirs)
	fmt.Println("unido con ' > ':", strings.Join(dirs, " > "))

	// ─────────────────────────────────────────────────────────
	// TrimSpace y familia Trim: limpiar los bordes
	// ─────────────────────────────────────────────────────────
	//   TrimSpace(s)        → saca espacios/tabs/saltos de AMBOS lados
	//   Trim(s, "xy")       → saca los caracteres x e y de ambos lados
	//   TrimLeft / TrimRight→ solo un lado
	//   TrimPrefix(s, pre)  → saca pre del inicio (si está)
	//   TrimSuffix(s, suf)  → saca suf del final (si está)
	// Solo tocan los BORDES, nunca el medio.

	fmt.Println("\n=== TrimSpace y familia ===")
	input := "   matias@gmail.com  \n"
	fmt.Printf("crudo:    %q\n", input)
	fmt.Printf("limpio:   %q\n", strings.TrimSpace(input))

	fmt.Println(strings.Trim("***oferta***", "*"))                // oferta
	fmt.Println(strings.TrimPrefix("https://go.dev", "https://")) // go.dev
	fmt.Println(strings.TrimSuffix("informe.pdf", ".pdf"))        // informe

	// TrimPrefix con un prefijo que NO está: devuelve el string igual
	fmt.Println(strings.TrimPrefix("go.dev", "https://")) // go.dev (sin cambios)

	// ─────────────────────────────────────────────────────────
	// Replace / ReplaceAll: reemplazar
	// ─────────────────────────────────────────────────────────
	//   ReplaceAll(s, viejo, nuevo) → reemplaza TODAS
	//   Replace(s, viejo, nuevo, n) → reemplaza las primeras n (-1 = todas)

	fmt.Println("\n=== Replace / ReplaceAll ===")
	texto := "me gusta java, java es lo mejor, viva java"
	fmt.Println(strings.ReplaceAll(texto, "java", "go"))
	fmt.Println(strings.Replace(texto, "java", "go", 1)) // solo la primera

	// Uso típico: normalizar separadores
	tel := "11-5555-4444"
	fmt.Println("tel sin guiones:", strings.ReplaceAll(tel, "-", ""))

	// ─────────────────────────────────────────────────────────
	// COMBO REAL: normalizar input del usuario
	// ─────────────────────────────────────────────────────────
	// Lo que tipea un usuario llega con espacios y mayúsculas
	// aleatorias. Antes de comparar o guardar: limpiar.

	fmt.Println("\n=== Combo: normalizar input ===")
	respuestas := []string{"  SI ", "si", "Si  ", " NO", "nO "}
	for _, r := range respuestas {
		norm := strings.ToLower(strings.TrimSpace(r))
		fmt.Printf("%q → %q → ¿es sí?: %v\n", r, norm, norm == "si")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  ToUpper/ToLower      → MAYÚS/minús (maneja ñ y tildes)")
	fmt.Println("  Split(s, sep)        → []string partido por sep")
	fmt.Println("  Fields(s)            → []string partido por espacios (robusto)")
	fmt.Println("  Join(slice, sep)     → la inversa de Split")
	fmt.Println("  TrimSpace(s)         → limpia espacios de los bordes")
	fmt.Println("  TrimPrefix/Suffix    → saca un prefijo/sufijo exacto")
	fmt.Println("  ReplaceAll(s, v, n)  → reemplaza todas las apariciones")
	fmt.Println("  Todas RETORNAN un string nuevo: hay que asignar.")
}
