package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// =========================================================
	// CICLO FOR DE SOLO CONDICIÓN (equivalente a "while")
	// =========================================================
	// En Go no existe la palabra "while". En cambio, el "for"
	// con solo una condición (sin inicialización ni post)
	// se comporta exactamente igual.
	//
	// Sintaxis:  for condición { }
	//
	// Funciona así:
	//   1. Evalúa la condición
	//   2. Si es true  → ejecuta el cuerpo → vuelve al paso 1
	//   3. Si es false → sale del bucle
	//
	// Equivalente exacto en otros lenguajes:
	//   while (condición) { }   (Java, C, C++)
	//   while condición:         (Python)

	// ─────────────────────────────────────────────────────────
	// EJEMPLO BÁSICO
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== For de condición básico ===")

	n := 1
	for n <= 5 { // equivalente a: while n <= 5
		fmt.Printf("n = %d\n", n)
		n++ // importante: si olvidás esto, el bucle es infinito
	}
	fmt.Println("Salimos del for, n =", n)

	// ─────────────────────────────────────────────────────────
	// COMPARACIÓN: FOR CLÁSICO vs FOR CONDICIÓN
	// ─────────────────────────────────────────────────────────
	// Estos dos hacen exactamente lo mismo:

	fmt.Println("\n=== For clásico vs For condición (equivalentes) ===")

	// For CLÁSICO:
	fmt.Print("For clásico: ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// For CONDICIÓN (while):
	fmt.Print("For condición: ")
	j := 1
	for j <= 5 {
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	// ¿Cuándo conviene el for-condición?
	// Cuando la condición de parada depende de algo que no es un simple contador.

	// ─────────────────────────────────────────────────────────
	// CASO 1: Leer hasta una condición dinámica
	// ─────────────────────────────────────────────────────────
	// Simulamos una cola de trabajo que se procesa hasta que queda vacía.
	fmt.Println("\n=== Procesar cola hasta vaciar ===")

	cola := []string{"tarea1", "tarea2", "tarea3", "tarea4"}

	for len(cola) > 0 {
		// Tomamos la primera tarea
		tarea := cola[0]
		cola = cola[1:] // eliminamos la primera

		fmt.Printf("Procesando: %s (quedan %d en cola)\n", tarea, len(cola))
	}
	fmt.Println("Cola vacía, listo.")

	// ─────────────────────────────────────────────────────────
	// CASO 2: Buscar en datos hasta encontrar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Buscar hasta encontrar (simulación) ===")

	registros := []struct {
		id    int
		valor string
	}{
		{1, "alfa"}, {2, "beta"}, {3, "objetivo"}, {4, "delta"}, {5, "épsilon"},
	}

	idx := 0
	for idx < len(registros) && registros[idx].valor != "objetivo" {
		fmt.Printf("  saltando id=%d (%s)\n", registros[idx].id, registros[idx].valor)
		idx++
	}

	if idx < len(registros) {
		fmt.Printf("  Encontrado: id=%d (%s)\n", registros[idx].id, registros[idx].valor)
	}

	// ─────────────────────────────────────────────────────────
	// CASO 3: Reintentos hasta éxito (patrón retry)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrón retry ===")

	const maxIntentos = 5
	intentos := 0
	exito := false

	// Simulamos una operación que puede fallar aleatoriamente
	rand.New(rand.NewSource(42)) // seed fijo para reproducibilidad

	for intentos < maxIntentos && !exito {
		intentos++
		if operacionFallida(intentos) {
			fmt.Printf("  Intento %d: falló\n", intentos)
		} else {
			fmt.Printf("  Intento %d: éxito!\n", intentos)
			exito = true
		}
	}

	if exito {
		fmt.Printf("Completado en %d intento(s)\n", intentos)
	} else {
		fmt.Printf("Falló después de %d intentos\n", maxIntentos)
	}

	// ─────────────────────────────────────────────────────────
	// CASO 4: Algoritmo de aproximación (convergencia)
	// ─────────────────────────────────────────────────────────
	// Método de Newton-Raphson para calcular raíz cuadrada de 2.
	// Iteramos hasta que la aproximación sea suficientemente precisa.

	fmt.Println("\n=== Aproximación: raíz cuadrada de 2 ===")

	aproximacion := 1.0
	precision := 0.000001
	iteracion := 0

	for {
		mejora := (aproximacion + 2.0/aproximacion) / 2.0
		iteracion++
		fmt.Printf("  iter %d: %.10f\n", iteracion, mejora)

		if abs(mejora-aproximacion) < precision {
			aproximacion = mejora
			break
		}
		aproximacion = mejora
	}

	fmt.Printf("√2 ≈ %.10f (en %d iteraciones)\n", aproximacion, iteracion)
	fmt.Println("√2 real: 1.4142135623")

	// ─────────────────────────────────────────────────────────
	// CASO 5: Validación de entrada (simulación)
	// ─────────────────────────────────────────────────────────
	// En una app real, el usuario ingresaría un número válido.
	// Aquí lo simulamos con un slice de "intentos de ingreso".

	fmt.Println("\n=== Validar entrada (simulado) ===")

	entradas := []int{-5, 0, 200, 42} // simula lo que el usuario ingresaría
	idx2 := 0
	entradaValida := -1

	for entradaValida < 0 && idx2 < len(entradas) {
		intento := entradas[idx2]
		idx2++

		if intento <= 0 || intento > 100 {
			fmt.Printf("  '%d' no es válido (debe ser 1-100)\n", intento)
		} else {
			entradaValida = intento
		}
	}

	if entradaValida > 0 {
		fmt.Printf("  Entrada aceptada: %d\n", entradaValida)
	}

	// ─────────────────────────────────────────────────────────
	// PELIGRO: BUCLE INFINITO ACCIDENTAL
	// ─────────────────────────────────────────────────────────
	// Si la condición NUNCA se vuelve false, el programa cuelga.
	// Esto es un error típico con el for-condición.

	fmt.Println("\n=== ⚠️ Precaución: condición que nunca cambia ===")
	fmt.Println("Este for sería INFINITO si no tuviéramos el break:")

	contador := 10
	pasos := 0
	for contador > 0 { // condición depende de 'contador'
		pasos++
		// ¡IMPORTANTE: 'contador' DEBE cambiar dentro del cuerpo!
		contador -= 3
		if pasos > 20 { // seguro de emergencia para el ejemplo
			fmt.Println("  (cortado por seguridad)")
			break
		}
	}
	fmt.Printf("  Terminó: contador=%d, pasos=%d\n", contador, pasos)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("for condicion { }  →  equivale a  while (condicion) { }")
	fmt.Println("")
	fmt.Println("Usalo cuando:")
	fmt.Println("  - No sabés cuántas iteraciones habrá de antemano")
	fmt.Println("  - La condición depende de algo que cambia dentro del cuerpo")
	fmt.Println("  - Procesás cola/pila hasta que se vacíe")
	fmt.Println("  - Reintentos hasta éxito")
	fmt.Println("  - Algoritmos de convergencia")
	fmt.Println("")
	fmt.Println("⚠️ Siempre asegurate de que la condición pueda volverse false!")
}

// operacionFallida simula una operación que falla las primeras veces
func operacionFallida(intento int) bool {
	return intento < 3 // falla en intentos 1 y 2, éxito en 3
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
