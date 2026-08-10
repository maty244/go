package main

import "fmt"

type Dinero int64 // centavos

func dineroTexto(monto Dinero) string {
	return fmt.Sprintf("$%d.%02d", monto/100, monto%100)
}

// =========================================================
// EJERCICIO PRÁCTICO: SISTEMA DE CUENTAS BANCARIAS
// =========================================================
// Aplicamos todo lo aprendido sobre punteros en un ejercicio
// real y concreto. Construimos un sistema bancario simple
// donde varias funciones reciben punteros a las mismas cuentas.
//
// El objetivo es ver CUÁNDO usar punteros y CUÁNDO no,
// en el contexto de un problema real.

type Cuenta struct {
	ID      int
	Titular string
	Saldo   Dinero
	Activa  bool
}

type Transaccion struct {
	Tipo  string // "depósito", "retiro", "transferencia"
	Monto Dinero
	Desde *Cuenta // puede ser nil (en depósitos)
	Hacia *Cuenta // puede ser nil (en retiros)
}

// ─────────────────────────────────────────────────────────
// OPERACIONES QUE MODIFICAN → puntero
// ─────────────────────────────────────────────────────────

func depositar(c *Cuenta, monto Dinero) error {
	if !c.Activa {
		return fmt.Errorf("cuenta %d de %s está inactiva", c.ID, c.Titular)
	}
	if monto <= 0 {
		return fmt.Errorf("monto inválido: %s", dineroTexto(monto))
	}
	c.Saldo += monto
	return nil
}

func retirar(c *Cuenta, monto Dinero) error {
	if !c.Activa {
		return fmt.Errorf("cuenta %d de %s está inactiva", c.ID, c.Titular)
	}
	if monto <= 0 {
		return fmt.Errorf("monto inválido: %s", dineroTexto(monto))
	}
	if c.Saldo < monto {
		return fmt.Errorf("saldo insuficiente: tenés %s, querés retirar %s",
			dineroTexto(c.Saldo), dineroTexto(monto))
	}
	c.Saldo -= monto
	return nil
}

// transferir modifica DOS cuentas → dos punteros
func transferir(origen, destino *Cuenta, monto Dinero) error {
	if err := retirar(origen, monto); err != nil {
		return fmt.Errorf("transferencia fallida (origen): %w", err)
	}
	if err := depositar(destino, monto); err != nil {
		// Si el depósito falla, revertimos el retiro
		origen.Saldo += monto
		return fmt.Errorf("transferencia fallida (destino): %w", err)
	}
	return nil
}

func cerrarCuenta(c *Cuenta) {
	c.Activa = false
}

// ─────────────────────────────────────────────────────────
// OPERACIONES DE SOLO LECTURA → valor (no puntero)
// ─────────────────────────────────────────────────────────

func mostrarCuenta(c Cuenta) {
	estado := "activa"
	if !c.Activa {
		estado = "INACTIVA"
	}
	fmt.Printf("  [%d] %-12s | Saldo: %9s | %s\n",
		c.ID, c.Titular, dineroTexto(c.Saldo), estado)
}

func generarExtracto(c Cuenta) string {
	estado := "Activa"
	if !c.Activa {
		estado = "Inactiva"
	}
	return fmt.Sprintf(
		"=== Extracto Cuenta #%d ===\nTitular: %s\nSaldo: %s\nEstado: %s",
		c.ID, c.Titular, dineroTexto(c.Saldo), estado,
	)
}

// ─────────────────────────────────────────────────────────
// FUNCIÓN QUE RETORNA UN PUNTERO (new cuenta)
// ─────────────────────────────────────────────────────────
var proximoID = 1

func nuevaCuenta(titular string, deposito Dinero) (*Cuenta, error) {
	if titular == "" {
		return nil, fmt.Errorf("el titular no puede estar vacío")
	}
	if deposito < 0 {
		return nil, fmt.Errorf("el depósito inicial no puede ser negativo")
	}
	c := &Cuenta{
		ID:      proximoID,
		Titular: titular,
		Saldo:   deposito,
		Activa:  true,
	}
	proximoID++
	return c, nil
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║   EJERCICIO: SISTEMA BANCARIO     ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// CREAR CUENTAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Abriendo cuentas ---")

	cuentaAna, err := nuevaCuenta("Ana García", 100_000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mostrarCuenta(*cuentaAna)

	cuentaCarlos, err := nuevaCuenta("Carlos López", 50_000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mostrarCuenta(*cuentaCarlos)

	_, err = nuevaCuenta("", 100) // cuenta inválida
	if err != nil {
		fmt.Println("Error esperado:", err)
	}

	// ─────────────────────────────────────────────────────────
	// DEPOSITAR Y RETIRAR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Operaciones ---")

	ops := []struct {
		fn     func() error
		nombre string
	}{
		{func() error { return depositar(cuentaAna, 50_000) }, "Depósito $500 → Ana"},
		{func() error { return retirar(cuentaCarlos, 20_000) }, "Retiro $200 de Carlos"},
		{func() error { return retirar(cuentaAna, 500_000) }, "Retiro $5000 de Ana (fallo esperado)"},
		{func() error { return depositar(cuentaCarlos, -5_000) }, "Depósito -$50 (fallo esperado)"},
	}

	for _, op := range ops {
		if err := op.fn(); err != nil {
			fmt.Printf("  ✗ %s: %v\n", op.nombre, err)
		} else {
			fmt.Printf("  ✓ %s: OK\n", op.nombre)
		}
	}

	fmt.Println("\n--- Saldos tras operaciones ---")
	mostrarCuenta(*cuentaAna)
	mostrarCuenta(*cuentaCarlos)

	// ─────────────────────────────────────────────────────────
	// TRANSFERENCIA: dos punteros a la vez
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Transferencia ---")
	fmt.Println("Antes:")
	mostrarCuenta(*cuentaAna)
	mostrarCuenta(*cuentaCarlos)

	if err := transferir(cuentaAna, cuentaCarlos, 30_000); err != nil {
		fmt.Println("Error en transferencia:", err)
	} else {
		fmt.Println("Transferencia $300 de Ana → Carlos: OK")
	}

	fmt.Println("Después:")
	mostrarCuenta(*cuentaAna)
	mostrarCuenta(*cuentaCarlos)

	// ─────────────────────────────────────────────────────────
	// CERRAR CUENTA Y VERIFICAR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Cerrar cuenta ---")
	cerrarCuenta(cuentaCarlos)
	mostrarCuenta(*cuentaCarlos)

	if err := depositar(cuentaCarlos, 10_000); err != nil {
		fmt.Println("Depósito en cuenta cerrada:", err)
	}

	// ─────────────────────────────────────────────────────────
	// EXTRACTO (solo lectura, recibe valor)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Extracto ---")
	// Pasamos *cuentaAna (valor del struct) porque solo leemos
	fmt.Println(generarExtracto(*cuentaAna))

	// ─────────────────────────────────────────────────────────
	// LO QUE APRENDIMOS EN ESTE EJERCICIO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Lo aprendido ===")
	fmt.Println("✓ nuevaCuenta() retorna *Cuenta (la ubicación física la decide Go)")
	fmt.Println("✓ depositar/retirar/cerrar reciben *Cuenta → modifican el original")
	fmt.Println("✓ mostrarCuenta/generarExtracto reciben Cuenta → solo leen")
	fmt.Println("✓ transferir recibe dos *Cuenta → modifica ambas")
	fmt.Println("✓ Pasar &cuenta vs *puntero: ambos dan el valor del struct")
	fmt.Println("✓ Cuando una operación falla a mitad, hay que revertir (rollback)")
}
