package main

import "fmt"

func main() {
	// =========================================================
	// BREAK Y CONTINUE
	// =========================================================
	// Estas dos palabras clave controlan el flujo DENTRO de un bucle.
	//
	// break    → sale completamente del bucle más cercano
	// continue → salta el resto del cuerpo actual y va a la SIGUIENTE iteración
	//
	// Ambas funcionan en: for (las 3 formas) y switch.
	// Con labels (etiquetas) pueden afectar a bucles externos.

	// =========================================================
	// BREAK
	// =========================================================
	// Detiene la ejecución del bucle y continúa con el código
	// que viene DESPUÉS del for.

	// ─────────────────────────────────────────────────────────
	// BREAK BÁSICO
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== break básico ===")

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Encontramos 5, saliendo del loop")
			break // el for termina aquí
		}
		fmt.Println("i:", i)
	}
	fmt.Println("Código después del for")

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Buscar en un slice y parar al encontrar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== break: buscar en slice ===")

	emails := []string{
		"usuario1@gmail.com",
		"admin@empresa.com",
		"soporte@empresa.com",
		"usuario2@gmail.com",
	}

	buscar := "admin@empresa.com"
	encontrado := false

	for i, email := range emails {
		if email == buscar {
			fmt.Printf("Email '%s' encontrado en posición %d\n", buscar, i)
			encontrado = true
			break // no tiene sentido seguir buscando
		}
	}
	if !encontrado {
		fmt.Printf("Email '%s' no encontrado\n", buscar)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Validación con límite de intentos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== break: límite de intentos ===")

	contraseñaCorrecta := "golang123"
	intentos := []string{"1234", "pass", "golang123", "otracosa"} // simulamos intentos
	const maxIntentos = 3
	accedido := false

	for i := 0; i < maxIntentos; i++ {
		intento := intentos[i]
		fmt.Printf("Intento %d: '%s'\n", i+1, intento)

		if intento == contraseñaCorrecta {
			fmt.Println("¡Contraseña correcta! Acceso concedido.")
			accedido = true
			break
		}
		fmt.Println("Contraseña incorrecta")
	}

	if !accedido {
		fmt.Println("Cuenta bloqueada por demasiados intentos")
	}

	// ─────────────────────────────────────────────────────────
	// BREAK EN SWITCH (comportamiento diferente al de C)
	// ─────────────────────────────────────────────────────────
	// En Go, el switch NO hace fallthrough automático.
	// El break en un case sale del switch, NO del for exterior.
	fmt.Println("\n=== break en switch (sale del switch, no del for) ===")

	for i := 0; i < 5; i++ {
		switch i {
		case 3:
			fmt.Println("Caso 3: break sale del switch, no del for")
			break // sale del switch, el for sigue
		default:
			fmt.Println("Caso default, i =", i)
		}
	}

	// =========================================================
	// CONTINUE
	// =========================================================
	// Salta el resto del cuerpo de la iteración actual
	// y va directo a la siguiente iteración.

	// ─────────────────────────────────────────────────────────
	// CONTINUE BÁSICO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== continue básico ===")

	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			continue // salta los pares, no ejecuta el Println de abajo
		}
		fmt.Println("Impar:", i)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Filtrar y procesar datos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== continue: filtrar datos inválidos ===")

	temperaturas := []float64{22.5, -999, 18.0, -999, 25.3, 30.1, -999, 19.7}
	// -999 es un marcador de "dato no disponible" (sensor sin señal)

	suma := 0.0
	validos := 0

	for _, temp := range temperaturas {
		if temp == -999 {
			fmt.Println("  Saltando dato inválido (-999)")
			continue // no contamos este dato
		}
		suma += temp
		validos++
		fmt.Printf("  Dato válido: %.1f°C\n", temp)
	}

	if validos > 0 {
		fmt.Printf("Promedio de %d datos válidos: %.2f°C\n", validos, suma/float64(validos))
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Procesar líneas de un archivo (saltando comentarios)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== continue: saltar comentarios en un archivo ===")

	lineas := []string{
		"# Configuración del servidor",
		"host=localhost",
		"# Puerto de escucha",
		"port=8080",
		"",
		"# Número máximo de conexiones",
		"max_conn=100",
		"debug=false",
	}

	config := make(map[string]string)

	for _, linea := range lineas {
		// Saltar líneas vacías
		if len(linea) == 0 {
			continue
		}
		// Saltar comentarios (empiezan con #)
		if linea[0] == '#' {
			continue
		}
		// Procesar la línea: buscar el signo =
		for i, ch := range linea {
			if ch == '=' {
				clave := linea[:i]
				valor := linea[i+1:]
				config[clave] = valor
				break
			}
		}
	}

	fmt.Println("Configuración cargada:")
	for k, v := range config {
		fmt.Printf("  %s = %s\n", k, v)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: Pipeline de transformación
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== continue: pipeline de validación ===")

	pedidos := []struct {
		id      int
		monto   float64
		cliente string
	}{
		{1, 150.00, "Ana"},
		{2, -50.00, "Carlos"}, // monto inválido
		{3, 200.00, ""},       // cliente vacío
		{4, 75.00, "Mia"},
		{5, 0.00, "Juan"}, // monto cero
	}

	fmt.Println("Procesando pedidos:")
	var pedidosValidos int

	for _, p := range pedidos {
		if p.monto <= 0 {
			fmt.Printf("  [SKIP] Pedido #%d: monto inválido (%.2f)\n", p.id, p.monto)
			continue
		}
		if p.cliente == "" {
			fmt.Printf("  [SKIP] Pedido #%d: cliente vacío\n", p.id)
			continue
		}
		pedidosValidos++
		fmt.Printf("  [OK]   Pedido #%d: %s — $%.2f\n", p.id, p.cliente, p.monto)
	}
	fmt.Printf("Pedidos válidos: %d de %d\n", pedidosValidos, len(pedidos))

	// ─────────────────────────────────────────────────────────
	// BREAK Y CONTINUE EN FOR ANIDADO (afectan solo el más cercano)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== break/continue afectan solo el for más cercano ===")

	for i := 0; i < 3; i++ {
		fmt.Printf("Outer i=%d\n", i)
		for j := 0; j < 5; j++ {
			if j == 2 {
				fmt.Printf("  break en j=%d (solo sale del for interno)\n", j)
				break // sale del for de j, pero el for de i sigue
			}
			fmt.Printf("  j=%d\n", j)
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN VISUAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("break:")
	fmt.Println("  - Sale COMPLETAMENTE del bucle más cercano")
	fmt.Println("  - El código DESPUÉS del for se ejecuta")
	fmt.Println("  - También sale de un switch (pero no de un for exterior)")
	fmt.Println()
	fmt.Println("continue:")
	fmt.Println("  - Salta al SIGUIENTE ciclo del bucle")
	fmt.Println("  - El código DESPUÉS del continue en ese ciclo NO se ejecuta")
	fmt.Println("  - Útil para filtrar, saltar datos inválidos")
	fmt.Println()
	fmt.Println("Ambos afectan SOLO el bucle más cercano.")
	fmt.Println("Para afectar bucles externos: usar labels (siguiente archivo).")
}
