package main

import "fmt"

func main() {
	// =========================================================
	// MÚLTIPLES VARIABLES Y PATRONES COMUNES
	// =========================================================

	// ─────────────────────────────────────────────────────────
	// DECLARACIÓN MÚLTIPLE EN UNA LÍNEA
	// ─────────────────────────────────────────────────────────
	a, b, c := 10, 20, 30
	x, y := "hola", true

	fmt.Println("=== Múltiples variables en una línea ===")
	fmt.Println(a, b, c)
	fmt.Println(x, y)

	// ─────────────────────────────────────────────────────────
	// INTERCAMBIO DE VALORES (swap)
	// ─────────────────────────────────────────────────────────
	// Go permite intercambiar valores sin variable temporal,
	// algo que en otros lenguajes requiere una variable auxiliar.
	fmt.Println("\n=== Swap de valores ===")
	fmt.Printf("Antes: a=%d, b=%d\n", a, b)
	a, b = b, a // intercambio en una línea
	fmt.Printf("Después: a=%d, b=%d\n", a, b)

	// En otros lenguajes necesitarías:
	// temp := a; a = b; b = temp

	// ─────────────────────────────────────────────────────────
	// RETORNO MÚLTIPLE DE FUNCIONES
	// ─────────────────────────────────────────────────────────
	// Go soporta retorno de múltiples valores nativamente.
	// El caso más común es retornar (resultado, error).
	fmt.Println("\n=== Retorno múltiple de funciones ===")

	cociente, resto := dividirConResto(17, 5)
	fmt.Printf("17 ÷ 5 = %d con resto %d\n", cociente, resto)

	nombre, apellido, edad := datosPersona()
	fmt.Printf("Persona: %s %s, %d años\n", nombre, apellido, edad)

	// Ignorar valores con _
	soloNombre, _, _ := datosPersona()
	fmt.Println("Solo el nombre:", soloNombre)

	// ─────────────────────────────────────────────────────────
	// RETORNO NOMBRADO (named return values)
	// ─────────────────────────────────────────────────────────
	// Podemos nombrar los valores de retorno en la firma de la función.
	// Esto permite usar "return" sin argumentos (naked return).
	area, perimetro := calcularRectangulo(5, 3)
	fmt.Printf("\nRectángulo 5x3: área=%d, perímetro=%d\n", area, perimetro)

	// ─────────────────────────────────────────────────────────
	// BLOQUE var PARA VARIABLES RELACIONADAS
	// ─────────────────────────────────────────────────────────
	// Agrupa variables que pertenecen al mismo concepto.
	var (
		servidorHost  string = "localhost"
		servidorPort  int    = 8080
		maxConexiones int    = 100
		timeout       int    = 30
	)

	fmt.Println("\n=== Bloque var para configuración ===")
	fmt.Printf("Servidor: %s:%d\n", servidorHost, servidorPort)
	fmt.Printf("Max conexiones: %d | Timeout: %ds\n", maxConexiones, timeout)

	// ─────────────────────────────────────────────────────────
	// VARIABLES EN LOOPS (patrón acumulador)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Acumuladores con variables ===")

	numeros := []int{3, 7, 2, 9, 1, 5, 8, 4, 6}
	suma, maximo, minimo := 0, numeros[0], numeros[0]

	for _, n := range numeros {
		suma += n
		if n > maximo {
			maximo = n
		}
		if n < minimo {
			minimo = n
		}
	}

	promedio := float64(suma) / float64(len(numeros))
	fmt.Printf("Números: %v\n", numeros)
	fmt.Printf("Suma: %d | Promedio: %.2f\n", suma, promedio)
	fmt.Printf("Máximo: %d | Mínimo: %d\n", maximo, minimo)

	// ─────────────────────────────────────────────────────────
	// VARIABLES DENTRO DE IF (scope reducido)
	// ─────────────────────────────────────────────────────────
	// Go permite declarar una variable directamente en el if.
	// Esa variable SOLO existe dentro del if/else, no afuera.
	// Es un patrón muy limpio para manejar errores y retornos.

	fmt.Println("\n=== Variables en if (scope reducido) ===")

	// Patrón clásico: declarar y verificar en el mismo if
	if resultado, err := operacionRiesgosa(10); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Éxito:", resultado)
	}
	// "resultado" y "err" NO existen aquí afuera

	if resultado, err := operacionRiesgosa(-1); err != nil {
		fmt.Println("Error capturado:", err)
	} else {
		fmt.Println("Éxito:", resultado) // esto no se ejecuta
	}

	// ─────────────────────────────────────────────────────────
	// CONSTANTES AGRUPADAS CON IOTA (enumeraciones)
	// ─────────────────────────────────────────────────────────
	type DiaSemana int
	const (
		Domingo DiaSemana = iota
		Lunes
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
	)

	hoy := Miercoles
	fmt.Printf("\n=== Enumeración tipada ===\n")
	fmt.Printf("Hoy es día %d de la semana\n", hoy)
	fmt.Printf("¿Es fin de semana? %v\n", hoy == Domingo || hoy == Sabado)

	// ─────────────────────────────────────────────────────────
	// DESESTRUCTURACIÓN DE STRUCTS (patrón común)
	// ─────────────────────────────────────────────────────────
	type Coordenada struct {
		X, Y float64
	}

	c1 := Coordenada{3.0, 4.0}
	c2 := Coordenada{6.0, 8.0}

	// Asignamos múltiples campos en una expresión
	distX := c2.X - c1.X
	distY := c2.Y - c1.Y
	fmt.Printf("\n=== Coordenadas ===\n")
	fmt.Printf("Punto 1: (%.1f, %.1f)\n", c1.X, c1.Y)
	fmt.Printf("Punto 2: (%.1f, %.1f)\n", c2.X, c2.Y)
	fmt.Printf("Diferencia: (%.1f, %.1f)\n", distX, distY)

	// ─────────────────────────────────────────────────────────
	// VARIABLES COMO CONTADORES DE ESTADO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Contadores de estado ===")

	palabras := []string{"go", "es", "genial", "go", "es", "rápido", "go"}
	frecuencias := make(map[string]int)

	for _, palabra := range palabras {
		frecuencias[palabra]++ // incrementa el contador, parte de 0 (zero value)
	}

	for palabra, count := range frecuencias {
		fmt.Printf("'%s' aparece %d veces\n", palabra, count)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE PATRONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrones clave ===")
	fmt.Println("a, b := 1, 2          → múltiples variables a la vez")
	fmt.Println("a, b = b, a           → swap sin variable temporal")
	fmt.Println("val, err := fn()      → capturar múltiples retornos")
	fmt.Println("_, err := fn()        → ignorar valor con _")
	fmt.Println("if v, err := fn(); err != nil → scope reducido en if")
	fmt.Println("var (...)             → agrupar variables relacionadas")
}

// Función que retorna dos valores
func dividirConResto(a, b int) (int, int) {
	return a / b, a % b
}

// Función que retorna tres valores
func datosPersona() (string, string, int) {
	return "Matias", "García", 25
}

// Función con retorno nombrado
func calcularRectangulo(largo, ancho int) (area int, perimetro int) {
	area = largo * ancho // asignamos a los nombres de retorno
	perimetro = 2 * (largo + ancho)
	return // "naked return": retorna los valores nombrados
}

// Función que puede retornar error
func operacionRiesgosa(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("el número %d no puede ser negativo", n)
	}
	return fmt.Sprintf("procesado: %d", n*2), nil
}
