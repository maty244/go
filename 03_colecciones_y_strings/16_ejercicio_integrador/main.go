package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// =========================================================
// EJERCICIO INTEGRADOR: PROCESADOR DE TEXTO Y FICHAS
// =========================================================
// Acá juntamos TODO lo de la sección Strings con lo que ya
// sabés de la biblia: maps, slices, structs, funciones,
// múltiples retornos y manejo de errores.
//
// Construimos dos mini-herramientas del mundo real:
//
//   PARTE A: analizador de texto
//     - contar palabras (Fields)
//     - frecuencia de cada palabra (map + normalización)
//     - top de palabras más usadas (sort.Slice)
//
//   PARTE B: parser de fichas "nombre,edad,ciudad"
//     - partir con Split
//     - limpiar con TrimSpace
//     - convertir la edad con strconv.Atoi (con error)
//     - capitalizar el nombre ([]rune + ToUpper)
//     - devolver (Persona, error) como en 06_interfaces_y_errores/12_patron_de_errores

// ─────────────────────────────────────────────────────────
// PARTE A: ANALIZADOR DE TEXTO
// ─────────────────────────────────────────────────────────

// normalizar deja una palabra lista para comparar:
// minúsculas y sin signos de puntuación en los bordes.
// "¡Hola," y "hola" deben contar como LA MISMA palabra.
func normalizar(palabra string) string {
	p := strings.ToLower(palabra)
	p = strings.Trim(p, ".,;:!?¡¿\"()") // saca puntuación de los bordes
	return p
}

// contarFrecuencias devuelve un map palabra → cantidad.
// Patrón clásico: map[string]int como contador (03_colecciones_y_strings/04_maps).
func contarFrecuencias(texto string) map[string]int {
	frecuencia := make(map[string]int)
	for _, palabra := range strings.Fields(texto) {
		p := normalizar(palabra)
		if p == "" {
			continue // era pura puntuación
		}
		frecuencia[p]++ // si no existe, arranca en 0 (zero value)
	}
	return frecuencia
}

// topPalabras devuelve las n palabras más frecuentes.
// Los maps NO tienen orden → pasamos a slice y ordenamos.
func topPalabras(frecuencia map[string]int, n int) []string {
	// 1. juntar las claves en un slice
	palabras := make([]string, 0, len(frecuencia))
	for p := range frecuencia {
		palabras = append(palabras, p)
	}

	// 2. ordenar por frecuencia (mayor primero);
	//    a igual frecuencia, alfabético para que sea estable
	sort.Slice(palabras, func(i, j int) bool {
		fi, fj := frecuencia[palabras[i]], frecuencia[palabras[j]]
		if fi != fj {
			return fi > fj
		}
		return palabras[i] < palabras[j]
	})

	// 3. recortar al top n (cuidando no pasarnos)
	if n > len(palabras) {
		n = len(palabras)
	}
	return palabras[:n]
}

// ─────────────────────────────────────────────────────────
// PARTE B: PARSER DE FICHAS "nombre,edad,ciudad"
// ─────────────────────────────────────────────────────────

type Persona struct {
	Nombre string
	Edad   int
	Ciudad string
}

// capitalizar pone la primera letra en mayúscula.
// Usamos []rune para que funcione con tildes y eñes
// ("ángela" → "Ángela", no basura).
func capitalizar(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	primera := strings.ToUpper(string(runes[0]))
	return primera + string(runes[1:])
}

// parsearFicha convierte una línea "nombre, edad, ciudad"
// en una Persona. Devuelve (Persona, error): si la línea
// está mal, el error explica POR QUÉ (06_interfaces_y_errores/12_patron_de_errores).
func parsearFicha(linea string) (Persona, error) {
	campos := strings.Split(linea, ",")
	if len(campos) != 3 {
		return Persona{}, fmt.Errorf("se esperaban 3 campos, llegaron %d: %q", len(campos), linea)
	}

	// limpiar espacios de cada campo
	nombre := strings.TrimSpace(campos[0])
	edadTexto := strings.TrimSpace(campos[1])
	ciudad := strings.TrimSpace(campos[2])

	if nombre == "" {
		return Persona{}, fmt.Errorf("el nombre está vacío en: %q", linea)
	}

	// la edad llega como texto → strconv.Atoi (puede fallar)
	edad, err := strconv.Atoi(edadTexto)
	if err != nil {
		return Persona{}, fmt.Errorf("edad inválida %q: no es un número", edadTexto)
	}
	if edad < 0 || edad > 130 {
		return Persona{}, fmt.Errorf("edad fuera de rango: %d", edad)
	}

	return Persona{
		Nombre: capitalizar(strings.ToLower(nombre)),
		Edad:   edad,
		Ciudad: capitalizar(strings.ToLower(ciudad)),
	}, nil
}

// ─────────────────────────────────────────────────────────
// MAIN: usar las dos herramientas
// ─────────────────────────────────────────────────────────

func main() {
	// ═════════════════════════════════════════════════════
	// PARTE A en acción
	// ═════════════════════════════════════════════════════
	fmt.Println("=== PARTE A: analizador de texto ===")

	texto := `Go es simple. Go es rápido y Go es divertido.
	¡Aprender Go es la mejor decisión! Simple de leer, simple de escribir.`

	fmt.Println("texto:", strings.TrimSpace(texto))

	palabras := strings.Fields(texto)
	fmt.Println("\ntotal de palabras:", len(palabras))

	frecuencia := contarFrecuencias(texto)
	fmt.Println("palabras distintas:", len(frecuencia))

	fmt.Println("\ntop 3 más usadas:")
	for i, p := range topPalabras(frecuencia, 3) {
		fmt.Printf("  %d. %q → %d veces\n", i+1, p, frecuencia[p])
	}
	// "go" aparece 4 veces aunque en el texto está como "Go":
	// la normalización (ToLower + Trim de puntuación) hizo su trabajo.

	// ═════════════════════════════════════════════════════
	// PARTE B en acción
	// ═════════════════════════════════════════════════════
	fmt.Println("\n=== PARTE B: parser de fichas ===")

	// Líneas como llegarían de un archivo CSV: algunas bien,
	// otras con problemas. El programa NO se cae: reporta y sigue.
	lineas := []string{
		"  matias , 28 , buenos aires ",
		"ÁNGELA,35,córdoba",
		"luis,treinta,rosario", // edad no numérica
		"ana,29",               // faltan campos
		", 40, mendoza",        // nombre vacío
		"sofía, 31, ushuaia",
	}

	var personas []Persona
	for _, linea := range lineas {
		p, err := parsearFicha(linea)
		if err != nil {
			fmt.Println("  ✗ descartada:", err)
			continue
		}
		personas = append(personas, p)
		fmt.Printf("  ✓ %s (%d) de %s\n", p.Nombre, p.Edad, p.Ciudad)
	}

	fmt.Printf("\ncargadas %d de %d fichas\n", len(personas), len(lineas))

	// Broche final: ordenar las personas por nombre
	sort.Slice(personas, func(i, j int) bool {
		return personas[i].Nombre < personas[j].Nombre
	})
	fmt.Println("\nordenadas por nombre:")
	for _, p := range personas {
		fmt.Printf("  - %-8s %3d años, %s\n", p.Nombre, p.Edad, p.Ciudad)
	}

	// ─────────────────────────────────────────────────────────
	// QUÉ USAMOS DE LA BIBLIA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Qué usamos ===")
	fmt.Println("  strings.Fields/Split/Trim/ToLower → Strings 06-07")
	fmt.Println("  strconv.Atoi con error            → Strings 08")
	fmt.Println("  []rune para capitalizar           → Strings 03-04")
	fmt.Println("  map como contador                 → Fundamentos 17/26")
	fmt.Println("  struct Persona                    → Fundamentos 18")
	fmt.Println("  (valor, error) y validación       → Fundamentos 37-38")
	fmt.Println("  sort.Slice con criterio propio    → Strings 09")
}
