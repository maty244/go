package main

import "fmt"

// =========================================================
// EN GO TODO SE PASA POR VALOR
// =========================================================
// Este es uno de los conceptos más importantes en Go.
// Entender la diferencia evita bugs muy difíciles de rastrear.
//
// Go copia cada argumento dentro del parámetro de la función. La diferencia
// está en QUÉ contiene el valor copiado:
//
//   - int, bool, arrays y structs contienen sus datos directamente;
//   - un puntero contiene una dirección: la copia apunta al mismo destino;
//   - un slice contiene un descriptor que puede compartir un array;
//   - un map contiene un descriptor que comparte una tabla interna.
//
// Dos valores pueden observar los mismos datos sin que exista una modalidad
// separada de "paso por referencia".

// ─────────────────────────────────────────────────────────
// EJEMPLO 1: int — paso por VALOR
// ─────────────────────────────────────────────────────────
func doblarPorValor(n int) {
	n = n * 2 // modifica la COPIA local, no el original
	fmt.Println("  dentro de doblarPorValor, n =", n)
}

// Para modificar el int, necesitamos pasar un puntero
func doblarConPuntero(n *int) {
	*n = *n * 2 // * desreferencia: accede al valor apuntado
	fmt.Println("  dentro de doblarConPuntero, *n =", *n)
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 2: struct — paso por VALOR vs PUNTERO
// ─────────────────────────────────────────────────────────
type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

// Por valor: recibe copia, el original no cambia
func aplicarIVAPorValor(p Producto) Producto {
	p.Precio *= 1.21 // modifica la copia
	return p         // retorna la copia modificada
}

// Por puntero: modifica el original directamente
func aplicarIVAPorPuntero(p *Producto) {
	p.Precio *= 1.21 // modifica el original
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 3: slice — descriptor copiado, array compartido
// ─────────────────────────────────────────────────────────
// Los slices internamente son {puntero al array, len, cap}.
// Al pasarlos a una función, el puntero al array se copia,
// pero APUNTA al mismo array → modificar elementos SÍ afecta al original.
// PERO: append() puede crear un nuevo array → el caller no ve el cambio.

func duplicarElementos(s []int) {
	for i := range s {
		s[i] *= 2 // modifica los elementos → SÍ afecta al original
	}
}

func agregarElemento(s []int, n int) []int {
	// append puede crear un nuevo array
	// por eso hay que RETORNAR el slice
	return append(s, n)
}

// ─────────────────────────────────────────────────────────
// EJEMPLO 4: map — descriptor copiado, tabla compartida
// ─────────────────────────────────────────────────────────
// Copiar un map no clona sus entradas. Ambos descriptores observan la misma
// tabla, salvo que una variable se reasigne a un map diferente.
func agregarPrecio(precios map[string]float64, producto string, precio float64) {
	precios[producto] = precio // modifica el map original directamente
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║     TODO SE PASA POR VALOR     ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// INT: valor vs puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== int: valor vs puntero ===")

	precio := 100

	fmt.Println("Antes:", precio)
	doblarPorValor(precio)
	fmt.Println("Después de doblarPorValor:", precio, "← sin cambio!")

	doblarConPuntero(&precio) // & = "dame la dirección de precio"
	fmt.Println("Después de doblarConPuntero:", precio, "← modificado!")

	// ─────────────────────────────────────────────────────────
	// STRUCT: valor vs puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== struct: valor vs puntero ===")

	notebook := Producto{"Notebook", 1000.00, 5}
	mouse := Producto{"Mouse", 20.00, 30}

	fmt.Printf("notebook antes: $%.2f\n", notebook.Precio)

	// Por valor: recibimos la copia modificada
	notebookConIVA := aplicarIVAPorValor(notebook)
	fmt.Printf("notebook ORIGINAL sin cambios: $%.2f\n", notebook.Precio)
	fmt.Printf("notebookConIVA (nuevo): $%.2f\n", notebookConIVA.Precio)

	// Por puntero: el original se modifica
	fmt.Printf("\nmouse antes: $%.2f\n", mouse.Precio)
	aplicarIVAPorPuntero(&mouse)
	fmt.Printf("mouse ORIGINAL modificado: $%.2f\n", mouse.Precio)

	// ─────────────────────────────────────────────────────────
	// SLICE: descriptor copiado, array compartido
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== slice: descriptor por valor, array compartido ===")

	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Antes:", numeros)

	duplicarElementos(numeros) // modifica los elementos
	fmt.Println("Después de duplicarElementos:", numeros, "← SÍ modificado!")

	// Pero append no se propaga sin retorno
	fmt.Printf("\nLen antes: %d\n", len(numeros))
	agregarElemento(numeros, 999) // el append no se ve afuera
	fmt.Printf("Len después (sin retorno): %d ← NO cambió!\n", len(numeros))

	numeros = agregarElemento(numeros, 999) // así sí funciona
	fmt.Printf("Len después (con retorno): %d ← SÍ cambió!\n", len(numeros))

	// ─────────────────────────────────────────────────────────
	// MAP: descriptor copiado, tabla compartida
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== map: descriptor por valor, tabla compartida ===")

	precios := map[string]float64{
		"Notebook": 1000.00,
	}
	fmt.Println("Antes:", precios)

	agregarPrecio(precios, "Mouse", 20.00) // modifica el map original
	agregarPrecio(precios, "Teclado", 60.00)

	fmt.Println("Después:", precios, "← SÍ modificado!")

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR CADA UNO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Guía de uso ===")
	fmt.Println("Por VALOR (copia):")
	fmt.Println("  - Tipos pequeños: int, float, bool")
	fmt.Println("  - Struct pequeño que no necesitás modificar")
	fmt.Println("  - Cuando querés garantizar inmutabilidad del argumento")
	fmt.Println()
	fmt.Println("Puntero pasado POR VALOR (*T):")
	fmt.Println("  - Struct medido como costoso de copiar")
	fmt.Println("  - Cuando la función NECESITA modificar el valor")
	fmt.Println("  - Cuando querés señalar que algo puede ser nil")
	fmt.Println()
	fmt.Println("Valores con datos internos compartidos:")
	fmt.Println("  - slices y maps se copian, pero pueden compartir almacenamiento")
	fmt.Println("  - Para modificar ELEMENTOS de un slice: función sin *")
	fmt.Println("  - Para agregar/quitar con append: necesitás retornar el slice")

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla resumen ===")
	fmt.Println("Tipo       | Se pasa     | Modif. elementos | Modif. tamaño")
	fmt.Println("-----------|-------------|------------------|---------------")
	fmt.Println("int/float  | por valor   | necesita *       | N/A")
	fmt.Println("string     | por valor   | necesita *       | N/A (inmutable)")
	fmt.Println("struct     | por valor   | necesita *       | necesita *")
	fmt.Println("array      | por valor   | necesita *       | tamaño fijo")
	fmt.Println("slice      | descriptor  | directo ✓        | retornar slice")
	fmt.Println("map        | descriptor  | directo ✓        | directo ✓")
}
