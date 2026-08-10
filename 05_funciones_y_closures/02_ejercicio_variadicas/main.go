package main

import (
	"fmt"
	"strings"
)

// =========================================================
// APLICANDO CUPONES DE DESCUENTO (VARIADIC)
// =========================================================
// Caso de uso real: un sistema de cupones donde podés
// aplicar 0, 1, o varios descuentos a una compra.
// El variádico es perfecto porque no siempre hay cupones.

type Producto struct {
	Nombre string
	Precio float64
}

type Cupon struct {
	Codigo    string
	Descuento float64 // porcentaje: 0.10 = 10%
	MinCompra float64 // compra mínima para aplicar
}

// ─────────────────────────────────────────────────────────
// CATÁLOGO DE CUPONES DISPONIBLES
// ─────────────────────────────────────────────────────────
var cuponesCatalogo = map[string]Cupon{
	"VERANO10": {Codigo: "VERANO10", Descuento: 0.10, MinCompra: 0},
	"TECH20":   {Codigo: "TECH20", Descuento: 0.20, MinCompra: 500},
	"PROMO5":   {Codigo: "PROMO5", Descuento: 0.05, MinCompra: 0},
	"VIP30":    {Codigo: "VIP30", Descuento: 0.30, MinCompra: 1000},
}

// ─────────────────────────────────────────────────────────
// calcularTotal: función variádica que aplica cupones
// ─────────────────────────────────────────────────────────
// El tercer parámetro es variádico: podés pasar 0 o más cupones.
func calcularTotal(productos []Producto, impuesto float64, codigos ...string) (float64, []string) {
	// 1. Calcular subtotal
	subtotal := 0.0
	for _, p := range productos {
		subtotal += p.Precio
	}

	descuentoTotal := 0.0
	var cuponesAplicados []string
	var cuponesRechazados []string

	// 2. Aplicar cada cupón ingresado
	for _, codigo := range codigos {
		codigoUpper := strings.ToUpper(codigo)
		cupon, existe := cuponesCatalogo[codigoUpper]

		if !existe {
			cuponesRechazados = append(cuponesRechazados, fmt.Sprintf("%s (inválido)", codigo))
			continue
		}
		if subtotal < cupon.MinCompra {
			cuponesRechazados = append(cuponesRechazados,
				fmt.Sprintf("%s (mínimo $%.0f)", codigo, cupon.MinCompra))
			continue
		}
		descuentoTotal += cupon.Descuento
		cuponesAplicados = append(cuponesAplicados,
			fmt.Sprintf("%s (-%.0f%%)", cupon.Codigo, cupon.Descuento*100))
	}

	// 3. Aplicar descuento total (no puede pasar del 50%)
	if descuentoTotal > 0.50 {
		descuentoTotal = 0.50
	}

	subtotalConDescuento := subtotal * (1 - descuentoTotal)
	totalConImpuesto := subtotalConDescuento * (1 + impuesto)

	// 4. Mostrar desglose
	fmt.Printf("\n  Subtotal:             $%8.2f\n", subtotal)
	if descuentoTotal > 0 {
		ahorro := subtotal * descuentoTotal
		fmt.Printf("  Descuento (%.0f%%):     -$%8.2f\n", descuentoTotal*100, ahorro)
		fmt.Printf("  Subtotal c/descuento: $%8.2f\n", subtotalConDescuento)
	}
	if len(cuponesRechazados) > 0 {
		fmt.Printf("  Cupones rechazados:   %s\n", strings.Join(cuponesRechazados, ", "))
	}
	fmt.Printf("  IVA (%.0f%%):           +$%8.2f\n",
		impuesto*100, subtotalConDescuento*impuesto)
	fmt.Printf("  ─────────────────────────────\n")
	fmt.Printf("  TOTAL:                $%8.2f\n", totalConImpuesto)

	return totalConImpuesto, cuponesAplicados
}

func mostrarCarrito(productos []Producto) {
	fmt.Println("  Carrito:")
	for _, p := range productos {
		fmt.Printf("    %-15s $%.2f\n", p.Nombre, p.Precio)
	}
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  SISTEMA DE CUPONES VARIÁDICO ║")
	fmt.Println("╚══════════════════════════════╝")

	const IVA = 0.21

	notebook := Producto{"Notebook", 1500.00}
	mouse := Producto{"Mouse", 25.99}
	teclado := Producto{"Teclado", 75.50}

	// ─────────────────────────────────────────────────────────
	// CASO 1: Sin cupones (carrito simple)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Compra sin cupones ---")
	carrito1 := []Producto{mouse, teclado}
	mostrarCarrito(carrito1)
	total1, _ := calcularTotal(carrito1, IVA) // sin cupones → variádico vacío
	fmt.Printf("  Pagás: $%.2f\n", total1)

	// ─────────────────────────────────────────────────────────
	// CASO 2: Un cupón válido
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Compra con un cupón ---")
	carrito2 := []Producto{mouse, teclado}
	mostrarCarrito(carrito2)
	total2, aplicados := calcularTotal(carrito2, IVA, "VERANO10")
	fmt.Printf("  Cupones aplicados: %v\n", aplicados)
	fmt.Printf("  Pagás: $%.2f\n", total2)

	// ─────────────────────────────────────────────────────────
	// CASO 3: Múltiples cupones
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Compra con múltiples cupones ---")
	carrito3 := []Producto{notebook, mouse, teclado}
	mostrarCarrito(carrito3)
	// TECH20 requiere mínimo $500 → aplica (total > 500)
	// VIP30 requiere mínimo $1000 → aplica
	total3, aplicados3 := calcularTotal(carrito3, IVA, "TECH20", "VIP30", "VERANO10")
	fmt.Printf("  Cupones aplicados: %v\n", aplicados3)
	fmt.Printf("  Pagás: $%.2f\n", total3)

	// ─────────────────────────────────────────────────────────
	// CASO 4: Cupón que no cumple el mínimo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Cupón que no cumple mínimo ---")
	carrito4 := []Producto{mouse} // solo $25.99
	mostrarCarrito(carrito4)
	// TECH20 requiere $500 → se rechaza
	total4, aplicados4 := calcularTotal(carrito4, IVA, "TECH20", "PROMO5")
	fmt.Printf("  Cupones aplicados: %v\n", aplicados4)
	fmt.Printf("  Pagás: $%.2f\n", total4)

	// ─────────────────────────────────────────────────────────
	// CASO 5: Cupón inválido
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Cupón inválido ---")
	carrito5 := []Producto{teclado}
	mostrarCarrito(carrito5)
	total5, _ := calcularTotal(carrito5, IVA, "DESCUENTO99", "VERANO10")
	fmt.Printf("  Pagás: $%.2f\n", total5)

	// ─────────────────────────────────────────────────────────
	// EXPANDIR SLICE DE CUPONES CON ...
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Aplicar cupones desde un slice ---")
	cuponesDelUsuario := []string{"PROMO5", "TECH20"}
	carrito6 := []Producto{notebook}
	mostrarCarrito(carrito6)
	total6, _ := calcularTotal(carrito6, IVA, cuponesDelUsuario...)
	fmt.Printf("  Pagás: $%.2f\n", total6)

	fmt.Println("\n=== Cupones disponibles ===")
	for codigo, c := range cuponesCatalogo {
		if c.MinCompra > 0 {
			fmt.Printf("  %s: %.0f%% off (mínimo $%.0f)\n",
				codigo, c.Descuento*100, c.MinCompra)
		} else {
			fmt.Printf("  %s: %.0f%% off (sin mínimo)\n",
				codigo, c.Descuento*100)
		}
	}
}
