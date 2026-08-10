package main

import "fmt"

// =========================================================
// FUNCIONES ANÓNIMAS
// =========================================================
// Ya viste en 05_funciones_y_closures/05_funciones_como_valores (funciones como valores) que una
// función se puede guardar en una variable. Una función ANÓNIMA
// es una función SIN NOMBRE, definida justo donde se usa.
//
// Sirven para: código que se usa una sola vez y no amerita un
// nombre propio, o para armar closures (el tema central de esta
// sección, que vemos a partir del archivo 02).

func main() {
	// ─────────────────────────────────────────────────────────
	// FUNCIÓN ANÓNIMA GUARDADA EN UNA VARIABLE
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Función anónima en una variable ===")

	duplicar := func(n int) int {
		return n * 2
	}

	fmt.Println("duplicar(21) =", duplicar(21))

	// ─────────────────────────────────────────────────────────
	// FUNCIÓN ANÓNIMA INVOCADA INMEDIATAMENTE (IIFE)
	// ─────────────────────────────────────────────────────────
	// A veces querés ejecutar código una sola vez, en un bloque
	// aislado (por ejemplo, para no "ensuciar" main con variables
	// temporales). Se define la función y se la llama al toque
	// con () al final.

	fmt.Println("\n=== Función anónima invocada al instante (IIFE) ===")

	resultado := func(a, b int) int {
		suma := a + b
		return suma * suma
	}(3, 4) // ← se llama inmediatamente con estos argumentos

	fmt.Println("(3+4)² =", resultado)

	// ─────────────────────────────────────────────────────────
	// FUNCIÓN ANÓNIMA COMO ARGUMENTO
	// ─────────────────────────────────────────────────────────
	// Ya lo viste en 05_funciones_y_closures/06_funciones_como_parametros (funciones como parámetros).
	// Es MUY común pasar una función anónima directo, sin
	// declararla antes, cuando solo se usa una vez.

	fmt.Println("\n=== Función anónima como argumento ===")

	numeros := []int{1, 2, 3, 4, 5}
	pares := filtrar(numeros, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Pares:", pares)

	// ─────────────────────────────────────────────────────────
	// defer CON FUNCIÓN ANÓNIMA
	// ─────────────────────────────────────────────────────────
	// Ya usaste defer en 05_funciones_y_closures/04_defer. Combinarlo con una función
	// anónima permite ejecutar VARIAS líneas al salir de la función,
	// no solo una llamada simple.

	fmt.Println("\n=== defer con función anónima ===")
	procesarConLimpieza()

	// ─────────────────────────────────────────────────────────
	// CASO REAL: validaciones puntuales
	// ─────────────────────────────────────────────────────────
	// Un checkout que valida una lista de reglas puntuales, cada
	// una definida ahí mismo porque no se reutiliza en otro lado.

	fmt.Println("\n=== Caso real: validar un carrito ===")
	total := 15000.0
	items := 3

	reglas := []func() error{
		func() error {
			if total <= 0 {
				return fmt.Errorf("el total debe ser mayor a 0")
			}
			return nil
		},
		func() error {
			if items == 0 {
				return fmt.Errorf("el carrito no puede estar vacío")
			}
			return nil
		},
	}

	for _, regla := range reglas {
		if err := regla(); err != nil {
			fmt.Println("  Regla incumplida:", err)
		}
	}
	fmt.Println("  Carrito válido, se puede procesar el pago")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  func(...) {...}         → función sin nombre")
	fmt.Println("  func(...) {...}()       → IIFE: se define y se llama en el momento")
	fmt.Println("  Se puede guardar        → en una variable, pasar como argumento,")
	fmt.Println("                             o usar directo en defer/go/etc.")
	fmt.Println("  Usalas cuando           → el código se usa una sola vez")
}

func filtrar(nums []int, cumple func(int) bool) []int {
	var resultado []int
	for _, n := range nums {
		if cumple(n) {
			resultado = append(resultado, n)
		}
	}
	return resultado
}

func procesarConLimpieza() {
	fmt.Println("  Abriendo recurso...")

	defer func() {
		fmt.Println("  Cerrando recurso...")
		fmt.Println("  Notificando que se liberó la memoria")
	}()

	fmt.Println("  Procesando...")
}
