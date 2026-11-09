package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// =========================================================
	// MAPS (Diccionarios / Hash tables)
	// =========================================================
	// Un map es una colección de pares CLAVE → VALOR.
	// Las claves son únicas. Los valores pueden repetirse.
	// Acceso promedio O(1) (muy rápido).
	// El zero value de un map es nil.
	// La clave K debe ser comparable: bool, números, strings, channels,
	// punteros, interfaces y arrays/structs cuyos componentes sean comparables.
	// Slices, maps y funciones no pueden ser claves.
	//
	// ⚠️ Los maps NO se pueden comparar con == (salvo contra nil: m == nil).
	// Para comparar si dos maps tienen las mismas claves y valores, usá maps.Equal(m1, m2).
	//
	// Sintaxis: map[tipoClave]tipoValor

	// ─────────────────────────────────────────────────────────
	// CREACIÓN DE MAPS
	// ─────────────────────────────────────────────────────────

	// Forma 1: make (mapa vacío listo para usar)
	edades := make(map[string]int)

	// Forma 2: map literal
	capitales := map[string]string{
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
		"Uruguay":   "Montevideo",
	}

	// Forma 3: map vacío con literal
	precios := map[string]float64{}

	fmt.Println("=== Maps básicos ===")
	fmt.Println("edades (vacío):", edades)
	fmt.Println("capitales:", capitales)
	fmt.Println("precios (vacío):", precios)

	// ─────────────────────────────────────────────────────────
	// AGREGAR Y MODIFICAR
	// ─────────────────────────────────────────────────────────
	// La misma sintaxis sirve para crear y actualizar.
	edades["Ana"] = 28
	edades["Carlos"] = 35
	edades["Mia"] = 22

	fmt.Println("\n=== Agregar y modificar ===")
	fmt.Println("edades:", edades)

	edades["Ana"] = 29 // actualizar
	fmt.Println("Ana actualizada:", edades["Ana"])

	// ─────────────────────────────────────────────────────────
	// LEER VALORES
	// ─────────────────────────────────────────────────────────
	// Acceso simple: retorna el zero value si la clave no existe
	// Acceso con verificación: retorna valor + bool (existe?)

	fmt.Println("\n=== Leer valores ===")

	// Sin verificación (puede retornar 0 si la clave no existe)
	edadAna := edades["Ana"]
	edadDesconocido := edades["Pedro"] // Pedro no existe → retorna 0
	fmt.Printf("Ana: %d | Pedro (no existe): %d\n", edadAna, edadDesconocido)

	// CON verificación (el patrón correcto: "comma-ok idiom")
	if edad, existe := edades["Carlos"]; existe {
		fmt.Printf("Carlos tiene %d años\n", edad)
	} else {
		fmt.Println("Carlos no encontrado")
	}

	if _, existe := edades["Pedro"]; !existe {
		fmt.Println("Pedro no está en el mapa")
	}

	// ─────────────────────────────────────────────────────────
	// ELIMINAR Y COMPARAR ELEMENTOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Eliminar y comparar ===")
	fmt.Println("Antes:", edades)
	delete(edades, "Mia")
	fmt.Println("Después de delete('Mia'):", edades)

	// delete en una clave que no existe no da error (silencioso)
	delete(edades, "Fantasma") // no hace nada

	// Comparar dos maps con maps.Equal (Go 1.21+)
	copiaEdades := map[string]int{"Ana": 29, "Carlos": 35}
	fmt.Println("¿edades y copiaEdades son iguales? (maps.Equal):", maps.Equal(edades, copiaEdades))

	// clear (Go 1.21+) elimina todas las entradas de una vez.
	clear(copiaEdades)
	fmt.Println("Después de clear:", copiaEdades, "len:", len(copiaEdades))

	// ─────────────────────────────────────────────────────────
	// ITERAR SOBRE UN MAP (con for range)
	// ─────────────────────────────────────────────────────────
	// IMPORTANTE: el orden de iteración NO ESTÁ ESPECIFICADO.
	// No dependas de él; tampoco asumas que cambia en cada recorrido.

	fmt.Println("\n=== Iterar map ===")
	inventario := map[string]int{
		"manzanas": 50,
		"bananas":  30,
		"naranjas": 75,
		"uvas":     20,
	}

	fmt.Println("Inventario (orden no especificado):")
	for producto, cantidad := range inventario {
		fmt.Printf("  %s: %d unidades\n", producto, cantidad)
	}

	// Para orden predecible en Go moderno (1.23+): slices.Sorted(maps.Keys(m))
	fmt.Println("\nInventario (orden alfabético con slices.Sorted + maps.Keys):")
	for _, k := range slices.Sorted(maps.Keys(inventario)) {
		fmt.Printf("  %s: %d unidades\n", k, inventario[k])
	}

	// ─────────────────────────────────────────────────────────
	// MAPS COMO CONTADORES (patrón frecuencia)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Map como contador ===")

	texto := "go es genial go me gusta go es rápido"
	palabras := splitPalabras(texto)
	frecuencia := make(map[string]int)

	for _, palabra := range palabras {
		frecuencia[palabra]++ // si la clave no existe, empieza en 0 (zero value)
	}

	fmt.Println("Frecuencia de palabras:")
	for _, palabra := range slices.Sorted(maps.Keys(frecuencia)) {
		fmt.Printf("  '%s': %d veces\n", palabra, frecuencia[palabra])
	}

	// ─────────────────────────────────────────────────────────
	// MAPS ANIDADOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Maps anidados ===")

	// Map de maps: útil para datos jerárquicos
	alumnos := map[string]map[string]int{
		"Ana":    {"matemáticas": 9, "física": 8, "química": 7},
		"Carlos": {"matemáticas": 6, "física": 9, "química": 8},
		"Mia":    {"matemáticas": 10, "física": 7, "química": 9},
	}

	for alumno, notas := range alumnos {
		promedio := 0
		for _, nota := range notas {
			promedio += nota
		}
		fmt.Printf("  %s: promedio %.1f\n", alumno, float64(promedio)/float64(len(notas)))
	}

	// Acceder a un mapa anidado
	if notasAna, ok := alumnos["Ana"]; ok {
		fmt.Printf("  Ana en matemáticas: %d\n", notasAna["matemáticas"])
	}

	// ─────────────────────────────────────────────────────────
	// MAP CON STRUCTS COMO VALORES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Map con structs ===")

	type Producto struct {
		Nombre string
		Precio float64
		Stock  int
	}

	catalogo := map[string]Producto{
		"P001": {Nombre: "Notebook", Precio: 1500.00, Stock: 5},
		"P002": {Nombre: "Mouse", Precio: 25.99, Stock: 42},
		"P003": {Nombre: "Teclado", Precio: 75.50, Stock: 18},
	}

	// ⚠️ DETALLE CLAVE DE GO: los valores dentro de un map NO son direccionables
	// (porque el map puede crecer y mover sus buckets en memoria).
	// Por eso NO compila hacer directamente:
	//   catalogo["P001"].Stock = 10  // ❌ cannot assign to struct field in map
	//
	// Para modificar un campo de un struct en un map: leer → modificar → reasignar:
	prod := catalogo["P001"]
	prod.Stock = 10
	catalogo["P001"] = prod

	for _, codigo := range slices.Sorted(maps.Keys(catalogo)) {
		p := catalogo[codigo]
		fmt.Printf("  [%s] %s - $%.2f (stock: %d)\n",
			codigo, p.Nombre, p.Precio, p.Stock)
	}

	// ─────────────────────────────────────────────────────────
	// MAP COMO SET (conjunto sin duplicados)
	// ─────────────────────────────────────────────────────────
	// Go no tiene tipo Set nativo. Se simula con map[tipo]bool
	// o map[tipo]struct{} (más eficiente, struct{} ocupa 0 bytes)

	fmt.Println("\n=== Map como Set ===")
	numeros := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	set := make(map[int]struct{})

	for _, n := range numeros {
		set[n] = struct{}{} // struct{}{} es la instancia vacía de struct{}
	}

	fmt.Printf("Números con duplicados: %v\n", numeros)
	fmt.Println("Sin duplicados (set):", slices.Sorted(maps.Keys(set)))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Crear:         make(map[K]V) o map[K]V{...}")
	fmt.Println("Agregar:       m[clave] = valor")
	fmt.Println("Leer seguro:   val, ok := m[clave]")
	fmt.Println("Eliminar:      delete(m, clave)")
	fmt.Println("Vaciar:        clear(m)")
	fmt.Println("Comparar:      maps.Equal(m1, m2)")
	fmt.Println("Iterar:        for k, v := range m { ... }")
	fmt.Println("Ordenar keys:  slices.Sorted(maps.Keys(m))")
	fmt.Println("Nil map:       no se puede escribir en un map nil (panic)")
}

func splitPalabras(s string) []string {
	var palabras []string
	palabra := ""
	for _, r := range s {
		if r == ' ' {
			if palabra != "" {
				palabras = append(palabras, palabra)
				palabra = ""
			}
		} else {
			palabra += string(r)
		}
	}
	if palabra != "" {
		palabras = append(palabras, palabra)
	}
	return palabras
}
