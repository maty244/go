// Package clientes maneja el registro de clientes del kiosco
// y su historial de compras.
package clientes

import "fmt"

// Cliente representa a un comprador registrado.
type Cliente struct {
	Nombre  string
	compras []string // historial, privado: solo se modifica con RegistrarCompra
}

// Registro agrupa clientes por nombre.
type Registro struct {
	clientes map[string]*Cliente
}

// NuevoRegistro crea un registro de clientes vacío.
func NuevoRegistro() *Registro {
	return &Registro{clientes: map[string]*Cliente{}}
}

// Alta registra un cliente nuevo (o devuelve el existente).
func (r *Registro) Alta(nombre string) *Cliente {
	if c, ok := r.clientes[nombre]; ok {
		return c
	}
	c := &Cliente{Nombre: nombre}
	r.clientes[nombre] = c
	return c
}

// RegistrarCompra agrega una compra al historial del cliente.
func (c *Cliente) RegistrarCompra(descripcion string) {
	c.compras = append(c.compras, descripcion)
}

// Historial devuelve las compras registradas (solo lectura).
func (c *Cliente) Historial() []string {
	return c.compras
}

// ResumenCliente arma un texto legible del historial de un cliente.
func ResumenCliente(c *Cliente) string {
	if len(c.compras) == 0 {
		return fmt.Sprintf("%s todavía no compró nada", c.Nombre)
	}
	return fmt.Sprintf("%s compró %d producto(s): %v", c.Nombre, len(c.compras), c.compras)
}
