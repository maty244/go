package main

import "fmt"

// =========================================================
// FUNCIONES COMO VALORES + FUNCIONES ANÓNIMAS + CLOSURES
// =========================================================
// En Go, las funciones son "ciudadanos de primera clase".
// Esto significa que las funciones SON valores: podés:
//   - Asignarlas a variables
//   - Pasarlas como parámetros
//   - Retornarlas desde otras funciones
//   - Almacenarlas en slices y maps
//
// FUNCIÓN ANÓNIMA: función sin nombre, declarada inline.
// CLOSURE: función que "captura" variables del scope exterior.

// ─────────────────────────────────────────────────────────
// TIPOS DE FUNCIÓN
// ─────────────────────────────────────────────────────────
// Podemos crear un alias de tipo para un tipo de función.
type OperacionPrecio func(float64) float64

type Producto struct {
	Nombre string
	Precio float64
}

// ─────────────────────────────────────────────────────────
// FUNCIONES COMO VALORES
// ─────────────────────────────────────────────────────────

func aplicarIVA(p float64) float64         { return p * 1.21 }
func aplicarDescuento10(p float64) float64 { return p * 0.90 }
func aplicarDescuento20(p float64) float64 { return p * 0.80 }
func sinModificacion(p float64) float64    { return p }

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║   FUNCIONES COMO VALORES      ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// 1. ASIGNAR FUNCIÓN A VARIABLE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Función asignada a variable ===")

	// "calcular" es una variable que contiene una función
	var calcular OperacionPrecio = aplicarIVA
	fmt.Printf("Precio con IVA: $%.2f\n", calcular(100.0))

	// Reasignar la variable a otra función del mismo tipo
	calcular = aplicarDescuento10
	fmt.Printf("Precio con descuento 10%%: $%.2f\n", calcular(100.0))

	// ─────────────────────────────────────────────────────────
	// 2. SLICE DE FUNCIONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slice de funciones ===")

	// Definimos un pipeline de transformaciones
	pipeline := []OperacionPrecio{
		aplicarIVA,
		aplicarDescuento10,
	}

	precio := 100.0
	for _, fn := range pipeline {
		precio = fn(precio)
	}
	fmt.Printf("Precio después del pipeline: $%.2f\n", precio)

	// ─────────────────────────────────────────────────────────
	// 3. MAP DE FUNCIONES (selector de estrategia)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Map de funciones ===")

	estrategias := map[string]OperacionPrecio{
		"normal":  sinModificacion,
		"vip10":   aplicarDescuento10,
		"vip20":   aplicarDescuento20,
		"con_iva": aplicarIVA,
	}

	clientes := map[string]string{
		"Ana":    "vip20",
		"Carlos": "normal",
		"Mia":    "vip10",
		"Admin":  "con_iva",
	}

	precioBase := 500.0
	for cliente, tipo := range clientes {
		fn := estrategias[tipo]
		fmt.Printf("  %-8s (%-8s): $%.2f\n", cliente, tipo, fn(precioBase))
	}

	// ─────────────────────────────────────────────────────────
	// 4. FUNCIONES ANÓNIMAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Funciones anónimas ===")

	// Asignar una función anónima a una variable
	cuadrado := func(n float64) float64 {
		return n * n
	}
	fmt.Printf("Cuadrado de 5: %.0f\n", cuadrado(5))

	// IIFE (Immediately Invoked Function Expression)
	// Función anónima que se llama inmediatamente
	resultado := func(a, b float64) float64 {
		return a*a + b*b
	}(3, 4) // se llama aquí mismo con (3, 4)
	fmt.Printf("3² + 4² = %.0f\n", resultado)

	// Función anónima como argumento inline
	productos := []Producto{
		{"Notebook", 1500.00},
		{"Mouse", 25.99},
		{"Teclado", 75.50},
	}

	fmt.Println("\nProductos con IVA (función anónima inline):")
	aplicarATodos(productos, func(p Producto) Producto {
		p.Precio *= 1.21
		return p
	})

	// ─────────────────────────────────────────────────────────
	// 5. CLOSURES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Closures ===")
	// Un closure es una función anónima que "recuerda" variables
	// del entorno donde fue creada.

	// Contador con closure (estado encapsulado)
	crearContador := func(inicio int) func() int {
		n := inicio // "n" es capturada por el closure
		return func() int {
			n++
			return n
		}
	}

	contarPedidos := crearContador(1000)
	fmt.Printf("Pedido #%d\n", contarPedidos())
	fmt.Printf("Pedido #%d\n", contarPedidos())
	fmt.Printf("Pedido #%d\n", contarPedidos())

	contarFacturas := crearContador(5000) // contador independiente
	fmt.Printf("Factura #%d\n", contarFacturas())
	fmt.Printf("Factura #%d\n", contarFacturas())

	// Closure que captura la tasa de impuesto
	crearCalculadoraImpuesto := func(tasa float64) func(float64) float64 {
		return func(precio float64) float64 {
			return precio * (1 + tasa)
		}
	}

	conIVA := crearCalculadoraImpuesto(0.21)
	conIIBB := crearCalculadoraImpuesto(0.035)

	fmt.Printf("\n$100 con IVA:  $%.2f\n", conIVA(100))
	fmt.Printf("$100 con IIBB: $%.2f\n", conIIBB(100))

	// ─────────────────────────────────────────────────────────
	// 6. FUNCIONES QUE RETORNAN FUNCIONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Función que retorna función ===")

	// "Fábrica" de funciones de descuento
	hacerDescuento := func(porcentaje float64) OperacionPrecio {
		return func(precio float64) float64 {
			return precio * (1 - porcentaje)
		}
	}

	desc5 := hacerDescuento(0.05)
	desc15 := hacerDescuento(0.15)
	desc25 := hacerDescuento(0.25)

	fmt.Printf("$200 con 5%% desc:  $%.2f\n", desc5(200))
	fmt.Printf("$200 con 15%% desc: $%.2f\n", desc15(200))
	fmt.Printf("$200 con 25%% desc: $%.2f\n", desc25(200))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("Función como valor:  var f func(float64) float64 = miFunc")
	fmt.Println("Función anónima:     f := func(x int) int { return x * 2 }")
	fmt.Println("IIFE:                resultado := func(x int) int { ... }(5)")
	fmt.Println("Closure:             captura variables del scope donde se creó")
	fmt.Println("Slice de funciones:  []func(float64) float64{fn1, fn2, fn3}")
	fmt.Println("Map de funciones:    map[string]func(float64) float64{\"iva\": aplicarIVA}")
}

// Función que recibe una función como parámetro
func aplicarATodos(ps []Producto, transform func(Producto) Producto) {
	for _, p := range ps {
		transformado := transform(p)
		fmt.Printf("  %-15s → $%.2f\n", transformado.Nombre, transformado.Precio)
	}
}
