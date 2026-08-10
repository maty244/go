package main

import "fmt"

// =========================================================
// ERRORES COMUNES EN GO
// =========================================================
// Go tiene un compilador muy estricto que convierte muchos
// errores de runtime en errores de compilación.
// Esto puede frustrar al principiante, pero te protege de bugs.
// Este archivo muestra los errores más frecuentes y cómo evitarlos.

func main() {
	fmt.Printf("=== Errores comunes en Go y cómo evitarlos ===\n\n")

	// ─────────────────────────────────────────────────────────
	// ERROR 1: Variable declarada pero no usada
	// ─────────────────────────────────────────────────────────
	// Go no compila si declarás una variable y no la usás.
	// Esto evita variables "fantasma" que confunden al lector.
	//
	// var x int = 5 // ERROR: x declared and not used
	//
	// SOLUCIÓN: usá la variable o usá _ para descartarla
	x := 42
	fmt.Println("ERROR 1 - Variable usada:", x)

	// ─────────────────────────────────────────────────────────
	// ERROR 2: Importación no usada
	// ─────────────────────────────────────────────────────────
	// Si importás un paquete y no lo usás, Go no compila.
	// import "os" // ERROR si no usás nada de "os"
	//
	// SOLUCIÓN: importar solo lo que usás, o usar _ para importación
	// con efectos secundarios (muy raro, para plugins/drivers).
	fmt.Println("\nERROR 2 - Importaciones: solo importá lo que usás")

	// ─────────────────────────────────────────────────────────
	// ERROR 3: Tipo incorrecto en una operación
	// ─────────────────────────────────────────────────────────
	// Go NO hace conversiones automáticas entre tipos.
	// var a int = 5
	// var b float64 = 3.14
	// resultado := a + b  // ERROR: mismatched types int and float64
	//
	// SOLUCIÓN: convertir explícitamente
	var a int = 5
	var b float64 = 3.14
	resultado := float64(a) + b // convertimos int a float64
	fmt.Printf("\nERROR 3 - Tipos: float64(%d) + %.2f = %.2f\n", a, b, resultado)

	// ─────────────────────────────────────────────────────────
	// ERROR 4: := a nivel de paquete
	// ─────────────────────────────────────────────────────────
	// := solo funciona DENTRO de funciones.
	// A nivel de paquete se debe usar var.
	// miVar := 10 // ERROR si está fuera de una función
	//
	// SOLUCIÓN: usar var a nivel de paquete
	fmt.Println("\nERROR 4 - := solo va dentro de funciones, ver variableGlobal:", variableGlobal)

	// ─────────────────────────────────────────────────────────
	// ERROR 5: Usar = en vez de == para comparar
	// ─────────────────────────────────────────────────────────
	// En Go (a diferencia de lenguajes más permisivos), el if requiere
	// una expresión bool. Un = dentro de if daría error de compilación.
	// if x = 5 { } // ERROR: cannot use assignment as value
	//
	// SOLUCIÓN: usar == para comparar
	if x == 42 {
		fmt.Println("\nERROR 5 - == para comparar, = para asignar: x es 42")
	}

	// ─────────────────────────────────────────────────────────
	// ERROR 6: Olvidarse del & en fmt.Scan
	// ─────────────────────────────────────────────────────────
	// fmt.Scan necesita la DIRECCIÓN de la variable, no su valor.
	// var n int
	// fmt.Scan(n)  // ERROR: no compila o comportamiento undefined
	// fmt.Scan(&n) // CORRECTO: & da la dirección de memoria
	fmt.Println("\nERROR 6 - Scan necesita &variable (dirección, no valor)")
	fmt.Println("  Mal:  fmt.Scan(n)")
	fmt.Println("  Bien: fmt.Scan(&n)")

	// ─────────────────────────────────────────────────────────
	// ERROR 7: Overflow silencioso
	// ─────────────────────────────────────────────────────────
	// Si un número supera el rango de su tipo, "da vuelta" silenciosamente.
	var maxInt8 int8 = 127
	maxInt8++ // ahora es -128 (overflow), NO da error en runtime
	fmt.Printf("\nERROR 7 - Overflow: int8(127)++ = %d (¡se volvió negativo!)\n", maxInt8)
	fmt.Println("  PREVENCIÓN: elegí el tipo de dato correcto para tu rango")

	// ─────────────────────────────────────────────────────────
	// ERROR 8: División entera (truncamiento)
	// ─────────────────────────────────────────────────────────
	// Cuando dividís dos enteros en Go, el resultado es entero.
	// Los decimales se descartan (NO se redondea).
	cociente := 7 / 2     // resultado: 3, NO 3.5
	modulo := 7 % 2       // resultado: 1 (el resto)
	flotante := 7.0 / 2.0 // resultado: 3.5 (división float)

	fmt.Printf("\nERROR 8 - División entera:\n")
	fmt.Printf("  7 / 2 = %d (entero, trunca)\n", cociente)
	fmt.Printf("  7 %% 2 = %d (resto)\n", modulo)
	fmt.Printf("  7.0 / 2.0 = %f (float, correcto)\n", flotante)

	// ─────────────────────────────────────────────────────────
	// ERROR 9: No manejar errores retornados por funciones
	// ─────────────────────────────────────────────────────────
	// En Go, las funciones retornan errores explícitamente.
	// Ignorar el error puede llevar a bugs silenciosos.

	fmt.Println("\nERROR 9 - Ignorar errores:")

	// MAL: ignorar el error puede causar comportamiento inesperado
	resultado2, _ := dividir(10, 0) // ignoramos el error con _
	fmt.Printf("  10/0 ignorando error = %v\n", resultado2)

	// BIEN: siempre verificar el error
	resultado3, err := dividir(10, 0)
	if err != nil {
		fmt.Printf("  10/0 manejando error: %v\n", err)
	} else {
		fmt.Printf("  Resultado: %v\n", resultado3)
	}

	resultado4, err2 := dividir(10, 2)
	if err2 != nil {
		fmt.Println("  Error:", err2)
	} else {
		fmt.Printf("  10/2 = %v ✓\n", resultado4)
	}

	// ─────────────────────────────────────────────────────────
	// ERROR 10: Modificar un slice mientras se itera
	// ─────────────────────────────────────────────────────────
	// Modificar los índices/longitud de un slice mientras se recorre
	// puede dar resultados inesperados. Lo ideal es crear una copia
	// o recolectar los índices y modificar después.
	fmt.Println("\nERROR 10 - Modificar slice durante iteración:")
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("  Original:", numeros)

	// MAL PATRÓN (descomentar solo para ver el efecto):
	// for i, v := range numeros {
	// 	if v == 3 {
	// 		numeros = append(numeros[:i], numeros[i+1:]...) // eliminar elemento
	// 	}
	// }

	// BIEN: filtrar creando un nuevo slice
	filtrado := filtrar(numeros, func(n int) bool { return n != 3 })
	fmt.Println("  Sin el 3:", filtrado)

	// ─────────────────────────────────────────────────────────
	// ERROR 11: String vs []byte cuando se necesita mutabilidad
	// ─────────────────────────────────────────────────────────
	// Los strings en Go son INMUTABLES. No podés hacer: s[0] = 'H'
	// Si necesitás modificar caracteres, convertí a []byte primero.

	fmt.Println("\nERROR 11 - Strings son inmutables:")
	s := "hola"
	// s[0] = 'H' // ERROR: cannot assign to s[0] (strings are immutable)

	// SOLUCIÓN: convertir a []byte, modificar, y volver a string
	bytes := []byte(s)
	bytes[0] = 'H'
	sModificado := string(bytes)
	fmt.Printf("  '%s' → '%s' (via []byte)\n", s, sModificado)

	// ─────────────────────────────────────────────────────────
	// ERROR 12: Comparar structs que contienen slices/maps
	// ─────────────────────────────────────────────────────────
	// En Go podés comparar structs con == solo si todos sus campos
	// son comparables. Los slices y maps NO son comparables con ==.

	fmt.Println("\nERROR 12 - Structs con slices no son comparables con ==:")

	type Persona struct {
		Nombre string
		Edad   int
	}
	p1 := Persona{"Ana", 25}
	p2 := Persona{"Ana", 25}
	fmt.Printf("  p1 == p2: %v (struct sin slices: OK)\n", p1 == p2)

	// Para comparar slices, usar reflect.DeepEqual o comparar manualmente
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	// s1 == s2 // ERROR de compilación
	fmt.Printf("  Slices iguales manualmente: %v\n", slicesIguales(s1, s2))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen de errores frecuentes ===")
	errores := []string{
		"1. Variable no usada → usarla o descartarla con _",
		"2. Import no usado → eliminar la importación",
		"3. Tipos incompatibles → convertir explícitamente",
		"4. := fuera de función → usar var",
		"5. = en lugar de == para comparar",
		"6. Olvidar & en fmt.Scan",
		"7. Overflow silencioso → elegir el tipo correcto",
		"8. División entera trunca → usar floats si necesitás decimales",
		"9. No manejar errores → siempre verificar err != nil",
		"10. Modificar slice en iteración → crear copia",
		"11. Intentar mutar string → convertir a []byte",
		"12. Comparar structs con slices/maps con == → comparación manual",
	}
	for _, e := range errores {
		fmt.Println(" ", e)
	}
}

// Variable global para demostrar error 4
var variableGlobal int = 100

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: división por cero")
	}
	return a / b, nil
}

func filtrar(slice []int, fn func(int) bool) []int {
	resultado := []int{}
	for _, v := range slice {
		if fn(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func slicesIguales(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
