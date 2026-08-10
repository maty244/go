package main

import "fmt"

// =========================================================
// RETORNANDO MÚLTIPLES VALORES
// =========================================================
// Go permite que una función retorne más de un valor.
// Esto es nativo del lenguaje, no un hack.
//
// Sintaxis:
//   func nombre() (tipo1, tipo2) {
//       return val1, val2
//   }
//
// El caso más común en Go: retornar (resultado, error).

// ─────────────────────────────────────────────────────────
// RETORNO BÁSICO: dos valores
// ─────────────────────────────────────────────────────────
// Retorna cociente y resto de una división
func dividir(a, b int) (int, int) {
	return a / b, a % b
}

// Retorna mínimo y máximo de un slice
func minMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// ─────────────────────────────────────────────────────────
// EL PATRÓN MÁS COMÚN EN GO: (resultado, error)
// ─────────────────────────────────────────────────────────
// La convención de Go: si algo puede fallar, retorná
// el resultado Y un error. El error es el ÚLTIMO valor.

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

// buscarProducto retorna el producto encontrado y un error
func buscarProducto(catalogo []Producto, id int) (Producto, error) {
	for _, p := range catalogo {
		if p.ID == id {
			return p, nil // nil = "sin error"
		}
	}
	return Producto{}, fmt.Errorf("producto con ID %d no encontrado", id)
}

// procesarCompra retorna el total pagado y un error si hay problema
func procesarCompra(p Producto, cantidad int) (float64, error) {
	if cantidad <= 0 {
		return 0, fmt.Errorf("cantidad inválida: %d", cantidad)
	}
	if p.Stock < cantidad {
		return 0, fmt.Errorf("stock insuficiente: pedido %d, disponible %d",
			cantidad, p.Stock)
	}
	total := p.Precio * float64(cantidad)
	return total, nil
}

// calcularEstadisticas retorna tres valores: suma, promedio, cantidad
func calcularEstadisticas(precios []float64) (suma float64, promedio float64, cantidad int) {
	cantidad = len(precios)
	if cantidad == 0 {
		return 0, 0, 0
	}
	for _, p := range precios {
		suma += p
	}
	promedio = suma / float64(cantidad)
	return // naked return, retorna los valores nombrados
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║   RETORNO MÚLTIPLE DE VALORES ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// CAPTURAR MÚLTIPLES RETORNOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Retorno básico de dos valores ===")

	cociente, resto := dividir(17, 5)
	fmt.Printf("17 ÷ 5 = %d con resto %d\n", cociente, resto)

	numeros := []int{7, 2, 9, 1, 5, 8, 3}
	minimo, maximo := minMax(numeros)
	fmt.Printf("Números: %v\n", numeros)
	fmt.Printf("Mín: %d | Máx: %d\n", minimo, maximo)

	// ─────────────────────────────────────────────────────────
	// IGNORAR UN VALOR CON _
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Ignorar valores con _ ===")

	_, soloResto := dividir(17, 5)    // solo nos interesa el resto
	soloCociente, _ := dividir(17, 5) // solo nos interesa el cociente
	fmt.Printf("Solo resto: %d\n", soloResto)
	fmt.Printf("Solo cociente: %d\n", soloCociente)

	// ─────────────────────────────────────────────────────────
	// PATRÓN (resultado, error): la forma Go de manejar fallos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrón (resultado, error) ===")

	catalogo := []Producto{
		{1, "Notebook", 1500.00, 5},
		{2, "Mouse", 25.99, 20},
		{3, "Teclado", 75.50, 3},
	}

	// Buscar producto existente
	if producto, err := buscarProducto(catalogo, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Encontrado: %s — $%.2f\n", producto.Nombre, producto.Precio)
	}

	// Buscar producto inexistente
	if producto, err := buscarProducto(catalogo, 99); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Producto:", producto.Nombre)
	}

	// ─────────────────────────────────────────────────────────
	// CADENA DE RETORNOS: usar el resultado de uno para llamar al otro
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cadena de operaciones ===")

	producto, err := buscarProducto(catalogo, 3) // Teclado, stock=3
	if err != nil {
		fmt.Println("Error buscando:", err)
		return
	}

	total, err := procesarCompra(producto, 2) // 2 teclados
	if err != nil {
		fmt.Println("Error comprando:", err)
		return
	}
	fmt.Printf("Compra exitosa: 2x %s = $%.2f\n", producto.Nombre, total)

	// Intentar comprar más de lo disponible
	_, err = procesarCompra(producto, 10)
	if err != nil {
		fmt.Println("Error (esperado):", err)
	}

	// ─────────────────────────────────────────────────────────
	// TRES VALORES DE RETORNO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tres valores de retorno ===")

	precios := []float64{1500.00, 25.99, 75.50, 450.00, 80.00}
	suma, promedio, cantidad := calcularEstadisticas(precios)
	fmt.Printf("Precios: %v\n", precios)
	fmt.Printf("Suma: $%.2f | Promedio: $%.2f | Cantidad: %d\n",
		suma, promedio, cantidad)

	// ─────────────────────────────────────────────────────────
	// USAR EL RETORNO DIRECTAMENTE SIN VARIABLE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Usar retorno directo ===")
	fmt.Println("Suma directa:", sumarVariosEnteros(dividir(17, 5))) // usa ambos retornos

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("func f() (int, int)           → retorna dos enteros")
	fmt.Println("func f() (Producto, error)    → patrón más común en Go")
	fmt.Println("a, b := f()                   → capturar ambos")
	fmt.Println("a, _ := f()                   → ignorar uno con _")
	fmt.Println("_, b := f()                   → ignorar el primero")
	fmt.Println("El error siempre va al FINAL por convención")
}

func sumarVariosEnteros(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
