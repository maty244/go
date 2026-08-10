package main

import "fmt"

// =========================================================
// ¿POR QUÉ GENERICS?
// =========================================================
// Ya viste dos formas de escribir código que funciona con
// "cualquier tipo":
//
//   1. Interfaces (06_interfaces_y_errores/): "cualquier tipo que sepa hacer X"
//   2. any (06_interfaces_y_errores/07_interfaz_vacia_any): "literalmente cualquier tipo"
//
// Generics es una TERCERA herramienta, para un problema distinto:
// funciones/tipos que hacen EXACTAMENTE LO MISMO sin importar el
// tipo concreto, pero SIN perder la seguridad de tipos de `any`.
//
// El problema que resuelven: mirá estas dos funciones.

func maximoInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximoFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Son IDÉNTICAS salvo el tipo. Sin generics, tenés dos opciones
// (ambas malas):
//   A) copiar y pegar una función por cada tipo (lo de arriba)
//   B) usar "any" y perder toda garantía de tipos:

func maximoAny(a, b any) any {
	// Esto NO COMPILA: any no tiene el operador >
	// return a > b // ERROR
	// Habría que hacer type assertion, manejar cada tipo posible...
	// muchísimo código extra para algo que debería ser trivial.
	return nil
}

func main() {
	fmt.Println("=== El problema sin generics ===")
	fmt.Println("maximoInt(3, 7)     =", maximoInt(3, 7))
	fmt.Println("maximoFloat(3.5, 2) =", maximoFloat(3.5, 2))
	fmt.Println("(hicimos falta DOS funciones para la MISMA lógica)")

	// ─────────────────────────────────────────────────────────
	// LA SOLUCIÓN: una función GENÉRICA
	// ─────────────────────────────────────────────────────────
	// En el próximo tema vas a ver la sintaxis exacta, pero
	// adelantamos el resultado: con generics, se escribe UNA sola
	// función que funciona para int, float64, string, y cualquier
	// tipo "comparable con >", manteniendo el chequeo de tipos del
	// compilador (a diferencia de "any").

	fmt.Println("\n=== Adelanto: con generics, UNA función alcanza ===")
	fmt.Println("Maximo(3, 7)       =", Maximo(3, 7))
	fmt.Println("Maximo(3.5, 2.1)   =", Maximo(3.5, 2.1))
	fmt.Println("Maximo(\"pera\", \"manzana\") =", Maximo("pera", "manzana"))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Sin generics    → copiar función por tipo, O usar any (sin seguridad)")
	fmt.Println("  Con generics    → una función, funciona para varios tipos,")
	fmt.Println("                     CON chequeo de tipos en compilación")
	fmt.Println("  Dónde usarlos   → estructuras de datos y algoritmos genéricos")
	fmt.Println("                     (listas, comparaciones, filtros, no lógica de negocio)")
}

// Maximo es la versión genérica: la vas a entender línea por
// línea en el próximo tema. Por ahora, notá que es UNA función
// para los tres casos de arriba.
func Maximo[T int | float64 | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}
