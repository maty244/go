package main

import "fmt"

// =========================================================
// ¿QUÉ ES UN CLOSURE?
// =========================================================
// Un closure es una función anónima que "RECUERDA" las variables
// del lugar donde nació, incluso después de que ese lugar (la
// función que la creó) ya terminó de ejecutarse.
//
// La palabra viene de "cerrar sobre" (close over) esas variables:
// la función anónima las captura y las mantiene vivas mientras
// ella misma exista.
//
// Es la diferencia entre una función anónima cualquiera (tema 01)
// y un closure: el closure usa una variable de AFUERA de su propio
// cuerpo, y esa variable persiste en memoria por él.

// ─────────────────────────────────────────────────────────
// EJEMPLO MÍNIMO: capturar una variable externa
// ─────────────────────────────────────────────────────────

func crearSaludador(nombre string) func() string {
	// "nombre" es un parámetro de crearSaludador.
	// La función que retornamos usa "nombre" en su cuerpo:
	// eso la convierte en un closure.
	return func() string {
		return "Hola, " + nombre
	}
}

func main() {
	fmt.Println("=== Closure básico ===")

	saludarMati := crearSaludador("Matias")
	saludarAna := crearSaludador("Ana")

	// crearSaludador YA TERMINÓ de ejecutarse en ambos casos,
	// pero cada función devuelta "recuerda" su propio nombre.
	fmt.Println(saludarMati())
	fmt.Println(saludarAna())

	// ─────────────────────────────────────────────────────────
	// CADA LLAMADA CREA UN CLOSURE INDEPENDIENTE
	// ─────────────────────────────────────────────────────────
	// saludarMati y saludarAna NO comparten memoria: cada llamada
	// a crearSaludador() generó su propia copia de "nombre".

	fmt.Println("\n=== Closures independientes ===")
	fmt.Println("saludarMati sigue diciendo:", saludarMati())
	fmt.Println("saludarAna sigue diciendo:", saludarAna())

	// ─────────────────────────────────────────────────────────
	// UN CLOSURE PUEDE MODIFICAR LA VARIABLE CAPTURADA
	// ─────────────────────────────────────────────────────────
	// No solo LEE la variable externa: también puede MODIFICARLA,
	// y ese cambio persiste entre llamadas. Esto es lo que permite
	// que un closure tenga "estado" (lo profundizamos en el tema 03).

	fmt.Println("\n=== Closure que modifica su variable capturada ===")

	sumador := crearSumador()
	fmt.Println("sumador() =", sumador()) // 1
	fmt.Println("sumador() =", sumador()) // 2
	fmt.Println("sumador() =", sumador()) // 3

	otroSumador := crearSumador()                 // OTRO closure, con su propio contador
	fmt.Println("otroSumador() =", otroSumador()) // 1 (independiente del anterior)
	fmt.Println("sumador() =", sumador())         // 4 (el original sigue su cuenta)

	// ─────────────────────────────────────────────────────────
	// POR QUÉ FUNCIONA: la vida de la variable se extiende
	// ─────────────────────────────────────────────────────────
	// La semántica del lenguaje mantiene viva la variable mientras el closure la
	// necesite. El compilador puede detectar que escapa, aplicar inlining u otras
	// optimizaciones y elegir su representación; no conviene prometer una
	// ubicación física específica como parte del comportamiento del programa.

	// ─────────────────────────────────────────────────────────
	// CASO REAL: generador de IDs únicos
	// ─────────────────────────────────────────────────────────
	// Un closure es perfecto para mantener un contador sin usar
	// una variable global (que cualquier parte del código podría
	// modificar por error).

	fmt.Println("\n=== Caso real: generador de IDs de pedido ===")

	siguienteID := crearGeneradorID("PED")
	fmt.Println(siguienteID()) // PED-001
	fmt.Println(siguienteID()) // PED-002
	fmt.Println(siguienteID()) // PED-003

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Closure          → función que 'recuerda' variables de afuera")
	fmt.Println("  Cada llamada a   → crea un closure NUEVO e independiente")
	fmt.Println("  la función que")
	fmt.Println("  lo genera")
	fmt.Println("  Puede modificar  → la variable capturada, y el cambio persiste")
	fmt.Println("  Por qué funciona → Go extiende la vida de la variable capturada")
}

func crearSumador() func() int {
	total := 0
	return func() int {
		total++
		return total
	}
}

func crearGeneradorID(prefijo string) func() string {
	contador := 0
	return func() string {
		contador++
		return fmt.Sprintf("%s-%03d", prefijo, contador)
	}
}
