package main

import "fmt"

// =========================================================
// FUNCIONES QUE DEVUELVEN FUNCIONES CONFIGURADAS
// =========================================================
// Un uso muy práctico de los closures: fabricar funciones YA
// CONFIGURADAS con ciertos parámetros fijos, para no tener que
// repetirlos cada vez que las usás. Es una forma de "precargar"
// argumentos.
//
// Este patrón se llama a veces "fábrica de funciones" o, en
// términos más técnicos, "aplicación parcial".

// ─────────────────────────────────────────────────────────
// EJEMPLO: calculadora de descuentos preconfigurada
// ─────────────────────────────────────────────────────────
// En vez de llamar siempre aplicarDescuento(precio, 0.10), armamos
// UNA función ya configurada con el 10%, lista para reusar.

func crearDescuento(porcentaje float64) func(precio float64) float64 {
	return func(precio float64) float64 {
		return precio * (1 - porcentaje)
	}
}

func main() {
	fmt.Println("=== Fábrica de descuentos ===")

	descuento10 := crearDescuento(0.10)
	descuentoVIP := crearDescuento(0.25)

	precio := 1000.0
	fmt.Printf("Precio original:      $%.2f\n", precio)
	fmt.Printf("Con 10%% de descuento: $%.2f\n", descuento10(precio))
	fmt.Printf("Con 25%% VIP:          $%.2f\n", descuentoVIP(precio))

	// La misma función preconfigurada sirve para cualquier precio,
	// sin repetir el porcentaje cada vez.
	fmt.Println("\nAplicando descuento10 a varios precios:")
	for _, p := range []float64{500, 1200, 3000} {
		fmt.Printf("  $%.2f → $%.2f\n", p, descuento10(p))
	}

	// ─────────────────────────────────────────────────────────
	// DECORADORES: envolver una función con comportamiento extra
	// ─────────────────────────────────────────────────────────
	// Un closure puede RECIBIR una función y devolver una VERSIÓN
	// mejorada de ella, que hace algo antes y/o después de llamarla.
	// Esto es la base de los "middlewares" que vas a ver en HTTP.

	fmt.Println("\n=== Decorador: loguear entrada y salida de una función ===")

	sumaLenta := func(a, b int) int {
		return a + b
	}

	sumaConLog := conLog("suma", sumaLenta)
	resultado := sumaConLog(4, 5)
	fmt.Println("Resultado:", resultado)

	// ─────────────────────────────────────────────────────────
	// RATE LIMITER SIMPLE: un closure que "recuerda" cuántas veces
	// se lo llamó
	// ─────────────────────────────────────────────────────────
	// Otro caso clásico: limitar cuántas veces se puede ejecutar
	// algo. Acá no medimos tiempo real (eso lo vemos en 09_biblioteca_estandar/04_tiempo/),
	// pero mostramos la idea con un contador simple.

	fmt.Println("\n=== Rate limiter simple (por cantidad de intentos) ===")

	intentarLogin := crearLimitador(3)

	for i := 1; i <= 5; i++ {
		permitido := intentarLogin()
		fmt.Printf("  Intento %d: permitido=%v\n", i, permitido)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: validadores configurables
	// ─────────────────────────────────────────────────────────
	// Un formulario de registro necesita validar que un campo
	// tenga cierta longitud mínima. En vez de escribir una función
	// por cada mínimo posible, fabricamos la validación.

	fmt.Println("\n=== Caso real: validador de longitud configurable ===")

	validarUsername := longitudMinima(4)
	validarPassword := longitudMinima(8)

	fmt.Println(`validarUsername("ana"):`, validarUsername("ana"))
	fmt.Println(`validarUsername("carlos"):`, validarUsername("carlos"))
	fmt.Println(`validarPassword("1234"):`, validarPassword("1234"))
	fmt.Println(`validarPassword("unaClaveSegura"):`, validarPassword("unaClaveSegura"))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  func(config) func(...)  → fábrica: precarga argumentos fijos")
	fmt.Println("  func(f) func(...)       → decorador: envuelve una función con extras")
	fmt.Println("  El estado (contador,    → vive en el closure, no en una global")
	fmt.Println("  configuración) queda")
	fmt.Println("  Base de                 → middlewares (HTTP/) y opciones (config)")
}

// conLog decora cualquier función suma(int,int)int agregando un log
// antes y después, sin modificar la función original.
func conLog(nombre string, f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Printf("  [%s] llamando con a=%d, b=%d\n", nombre, a, b)
		resultado := f(a, b)
		fmt.Printf("  [%s] resultado=%d\n", nombre, resultado)
		return resultado
	}
}

func crearLimitador(maximo int) func() bool {
	intentos := 0
	return func() bool {
		intentos++
		return intentos <= maximo
	}
}

// Usamos "minimo" en lugar de "min" para no sombrear la función predeclarada min() de Go 1.21+.
// Nota: len(s) mide bytes (para contar caracteres Unicode/runas usarías utf8.RuneCountInString(s)).
func longitudMinima(minimo int) func(string) bool {
	return func(s string) bool {
		return len(s) >= minimo
	}
}
