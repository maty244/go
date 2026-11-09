package main

import "fmt"

// =========================================================
// VARIABLES DE SOMBRA (SHADOWING)
// =========================================================
// El "shadowing" ocurre cuando declarás una variable con el MISMO
// nombre que otra variable en un scope exterior.
// La variable interior "cubre" o "sombrea" a la exterior.
// La exterior NO se modifica; simplemente no es visible desde adentro.
//
// Go PERMITE el shadowing pero puede ser fuente de bugs sutiles
// si no se comprende bien.

// Variable global para los ejemplos
var mensaje = "soy global"

func main() {
	// ─────────────────────────────────────────────────────────
	// EJEMPLO BÁSICO: shadowing con bloques {}
	// ─────────────────────────────────────────────────────────
	x := "exterior"

	fmt.Println("=== Shadowing básico ===")
	fmt.Println("x antes del bloque:", x) // "exterior"

	{
		// Esta := crea una NUEVA variable x en este scope.
		// No modifica la x exterior.
		x := "interior"                        // ← nueva variable, sombrea a la exterior
		fmt.Println("x dentro del bloque:", x) // "interior"
	}

	fmt.Println("x después del bloque:", x) // "exterior" (sin cambios)

	// ─────────────────────────────────────────────────────────
	// DIFERENCIA CRÍTICA: := vs =
	// ─────────────────────────────────────────────────────────
	// := declara al menos una variable nueva; en el MISMO bloque puede
	// reasignar otras ya existentes: valor, nuevo := 20, true.
	// En un bloque interior, declarar valor con := sí sombrea a la exterior.
	// =  siempre MODIFICA la variable existente (no hace shadowing)

	valor := 10
	fmt.Println("\n=== := vs = en bloques ===")
	fmt.Println("valor antes:", valor) // 10

	{
		valor := 99                         // := → nueva variable LOCAL, sombrea a la exterior
		fmt.Println("dentro con :=", valor) // 99
	}
	fmt.Println("valor después de :=:", valor) // 10 (sin cambios!)

	{
		valor = 99                          // = → modifica la variable EXTERIOR
		fmt.Println("dentro con =:", valor) // 99
	}
	fmt.Println("valor después de =:", valor) // 99 (modificado!)

	// CASO SUTIL: := con una variable existente y una nueva
	// 1. En el MISMO bloque: "valor" se REASIGNA y "nuevo" se DECLARA:
	valor, nuevo := 20, true
	fmt.Printf("en el mismo bloque (valor, nuevo := 20, true) → valor=%d, nuevo=%v\n", valor, nuevo)
	// 2. En un bloque INTERIOR: AMBAS se declaran como nuevas (valor sombrea a la exterior):
	{
		valor, otroNuevo := 500, false
		fmt.Printf("en bloque interior (valor, otroNuevo := 500, false) → valor interno=%d, otroNuevo=%v\n", valor, otroNuevo)
	}
	fmt.Println("valor exterior después del bloque interior:", valor) // sigue siendo 20!

	// ─────────────────────────────────────────────────────────
	// SHADOWING CON VARIABLES GLOBALES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Sombra sobre variable global ===")
	fmt.Println("global:", mensaje) // "soy global"

	mensaje := "soy local"                          // := crea una variable LOCAL que sombrea a la global
	fmt.Println("local (sombrea global):", mensaje) // "soy local"

	mostrarGlobal() // la función ve la global, no la local
	// El package-level "mensaje" sigue siendo "soy global"

	// ─────────────────────────────────────────────────────────
	// SHADOWING EN IF (la trampa más común)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Trampa clásica: shadowing en if ===")

	var resultado int

	// VERSIÓN CON BUG DE SOMBRA:
	if true {
		resultado := calcular(5)                            // := crea una NUEVA variable local
		fmt.Println("dentro del if, resultado:", resultado) // 25
		_ = resultado                                       // suprimir unused
	}
	fmt.Println("resultado afuera del if:", resultado) // 0 ← BUG! no se modificó

	fmt.Println()

	// VERSIÓN CORRECTA sin sombra:
	if true {
		resultado = calcular(5)                             // = modifica la variable externa
		fmt.Println("dentro del if, resultado:", resultado) // 25
	}
	fmt.Println("resultado afuera del if:", resultado) // 25 ← correcto!

	// ─────────────────────────────────────────────────────────
	// SHADOWING CON err (MUY FRECUENTE EN GO)
	// ─────────────────────────────────────────────────────────
	// "val, err := fn()" puede reutilizar err en el mismo bloque si val es
	// nuevo. En un bloque interior puede declarar otra err y sombrear a la
	// exterior. El scope, no el nombre por sí solo, decide qué variable es.

	fmt.Println("\n=== Shadowing con err ===")

	var errFinal error

	if true {
		resultado, err := dividir(10, 2) // err es NUEVA aquí
		if err != nil {
			errFinal = err
		}
		fmt.Printf("10/2 = %.1f (err local: %v)\n", resultado, err)
	}

	if true {
		resultado, err := dividir(10, 0) // otra err NUEVA
		if err != nil {
			errFinal = err
		}
		fmt.Printf("10/0 = %.1f (err local: %v)\n", resultado, err)
	}

	fmt.Println("errFinal:", errFinal)

	// ─────────────────────────────────────────────────────────
	// SHADOWING EN FOR RANGE (efecto sorpresa)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Shadowing en for range ===")

	numeros := []int{1, 2, 3}

	for _, n := range numeros {
		// Desde Go 1.22, n DECLARADA con := es nueva en cada iteración.
		// Con una n declarada afuera y "for _, n = range numeros", se
		// reasigna la misma variable. Ver 05_funciones_y_closures/10_trampa_closure_en_for.
		fmt.Println("n:", n)
	}

	// ─────────────────────────────────────────────────────────
	// SHADOWING INTENCIONAL Y ÚTIL
	// ─────────────────────────────────────────────────────────
	// A veces el shadowing es intencional y hace el código más limpio.
	// Por ejemplo, en el if con inicialización:

	fmt.Println("\n=== Shadowing útil: if con init ===")

	err := hacerAlgo()
	fmt.Println("err inicial:", err)

	if err := hacerAlgoConError(); err != nil {
		// "err" aquí sombrea a la exterior, está acotada a este if
		fmt.Println("Error capturado en el if:", err)
	}

	fmt.Println("err exterior sin cambios:", err)

	// ─────────────────────────────────────────────────────────
	// CÓMO DETECTAR SHADOWING: go vet y staticcheck
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cómo detectar sombras ===")
	fmt.Println("Herramientas para detectar shadowing no intencional:")
	fmt.Println("  go vet ./...                    → análisis básico; no detecta toda sombra")
	fmt.Println("  staticcheck ./...               → análisis estático avanzado")
	fmt.Println("  shadow (golang.org/x/tools)     → detecta shadowing específicamente")
	fmt.Println()
	fmt.Println("En IDEs como GoLand o VS Code con gopls,")
	fmt.Println("las variables sombreadas suelen resaltarse visualmente.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN: CUÁNDO PRESTAR ATENCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reglas para evitar bugs de shadowing ===")
	fmt.Println("1. Dentro de if/for/switch, usar = si querés modificar variable externa")
	fmt.Println("2. Usar := solo si necesitás una variable nueva en ese scope")
	fmt.Println("3. Cuidado con el patrón:  if v, err := fn(); err != nil {")
	fmt.Println("   La 'err' dentro del if ES diferente a cualquier 'err' exterior")
	fmt.Println("4. Los IDEs modernos resaltan el shadowing visualmente")
	fmt.Println("5. Dar nombres distintos cuando sea posible para mayor claridad")
}

func calcular(n int) int {
	return n * n
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

func mostrarGlobal() {
	fmt.Println("desde función, global:", mensaje) // ve la global original
}

func hacerAlgo() error {
	return nil
}

func hacerAlgoConError() error {
	return fmt.Errorf("algo salió mal")
}
