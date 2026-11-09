package main

import (
	"fmt"

	"gocito/07_paquetes_modulos_y_herramientas/02_organizar_en_carpetas/precios"
)

// =========================================================
// ORGANIZAR CÓDIGO EN CARPETAS (PAQUETES PROPIOS)
// =========================================================
// Este es el primer ejemplo de esta biblia con MÁS DE UN ARCHIVO.
// Mirá la carpeta de este tema:
//
//   02_organizar_en_carpetas/
//   ├── main.go              ← package main (este archivo)
//   └── precios/
//       └── precios.go       ← package precios
//
// "precios" es un paquete propio, una carpeta con su propio
// nombre de paquete. Para usarlo desde main.go, hay que IMPORTARLO
// con su ruta completa a partir del nombre del módulo.
//
// El módulo de este repo se llama "gocito" (mirá el archivo
// go.mod en la raíz). Por eso la ruta de import es:
//
//   "gocito/07_paquetes_modulos_y_herramientas/02_organizar_en_carpetas/precios"
//
// Es literalmente: nombre del módulo + ruta de carpetas hasta
// llegar al paquete que querés usar.

func main() {
	fmt.Println("=== Usando un paquete propio ===")

	precioBase := 1000.0

	// precios.AplicarDescuento y precios.AplicarIVA viven en OTRO
	// archivo, en OTRO paquete, pero los usamos anteponiendo el
	// nombre del paquete, exactamente como con fmt.Println.
	conDescuento := precios.AplicarDescuento(precioBase, 0.10)
	conIVA := precios.AplicarIVA(conDescuento)

	fmt.Printf("Precio base:        $%.2f\n", precioBase)
	fmt.Printf("Con 10%% descuento:  $%.2f\n", conDescuento)
	fmt.Printf("Con IVA incluido:   $%.2f\n", conIVA)

	// ─────────────────────────────────────────────────────────
	// POR QUÉ SEPARAR ESTO EN OTRO PAQUETE
	// ─────────────────────────────────────────────────────────
	// La lógica de precios (descuentos, IVA) es algo que:
	//   - Se puede testear por separado (10_testing/, más adelante)
	//   - Se puede reusar desde OTRO programa que también necesite
	//     calcular precios, sin copiar y pegar código
	//   - Tiene una responsabilidad clara y un nombre que la describe

	// ─────────────────────────────────────────────────────────
	// DOS REGLAS DE ORO DE LOS PAQUETES EN GO
	// ─────────────────────────────────────────────────────────
	// 1. Prohibidos los imports circulares: si el paquete A importa al
	//    paquete B, entonces B NO puede importar a A (el compilador falla
	//    con "import cycle not allowed"). Esto fuerza un diseño limpio.
	// 2. Carpeta especial "internal/": cualquier paquete dentro de una
	//    carpeta llamada "internal" solo puede ser importado por código
	//    que viva dentro del árbol padre de ese "internal".

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Una carpeta = un paquete (con su propio \"package nombre\")")
	fmt.Println("  Se importa con: \"nombreDelModulo/ruta/de/carpetas\"")
	fmt.Println("  Se usa como:    nombrePaquete.Funcion(...)")
	fmt.Println("  Sin ciclos:     Go prohíbe imports circulares (A ⇄ B)")
	fmt.Println("  Ventaja:        separar responsabilidades y poder reusar código")
}
