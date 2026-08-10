package main

import (
	"fmt"
	"strings"
)

// =========================================================
// RECURSIVIDAD
// =========================================================
// Una función recursiva es una función que se LLAMA A SÍ MISMA.
// Toda función recursiva necesita:
//   1. CASO BASE: la condición que detiene la recursión.
//   2. CASO RECURSIVO: la llamada a sí misma con un problema MÁS PEQUEÑO.
//
// Sin caso base → stack overflow (se llama infinitamente).
// Sin reducción del problema → nunca llega al caso base.

// ─────────────────────────────────────────────────────────
// EJEMPLO CLÁSICO 1: Factorial
// ─────────────────────────────────────────────────────────
// n! = n × (n-1) × (n-2) × ... × 1
// Caso base:     factorial(0) = 1
// Caso recursivo: factorial(n) = n × factorial(n-1)

func factorial(n int) int {
	if n <= 1 { // CASO BASE: factorial(0) = factorial(1) = 1
		return 1
	}
	return n * factorial(n-1) // CASO RECURSIVO: se llama con n-1
}

// ─────────────────────────────────────────────────────────
// EJEMPLO CLÁSICO 2: Fibonacci
// ─────────────────────────────────────────────────────────
// Secuencia: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
// fib(n) = fib(n-1) + fib(n-2)

func fibonacci(n int) int {
	if n <= 1 { // CASOS BASE: fib(0)=0, fib(1)=1
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2) // CASO RECURSIVO
}

// ─────────────────────────────────────────────────────────
// EJEMPLO CLÁSICO 3: Potencia
// ─────────────────────────────────────────────────────────
func potencia(base, exponente int) int {
	if exponente == 0 { // CASO BASE: cualquier número a la 0 = 1
		return 1
	}
	return base * potencia(base, exponente-1)
}

// ─────────────────────────────────────────────────────────
// EJEMPLO CLÁSICO 4: Suma de dígitos
// ─────────────────────────────────────────────────────────
// sumaDigitos(1234) = 1+2+3+4 = 10
func sumaDigitos(n int) int {
	if n < 10 { // CASO BASE: n es un solo dígito
		return n
	}
	return (n % 10) + sumaDigitos(n/10) // último dígito + suma del resto
}

// ─────────────────────────────────────────────────────────
// RECURSIÓN CON ACUMULADOR (más eficiente, evita re-cálculos)
// ─────────────────────────────────────────────────────────
func factorialConAcum(n, acum int) int {
	if n <= 1 {
		return acum
	}
	return factorialConAcum(n-1, n*acum) // pasa el resultado acumulado
}

// ─────────────────────────────────────────────────────────
// VISUALIZAR EL STACK DE LLAMADAS
// ─────────────────────────────────────────────────────────
func factorialVisual(n, nivel int) int {
	espacio := ""
	for i := 0; i < nivel; i++ {
		espacio += "  "
	}
	fmt.Printf("%s→ factorial(%d)\n", espacio, n)

	if n <= 1 {
		fmt.Printf("%s← retorna 1\n", espacio)
		return 1
	}
	resultado := n * factorialVisual(n-1, nivel+1)
	fmt.Printf("%s← retorna %d × ... = %d\n", espacio, n, resultado)
	return resultado
}

// ─────────────────────────────────────────────────────────
// APLICACIÓN A LA TIENDA: Calcular total de carrito anidado
// ─────────────────────────────────────────────────────────
type Nodo struct {
	Nombre   string
	Precio   float64
	SubNodos []Nodo // un producto puede tener sub-componentes
}

// sumarArbol recorre un árbol de productos recursivamente
func sumarArbol(nodo Nodo) float64 {
	total := nodo.Precio
	for _, hijo := range nodo.SubNodos {
		total += sumarArbol(hijo) // CASO RECURSIVO: sumamos cada hijo
	}
	return total
}

func mostrarArbol(nodo Nodo, nivel int) {
	prefijo := strings.Repeat("  ", nivel) // repite "  " nivel veces (indentación)
	fmt.Printf("%s- %s: $%.2f\n", prefijo, nodo.Nombre, nodo.Precio)
	for _, hijo := range nodo.SubNodos {
		mostrarArbol(hijo, nivel+1)
	}
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║         RECURSIVIDAD          ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// FACTORIAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Factorial ===")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  %2d! = %d\n", i, factorial(i))
	}

	// ─────────────────────────────────────────────────────────
	// FIBONACCI
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Fibonacci (primeros 10) ===")
	fmt.Print("  ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// POTENCIA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Potencias ===")
	fmt.Printf("  2^10 = %d\n", potencia(2, 10))
	fmt.Printf("  3^5  = %d\n", potencia(3, 5))
	fmt.Printf("  10^3 = %d\n", potencia(10, 3))

	// ─────────────────────────────────────────────────────────
	// SUMA DE DÍGITOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Suma de dígitos ===")
	nums := []int{0, 9, 123, 9999, 12345}
	for _, n := range nums {
		fmt.Printf("  sumaDigitos(%d) = %d\n", n, sumaDigitos(n))
	}

	// ─────────────────────────────────────────────────────────
	// VISUALIZAR EL STACK
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Stack de llamadas de factorial(4) ===")
	resultado := factorialVisual(4, 0)
	fmt.Printf("Resultado: %d\n", resultado)

	// ─────────────────────────────────────────────────────────
	// ÁRBOL DE PRODUCTOS (recursión sobre árbol)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Árbol de productos (categorías anidadas) ===")

	arbol := Nodo{
		Nombre: "Computación", Precio: 0,
		SubNodos: []Nodo{
			{
				Nombre: "Portátiles", Precio: 0,
				SubNodos: []Nodo{
					{Nombre: "Notebook Pro", Precio: 1500.00},
					{Nombre: "Notebook Air", Precio: 1200.00},
				},
			},
			{
				Nombre: "Accesorios", Precio: 0,
				SubNodos: []Nodo{
					{Nombre: "Mouse", Precio: 25.99},
					{Nombre: "Teclado", Precio: 75.50},
					{
						Nombre: "Monitores", Precio: 0,
						SubNodos: []Nodo{
							{Nombre: "Monitor 24\"", Precio: 250.00},
							{Nombre: "Monitor 27\"", Precio: 450.00},
						},
					},
				},
			},
		},
	}

	mostrarArbol(arbol, 0)
	fmt.Printf("\nValor total del árbol: $%.2f\n", sumarArbol(arbol))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Toda función recursiva necesita:")
	fmt.Println("  1. CASO BASE:      condición que detiene la recursión")
	fmt.Println("  2. CASO RECURSIVO: llamada a sí misma con problema MÁS PEQUEÑO")
	fmt.Println()
	fmt.Println("Ventajas: código elegante para problemas naturalmente recursivos")
	fmt.Println("          (árboles, estructuras anidadas, divide y conquista)")
	fmt.Println("Desventajas: más lenta que loops, puede causar stack overflow")
	fmt.Println("             con N muy grandes")
}
