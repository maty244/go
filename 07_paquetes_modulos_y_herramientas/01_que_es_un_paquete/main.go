package main

import "fmt"

// =========================================================
// ¿QUÉ ES UN PAQUETE?
// =========================================================
// Hasta ahora, TODOS tus archivos empiezan con "package main".
// Un paquete es simplemente la unidad en la que Go organiza el
// código: cada carpeta con archivos .go es (normalmente) UN
// paquete, y todos los archivos de esa carpeta DECLARAN
// pertenecer al mismo paquete con la línea "package NOMBRE".
//
// Hay dos tipos de paquete:
//
//   1. package main  → un PROGRAMA ejecutable. Tiene que tener
//      una función main() que es el punto de entrada. Es lo único
//      que viste hasta ahora en esta biblia.
//
//   2. package OTRONOMBRE → una LIBRERÍA. No tiene main(), no se
//      ejecuta sola: otros paquetes la IMPORTAN para usar sus
//      funciones, tipos y variables.
//
// Todo lo que usaste de la librería estándar (fmt, strings,
// errors, math...) son justamente paquetes tipo librería.

func main() {
	// ─────────────────────────────────────────────────────────
	// LO QUE YA VENÍAS HACIENDO ES USAR PAQUETES
	// ─────────────────────────────────────────────────────────
	// Cada "import" trae un paquete ajeno para usarlo.

	fmt.Println("=== fmt es un paquete de la librería estándar ===")
	fmt.Println("Esta línea usa fmt.Println, del paquete \"fmt\"")

	// ─────────────────────────────────────────────────────────
	// POR QUÉ IMPORTA ORGANIZAR EN PAQUETES
	// ─────────────────────────────────────────────────────────
	// Hasta ahora, cada tema de esta biblia es UN archivo package
	// main independiente: no se relacionan entre sí. Un proyecto
	// real no funciona así. A medida que un programa crece:
	//
	//   - Un solo archivo de 5000 líneas es imposible de mantener
	//   - Distintas partes del programa (pagos, usuarios, productos)
	//     deberían poder desarrollarse y entenderse por separado
	//   - Vas a querer REUSAR código entre distintos programas
	//
	// La solución: dividir el código en paquetes, cada uno con
	// una responsabilidad clara. En los próximos temas vas a ver
	// cómo se arma esto en la práctica: crear tus propios paquetes,
	// qué se puede usar desde afuera (visibilidad), y cómo
	// go.mod ata todo junto.

	fmt.Println("\n=== Lo que viene en esta sección ===")
	fmt.Println("  02: organizar código en carpetas/paquetes propios")
	fmt.Println("  03: visibilidad (qué se exporta y qué no)")
	fmt.Println("  04: go.mod a fondo")
	fmt.Println("  05: un mismo paquete repartido en varios archivos")
	fmt.Println("  06: ejercicio integrador con un mini proyecto real")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Paquete            → carpeta de archivos .go que comparten \"package X\"")
	fmt.Println("  package main       → programa ejecutable, necesita func main()")
	fmt.Println("  package otronombre → librería, se importa desde otro código")
	fmt.Println("  import             → trae un paquete para poder usarlo")
}
