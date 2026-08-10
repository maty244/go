package main

import "fmt"

// =========================================================
// PANIC Y RECOVER
// =========================================================
// Todo lo que viste hasta ahora (error, errors.Is/As, wrapping)
// es el manejo NORMAL de errores en Go: cosas que PUEDEN fallar,
// y el código llamador decide qué hacer.
//
// panic es distinto: es para situaciones EXCEPCIONALES, errores
// de PROGRAMACIÓN de los que el programa no puede "seguir
// normalmente" (un índice fuera de rango, un puntero nil que se
// desreferencia, dividir por cero con enteros). Cuando ocurre un
// panic, el programa empieza a "desenrollarse": termina la
// función actual, después la que la llamó, y así sucesivamente,
// ejecutando los defer que encuentre en el camino, hasta que
// alguien lo detiene con recover() o el programa termina.
//
// REGLA DE ORO: en Go, un error normal SIEMPRE se maneja con
// error, NUNCA con panic. panic es la excepción, no la regla.
// Vas a panic() casi nunca en tu propio código de negocio.

func main() {
	fmt.Println("=== panic: interrumpe la ejecución ===")

	demoPanicControlado()

	fmt.Println("\n(el programa sigue vivo después del panic recuperado)")

	// ─────────────────────────────────────────────────────────
	// CÓMO SE VE UN PANIC SIN RECOVER (comentado para no cortar
	// la ejecución de este archivo)
	// ─────────────────────────────────────────────────────────
	// var s []int
	// fmt.Println(s[5]) // panic: runtime error: index out of range [5] with length 0
	// Si esto se ejecuta y nadie hace recover(), el PROGRAMA ENTERO
	// termina abruptamente, imprimiendo el panic y el "stack trace".

	fmt.Println("\n=== recover(): atrapar un panic ===")
	resultado, err := dividirSeguro(10, 0)
	if err != nil {
		fmt.Println("Error controlado:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	resultado, err = dividirSeguro(10, 2)
	fmt.Println("10 / 2 =", resultado, "| error:", err)

	// ─────────────────────────────────────────────────────────
	// recover() SOLO FUNCIONA DENTRO DE defer
	// ─────────────────────────────────────────────────────────
	// Es la única forma en que funciona: llamar recover() fuera
	// de una función diferida no hace nada (devuelve nil).

	// ─────────────────────────────────────────────────────────
	// CASO REAL: proteger un servidor de que UN pedido roto tire
	// abajo TODO el proceso
	// ─────────────────────────────────────────────────────────
	// Esto es EXACTAMENTE lo que hacen los frameworks web: envuelven
	// cada request en un recover(), para que un bug en un handler
	// no tumbe el servidor completo. Lo vas a ver de nuevo en HTTP/.

	fmt.Println("\n=== Caso real: procesar pedidos sin que uno tumbe todo el lote ===")
	pedidos := []int{10, 5, 0, 8} // el 0 va a generar un panic simulado

	for _, cantidad := range pedidos {
		procesarPedidoSeguro(cantidad)
	}
	fmt.Println("El lote completo terminó de procesarse, a pesar del error")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  panic(valor)     → interrumpe todo, desenrolla ejecutando defers")
	fmt.Println("  recover()        → SOLO dentro de un defer, atrapa el panic activo")
	fmt.Println("  Sin recover      → el panic termina TODO el programa")
	fmt.Println("  Regla de oro     → error para lo esperable, panic para lo excepcional")
	fmt.Println("  Uso típico       → 'red de seguridad' en el punto de entrada (servidores, workers)")
}

func demoPanicControlado() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de un panic:", r)
		}
	}()

	fmt.Println("Antes del panic")
	panic("algo salió muy mal")
	// Esta línea NUNCA se ejecuta: panic corta el flujo acá mismo.
}

// dividirSeguro convierte un panic (división entera por cero) en
// un error normal, para que el llamador lo maneje con el patrón
// habitual (if err != nil), sin saber que por dentro hubo un panic.
func dividirSeguro(a, b int) (resultado int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("dividirSeguro: recuperado de panic: %v", r)
		}
	}()

	resultado = a / b // si b es 0, esto genera panic: integer divide by zero
	return resultado, nil
}

func procesarPedidoSeguro(cantidad int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Pedido con cantidad=%d falló (%v), seguimos con el resto\n", cantidad, r)
		}
	}()

	if cantidad == 0 {
		panic("cantidad no puede ser 0")
	}
	fmt.Printf("  Pedido de %d unidades procesado OK\n", cantidad)
}
