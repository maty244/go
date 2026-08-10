package main

import (
	"errors"
	"fmt"
)

// =========================================================
// EJERCICIO INTEGRADOR: KIOSCO DIGITAL
// =========================================================
// Combinamos todo lo visto en Interfaces: contratos, implementación
// implícita, Stringer, type assertion, type switch, any y error
// como interfaz. Armamos el checkout de un kiosco con:
//   - Productos con precio (Stringer para mostrarlos lindo)
//   - Métodos de pago intercambiables (polimorfismo)
//   - Notificaciones por distintos canales
//   - Un error personalizado para stock insuficiente
//   - Un "historial" que procesa eventos de tipos distintos

// ─────────────────────────────────────────────────────────
// PRODUCTO (con Stringer)
// ─────────────────────────────────────────────────────────

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func (p Producto) String() string {
	return fmt.Sprintf("%s ($%.2f, stock: %d)", p.Nombre, p.Precio, p.Stock)
}

// ─────────────────────────────────────────────────────────
// ERROR PERSONALIZADO
// ─────────────────────────────────────────────────────────

type ErrorStockInsuficiente struct {
	Producto string
	Pedido   int
	Stock    int
}

func (e *ErrorStockInsuficiente) Error() string {
	return fmt.Sprintf("stock insuficiente de %s (pedido: %d, stock: %d)",
		e.Producto, e.Pedido, e.Stock)
}

// ─────────────────────────────────────────────────────────
// MÉTODO DE PAGO (interfaz + dos implementaciones)
// ─────────────────────────────────────────────────────────

type MetodoPago interface {
	Pagar(monto float64) error
	String() string
}

type Efectivo struct{}

func (Efectivo) Pagar(monto float64) error { return nil } // siempre funciona
func (Efectivo) String() string            { return "Efectivo" }

type MercadoPago struct {
	Saldo float64
}

func (m *MercadoPago) Pagar(monto float64) error {
	if monto > m.Saldo {
		return fmt.Errorf("mercadopago: saldo insuficiente ($%.2f < $%.2f)", m.Saldo, monto)
	}
	m.Saldo -= monto
	return nil
}
func (m *MercadoPago) String() string { return "MercadoPago" }

// ─────────────────────────────────────────────────────────
// NOTIFICACIÓN (interfaz)
// ─────────────────────────────────────────────────────────

type Notificador interface {
	Notificar(mensaje string)
}

type NotificadorConsola struct{}

func (NotificadorConsola) Notificar(mensaje string) {
	fmt.Println("  [Notificación]", mensaje)
}

// ─────────────────────────────────────────────────────────
// EVENTOS DEL HISTORIAL (para el type switch)
// ─────────────────────────────────────────────────────────

type Evento interface {
	Resumen() string
}

type EventoVenta struct {
	Producto string
	Total    float64
}

func (e EventoVenta) Resumen() string {
	return fmt.Sprintf("Venta: %s por $%.2f", e.Producto, e.Total)
}

type EventoError struct {
	Motivo string
}

func (e EventoError) Resumen() string {
	return fmt.Sprintf("Error: %s", e.Motivo)
}

// ─────────────────────────────────────────────────────────
// EL KIOSCO
// ─────────────────────────────────────────────────────────

type Kiosco struct {
	catalogo    map[string]*Producto
	notificador Notificador
	historial   []Evento
}

func nuevoKiosco(n Notificador) *Kiosco {
	return &Kiosco{
		catalogo:    map[string]*Producto{},
		notificador: n,
	}
}

func (k *Kiosco) agregarProducto(p *Producto) {
	k.catalogo[p.Nombre] = p
}

// comprar es la función polimórfica central: acepta CUALQUIER
// MetodoPago, y registra el resultado en el historial (que mezcla
// tipos de Evento distintos, resueltos con type switch al mostrarlo).
func (k *Kiosco) comprar(nombreProducto string, cantidad int, metodo MetodoPago) error {
	prod, existe := k.catalogo[nombreProducto]
	if !existe {
		return fmt.Errorf("producto %q no existe en el catálogo", nombreProducto)
	}

	if cantidad > prod.Stock {
		err := &ErrorStockInsuficiente{Producto: nombreProducto, Pedido: cantidad, Stock: prod.Stock}
		k.historial = append(k.historial, EventoError{Motivo: err.Error()})
		return err
	}

	total := prod.Precio * float64(cantidad)

	if err := metodo.Pagar(total); err != nil {
		k.historial = append(k.historial, EventoError{Motivo: err.Error()})
		return err
	}

	prod.Stock -= cantidad
	k.historial = append(k.historial, EventoVenta{Producto: nombreProducto, Total: total})
	k.notificador.Notificar(fmt.Sprintf("Compraste %d x %s con %s — total $%.2f",
		cantidad, nombreProducto, metodo, total))

	return nil
}

func (k *Kiosco) mostrarHistorial() {
	fmt.Println("\n=== Historial (type switch sobre Evento) ===")
	ventas, errores := 0, 0
	for _, ev := range k.historial {
		switch e := ev.(type) {
		case EventoVenta:
			ventas++
			fmt.Println("  ✓", e.Resumen())
		case EventoError:
			errores++
			fmt.Println("  ✗", e.Resumen())
		}
	}
	fmt.Printf("  Total: %d ventas, %d errores\n", ventas, errores)
}

func main() {
	fmt.Println("=== KIOSCO DIGITAL: ejercicio integrador de Interfaces ===")

	kiosco := nuevoKiosco(NotificadorConsola{})
	kiosco.agregarProducto(&Producto{Nombre: "Alfajor", Precio: 800, Stock: 10})
	kiosco.agregarProducto(&Producto{Nombre: "Gaseosa", Precio: 1200, Stock: 3})

	fmt.Println("\n=== Catálogo (usa Stringer) ===")
	for _, p := range kiosco.catalogo {
		fmt.Println(" -", p)
	}

	mp := &MercadoPago{Saldo: 5000}
	efectivo := Efectivo{}

	fmt.Println("\n=== Compras ===")

	if err := kiosco.comprar("Alfajor", 2, efectivo); err != nil {
		fmt.Println("  ERROR:", err)
	}

	if err := kiosco.comprar("Gaseosa", 5, mp); err != nil {
		// Type assertion para reaccionar distinto según el error
		var errStock *ErrorStockInsuficiente
		if errors.As(err, &errStock) {
			fmt.Printf("  ERROR (stock): faltan %d unidades\n", errStock.Pedido-errStock.Stock)
		} else {
			fmt.Println("  ERROR:", err)
		}
	}

	if err := kiosco.comprar("Alfajor", 1, mp); err != nil {
		fmt.Println("  ERROR:", err)
	}

	kiosco.mostrarHistorial()

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE TODO LO VISTO EN 06_interfaces_y_errores/
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué interfaz resolvió cada problema ===")
	fmt.Println("  MetodoPago    → polimorfismo: Efectivo y MercadoPago, un solo comprar()")
	fmt.Println("  Notificador   → desacopla el aviso del canal (consola, SMS, email...)")
	fmt.Println("  error         → ErrorStockInsuficiente es un error como cualquier otro")
	fmt.Println("  Evento        → type switch para mostrar historial mixto")
	fmt.Println("  Stringer      → Producto y MetodoPago se imprimen lindo solos")
}
