package main

import "fmt"

// =========================================================
// var vs := (short variable declaration)
// =========================================================
// Go tiene DOS formas principales de declarar variables:
//   1. var nombreVar tipo = valor
//   2. nombreVar := valor
//
// La segunda (:=) se llama "short variable declaration" o
// "walrus operator" (operador morsa, porque := parece una morsa).

// Nivel de paquete: SOLO se puede usar "var", nunca ":="
// (las variables de paquete son accesibles desde todas las funciones)
var contadorGlobal int = 0
var appNombre = "MiApp" // var con inferencia de tipo (sin especificar string)
// contadorGlobal2 := 0 // ERROR: esto no compila a nivel de paquete

func main() {
	// ─────────────────────────────────────────────────────────
	// FORMA 1: var (declaración completa)
	// ─────────────────────────────────────────────────────────
	// Sintaxis: var nombre tipo = valor
	// El tipo es opcional si se provee un valor (inferencia).
	// El valor es opcional si se provee el tipo (zero value).

	var a int = 10       // tipo y valor explícitos
	var b int            // solo tipo → zero value (0)
	var c = "hola"       // inferencia de tipo → string
	var d, e int = 5, 15 // múltiples variables del mismo tipo

	fmt.Println("=== var (forma larga) ===")
	fmt.Printf("a = %d (tipo: %T)\n", a, a)
	fmt.Printf("b = %d (zero value de int)\n", b)
	fmt.Printf("c = '%s' (tipo inferido: %T)\n", c, c)
	fmt.Printf("d = %d, e = %d\n", d, e)

	// Bloque var: útil para agrupar declaraciones relacionadas
	var (
		nombre  string  = "Matias"
		edad    int     = 25
		activo  bool    = true
		balance float64 // zero value: 0.0
	)

	fmt.Println("\n=== Bloque var ===")
	fmt.Printf("nombre=%s, edad=%d, activo=%v, balance=%.2f\n",
		nombre, edad, activo, balance)

	// ─────────────────────────────────────────────────────────
	// FORMA 2: := (short declaration, "walrus operator")
	// ─────────────────────────────────────────────────────────
	// Sintaxis: nombre := valor
	// Go INFIERE el tipo automáticamente.
	// SOLO funciona dentro de funciones.
	// DEBE tener un valor (no hay zero value con :=).

	x := 42      // int (Go infiere)
	y := 3.14    // float64 (Go infiere)
	z := "mundo" // string (Go infiere)
	ok := false  // bool (Go infiere)

	fmt.Println("\n=== := (short declaration) ===")
	fmt.Printf("x = %d (tipo: %T)\n", x, x)
	fmt.Printf("y = %f (tipo: %T)\n", y, y)
	fmt.Printf("z = '%s' (tipo: %T)\n", z, z)
	fmt.Printf("ok = %v (tipo: %T)\n", ok, ok)

	// Múltiples variables con :=
	largo, ancho := 10, 20
	area := largo * ancho
	fmt.Printf("\nlargoXancho: %d x %d = %d\n", largo, ancho, area)

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR CADA UNO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== ¿Cuándo usar cada uno? ===")

	// Usá "var" cuando:
	// 1. El tipo es importante de hacer explícito (documentación)
	var maxReintentos int = 3 // tipo explícito deja claro que es int, no int8
	fmt.Printf("maxReintentos: %d (tipo explícito por claridad)\n", maxReintentos)

	// 2. Querés el zero value (sin valor inicial)
	var buffer []byte // nil, listo para ser inicializado después
	fmt.Printf("buffer es nil: %v\n", buffer == nil)

	// 3. A nivel de paquete (obligatorio)
	fmt.Println("contadorGlobal:", contadorGlobal)

	// 4. Cuando el tipo inferido NO es el que querés
	// := infiere float64 para decimales
	autoInferido := 3.14           // float64
	var comoFloat32 float32 = 3.14 // float32 (necesita var + tipo)
	fmt.Printf("inferido := 3.14 → %T\n", autoInferido)
	fmt.Printf("var float32 = 3.14 → %T\n", comoFloat32)

	// Usá ":=" cuando:
	// 1. Dentro de funciones, como forma habitual (el 90% de los casos)
	suma := calcularSuma(5, 3)
	fmt.Printf("\nsuma := calcularSuma(5,3) → %d\n", suma)

	// 2. Para capturar retornos de funciones
	valor, err := calcularDivision(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10/2 = %v\n", valor)
	}

	// ─────────────────────────────────────────────────────────
	// REDECLARACIÓN CON := (comportamiento especial)
	// ─────────────────────────────────────────────────────────
	// Go permite usar := con variables ya existentes SI al menos
	// UNA variable del lado izquierdo es NUEVA.
	// Esta es la razón por la que podemos reutilizar "err" en
	// múltiples llamadas a funciones.

	fmt.Println("\n=== Redeclaración con := ===")

	resultado, err := calcularDivision(20, 4)
	fmt.Printf("20/4 = %v, err = %v\n", resultado, err)

	// "err" ya existe, pero "resultado2" es nueva → ":=" es válido
	resultado2, err := calcularDivision(15, 3) // err se REASIGNA, no se redeclara
	fmt.Printf("15/3 = %v, err = %v\n", resultado2, err)

	// Si TODAS las variables ya existen, daría error:
	// resultado2, err = calcularDivision(8, 2) // esto usa = (no :=)
	resultado2, err = calcularDivision(8, 2) // correcto con =
	fmt.Printf("8/2 = %v, err = %v\n", resultado2, err)

	// ─────────────────────────────────────────────────────────
	// SHADOWING: cubriendo variables externas
	// ─────────────────────────────────────────────────────────
	// Si usás := dentro de un bloque con el mismo nombre que
	// una variable externa, creás una NUEVA variable local que
	// "cubre" (shadow) a la externa. La externa no se modifica.

	variable := "exterior"
	fmt.Println("\n=== Shadowing ===")
	fmt.Println("Antes del bloque:", variable)

	{
		variable := "interior" // NUEVA variable, cubre a la exterior
		fmt.Println("Dentro del bloque:", variable)
	}

	fmt.Println("Después del bloque:", variable) // sigue siendo "exterior"

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla comparativa ===")
	fmt.Println("Característica          | var                  | :=")
	fmt.Println("------------------------|----------------------|----------------")
	fmt.Println("Nivel de paquete        | ✓ sí                | ✗ no")
	fmt.Println("Dentro de función       | ✓ sí                | ✓ sí")
	fmt.Println("Tipo explícito          | ✓ opcional          | ✗ no (inferido)")
	fmt.Println("Zero value sin valor    | ✓ sí                | ✗ no")
	fmt.Println("Cantidad de código      | más verboso          | más corto")
	fmt.Println("Uso común en funciones  | ~20%                | ~80%")
}

func calcularSuma(a, b int) int {
	return a + b
}

func calcularDivision(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}
