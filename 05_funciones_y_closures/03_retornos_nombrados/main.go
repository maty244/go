package main

import "fmt"

// =========================================================
// RETORNANDO VALORES CON NOMBRE Y VALORES EN BLANCO
// =========================================================
// Go permite darle nombres a los valores de retorno.
// Esto cambia cómo se declaran y documentan las funciones.
//
// RETORNO NOMBRADO:
//   func f() (resultado int, err error) { }
//   - Los nombres actúan como variables pre-declaradas en cero.
//   - Podés usar "return" sin argumentos (naked return).
//   - Documenta claramente qué retorna la función.
//
// VALOR EN BLANCO (_):
//   - Descarta un valor de retorno que no necesitamos.
//   - "Le decimos a Go: sé que esto retorna algo, pero no me importa."

// ─────────────────────────────────────────────────────────
// RETORNO NOMBRADO BÁSICO
// ─────────────────────────────────────────────────────────

// Sin retorno nombrado (forma normal)
func divisionNormal(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

// Con retorno nombrado (los nombres documentan el código)
func divisionNombrada(a, b float64) (resultado float64, err error) {
	// "resultado" y "err" son variables en zero value (0 y nil)
	if b == 0 {
		err = fmt.Errorf("división por cero") // asignamos al nombre
		return                                // naked return: retorna resultado(0) y err
	}
	resultado = a / b // asignamos al nombre
	return            // naked return: retorna resultado y nil
}

// ─────────────────────────────────────────────────────────
// RETORNO NOMBRADO: CASO TIENDA
// ─────────────────────────────────────────────────────────
type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

// La firma documenta exactamente qué retorna:
// subtotal, descuento aplicado, total final
func calcularOrden(productos []Producto, descuento float64) (
	subtotal float64,
	montoDescuento float64,
	total float64,
) {
	// Las tres variables ya existen con zero value (0.0)
	for _, p := range productos {
		subtotal += p.Precio
	}
	montoDescuento = subtotal * descuento
	total = subtotal - montoDescuento
	return // naked return: retorna los tres valores
}

// Buscar producto con retorno nombrado
func buscarEnCatalogo(catalogo []Producto, nombre string) (p Producto, encontrado bool) {
	// "p" es Producto{} (zero value) y "encontrado" es false
	for _, prod := range catalogo {
		if prod.Nombre == nombre {
			p = prod
			encontrado = true
			return // retorna el producto encontrado y true
		}
	}
	return // retorna Producto{} y false (zero values)
}

// ─────────────────────────────────────────────────────────
// CUÁNDO LOS RETORNOS NOMBRADOS AYUDAN Y CUÁNDO NO
// ─────────────────────────────────────────────────────────

// ÚTIL: función con lógica compleja donde los nombres documentan
func parsearPedido(raw string) (cliente string, monto float64, valido bool) {
	// Simulación: en la vida real parsearías JSON, CSV, etc.
	if raw == "" {
		// valido ya es false, cliente="" y monto=0 → naked return documenta el caso error
		return
	}
	// Parseo simulado
	cliente = "Ana García"
	monto = 150.00
	valido = true
	return
}

// MENOS ÚTIL (naked return en funciones cortas puede confundir):
func sumar(a, b int) (resultado int) {
	resultado = a + b
	return // equivalente a: return a + b
	// En este caso sería más claro: return a + b
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║  RETORNO NOMBRADO Y EN BLANCO ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// COMPARACIÓN: normal vs nombrado
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Normal vs Nombrado (mismo resultado) ===")

	r1, e1 := divisionNormal(10, 3)
	r2, e2 := divisionNombrada(10, 3)
	fmt.Printf("Normal:   %.4f, %v\n", r1, e1)
	fmt.Printf("Nombrado: %.4f, %v\n", r2, e2)

	_, e3 := divisionNombrada(10, 0)
	fmt.Printf("Error nombrado: %v\n", e3)

	// ─────────────────────────────────────────────────────────
	// RETORNO NOMBRADO EN LA TIENDA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== calcularOrden con retorno nombrado ===")

	carrito := []Producto{
		{"Notebook", 1500.00, 5},
		{"Mouse", 25.99, 20},
		{"Teclado", 75.50, 3},
	}

	subtotal, descAplicado, total := calcularOrden(carrito, 0.10)
	fmt.Printf("  Subtotal:    $%.2f\n", subtotal)
	fmt.Printf("  Descuento:   -$%.2f\n", descAplicado)
	fmt.Printf("  Total:       $%.2f\n", total)

	// ─────────────────────────────────────────────────────────
	// BUSCAR CON RETORNO NOMBRADO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== buscarEnCatalogo con retorno nombrado ===")

	if p, ok := buscarEnCatalogo(carrito, "Mouse"); ok {
		fmt.Printf("  Encontrado: %s — $%.2f\n", p.Nombre, p.Precio)
	}

	if _, ok := buscarEnCatalogo(carrito, "Monitor"); !ok {
		fmt.Println("  'Monitor' no está en el catálogo")
	}

	// ─────────────────────────────────────────────────────────
	// VALORES EN BLANCO (_)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Valores en blanco (_) ===")

	// Ignorar el error (solo cuando estás SEGURO de que no falla)
	resultado, _ := divisionNormal(10, 2)
	fmt.Printf("Solo el resultado: %.1f\n", resultado)

	// Ignorar el resultado, solo nos importa si hay error
	_, err := divisionNormal(10, 0)
	if err != nil {
		fmt.Println("Hubo un error:", err)
	}

	// Ignorar valores en parseo (patrón ok)
	_, monto, valido := parsearPedido("datos del pedido")
	if valido {
		fmt.Printf("Pedido válido por $%.2f\n", monto)
	}

	// Ignorar todo (raramente justificado en producción)
	_, _, _ = calcularOrden(carrito, 0.05) // solo ejecutamos por efecto secundario

	// ─────────────────────────────────────────────────────────
	// CASO ESPECIAL: _ en imports y range
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== _ en otros contextos ===")

	numeros := []int{10, 20, 30}

	// _ en range: ignorar el índice
	for _, n := range numeros {
		fmt.Print(n, " ")
	}
	fmt.Println()

	// _ en range: ignorar el valor
	for i := range numeros {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Retorno nombrado:")
	fmt.Println("  func f() (r int, err error)  → nombres = variables pre-declaradas")
	fmt.Println("  return                        → naked return: retorna los nombres")
	fmt.Println("  Útil: documenta el significado de cada retorno")
	fmt.Println("  Útil: funciones con lógica compleja y múltiples puntos de retorno")
	fmt.Println("  Evitar: en funciones cortas (puede confundir)")
	fmt.Println()
	fmt.Println("Valor en blanco (_):")
	fmt.Println("  a, _ := f()    → ignorar el segundo retorno")
	fmt.Println("  _, b := f()    → ignorar el primer retorno")
	fmt.Println("  for _, v := range   → ignorar el índice")
	fmt.Println("  for i := range      → ignorar el valor")
}
