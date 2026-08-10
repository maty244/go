package main

import (
	"fmt"
	"strings"
)

// =========================================================
// INMUTABILIDAD DE STRINGS
// =========================================================
// Un string en Go NO se puede modificar. Nunca. Punto.
//
//   s := "hola"
//   s[0] = 'H'   // ❌ NO COMPILA: cannot assign to s[0]
//
// ¿Por qué Go decidió esto?
//   1. SEGURIDAD: si pasás un string a una función, tenés
//      garantía de que no te lo van a cambiar.
//   2. EFICIENCIA: como nadie puede modificarlo, varios
//      strings pueden COMPARTIR los mismos bytes en memoria
//      sin riesgo. Copiar un string es copiar un encabezado
//      chiquito (puntero + len), no todo el texto.
//   3. Se pueden usar como claves de map sin sorpresas.
//
// Conexión con Punteros: pasar un string por valor es barato
// por lo mismo que un slice: viaja el encabezado, no los datos.
// La diferencia es que el string NO permite mutar esos datos.
//
// "Modificar" un string siempre significa CREAR UNO NUEVO.

func main() {
	// ─────────────────────────────────────────────────────────
	// REASIGNAR ≠ MODIFICAR
	// ─────────────────────────────────────────────────────────
	// Esto es válido, pero no "modifica" el string:
	// crea un string nuevo y la variable apunta al nuevo.

	s := "hola"
	fmt.Println("=== Reasignar crea un string nuevo ===")
	fmt.Println("antes: ", s)

	s = "chau" // s ahora apunta a OTRO string; "hola" quedó intacto
	fmt.Println("después:", s)

	// Lo que NO se puede (descomentá y mirá el error del compilador):
	// s[0] = 'C' // ❌ cannot assign to s[0] (strings are immutable)

	// ─────────────────────────────────────────────────────────
	// CÓMO "CAMBIAR UNA LETRA": vía []rune
	// ─────────────────────────────────────────────────────────
	// Receta de 3 pasos:
	//   1. convertir a []rune (esto COPIA los datos a un slice mutable)
	//   2. modificar el slice
	//   3. convertir de vuelta a string (otra copia)

	fmt.Println("\n=== Cambiar una letra (vía []rune) ===")
	nombre := "matias"
	runes := []rune(nombre)
	runes[0] = 'M' // ahora sí: es un slice, se puede
	nombreCapital := string(runes)

	fmt.Println("original:  ", nombre)        // matias (intacto)
	fmt.Println("modificado:", nombreCapital) // Matias

	// Con tildes funciona igual, porque []rune trabaja por LETRA
	pais := "perú"
	rp := []rune(pais)
	rp[0] = 'P'
	fmt.Println("perú →", string(rp)) // Perú

	// ─────────────────────────────────────────────────────────
	// VÍA []byte: solo si es ASCII puro
	// ─────────────────────────────────────────────────────────
	// []byte también sirve y es más barato, pero trabaja por BYTE:
	// con tildes o eñes podés romper la codificación UTF-8.

	fmt.Println("\n=== Vía []byte (solo ASCII) ===")
	código := "abc-123"
	bs := []byte(código)
	bs[3] = '_'
	fmt.Println(string(bs)) // abc_123

	// ─────────────────────────────────────────────────────────
	// LAS FUNCIONES DE strings TAMPOCO MODIFICAN: RETORNAN
	// ─────────────────────────────────────────────────────────
	// Todas las funciones del paquete strings devuelven un
	// string NUEVO. Si no asignás el resultado, no pasó nada.

	fmt.Println("\n=== Las funciones retornan, no modifican ===")
	texto := "hola mundo"

	// MAL: llamar y tirar el resultado
	strings.ToUpper(texto)                          // no hace nada visible
	fmt.Println("tras ToUpper sin asignar:", texto) // sigue en minúscula

	// BIEN: asignar el resultado
	mayus := strings.ToUpper(texto)
	fmt.Println("asignando el resultado:  ", mayus)

	// (Es el mismo patrón que append con slices:
	//  s = append(s, x) → siempre asignás el resultado)

	// ─────────────────────────────────────────────────────────
	// CONSECUENCIA: comparar y compartir es seguro
	// ─────────────────────────────────────────────────────────
	// Dos variables pueden compartir los mismos bytes sin peligro.

	fmt.Println("\n=== Compartir es seguro ===")
	original := "inmutable"
	copia := original      // comparten los bytes (no se copian)
	parte := original[2:5] // "mut" — también comparte memoria

	fmt.Println("original:", original)
	fmt.Println("copia:   ", copia)
	fmt.Println("parte:   ", parte)
	// No existe forma de que 'copia' o 'parte' rompan 'original'.

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  s[0] = 'X'         → ❌ no compila, strings inmutables")
	fmt.Println("  s = \"otro\"         → ✅ reasigna (string nuevo)")
	fmt.Println("  []rune(s) → mod → string() → ✅ cambiar letras (UTF-8 seguro)")
	fmt.Println("  []byte(s)          → ✅ solo si es ASCII puro")
	fmt.Println("  strings.ToUpper(s) → retorna NUEVO; hay que asignarlo")
	fmt.Println("  Compartir/copiar strings es barato y seguro.")
}
