package main

import (
	"fmt"
	"sync"
)

type Catalogo struct {
	mu    sync.RWMutex
	stock map[string]int
}

func (c *Catalogo) Consultar(nombre string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	stock, existe := c.stock[nombre]
	return stock, existe
}

func (c *Catalogo) Reponer(nombre string, cantidad int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stock[nombre] += cantidad
}

var cargarConfiguracion = sync.OnceValue(func() map[string]string {
	return map[string]string{"entorno": "desarrollo"}
})

func main() {
	catalogo := &Catalogo{stock: map[string]int{"mouse": 2}}
	catalogo.Reponer("mouse", 3)
	stock, _ := catalogo.Consultar("mouse")
	fmt.Println("stock:", stock)
	fmt.Println("config:", cargarConfiguracion()["entorno"])

	fmt.Println("\nMemory model práctico:")
	fmt.Println("- Unlock publica cambios para un Lock/RLock posterior.")
	fmt.Println("- Un envío por channel sincroniza con la recepción correspondiente.")
	fmt.Println("- time.Sleep no sincroniza memoria: usá channels o sync.")
}
