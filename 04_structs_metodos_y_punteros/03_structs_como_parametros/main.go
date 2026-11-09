package main

import (
	"fmt"
	"strings"
)

// =========================================================
// PARÁMETROS CON STRUCT
// =========================================================
// Los structs son el tipo más común que pasamos como parámetro
// en Go. Hay dos maneras de pasarlos:
//
//   func f(p Producto)    → pasa una COPIA del struct
//   func f(p *Producto)   → copia un puntero que apunta al original
//
// Veremos la diferencia y cuándo usar cada uno.

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Orden struct {
	ID      int
	Cliente string
	Items   []ItemOrden
	Enviada bool
}

type ItemOrden struct {
	Producto Producto
	Cantidad int
}

// ─────────────────────────────────────────────────────────
// FUNCIONES QUE RECIBEN STRUCT POR VALOR (copia)
// ─────────────────────────────────────────────────────────
// Reciben una copia → no pueden modificar el original.
// Útil para funciones que solo LEEN el struct.

func mostrarProducto(p Producto) {
	fmt.Printf("  [ID:%d] %-15s | $%8.2f | Stock: %3d | Categoría: %s\n",
		p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
}

func estaDisponible(p Producto, cantidad int) bool {
	return p.Stock >= cantidad
}

func calcularSubtotal(p Producto, cantidad int) float64 {
	return p.Precio * float64(cantidad)
}

func descripcionProducto(p Producto) string {
	disponibilidad := "✓ Disponible"
	if p.Stock == 0 {
		disponibilidad = "✗ Sin stock"
	} else if p.Stock <= 5 {
		disponibilidad = "⚠ Pocas unidades"
	}
	return fmt.Sprintf("%s — $%.2f [%s]", p.Nombre, p.Precio, disponibilidad)
}

// ─────────────────────────────────────────────────────────
// FUNCIONES QUE RECIBEN UN PUNTERO AL STRUCT
// ─────────────────────────────────────────────────────────
// Reciben el puntero → PUEDEN modificar el original.
// Útil cuando la función necesita cambiar el struct.
// También más eficiente para structs grandes (no copia toda la memoria).

func aplicarDescuento(p *Producto, porcentaje float64) {
	descuento := p.Precio * porcentaje
	p.Precio -= descuento // modifica el precio del ORIGINAL
	fmt.Printf("  Descuento %.0f%% aplicado: %s → $%.2f\n",
		porcentaje*100, p.Nombre, p.Precio)
}

func reducirStock(p *Producto, cantidad int) bool {
	if p.Stock < cantidad {
		return false // no hay suficiente stock
	}
	p.Stock -= cantidad // modifica el stock del ORIGINAL
	return true
}

func reponerStock(p *Producto, cantidad int) {
	p.Stock += cantidad
	fmt.Printf("  Repuesto: %s → nuevo stock: %d\n", p.Nombre, p.Stock)
}

// ─────────────────────────────────────────────────────────
// FUNCIÓN QUE RECIBE ORDEN (struct más complejo)
// ─────────────────────────────────────────────────────────
func calcularTotalOrden(o Orden) float64 {
	total := 0.0
	for _, item := range o.Items {
		total += calcularSubtotal(item.Producto, item.Cantidad)
	}
	return total
}

func mostrarOrden(o Orden) {
	fmt.Printf("\n  === Orden #%d — Cliente: %s ===\n", o.ID, o.Cliente)
	fmt.Printf("  %-15s %8s %6s %10s\n", "Producto", "Precio", "Cant.", "Subtotal")
	fmt.Println("  " + repetir("-", 45))
	for _, item := range o.Items {
		sub := calcularSubtotal(item.Producto, item.Cantidad)
		fmt.Printf("  %-15s %8.2f %6d %10.2f\n",
			item.Producto.Nombre, item.Producto.Precio, item.Cantidad, sub)
	}
	fmt.Println("  " + repetir("-", 45))
	fmt.Printf("  %-30s %10.2f\n", "TOTAL:", calcularTotalOrden(o))
	estado := "Pendiente"
	if o.Enviada {
		estado = "Enviada ✓"
	}
	fmt.Printf("  Estado: %s\n", estado)
}

func enviarOrden(o *Orden) {
	o.Enviada = true
	fmt.Printf("  Orden #%d marcada como enviada\n", o.ID)
}

func repetir(s string, n int) string {
	return strings.Repeat(s, n)
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║   PARÁMETROS CON STRUCTS      ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// CATÁLOGO
	// ─────────────────────────────────────────────────────────
	notebook := Producto{1, "Notebook", 1500.00, 5, "computación"}
	mouse := Producto{2, "Mouse", 25.99, 42, "periférico"}
	teclado := Producto{3, "Teclado", 75.50, 3, "periférico"}

	fmt.Println("\n--- Catálogo inicial ---")
	mostrarProducto(notebook)
	mostrarProducto(mouse)
	mostrarProducto(teclado)

	// ─────────────────────────────────────────────────────────
	// PASO POR VALOR: solo lectura
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Funciones de solo lectura (por valor) ---")
	fmt.Println(descripcionProducto(notebook))
	fmt.Println(descripcionProducto(teclado)) // stock bajo

	fmt.Printf("\n¿Hay 3 Teclados disponibles? %v\n", estaDisponible(teclado, 3))
	fmt.Printf("¿Hay 5 Teclados disponibles? %v\n", estaDisponible(teclado, 5))
	fmt.Printf("Subtotal 2 notebooks: $%.2f\n", calcularSubtotal(notebook, 2))

	// ─────────────────────────────────────────────────────────
	// PASO POR PUNTERO: modifica el original
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Modificar struct (por puntero) ---")
	fmt.Printf("Precio notebook antes: $%.2f\n", notebook.Precio)
	aplicarDescuento(&notebook, 0.10) // pasamos la dirección con &
	fmt.Printf("Precio notebook después: $%.2f\n", notebook.Precio)

	fmt.Printf("\nStock teclado antes: %d\n", teclado.Stock)
	ok := reducirStock(&teclado, 2)
	fmt.Printf("¿Reducción exitosa? %v | Stock después: %d\n", ok, teclado.Stock)

	okFallo := reducirStock(&teclado, 5) // no hay stock suficiente
	fmt.Printf("¿Reducción de 5 exitosa? %v | Stock sin cambio: %d\n", okFallo, teclado.Stock)

	reponerStock(&mouse, 10)

	// ─────────────────────────────────────────────────────────
	// STRUCT COMPLEJO COMO PARÁMETRO
	// ─────────────────────────────────────────────────────────
	orden := Orden{
		ID:      1001,
		Cliente: "Ana García",
		Items: []ItemOrden{
			{Producto: notebook, Cantidad: 1},
			{Producto: mouse, Cantidad: 2},
			{Producto: teclado, Cantidad: 1},
		},
		Enviada: false,
	}

	mostrarOrden(orden)

	fmt.Println("\n--- Procesando orden ---")
	enviarOrden(&orden) // modifica Enviada en el original
	mostrarOrden(orden) // ahora muestra "Enviada ✓"

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("func f(p Producto)    → copia, para lectura o structs pequeños")
	fmt.Println("func f(p *Producto)   → puntero, para modificar o structs grandes")
	fmt.Println("Pasar con &: f(&miProducto)")
	fmt.Println("Go auto-derefencia: p.Campo en vez de (*p).Campo")
}
