package main

import "fmt"

// =========================================================
// PARÁMETROS VARIABLES (VARIADIC)
// =========================================================
// Una función variádica acepta un número variable de argumentos.
// El parámetro variádico se declara con "..." antes del tipo.
//
//   func nombre(params ...tipo) { }
//
// Dentro de la función, el parámetro variádico es un SLICE del tipo.
// Puede llamarse con 0, 1, 2, o N argumentos.
// Siempre debe ser el ÚLTIMO parámetro de la función.

// ─────────────────────────────────────────────────────────
// FUNCIÓN VARIÁDICA BÁSICA
// ─────────────────────────────────────────────────────────
func sumar(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

// ─────────────────────────────────────────────────────────
// VARIÁDICO COMBINADO CON PARÁMETROS NORMALES
// ─────────────────────────────────────────────────────────
// El parámetro variádico SIEMPRE va al final.
func calcularTotalConImpuesto(impuesto float64, precios ...float64) float64 {
	subtotal := 0.0
	for _, p := range precios {
		subtotal += p
	}
	return subtotal * (1 + impuesto)
}

// ─────────────────────────────────────────────────────────
// CASO DE LA TIENDA: calcular total de varios productos
// ─────────────────────────────────────────────────────────
type Producto struct {
	Nombre string
	Precio float64
}

// calcularTotalProductos recibe cualquier cantidad de productos
func calcularTotalProductos(productos ...Producto) float64 {
	total := 0.0
	for _, p := range productos {
		total += p.Precio
	}
	return total
}

// listarProductos muestra una lista de productos con numeración
func listarProductos(titulo string, productos ...Producto) {
	fmt.Printf("\n  --- %s ---\n", titulo)
	if len(productos) == 0 {
		fmt.Println("  (sin productos)")
		return
	}
	for i, p := range productos {
		fmt.Printf("  %d. %-15s $%.2f\n", i+1, p.Nombre, p.Precio)
	}
	fmt.Printf("  Total: $%.2f\n", calcularTotalProductos(productos...))
}

// imprimirLog acepta cualquier cantidad de mensajes (como fmt.Println)
func imprimirLog(nivel string, mensajes ...string) {
	for _, msg := range mensajes {
		fmt.Printf("[%s] %s\n", nivel, msg)
	}
}

// max retorna el mayor de varios números
func max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	mayor := nums[0]
	for _, n := range nums[1:] {
		if n > mayor {
			mayor = n
		}
	}
	return mayor
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║   PARÁMETROS VARIÁDICOS       ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// LLAMAR CON DIFERENTE CANTIDAD DE ARGUMENTOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== sumar con distintas cantidades ===")
	fmt.Println("sumar():           ", sumar())              // 0
	fmt.Println("sumar(5):          ", sumar(5))             // 5
	fmt.Println("sumar(1,2,3):      ", sumar(1, 2, 3))       // 6
	fmt.Println("sumar(1,2,3,4,5):  ", sumar(1, 2, 3, 4, 5)) // 15

	// ─────────────────────────────────────────────────────────
	// EXPANDIR UN SLICE CON ...
	// ─────────────────────────────────────────────────────────
	// Si ya tenés un slice, podés expandirlo como argumentos con ...
	fmt.Println("\n=== Expandir slice con ... ===")
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("slice:", numeros)
	fmt.Println("sumar(numeros...):", sumar(numeros...)) // expande el slice

	// ─────────────────────────────────────────────────────────
	// VARIÁDICO CON PARÁMETROS NORMALES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== calcularTotalConImpuesto ===")
	const IVA = 0.21
	total1 := calcularTotalConImpuesto(IVA, 100.0)
	total2 := calcularTotalConImpuesto(IVA, 100.0, 50.0, 25.0)
	fmt.Printf("Un precio: $%.2f\n", total1)
	fmt.Printf("Tres precios: $%.2f\n", total2)

	// ─────────────────────────────────────────────────────────
	// PRODUCTOS EN LA TIENDA
	// ─────────────────────────────────────────────────────────
	notebook := Producto{"Notebook", 1500.00}
	mouse := Producto{"Mouse", 25.99}
	teclado := Producto{"Teclado", 75.50}
	monitor := Producto{"Monitor", 450.00}
	auriculares := Producto{"Auriculares", 80.00}

	fmt.Println("\n=== Catálogo variádico ===")
	listarProductos("Carrito de Ana", notebook, mouse)
	listarProductos("Carrito de Carlos", teclado, monitor, auriculares)
	listarProductos("Carrito vacío")

	// También podemos pasar un slice existente
	todoElCatalogo := []Producto{notebook, mouse, teclado, monitor, auriculares}
	listarProductos("Todo el catálogo", todoElCatalogo...) // expande el slice

	// ─────────────────────────────────────────────────────────
	// MÁXIMO DE VARIOS NÚMEROS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Máximo variádico ===")
	fmt.Printf("max(3, 1, 4, 1, 5, 9, 2, 6): %.0f\n",
		max(3, 1, 4, 1, 5, 9, 2, 6))

	// ─────────────────────────────────────────────────────────
	// LOG CON VARIÁDICO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Log variádico ===")
	imprimirLog("INFO", "Servidor iniciado")
	imprimirLog("INFO", "Conexión establecida", "Base de datos lista")
	imprimirLog("ERROR", "No se pudo conectar", "Reintentando...", "Fallo total")

	// ─────────────────────────────────────────────────────────
	// CÓMO EL VARIÁDICO SE VE DENTRO DE LA FUNCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El variádico es un slice adentro ===")
	mostrarTipo(1, 2, 3)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("func f(args ...int)         → variádico, recibe 0 o más ints")
	fmt.Println("func f(s string, a ...int)  → parámetros normales + variádico al final")
	fmt.Println("f(1, 2, 3)                  → llamar con argumentos individuales")
	fmt.Println("f(slice...)                 → expandir un slice como argumentos")
	fmt.Println("Dentro de la función: args es un []int (un slice normal)")
}

func mostrarTipo(nums ...int) {
	fmt.Printf("  Tipo de 'nums' dentro: %T\n", nums)
	fmt.Printf("  Valor: %v | len: %d\n", nums, len(nums))
}
