package main

import (
	"fmt"
	"time"
)

// =========================================================
// GOROUTINES: LO MÁS DISTINTIVO DE GO
// =========================================================
// Una goroutine es una función que corre de forma CONCURRENTE
// con el resto del programa: no espera a que termine para seguir.
// Se lanza con la palabra clave "go" antes de una llamada.
//
//   go miFuncion()   // arranca concurrentemente, y el código SIGUE
//                    // a la línea siguiente sin esperarla
//
// Tienen un stack inicial pequeño (~2 KB, que crece y se achica solo) y suelen
// ser mucho más livianas que un thread del sistema operativo, pero no son gratis:
// consumen memoria, tiempo de scheduler y los recursos que abran.
//
// CONCURRENCIA VS PARALELISMO:
// - Concurrencia es estructurar el programa para lidiar con varias tareas a la vez.
// - Paralelismo es ejecutar varias tareas físicamente en el mismo instante en
//   distintos núcleos de CPU (el runtime de Go reparte las goroutines entre los
//   núcleos disponibles según GOMAXPROCS).

func saludarDespacio(nombre string) {
	time.Sleep(50 * time.Millisecond) // simula trabajo que tarda
	fmt.Println("  Hola,", nombre)
}

func main() {
	fmt.Println("=== Sin goroutine: todo en orden, uno espera al otro ===")

	saludarDespacio("Ana")
	saludarDespacio("Carlos")
	fmt.Println("(esto tardó el doble: cada llamada esperó a la anterior)")

	// ─────────────────────────────────────────────────────────
	// CON goroutine: el programa NO espera
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Con 'go': se lanza y el programa sigue ===")

	go saludarDespacio("Matias") // arranca... pero main() NO espera

	fmt.Println("Esta línea se imprime CASI seguro antes que el saludo de arriba")

	// ─────────────────────────────────────────────────────────
	// EL PROBLEMA: SI main() TERMINA, EL PROGRAMA TERMINA
	// ─────────────────────────────────────────────────────────
	// Esta es LA trampa número uno con goroutines: si main() se
	// termina, TODAS las goroutines en curso se cortan de golpe,
	// hayan terminado o no. Por eso, sin algo que las "espere"
	// (tema 02: sync.WaitGroup), el saludo de Matias podría no
	// llegar a imprimirse nunca.

	fmt.Println("\n=== El problema: main() no espera a las goroutines ===")
	fmt.Println("Si el programa terminara ACÁ, el saludo de Matias podría perderse")

	// Para este ejemplo, usamos Sleep como "parche" temporal (MAL,
	// mostrado a propósito): funciona, pero es frágil y lento.
	// El tema 02 muestra la forma CORRECTA de esperar.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("(esperamos con Sleep para que el saludo alcance a imprimirse)")

	// ─────────────────────────────────────────────────────────
	// VARIAS GOROUTINES A LA VEZ
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Varias goroutines concurrentes ===")

	nombres := []string{"Sofía", "Diego", "Lucía"}
	for _, nombre := range nombres {
		go saludarDespacio(nombre)
	}

	fmt.Println("Los 3 saludos se lanzaron 'a la vez' (orden de impresión no garantizado)")
	time.Sleep(100 * time.Millisecond) // de nuevo, un parche temporal

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  go miFuncion()      → la lanza concurrentemente, sin esperarla")
	fmt.Println("  Relativamente barata→ igual consume memoria y otros recursos")
	fmt.Println("  Paralelismo         → depende del runtime, CPU y trabajo")
	fmt.Println("  Orden NO garantizado → cuál termina primero puede variar")
	fmt.Println("  PELIGRO             → si main() termina, se cortan todas las goroutines")
	fmt.Println("  Sleep como 'espera'  → funciona para aprender, pero es un ANTI-PATRÓN")
	fmt.Println("  La forma correcta   → sync.WaitGroup (tema 02)")
}
