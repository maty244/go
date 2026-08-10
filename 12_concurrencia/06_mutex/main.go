package main

import (
	"fmt"
	"sync"
)

// =========================================================
// sync.Mutex: PROTEGER DATOS COMPARTIDOS
// =========================================================
// Los channels (temas 03-05) son geniales para COMUNICAR datos
// entre goroutines. Pero a veces varias goroutines necesitan
// LEER Y MODIFICAR la MISMA variable (un contador, un map, un
// balance de cuenta). Si lo hacen al mismo tiempo sin cuidado,
// se produce una CONDICIÓN DE CARRERA (race condition): el
// resultado final depende de un orden de ejecución que NO
// controlás, y suele dar resultados incorrectos e inconsistentes.
//
// sync.Mutex ("mutual exclusion") es un candado: solo UNA
// goroutine puede tener el candado a la vez. El patrón:
//
//   var mu sync.Mutex
//   mu.Lock()
//   // ... tocar el dato compartido ...
//   mu.Unlock() // (casi siempre con defer, 05_funciones_y_closures/04_defer)

// ContadorInseguro NO usa Mutex: vamos a ver que da resultados
// incorrectos con concurrencia.
type ContadorInseguro struct {
	valor int
}

func (c *ContadorInseguro) Incrementar() {
	c.valor++ // esto NO es una sola operación atómica: leer + sumar + guardar
}

// ContadorSeguro SÍ usa Mutex para proteger "valor".
type ContadorSeguro struct {
	mu    sync.Mutex
	valor int
}

func (c *ContadorSeguro) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *ContadorSeguro) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

func main() {
	fmt.Println("=== Contador SIN Mutex: resultado incorrecto ===")

	inseguro := &ContadorInseguro{}
	var wg sync.WaitGroup

	for range 1000 {
		wg.Go(func() {
			inseguro.Incrementar()
		})
	}
	wg.Wait()

	fmt.Println("Esperado: 1000 | Obtenido:", inseguro.valor)
	fmt.Println("(el número real puede variar en cada corrida, y rara vez da 1000)")

	// ─────────────────────────────────────────────────────────
	// EL MISMO CASO, CON Mutex: SIEMPRE correcto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Contador CON Mutex: siempre correcto ===")

	seguro := &ContadorSeguro{}
	var wg2 sync.WaitGroup

	for range 1000 {
		wg2.Go(func() {
			seguro.Incrementar()
		})
	}
	wg2.Wait()

	fmt.Println("Esperado: 1000 | Obtenido:", seguro.Valor())

	// ─────────────────────────────────────────────────────────
	// POR QUÉ PASA ESTO: c.valor++ NO ES UNA SOLA OPERACIÓN
	// ─────────────────────────────────────────────────────────
	// "valor++" en realidad son TRES pasos: leer valor, sumarle 1,
	// guardar el resultado. Si dos goroutines hacen esto AL MISMO
	// TIEMPO, pueden las dos leer el mismo valor viejo, sumar 1
	// cada una, y guardar el mismo resultado: se "pierde" un
	// incremento. Con miles de intentos simultáneos, esto pasa
	// todo el tiempo.

	fmt.Println("\n=== Por qué pasa: valor++ son 3 pasos, no 1 ===")
	fmt.Println("  1. Leer el valor actual")
	fmt.Println("  2. Sumarle 1")
	fmt.Println("  3. Guardar el resultado")
	fmt.Println("  Si dos goroutines hacen esto a la vez, uno de los incrementos se pierde")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  sync.Mutex           → candado: solo una goroutine adentro a la vez")
	fmt.Println("  mu.Lock()/Unlock()   → SIEMPRE en pares, casi siempre con defer")
	fmt.Println("  Sin Mutex + dato     → condición de carrera: resultados incorrectos")
	fmt.Println("  compartido")
	fmt.Println("  Regla general        → channels para COMUNICAR, Mutex para PROTEGER estado")
}
