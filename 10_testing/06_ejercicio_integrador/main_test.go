package main

import (
	"errors"
	"testing"
)

// =========================================================
// SUITE COMPLETA: tabla + subtests + errors.Is
// =========================================================
// Una tabla que cubre las 4 ramas de ValidarPedido (caso válido +
// 3 casos de error), cada una como un subtest con nombre propio,
// verificando el error EXACTO con errors.Is.

func TestValidarPedido(t *testing.T) {
	casos := []struct {
		nombre      string
		pedido      Pedido
		errEsperado error // nil significa "no debería fallar"
	}{
		{
			nombre:      "pedido válido",
			pedido:      Pedido{Producto: "Notebook", Cantidad: 1, Precio: 450000},
			errEsperado: nil,
		},
		{
			nombre:      "producto vacío",
			pedido:      Pedido{Producto: "", Cantidad: 1, Precio: 100},
			errEsperado: ErrProductoVacio,
		},
		{
			nombre:      "cantidad cero",
			pedido:      Pedido{Producto: "Mouse", Cantidad: 0, Precio: 8500},
			errEsperado: ErrCantidadInvalida,
		},
		{
			nombre:      "cantidad negativa",
			pedido:      Pedido{Producto: "Mouse", Cantidad: -3, Precio: 8500},
			errEsperado: ErrCantidadInvalida,
		},
		{
			nombre:      "precio negativo",
			pedido:      Pedido{Producto: "Teclado", Cantidad: 1, Precio: -500},
			errEsperado: ErrPrecioInvalido,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := ValidarPedido(c.pedido)

			if c.errEsperado == nil {
				if err != nil {
					t.Errorf("no se esperaba error, se obtuvo: %v", err)
				}
				return
			}

			if !errors.Is(err, c.errEsperado) {
				t.Errorf("se esperaba %v, se obtuvo: %v", c.errEsperado, err)
			}
		})
	}
}

// ─────────────────────────────────────────────────────────
// UN TEST ADICIONAL: verificar que el orden de validación es
// el esperado (producto vacío se detecta ANTES que cantidad)
// ─────────────────────────────────────────────────────────

func TestValidarPedidoOrdenDeValidacion(t *testing.T) {
	// Este pedido tiene DOS problemas: producto vacío Y cantidad
	// inválida. Confirmamos cuál de los dos se reporta primero,
	// documentando el comportamiento esperado de la función.
	pedido := Pedido{Producto: "", Cantidad: -1, Precio: 100}

	err := ValidarPedido(pedido)

	if !errors.Is(err, ErrProductoVacio) {
		t.Errorf("se esperaba que ErrProductoVacio se detectara primero, se obtuvo: %v", err)
	}
}
