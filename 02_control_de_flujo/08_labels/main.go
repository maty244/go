package main

import "fmt"

func main() {
	// =========================================================
	// ETIQUETAS (LABELS) EN EL CICLO FOR
	// =========================================================
	// Por defecto, break y continue solo afectan al bucle
	// más CERCANO (el que los contiene directamente).
	//
	// Con etiquetas podemos decirle a break/continue que afecte
	// a un bucle EXTERNO específico.
	//
	// Sintaxis de la etiqueta:
	//   NombreEtiqueta:
	//   for ... { }
	//
	// Uso:
	//   break NombreEtiqueta    → sale del for marcado
	//   continue NombreEtiqueta → va a la siguiente iteración del for marcado
	//
	// Las etiquetas se ponen ANTES del for (en la línea anterior).
	// Por convención se escriben en MAYÚSCULAS.
	// Se usan MUY poco, solo cuando realmente simplifican el código.

	// ─────────────────────────────────────────────────────────
	// PROBLEMA SIN ETIQUETA (break solo sale del for interno)
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Sin etiqueta: break solo sale del for interno ===")

	encontrado := false
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				encontrado = true
				break // sale del for j, pero el for i SIGUE!
			}
			fmt.Printf("  i=%d j=%d\n", i, j)
		}
		if encontrado {
			break // necesitamos este segundo break para salir del for i
		}
	}
	fmt.Println("  Terminó (necesitó dos breaks)")

	// ─────────────────────────────────────────────────────────
	// SOLUCIÓN CON ETIQUETA: break sale del for externo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Con etiqueta: break sale del for externo ===")

EXTERNO: // ← etiqueta para el for externo
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Printf("  Encontrado en i=%d j=%d → saliendo de EXTERNO\n", i, j)
				break EXTERNO // sale directamente del for marcado EXTERNO
			}
			fmt.Printf("  i=%d j=%d\n", i, j)
		}
	}
	fmt.Println("  Terminó (un solo break gracias a la etiqueta)")

	// ─────────────────────────────────────────────────────────
	// CONTINUE CON ETIQUETA
	// ─────────────────────────────────────────────────────────
	// continue NombreEtiqueta va a la SIGUIENTE iteración del for
	// marcado con esa etiqueta (saltando el resto del for interno).

	fmt.Println("\n=== continue con etiqueta ===")

OUTER:
	for i := 0; i < 4; i++ {
		fmt.Printf("Outer i=%d\n", i)
		for j := 0; j < 4; j++ {
			if j == 2 {
				fmt.Printf("  j=%d → continue OUTER (saltamos al siguiente i)\n", j)
				continue OUTER // salta al siguiente ciclo del for i
			}
			fmt.Printf("  j=%d\n", j)
		}
		fmt.Println("  (esta línea nunca se ejecuta porque continue OUTER la saltea)")
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 1: Búsqueda en matriz 2D
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: buscar en matriz 2D ===")

	matriz := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 42, 12}, // ← el 42 está aquí
		{13, 14, 15, 16},
	}

	objetivo := 42
	filaEncontrada, colEncontrada := -1, -1

BUSCAR:
	for fila, row := range matriz {
		for col, val := range row {
			if val == objetivo {
				filaEncontrada = fila
				colEncontrada = col
				break BUSCAR // encontramos, salimos de ambos fors
			}
		}
	}

	if filaEncontrada >= 0 {
		fmt.Printf("  %d encontrado en fila=%d, col=%d\n", objetivo, filaEncontrada, colEncontrada)
	} else {
		fmt.Printf("  %d no encontrado\n", objetivo)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 2: Procesar lotes con error → reintentar el lote
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: continuar al siguiente lote ===")

	lotes := [][]string{
		{"item1", "item2", "item3"},
		{"item4", "INVALID", "item6"}, // este lote tiene un item inválido
		{"item7", "item8", "item9"},
	}

LOTE:
	for i, lote := range lotes {
		fmt.Printf("Procesando lote %d:\n", i+1)
		for _, item := range lote {
			if item == "INVALID" {
				fmt.Printf("  ⚠️ Item inválido '%s', saltando al siguiente lote\n", item)
				continue LOTE // salta todo el lote y va al siguiente
			}
			fmt.Printf("  ✓ Procesado: %s\n", item)
		}
		fmt.Printf("  Lote %d completado\n", i+1)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 3: Juego de tablero (buscar posición libre)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: buscar posición libre en tablero ===")

	tablero := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", " ", "O"}, // ← hay una posición libre
	}

	fmt.Println("Tablero:")
	for _, fila := range tablero {
		fmt.Println(" ", fila)
	}

	libreI, libreJ := -1, -1

TABLERO:
	for i, fila := range tablero {
		for j, celda := range fila {
			if celda == " " {
				libreI, libreJ = i, j
				break TABLERO
			}
		}
	}

	if libreI >= 0 {
		fmt.Printf("Posición libre encontrada: [%d][%d]\n", libreI, libreJ)
	}

	// ─────────────────────────────────────────────────────────
	// ETIQUETAS EN FOR INFINITO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Etiqueta en for infinito ===")

	datos := [][]int{
		{1, 2, 3},
		{4, -99, 6}, // -99 es señal de parada
		{7, 8, 9},
	}

PROCESO:
	for _, fila := range datos {
		for _, val := range fila {
			if val == -99 {
				fmt.Println("  Señal de parada recibida (-99)")
				break PROCESO
			}
			fmt.Printf("  Procesando valor: %d\n", val)
		}
	}
	fmt.Println("  Proceso terminado")

	// ─────────────────────────────────────────────────────────
	// ALTERNATIVAS A LAS ETIQUETAS
	// ─────────────────────────────────────────────────────────
	// Las etiquetas hacen el código difícil de seguir si se abusan.
	// Alternativas más claras en muchos casos:
	//   1. Mover el bucle interno a una función con return
	//   2. Usar una variable booleana de control

	fmt.Println("\n=== Alternativa a labels: función con return ===")

	if fila, col, ok := buscarEnMatriz(matriz, 42); ok {
		fmt.Printf("  42 encontrado en [%d][%d] (con función)\n", fila, col)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Etiqueta:")
	fmt.Println("  NOMBRE:")
	fmt.Println("  for { }")
	fmt.Println()
	fmt.Println("break NOMBRE    → sale del for marcado con NOMBRE")
	fmt.Println("continue NOMBRE → va al siguiente ciclo del for marcado")
	fmt.Println()
	fmt.Println("¿Cuándo usarlas?")
	fmt.Println("  - Búsqueda en estructuras anidadas (matriz, slice de slices)")
	fmt.Println("  - Cuando necesitás salir de 2+ bucles de golpe")
	fmt.Println("  - Cuando la alternativa (bandera booleana) sería más confusa")
	fmt.Println()
	fmt.Println("Evitalas si podés: mover el bucle interno a una función")
	fmt.Println("es generalmente más limpio.")
}

// Extraer la búsqueda a una función elimina la necesidad de labels
func buscarEnMatriz(m [][]int, objetivo int) (fila, col int, encontrado bool) {
	for i, row := range m {
		for j, val := range row {
			if val == objetivo {
				return i, j, true // return sale de toda la función
			}
		}
	}
	return -1, -1, false
}
