package main

import "fmt"

// =========================================================
// CHANNELS CON BUFFER: NO SIEMPRE HAY QUE ESPERAR
// =========================================================
// En el tema 03 vimos canales SIN buffer: enviar bloquea hasta
// que alguien recibe (sincrónico, "de mano en mano"). Un canal
// CON buffer tiene espacio para guardar varios valores SIN que
// alguien los esté recibiendo todavía.
//
//   ch := make(chan int, 3)  // buffer de capacidad 3
//
// Mientras el buffer no esté lleno, enviar NO bloquea. Recién
// cuando se llena, el próximo envío sí espera a que alguien saque
// algo. Es como un buzón con espacio para 3 cartas: podés meter
// hasta 3 sin que nadie las retire, pero la cuarta espera.

func main() {
	fmt.Println("=== Canal con buffer: enviar sin esperar (hasta llenarlo) ===")

	ch := make(chan string, 3)

	// Estos 3 envíos NO bloquean: hay espacio en el buffer.
	ch <- "pedido 1"
	ch <- "pedido 2"
	ch <- "pedido 3"
	fmt.Println("Los 3 pedidos entraron al buffer sin que nadie los reciba todavía")

	// Ahora los recibimos, en el mismo orden en que entraron (FIFO)
	fmt.Println("Recibiendo:", <-ch)
	fmt.Println("Recibiendo:", <-ch)
	fmt.Println("Recibiendo:", <-ch)

	// ─────────────────────────────────────────────────────────
	// len() Y cap(): cuánto hay adentro, y cuánto entra en total
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== len y cap de un channel ===")

	ch2 := make(chan int, 5)
	ch2 <- 10
	ch2 <- 20

	fmt.Println("len(ch2):", len(ch2)) // cuántos valores hay ESPERANDO ser recibidos
	fmt.Println("cap(ch2):", cap(ch2)) // capacidad total del buffer

	// ─────────────────────────────────────────────────────────
	// SI EL BUFFER SE LLENA, EL PRÓXIMO ENVÍO SÍ BLOQUEA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Buffer lleno: el próximo envío bloquea ===")

	chChico := make(chan int, 2)
	chChico <- 1
	chChico <- 2
	fmt.Println("Buffer lleno (len=2, cap=2)")

	// Si hiciéramos "chChico <- 3" en esta misma goroutine, se
	// colgaría (nadie está recibiendo). Por eso lo mandamos desde
	// una goroutine, y recibimos acá para liberar espacio.
	go func() {
		chChico <- 3 // esto espera a que main() reciba algo primero
		fmt.Println("  (goroutine) el 3 pudo entrar tras liberarse espacio")
	}()

	fmt.Println("Recibiendo para liberar espacio:", <-chChico)
	fmt.Println("Recibiendo:", <-chChico)
	fmt.Println("Recibiendo:", <-chChico)

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR BUFFER Y CUÁNDO NO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Sin buffer vs con buffer ===")
	fmt.Println("  Sin buffer (make(chan T))     → sincronización estricta, 'de mano en mano'")
	fmt.Println("  Con buffer (make(chan T, n))  → desacopla un poco emisor de receptor")
	fmt.Println("  Buffer MUY grande             → puede ocultar problemas de diseño")
	fmt.Println("  Regla general                 → empezá SIN buffer, agregalo si medís que hace falta")

	// ─────────────────────────────────────────────────────────
	// CASO REAL: cola de pedidos con capacidad limitada
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: cola de pedidos con capacidad máxima ===")

	colaPedidos := make(chan string, 3)

	pedidos := []string{"Pedido A", "Pedido B", "Pedido C"}
	for _, p := range pedidos {
		colaPedidos <- p
		fmt.Printf("  %s encolado (en cola: %d/%d)\n", p, len(colaPedidos), cap(colaPedidos))
	}

	close(colaPedidos)
	fmt.Println("Procesando la cola:")
	for p := range colaPedidos {
		fmt.Println("  Procesando:", p)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  make(chan T, n)  → canal con buffer de capacidad n")
	fmt.Println("  Enviar no bloquea → mientras el buffer no esté lleno")
	fmt.Println("  len(ch)          → cuántos valores hay esperando ser recibidos")
	fmt.Println("  cap(ch)          → capacidad total del buffer")
	fmt.Println("  Buffer lleno     → el próximo envío bloquea, como sin buffer")
}
