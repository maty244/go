package main

import (
	"fmt"
	"math"
)

func main() {
	// =========================================================
	// TIPOS PRIMITIVOS - PARTE 1: NÚMEROS
	// =========================================================
	// Los tipos primitivos son los bloques básicos de construcción.
	// Son los tipos que Go conoce "de fábrica", sin necesidad de
	// importar nada. En esta parte vemos todos los tipos NUMÉRICOS.

	// ─────────────────────────────────────────────────────────
	// ENTEROS CON SIGNO (pueden ser negativos)
	// ─────────────────────────────────────────────────────────
	// Go tiene múltiples tipos de enteros según cuántos bits usan.
	// Más bits = puede guardar números más grandes.
	// Menos bits = ocupa menos memoria.
	//
	// Regla práctica: usa "int" para la mayoría de los casos.
	// Go elegirá automáticamente 32 o 64 bits según tu sistema.

	var i8 int8 = 127                   // 8 bits  → rango: -128 a 127
	var i16 int16 = 32767               // 16 bits → rango: -32768 a 32767
	var i32 int32 = 2147483647          // 32 bits → rango: -2.1B a 2.1B
	var i64 int64 = 9223372036854775807 // 64 bits → enorme
	var i int = 42                      // int → 32 o 64 bits según el sistema (el más común)

	fmt.Println("=== Enteros con signo ===")
	fmt.Printf("int8  (máx): %d\n", i8)
	fmt.Printf("int16 (máx): %d\n", i16)
	fmt.Printf("int32 (máx): %d\n", i32)
	fmt.Printf("int64 (máx): %d\n", i64)
	fmt.Printf("int   (uso común): %d\n", i)

	// ¿Cuándo usar cada uno?
	// int8  → edad, nivel, etc. (valores pequeños conocidos)
	// int16 → año, temperatura en Kelvin
	// int32 → IDs en bases de datos pequeñas
	// int64 → timestamps (nanosegundos), IDs grandes, dinero en centavos
	// int   → la mayoría de los casos

	// ─────────────────────────────────────────────────────────
	// ENTEROS SIN SIGNO (solo positivos, más rango positivo)
	// ─────────────────────────────────────────────────────────
	// El prefijo "u" significa "unsigned" (sin signo).
	// Al no tener negativos, el rango positivo se duplica.

	var u8 uint8 = 255          // 8 bits → rango: 0 a 255
	var u16 uint16 = 65535      // 16 bits → rango: 0 a 65535
	var u32 uint32 = 4294967295 // 32 bits → 0 a ~4 billones
	var u uint = 42             // uint → 32 o 64 bits según sistema

	fmt.Println("\n=== Enteros sin signo (unsigned) ===")
	fmt.Printf("uint8  (máx): %d\n", u8)
	fmt.Printf("uint16 (máx): %d\n", u16)
	fmt.Printf("uint32 (máx): %d\n", u32)
	fmt.Printf("uint   (uso común): %d\n", u)

	// Ejemplo real: colores RGB
	// Cada canal de color (R, G, B) va de 0 a 255 → uint8 es perfecto
	var rojo uint8 = 255
	var verde uint8 = 128
	var azul uint8 = 0
	fmt.Printf("\nColor RGB: rgb(%d, %d, %d)\n", rojo, verde, azul)

	// ─────────────────────────────────────────────────────────
	// DESBORDAMIENTO (OVERFLOW) - ¡PELIGRO!
	// ─────────────────────────────────────────────────────────
	// Si un número supera el límite del tipo, NO da error en runtime,
	// simplemente "da vuelta" al otro extremo. Esto es un bug silencioso.

	var maxInt8 int8 = 127
	// maxInt8++ daría -128 en runtime (overflow silencioso)
	fmt.Println("\n=== Overflow ===")
	fmt.Println("int8 máximo:", maxInt8)
	fmt.Println("Si le sumamos 1 en runtime: da la vuelta a -128 (¡bug silencioso!)")

	// ─────────────────────────────────────────────────────────
	// NÚMEROS DECIMALES (PUNTO FLOTANTE)
	// ─────────────────────────────────────────────────────────
	// Para números con parte decimal Go tiene float32 y float64.
	// float64 es el más usado porque es más preciso.
	// Los literales decimales en Go son float64 por defecto.

	var f32 float32 = 3.14              // 32 bits, ~6-7 dígitos de precisión
	var f64 float64 = 3.141592653589793 // 64 bits, ~15-17 dígitos de precisión

	fmt.Println("\n=== Números decimales (float) ===")
	fmt.Printf("float32: %.10f\n", f32) // nota la pérdida de precisión
	fmt.Printf("float64: %.15f\n", f64) // mucho más preciso

	// Diferencia de precisión (¡importante en sistemas financieros!)
	var a float32 = 0.1
	var b float32 = 0.2
	var c float64 = 0.1
	var d float64 = 0.2

	fmt.Printf("\nfloat32: 0.1 + 0.2 = %.20f\n", a+b) // resultado impreciso
	fmt.Printf("float64: 0.1 + 0.2 = %.20f\n", c+d)   // más preciso, pero tampoco exacto

	// Regla: NUNCA uses float para dinero. Usa enteros (centavos).
	// 1.99 dólares = 199 centavos (int)

	// ─────────────────────────────────────────────────────────
	// NÚMEROS ESPECIALES EN FLOAT
	// ─────────────────────────────────────────────────────────
	positiveInf := math.Inf(1)  // +Infinito
	negativeInf := math.Inf(-1) // -Infinito
	notANumber := math.NaN()    // No es un número (resultado de 0/0, etc.)

	fmt.Println("\n=== Valores especiales float ===")
	fmt.Println("+Infinito:", positiveInf)
	fmt.Println("-Infinito:", negativeInf)
	fmt.Println("NaN:", notANumber)
	fmt.Println("¿Es NaN?:", math.IsNaN(notANumber))

	// ─────────────────────────────────────────────────────────
	// NÚMEROS COMPLEJOS
	// ─────────────────────────────────────────────────────────
	// Go tiene soporte nativo para números complejos (raramente usados
	// excepto en matemática/física/ingeniería).
	// Formato: parte_real + parte_imaginaria*i

	var comp64 complex64 = 3 + 4i   // float32 para cada parte
	var comp128 complex128 = 3 + 4i // float64 para cada parte (más preciso)

	fmt.Println("\n=== Números complejos ===")
	fmt.Printf("complex64:  %v\n", comp64)
	fmt.Printf("complex128: %v\n", comp128)
	fmt.Printf("Parte real: %v, Parte imaginaria: %v\n",
		real(comp128), imag(comp128))

	// Módulo de un número complejo: sqrt(a² + b²)
	modulo := math.Sqrt(real(comp128)*real(comp128) + imag(comp128)*imag(comp128))
	fmt.Printf("Módulo (|3+4i|): %.2f\n", modulo) // 5.00 (triángulo 3-4-5)

	// ─────────────────────────────────────────────────────────
	// OPERACIONES COMUNES CON ENTEROS
	// ─────────────────────────────────────────────────────────
	a2, b2 := 17, 5

	fmt.Println("\n=== Operaciones con enteros ===")
	fmt.Printf("%d + %d = %d\n", a2, b2, a2+b2)
	fmt.Printf("%d - %d = %d\n", a2, b2, a2-b2)
	fmt.Printf("%d * %d = %d\n", a2, b2, a2*b2)
	fmt.Printf("%d / %d = %d  (división entera, descarta decimales)\n", a2, b2, a2/b2)
	fmt.Printf("%d %% %d = %d (resto/módulo: qué sobra)\n", a2, b2, a2%b2)

	// Aplicación real del módulo: saber si un número es par o impar
	numero := 42
	if numero%2 == 0 {
		fmt.Printf("\n%d es par (resto al dividir por 2 es 0)\n", numero)
	} else {
		fmt.Printf("\n%d es impar\n", numero)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE CUÁNDO USAR CADA TIPO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Guía de uso ===")
	fmt.Println("int     → contadores, índices, cantidades generales")
	fmt.Println("int64   → IDs grandes, timestamps, cálculos donde el rango importa")
	fmt.Println("uint8   → bytes, canales de color RGB, datos binarios")
	fmt.Println("float64 → decimales en general, coordenadas, porcentajes")
	fmt.Println("float32 → gráficos 3D donde la memoria importa más que la precisión")
}
