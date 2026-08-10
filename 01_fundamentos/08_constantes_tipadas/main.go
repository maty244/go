package main

import "fmt"

func main() {
	// =========================================================
	// CONSTANTES TIPADAS Y NO TIPADAS
	// =========================================================
	// Este es uno de los temas más sutiles pero importantes de Go.
	// Las constantes en Go pueden ser de dos categorías:
	//   1. NO tipadas (untyped): más flexibles
	//   2. Tipadas (typed): más estrictas, como variables

	// ─────────────────────────────────────────────────────────
	// CONSTANTES NO TIPADAS (untyped constants)
	// ─────────────────────────────────────────────────────────
	// Cuando declarás una constante SIN especificar el tipo,
	// Go la trata como "no tipada". Tiene un "tipo predeterminado"
	// pero puede usarse con diferentes tipos sin conversión explícita.
	// Esto las hace mucho más flexibles.

	const sinTipo = 42      // constante no tipada, "tipo ideal": integer
	const sinTipoF = 3.14   // constante no tipada, "tipo ideal": float
	const sinTipoS = "hola" // constante no tipada, "tipo ideal": string
	const sinTipoB = true   // constante no tipada, "tipo ideal": bool

	// Lo poderoso de las constantes no tipadas:
	// podemos asignarlas a DISTINTOS tipos sin conversión
	var x int = sinTipo     // funciona
	var y int64 = sinTipo   // también funciona (sin conversión!)
	var z float64 = sinTipo // también funciona!
	var w float32 = sinTipo // ¡también!

	fmt.Println("=== Constantes NO tipadas ===")
	fmt.Printf("int     x = %d (tipo: %T)\n", x, x)
	fmt.Printf("int64   y = %d (tipo: %T)\n", y, y)
	fmt.Printf("float64 z = %f (tipo: %T)\n", z, z)
	fmt.Printf("float32 w = %f (tipo: %T)\n", w, w)

	// Con tipos no numéricos
	var s1 string = sinTipoS
	var b1 bool = sinTipoB
	var f1 float64 = sinTipoF
	fmt.Printf("string  s1 = '%s'\n", s1)
	fmt.Printf("bool    b1 = %v\n", b1)
	fmt.Printf("float64 f1 = %f\n", f1)

	// ─────────────────────────────────────────────────────────
	// CONSTANTES TIPADAS (typed constants)
	// ─────────────────────────────────────────────────────────
	// Cuando especificás el tipo de la constante, se comporta
	// igual que una variable en cuanto a compatibilidad de tipos.
	// Solo puede usarse donde se espere ESE tipo específico.

	const conTipoInt int = 42
	const conTipoF64 float64 = 3.14
	const conTipoStr string = "mundo"
	const conTipoBool bool = false

	fmt.Println("\n=== Constantes TIPADAS ===")
	fmt.Printf("int     = %d (tipo: %T)\n", conTipoInt, conTipoInt)
	fmt.Printf("float64 = %f (tipo: %T)\n", conTipoF64, conTipoF64)
	fmt.Printf("string  = '%s' (tipo: %T)\n", conTipoStr, conTipoStr)
	fmt.Printf("bool    = %v (tipo: %T)\n", conTipoBool, conTipoBool)

	// Constante tipada NO puede asignarse a otro tipo sin conversión:
	var a int = conTipoInt
	// var b int64 = conTipoInt // ERROR: cannot use conTipoInt (type int) as int64
	var b int64 = int64(conTipoInt) // necesita conversión explícita
	fmt.Printf("\nint = %d, convertido a int64 = %d\n", a, b)

	// ─────────────────────────────────────────────────────────
	// DIFERENCIA PRÁCTICA: NO tipada vs TIPADA
	// ─────────────────────────────────────────────────────────
	// Caso clásico: usamos la constante en una expresión

	const untyped = 100   // no tipada
	const typed int = 100 // tipada

	var resultado64 float64

	// Con no tipada: funciona directamente (Go adapta el tipo)
	resultado64 = untyped * 1.5
	fmt.Println("\n=== Diferencia práctica ===")
	fmt.Printf("untyped(100) * 1.5 = %f  ✓ funciona\n", resultado64)

	// Con tipada: hay que convertir primero
	resultado64 = float64(typed) * 1.5
	fmt.Printf("float64(typed(100)) * 1.5 = %f  (necesitó conversión)\n", resultado64)

	// ─────────────────────────────────────────────────────────
	// PRECISIÓN DE CONSTANTES NO TIPADAS
	// ─────────────────────────────────────────────────────────
	// Las constantes no tipadas en Go tienen PRECISIÓN ARBITRARIA.
	// Pueden representar números enormes sin perder precisión.
	// Go solo convierte cuando se asignan a una variable.

	const numeroEnorme = 1_000_000_000_000_000_000 // 10^18
	const piPreciso = 3.14159265358979323846264338327950288

	fmt.Println("\n=== Precisión arbitraria de constantes no tipadas ===")
	fmt.Printf("número enorme: %d\n", numeroEnorme)
	fmt.Printf("pi preciso como float64: %.20f\n", float64(piPreciso))
	// float64 solo puede representar ~15-17 dígitos significativos,
	// pero la CONSTANTE tiene precisión arbitraria.

	// ─────────────────────────────────────────────────────────
	// CONSTANTES TIPADAS CON iota (enumeraciones con tipo)
	// ─────────────────────────────────────────────────────────
	// Un patrón muy común es crear un tipo personalizado y
	// usarlo con iota para hacer enumeraciones seguras.

	// Definimos un tipo basado en int para los estados
	type EstadoOrden int

	const (
		Borrador   EstadoOrden = iota // 0
		Pendiente                     // 1
		Procesando                    // 2
		Completado                    // 3
		Cancelado                     // 4
	)

	// Ahora la función solo acepta EstadoOrden, no cualquier int
	estadoActual := Procesando

	fmt.Println("\n=== Enumeración tipada con iota ===")
	fmt.Printf("Estado actual: %d (tipo: %T)\n", estadoActual, estadoActual)
	fmt.Println("Descripción:", describir(estadoActual))

	// Esto es type-safe: no podés pasar un int cualquiera
	// describir(2) // ERROR: no se puede pasar int donde se espera EstadoOrden

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("NO tipadas: const PI = 3.14      → más flexibles, se adaptan al contexto")
	fmt.Println("Tipadas:    const PI float64 = 3.14 → más estrictas, como variables")
	fmt.Println("")
	fmt.Println("Usa NO tipadas por defecto (son la forma estándar en Go).")
	fmt.Println("Usa TIPADAS cuando querés crear enumeraciones type-safe con iota.")
}

// Función que solo acepta EstadoOrden (no cualquier int)
// Necesitamos definir el tipo fuera de main para usarlo aquí.
// Para este ejemplo la dejamos dentro del main, pero en código real
// iría a nivel de paquete.
func describir(estado any) string {
	// Simplificación para el ejemplo
	switch estado {
	case 0:
		return "Borrador"
	case 1:
		return "Pendiente"
	case 2:
		return "Procesando"
	case 3:
		return "Completado"
	case 4:
		return "Cancelado"
	}
	return "Desconocido"
}
