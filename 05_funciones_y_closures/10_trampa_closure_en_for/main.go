package main

import "fmt"

// =========================================================
// LA TRAMPA CLÁSICA: CLOSURES DENTRO DE UN FOR
// =========================================================
// Este archivo documenta el bug más famoso de Go con closures...
// que HOY YA NO EXISTE por defecto. Te lo mostramos igual porque:
//   1. Vas a encontrarlo en código viejo, tutoriales y Stack
//      Overflow (todo lo escrito antes de Go 1.22).
//   2. Entender POR QUÉ pasaba te ayuda a entender mejor cómo
//      funcionan los closures y las variables de un for.
//
// ANTES de Go 1.22: la variable del for (el "i" de "for i := ...")
// era UNA SOLA variable, reusada en cada vuelta del loop. Si un
// closure la capturaba, capturaba ESA variable compartida, no
// "una copia por vuelta". Entonces todos los closures terminaban
// viendo el ÚLTIMO valor que tomó la variable.
//
// DESDE Go 1.22 (mediados 2024): cada vuelta del for crea una
// variable NUEVA. El bug de abajo simplemente no ocurre más.
// Este repo usa Go 1.26, así que en TU código esto ya está resuelto.

func main() {
	// ─────────────────────────────────────────────────────────
	// CÓMO SE VE HOY (Go 1.22+): funciona como uno esperaría
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Closures capturando la variable del for (Go 1.26) ===")

	var funciones []func()

	for i := 0; i < 3; i++ {
		funciones = append(funciones, func() {
			fmt.Println("  i =", i)
		})
	}

	for _, f := range funciones {
		f()
	}
	// Salida en Go 1.26: i = 0, i = 1, i = 2 (cada uno con SU valor)
	//
	// EN GO VIEJO (antes de 1.22), esto imprimía "i = 3" tres veces:
	// las tres funciones compartían la misma variable "i", y para
	// cuando se ejecutaban (después de que el for ya había terminado),
	// "i" ya valía 3.

	// ─────────────────────────────────────────────────────────
	// EL MISMO PATRÓN CON for...range
	// ─────────────────────────────────────────────────────────
	// Se aplica igual a range: antes de 1.22, la variable del range
	// también era compartida. Hoy, cada iteración tiene la suya.

	fmt.Println("\n=== Lo mismo con for...range ===")

	nombres := []string{"Ana", "Carlos", "Sofía"}
	var saludos []func() string

	for _, nombre := range nombres {
		saludos = append(saludos, func() string {
			return "Hola, " + nombre
		})
	}

	for _, s := range saludos {
		fmt.Println(" -", s())
	}
	// Salida en Go 1.26: Hola, Ana / Hola, Carlos / Hola, Sofía

	// ─────────────────────────────────────────────────────────
	// CÓMO SE SOLUCIONABA ANTES (por si lo ves en código viejo)
	// ─────────────────────────────────────────────────────────
	// El arreglo clásico era crear una variable NUEVA dentro del
	// loop, sombreando la del for (viste "variables sombra" en
	// 01_fundamentos/16_scope_y_shadowing). Hoy es innecesario, pero es 100% válido y
	// vas a verlo en repositorios que no se actualizaron:

	fmt.Println("\n=== Solución 'de la vieja escuela' (ya no hace falta) ===")

	var funcionesViejoEstilo []func()
	for i := 0; i < 3; i++ {
		i := i // sombra: crea una copia NUEVA de "i" en cada vuelta
		funcionesViejoEstilo = append(funcionesViejoEstilo, func() {
			fmt.Println("  i =", i)
		})
	}
	for _, f := range funcionesViejoEstilo {
		f()
	}

	// ─────────────────────────────────────────────────────────
	// CUÁNDO SIGUE IMPORTANDO ESTE TEMA
	// ─────────────────────────────────────────────────────────
	// Ojo con las goroutines (lo vas a ver en 12_concurrencia/): el
	// mismo concepto de "qué variable captura cada función" aplica
	// ahí, y aunque el bug del for está resuelto, entender la
	// captura de variables por closures sigue siendo esencial.

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Antes de Go 1.22  → la variable del for/range era UNA sola,")
	fmt.Println("                       compartida por todos los closures del loop")
	fmt.Println("  Desde Go 1.22     → cada vuelta tiene su PROPIA variable")
	fmt.Println("  En Go 1.26 (acá)  → el bug clásico ya NO ocurre por defecto")
	fmt.Println("  Igual conocelo    → vas a verlo en código y tutoriales viejos")
}
