package main

import "fmt"

// =========================================================
// TYPE SWITCH: type assertion para varios tipos a la vez
// =========================================================
// Cuando necesitás distinguir entre VARIOS tipos concretos
// posibles (no solo uno), en vez de encadenar varios
// "if v, ok := i.(T)" existe el type switch:
//
//   switch v := i.(type) {
//   case TipoA:
//       // v ya es TipoA acá adentro
//   case TipoB:
//       // v ya es TipoB acá adentro
//   default:
//       // ningún caso matcheó
//   }
//
// Es sintaxis especial: "i.(type)" SOLO existe dentro de un
// switch. En cada "case", la variable v automáticamente toma
// el tipo concreto de ese caso.

type Forma interface {
	Area() float64
}

type Rectangulo struct{ Ancho, Alto float64 }

func (r Rectangulo) Area() float64 { return r.Ancho * r.Alto }

type Cuadrado struct{ Lado float64 }

func (c Cuadrado) Area() float64 { return c.Lado * c.Lado }

type Circulo struct{ Radio float64 }

func (c Circulo) Area() float64 { return 3.14159 * c.Radio * c.Radio }

// describir usa type switch para dar un mensaje distinto según
// el tipo concreto, algo que Area() (el único método del contrato)
// no puede resolver por sí solo.
func describir(f Forma) string {
	switch v := f.(type) {
	case Rectangulo:
		return fmt.Sprintf("Rectángulo de %vx%v (área %.2f)", v.Ancho, v.Alto, v.Area())
	case Cuadrado:
		return fmt.Sprintf("Cuadrado de lado %v (área %.2f)", v.Lado, v.Area())
	case Circulo:
		return fmt.Sprintf("Círculo de radio %v (área %.2f)", v.Radio, v.Area())
	default:
		return fmt.Sprintf("Forma desconocida (área %.2f)", f.Area())
	}
}

func main() {
	fmt.Println("=== Type switch ===")

	formas := []Forma{
		Rectangulo{Ancho: 4, Alto: 3},
		Cuadrado{Lado: 5},
		Circulo{Radio: 2},
	}

	for _, f := range formas {
		fmt.Println(" -", describir(f))
	}

	// ─────────────────────────────────────────────────────────
	// VARIOS TIPOS EN UN MISMO CASE
	// ─────────────────────────────────────────────────────────
	// Podés agrupar tipos que se tratan igual. Ojo: acá "v" NO
	// toma un tipo concreto (porque hay más de uno posible),
	// sigue siendo del tipo de la interfaz original.

	fmt.Println("\n=== Agrupar tipos en un case ===")
	for _, f := range formas {
		switch f.(type) {
		case Cuadrado, Rectangulo:
			fmt.Println(" - Es una figura de 4 lados")
		case Circulo:
			fmt.Println(" - Es una figura redonda")
		}
	}

	// ─────────────────────────────────────────────────────────
	// TYPE SWITCH SOBRE interface{} / any
	// ─────────────────────────────────────────────────────────
	// Un uso MUY común: procesar un valor de tipo "any" (lo vemos
	// a fondo en el próximo tema) y reaccionar según qué vino.

	fmt.Println("\n=== Type switch sobre 'any' ===")
	valores := []any{42, "hola", 3.14, true, Cuadrado{Lado: 2}}

	for _, v := range valores {
		procesar(v)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: procesar eventos de distinto tipo
	// ─────────────────────────────────────────────────────────
	// Un sistema de pedidos que recibe distintos "eventos" y actúa
	// distinto según cuál llegó — el patrón típico de un webhook
	// o una cola de mensajes.

	fmt.Println("\n=== Caso real: procesar eventos de un pedido ===")
	eventos := []Evento{
		EventoCreado{ID: 101},
		EventoCancelado{ID: 102, Motivo: "sin stock"},
		EventoEntregado{ID: 101},
	}
	for _, e := range eventos {
		procesarEvento(e)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  switch v := i.(type) { ... }  → distinguir entre varios tipos`)
	fmt.Println(`  case TipoA:                   → v ya es TipoA ahí adentro`)
	fmt.Println(`  case A, B:                    → agrupa tipos, v sigue siendo la interfaz`)
	fmt.Println(`  default:                      → ningún tipo matcheó`)
	fmt.Println(`  Muy usado con                 → any (interface vacía) y eventos`)
}

func procesar(v any) {
	switch valor := v.(type) {
	case int:
		fmt.Printf("  int: %d (doble: %d)\n", valor, valor*2)
	case string:
		fmt.Printf("  string: %q (len: %d)\n", valor, len(valor))
	case float64:
		fmt.Printf("  float64: %.2f\n", valor)
	case bool:
		fmt.Printf("  bool: %v\n", valor)
	default:
		fmt.Printf("  tipo no manejado: %v (%T)\n", valor, valor)
	}
}

type Evento interface {
	PedidoID() int
}

type EventoCreado struct{ ID int }

func (e EventoCreado) PedidoID() int { return e.ID }

type EventoCancelado struct {
	ID     int
	Motivo string
}

func (e EventoCancelado) PedidoID() int { return e.ID }

type EventoEntregado struct{ ID int }

func (e EventoEntregado) PedidoID() int { return e.ID }

func procesarEvento(e Evento) {
	switch ev := e.(type) {
	case EventoCreado:
		fmt.Printf("  Pedido #%d creado, notificar al depósito\n", ev.ID)
	case EventoCancelado:
		fmt.Printf("  Pedido #%d cancelado: %s\n", ev.ID, ev.Motivo)
	case EventoEntregado:
		fmt.Printf("  Pedido #%d entregado, pedir reseña\n", ev.ID)
	}
}
