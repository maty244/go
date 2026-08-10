package main

import (
	"errors"
	"fmt"
)

// =========================================================
// EJERCICIO INTEGRADOR: SUITE DE TESTS COMPLETA
// =========================================================
// Combinamos todo lo visto en 10_testing/: tabla de casos, subtests
// con t.Run, testing de errores con errors.Is, y coverage. Le
// escribimos una suite de tests completa a una función de
// validación de pedidos del kiosco.

var (
	ErrCantidadInvalida = errors.New("la cantidad debe ser mayor a 0")
	ErrPrecioInvalido   = errors.New("el precio no puede ser negativo")
	ErrProductoVacio    = errors.New("el nombre del producto no puede estar vacío")
)

type Pedido struct {
	Producto string
	Cantidad int
	Precio   float64
}

// ValidarPedido revisa las reglas de negocio de un pedido antes
// de procesarlo. Es la función que vamos a testear a fondo en
// main_test.go.
func ValidarPedido(p Pedido) error {
	if p.Producto == "" {
		return ErrProductoVacio
	}
	if p.Cantidad <= 0 {
		return fmt.Errorf("ValidarPedido(%s): %w", p.Producto, ErrCantidadInvalida)
	}
	if p.Precio < 0 {
		return fmt.Errorf("ValidarPedido(%s): %w", p.Producto, ErrPrecioInvalido)
	}
	return nil
}

func main() {
	fmt.Println("=== ValidarPedido ===")

	pedidos := []Pedido{
		{Producto: "Notebook", Cantidad: 1, Precio: 450000},
		{Producto: "", Cantidad: 1, Precio: 100},
		{Producto: "Mouse", Cantidad: 0, Precio: 8500},
		{Producto: "Teclado", Cantidad: 1, Precio: -500},
	}

	for _, p := range pedidos {
		err := ValidarPedido(p)
		if err != nil {
			fmt.Printf("  %+v → ERROR: %v\n", p, err)
		} else {
			fmt.Printf("  %+v → OK\n", p)
		}
	}

	fmt.Println("\n(mirá main_test.go: la suite completa de tests, con todo lo de 10_testing/)")
	fmt.Println("Corré:")
	fmt.Println("  go test -v ./10_testing/06_ejercicio_integrador/")
	fmt.Println("  go test -cover ./10_testing/06_ejercicio_integrador/")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 10_testing/ ===")
	fmt.Println("  Tabla de casos (02)     → un caso por regla de negocio a validar")
	fmt.Println("  t.Run subtests (03)     → cada caso identificado por su nombre")
	fmt.Println("  errors.Is (05)          → verificar EXACTAMENTE qué error se esperaba")
	fmt.Println("  go test -cover (04)     → confirmar que las 4 ramas están cubiertas")
}
