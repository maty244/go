package main

import "fmt"

// =========================================================
// INTERFACES COMO PARÁMETROS: POLIMORFISMO
// =========================================================
// "Polimorfismo" es una palabra grande para una idea simple:
// una misma función se comporta distinto según QUÉ tipo concreto
// le pases, sin que la función tenga que saber cuál es.
//
// Este es EL caso de uso más importante de las interfaces, y el
// que resuelve el problema que mencionamos en el tema 01: en vez
// de escribir una función por cada tipo de pago, escribís UNA
// función que acepta la interfaz.

// ─────────────────────────────────────────────────────────
// EL CONTRATO: MetodoPago
// ─────────────────────────────────────────────────────────

type MetodoPago interface {
	Pagar(monto float64) error
	Comision() float64 // % que se descuenta por usar este método
}

// ─────────────────────────────────────────────────────────
// TARJETA DE CRÉDITO
// ─────────────────────────────────────────────────────────

type Tarjeta struct {
	Numero string
	Limite float64
}

func (t *Tarjeta) Pagar(monto float64) error {
	if monto > t.Limite {
		return fmt.Errorf("límite excedido: monto $%.2f > límite $%.2f", monto, t.Limite)
	}
	t.Limite -= monto
	fmt.Printf("  Cobrando $%.2f a tarjeta terminada en %s\n", monto, t.Numero[len(t.Numero)-4:])
	return nil
}

func (t *Tarjeta) Comision() float64 { return 0.03 } // 3%

// ─────────────────────────────────────────────────────────
// EFECTIVO
// ─────────────────────────────────────────────────────────

type Efectivo struct{}

func (Efectivo) Pagar(monto float64) error {
	fmt.Printf("  Recibiendo $%.2f en efectivo\n", monto)
	return nil
}

func (Efectivo) Comision() float64 { return 0 } // sin comisión

// ─────────────────────────────────────────────────────────
// MERCADOPAGO (o cualquier billetera virtual)
// ─────────────────────────────────────────────────────────

type MercadoPago struct {
	Email string
	Saldo float64
}

func (m *MercadoPago) Pagar(monto float64) error {
	if monto > m.Saldo {
		return fmt.Errorf("saldo insuficiente: tenés $%.2f, necesitás $%.2f", m.Saldo, monto)
	}
	m.Saldo -= monto
	fmt.Printf("  Cobrando $%.2f desde cuenta MP de %s\n", monto, m.Email)
	return nil
}

func (m *MercadoPago) Comision() float64 { return 0.015 } // 1.5%

// ─────────────────────────────────────────────────────────
// LA FUNCIÓN POLIMÓRFICA: procesarCompra
// ─────────────────────────────────────────────────────────
// Esta función NO sabe (ni le importa) si mp es una Tarjeta, un
// Efectivo o un MercadoPago. Solo sabe que cumple MetodoPago.
// Si mañana agregás "Transferencia" o "Cripto", esta función NO
// cambia una sola línea.

func procesarCompra(mp MetodoPago, montoBase float64) error {
	comision := montoBase * mp.Comision()
	total := montoBase + comision

	fmt.Printf("Compra de $%.2f (+ $%.2f de comisión = $%.2f total)\n",
		montoBase, comision, total)

	return mp.Pagar(total)
}

func main() {
	fmt.Println("=== Polimorfismo: un checkout, tres métodos de pago ===")

	tarjeta := &Tarjeta{Numero: "4532111122223333", Limite: 50000}
	efectivo := Efectivo{}
	mp := &MercadoPago{Email: "cliente@mail.com", Saldo: 20000}

	metodos := []MetodoPago{tarjeta, efectivo, mp}
	montos := []float64{15000, 5000, 8000}

	for i, metodo := range metodos {
		fmt.Printf("\n--- Pago %d ---\n", i+1)
		if err := procesarCompra(metodo, montos[i]); err != nil {
			fmt.Println("  ERROR:", err)
		}
	}

	// ─────────────────────────────────────────────────────────
	// EL PODER REAL: agregar un método de pago nuevo sin tocar
	// procesarCompra
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Agregando Transferencia sin tocar procesarCompra ===")

	transferencia := Transferencia{CBU: "0000003100012345678901"}
	if err := procesarCompra(transferencia, 3000); err != nil {
		fmt.Println("  ERROR:", err)
	}

	// ─────────────────────────────────────────────────────────
	// COMPARÁ ESTO CON EL "MAL": una función por cada tipo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== MAL: sin interfaces habría que hacer esto ===")
	fmt.Println("  func procesarConTarjeta(t *Tarjeta, monto float64) error")
	fmt.Println("  func procesarConEfectivo(e Efectivo, monto float64) error")
	fmt.Println("  func procesarConMP(m *MercadoPago, monto float64) error")
	fmt.Println("  ... una función más por cada método nuevo, código duplicado")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Polimorfismo    → una función, comportamiento distinto según el tipo")
	fmt.Println("  La función      → solo conoce la interfaz, nunca el tipo concreto")
	fmt.Println("  Agregar un tipo → CERO cambios en las funciones que ya usan la interfaz")
	fmt.Println("  Esto es lo que  → hace que agregar features no rompa código viejo")
}

type Transferencia struct {
	CBU string
}

func (Transferencia) Pagar(monto float64) error {
	fmt.Printf("  Transferencia de $%.2f iniciada\n", monto)
	return nil
}

func (Transferencia) Comision() float64 { return 0 }
