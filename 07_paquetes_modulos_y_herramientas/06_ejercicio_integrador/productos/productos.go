// Package productos maneja el catálogo del kiosco: crear productos,
// consultarlos y vender, protegiendo el stock de modificaciones inválidas.
package productos

import "fmt"

// Producto representa un ítem del catálogo. Nombre y Precio son
// exportados porque cualquiera que use el paquete necesita leerlos;
// stock es privado, se protege con Vender/Stock.
type Producto struct {
	Nombre string
	Precio float64
	stock  int
}

// Catalogo agrupa productos por nombre.
type Catalogo struct {
	items map[string]*Producto
}

// NuevoCatalogo crea un catálogo vacío.
func NuevoCatalogo() *Catalogo {
	return &Catalogo{items: map[string]*Producto{}}
}

// Agregar registra un producto nuevo en el catálogo.
func (c *Catalogo) Agregar(nombre string, precio float64, stockInicial int) *Producto {
	p := &Producto{Nombre: nombre, Precio: precio, stock: stockInicial}
	c.items[nombre] = p
	return p
}

// Buscar devuelve el producto por nombre, si existe.
func (c *Catalogo) Buscar(nombre string) (*Producto, bool) {
	p, ok := c.items[nombre]
	return p, ok
}

// Stock expone el stock actual de solo lectura.
func (p *Producto) Stock() int {
	return p.stock
}

// Vender reduce el stock, validando que alcance.
func (p *Producto) Vender(cantidad int) error {
	if cantidad > p.stock {
		return fmt.Errorf("stock insuficiente de %s: pediste %d, hay %d",
			p.Nombre, cantidad, p.stock)
	}
	p.stock -= cantidad
	return nil
}
