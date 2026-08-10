package main

import "fmt"

// =========================================================
// TIPOS GENÉRICOS: STRUCTS QUE FUNCIONAN CON CUALQUIER TIPO
// =========================================================
// Los generics no son solo para funciones: también podés definir
// un STRUCT genérico, que guarda datos de un tipo T que se decide
// al usarlo. El caso más común: estructuras de datos (pilas,
// colas, listas) que deberían funcionar igual sin importar QUÉ
// guardan adentro.

// ─────────────────────────────────────────────────────────
// PILA GENÉRICA (Stack): último en entrar, primero en salir
// ─────────────────────────────────────────────────────────
// El [T any] va después del nombre del tipo. A partir de ahí, T
// se puede usar en los campos y en los métodos del tipo.

type Pila[T any] struct {
	items []T
}

// NuevaPila crea una pila vacía. También es una función genérica:
// el [T any] se repite porque una función libre no "hereda" el T
// del struct automáticamente.
func NuevaPila[T any]() *Pila[T] {
	return &Pila[T]{}
}

func (p *Pila[T]) Apilar(valor T) {
	p.items = append(p.items, valor)
}

// Desapilar devuelve el último elemento y un bool indicando si
// había algo (patrón similar al de un map: valor, ok).
func (p *Pila[T]) Desapilar() (T, bool) {
	var cero T // zero value de T, para el caso "no hay nada"
	if len(p.items) == 0 {
		return cero, false
	}
	ultimo := p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]
	return ultimo, true
}

func (p *Pila[T]) EstaVacia() bool {
	return len(p.items) == 0
}

func main() {
	fmt.Println("=== Pila genérica de int ===")

	pilaEnteros := NuevaPila[int]()
	pilaEnteros.Apilar(1)
	pilaEnteros.Apilar(2)
	pilaEnteros.Apilar(3)

	for !pilaEnteros.EstaVacia() {
		v, _ := pilaEnteros.Desapilar()
		fmt.Println("Desapilado:", v)
	}

	// ─────────────────────────────────────────────────────────
	// LA MISMA PILA, CON OTRO TIPO
	// ─────────────────────────────────────────────────────────
	// Sin generics, tendríamos que escribir PilaInt, PilaString,
	// PilaProducto... Con generics, es LA MISMA definición.

	fmt.Println("\n=== La misma Pila, ahora de string ===")

	pilaTextos := NuevaPila[string]()
	pilaTextos.Apilar("primero")
	pilaTextos.Apilar("segundo")

	v, ok := pilaTextos.Desapilar()
	fmt.Println("Desapilado:", v, "| había algo:", ok)

	// ─────────────────────────────────────────────────────────
	// CASO REAL: pila de un struct propio
	// ─────────────────────────────────────────────────────────
	// Funciona igual con tipos definidos por vos, no solo tipos
	// primitivos.

	fmt.Println("\n=== Pila de structs propios ===")

	type Accion struct {
		Descripcion string
	}

	historial := NuevaPila[Accion]()
	historial.Apilar(Accion{Descripcion: "Agregar Mouse al carrito"})
	historial.Apilar(Accion{Descripcion: "Aplicar cupón 10%"})

	ultimaAccion, _ := historial.Desapilar()
	fmt.Println("Deshaciendo:", ultimaAccion.Descripcion)

	// ─────────────────────────────────────────────────────────
	// DESAPILAR DE UNA PILA VACÍA: el "ok" avisa
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Pila vacía ===")
	vacia := NuevaPila[int]()
	_, ok = vacia.Desapilar()
	fmt.Println("¿Había algo para desapilar?", ok)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  type Pila[T any] struct {...}  → struct genérico")
	fmt.Println("  func (p *Pila[T]) Metodo()     → métodos usan el mismo T")
	fmt.Println("  NuevaPila[int](), [string]...  → misma definición, distinto T")
	fmt.Println("  Uso típico                     → estructuras de datos reutilizables")
}
