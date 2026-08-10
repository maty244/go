package main

import (
	"errors"
	"fmt"
)

// =========================================================
// errors.Is A FONDO
// =========================================================
// errors.Is(err, objetivo) responde: "¿es err (o alguno de los
// errores que envuelve) el mismo que objetivo?"
//
// Para errores centinela SIN wrapping, esto es equivalente a
// "err == objetivo". La diferencia aparece con wrapping (tema 02):
// errors.Is SIGUE buscando adentro de las capas, mientras que ==
// solo compara el valor de más afuera.
//
// Es el patrón correcto para preguntar "¿pasó ESTO específico?"
// cuando tenés varios errores centinela posibles.

var (
	ErrSinStock         = errors.New("sin stock")
	ErrProductoInvalido = errors.New("producto inválido")
	ErrPagoRechazado    = errors.New("pago rechazado")
)

func comprar(producto string, cantidad int, saldoSuficiente bool) error {
	if producto == "" {
		return ErrProductoInvalido
	}
	if cantidad > 10 {
		return fmt.Errorf("comprar(%s): %w", producto, ErrSinStock)
	}
	if !saldoSuficiente {
		return fmt.Errorf("comprar(%s): %w", producto, ErrPagoRechazado)
	}
	return nil
}

func main() {
	fmt.Println("=== errors.Is con varios errores centinela ===")

	casos := []struct {
		producto string
		cantidad int
		saldo    bool
	}{
		{"Mouse", 15, true},
		{"", 1, true},
		{"Teclado", 2, false},
		{"Monitor", 1, true},
	}

	for _, c := range casos {
		err := comprar(c.producto, c.cantidad, c.saldo)
		manejarError(err)
	}

	// ─────────────────────────────────────────────────────────
	// == vs errors.Is: LA DIFERENCIA REAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== == vs errors.Is ===")

	errEnvuelto := fmt.Errorf("contexto: %w", ErrSinStock)

	fmt.Println("errEnvuelto == ErrSinStock:      ", errEnvuelto == ErrSinStock)
	fmt.Println("errors.Is(errEnvuelto, ErrSinStock):", errors.Is(errEnvuelto, ErrSinStock))

	// == da false porque errEnvuelto es un valor DISTINTO (el
	// wrapper), aunque adentro tenga a ErrSinStock. errors.Is sabe
	// "desenvolver" para encontrarlo.

	// ─────────────────────────────────────────────────────────
	// errors.Is TAMBIÉN FUNCIONA SIN WRAPPING
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Sin wrapping, errors.Is se comporta como == ===")
	fmt.Println("errors.Is(ErrSinStock, ErrSinStock):", errors.Is(ErrSinStock, ErrSinStock))
	fmt.Println("errors.Is(ErrSinStock, ErrPagoRechazado):", errors.Is(ErrSinStock, ErrPagoRechazado))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  errors.Is(err, objetivo)  → ¿err ES (o envuelve a) objetivo?")
	fmt.Println("  Atraviesa el wrapping     → a diferencia de ==")
	fmt.Println("  Usalo para                → decidir QUÉ pasó y reaccionar distinto")
	fmt.Println("  Con varios centinela      → si/else if encadenados con errors.Is")
}

func manejarError(err error) {
	switch {
	case err == nil:
		fmt.Println("  Compra exitosa")
	case errors.Is(err, ErrSinStock):
		fmt.Println("  → Avisar al depósito:", err)
	case errors.Is(err, ErrProductoInvalido):
		fmt.Println("  → Pedirle al usuario que revise el producto:", err)
	case errors.Is(err, ErrPagoRechazado):
		fmt.Println("  → Sugerir otro método de pago:", err)
	default:
		fmt.Println("  → Error desconocido:", err)
	}
}
