package main

import (
	"errors"
	"fmt"
)

// Dinero representa centavos, la unidad mínima elegida por esta aplicación.
// El tipo definido impide mezclar por accidente un importe con cualquier int64.
type Dinero int64

var ErrCantidadInvalida = errors.New("cantidad inválida")

func Total(precioUnitario Dinero, cantidad int) (Dinero, error) {
	if precioUnitario < 0 || cantidad < 0 {
		return 0, ErrCantidadInvalida
	}
	return precioUnitario * Dinero(cantidad), nil
}

func formatear(importe Dinero) string {
	centavos := importe % 100
	pesos := importe / 100
	return fmt.Sprintf("$%d.%02d", pesos, centavos)
}

func main() {
	total, err := Total(1_299, 3) // $12.99 por tres unidades
	if err != nil {
		fmt.Println("calcular total:", err)
		return
	}
	fmt.Println(formatear(total))

	fmt.Println("\nRegla del curso:")
	fmt.Println("- dinero: entero + unidad explícita (centavos)")
	fmt.Println("- geometría/mediciones: float64 cuando corresponde")
	fmt.Println("- porcentajes: aplicarlos con una política de redondeo explícita")
}
