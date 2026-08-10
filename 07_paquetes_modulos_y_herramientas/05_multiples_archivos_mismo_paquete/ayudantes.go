package main

import "fmt"

// Este archivo declara "package main", igual que main.go.
// Por eso todo lo de acá está disponible en main.go sin import.

// formatearPrecio da formato de moneda a un precio.
func formatearPrecio(precio float64) string {
	return fmt.Sprintf("$%.2f", precio)
}

// validarStock valida que haya suficiente stock para una compra.
func validarStock(pedido, stock int) error {
	if pedido > stock {
		return fmt.Errorf("stock insuficiente: pediste %d, hay %d", pedido, stock)
	}
	return nil
}
