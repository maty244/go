package main

import (
	"fmt"
	"sync"
)

// =========================================================
// sync.RWMutex, sync.OnceValue Y EL MEMORY MODEL DE GO
// =========================================================
// 1. sync.RWMutex (Read/Write Mutex):
//    Un sync.Mutex normal (tema 06) bloquea a TODOS por igual,
//    incluso si dos goroutines solo quieren LEER sin modificar nada.
//    sync.RWMutex distingue entre lectores y escritores:
//      - mu.RLock() / mu.RUnlock(): permite que MUCHAS goroutines
//        lean al mismo tiempo (mientras nadie esté escribiendo).
//      - mu.Lock() / mu.Unlock(): pide acceso EXCLUSIVO para escribir
//        (espera a que terminen los lectores y bloquea nuevos lectores).
//
// 2. sync.Once / sync.OnceValue / sync.OnceValues:
//    Garantiza que una inicialización (cargar config, abrir un pool,
//    compilar una regex) se ejecute EXACTAMENTE UNA VEZ aunque
//    100 goroutines la llamen al mismo tiempo.

type Catalogo struct {
	mu    sync.RWMutex
	stock map[string]int
}

func (c *Catalogo) Consultar(nombre string) (int, bool) {
	c.mu.RLock() // múltiples lectores pueden entrar en paralelo
	defer c.mu.RUnlock()
	stock, existe := c.stock[nombre]
	return stock, existe
}

func (c *Catalogo) Reponer(nombre string, cantidad int) {
	c.mu.Lock() // escritor exclusivo: nadie más lee ni escribe durante esto
	defer c.mu.Unlock()
	c.stock[nombre] += cantidad
}

// sync.OnceValue devuelve una función que ejecuta el cálculo la
// primera vez que alguien la llama y guarda el resultado para todas
// las llamadas siguientes, de forma 100% segura entre goroutines.
var cargarConfiguracion = sync.OnceValue(func() map[string]string {
	fmt.Println("  (inicializando configuración por única vez...)")
	return map[string]string{"entorno": "desarrollo"}
})

func main() {
	fmt.Println("=== sync.RWMutex y sync.OnceValue ===")

	catalogo := &Catalogo{stock: map[string]int{"mouse": 2}}
	catalogo.Reponer("mouse", 3)
	stock, _ := catalogo.Consultar("mouse")
	fmt.Println("stock:", stock)

	// Aunque llamemos dos veces a cargarConfiguracion(), el mensaje
	// interno se imprime una sola vez:
	fmt.Println("config (1ra llamada):", cargarConfiguracion()["entorno"])
	fmt.Println("config (2da llamada):", cargarConfiguracion()["entorno"])

	fmt.Println("\nMemory model práctico:")
	fmt.Println("- Unlock publica cambios para un Lock/RLock posterior.")
	fmt.Println("- Un envío por channel sincroniza con la recepción correspondiente.")
	fmt.Println("- time.Sleep no sincroniza memoria: usá channels o sync.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  mu.RLock()/RUnlock()  → candado de lectura (varios lectores a la vez)")
	fmt.Println("  mu.Lock()/Unlock()    → candado de escritura (exclusivo)")
	fmt.Println("  sync.OnceValue(fn)    → inicializa algo una sola vez, thread-safe")
	fmt.Println("  Memory Model          → solo channels, Mutex, WaitGroup y atomic sincronizan")
}
