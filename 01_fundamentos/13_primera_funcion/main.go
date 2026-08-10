package main

import "fmt"

// =========================================================
// PRIMERA FUNCIÓN EN GO
// =========================================================
// Una función es un bloque de código con nombre que podés
// ejecutar ("llamar") desde cualquier parte del programa.
//
// ¿Para qué sirven?
//   - Evitar repetir el mismo código (principio DRY: Don't Repeat Yourself)
//   - Organizar el código en piezas pequeñas y comprensibles
//   - Facilitar el testing y el mantenimiento
//
// Anatomía de una función en Go:
//
//   func nombreFuncion(param1 tipo, param2 tipo) tipoRetorno {
//       // cuerpo de la función
//       return valor
//   }
//
//   func     → palabra clave para declarar una función
//   nombre   → identificador único (usa camelCase en Go)
//   params   → lista de entrada (puede ser vacía)
//   retorno  → tipo del valor que devuelve (puede omitirse si no retorna nada)
//   return   → devuelve el control (y opcionalmente un valor) al llamador

// ─────────────────────────────────────────────────────────
// Nuestra primera función: sin parámetros, sin retorno
// ─────────────────────────────────────────────────────────
// Esto se llama función "void" en otros lenguajes.
// En Go simplemente no declaramos tipo de retorno.
func saludar() {
	fmt.Println("¡Bienvenido a la Tienda Go!")
	fmt.Println("Tu lugar de confianza para comprar tecnología.")
}

// ─────────────────────────────────────────────────────────
// Función con un parámetro
// ─────────────────────────────────────────────────────────
// "nombre" es el parámetro: una variable local que recibe
// el valor que pasamos al llamar la función.
func saludarCliente(nombre string) {
	fmt.Printf("¡Hola, %s! Bienvenido a la Tienda Go.\n", nombre)
}

// ─────────────────────────────────────────────────────────
// Función con parámetro y retorno
// ─────────────────────────────────────────────────────────
// Esta función toma un precio y le aplica el IVA (21%).
// Retorna el precio final como float64.
func calcularPrecioConIVA(precioBase float64) float64 {
	const IVA = 0.21
	return precioBase * (1 + IVA)
}

// ─────────────────────────────────────────────────────────
// Función con múltiples parámetros
// ─────────────────────────────────────────────────────────
func mostrarProducto(nombre string, precio float64, stock int) {
	fmt.Printf("  📦 %-15s | $%8.2f | Stock: %d\n", nombre, precio, stock)
}

func main() {
	fmt.Printf("=== Mi primera función en Go ===\n\n")

	// Llamar una función: escribís el nombre seguido de ()
	saludar()

	fmt.Println()

	// Llamar con argumento
	saludarCliente("Matias")
	saludarCliente("Ana")
	saludarCliente("Carlos")

	fmt.Println()

	// Llamar y usar el valor retornado
	precioBase := 100.0
	precioFinal := calcularPrecioConIVA(precioBase)
	fmt.Printf("Precio base: $%.2f\n", precioBase)
	fmt.Printf("Con IVA:     $%.2f\n", precioFinal)

	fmt.Println("\nCatálogo de productos:")
	fmt.Println("  Producto        | Precio   | Stock")
	fmt.Println("  ----------------|----------|------")
	mostrarProducto("Notebook", 1500.00, 5)
	mostrarProducto("Mouse", 25.99, 42)
	mostrarProducto("Teclado", 75.50, 18)
	mostrarProducto("Monitor", 450.00, 3)

	// ─────────────────────────────────────────────────────────
	// PUNTOS CLAVE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Lo que aprendimos ===")
	fmt.Println("1. func nombre()               → función sin params ni retorno")
	fmt.Println("2. func nombre(p tipo)         → función con parámetro")
	fmt.Println("3. func nombre(p tipo) tipo    → función con parámetro y retorno")
	fmt.Println("4. nombreFuncion(argumento)    → llamar una función")
	fmt.Println("5. variable := nombreFuncion() → capturar el valor retornado")
}
