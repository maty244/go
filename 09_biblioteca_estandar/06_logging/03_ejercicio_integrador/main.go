package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

// =========================================================
// EJERCICIO INTEGRADOR: LOGGING DE UN CHECKOUT COMPLETO
// =========================================================
// Combinamos slog con lo visto en 06_interfaces_y_errores/: cada paso
// del checkout loguea lo que corresponde (Info para lo normal,
// Warn para lo raro-pero-no-grave, Error para lo que falla), y
// usamos logger.With() para que TODOS los logs de un pedido
// lleven su ID automáticamente, sin repetirlo a mano en cada línea.

var ErrStockInsuficiente = errors.New("stock insuficiente")
var ErrPagoRechazado = errors.New("pago rechazado")

type Pedido struct {
	ID       int
	Cliente  string
	Producto string
	Cantidad int
	Stock    int
}

func procesarCheckout(logger *slog.Logger, p Pedido) error {
	// logger.With: a partir de acá, TODOS los logs de esta función
	// incluyen pedido_id y cliente automáticamente.
	log := logger.With("pedido_id", p.ID, "cliente", p.Cliente)

	log.Info("iniciando checkout", "producto", p.Producto, "cantidad", p.Cantidad)

	if p.Cantidad > p.Stock {
		log.Error("no se pudo completar el checkout",
			"motivo", ErrStockInsuficiente,
			"pedido_cantidad", p.Cantidad,
			"stock_disponible", p.Stock,
		)
		return fmt.Errorf("procesarCheckout: %w", ErrStockInsuficiente)
	}

	if p.Stock-p.Cantidad <= 2 {
		log.Warn("stock quedará muy bajo tras esta venta",
			"producto", p.Producto,
			"stock_restante", p.Stock-p.Cantidad,
		)
	}

	// Simulamos que los pedidos de más de 10 unidades fallan el pago
	if p.Cantidad > 10 {
		log.Error("pago rechazado", "motivo", ErrPagoRechazado)
		return fmt.Errorf("procesarCheckout: %w", ErrPagoRechazado)
	}

	log.Info("checkout completado con éxito")
	return nil
}

func main() {
	// Usamos JSON (tema 02): es el formato que un sistema de
	// monitoreo real leería en producción.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	pedidos := []Pedido{
		{ID: 1001, Cliente: "Matias", Producto: "Alfajor", Cantidad: 3, Stock: 20},
		{ID: 1002, Cliente: "Ana", Producto: "Gaseosa", Cantidad: 5, Stock: 6},       // stock quedará bajo
		{ID: 1003, Cliente: "Carlos", Producto: "Notebook", Cantidad: 15, Stock: 20}, // pago rechazado
		{ID: 1004, Cliente: "Sofía", Producto: "Mouse", Cantidad: 30, Stock: 10},     // sin stock
	}

	exitosos, fallidos := 0, 0

	for _, p := range pedidos {
		if err := procesarCheckout(logger, p); err != nil {
			fallidos++
			continue
		}
		exitosos++
	}

	// El resumen final lo mostramos con fmt (no todo tiene que ser
	// slog: los logs son para EVENTOS del sistema, un resumen para
	// el usuario de la terminal puede seguir siendo fmt.Println).
	fmt.Println("\n=== Resumen de la corrida ===")
	fmt.Printf("Exitosos: %d | Fallidos: %d\n", exitosos, fallidos)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 09_biblioteca_estandar/06_logging/ ===")
	fmt.Println("  slog.NewJSONHandler   → logs listos para un sistema de monitoreo")
	fmt.Println("  logger.With(...)      → pedido_id y cliente en TODOS los logs de la función")
	fmt.Println("  Info/Warn/Error       → según qué tan grave es cada situación")
	fmt.Println("  errores centinela     → conectan el log con el error real devuelto")
}
