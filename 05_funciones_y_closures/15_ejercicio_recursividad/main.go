package main

import (
	"fmt"
	"math"
)

// =========================================================
// APLICANDO RECURSIVIDAD AL PROYECTO + DESCUENTO RECURSIVO
// + MEJORANDO FORMATO DE MONEDA
// =========================================================
// Aplicamos recursividad a nuestra tienda:
//   1. Calcular total de carritos anidados recursivamente
//   2. Aplicar descuentos por volumen de forma recursiva
//   3. Formatear precios en moneda con separadores de miles

// ─────────────────────────────────────────────────────────
// TIPOS
// ─────────────────────────────────────────────────────────
type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

type Item struct {
	Producto Producto
	Cantidad int
}

type Carrito struct {
	Cliente string
	Items   []Item
}

// ─────────────────────────────────────────────────────────
// FORMATEO DE MONEDA (mejorando la presentación)
// ─────────────────────────────────────────────────────────

// formatearEnteroConComas convierte 1500000 → "1,500,000"
// Usa recursión para insertar la coma cada 3 dígitos
func formatearEnteroConComas(n int) string {
	if n < 0 {
		return "-" + formatearEnteroConComas(-n)
	}
	if n < 1000 { // CASO BASE: menos de 1000, no necesita coma
		return fmt.Sprintf("%d", n)
	}
	// CASO RECURSIVO: formatear el resto + "," + últimos 3 dígitos
	return formatearEnteroConComas(n/1000) + fmt.Sprintf(",%03d", n%1000)
}

// formatearPrecio retorna precio formateado para tabla
func formatearPrecio(precio float64) string {
	entero := int(precio)
	centavos := int(math.Round((precio - float64(entero)) * 100))
	return fmt.Sprintf("$%s.%02d", formatearEnteroConComas(entero), centavos)
}

// ─────────────────────────────────────────────────────────
// DESCUENTO RECURSIVO POR VOLUMEN
// ─────────────────────────────────────────────────────────
// Reglas de descuento:
//   - $0      - $999:    0% descuento
//   - $1,000  - $4,999:  5% descuento
//   - $5,000  - $9,999: 10% descuento
//   - $10,000+:         15% descuento
//
// El descuento aplica recursivamente:
//   Si tenés 3 items y el total > $10,000, aplicamos 15%.
//   Pero también podemos aplicar descuentos adicionales por cada $5,000 extra.

func calcularDescuentoRecursivo(total float64) float64 {
	switch {
	case total < 1000:
		return 0 // CASO BASE: sin descuento
	case total < 5000:
		return 0.05 // CASO BASE: 5%
	case total < 10000:
		return 0.10 // CASO BASE: 10%
	default:
		// CASO RECURSIVO: 15% base + descuento adicional por cada $10,000 extra
		// Por cada $10,000 sobre los primeros $10,000 → +1% adicional (máximo 25%)
		descuentoBase := 0.15
		extra := total - 10000
		descuentoAdicional := calcularDescuentoExtra(extra, 0.01, 0.10)
		descuento := descuentoBase + descuentoAdicional
		if descuento > 0.25 {
			return 0.25 // tope máximo
		}
		return descuento
	}
}

// calcularDescuentoExtra calcula descuento adicional de forma recursiva
func calcularDescuentoExtra(montoExtra, incremento, tope float64) float64 {
	if montoExtra <= 0 || incremento >= tope { // CASO BASE
		return 0
	}
	// Por cada $5,000 extra → +incremento
	if montoExtra < 5000 {
		return 0
	}
	return incremento + calcularDescuentoExtra(montoExtra-5000, incremento, tope)
}

// ─────────────────────────────────────────────────────────
// CARRITO CON RECURSIÓN
// ─────────────────────────────────────────────────────────

// calcularTotalRecursivo calcula el total del carrito de forma recursiva
func calcularTotalRecursivo(items []Item) float64 {
	if len(items) == 0 { // CASO BASE: carrito vacío
		return 0
	}
	// CASO RECURSIVO: precio del primer item + total del resto
	primero := items[0]
	subtotalPrimero := primero.Producto.Precio * float64(primero.Cantidad)
	return subtotalPrimero + calcularTotalRecursivo(items[1:])
}

func mostrarCarrito(c Carrito) {
	fmt.Printf("\n  ┌─ Carrito de %s ─────────────────\n", c.Cliente)
	fmt.Printf("  │  %-15s %10s %6s %12s\n", "Producto", "Precio", "Cant.", "Subtotal")
	fmt.Printf("  │  %s\n", repetir("─", 49))

	for _, item := range c.Items {
		subtotal := item.Producto.Precio * float64(item.Cantidad)
		fmt.Printf("  │  %-15s %10s %6d %12s\n",
			item.Producto.Nombre,
			formatearPrecio(item.Producto.Precio),
			item.Cantidad,
			formatearPrecio(subtotal))
	}

	subtotal := calcularTotalRecursivo(c.Items)
	descuento := calcularDescuentoRecursivo(subtotal)
	montoDescuento := subtotal * descuento
	total := subtotal - montoDescuento

	fmt.Printf("  │  %s\n", repetir("─", 49))
	fmt.Printf("  │  %-33s %12s\n", "Subtotal:", formatearPrecio(subtotal))

	if descuento > 0 {
		fmt.Printf("  │  %-33s %12s\n",
			fmt.Sprintf("Descuento (%.0f%%):", descuento*100),
			"-"+formatearPrecio(montoDescuento))
	}

	fmt.Printf("  │  %-33s %12s\n", "TOTAL:", formatearPrecio(total))
	fmt.Printf("  └%s\n", repetir("─", 50))
}

func repetir(s string, n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r += s
	}
	return r
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  RECURSIÓN + FORMATO MONEDA   ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// FORMATO DE MONEDA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== formatearEnteroConComas (recursivo) ===")
	montos := []float64{0, 99.5, 1000, 12345.99, 1000000, 9999999.99}
	for _, m := range montos {
		fmt.Printf("  %12.2f → %s\n", m, formatearPrecio(m))
	}

	// ─────────────────────────────────────────────────────────
	// DESCUENTO RECURSIVO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Descuentos por volumen (recursivo) ===")
	totalesPrueba := []float64{500, 2000, 7500, 10000, 15000, 25000, 50000}
	for _, t := range totalesPrueba {
		desc := calcularDescuentoRecursivo(t)
		ahorro := t * desc
		fmt.Printf("  %12s → descuento: %3.0f%% → ahorrás: %s\n",
			formatearPrecio(t), desc*100, formatearPrecio(ahorro))
	}

	// ─────────────────────────────────────────────────────────
	// CARRITOS DE PRUEBA
	// ─────────────────────────────────────────────────────────
	productos := []Producto{
		{1, "Notebook", 1500.00, 10},
		{2, "Mouse", 25.99, 50},
		{3, "Teclado", 75.50, 20},
		{4, "Monitor", 450.00, 8},
		{5, "Auriculares", 80.00, 15},
	}

	// Carrito chico (sin descuento)
	carritoChico := Carrito{
		Cliente: "Ana García",
		Items: []Item{
			{productos[1], 1},
			{productos[2], 1},
		},
	}
	mostrarCarrito(carritoChico)

	// Carrito mediano (descuento 5%)
	carritoMediano := Carrito{
		Cliente: "Carlos López",
		Items: []Item{
			{productos[0], 1},
			{productos[3], 1},
		},
	}
	mostrarCarrito(carritoMediano)

	// Carrito grande (descuento 15%+)
	carritoGrande := Carrito{
		Cliente: "Empresa Tech S.A.",
		Items: []Item{
			{productos[0], 5},
			{productos[3], 3},
			{productos[2], 10},
			{productos[1], 20},
		},
	}
	mostrarCarrito(carritoGrande)

	// ─────────────────────────────────────────────────────────
	// VISUALIZAR RECURSIÓN DE TOTAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== calcularTotalRecursivo en acción ===")
	items := carritoChico.Items
	fmt.Printf("Items: %v\n", func() []string {
		names := []string{}
		for _, i := range items {
			names = append(names, fmt.Sprintf("%s×%d", i.Producto.Nombre, i.Cantidad))
		}
		return names
	}())
	fmt.Printf("Total recursivo: %s\n", formatearPrecio(calcularTotalRecursivo(items)))

	// ─────────────────────────────────────────────────────────
	// RESUMEN FINAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen del módulo de Funciones ===")
	fmt.Println("✓ Declarar y llamar funciones")
	fmt.Println("✓ Parámetros con structs (valor vs puntero)")
	fmt.Println("✓ Parámetros variádicos (...tipo)")
	fmt.Println("✓ Retorno múltiple y (resultado, error)")
	fmt.Println("✓ Retorno nombrado y valor en blanco (_)")
	fmt.Println("✓ defer para garantizar limpieza de recursos")
	fmt.Println("✓ Funciones como valores, anónimas, closures")
	fmt.Println("✓ Funciones como parámetros (HOF)")
	fmt.Println("✓ Métodos con receiver por valor y puntero")
	fmt.Println("✓ Recursividad con caso base y caso recursivo")
	fmt.Println("✓ Formateo de moneda con separadores de miles")
}
