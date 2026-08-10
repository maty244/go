package main

import "fmt"

func main() {
	// =========================================================
	// VARIABLES EN GO
	// =========================================================
	// Una variable es un espacio en memoria con un nombre,
	// donde guardamos un valor que puede cambiar en el tiempo.
	// En Go, toda variable tiene un TIPO definido (int, string, bool, etc.)
	// y ese tipo NO puede cambiar una vez declarado.

	// ─────────────────────────────────────────────────────────
	// FORMA 1: Declaración con "var" + tipo explícito
	// ─────────────────────────────────────────────────────────
	// Sintaxis: var nombreVariable tipo = valor
	var nombre string = "Matias"
	var edad int = 25
	var altura float64 = 1.75

	fmt.Println("=== Forma 1: var con tipo explícito ===")
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)

	// ─────────────────────────────────────────────────────────
	// FORMA 2: var con inferencia de tipo
	// ─────────────────────────────────────────────────────────
	// Go puede deducir el tipo del valor que le asignamos.
	// No hace falta escribir el tipo, Go lo adivina solo.
	var ciudad = "Buenos Aires" // Go infiere que es string
	var poblacion = 3_000_000   // Go infiere que es int (el _ es separador visual)

	fmt.Println("\n=== Forma 2: var con inferencia de tipo ===")
	fmt.Println("Ciudad:", ciudad)
	fmt.Println("Población:", poblacion)

	// ─────────────────────────────────────────────────────────
	// FORMA 3: Declaración corta con := (short variable declaration)
	// ─────────────────────────────────────────────────────────
	// La forma más común dentro de funciones.
	// SOLO funciona dentro de una función, no a nivel de paquete.
	// Sintaxis: nombreVariable := valor
	pais := "Argentina"
	temperatura := 23.5
	estaLloviendo := false

	fmt.Println("\n=== Forma 3: Declaración corta := ===")
	fmt.Println("País:", pais)
	fmt.Println("Temperatura:", temperatura)
	fmt.Println("¿Está lloviendo?", estaLloviendo)

	// ─────────────────────────────────────────────────────────
	// REASIGNACIÓN DE VARIABLES
	// ─────────────────────────────────────────────────────────
	// Una vez declarada, usamos = para cambiar su valor.
	// NO usamos := de nuevo (eso crearía una variable nueva).
	edad = 26 // cumpleaños! cambiamos el valor
	fmt.Println("\n=== Reasignación ===")
	fmt.Println("Nueva edad:", edad)

	// ─────────────────────────────────────────────────────────
	// DECLARACIÓN EN BLOQUE con var()
	// ─────────────────────────────────────────────────────────
	// Útil cuando queremos declarar varias variables juntas,
	// hace el código más limpio y legible.
	var (
		producto   string  = "Notebook"
		precio     float64 = 1500.99
		stock      int     = 42
		disponible bool    = true
	)

	fmt.Println("\n=== Declaración en bloque var() ===")
	fmt.Printf("Producto: %s | Precio: $%.2f | Stock: %d | Disponible: %v\n",
		producto, precio, stock, disponible)

	// ─────────────────────────────────────────────────────────
	// MÚLTIPLES VARIABLES EN UNA LÍNEA
	// ─────────────────────────────────────────────────────────
	// Podemos declarar y asignar varias variables a la vez.
	x, y, z := 10, 20, 30

	fmt.Println("\n=== Múltiples variables en una línea ===")
	fmt.Println("x:", x, "| y:", y, "| z:", z)

	// Truco clásico: intercambiar valores sin variable temporal
	x, y = y, x
	fmt.Println("Después de intercambiar x e y:")
	fmt.Println("x:", x, "| y:", y)

	// ─────────────────────────────────────────────────────────
	// VARIABLE BLANK (identificador vacío "_")
	// ─────────────────────────────────────────────────────────
	// En Go, si declaras una variable y no la usas, el compilador
	// da error. Pero a veces una función retorna múltiples valores
	// y no necesitamos todos. Para eso usamos "_" (blank identifier).
	// "_" descarta el valor, Go no lo guarda ni lo verifica.
	valorUtil, _ := obtenerDatos() // ignoramos el segundo valor
	fmt.Println("\n=== Blank identifier _ ===")
	fmt.Println("Solo usamos el primer valor:", valorUtil)

	// ─────────────────────────────────────────────────────────
	// SCOPE (ÁMBITO) DE LAS VARIABLES
	// ─────────────────────────────────────────────────────────
	// Las variables solo existen dentro del bloque {} donde se crean.
	{
		dentroDelBloque := "Solo existo aquí adentro"
		fmt.Println("\n=== Scope ===")
		fmt.Println(dentroDelBloque)
	}
	// Si intentáramos usar dentroDelBloque aquí, el compilador daría error.
	// fmt.Println(dentroDelBloque) // ERROR: undefined: dentroDelBloque

	// ─────────────────────────────────────────────────────────
	// VERIFICAR EL TIPO DE UNA VARIABLE
	// ─────────────────────────────────────────────────────────
	// %T en fmt.Printf nos dice el tipo de una variable.
	numero := 42
	decimal := 3.14
	texto := "hola"
	verdadero := true

	fmt.Println("\n=== Tipos de variables ===")
	fmt.Printf("numero = %v  → tipo: %T\n", numero, numero)
	fmt.Printf("decimal = %v  → tipo: %T\n", decimal, decimal)
	fmt.Printf("texto = %v  → tipo: %T\n", texto, texto)
	fmt.Printf("verdadero = %v  → tipo: %T\n", verdadero, verdadero)
}

// Función auxiliar que retorna dos valores (útil para demo del blank identifier)
func obtenerDatos() (string, error) {
	return "dato importante", nil
}
