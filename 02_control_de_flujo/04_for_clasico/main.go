package main

import "fmt"

func main() {
	// =========================================================
	// CICLO FOR COMPLETO (forma clásica / C-style)
	// =========================================================
	// En Go, "for" es el ÚNICO bucle disponible.
	// No existe while, do-while, foreach, loop, etc.
	// Pero "for" puede comportarse como todos ellos.
	//
	// La forma COMPLETA del for tiene tres partes:
	//   for inicialización; condición; post { }
	//
	//   inicialización → se ejecuta UNA sola vez al comienzo
	//   condición      → se evalúa ANTES de cada iteración
	//   post           → se ejecuta DESPUÉS de cada iteración
	//
	// Flujo de ejecución:
	//   1. inicialización
	//   2. evalúa condición → si false, termina
	//   3. ejecuta el cuerpo del bucle
	//   4. ejecuta post
	//   5. vuelve al paso 2

	// ─────────────────────────────────────────────────────────
	// FOR BÁSICO: contar del 0 al 4
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== For básico: 0 al 4 ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}
	// "i" no existe aquí afuera (scope del for)

	// ─────────────────────────────────────────────────────────
	// CONTAR HACIA ATRÁS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cuenta regresiva ===")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d...\n", i)
	}
	fmt.Println("¡Despegue!")

	// ─────────────────────────────────────────────────────────
	// POST PASO DE A N (incremento/decremento de más de 1)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Paso de a 2 (números pares) ===")
	for i := 0; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Println("\n=== Paso de a 3 ===")
	for i := 0; i <= 30; i += 3 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// VARIABLE DECLARADA ANTES DEL FOR
	// ─────────────────────────────────────────────────────────
	// Si necesitás acceder al contador DESPUÉS del for,
	// declaralo fuera.
	fmt.Println("\n=== Variable accesible después del for ===")

	var i int
	for i = 0; i < 5; i++ {
		// ... hacemos algo ...
	}
	fmt.Println("Valor de i al terminar el for:", i) // 5

	// ─────────────────────────────────────────────────────────
	// MÚLTIPLES VARIABLES EN UN FOR
	// ─────────────────────────────────────────────────────────
	// Go no tiene el operador coma (,) en el post de for como C,
	// pero podemos usar asignación paralela.
	fmt.Println("\n=== Múltiples variables ===")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d  j=%d\n", i, j)
	}

	// ─────────────────────────────────────────────────────────
	// CASOS DE USO REALES
	// ─────────────────────────────────────────────────────────

	// Caso 1: Tabla de multiplicar
	fmt.Println("\n=== Tabla del 7 ===")
	for n := 1; n <= 10; n++ {
		fmt.Printf("7 × %2d = %2d\n", n, 7*n)
	}

	// Caso 2: Suma acumulada (sumatoria)
	fmt.Println("\n=== Suma 1..100 (fórmula de Gauss) ===")
	suma := 0
	for n := 1; n <= 100; n++ {
		suma += n
	}
	fmt.Println("Suma de 1 a 100:", suma) // 5050

	// Caso 3: Factorial (n!)
	fmt.Println("\n=== Factorial de 10 ===")
	factorial := 1
	for k := 1; k <= 10; k++ {
		factorial *= k
	}
	fmt.Printf("10! = %d\n", factorial) // 3628800

	// Caso 4: Imprimir elementos de un slice por índice
	fmt.Println("\n=== Iterar slice por índice ===")
	productos := []string{"notebook", "mouse", "teclado", "monitor"}
	for idx := 0; idx < len(productos); idx++ {
		fmt.Printf("[%d] %s\n", idx, productos[idx])
	}

	// Caso 5: Recorrer string byte por byte (ASCII)
	fmt.Println("\n=== Recorrer string byte a byte ===")
	palabra := "GoLang"
	for pos := 0; pos < len(palabra); pos++ {
		fmt.Printf("byte[%d] = %d = '%c'\n", pos, palabra[pos], palabra[pos])
	}

	// Caso 6: Buscar un elemento en un slice
	fmt.Println("\n=== Buscar elemento en slice ===")
	numeros := []int{3, 7, 1, 9, 4, 6, 2, 8, 5}
	objetivo := 9
	encontrado := -1
	for idx := 0; idx < len(numeros); idx++ {
		if numeros[idx] == objetivo {
			encontrado = idx
			break // encontramos, salimos
		}
	}
	if encontrado >= 0 {
		fmt.Printf("Encontrado %d en índice %d\n", objetivo, encontrado)
	}

	// Caso 7: Máximo y mínimo en un slice
	fmt.Println("\n=== Máximo y mínimo ===")
	maximo, minimo := numeros[0], numeros[0]
	for idx := 1; idx < len(numeros); idx++ {
		if numeros[idx] > maximo {
			maximo = numeros[idx]
		}
		if numeros[idx] < minimo {
			minimo = numeros[idx]
		}
	}
	fmt.Printf("Slice: %v\n", numeros)
	fmt.Printf("Máximo: %d | Mínimo: %d\n", maximo, minimo)

	// ─────────────────────────────────────────────────────────
	// FOR CLÁSICO vs FOR RANGE: ¿cuándo usar cada uno?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== For clásico vs For range ===")
	fmt.Println("For CLÁSICO cuando:")
	fmt.Println("  - Necesitás control total del índice (saltar, retroceder, paso personalizado)")
	fmt.Println("  - Necesitás acceder al índice Y al siguiente elemento a la vez")
	fmt.Println("  - Iterás múltiples slices en paralelo con el mismo índice")
	fmt.Println("  - Recorrer string byte por byte")
	fmt.Println()
	fmt.Println("For RANGE cuando:")
	fmt.Println("  - Solo querés recorrer todos los elementos de corrido")
	fmt.Println("  - Iterás slices, maps, channels, strings con runes")
	fmt.Println("  - El código queda más legible (el 80% de los casos)")

	// Ejemplo donde el for clásico brilla: acceder a i y i+1
	fmt.Println("\n=== Acceder a elemento actual y siguiente ===")
	precios := []float64{10.0, 15.0, 12.0, 18.0, 9.0}
	for i := 0; i < len(precios)-1; i++ {
		diff := precios[i+1] - precios[i]
		if diff > 0 {
			fmt.Printf("precios[%d]=%.0f → precios[%d]=%.0f: subió %.0f\n", i, precios[i], i+1, precios[i+1], diff)
		} else {
			fmt.Printf("precios[%d]=%.0f → precios[%d]=%.0f: bajó %.0f\n", i, precios[i], i+1, precios[i+1], -diff)
		}
	}
}
