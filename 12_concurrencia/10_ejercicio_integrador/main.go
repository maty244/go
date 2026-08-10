package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// =========================================================
// EJERCICIO INTEGRADOR: PROCESAR PEDIDOS DEL KIOSCO EN PARALELO
// =========================================================
// Combinamos TODO lo visto en 12_concurrencia/: un worker pool que
// procesa pedidos, un Mutex protegiendo las estadísticas
// compartidas (nada de channels ahí, a propósito, para mostrar
// cuándo conviene Mutex sobre channels), y un context con
// timeout para todo el lote.

type Pedido struct {
	ID       int
	Producto string
	Monto    Dinero
}

// Dinero guarda centavos. Para negocio evitamos float64: 0.1 no tiene una
// representación binaria exacta y los redondeos acumulados son peligrosos.
type Dinero int64

// EstadisticasDelDia se actualiza desde VARIOS workers al mismo
// tiempo: necesita Mutex (tema 06), no channels, porque es
// ESTADO COMPARTIDO que se lee y modifica, no un mensaje puntual.
type EstadisticasDelDia struct {
	mu              sync.Mutex
	pedidosOK       int
	pedidosFallidos int
	totalFacturado  Dinero
}

func (e *EstadisticasDelDia) registrarExito(monto Dinero) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.pedidosOK++
	e.totalFacturado += monto
}

func (e *EstadisticasDelDia) registrarFallo() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.pedidosFallidos++
}

func (e *EstadisticasDelDia) resumen() (int, int, Dinero) {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.pedidosOK, e.pedidosFallidos, e.totalFacturado
}

// worker procesa pedidos del channel hasta que se cierra o el
// context se cancela (combina el patrón del tema 08 con el 09).
func worker(ctx context.Context, id int, pedidos <-chan Pedido, stats *EstadisticasDelDia) {
	for {
		select {
		case <-ctx.Done():
			return
		case p, abierto := <-pedidos:
			if !abierto {
				return
			}
			select {
			case <-ctx.Done():
				stats.registrarFallo()
				return
			case <-time.After(20 * time.Millisecond):
				fmt.Printf("  worker %d: procesó pedido #%d (%s, %d centavos)\n", id, p.ID, p.Producto, p.Monto)
				stats.registrarExito(p.Monto)
			}
		}
	}
}

func main() {
	fmt.Println("=== KIOSCO DIGITAL: procesamiento concurrente de pedidos ===")

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	pedidosCh := make(chan Pedido, 10)
	stats := &EstadisticasDelDia{}
	var wg sync.WaitGroup

	const cantidadWorkers = 3
	for w := 1; w <= cantidadWorkers; w++ {
		wg.Go(func() { worker(ctx, w, pedidosCh, stats) })
	}

	pedidos := []Pedido{
		{ID: 1, Producto: "Alfajor", Monto: 80000},
		{ID: 2, Producto: "Gaseosa", Monto: 120000},
		{ID: 3, Producto: "Notebook", Monto: 45000000},
		{ID: 4, Producto: "Mouse", Monto: 850000},
		{ID: 5, Producto: "Teclado", Monto: 1200000},
		{ID: 6, Producto: "Monitor", Monto: 9500000},
		{ID: 7, Producto: "Auriculares", Monto: 1500000},
	}

	for _, p := range pedidos {
		pedidosCh <- p
	}
	close(pedidosCh)

	wg.Wait()

	ok, fallidos, total := stats.resumen()

	fmt.Println("\n=== Resumen del lote ===")
	fmt.Printf("  Pedidos procesados: %d\n", ok)
	fmt.Printf("  Pedidos fallidos:   %d\n", fallidos)
	fmt.Printf("  Total facturado:    %d centavos\n", total)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 12_concurrencia/ ===")
	fmt.Println("  Worker pool (08)       → 3 workers fijos procesan 7 pedidos")
	fmt.Println("  channels (03-04)       → cola de pedidos y coordinación")
	fmt.Println("  select (05)            → cada worker respeta el context o procesa")
	fmt.Println("  context.WithTimeout (09) → límite de tiempo para TODO el lote")
	fmt.Println("  sync.Mutex (06)        → estadísticas compartidas, protegidas")
	fmt.Println("  sync.WaitGroup (02)    → esperar a que los 3 workers terminen")
	fmt.Println("\nEste es el mismo patrón que vas a usar para procesar requests")
	fmt.Println("concurrentes en un servidor HTTP real.")
}
