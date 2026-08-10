package main

import "fmt"

// =========================================================
// DEFER
// =========================================================
// "defer" programa la ejecución de una función para cuando
// la función CONTENEDORA retorne (sin importar por qué retorna:
// return normal, panic, o final del cuerpo).
//
// Características clave:
//   - Los defers se ejecutan en orden LIFO (último en entrar, primero en salir)
//   - Los argumentos del defer se evalúan EN EL MOMENTO de la declaración,
//     no cuando se ejecuta.
//   - Es ideal para garantizar limpieza de recursos.
//
// Casos de uso reales:
//   - Cerrar archivos, conexiones a DB, clientes HTTP
//   - Liberar locks (mutexes)
//   - Medir tiempo de ejecución
//   - Logging de entrada/salida de funciones
//   - Recover después de un panic

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║           DEFER               ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// DEFER BÁSICO: se ejecuta al salir de la función
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Defer básico ===")
	fmt.Println("1: antes del defer")
	defer fmt.Println("3: defer → se ejecuta al final de main")
	fmt.Println("2: después del defer")
	// Al terminar main, imprime: "3: defer → se ejecuta al final"

	// ─────────────────────────────────────────────────────────
	// ORDEN LIFO: último en entrar, primero en salir
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Orden LIFO de múltiples defers ===")
	demostrarLIFO()

	// ─────────────────────────────────────────────────────────
	// LOS ARGUMENTOS SE EVALÚAN AHORA, NO DESPUÉS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Argumentos evaluados en el momento del defer ===")
	demostrarEvaluacionArgumenots()

	// ─────────────────────────────────────────────────────────
	// DEFER PARA SIMULAR APERTURA/CIERRE DE RECURSOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Defer para recursos ===")
	procesarArchivo("ventas.csv")

	// ─────────────────────────────────────────────────────────
	// DEFER CON RETURN: el defer se ejecuta ANTES de que el
	// valor de retorno llegue al llamador
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Defer y return ===")
	resultado := contarConDefer()
	fmt.Println("Resultado retornado:", resultado)

	// ─────────────────────────────────────────────────────────
	// MEDIR TIEMPO CON DEFER (patrón real muy usado)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Medir tiempo con defer ===")
	operacionLenta()

	// ─────────────────────────────────────────────────────────
	// DEFER EN LOOPS (¡CUIDADO!)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cuidado: defer en loops ===")
	// MAL: el defer no se ejecuta en cada iteración, sino al final de la función
	// Esto puede acumular muchos defers si el loop es grande.
	fmt.Println("Los defers en un loop se acumulan hasta que la función retorna:")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  defer del loop, i=%d\n", i) // se acumula!
	}
	fmt.Println("(los defers del loop se ejecutarán al terminar main)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("defer f()      → ejecuta f() cuando la función actual retorne")
	fmt.Println("Múltiples:     se ejecutan en orden LIFO (stack)")
	fmt.Println("Argumentos:    se evalúan YA, no cuando se ejecuta")
	fmt.Println("Uso principal: limpiar recursos (close, unlock, etc.)")
	fmt.Println("Patrón:        abrir recurso → defer cerrar recurso")
}

func demostrarLIFO() {
	defer fmt.Println("  defer 1 (primero declarado, último en ejecutarse)")
	defer fmt.Println("  defer 2")
	defer fmt.Println("  defer 3 (último declarado, primero en ejecutarse)")
	fmt.Println("  Cuerpo de la función")
	// Al retornar: 3, 2, 1
}

func demostrarEvaluacionArgumenots() {
	x := 10
	defer fmt.Println("  valor en defer (capturado ahora):", x) // captura x=10

	x = 99
	fmt.Println("  x modificado a:", x) // x ahora es 99
	// Pero el defer imprimirá 10, porque se capturó cuando se declaró
}

func procesarArchivo(nombre string) {
	fmt.Printf("  Abriendo archivo: %s\n", nombre)
	defer fmt.Printf("  Cerrando archivo: %s\n", nombre) // garantiza el cierre

	// Procesamiento del archivo...
	fmt.Println("  Leyendo datos...")
	fmt.Println("  Procesando registros...")
	fmt.Println("  Guardando resultados...")
	// Al retornar (sin importar si hay error), el archivo se cierra
}

func contarConDefer() int {
	i := 0
	defer func() {
		fmt.Println("  defer ejecutado, i en ese momento:", i)
	}()
	i = 5
	return i // defer se ejecuta ANTES de que 5 llegue al llamador
}

// Patrón real: medir cuánto tarda una función
func operacionLenta() {
	defer medirTiempo("operacionLenta")()
	// Simulamos trabajo pesado
	suma := 0
	for i := 0; i < 10_000_000; i++ {
		suma += i
	}
	fmt.Println("  Resultado:", suma)
}

func medirTiempo(nombre string) func() {
	fmt.Printf("  [%s] Iniciando...\n", nombre)
	return func() {
		// En producción: time.Since(start)
		fmt.Printf("  [%s] Finalizado\n", nombre)
	}
}
