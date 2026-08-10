package main

import "fmt"

// =========================================================
// CLOSURES CON ESTADO: EL PATRÓN "CONTADOR"
// =========================================================
// Profundizamos el ejemplo del tema 02: un closure puede actuar
// como un objeto con estado privado, sin necesidad de structs
// ni de exponer la variable interna. Esto es un patrón MUY usado
// en Go para encapsular datos sin escribir un struct completo.

// ─────────────────────────────────────────────────────────
// UN CLOSURE QUE EXPONE VARIAS OPERACIONES
// ─────────────────────────────────────────────────────────
// En vez de devolver UNA sola función, podemos devolver VARIAS,
// todas compartiendo la MISMA variable capturada. Es la versión
// "closure" de un objeto con métodos que comparten estado.

func crearContador() (incrementar func() int, decrementar func() int, valorActual func() int) {
	contador := 0

	incrementar = func() int {
		contador++
		return contador
	}
	decrementar = func() int {
		contador--
		return contador
	}
	valorActual = func() int {
		return contador
	}

	return incrementar, decrementar, valorActual
}

func main() {
	fmt.Println("=== Contador con varias operaciones ===")

	sumar, restar, ver := crearContador()

	fmt.Println("sumar()  →", sumar())  // 1
	fmt.Println("sumar()  →", sumar())  // 2
	fmt.Println("sumar()  →", sumar())  // 3
	fmt.Println("restar() →", restar()) // 2
	fmt.Println("ver()    →", ver())    // 2 (no lo modifica, solo lo lee)

	// ─────────────────────────────────────────────────────────
	// LA VARIABLE "contador" ESTÁ PROTEGIDA
	// ─────────────────────────────────────────────────────────
	// Nadie fuera de crearContador() puede tocar "contador"
	// directamente. Solo se puede modificar a través de las
	// funciones que la capturaron. Esto es encapsulamiento real,
	// sin structs ni campos privados.

	// ─────────────────────────────────────────────────────────
	// STRUCT + CLOSURE: guardar closures como campos
	// ─────────────────────────────────────────────────────────
	// También podés guardar closures dentro de un struct, mezclando
	// las dos herramientas que ya conocés.

	fmt.Println("\n=== Struct que contiene un closure ===")

	acumulador := nuevoAcumulador()
	acumulador.Sumar(100)
	acumulador.Sumar(50)
	acumulador.Sumar(-30)
	fmt.Println("Total acumulado:", acumulador.Total())

	// ─────────────────────────────────────────────────────────
	// CASO REAL: contador de visitas por producto
	// ─────────────────────────────────────────────────────────
	// En un kiosco digital, cada producto podría tener su propio
	// contador de "veces visto", sin necesitar una base de datos
	// para una métrica simple mientras el programa corre.

	fmt.Println("\n=== Caso real: contador de visitas ===")

	verMouse := crearContadorVisitas("Mouse")
	verTeclado := crearContadorVisitas("Teclado")

	verMouse()
	verMouse()
	verTeclado()
	verMouse()

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Varios closures        → pueden compartir la MISMA variable capturada")
	fmt.Println("  Esto encapsula estado  → sin exponer la variable directamente")
	fmt.Println("  Se puede combinar      → closures guardados como campos de un struct")
	fmt.Println("  Útil para              → contadores, generadores, caches simples")
}

type Acumulador struct {
	Sumar func(monto float64)
	Total func() float64
}

func nuevoAcumulador() Acumulador {
	total := 0.0
	return Acumulador{
		Sumar: func(monto float64) { total += monto },
		Total: func() float64 { return total },
	}
}

func crearContadorVisitas(producto string) func() {
	visitas := 0
	return func() {
		visitas++
		fmt.Printf("  %s visto %d vez/veces\n", producto, visitas)
	}
}
