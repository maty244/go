package main

import (
	"errors"
	"fmt"
)

// =========================================================
// REPRESENTAR DINERO CON ENTEROS (CENTAVOS)
// =========================================================
// ¿Por qué NO usar float64 para dinero real?
// Porque float64 usa punto flotante binario (IEEE 754) y fracciones
// decimales simples como 0.1 o 0.2 no tienen representación exacta en base 2:
//   0.1 + 0.2 en float64 da 0.30000000000000004
// En un sistema de cobros o contabilidad, esos errores se acumulan.
//
// Solución básica: trabajar en la unidad mínima entera (por ejemplo, centavos)
// y crear un TIPO DEFINIDO para no mezclar centavos con enteros comunes.

// Dinero representa centavos (ej: $12.99 se guarda como 1299).
// Al ser un tipo definido (`type Dinero int64`), el compilador impide sumar
// o multiplicar por accidente un `Dinero` con un `int` o `int64` sin convertirlo.
type Dinero int64

// ErrCantidadInvalida es un error centinela (sentinel error) reutilizable.
var ErrCantidadInvalida = errors.New("cantidad o precio inválido")

// Total calcula el importe total multiplicando el precio unitario en centavos
// por la cantidad de unidades.
// Nota: como cantidad es `int` y precioUnitario es `Dinero`, Go exige convertir
// explícitamente `Dinero(cantidad)`.
// (Advertencia: esta versión básica no comprueba overflow de int64 en números gigantes).
func Total(precioUnitario Dinero, cantidad int) (Dinero, error) {
	if precioUnitario < 0 || cantidad < 0 {
		return 0, ErrCantidadInvalida
	}
	return precioUnitario * Dinero(cantidad), nil
}

// formatear convierte los centavos a una cadena legible "$pesos.centavos".
// Usa división entera (/ 100) para obtener la parte entera y módulo (% 100)
// para obtener los centavos restantes.
func formatear(importe Dinero) string {
	if importe < 0 {
		return "-" + formatear(-importe)
	}
	centavos := importe % 100 // resto al dividir por 100 (0 a 99)
	pesos := importe / 100    // división entera (descarta los centavos)
	return fmt.Sprintf("$%d.%02d", pesos, centavos)
}

func main() {
	// Ejemplo: sumar variables float64 en runtime (0.10 + 0.20) vs centavos (Dinero)
	a, b := 0.10, 0.20
	fmt.Printf("Problema con float64: 0.10 + 0.20 = %.17f (¿es == 0.30? %v)\n", a+b, a+b == 0.30)
	fmt.Printf("Exactitud con Dinero: 10 + 20 centavos = %s\n\n", formatear(Dinero(10)+Dinero(20)))

	// 1_299 centavos = $12.99 por unidad, comprando 3 unidades:
	total, err := Total(1_299, 3)
	if err != nil {
		fmt.Println("calcular total:", err)
		return
	}
	fmt.Println("Total por 3 unidades de $12.99:", formatear(total)) // $38.97

	fmt.Println("\nRegla del curso:")
	fmt.Println("- dinero: entero + unidad explícita (centavos)")
	fmt.Println("- geometría/mediciones: float64 cuando corresponde")
	fmt.Println("- porcentajes: aplicarlos con una política de redondeo explícita")
}
