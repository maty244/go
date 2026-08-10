package main

import "fmt"

// =========================================================
// COVERAGE: ¿QUÉ PORCENTAJE DE TU CÓDIGO TESTEÁS?
// =========================================================
// "Coverage" (cobertura) mide qué porcentaje de los STATEMENTS de tu
// código se ejecutaron durante los tests. No es exactamente un conteo de
// líneas. Go lo mide de forma
// integrada, sin herramientas externas:
//
//   go test -cover ./ruta/
//
// Esto NO se ve corriendo `go run` (por eso este archivo no puede
// "mostrar" el coverage directamente): es un reporte que genera
// el comando `go test`. Correlo vos mismo para ver el resultado:
//
//   go test -cover ./10_testing/04_coverage/
//
// 100% de cobertura NO significa "sin bugs" (podés ejecutar una
// línea con un valor que no revela el problema). Pero 0% o muy
// bajo SÍ significa con certeza que hay código sin ninguna
// verificación.

// AplicarDescuento calcula el precio final según el tipo de
// cliente. Tiene VARIAS ramas: en main_test.go vamos a testear
// solo ALGUNAS a propósito, para ver un coverage incompleto.
type Dinero int64 // centavos

func AplicarDescuento(precio Dinero, tipoCliente string) Dinero {
	switch tipoCliente {
	case "VIP":
		return precio * 80 / 100 // 20% off
	case "frecuente":
		return precio * 90 / 100 // 10% off
	case "empleado":
		return precio * 50 / 100 // 50% off
	default:
		return precio // sin descuento
	}
}

func main() {
	fmt.Println("=== AplicarDescuento con varias ramas ===")

	tipos := []string{"VIP", "frecuente", "empleado", "regular"}
	for _, tipo := range tipos {
		fmt.Printf("  %-10s → %d centavos\n", tipo, AplicarDescuento(100_000, tipo))
	}

	fmt.Println("\n=== Cómo ver el coverage de este archivo ===")
	fmt.Println("  go test -cover ./10_testing/04_coverage/")
	fmt.Println(`  → un porcentaje, por ejemplo: "coverage: 13.0% of statements"`)
	fmt.Println("  (se ve más bajo de lo esperado: func main() NUNCA se ejecuta")
	fmt.Println("   durante go test, así que sus líneas cuentan como no cubiertas.")
	fmt.Println("   Por eso el coverage real de un proyecto se mide sobre la")
	fmt.Println("   lógica de negocio, no sobre main())")
	fmt.Println()
	fmt.Println("  Para un reporte VISUAL (qué líneas exactas faltan):")
	fmt.Println("  go test -coverprofile=coverage.out ./10_testing/04_coverage/")
	fmt.Println("  go tool cover -html=coverage.out")
	fmt.Println("  (abre el navegador con el código coloreado: verde = cubierto, rojo = no)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  go test -cover              → % de statements ejecutados por los tests")
	fmt.Println("  go test -coverprofile=x.out → guarda el detalle línea por línea")
	fmt.Println("  go tool cover -html=x.out   → reporte visual en el navegador")
	fmt.Println("  100% coverage               → NO garantiza que no haya bugs")
	fmt.Println("  Coverage muy bajo           → SÍ garantiza que falta testear")
}
