package main

import (
	"fmt"
	"time"
)

func main() {
	// =========================================================
	// CICLO FOR INFINITO
	// =========================================================
	// Un for sin condición corre para siempre. Solo se detiene
	// con break, return, os.Exit() o un panic.
	//
	// Sintaxis:  for { }
	//
	// En otros lenguajes equivale a:
	//   while (true) { }    (Java, C)
	//   while True:          (Python)
	//   loop { }             (Rust)
	//
	// Usos legítimos en Go:
	//   - Servidores que escuchan conexiones indefinidamente
	//   - Loops de juegos (game loop)
	//   - Workers que procesan tareas de una cola
	//   - CLIs que muestran un menú hasta que el usuario sale
	//   - Reintentos con lógica de salida compleja

	// ─────────────────────────────────────────────────────────
	// EJEMPLO BÁSICO
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== For infinito básico ===")

	cuenta := 0
	for {
		cuenta++
		fmt.Println("cuenta:", cuenta)
		if cuenta >= 3 {
			break // condición de salida
		}
	}
	fmt.Println("Salimos del for infinito")

	// ─────────────────────────────────────────────────────────
	// PATRÓN REAL: Menú de consola interactivo (simulado)
	// ─────────────────────────────────────────────────────────
	// En una app real, el usuario ingresaría la opción.
	// Aquí lo simulamos con un slice de "entradas".

	fmt.Println("\n=== Menú de consola (simulado) ===")

	entradas := []string{"1", "2", "3", "4"} // simula inputs del usuario
	idx := 0
	saldo := 1000.0

	for {
		if idx >= len(entradas) {
			break
		}
		opcion := entradas[idx]
		idx++

		fmt.Println("\n--- BANCO GO ---")
		fmt.Printf("Saldo actual: $%.2f\n", saldo)
		fmt.Println("1. Depositar")
		fmt.Println("2. Extraer")
		fmt.Println("3. Ver saldo")
		fmt.Println("4. Salir")
		fmt.Printf("Opción elegida (simulada): %s\n", opcion)

		switch opcion {
		case "1":
			saldo += 500
			fmt.Println("Depositado $500")
		case "2":
			if saldo >= 200 {
				saldo -= 200
				fmt.Println("Extraído $200")
			} else {
				fmt.Println("Saldo insuficiente")
			}
		case "3":
			fmt.Printf("Tu saldo es: $%.2f\n", saldo)
		case "4":
			fmt.Println("¡Hasta luego!")
			goto salir // salida controlada
		default:
			fmt.Println("Opción inválida")
		}
	}
salir:
	fmt.Printf("Saldo final: $%.2f\n", saldo)

	// ─────────────────────────────────────────────────────────
	// PATRÓN REAL: Worker que procesa una cola de trabajos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Worker con cola de trabajos ===")

	type Trabajo struct {
		ID    int
		Tarea string
	}

	trabajos := []Trabajo{
		{1, "enviar email"},
		{2, "generar reporte"},
		{3, "comprimir archivos"},
		{4, "sincronizar base de datos"},
	}

	procesados := 0
	colaIdx := 0

	for {
		// Verificamos si hay trabajos disponibles
		if colaIdx >= len(trabajos) {
			fmt.Println("No hay más trabajos, worker se detiene")
			break
		}

		trabajo := trabajos[colaIdx]
		colaIdx++

		fmt.Printf("  [Worker] Procesando trabajo #%d: %s\n", trabajo.ID, trabajo.Tarea)
		procesados++

		// Simulamos que ocasionalmente hay un trabajo inválido
		if trabajo.ID == 3 {
			fmt.Println("  [Worker] ⚠️ Trabajo #3 marcado como problemático, continuando")
			continue // saltamos al siguiente (aunque en for{} esto vuelve al inicio)
		}
	}
	fmt.Printf("Worker procesó %d trabajos\n", procesados)

	// ─────────────────────────────────────────────────────────
	// PATRÓN REAL: Polling / monitoreo periódico
	// ─────────────────────────────────────────────────────────
	// Muy común en sistemas que necesitan revisar algo regularmente.
	// En producción se usaría time.Sleep o un ticker de verdad.

	fmt.Println("\n=== Polling periódico (simulado, 3 ciclos) ===")

	ciclos := 0
	estado := "iniciando"
	estados := []string{"iniciando", "corriendo", "corriendo", "completado"}

	for {
		if ciclos >= len(estados) {
			break
		}
		estado = estados[ciclos]
		ciclos++

		fmt.Printf("  [Monitor] Estado: %s\n", estado)

		if estado == "completado" {
			fmt.Println("  [Monitor] Proceso finalizado, deteniendo monitoreo")
			break
		}

		// En producción: time.Sleep(5 * time.Second)
		_ = time.Millisecond // referenciamos time para el import
	}

	// ─────────────────────────────────────────────────────────
	// PATRÓN REAL: Game loop (loop de juego)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Game loop (simulado) ===")

	type EstadoJuego struct {
		vida      int
		nivel     int
		puntos    int
		terminado bool
	}

	juego := EstadoJuego{vida: 3, nivel: 1, puntos: 0}

	// Simulamos eventos del juego
	eventos := []string{"puntos", "puntos", "daño", "puntos", "daño", "daño", "puntos"}
	eIdx := 0

	for !juego.terminado { // también válido: for con condición booleana
		if eIdx >= len(eventos) {
			fmt.Println("  Fin de la demo")
			break
		}

		evento := eventos[eIdx]
		eIdx++

		switch evento {
		case "puntos":
			juego.puntos += 10
			fmt.Printf("  +10 puntos! Total: %d\n", juego.puntos)
		case "daño":
			juego.vida--
			fmt.Printf("  ¡Daño! Vida restante: %d\n", juego.vida)
			if juego.vida <= 0 {
				juego.terminado = true
				fmt.Println("  GAME OVER")
			}
		}
	}

	fmt.Printf("Puntaje final: %d puntos\n", juego.puntos)

	// ─────────────────────────────────────────────────────────
	// PATRÓN: For infinito con return (dentro de goroutines)
	// ─────────────────────────────────────────────────────────
	// En funciones que corren indefinidamente (servidores, workers),
	// el for{} corre hasta que el programa termina.
	// No se usa break, simplemente el proceso se cierra.

	fmt.Println("\n=== Patrón servidor (ejemplo conceptual) ===")
	fmt.Println("En un servidor real:")
	fmt.Println("")
	fmt.Println("  func servirConexiones() {")
	fmt.Println("      for {                           // acepta conexiones para siempre")
	fmt.Println("          conn, err := listener.Accept()")
	fmt.Println("          if err != nil { continue }  // error temporal, reintentar")
	fmt.Println("          go manejarConexion(conn)     // nueva goroutine por conexión")
	fmt.Println("      }")
	fmt.Println("  }")

	// ─────────────────────────────────────────────────────────
	// DIFERENCIAS ENTRE LAS FORMAS DE FOR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cuándo usar cada forma ===")
	fmt.Println("for i := 0; i < n; i++    → sabés cuántas veces (contador)")
	fmt.Println("for condicion { }          → no sabés cuántas, depende de estado")
	fmt.Println("for { }                    → loop eterno (servidor, menú, worker)")
	fmt.Println("")
	fmt.Println("for {} siempre necesita una salida:")
	fmt.Println("  break        → sale del for")
	fmt.Println("  return       → sale de toda la función")
	fmt.Println("  os.Exit(0)   → termina el programa")
	fmt.Println("  panic(...)   → termina con error")
}
