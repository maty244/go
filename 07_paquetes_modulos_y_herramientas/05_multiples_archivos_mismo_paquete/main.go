package main

import "fmt"

// =========================================================
// UN MISMO PAQUETE REPARTIDO EN VARIOS ARCHIVOS
// =========================================================
// En los temas 02 y 03 separamos código en paquetes DISTINTOS
// (carpetas distintas). Pero también podés dividir UN MISMO
// paquete en varios archivos DENTRO de la misma carpeta.
//
// Mirá esta carpeta: tiene DOS archivos, main.go y ayudantes.go,
// y AMBOS dicen "package main" arriba. Para Go, son EL MISMO
// paquete: todo lo declarado en ayudantes.go está disponible acá
// en main.go directamente, SIN IMPORT y sin prefijo.
//
// Esto es distinto de los temas 02/03: ahí "precios" y "catalogo"
// eran paquetes DISTINTOS (carpetas distintas), y había que
// importarlos y usar el prefijo (precios.AplicarDescuento).
// Acá, al ser EL MISMO paquete, no hace falta nada de eso.

func main() {
	fmt.Println("=== Un paquete, varios archivos ===")

	// formatearPrecio y validarStock están definidas en
	// ayudantes.go, pero las usamos directo, sin prefijo ni import,
	// porque comparten el mismo "package main".
	fmt.Println(formatearPrecio(15499.9))

	if err := validarStock(3, 10); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stock válido para la compra")
	}

	if err := validarStock(15, 10); err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────────────────
	// ¿CUÁNDO CONVIENE DIVIDIR EN ARCHIVOS (sin crear un paquete
	// nuevo) VS CREAR UN PAQUETE NUEVO?
	// ─────────────────────────────────────────────────────────
	// - Varios archivos, MISMO paquete: cuando el código está
	//   MUY relacionado y un solo archivo quedaría gigante e
	//   incómodo de navegar, pero conceptualmente es "una sola
	//   cosa" (por ejemplo: separar tipos, validaciones y lógica
	//   principal de una misma feature).
	//
	// - Paquete nuevo (carpeta nueva): cuando el código tiene una
	//   RESPONSABILIDAD DISTINTA y podría reusarse desde otro
	//   programa por separado (como "precios" en el tema 02).

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Varios archivos, mismo \"package X\" → mismo paquete")
	fmt.Println("  Todo lo declarado en un archivo      → visible en los demás")
	fmt.Println("  del mismo paquete, SIN import ni prefijo")
	fmt.Println("  Usalo para                            → dividir un paquete grande")
	fmt.Println("  Convention típica                     → main.go + archivos por tema")
}
