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
	// := siempre CREA una nueva variable (puede sombrear)
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
	// El patrón "val, err := fn()" crea un shadowing de err
	// en cada llamada. Esto es intencional y es el estilo de Go.
	// Pero puede generar confusión si no se entiende bien.

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
		// "n" en cada iteración es la misma variable (reutilizada)
		// No es shadowing en sí, pero el comportamiento puede sorprender
		// cuando creás goroutines (hilos) dentro del for.
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
	fmt.Println("  go vet ./...                    → análisis básico")
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
