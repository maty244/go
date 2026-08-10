// Package catalogo maneja el listado de productos del kiosco.
package catalogo

import "fmt"

// ─────────────────────────────────────────────────────────
// MAYÚSCULA = EXPORTADO (público, visible desde otros paquetes)
// MINÚSCULA = NO exportado (privado, solo visible DENTRO de este paquete)
// ─────────────────────────────────────────────────────────
// Esta es la ÚNICA regla de visibilidad en Go. No hay palabras
// clave "public"/"private" como en otros lenguajes: todo depende
// de la PRIMERA LETRA del nombre.
//
// Aplica a todo: funciones, tipos, campos de structs, variables,
// constantes.

// Producto es EXPORTADO (empieza con mayúscula): otros paquetes
// pueden crear y usar valores de este tipo.
type Producto struct {
	Nombre string // exportado: otros paquetes pueden leer/escribir
	Precio float64
	stock  int // NO exportado: invisible fuera de este paquete
}

// margenGanancia es una constante NO exportada: es un detalle
// interno de cómo calculamos el stock mínimo, no le interesa a
// quien solo quiere USAR el catálogo.
const margenGanancia = 0.30

// catalogoInterno es una variable NO exportada: el paquete la usa
// para guardar el estado, pero nadie de afuera puede tocarla
// directamente. Así evitamos que otro paquete la deje en un
// estado inconsistente sin pasar por nuestras funciones.
var catalogoInterno = map[string]*Producto{}

// NuevoProducto es EXPORTADA: es la forma correcta (y única) de
// crear un producto y agregarlo al catálogo interno.
func NuevoProducto(nombre string, precio float64, stockInicial int) *Producto {
	p := &Producto{Nombre: nombre, Precio: precio, stock: stockInicial}
	catalogoInterno[nombre] = p
	return p
}

// Vender es EXPORTADA: reduce el stock validando que haya suficiente.
// Es la ÚNICA forma de modificar "stock" desde afuera, porque el
// campo en sí es privado.
func (p *Producto) Vender(cantidad int) error {
	if cantidad > p.stock {
		return fmt.Errorf("stock insuficiente de %s: pediste %d, hay %d",
			p.Nombre, cantidad, p.stock)
	}
	p.stock -= cantidad
	if necesitaReponer(p) {
		fmt.Printf("  [aviso interno] %s está por debajo del margen, reponer\n", p.Nombre)
	}
	return nil
}

// Stock es EXPORTADA: una función de SOLO LECTURA para exponer
// el valor de "stock" sin permitir que lo modifiquen directamente
// desde afuera (no hay setter, solo Vender() con su validación).
func (p *Producto) Stock() int {
	return p.stock
}

// necesitaReponer es un detalle interno: usa margenGanancia (que
// tampoco es exportada) para decidir si hay que avisar reposición.
// Nada de esto le importa a quien solo usa el paquete catalogo.
func necesitaReponer(p *Producto) bool {
	minimo := int(float64(p.stock) * margenGanancia)
	return p.stock <= minimo+1
}
