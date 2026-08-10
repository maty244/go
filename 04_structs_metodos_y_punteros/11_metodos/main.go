package main

import (
	"fmt"
	"strings"
)

// =========================================================
// MÉTODOS VS FUNCIONES + MÉTODOS EN GO (partes 1 y 2)
// =========================================================
// FUNCIÓN:  código independiente, no está asociado a ningún tipo.
//           func calcularTotal(p Producto) float64 { }
//
// MÉTODO:   función con un "receiver" (receptor) — está asociada
//           a un tipo específico y se llama sobre una instancia.
//           func (p Producto) Total() float64 { }
//           llamado como: miProducto.Total()
//
// ¿Cuándo usar cada uno?
//   - Método:   la operación PERTENECE al tipo (es su comportamiento)
//   - Función:  operación utilitaria que no pertenece a ningún tipo
//
// Receivers:
//   (p Producto)  → por valor, no modifica el original
//   (p *Producto) → por puntero, puede modificar el original

// ─────────────────────────────────────────────────────────
// TIPOS DE NUESTRO SISTEMA
// ─────────────────────────────────────────────────────────

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type ItemCarrito struct {
	Producto Producto
	Cantidad int
}

type Carrito struct {
	ID      int
	Cliente string
	Items   []ItemCarrito
}

// ─────────────────────────────────────────────────────────
// MÉTODOS DEL TIPO Producto
// ─────────────────────────────────────────────────────────

// String() implementa la interfaz fmt.Stringer
// Cuando hacés fmt.Println(producto), usa este método automáticamente
func (p Producto) String() string {
	return fmt.Sprintf("[%d] %-15s $%.2f (stock: %d)", p.ID, p.Nombre, p.Precio, p.Stock)
}

// EstaDisponible verifica si hay suficiente stock
func (p Producto) EstaDisponible(cantidad int) bool {
	return p.Stock >= cantidad
}

// PrecioConIVA retorna el precio con impuesto (no modifica, receiver por valor)
func (p Producto) PrecioConIVA() float64 {
	return p.Precio * 1.21
}

// Describir retorna una descripción formateada del producto
func (p Producto) Describir() string {
	disponibilidad := "disponible"
	if p.Stock == 0 {
		disponibilidad = "agotado"
	} else if p.Stock <= 3 {
		disponibilidad = fmt.Sprintf("últimas %d unidades", p.Stock)
	}
	return fmt.Sprintf("%s — $%.2f (%s)", p.Nombre, p.Precio, disponibilidad)
}

// AplicarDescuento MODIFICA el precio → receiver por puntero
func (p *Producto) AplicarDescuento(porcentaje float64) {
	p.Precio = p.Precio * (1 - porcentaje)
}

// ReducirStock MODIFICA el stock → receiver por puntero
func (p *Producto) ReducirStock(cantidad int) error {
	if p.Stock < cantidad {
		return fmt.Errorf("stock insuficiente: pedido %d, disponible %d", cantidad, p.Stock)
	}
	p.Stock -= cantidad
	return nil
}

// ReponerStock suma stock → receiver por puntero
func (p *Producto) ReponerStock(cantidad int) {
	p.Stock += cantidad
}

// ─────────────────────────────────────────────────────────
// MÉTODOS DEL TIPO ItemCarrito
// ─────────────────────────────────────────────────────────

func (i ItemCarrito) Subtotal() float64 {
	return i.Producto.Precio * float64(i.Cantidad)
}

func (i ItemCarrito) String() string {
	return fmt.Sprintf("  %-15s × %d = $%.2f",
		i.Producto.Nombre, i.Cantidad, i.Subtotal())
}

// ─────────────────────────────────────────────────────────
// MÉTODOS DEL TIPO Carrito
// ─────────────────────────────────────────────────────────

// Agregar agrega un producto al carrito
func (c *Carrito) Agregar(p Producto, cantidad int) {
	for i, item := range c.Items {
		if item.Producto.ID == p.ID {
			c.Items[i].Cantidad += cantidad
			return
		}
	}
	c.Items = append(c.Items, ItemCarrito{p, cantidad})
}

// Eliminar elimina un producto por ID
func (c *Carrito) Eliminar(productoID int) {
	for i, item := range c.Items {
		if item.Producto.ID == productoID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}

// Total calcula el total del carrito
func (c Carrito) Total() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Subtotal()
	}
	return total
}

// CantidadItems retorna cuántos items hay
func (c Carrito) CantidadItems() int {
	return len(c.Items)
}

// Mostrar imprime el carrito completo
func (c Carrito) Mostrar() {
	fmt.Printf("\n  ╔═ Carrito #%d — %s ═╗\n", c.ID, c.Cliente)
	if len(c.Items) == 0 {
		fmt.Println("  ║  (carrito vacío)")
	}
	for _, item := range c.Items {
		fmt.Println(" ", item)
	}
	fmt.Println("  ╠════════════════════════╣")
	fmt.Printf("  ║  TOTAL: $%.2f (%d items)\n", c.Total(), c.CantidadItems())
	fmt.Println("  ╚════════════════════════╝")
}

// ─────────────────────────────────────────────────────────
// FUNCIONES (no son métodos, operan sobre varios tipos)
// ─────────────────────────────────────────────────────────

// Esta es una FUNCIÓN, no un método.
// Compara dos productos por precio.
func compararPrecio(a, b Producto) int {
	switch {
	case a.Precio < b.Precio:
		return -1
	case a.Precio > b.Precio:
		return 1
	default:
		return 0
	}
}

// Buscar en catálogo: es una función, no pertenece a Producto
func buscarEnCatalogo(catalogo []Producto, nombre string) (Producto, bool) {
	nombreLower := strings.ToLower(nombre)
	for _, p := range catalogo {
		if strings.ToLower(p.Nombre) == nombreLower {
			return p, true
		}
	}
	return Producto{}, false
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  MÉTODOS VS FUNCIONES EN GO   ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// USANDO MÉTODOS DE Producto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Métodos de Producto ===")

	notebook := Producto{1, "Notebook", 1500.00, 5, "computación"}
	mouse := Producto{2, "Mouse", 25.99, 42, "periférico"}
	teclado := Producto{3, "Teclado", 75.50, 2, "periférico"}

	// String() se llama automáticamente con Println
	fmt.Println(notebook)
	fmt.Println(mouse)

	fmt.Println("\nDescripciones:")
	fmt.Println(" ", notebook.Describir())
	fmt.Println(" ", teclado.Describir()) // últimas 2 unidades

	fmt.Printf("\n¿Hay 3 notebooks? %v\n", notebook.EstaDisponible(3))
	fmt.Printf("¿Hay 10 notebooks? %v\n", notebook.EstaDisponible(10))
	fmt.Printf("Notebook con IVA: $%.2f\n", notebook.PrecioConIVA())

	// Métodos que MODIFICAN (receiver por puntero)
	fmt.Printf("\nDescuento 15%% en notebook (precio antes: $%.2f)\n", notebook.Precio)
	notebook.AplicarDescuento(0.15)
	fmt.Printf("Precio después: $%.2f\n", notebook.Precio)

	if err := teclado.ReducirStock(2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Teclado vendido: stock restante %d\n", teclado.Stock)
	}

	if err := teclado.ReducirStock(1); err != nil {
		fmt.Println("Stock agotado:", err)
	}

	teclado.ReponerStock(10)
	fmt.Printf("Teclado repuesto: stock %d\n", teclado.Stock)

	// ─────────────────────────────────────────────────────────
	// USANDO MÉTODOS DEL CARRITO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Carrito con métodos ===")

	carrito := Carrito{ID: 1001, Cliente: "Ana García"}
	carrito.Agregar(notebook, 1)
	carrito.Agregar(mouse, 2)
	carrito.Agregar(teclado, 1)
	carrito.Mostrar()

	fmt.Println("\nEliminando mouse...")
	carrito.Eliminar(2)
	carrito.Mostrar()

	// Agregar producto ya existente (suma cantidad)
	fmt.Println("\nAgregando otro notebook...")
	carrito.Agregar(notebook, 1)
	carrito.Mostrar()

	// ─────────────────────────────────────────────────────────
	// FUNCIONES vs MÉTODOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Función vs Método (comparación) ===")

	// Función: no pertenece a ningún tipo
	resultado := compararPrecio(notebook, mouse)
	if resultado > 0 {
		fmt.Printf("Función: %s es más caro que %s\n", notebook.Nombre, mouse.Nombre)
	}

	// Método en Producto: sí pertenece al tipo
	fmt.Printf("Método: %s está disponible en 5? %v\n",
		notebook.Nombre, notebook.EstaDisponible(5))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Método (valor):    func (p Producto) Nombre() → no modifica")
	fmt.Println("Método (puntero):  func (p *Producto) Nombre() → puede modificar")
	fmt.Println("Llamada:           miProducto.Nombre()")
	fmt.Println("Función:           func Nombre(p Producto) → no está en el tipo")
	fmt.Println()
	fmt.Println("Usar MÉTODO cuando la operación es el comportamiento del tipo.")
	fmt.Println("Usar FUNCIÓN cuando opera sobre múltiples tipos o es utilitaria.")
	fmt.Println()
	fmt.Println("fmt.Stringer: implementar String() string → fmt usa el método automáticamente")
}
