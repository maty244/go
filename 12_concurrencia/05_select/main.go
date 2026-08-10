package main

import (
	"fmt"
	"time"
)

// =========================================================
// select: ESPERAR EN VARIOS CHANNELS A LA VEZ
// =========================================================
// Ya usaste select "de prestado" en 09_biblioteca_estandar/05_contexto/02_with_timeout y 03, sin
// explicarlo a fondo. select es como un switch, pero cada "case"
// es una operación de channel (enviar o recibir). Se ejecuta el
// PRIMER caso que esté listo; si varios lo están a la vez, elige
// uno al azar entre ellos.
//
//   select {
//   case v := <-ch1:
//       // ch1 tenía algo listo
//   case ch2 <- valor:
//       // se pudo enviar a ch2
//   default:
//       // ningún canal estaba listo AHORA (no bloquea)
//   }

func main() {
	fmt.Println("=== select: el primero que esté listo, gana ===")

	rapido := make(chan string)
	lento := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		rapido <- "respuesta rápida"
	}()
	go func() {
		time.Sleep(150 * time.Millisecond)
		lento <- "respuesta lenta"
	}()

	select {
	case msg := <-rapido:
		fmt.Println("Ganó:", msg)
	case msg := <-lento:
		fmt.Println("Ganó:", msg)
	}

	// ─────────────────────────────────────────────────────────
	// select CON default: NO BLOQUEAR
	// ─────────────────────────────────────────────────────────
	// Sin default, select ESPERA a que algún caso esté listo. Con
	// default, si NINGÚN canal está listo en este instante, ejecuta
	// default inmediatamente, sin esperar nada.

	fmt.Println("\n=== select con default: chequear sin bloquear ===")

	canal := make(chan int)

	select {
	case v := <-canal:
		fmt.Println("Recibido:", v)
	default:
		fmt.Println("No había nada esperando, seguimos sin bloquear")
	}

	// ─────────────────────────────────────────────────────────
	// select EN UN LOOP: procesar mientras algo esté disponible
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== select en loop: drenar un canal sin bloquear al final ===")
	// Trampa clásica: un "break" SIN label adentro de un select
	// solo corta el select, NO el for que lo rodea (quedaría en
	// loop infinito). Por eso usamos un label ("Drenaje:") sobre
	// el for, igual que en 02_control_de_flujo/08_labels.

	resultados := make(chan int, 3)
	resultados <- 10
	resultados <- 20
	resultados <- 30

Drenaje:
	for {
		select {
		case v := <-resultados:
			fmt.Println("  Procesando:", v)
		default:
			fmt.Println("  No queda nada más por ahora")
			break Drenaje // corta el for con label, visto en 02_control_de_flujo/08_labels
		}
	}

	// ─────────────────────────────────────────────────────────
	// select + time.After: TIMEOUT PARA UNA OPERACIÓN CON CHANNEL
	// ─────────────────────────────────────────────────────────
	// Patrón MUY común (ya lo viste en 09_biblioteca_estandar/05_contexto/02_with_timeout, pero con
	// context.WithTimeout en vez de time.After directo): "esperá
	// esto, pero si tarda más de X, seguí sin colgarte".

	fmt.Println("\n=== select + time.After: timeout manual ===")

	datosLentos := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		datosLentos <- "estos datos tardaron demasiado"
	}()

	select {
	case dato := <-datosLentos:
		fmt.Println("Llegó:", dato)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Se agotó el tiempo de espera, seguimos sin los datos")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  select { case <-ch: ... }  → espera el PRIMER canal que esté listo")
	fmt.Println("  Varios listos a la vez     → elige uno al azar entre ellos")
	fmt.Println("  default                    → ejecuta si NINGUNO está listo (no bloquea)")
	fmt.Println("  select + time.After        → patrón clásico de timeout")
	fmt.Println("  Sin default                → select BLOQUEA hasta que algo esté listo")
}
