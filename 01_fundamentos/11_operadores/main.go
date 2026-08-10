package main

import (
	"fmt"
	"math"
)

func main() {
	// =========================================================
	// OPERADORES EN GO
	// =========================================================
	// Los operadores son símbolos que indican una operación
	// entre uno o más valores. Go tiene 4 categorías principales:
	// 1. Aritméticos, 2. Comparación, 3. Lógicos, 4. Bits

	// ─────────────────────────────────────────────────────────
	// 1. OPERADORES ARITMÉTICOS
	// ─────────────────────────────────────────────────────────

	a, b := 17, 5

	fmt.Println("=== Operadores Aritméticos ===")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)  // suma
	fmt.Printf("%d - %d = %d\n", a, b, a-b)  // resta
	fmt.Printf("%d * %d = %d\n", a, b, a*b)  // multiplicación
	fmt.Printf("%d / %d = %d\n", a, b, a/b)  // división ENTERA (descarta decimales)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b) // módulo (resto de la división)

	// División flotante
	fa, fb := float64(a), float64(b)
	fmt.Printf("%.1f / %.1f = %.4f (división flotante)\n", fa, fb, fa/fb)

	// ─────────────────────────────────────────────────────────
	// OPERADORES DE ASIGNACIÓN COMBINADOS
	// ─────────────────────────────────────────────────────────
	// Son atajos: x += 5 es lo mismo que x = x + 5
	x := 10
	fmt.Println("\n=== Asignación combinada (x inicial = 10) ===")
	x += 5
	fmt.Println("x += 5  →", x) // 15
	x -= 3
	fmt.Println("x -= 3  →", x) // 12
	x *= 2
	fmt.Println("x *= 2  →", x) // 24
	x /= 4
	fmt.Println("x /= 4  →", x) // 6
	x %= 4
	fmt.Println("x %= 4  →", x) // 2

	// ─────────────────────────────────────────────────────────
	// INCREMENTO Y DECREMENTO
	// ─────────────────────────────────────────────────────────
	// En Go, ++ y -- son SENTENCIAS, no expresiones.
	// Es decir: y = x++ NO existe en Go (a diferencia de C/Java).
	// Solo se puede usar solo: x++ o x-- en su propia línea.
	contador := 0
	fmt.Println("\n=== Incremento y decremento ===")
	contador++
	fmt.Println("contador++:", contador) // 1
	contador++
	fmt.Println("contador++:", contador) // 2
	contador--
	fmt.Println("contador--:", contador) // 1

	// ─────────────────────────────────────────────────────────
	// OPERADORES DE COMPARACIÓN
	// ─────────────────────────────────────────────────────────
	// Siempre retornan un bool (true o false)
	p, q := 10, 20
	fmt.Println("\n=== Operadores de Comparación ===")
	fmt.Printf("%d == %d → %v  (igual a)\n", p, q, p == q)
	fmt.Printf("%d != %d → %v  (distinto de)\n", p, q, p != q)
	fmt.Printf("%d <  %d → %v  (menor que)\n", p, q, p < q)
	fmt.Printf("%d >  %d → %v  (mayor que)\n", p, q, p > q)
	fmt.Printf("%d <= %d → %v  (menor o igual)\n", p, q, p <= q)
	fmt.Printf("%d >= %d → %v  (mayor o igual)\n", p, q, p >= q)

	// Comparación de strings (lexicográfica)
	s1, s2 := "apple", "banana"
	fmt.Println("\n=== Comparación de strings ===")
	fmt.Printf("'%s' == '%s' → %v\n", s1, s2, s1 == s2)
	fmt.Printf("'%s' <  '%s' → %v (orden lexicográfico)\n", s1, s2, s1 < s2)
	fmt.Printf("'%s' != '%s' → %v\n", s1, s2, s1 != s2)

	// ─────────────────────────────────────────────────────────
	// OPERADORES LÓGICOS
	// ─────────────────────────────────────────────────────────
	// Trabajan con booleanos y retornan booleanos.
	edad := 25
	tieneCarnet := true
	tieneMultas := false

	fmt.Println("\n=== Operadores Lógicos ===")

	// && (AND): ambos deben ser true
	puedeConducir := edad >= 18 && tieneCarnet
	fmt.Printf("¿Puede conducir? edad>=18(%v) && tieneCarnet(%v) → %v\n",
		edad >= 18, tieneCarnet, puedeConducir)

	// || (OR): al menos uno debe ser true
	necesitaRevision := tieneMultas || edad < 18
	fmt.Printf("¿Necesita revisión? tieneMultas(%v) || edad<18(%v) → %v\n",
		tieneMultas, edad < 18, necesitaRevision)

	// ! (NOT): invierte el booleano
	fmt.Printf("!tieneMultas → %v\n", !tieneMultas)
	fmt.Printf("!tieneCarnet → %v\n", !tieneCarnet)

	// Combinaciones complejas
	fmt.Println("\n=== Combinaciones lógicas ===")
	usuario := "admin"
	contraseñaOk := true
	esBloqueado := false

	// Puede acceder si: (es admin O contraseña ok) Y NO está bloqueado
	puedeAcceder := (usuario == "admin" || contraseñaOk) && !esBloqueado
	fmt.Printf("¿Puede acceder? %v\n", puedeAcceder)

	// ─────────────────────────────────────────────────────────
	// SHORT-CIRCUIT EVALUATION (evaluación en cortocircuito)
	// ─────────────────────────────────────────────────────────
	// Go (como la mayoría de lenguajes) evalúa de izquierda a derecha
	// y SE DETIENE cuando el resultado ya está determinado.
	// Con &&: si el primero es false, no evalúa el segundo.
	// Con ||: si el primero es true, no evalúa el segundo.
	// Esto es IMPORTANTE cuando el segundo operando tiene efectos secundarios.

	fmt.Println("\n=== Short-circuit evaluation ===")
	valor := 0
	// Si valor == 0 es true, && se detiene y NO llama a esPrimo(valor)
	// porque false && cualquier_cosa = false
	if valor != 0 && esPrimo(valor) {
		fmt.Println("Es primo")
	} else {
		fmt.Println("No evaluó esPrimo porque valor == 0 fue suficiente")
	}

	// ─────────────────────────────────────────────────────────
	// OPERADORES DE BITS (Bitwise)
	// ─────────────────────────────────────────────────────────
	// Trabajan a nivel de bits (0s y 1s). Muy usados en:
	// - Sistemas de permisos
	// - Optimización de código
	// - Protocolos de red
	// - Flags de configuración

	c, d := 0b1010, 0b1100 // 10 y 12 en binario (prefijo 0b)
	fmt.Println("\n=== Operadores de Bits ===")
	fmt.Printf("a = %04b (%d)\n", c, c)
	fmt.Printf("b = %04b (%d)\n", d, d)
	fmt.Printf("a &  b = %04b (%d)  AND bit a bit\n", c&d, c&d)
	fmt.Printf("a |  b = %04b (%d)  OR bit a bit\n", c|d, c|d)
	fmt.Printf("a ^  b = %04b (%d)  XOR bit a bit\n", c^d, c^d)
	fmt.Printf("^a     = %b  NOT (complemento a 1)\n", ^c)
	fmt.Printf("a << 1 = %04b (%d)  shift left (multiplica por 2)\n", c<<1, c<<1)
	fmt.Printf("a >> 1 = %04b (%d)  shift right (divide por 2)\n", c>>1, c>>1)

	// Aplicación real: sistema de permisos con bits
	// Cada bit representa un permiso diferente
	const (
		LEER     = 1 << 0 // bit 0 → 001 → 1
		ESCRIBIR = 1 << 1 // bit 1 → 010 → 2
		EJECUTAR = 1 << 2 // bit 2 → 100 → 4
	)

	permisoUsuario := LEER | ESCRIBIR // 001 | 010 = 011 = 3
	fmt.Println("\n=== Sistema de permisos con bits ===")
	fmt.Printf("permisos = %b (%d)\n", permisoUsuario, permisoUsuario)
	fmt.Printf("¿Puede leer?    %v\n", permisoUsuario&LEER != 0)
	fmt.Printf("¿Puede escribir? %v\n", permisoUsuario&ESCRIBIR != 0)
	fmt.Printf("¿Puede ejecutar? %v\n", permisoUsuario&EJECUTAR != 0)

	// ─────────────────────────────────────────────────────────
	// PRECEDENCIA DE OPERADORES
	// ─────────────────────────────────────────────────────────
	// Go respeta el orden matemático estándar (PEMDAS).
	// De mayor a menor precedencia:
	// 5: * / % << >> & &^
	// 4: + - | ^
	// 3: == != < <= > >=
	// 2: &&
	// 1: ||
	// Consejo: ante la duda, usá paréntesis para dejar claro el orden.

	fmt.Println("\n=== Precedencia de operadores ===")
	r1 := 2 + 3*4         // 14 (primero 3*4=12, luego 2+12)
	r2 := (2 + 3) * 4     // 20 (paréntesis primero)
	r3 := 10 > 5 && 3 < 7 // true (ambas comparaciones true)
	r4 := 10 > 5 || 3 > 7 // true (la primera es true)
	fmt.Printf("2 + 3*4   = %d\n", r1)
	fmt.Printf("(2+3) * 4 = %d\n", r2)
	fmt.Printf("10>5 && 3<7 = %v\n", r3)
	fmt.Printf("10>5 || 3>7 = %v\n", r4)

	// ─────────────────────────────────────────────────────────
	// EJEMPLO REAL: Cálculo de Hipotenusa
	// ─────────────────────────────────────────────────────────
	cateto1 := 3.0
	cateto2 := 4.0
	hipotenusa := math.Sqrt(cateto1*cateto1 + cateto2*cateto2)
	fmt.Printf("\n=== Hipotenusa del triángulo 3-4-x ===\n")
	fmt.Printf("√(%.0f² + %.0f²) = √%.0f = %.2f\n",
		cateto1, cateto2, cateto1*cateto1+cateto2*cateto2, hipotenusa)
}

// Función auxiliar para el ejemplo de short-circuit
func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
