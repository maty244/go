package main

import "fmt"

// =========================================================
// FUNCIONES GENÉRICAS: LA SINTAXIS
// =========================================================
// La sintaxis agrega un "parámetro de tipo" entre corchetes,
// antes de los parámetros normales:
//
//   func Nombre[T TipoPermitido](param T) T { ... }
//
// T es un NOMBRE que elegís (por convención, una letra mayúscula:
// T, K, V...). Representa "el tipo que se use en esta llamada
// específica". Go lo DEDUCE automáticamente de los argumentos,
// no hace falta escribirlo casi nunca.

// ─────────────────────────────────────────────────────────
// EJEMPLO MÍNIMO: any COMO RESTRICCIÓN (sin operaciones)
// ─────────────────────────────────────────────────────────
// "any" es la restricción más permisiva: acepta cualquier tipo,
// pero por eso mismo no podés usar operadores (+, >, etc.) sobre
// el valor, solo cosas que funcionan para CUALQUIER tipo (como
// asignarlo, devolverlo, o compararlo con == si es "comparable").

func Primero[T any](lista []T) T {
	return lista[0]
}

func Ultimo[T any](lista []T) T {
	return lista[len(lista)-1]
}

// ─────────────────────────────────────────────────────────
// RESTRICCIÓN CON OPERADORES: una lista de tipos permitidos
// ─────────────────────────────────────────────────────────
// Si tu función necesita usar +, >, etc., "any" no alcanza (esos
// operadores no existen para CUALQUIER tipo). Tenés que restringir
// a los tipos que SÍ los soportan.

func Sumar[T int | float64](numeros []T) T {
	var total T
	for _, n := range numeros {
		total += n
	}
	return total
}

func main() {
	fmt.Println("=== Funciones genéricas con any ===")

	nombres := []string{"Ana", "Carlos", "Sofía"}
	fmt.Println("Primero:", Primero(nombres))
	fmt.Println("Ultimo:", Ultimo(nombres))

	numeros := []int{10, 20, 30}
	fmt.Println("Primero:", Primero(numeros))
	fmt.Println("Ultimo:", Ultimo(numeros))

	// ─────────────────────────────────────────────────────────
	// GO DEDUCE EL TIPO SOLO (inferencia)
	// ─────────────────────────────────────────────────────────
	// No escribimos Primero[string](nombres): Go mira el argumento
	// y deduce que T=string. Se puede ser explícito si hace falta:

	fmt.Println("\n=== Inferencia de tipo (normalmente no hace falta ser explícito) ===")
	explicito := Primero[string](nombres)
	fmt.Println("Primero[string](nombres):", explicito)

	// ─────────────────────────────────────────────────────────
	// FUNCIÓN CON RESTRICCIÓN NUMÉRICA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Sumar con restricción int | float64 ===")

	enteros := []int{1, 2, 3, 4, 5}
	decimales := []float64{1.5, 2.5, 3.0}

	fmt.Println("Sumar(enteros)  =", Sumar(enteros))
	fmt.Println("Sumar(decimales) =", Sumar(decimales))

	// Sumar([]string{"a", "b"})  ← ERROR de compilación:
	// string no está en la lista de tipos permitidos (int | float64)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  func F[T any](x T) T       → T puede ser CUALQUIER tipo")
	fmt.Println("  func F[T int|float64](...)  → T solo puede ser uno de esos tipos")
	fmt.Println("  Go DEDUCE T                 → normalmente no hace falta escribirlo")
	fmt.Println("  any                         → sin operadores, solo asignar/devolver")
	fmt.Println("  Restricción específica      → habilita los operadores de esos tipos")
}
