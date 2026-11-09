package main

import "fmt"

func main() {
	// =========================================================
	// CICLO FOR CON RANGO (for range)
	// =========================================================
	// La forma "for range" es la más idiomática de Go para
	// iterar sobre colecciones. Es más segura y expresiva que
	// el for clásico porque Go maneja los índices internamente.
	//
	// Sintaxis general:
	//   for índice, valor := range colección { }
	//
	// Funciona con: slices, arrays, strings, maps, channels, enteros (Go 1.22+)
	// y funciones iteradoras (Go 1.23+, que veremos en 08_generics_e_iteradores).
	//
	// En cada iteración, range retorna:
	//   - slices/arrays : índice (int),    valor (copia del elemento)
	//   - strings       : índice (int, posición en bytes), rune (caracter Unicode)
	//   - maps          : clave,           valor
	//   - channels      : valor (uno solo, hasta que se cierre el canal)
	//   - enteros (N)   : entero de 0 a N-1 (o sin variable: `for range N`)
	//
	// Podés ignorar cualquiera con _:
	//   for _, v := range s  → solo valor
	//   for i := range s     → solo índice (sin la coma!)
	//   for range s          → repetir len(s) veces sin usar variables

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE SLICE / ARRAY
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== For range sobre slice ===")

	frutas := []string{"manzana", "banana", "naranja", "uva", "mango"}

	// Índice y valor
	for i, fruta := range frutas {
		fmt.Printf("[%d] %s\n", i, fruta)
	}

	// Solo el valor (ignorar índice con _)
	fmt.Println("\nSolo valores:")
	for _, fruta := range frutas {
		fmt.Print(fruta, " ")
	}
	fmt.Println()

	// Solo el índice
	fmt.Println("\nSolo índices:")
	for i := range frutas {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// IMPORTANTE: el valor es una COPIA, no una referencia
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El valor es COPIA (no referencia) ===")

	precios := []float64{10.0, 20.0, 30.0}

	// Esto NO modifica el slice original
	for _, p := range precios {
		p *= 2 // modifica la copia local, no el slice
		_ = p
	}
	fmt.Println("precios sin cambio:", precios) // [10 20 30]

	// Para MODIFICAR, usá el índice
	for i := range precios {
		precios[i] *= 2 // esto SÍ modifica el slice
	}
	fmt.Println("precios modificados:", precios) // [20 40 60]

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE ARRAY
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== For range sobre array ===")

	diasHabiles := [5]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes"}

	for dia, nombre := range diasHabiles {
		fmt.Printf("Día %d: %s\n", dia+1, nombre)
	}

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE STRING
	// ─────────────────────────────────────────────────────────
	// Al iterar un string, range retorna:
	//   - índice: posición del BYTE donde comienza el rune
	//   - valor:  el rune (caracter Unicode completo)
	//
	// Esto es diferente a recorrer byte a byte.
	// ¡Es la forma CORRECTA de iterar sobre strings con UTF-8!

	fmt.Println("\n=== For range sobre string (runes) ===")

	texto := "Hola Ñoño"
	fmt.Printf("String: '%s' | len=%d bytes\n\n", texto, len(texto))

	for i, r := range texto {
		fmt.Printf("byte[%d]: '%c' (Unicode: U+%04X)\n", i, r, r)
	}

	// Nótese: la Ñ y la ó ocupan más de 1 byte en UTF-8,
	// por eso los índices no son consecutivos de 1 en 1.

	// Comparación: byte a byte vs rune a rune
	fmt.Println("\n--- byte a byte (cuidado con UTF-8) ---")
	for i := 0; i < len(texto); i++ {
		fmt.Printf("texto[%d] = %d\n", i, texto[i])
	}

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE MAP
	// ─────────────────────────────────────────────────────────
	// Lo vemos en profundidad en 03_colecciones_y_strings/04_maps y 05_iteracion_de_maps,
	// pero aquí una muestra rápida:

	fmt.Println("\n=== For range sobre map ===")

	capitales := map[string]string{
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
	}

	// El orden de iteración sobre un map es ALEATORIO (por diseño de Go)
	for pais, capital := range capitales {
		fmt.Printf("%s → %s\n", pais, capital)
	}

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE CHANNEL
	// ─────────────────────────────────────────────────────────
	// Un channel es como un tubo por donde pasan datos entre goroutines.
	// El for range sobre un channel itera hasta que el channel se cierra.
	// (Lo veremos en detalle en concurrencia, pero aquí el ejemplo básico)

	fmt.Println("\n=== For range sobre channel ===")

	ch := make(chan int, 5) // channel con buffer de 5
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch) // cerramos para que el range termine

	for valor := range ch {
		fmt.Println("Recibido del channel:", valor)
	}

	// ─────────────────────────────────────────────────────────
	// FOR RANGE SOBRE ENTEROS (Go 1.22+)
	// ─────────────────────────────────────────────────────────
	// Desde Go 1.22, podés hacer `for i := range N` donde N es un entero.
	// Recorre desde 0 hasta N-1 (igual que `for i := 0; i < N; i++`).
	// Si N <= 0, el cuerpo del bucle no se ejecuta ninguna vez.
	fmt.Println("\n=== For range sobre enteros (Go 1.22+) ===")

	fmt.Print("for i := range 5 → ")
	for i := range 5 {
		fmt.Printf("%d ", i) // imprime: 0 1 2 3 4
	}
	fmt.Println()

	// También podés omitir la variable si solo querés repetir algo N veces:
	fmt.Print("for range 3      → ")
	for range 3 {
		fmt.Print("¡Go! ")
	}
	fmt.Println()

	// Con range 0 no entra nunca al cuerpo:
	for range 0 {
		fmt.Println("Esto jamás se ejecuta")
	}

	// ─────────────────────────────────────────────────────────
	// CASOS PRÁCTICOS
	// ─────────────────────────────────────────────────────────

	// Caso 1: Sumar todos los elementos
	fmt.Println("\n=== Caso: sumar elementos ===")
	numeros := []int{3, 7, 2, 9, 1, 5, 8, 4, 6}
	suma := 0
	for _, n := range numeros {
		suma += n
	}
	fmt.Printf("Suma de %v = %d\n", numeros, suma)

	// Caso 2: Contar coincidencias
	fmt.Println("\n=== Caso: contar ocurrencias ===")
	palabras := []string{"go", "python", "go", "java", "go", "rust"}
	contador := 0
	for _, p := range palabras {
		if p == "go" {
			contador++
		}
	}
	fmt.Printf("'go' aparece %d veces\n", contador)

	// Caso 3: Construir un nuevo slice a partir del rango
	fmt.Println("\n=== Caso: construir slice filtrado ===")
	todos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var pares []int
	for _, n := range todos {
		if n%2 == 0 {
			pares = append(pares, n)
		}
	}
	fmt.Println("Pares:", pares)

	// Caso 4: Encontrar el máximo
	fmt.Println("\n=== Caso: encontrar máximo ===")
	maximo := numeros[0]
	for _, n := range numeros[1:] {
		if n > maximo {
			maximo = n
		}
	}
	fmt.Println("Máximo:", maximo)

	// Caso 5: Tabla de frecuencia de caracteres
	fmt.Println("\n=== Caso: frecuencia de caracteres ===")
	frase := "golang rocks"
	frecuencia := make(map[rune]int)
	for _, r := range frase {
		if r != ' ' {
			frecuencia[r]++
		}
	}
	for ch2, count := range frecuencia {
		fmt.Printf("  '%c': %d\n", ch2, count)
	}

	// ─────────────────────────────────────────────────────────
	// FOR RANGE CON ÍNDICE SOLO (patrón idiomático)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== for i := range (solo índice) ===")

	// Inicializar slice con cuadrados
	cuadrados := make([]int, 6)
	for i := range cuadrados {
		cuadrados[i] = (i + 1) * (i + 1)
	}
	fmt.Println("Cuadrados:", cuadrados)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen de for range ===")
	fmt.Println("for i, v := range slice   → índice + copia del valor")
	fmt.Println("for _, v := range slice   → solo valor")
	fmt.Println("for i := range slice      → solo índice")
	fmt.Println("for i, r := range string  → índice int (posición en bytes) + rune Unicode")
	fmt.Println("for k, v := range map     → clave + valor (orden no especificado)")
	fmt.Println("for v := range channel    → valor (hasta que se cierre)")
	fmt.Println("for i := range N          → enteros de 0 a N-1 (Go 1.22+)")
	fmt.Println("for range N               → repite el bloque N veces (Go 1.22+)")
	fmt.Println()
	fmt.Println("⚠️ El valor en range es una COPIA.")
	fmt.Println("   Para modificar: usá el índice → slice[i] = nuevoValor")
}
