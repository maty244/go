package main

import "fmt"

// =========================================================
// SLICES: DESCRIPTOR, ALIASING Y ARRAY SUBYACENTE
// =========================================================
// Un slice en Go NO es un array. Es una ESTRUCTURA de tres campos:
//
//   type SliceHeader struct {
//       Data unsafe.Pointer  // puntero al array subyacente
//       Len  int             // longitud actual
//       Cap  int             // capacidad total
//   }
//
// Cuando hacés: s2 := s1
// Go COPIA el header (Len, Cap, Data), pero NO el array.
// Ambos slices apuntan al MISMO array subyacente.
//
//   s1 → [Data: 0xc000, Len: 3, Cap: 5]
//                  ↓
//         [1][2][3][ ][ ]   ← array; su ubicación la decide Go
//                  ↑
//   s2 → [Data: 0xc000, Len: 3, Cap: 5]
//
// Esta es la fuente de muchos bugs sutiles en Go.

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       SLICES Y ALIASING           ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// 1. EL HEADER SE COPIA, EL ARRAY SE COMPARTE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El header se copia, el array se comparte ===")

	s1 := []int{10, 20, 30}
	s2 := s1 // copia el header (puntero, len, cap) — NO el array

	s2[0] = 999 // modifica el array compartido

	fmt.Printf("s1 = %v  (modificado, comparte array)\n", s1)
	fmt.Printf("s2 = %v\n", s2)
	fmt.Println("  → s1[0] cambió porque s1 y s2 apuntan al mismo array")

	// ─────────────────────────────────────────────────────────
	// 2. APPEND Y EL REALLOC: cuándo se separan
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Append: cuándo se separan los slices ===")

	a := make([]int, 3, 5) // len=3, cap=5 — hay espacio extra
	a[0], a[1], a[2] = 1, 2, 3

	b := a // b comparte el mismo array subyacente

	// Append cabe dentro de la cap → usa el mismo array
	// PELIGRO: b y a aún comparten, el append modifica posición 3 del array común
	b = append(b, 4)
	fmt.Printf("Después de append(b, 4) — cap original era 5:\n")
	fmt.Printf("  a = %v (len=%d, cap=%d)\n", a, len(a), cap(a))
	fmt.Printf("  b = %v (len=%d, cap=%d)\n", b, len(b), cap(b))
	fmt.Println("  → a tiene len=3 así que no 've' el elemento 4, pero comparte array")

	// Forzar realloc superando la capacidad
	c := []int{1, 2, 3} // cap=3 exacto (sin espacio)
	d := c
	d = append(d, 4) // cap superada → Go crea un nuevo array para d

	d[0] = 999 // ahora modifica el NUEVO array de d, no el de c
	fmt.Printf("\nDespués de append que fuerza realloc:\n")
	fmt.Printf("  c = %v (array original intacto)\n", c)
	fmt.Printf("  d = %v (nuevo array)\n", d)
	fmt.Println("  → Después del realloc, c y d son INDEPENDIENTES")

	// ─────────────────────────────────────────────────────────
	// 3. SUBSLICE: comparte el mismo array
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Subslice: ventana sobre el mismo array ===")

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mitad := original[3:7] // [3, 4, 5, 6] — comparte array

	fmt.Printf("original = %v\n", original)
	fmt.Printf("mitad    = %v  (original[3:7])\n", mitad)

	mitad[0] = 300 // modifica original[3]
	fmt.Printf("\nDespués de mitad[0] = 300:\n")
	fmt.Printf("original = %v  (original[3] cambió!)\n", original)
	fmt.Printf("mitad    = %v\n", mitad)

	// ─────────────────────────────────────────────────────────
	// 4. COPY: la solución para copias independientes
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== copy(): copia real del array ===")

	fuente := []int{10, 20, 30, 40, 50}
	destino := make([]int, len(fuente)) // nuevo array

	n := copy(destino, fuente) // copia los valores, no el puntero
	fmt.Printf("Copiados: %d elementos\n", n)

	destino[0] = 9999 // no afecta a fuente

	fmt.Printf("fuente  = %v (intacto)\n", fuente)
	fmt.Printf("destino = %v (copia independiente)\n", destino)
	fmt.Println("  → copy() es la forma correcta de hacer una copia real")

	// ─────────────────────────────────────────────────────────
	// 5. FUNCIONES Y SLICES: qué se puede modificar y qué no
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Funciones y slices ===")

	nums := []int{1, 2, 3, 4, 5}

	// Modificar ELEMENTOS: funciona (compartimos array)
	duplicarElementos(nums)
	fmt.Printf("Después de duplicarElementos: %v\n", nums)

	// Modificar LONGITUD (append): NO afuera si no retornamos
	agregarElemento(nums)
	fmt.Printf("Después de agregarElemento (sin retorno): %v\n", nums)
	fmt.Println("  → La longitud no cambió afuera; el header es una copia")

	// Forma correcta: retornar el nuevo slice
	nums = agregarRetornando(nums, 99)
	fmt.Printf("Después de agregarRetornando: %v\n", nums)

	// ─────────────────────────────────────────────────────────
	// 6. SLICE DE SLICES: cada fila es independiente
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slice de slices (matriz) ===")

	// Forma CORRECTA: cada fila tiene su propio array
	matriz := make([][]int, 3)
	for i := range matriz {
		matriz[i] = make([]int, 3) // array independiente por fila
		for j := range matriz[i] {
			matriz[i][j] = i*3 + j
		}
	}

	fmt.Println("Matriz:")
	for _, fila := range matriz {
		fmt.Printf("  %v\n", fila)
	}

	matriz[0][0] = 999
	fmt.Printf("Después de matriz[0][0]=999:\n")
	for _, fila := range matriz {
		fmt.Printf("  %v\n", fila)
	}
	fmt.Println("  → Solo cambió [0][0], otras filas intactas (arrays independientes)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN VISUAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen: slice internals ===")
	fmt.Println("  s2 := s1           → comparten array (copia solo el header)")
	fmt.Println("  s2[i] = x          → modifica el array compartido")
	fmt.Println("  append(s2, x)      → puede crecer o no, depende de la cap")
	fmt.Println("  copy(dst, src)     → copia real, independiente")
	fmt.Println("  s[a:b]             → subslice, mismo array subyacente")
	fmt.Println()
	fmt.Println("  Para modificar len/cap desde una función:")
	fmt.Println("    → retornar el nuevo slice (como append)")
	fmt.Println("    → o pasar *[]T (puntero al slice header)")
}

// Modificar elementos: funciona porque compartimos el array
func duplicarElementos(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// Agregar con append sin retornar: el cambio NO se ve afuera
func agregarElemento(s []int) {
	s = append(s, 100) // modifica la copia LOCAL del header
	// cuando la función termina, la copia del header se descarta
	_ = s
}

// Forma correcta: retornar el slice actualizado
func agregarRetornando(s []int, v int) []int {
	return append(s, v)
}
