package main

import (
	"errors"
	"fmt"
)

// =========================================================
// errors.New VS fmt.Errorf
// =========================================================
// Ya usaste ambos en 06_interfaces_y_errores/12_patron_de_errores. Repasamos la diferencia
// exacta y cuándo conviene cada uno, ahora que sabés que error
// es una interfaz (06_interfaces_y_errores/09_error_interface_a_fondo).
//
//   errors.New("mensaje")        → crea un error con un texto FIJO
//   fmt.Errorf("mensaje %v", x)  → crea un error con texto FORMATEADO
//
// fmt.Errorf es, en esencia, "Sprintf + convertir el resultado en
// error". Si no necesitás insertar valores dinámicos, cualquiera
// de los dos funciona, pero fmt.Errorf es más flexible.

func main() {
	fmt.Println("=== errors.New: mensaje fijo ===")

	errFijo := errors.New("stock agotado")
	fmt.Println(errFijo)

	fmt.Println("\n=== fmt.Errorf: mensaje con datos dinámicos ===")

	producto := "Notebook"
	cantidad := 5
	errFormateado := fmt.Errorf("stock agotado de %s (pediste %d)", producto, cantidad)
	fmt.Println(errFormateado)

	// ─────────────────────────────────────────────────────────
	// ERRORES CENTINELA: variables error a nivel de paquete
	// ─────────────────────────────────────────────────────────
	// Un patrón MUY común: declarar errores conocidos como
	// variables globales del paquete, para poder compararlos
	// después con errors.Is (lo vemos a fondo en el tema 03).
	// Por convención se nombran empezando con "Err".

	fmt.Println("\n=== Errores centinela (variables reusables) ===")

	if err := comprar(0); err != nil {
		fmt.Println("Error:", err)
		fmt.Println("¿Es ErrCantidadInvalida?", errors.Is(err, ErrCantidadInvalida))
	}

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR CADA UNO
	// ─────────────────────────────────────────────────────────
	// errors.New   → mensaje fijo, reusable, se compara con errors.Is
	// fmt.Errorf   → mensaje con contexto específico de ESTA llamada
	//                (y además permite ENVOLVER otro error con %w,
	//                lo vemos en el próximo tema)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	pct := "%" // evita que go vet confunda esta línea con un Printf mal usado
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  errors.New("texto")       → error con mensaje fijo`)
	fmt.Println(`  fmt.Errorf("texto ` + pct + `v", x) → error con mensaje formateado`)
	fmt.Println(`  Errores centinela         → var ErrX = errors.New(...) a nivel paquete`)
	fmt.Println(`  Convención de nombres     → empezar con "Err" (ErrNoEncontrado, etc.)`)
}

// ErrCantidadInvalida es un error centinela: se compara por
// IDENTIDAD (la misma variable), no por el texto del mensaje.
var ErrCantidadInvalida = errors.New("la cantidad debe ser mayor a 0")

func comprar(cantidad int) error {
	if cantidad <= 0 {
		return ErrCantidadInvalida
	}
	return nil
}
