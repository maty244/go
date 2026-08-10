package main

import "fmt"

// =========================================================
// CHANNELS: CÓMO SE COMUNICAN LAS GOROUTINES
// =========================================================
// Ya usaste channels "de prestado" en 09_biblioteca_estandar/04_tiempo/05_timers_y_sleep (time.After) y
// 09_biblioteca_estandar/05_contexto/ (ctx.Done()), sin explicarlos a fondo. Ahora sí: un
// channel es un CONDUCTO tipado por el que las goroutines mandan
// y reciben valores, de forma segura (sin necesitar un Mutex
// para esto en particular, tema 06).
//
// Frase clásica de Go: "no te comuniques compartiendo memoria,
// compartí memoria comunicándote". Un channel es la herramienta
// para eso: en vez de que dos goroutines toquen la MISMA variable
// (peligroso, tema 06), se pasan valores a través del channel.
//
//   ch := make(chan int)  // canal de int, SIN buffer (tema 04: con buffer)
//   ch <- 5                // ENVIAR: mete 5 en el canal
//   valor := <-ch           // RECIBIR: saca un valor del canal
//
// Un canal SIN buffer es "sincrónico": el envío BLOQUEA hasta que
// alguien esté recibiendo, y viceversa. Es como pasarse un objeto
// de mano en mano: las dos manos tienen que estar ahí a la vez.

func main() {
	fmt.Println("=== Canal básico: enviar y recibir ===")

	ch := make(chan string)

	// Si hiciéramos "ch <- \"hola\"" acá mismo, sin una goroutine
	// que lo reciba, el programa se COLGARÍA para siempre (nadie
	// del otro lado). Por eso el envío va en una goroutine.
	go func() {
		ch <- "hola desde la goroutine"
	}()

	mensaje := <-ch // main() se bloquea acá hasta que llegue algo
	fmt.Println("Recibido:", mensaje)

	// ─────────────────────────────────────────────────────────
	// UN CHANNEL COMO "RESULTADO" DE UNA GOROUTINE
	// ─────────────────────────────────────────────────────────
	// Patrón MUY común: una goroutine hace un cálculo y manda el
	// resultado por un channel, en vez de un return normal (que no
	// tendría a quién devolverle nada, al correr en paralelo).

	fmt.Println("\n=== Recibir el resultado de un cálculo ===")

	resultado := make(chan int)
	go calcularTotal([]float64{800, 1200, 450000}, resultado)

	total := <-resultado
	fmt.Println("Total calculado:", total)

	// ─────────────────────────────────────────────────────────
	// CERRAR UN CHANNEL: close()
	// ─────────────────────────────────────────────────────────
	// Indica "no va a llegar NADA más por acá". Quien recibe puede
	// detectarlo con la forma de dos valores: valor, ok := <-ch
	// (ok=false significa "canal cerrado y vacío").

	fmt.Println("\n=== Cerrar un channel ===")

	numeros := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numeros <- i
		}
		close(numeros) // avisamos: no mandamos nada más
	}()

	for {
		n, ok := <-numeros
		if !ok {
			fmt.Println("  Canal cerrado, no hay más valores")
			break
		}
		fmt.Println("  Recibido:", n)
	}

	// ─────────────────────────────────────────────────────────
	// for range SOBRE UN CHANNEL: la forma idiomática
	// ─────────────────────────────────────────────────────────
	// Hace EXACTAMENTE lo del loop de arriba, pero más corto: recibe
	// hasta que el canal se cierra, y termina el loop solo.

	fmt.Println("\n=== for range sobre un channel ===")

	letras := make(chan string)
	go func() {
		for _, l := range []string{"a", "b", "c"} {
			letras <- l
		}
		close(letras)
	}()

	for l := range letras {
		fmt.Println("  Letra:", l)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  make(chan T)      → canal sin buffer, para valores de tipo T")
	fmt.Println("  ch <- valor       → enviar (bloquea hasta que alguien reciba)")
	fmt.Println("  valor := <-ch     → recibir (bloquea hasta que alguien envíe)")
	fmt.Println("  close(ch)         → avisa 'no mando nada más'")
	fmt.Println("  v, ok := <-ch     → ok=false si el canal está cerrado y vacío")
	fmt.Println("  for v := range ch → recibe hasta que el canal se cierra")
}

// El parámetro "chan<- int" es un canal de SOLO ENVÍO: esta función
// puede mandar valores por él, pero no recibir. Es un chequeo extra
// del compilador (documenta la intención); "chan int" también
// funcionaría, pero "chan<- int" dice explícitamente "esta función
// solo escribe acá".
func calcularTotal(precios []float64, resultado chan<- int) {
	total := 0.0
	for _, p := range precios {
		total += p
	}
	resultado <- int(total)
}
