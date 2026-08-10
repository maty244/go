package main

import "fmt"

// =========================================================
// EJERCICIO INTEGRADOR: KIOSCO DIGITAL CON REGLAS CONFIGURABLES
// =========================================================
// Combinamos closures, contadores con estado, decoradores y
// fábricas de funciones en un caso real: un sistema de reglas de
// descuento configurables para el checkout, más un contador de
// ventas que "recuerda" el total del día sin usar una variable
// global.

// ─────────────────────────────────────────────────────────
// REGLA DE DESCUENTO: una función que decide cuánto descontar
// ─────────────────────────────────────────────────────────
// Cada regla es un closure que ya trae su configuración adentro.

type ReglaDescuento func(total float64, cantidadItems int) float64

// descuentoPorMonto: fábrica que arma una regla "si superás $X, descuento Y%"
func descuentoPorMonto(umbral, porcentaje float64) ReglaDescuento {
	return func(total float64, _ int) float64 {
		if total >= umbral {
			return total * porcentaje
		}
		return 0
	}
}

// descuentoPorCantidad: fábrica "si comprás N items o más, descuento Y%"
func descuentoPorCantidad(minItems int, porcentaje float64) ReglaDescuento {
	return func(total float64, cantidad int) float64 {
		if cantidad >= minItems {
			return total * porcentaje
		}
		return 0
	}
}

// ─────────────────────────────────────────────────────────
// CAJA REGISTRADORA: closures con estado compartido
// ─────────────────────────────────────────────────────────
// Sigue el patrón del tema 03: varias funciones que comparten
// las mismas variables capturadas, simulando un "objeto" sin
// necesidad de un struct con métodos.

type CajaRegistradora struct {
	Cobrar       func(monto float64)
	TotalDelDia  func() float64
	VentasDelDia func() int
}

func nuevaCaja() CajaRegistradora {
	total := 0.0
	ventas := 0

	return CajaRegistradora{
		Cobrar: func(monto float64) {
			total += monto
			ventas++
		},
		TotalDelDia:  func() float64 { return total },
		VentasDelDia: func() int { return ventas },
	}
}

// ─────────────────────────────────────────────────────────
// DECORADOR: loguear cada cobro sin tocar la caja
// ─────────────────────────────────────────────────────────

func conLogDeCobro(cobrar func(float64)) func(float64) {
	return func(monto float64) {
		fmt.Printf("  [LOG] Cobrando $%.2f...\n", monto)
		cobrar(monto)
	}
}

// procesarVenta aplica todas las reglas de descuento configuradas
// y devuelve el monto final a cobrar.
func procesarVenta(reglas []ReglaDescuento, total float64, cantidadItems int) float64 {
	mejorDescuento := 0.0
	for _, regla := range reglas {
		d := regla(total, cantidadItems)
		if d > mejorDescuento {
			mejorDescuento = d // nos quedamos con el MEJOR descuento aplicable
		}
	}
	return total - mejorDescuento
}

func main() {
	fmt.Println("=== KIOSCO DIGITAL: reglas de descuento + caja con closures ===")

	// Reglas configuradas UNA vez, reusadas para todas las ventas
	reglas := []ReglaDescuento{
		descuentoPorMonto(10000, 0.10), // 10% si superás $10.000
		descuentoPorCantidad(5, 0.05),  // 5% si comprás 5 o más items
	}

	caja := nuevaCaja()
	cobrarConLog := conLogDeCobro(caja.Cobrar)

	ventas := []struct {
		Descripcion string
		Total       float64
		Items       int
	}{
		{"Compra chica", 3000, 2},
		{"Compra grande", 15000, 3},
		{"Compra con muchos items", 8000, 6},
	}

	for _, v := range ventas {
		fmt.Printf("\n--- %s ---\n", v.Descripcion)
		final := procesarVenta(reglas, v.Total, v.Items)
		fmt.Printf("  Total sin descuento: $%.2f → a cobrar: $%.2f\n", v.Total, final)
		cobrarConLog(final)
	}

	fmt.Println("\n=== Resumen del día (estado guardado en closures) ===")
	fmt.Printf("  Ventas realizadas: %d\n", caja.VentasDelDia())
	fmt.Printf("  Total recaudado:   $%.2f\n", caja.TotalDelDia())

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué resolvió cada closure ===")
	fmt.Println("  descuentoPorMonto/Cantidad → fábricas: reglas ya configuradas")
	fmt.Println("  nuevaCaja                  → estado compartido (total, ventas)")
	fmt.Println("  conLogDeCobro               → decorador: agrega logging sin tocar Cobrar")
}
