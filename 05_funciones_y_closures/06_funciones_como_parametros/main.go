package main

import (
	"fmt"
	"strings"
)

// =========================================================
// FUNCIONES COMO PARÁMETROS (Higher-Order Functions)
// =========================================================
// Una función de orden superior (HOF) es una función que:
//   a) Recibe una o más funciones como parámetros, y/o
//   b) Retorna una función
//
// Esto permite escribir código genérico y reutilizable.
// Es el corazón de la programación funcional en Go.

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

// ─────────────────────────────────────────────────────────
// FUNCIONES DE ORDEN SUPERIOR BÁSICAS
// ─────────────────────────────────────────────────────────

// filtrar retorna los productos que cumplen la condición
func filtrar(productos []Producto, condicion func(Producto) bool) []Producto {
	var resultado []Producto
	for _, p := range productos {
		if condicion(p) {
			resultado = append(resultado, p)
		}
	}
	return resultado
}

// transformar aplica una función a cada producto y retorna el nuevo slice
func transformar(productos []Producto, fn func(Producto) Producto) []Producto {
	resultado := make([]Producto, len(productos))
	for i, p := range productos {
		resultado[i] = fn(p)
	}
	return resultado
}

// reducir acumula un valor recorriendo todos los productos
func reducir(productos []Producto, inicial float64, fn func(float64, Producto) float64) float64 {
	acum := inicial
	for _, p := range productos {
		acum = fn(acum, p)
	}
	return acum
}

// forEach ejecuta una acción sobre cada producto (sin retorno)
func forEach(productos []Producto, accion func(Producto)) {
	for _, p := range productos {
		accion(p)
	}
}

// ─────────────────────────────────────────────────────────
// FUNCIONES CON IMPUESTO COMO PARÁMETRO
// ─────────────────────────────────────────────────────────

// calcularPrecioFinal aplica un cálculo de impuesto/precio parametrizable
func calcularPrecioFinal(p Producto, calcularImpuesto func(float64) float64) float64 {
	return p.Precio + calcularImpuesto(p.Precio)
}

// mostrarCatalogo muestra el catálogo aplicando un formateador custom
func mostrarCatalogo(productos []Producto, formatear func(Producto) string) {
	for _, p := range productos {
		fmt.Println(" ", formatear(p))
	}
}

func mostrarProductos(titulo string, ps []Producto) {
	fmt.Printf("\n  --- %s (%d items) ---\n", titulo, len(ps))
	for _, p := range ps {
		fmt.Printf("  [%d] %-15s $%8.2f | stock: %2d | %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
	if len(ps) == 0 {
		fmt.Println("  (sin resultados)")
	}
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  FUNCIONES COMO PARÁMETROS    ║")
	fmt.Println("╚══════════════════════════════╝")

	catalogo := []Producto{
		{1, "Notebook", 1500.00, 5, "computación"},
		{2, "Mouse", 25.99, 42, "periférico"},
		{3, "Teclado", 75.50, 3, "periférico"},
		{4, "Monitor", 450.00, 0, "pantalla"},
		{5, "Auriculares", 80.00, 15, "audio"},
		{6, "Tablet", 350.00, 8, "computación"},
		{7, "Impresora", 200.00, 2, "periférico"},
	}

	// ─────────────────────────────────────────────────────────
	// FILTRAR CON DISTINTAS CONDICIONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== filtrar (HOF) ===")

	// Filtrar por precio
	baratos := filtrar(catalogo, func(p Producto) bool {
		return p.Precio < 100
	})
	mostrarProductos("Menos de $100", baratos)

	// Filtrar por stock
	conStock := filtrar(catalogo, func(p Producto) bool {
		return p.Stock > 0
	})
	mostrarProductos("Con stock disponible", conStock)

	// Filtrar por categoría
	perifericos := filtrar(catalogo, func(p Producto) bool {
		return p.Categoria == "periférico"
	})
	mostrarProductos("Periféricos", perifericos)

	// Filtrar con condición compuesta
	ofertas := filtrar(catalogo, func(p Producto) bool {
		return p.Precio >= 100 && p.Precio <= 500 && p.Stock > 0
	})
	mostrarProductos("Ofertas ($100-$500 con stock)", ofertas)

	// ─────────────────────────────────────────────────────────
	// TRANSFORMAR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== transformar (HOF) ===")

	// Aplicar IVA a todos los precios
	conIVA := transformar(catalogo, func(p Producto) Producto {
		p.Precio = p.Precio * 1.21
		return p
	})
	mostrarProductos("Catálogo con IVA", conIVA[:3]) // mostramos solo 3

	// Normalizar nombres (todo mayúsculas)
	normalizados := transformar(catalogo, func(p Producto) Producto {
		p.Nombre = strings.ToUpper(p.Nombre)
		return p
	})
	fmt.Println("\n  Nombres normalizados:")
	for _, p := range normalizados {
		fmt.Printf("    %s\n", p.Nombre)
	}

	// ─────────────────────────────────────────────────────────
	// REDUCIR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== reducir (HOF) ===")

	// Suma de todos los precios
	totalPrecios := reducir(catalogo, 0, func(acc float64, p Producto) float64 {
		return acc + p.Precio
	})
	fmt.Printf("  Suma de precios: $%.2f\n", totalPrecios)

	// Precio máximo
	precioMax := reducir(catalogo, 0, func(max float64, p Producto) float64 {
		if p.Precio > max {
			return p.Precio
		}
		return max
	})
	fmt.Printf("  Precio máximo: $%.2f\n", precioMax)

	// Valor total del inventario (precio × stock)
	valorInventario := reducir(catalogo, 0, func(acc float64, p Producto) float64 {
		return acc + p.Precio*float64(p.Stock)
	})
	fmt.Printf("  Valor total inventario: $%.2f\n", valorInventario)

	// ─────────────────────────────────────────────────────────
	// FUNCIONES CON IMPUESTO PARAMETRIZABLE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== calcularPrecioFinal con impuesto parametrizable ===")

	notebook := catalogo[0]
	tablet := catalogo[5]

	// Con IVA 21%
	conIVAFn := func(precio float64) float64 { return precio * 0.21 }
	// Con IVA reducido 10%
	conIVAReducido := func(precio float64) float64 { return precio * 0.10 }
	// Sin impuesto
	sinImpuesto := func(precio float64) float64 { return 0 }

	fmt.Printf("  Notebook + IVA 21%%:       $%.2f\n", calcularPrecioFinal(notebook, conIVAFn))
	fmt.Printf("  Notebook + IVA reducido:  $%.2f\n", calcularPrecioFinal(notebook, conIVAReducido))
	fmt.Printf("  Tablet sin impuesto:      $%.2f\n", calcularPrecioFinal(tablet, sinImpuesto))

	// ─────────────────────────────────────────────────────────
	// FORMATEADOR PERSONALIZADO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== mostrarCatalogo con formato custom ===")

	formatoSimple := func(p Producto) string {
		return fmt.Sprintf("%-15s $%.2f", p.Nombre, p.Precio)
	}
	formatoCompleto := func(p Producto) string {
		estado := "✓"
		if p.Stock == 0 {
			estado = "✗"
		}
		return fmt.Sprintf("[%s] %-15s $%-8.2f (stock: %d)", estado, p.Nombre, p.Precio, p.Stock)
	}

	fmt.Println("  Formato simple:")
	mostrarCatalogo(catalogo[:3], formatoSimple)

	fmt.Println("  Formato completo:")
	mostrarCatalogo(catalogo, formatoCompleto)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("func filtrar(s []T, cond func(T) bool) []T  → patrón filter")
	fmt.Println("func transformar(s []T, fn func(T) T) []T   → patrón map")
	fmt.Println("func reducir(s []T, ini V, fn func(V,T) V)  → patrón reduce")
	fmt.Println()
	fmt.Println("Ventaja: el comportamiento se inyecta como argumento,")
	fmt.Println("la estructura del algoritmo queda separada de la lógica de negocio.")
}
