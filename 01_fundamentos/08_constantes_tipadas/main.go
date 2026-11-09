package main

import "fmt"

// Definimos un tipo propio basado en int a nivel de paquete para que
// tanto main() como la función describir() puedan usarlo con seguridad de tipos.
type EstadoOrden int

const (
	Borrador   EstadoOrden = iota // 0 (tipo EstadoOrden)
	Pendiente                     // 1 (repite tipo EstadoOrden e incrementa iota)
	Procesando                    // 2
	Completado                    // 3
	Cancelado                     // 4
)

func main() {
	// =========================================================
	// CONSTANTES TIPADAS Y NO TIPADAS
	// =========================================================
	// Este es uno de los temas más sutiles pero importantes de Go.
	// Las constantes en Go pueden ser de dos categorías:
	//   1. NO tipadas (untyped): más flexibles y con precisión arbitraria
	//   2. Tipadas (typed): más estrictas, atadas a un tipo concreto

	// ─────────────────────────────────────────────────────────
	// CONSTANTES NO TIPADAS (untyped constants)
	// ─────────────────────────────────────────────────────────
	// Cuando declarás una constante SIN especificar el tipo,
	// Go la trata como "no tipada". Tiene un "tipo por defecto" (int, float64,
	// rune, string, bool o complex128) cuando hace falta inferirlo (por ejemplo con :=),
	// pero puede adaptarse a cualquier tipo compatible sin conversión explícita.

	const sinTipo = 42      // constante entera no tipada (tipo por defecto: int)
	const sinTipoF = 3.14   // constante flotante no tipada (tipo por defecto: float64)
	const sinTipoS = "hola" // constante string no tipada (tipo por defecto: string)
	const sinTipoB = true   // constante booleana no tipada (tipo por defecto: bool)

	// Lo poderoso de las constantes no tipadas:
	// podemos asignarlas a DISTINTOS tipos numéricos sin conversión explícita,
	// siempre que el valor quepa en el tipo destino.
	var x int = sinTipo     // funciona
	var y int64 = sinTipo   // también funciona (¡sin conversión!)
	var z float64 = sinTipo // ¡también funciona!
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
	// Solo puede usarse donde se espere ESE tipo exacto.

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
	// var b int64 = conTipoInt // ERROR: cannot use conTipoInt (constant 42 of type int) as int64
	var b int64 = int64(conTipoInt) // necesita conversión explícita
	fmt.Printf("\nint = %d, convertido a int64 = %d\n", a, b)

	// ─────────────────────────────────────────────────────────
	// DIFERENCIA PRÁCTICA: NO tipada vs TIPADA
	// ─────────────────────────────────────────────────────────
	// Caso clásico: usamos la constante en una expresión con decimales

	const untyped = 100   // no tipada
	const typed int = 100 // tipada como int

	var resultado64 float64

	// Con no tipada: funciona directamente (Go adapta el 100 a float64)
	resultado64 = untyped * 1.5
	fmt.Println("\n=== Diferencia práctica ===")
	fmt.Printf("untyped(100) * 1.5 = %f  ✓ funciona sin convertir\n", resultado64)

	// Con tipada: typed * 1.5 daría error porque typed ya es int y 1.5 no entra en int.
	// Hay que convertir explícitamente a float64:
	resultado64 = float64(typed) * 1.5
	fmt.Printf("float64(typed(100)) * 1.5 = %f  (necesitó conversión)\n", resultado64)

	// ─────────────────────────────────────────────────────────
	// PRECISIÓN ARBITRARIA DE CONSTANTES NO TIPADAS
	// ─────────────────────────────────────────────────────────
	// Las constantes no tipadas en Go tienen PRECISIÓN ARBITRARIA en compilación:
	// pueden representar números mucho más grandes que int64 o uint64 sin desbordar.
	// Recién cuando las asignás a una variable en memoria deben caber en ese tipo.

	const unYottabyte = 1 << 80 // 2^80 = 1_208_925_819_614_629_174_706_176 (¡no cabe en 64 bits!)
	const unGigabyte = 1 << 30  // 2^30
	// Aunque unYottabyte no entra en un int64 directamente, podemos operar con él
	// en tiempo de compilación y guardar el resultado ya reducido en un int:
	gigabytesEnYottabyte := unYottabyte / unGigabyte // 2^50, sí cabe en int de 64 bits
	const piPreciso = 3.14159265358979323846264338327950288

	fmt.Println("\n=== Precisión arbitraria de constantes no tipadas ===")
	fmt.Printf("Gigabytes en 1 Yottabyte (calculado con 2^80 / 2^30): %d\n", gigabytesEnYottabyte)
	fmt.Printf("pi preciso reducido a float64: %.20f\n", float64(piPreciso))

	// ─────────────────────────────────────────────────────────
	// CONSTANTES TIPADAS CON iota (enumeraciones seguras)
	// ─────────────────────────────────────────────────────────
	// Un patrón muy común es crear un tipo personalizado (EstadoOrden) y
	// usarlo con iota para que una función acepte SOLO ese tipo de estado.

	estadoActual := Procesando

	fmt.Println("\n=== Enumeración tipada con iota ===")
	fmt.Printf("Estado actual: %d (tipo: %T)\n", estadoActual, estadoActual)
	fmt.Println("Descripción:", describir(estadoActual))

	// Esto es type-safe frente a variables de tipo int u otros tipos:
	// codigoSuero := 2
	// describir(codigoSuero) // ERROR: cannot use codigoSuero (variable of type int) as EstadoOrden

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("NO tipadas: const Pi = 3.14          → más flexibles, precisión arbitraria")
	fmt.Println("Tipadas:    const Pi float64 = 3.14  → estrictas, exigen su tipo exacto")
	fmt.Println("")
	fmt.Println("Usá NO tipadas por defecto para números y textos constantes.")
	fmt.Println("Usá TIPADAS cuando quieras crear enumeraciones type-safe con iota.")
}

// describir recibe específicamente un EstadoOrden (no un int genérico ni any).
func describir(estado EstadoOrden) string {
	switch estado {
	case Borrador:
		return "Borrador"
	case Pendiente:
		return "Pendiente"
	case Procesando:
		return "Procesando"
	case Completado:
		return "Completado"
	case Cancelado:
		return "Cancelado"
	default:
		return "Desconocido"
	}
}
