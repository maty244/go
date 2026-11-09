package main

import (
	"fmt"
	"slices"
)

func main() {
	// =========================================================
	// ARRAYS Y SLICES
	// =========================================================

	// ─────────────────────────────────────────────────────────
	// ARRAYS: tamaño FIJO, parte del tipo
	// ─────────────────────────────────────────────────────────
	// Un array tiene un tamaño fijo que no puede cambiar.
	// El tamaño es PARTE del tipo: [3]int y [5]int son tipos distintos.
	// Se usan cuando el tamaño es conocido y fijo (ej: coordenadas 3D).

	var arr [5]int // zero values: [0 0 0 0 0]
	arr[0] = 10
	arr[2] = 30
	arr[4] = 50

	fmt.Println("=== Arrays ===")
	fmt.Println("Array:", arr)
	fmt.Println("Longitud:", len(arr))

	// Declaración con valores (array literal)
	dias := [7]string{"Dom", "Lun", "Mar", "Mié", "Jue", "Vie", "Sáb"}
	fmt.Println("Días:", dias)

	// [...]int: Go deduce el tamaño automáticamente
	primos := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primos:", primos, "| len:", len(primos))

	// Los arrays se COPIAN cuando se asignan (valor, no referencia)
	original := [3]int{1, 2, 3}
	copia := original // copia completa
	copia[0] = 99
	fmt.Printf("\nOriginal: %v | Copia: %v (independientes)\n", original, copia)

	// Los arrays SÍ son comparables con == (si sus elementos son comparables)
	fmt.Println("¿[3]int{1,2,3} == original?:", [3]int{1, 2, 3} == original) // true

	// ─────────────────────────────────────────────────────────
	// SLICES: tamaño DINÁMICO, la herramienta principal de Go
	// ─────────────────────────────────────────────────────────
	// Un slice es una VISTA (descriptor) sobre un array subyacente.
	// Tiene tres componentes: puntero, longitud (len), capacidad (cap).
	// Es mutable y puede crecer con append().
	// En la práctica, casi siempre usamos slices, no arrays.
	//
	// ⚠️ A diferencia de los arrays, los slices NO se pueden comparar con ==
	// (salvo contra nil: s == nil). Para comparar dos slices usá slices.Equal(a, b).

	fmt.Println("\n=== Slices ===")

	// Declaración de slice vacío (nil)
	var s1 []int
	fmt.Printf("slice nil: %v | len=%d | cap=%d | es nil: %v\n",
		s1, len(s1), cap(s1), s1 == nil)

	// Slice literal
	notas := []int{9, 7, 8, 6, 10, 5}
	fmt.Printf("notas: %v | len=%d | cap=%d\n", notas, len(notas), cap(notas))
	fmt.Println("¿notas iguales a copia?:", slices.Equal(notas, []int{9, 7, 8, 6, 10, 5}))

	// Acceso por índice
	fmt.Println("Primera nota:", notas[0])
	fmt.Println("Última nota:", notas[len(notas)-1])

	// ─────────────────────────────────────────────────────────
	// MAKE: crear slice con longitud y capacidad específicas
	// ─────────────────────────────────────────────────────────
	// make([]tipo, longitud, capacidad)
	// capacidad es opcional, por defecto igual a longitud

	s2 := make([]int, 5)     // len=5, cap=5, todos en 0
	s3 := make([]int, 3, 10) // len=3, cap=10 (espacio pre-reservado)
	fmt.Printf("\nmake([]int,5): %v | len=%d | cap=%d\n", s2, len(s2), cap(s2))
	fmt.Printf("make([]int,3,10): %v | len=%d | cap=%d\n", s3, len(s3), cap(s3))

	// ─────────────────────────────────────────────────────────
	// APPEND: agregar elementos
	// ─────────────────────────────────────────────────────────
	// append retorna un NUEVO slice (puede o no ser el mismo array subyacente).
	// SIEMPRE asignamos el resultado: s = append(s, valor)

	var colores []string
	colores = append(colores, "rojo")
	colores = append(colores, "verde", "azul") // múltiples a la vez
	fmt.Println("\n=== append ===")
	fmt.Println("colores:", colores)

	// Append de un slice en otro con ...
	extras := []string{"amarillo", "naranja"}
	colores = append(colores, extras...)
	fmt.Println("colores con extras:", colores)

	// ─────────────────────────────────────────────────────────
	// SLICING: obtener sub-slices
	// ─────────────────────────────────────────────────────────
	// s[inicio:fin] → elementos desde inicio hasta fin-1
	// s[inicio:] → desde inicio hasta el final
	// s[:fin] → desde el inicio hasta fin-1
	// s[:] → todos los elementos

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("\n=== Slicing ===")
	fmt.Println("nums:", nums)
	fmt.Println("nums[2:5]:", nums[2:5]) // [2 3 4]
	fmt.Println("nums[:3]:", nums[:3])   // [0 1 2]
	fmt.Println("nums[7:]:", nums[7:])   // [7 8 9]
	fmt.Println("nums[:]:", nums[:])     // todos

	// ─────────────────────────────────────────────────────────
	// CUIDADO: Slices comparten memoria con el array subyacente
	// ─────────────────────────────────────────────────────────
	base := []int{1, 2, 3, 4, 5}
	vista := base[1:4] // comparte memoria con base
	vista[0] = 99      // modifica base también!
	fmt.Println("\n=== Slice comparte memoria ===")
	fmt.Println("base:", base)   // [1 99 3 4 5] ← modificado!
	fmt.Println("vista:", vista) // [99 3 4]

	// Para una copia independiente, usar copy() o slices.Clone()
	independiente := make([]int, len(base[1:4]))
	copy(independiente, base[1:4])
	independiente[0] = 777
	fmt.Println("independiente:", independiente) // [777 3 4]
	fmt.Println("base sin cambios:", base)       // no modificado

	// ─────────────────────────────────────────────────────────
	// CONVERSIÓN ENTRE ARRAY Y SLICE
	// ─────────────────────────────────────────────────────────
	// 1. Array → Slice: haciendo slicing con arr[:] (comparten memoria).
	// 2. Slice → Array (Go 1.20+): [N]T(slice) COPIA los primeros N elementos
	//    a un array nuevo (ojo: hace panic si len(slice) < N).
	// 3. Slice → Puntero a Array (Go 1.17+): (*[N]T)(slice) apunta a los
	//    primeros N elementos del array subyacente SIN copiar.
	fmt.Println("\n=== Conversión Array ↔ Slice ===")
	arrFijo := [4]int{10, 20, 30, 40}
	sliceDesdeArr := arrFijo[:] // Array → Slice (vista compartida)
	fmt.Println("Array a Slice (arrFijo[:]):", sliceDesdeArr)

	arrCopiado := [3]int(sliceDesdeArr) // Slice → Array (copia independiente)
	arrCopiado[0] = 999
	fmt.Printf("Slice a Array ([3]int(s)): %v | arrFijo sigue intacto: %v\n", arrCopiado, arrFijo)

	ptrArr := (*[2]int)(sliceDesdeArr) // Slice → *[2]int (comparte memoria)
	ptrArr[0] = 111
	fmt.Printf("Puntero (*[2]int)(s): %v | arrFijo modificado: %v\n", *ptrArr, arrFijo)

	// ─────────────────────────────────────────────────────────
	// COPY: copiar slices
	// ─────────────────────────────────────────────────────────
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3) // solo 3 elementos
	n := copy(dst, src)   // copia min(len(dst), len(src)) elementos
	fmt.Printf("\ncopy: copiados %d elementos, dst = %v\n", n, dst)

	// ─────────────────────────────────────────────────────────
	// ELIMINAR ELEMENTOS DE UN SLICE
	// ─────────────────────────────────────────────────────────
	// Se puede hacer manualmente con append o con slices.Delete (Go 1.21+).
	frutas := []string{"manzana", "banana", "naranja", "uva", "mango"}
	fmt.Println("\n=== Eliminar elemento ===")
	fmt.Println("Original:", frutas)

	i := 2                                       // índice de "naranja"
	frutas = append(frutas[:i], frutas[i+1:]...) // forma clásica con append
	fmt.Println("Sin 'naranja' (append):", frutas)

	frutas = slices.Delete(frutas, 1, 2) // forma moderna: elimina rango [1, 2) ("banana")
	fmt.Println("Sin 'banana' (slices.Delete):", frutas)

	// ─────────────────────────────────────────────────────────
	// ORDENAR Y BUSCAR EN SLICES (paquete moderno "slices")
	// ─────────────────────────────────────────────────────────
	// En Go moderno (1.21+) usamos el paquete "slices" para ordenar,
	// invertir y buscar en cualquier slice de tipos ordenables.
	// (En código viejo vas a ver sort.Ints, sort.Strings, sort.SearchInts).
	numeros := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Println("\n=== Ordenamiento (paquete slices) ===")
	fmt.Println("Sin ordenar:", numeros)
	slices.Sort(numeros)
	fmt.Println("Ordenado:", numeros)

	palabras := []string{"banana", "apple", "cherry", "date"}
	slices.Sort(palabras)
	fmt.Println("Strings ordenados:", palabras)

	// Orden descendente (invertir el slice ya ordenado)
	slices.Reverse(numeros)
	fmt.Println("Descendente (slices.Reverse):", numeros)

	// ─────────────────────────────────────────────────────────
	// BÚSQUEDA EN SLICE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Búsqueda ===")
	fmt.Println("¿Contiene 5? (slices.Contains):", slices.Contains(numeros, 5))

	// Búsqueda binaria (requisito: el slice DEBE estar ordenado en forma ascendente)
	slices.Sort(numeros)
	idx, encontrado := slices.BinarySearch(numeros, 5)
	fmt.Printf("BinarySearch de 5 en %v → índice: %d (encontrado: %v)\n", numeros, idx, encontrado)

	// ─────────────────────────────────────────────────────────
	// SLICES 2D (matriz)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slice 2D (matriz) ===")
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, fila := range matriz {
		for j, val := range fila {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}

	// ─────────────────────────────────────────────────────────
	// PATRONES COMUNES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrones comunes ===")

	// Stack (pila): push con append, pop cortando
	var stack []int
	stack = append(stack, 1, 2, 3) // push
	fmt.Println("Stack:", stack)
	top := stack[len(stack)-1]   // peek
	stack = stack[:len(stack)-1] // pop
	fmt.Printf("pop: top=%d, stack=%v\n", top, stack)

	// Queue (cola): enqueue con append, dequeue cortando
	var queue []string
	queue = append(queue, "primero", "segundo", "tercero")
	frente := queue[0]
	queue = queue[1:]
	fmt.Printf("dequeue: frente='%s', queue=%v\n", frente, queue)
}
