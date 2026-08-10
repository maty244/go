package main

import "fmt"

// =========================================================
// DECLARANDO Y LLAMANDO FUNCIONES
// =========================================================
// Acá profundizamos en todas las formas de declarar funciones
// y las reglas que gobiernan su uso en Go.
//
// REGLAS IMPORTANTES:
//   - En Go, el orden de declaración NO importa.
//     Podés llamar una función antes de declararla.
//   - Las funciones que empiezan con MAYÚSCULA son exportadas
//     (visibles desde otros paquetes).
//   - Las funciones que empiezan con minúscula son privadas
//     (solo visibles dentro del mismo paquete).

// ─────────────────────────────────────────────────────────
// FUNCIÓN EXPORTADA (Mayúscula) — visible desde otros paquetes
// ─────────────────────────────────────────────────────────
// Cuando otro paquete importa este, puede usar CalcularTotal.
func CalcularTotal(precio float64, cantidad int) float64 {
	return precio * float64(cantidad)
}

// ─────────────────────────────────────────────────────────
// FUNCIÓN PRIVADA (minúscula) — solo visible en este paquete
// ─────────────────────────────────────────────────────────
func aplicarDescuento(precio, descuento float64) float64 {
	return precio * (1 - descuento)
}

// ─────────────────────────────────────────────────────────
// PARÁMETROS DEL MISMO TIPO: forma abreviada
// ─────────────────────────────────────────────────────────
// En vez de: func sumar(a int, b int) int
// Podemos escribir:
func sumar(a, b int) int {
	return a + b
}

// Con tres parámetros del mismo tipo
func volumenCaja(largo, ancho, alto float64) float64 {
	return largo * ancho * alto
}

// ─────────────────────────────────────────────────────────
// LLAMAR FUNCIONES ANTES DE DECLARARLAS
// ─────────────────────────────────────────────────────────
// En Go esto es completamente válido.
// El compilador lee todo el paquete antes de compilar.
func main() {
	fmt.Printf("=== Declarando y llamando funciones ===\n\n")

	// ─────────────────────────────────────────────────────────
	// 1. LLAMADA SIMPLE (ignorando el retorno)
	// ─────────────────────────────────────────────────────────
	imprimirSeparador() // declarada más abajo, pero funciona igual

	// ─────────────────────────────────────────────────────────
	// 2. LLAMADA CAPTURANDO EL RETORNO
	// ─────────────────────────────────────────────────────────
	total := CalcularTotal(25.99, 3)
	fmt.Printf("3 × $25.99 = $%.2f\n", total)

	// ─────────────────────────────────────────────────────────
	// 3. LLAMADA USANDO EL RETORNO DIRECTAMENTE EN UNA EXPRESIÓN
	// ─────────────────────────────────────────────────────────
	// No siempre necesitás una variable intermedia.
	fmt.Printf("Total con descuento 10%%: $%.2f\n",
		aplicarDescuento(CalcularTotal(25.99, 3), 0.10))

	// ─────────────────────────────────────────────────────────
	// 4. PARÁMETROS ABREVIADOS
	// ─────────────────────────────────────────────────────────
	fmt.Printf("\n5 + 3 = %d\n", sumar(5, 3))
	fmt.Printf("Caja 2×3×4 = %.0f cm³\n", volumenCaja(2, 3, 4))

	// ─────────────────────────────────────────────────────────
	// 5. FUNCIONES COMO ARGUMENTOS DE OTRAS FUNCIONES
	// ─────────────────────────────────────────────────────────
	// Podés pasar el RESULTADO de una función como argumento de otra.
	fmt.Printf("\nDesde la tienda:\n")
	mostrarPrecioFinal("Notebook", CalcularTotal(1500.00, 1))
	mostrarPrecioFinal("Mouse x3", CalcularTotal(25.99, 3))

	imprimirSeparador()

	// ─────────────────────────────────────────────────────────
	// 6. FUNCIONES QUE LLAMAN A OTRAS FUNCIONES
	// ─────────────────────────────────────────────────────────
	// Es completamente normal y deseable.
	resumen := generarResumenPedido("Matias", []string{"Notebook", "Mouse", "Teclado"})
	fmt.Println(resumen)

	// ─────────────────────────────────────────────────────────
	// 7. VISIBILIDAD: EXPORTADA VS PRIVADA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Visibilidad de funciones ===")
	fmt.Println("CalcularTotal  → exportada (MAYÚSCULA): visible desde otros paquetes")
	fmt.Println("aplicarDescuento → privada (minúscula): solo en este paquete")
	fmt.Println()
	fmt.Println("Convención de nombres en Go:")
	fmt.Println("  CalcularTotal       → exportada, PascalCase")
	fmt.Println("  aplicarDescuento    → privada, camelCase")
	fmt.Println("  CONSTANTE_GLOBAL    → NO en Go (Go prefiere camelCase para todo)")

	imprimirSeparador()
}

// Declarada DESPUÉS de main pero se puede llamar desde main.
// Go lee todo el archivo antes de compilar.
func imprimirSeparador() {
	fmt.Println("─────────────────────────────")
}

func mostrarPrecioFinal(producto string, precio float64) {
	fmt.Printf("  %-15s → $%.2f\n", producto, precio)
}

// Función que llama a otras funciones internamente
func generarResumenPedido(cliente string, productos []string) string {
	encabezado := fmt.Sprintf("=== Pedido de %s ===", cliente)
	cuerpo := ""
	for i, p := range productos {
		cuerpo += fmt.Sprintf("\n  %d. %s", i+1, p)
	}
	pie := fmt.Sprintf("\n  Total de items: %d", len(productos))
	return encabezado + cuerpo + pie
}
