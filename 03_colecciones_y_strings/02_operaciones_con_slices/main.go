package main

import "fmt"

// =========================================================
// AGREGAR Y ELIMINAR ELEMENTOS A LA ORDEN
// =========================================================
// Construimos un sistema de carrito de compras real.
// Aprendemos a pasar slices a funciones y a modificarlos.
//
// REGLA IMPORTANTE QUE VEREMOS:
// En Go, los slices se pasan "por referencia" en cuanto a su
// contenido, PERO el slice header (len, cap, puntero) se copia.
// Esto significa que append() dentro de una función puede NO
// reflejarse afuera. Por eso, muchas funciones retornan el
// slice modificado.

// ─────────────────────────────────────────────────────────
// TIPOS
// ─────────────────────────────────────────────────────────
type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

type ItemCarrito struct {
	Producto Producto
	Cantidad int
}

// ─────────────────────────────────────────────────────────
// FUNCIONES DEL CARRITO
// ─────────────────────────────────────────────────────────

// agregarItem agrega un producto al carrito y retorna el carrito actualizado.
// Retornamos el slice porque append() puede crear uno nuevo internamente.
func agregarItem(carrito []ItemCarrito, producto Producto, cantidad int) []ItemCarrito {
	// Buscamos si el producto ya está en el carrito
	for i, item := range carrito {
		if item.Producto.ID == producto.ID {
			// Ya existe: solo sumamos la cantidad
			carrito[i].Cantidad += cantidad
			fmt.Printf("  ✓ Actualizado: %s (cantidad total: %d)\n",
				producto.Nombre, carrito[i].Cantidad)
			return carrito
		}
	}
	// No existe: agregamos un item nuevo
	nuevoItem := ItemCarrito{
		Producto: producto,
		Cantidad: cantidad,
	}
	fmt.Printf("  ✓ Agregado: %s × %d\n", producto.Nombre, cantidad)
	return append(carrito, nuevoItem) // ← por eso retornamos
}

// eliminarItem elimina un producto del carrito por su ID.
// Retornamos el nuevo slice porque su longitud cambia.
func eliminarItem(carrito []ItemCarrito, productoID int) []ItemCarrito {
	for i, item := range carrito {
		if item.Producto.ID == productoID {
			nombre := item.Producto.Nombre
			// Técnica de eliminación: unir todo lo que está antes y después del índice
			carrito = append(carrito[:i], carrito[i+1:]...)
			fmt.Printf("  ✗ Eliminado: %s\n", nombre)
			return carrito
		}
	}
	fmt.Println("  ⚠ Producto no encontrado en el carrito")
	return carrito
}

// actualizarCantidad modifica la cantidad de un item existente.
// Si la cantidad llega a 0, elimina el item.
func actualizarCantidad(carrito []ItemCarrito, productoID, nuevaCantidad int) []ItemCarrito {
	if nuevaCantidad <= 0 {
		return eliminarItem(carrito, productoID)
	}
	for i, item := range carrito {
		if item.Producto.ID == productoID {
			carrito[i].Cantidad = nuevaCantidad
			fmt.Printf("  ✎ Actualizado: %s → cantidad: %d\n",
				item.Producto.Nombre, nuevaCantidad)
			return carrito
		}
	}
	fmt.Println("  ⚠ Producto no encontrado")
	return carrito
}

// calcularTotal recorre el carrito y suma precio × cantidad de cada item.
func calcularTotal(carrito []ItemCarrito) float64 {
	total := 0.0
	for _, item := range carrito {
		total += item.Producto.Precio * float64(item.Cantidad)
	}
	return total
}

// mostrarCarrito imprime el carrito en formato tabla.
func mostrarCarrito(carrito []ItemCarrito) {
	if len(carrito) == 0 {
		fmt.Println("  [Carrito vacío]")
		return
	}
	fmt.Printf("  %-15s %8s %8s %10s\n", "Producto", "Precio", "Cant.", "Subtotal")
	fmt.Println("  " + linea(46))
	for _, item := range carrito {
		subtotal := item.Producto.Precio * float64(item.Cantidad)
		fmt.Printf("  %-15s %8.2f %8d %10.2f\n",
			item.Producto.Nombre,
			item.Producto.Precio,
			item.Cantidad,
			subtotal)
	}
	fmt.Println("  " + linea(46))
	fmt.Printf("  %-24s %10s %10.2f\n", "", "TOTAL:", calcularTotal(carrito))
}

func linea(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "-"
	}
	return s
}

func main() {
	// ─────────────────────────────────────────────────────────
	// CATÁLOGO DE PRODUCTOS
	// ─────────────────────────────────────────────────────────
	catalogo := []Producto{
		{ID: 1, Nombre: "Notebook", Precio: 1500.00},
		{ID: 2, Nombre: "Mouse", Precio: 25.99},
		{ID: 3, Nombre: "Teclado", Precio: 75.50},
		{ID: 4, Nombre: "Monitor", Precio: 450.00},
		{ID: 5, Nombre: "Auriculares", Precio: 80.00},
	}

	// ─────────────────────────────────────────────────────────
	// SIMULACIÓN DE COMPRA
	// ─────────────────────────────────────────────────────────
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║      TIENDA GO — CARRITO      ║")
	fmt.Println("╚══════════════════════════════╝")

	// Carrito vacío (slice nil → zero value de []ItemCarrito)
	var carrito []ItemCarrito

	fmt.Println("\n--- Agregando productos ---")
	carrito = agregarItem(carrito, catalogo[0], 1) // Notebook × 1
	carrito = agregarItem(carrito, catalogo[1], 2) // Mouse × 2
	carrito = agregarItem(carrito, catalogo[2], 1) // Teclado × 1

	fmt.Println("\nCarrito actual:")
	mostrarCarrito(carrito)

	fmt.Println("\n--- Agregar producto ya existente ---")
	carrito = agregarItem(carrito, catalogo[1], 1) // Mouse → ahora 3

	fmt.Println("\nCarrito actualizado:")
	mostrarCarrito(carrito)

	fmt.Println("\n--- Actualizar cantidad ---")
	carrito = actualizarCantidad(carrito, 3, 2) // Teclado → cantidad 2

	fmt.Println("\n--- Eliminar un producto ---")
	carrito = eliminarItem(carrito, 2) // eliminar Mouse

	fmt.Println("\nCarrito final:")
	mostrarCarrito(carrito)

	fmt.Println("\n--- Eliminar producto inexistente ---")
	carrito = eliminarItem(carrito, 99)

	fmt.Println("\n--- Vaciar carrito con cantidad 0 ---")
	carrito = actualizarCantidad(carrito, 1, 0) // Notebook → se elimina
	carrito = actualizarCantidad(carrito, 3, 0) // Teclado → se elimina

	fmt.Println("\nCarrito vacío:")
	mostrarCarrito(carrito)

	// ─────────────────────────────────────────────────────────
	// LECCIÓN APRENDIDA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Lección clave ===")
	fmt.Println("Las funciones que modifican slices con append() deben RETORNAR el slice.")
	fmt.Println("El llamador debe reasignar: carrito = agregarItem(carrito, ...)")
	fmt.Println()
	fmt.Println("Modificar elementos existentes (carrito[i].Campo = x) SÍ funciona")
	fmt.Println("sin retornar porque no cambia el header del slice.")
}
