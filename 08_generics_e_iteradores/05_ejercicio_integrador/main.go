package main

import (
	"cmp"
	"fmt"
)

// =========================================================
// EJERCICIO INTEGRADOR: Filtrar, Mapear y Reducir GENÉRICOS
// =========================================================
// El trío más famoso de la programación funcional, escrito UNA
// sola vez con generics, y reusado sobre el catálogo del kiosco
// (productos, precios, nombres) sin escribir una versión por tipo.
//
// Fijate que esto es la versión genérica del "filtrar" que
// escribiste en 05_funciones_y_closures/07_funciones_anonimas a mano, específico para []int.

// ─────────────────────────────────────────────────────────
// Filtrar: se queda con los elementos que cumplen una condición
// ─────────────────────────────────────────────────────────
func Filtrar[T any](lista []T, cumple func(T) bool) []T {
	var resultado []T
	for _, v := range lista {
		if cumple(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

// ─────────────────────────────────────────────────────────
// Mapear: transforma cada elemento en OTRO tipo (T → R)
// ─────────────────────────────────────────────────────────
// Usa DOS parámetros de tipo: T (entrada) y R (salida). No tienen
// por qué ser el mismo tipo.
func Mapear[T, R any](lista []T, transformar func(T) R) []R {
	resultado := make([]R, len(lista))
	for i, v := range lista {
		resultado[i] = transformar(v)
	}
	return resultado
}

// ─────────────────────────────────────────────────────────
// Reducir: combina todos los elementos en UN solo valor
// ─────────────────────────────────────────────────────────
func Reducir[T, R any](lista []T, inicial R, combinar func(R, T) R) R {
	acumulado := inicial
	for _, v := range lista {
		acumulado = combinar(acumulado, v)
	}
	return acumulado
}

// ─────────────────────────────────────────────────────────
// MaximoPor: encuentra el elemento con el mayor "valor" según
// una función clave, usando cmp.Ordered para la comparación
// ─────────────────────────────────────────────────────────
func MaximoPor[T any, K cmp.Ordered](lista []T, clave func(T) K) T {
	mejor := lista[0]
	mejorClave := clave(mejor)
	for _, v := range lista[1:] {
		if k := clave(v); k > mejorClave {
			mejor = v
			mejorClave = k
		}
	}
	return mejor
}

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func main() {
	fmt.Println("=== KIOSCO DIGITAL: Filtrar/Mapear/Reducir genéricos ===")

	catalogo := []Producto{
		{Nombre: "Alfajor", Precio: 800, Stock: 15},
		{Nombre: "Gaseosa", Precio: 1200, Stock: 0},
		{Nombre: "Notebook", Precio: 450000, Stock: 3},
		{Nombre: "Mouse", Precio: 8500, Stock: 20},
		{Nombre: "Monitor", Precio: 95000, Stock: 0},
	}

	// ─────────────────────────────────────────────────────────
	// Filtrar: productos CON stock
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Filtrar: productos con stock ===")
	conStock := Filtrar(catalogo, func(p Producto) bool {
		return p.Stock > 0
	})
	for _, p := range conStock {
		fmt.Printf("  %s (stock: %d)\n", p.Nombre, p.Stock)
	}

	// ─────────────────────────────────────────────────────────
	// Mapear: de []Producto a []string (solo nombres)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Mapear: solo los nombres ===")
	nombres := Mapear(catalogo, func(p Producto) string {
		return p.Nombre
	})
	fmt.Println(" ", nombres)

	// Mapear: de []Producto a []float64 (precios con IVA)
	fmt.Println("\n=== Mapear: precios con IVA (21%) ===")
	preciosConIVA := Mapear(catalogo, func(p Producto) float64 {
		return p.Precio * 1.21
	})
	for i, precio := range preciosConIVA {
		fmt.Printf("  %-10s $%.2f\n", catalogo[i].Nombre, precio)
	}

	// ─────────────────────────────────────────────────────────
	// Reducir: valor TOTAL del inventario (precio * stock)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reducir: valor total del inventario ===")
	valorTotal := Reducir(catalogo, 0.0, func(acumulado float64, p Producto) float64 {
		return acumulado + p.Precio*float64(p.Stock)
	})
	fmt.Printf("  Valor total: $%.2f\n", valorTotal)

	// ─────────────────────────────────────────────────────────
	// MaximoPor: el producto más caro
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== MaximoPor: producto más caro ===")
	masCaro := MaximoPor(catalogo, func(p Producto) float64 {
		return p.Precio
	})
	fmt.Printf("  %s a $%.2f\n", masCaro.Nombre, masCaro.Precio)

	// ─────────────────────────────────────────────────────────
	// COMBINANDO TODO: nombres de los productos caros con stock
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Combinando: nombres de productos caros (>$10000) con stock ===")
	caros := Filtrar(catalogo, func(p Producto) bool {
		return p.Precio > 10000 && p.Stock > 0
	})
	nombresCaros := Mapear(caros, func(p Producto) string { return p.Nombre })
	fmt.Println(" ", nombresCaros)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Filtrar[T]        → se queda con lo que cumple una condición")
	fmt.Println("  Mapear[T, R]      → transforma cada elemento (puede cambiar de tipo)")
	fmt.Println("  Reducir[T, R]     → combina todo en UN solo valor")
	fmt.Println("  MaximoPor[T, K]   → el 'mejor' según una clave comparable")
	fmt.Println("  Una sola vez      → funcionan para Producto, string, int, lo que sea")
}
