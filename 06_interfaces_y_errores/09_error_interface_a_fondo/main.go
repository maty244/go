package main

import (
	"errors"
	"fmt"
)

// =========================================================
// error ES UNA INTERFAZ (a fondo)
// =========================================================
// En Go vas a usar el tipo "error" todo el tiempo (y lo profundizamos a
// partir de 12_patron_de_errores). Ahora que sabés qué es una interfaz
// podemos ver su definición REAL, que vive en el paquete builtin de Go:
//
//   type error interface {
//       Error() string
//   }
//
// Eso es TODO. "error" no es un tipo especial mágico: es una
// interfaz de UN solo método. Por eso:
//   - errors.New("algo") devuelve un valor que cumple esa interfaz
//   - fmt.Errorf(...) también
//   - y CUALQUIER struct tuyo con un método Error() string
//     también es, automáticamente, un error válido

// ─────────────────────────────────────────────────────────
// UN ERROR PERSONALIZADO ES SOLO UN TIPO CON Error() string
// ─────────────────────────────────────────────────────────

type ErrorStock struct {
	Producto string
	Pedido   int
	Stock    int
}

// Con este único método, ErrorStock cumple la interfaz "error".
func (e *ErrorStock) Error() string {
	return fmt.Sprintf("stock insuficiente para %s: pediste %d, hay %d",
		e.Producto, e.Pedido, e.Stock)
}

func comprar(producto string, cantidad, stock int) error {
	if cantidad > stock {
		// Devolvemos *ErrorStock, pero la firma dice "error":
		// esto funciona porque *ErrorStock CUMPLE error.
		return &ErrorStock{Producto: producto, Pedido: cantidad, Stock: stock}
	}
	return nil
}

func main() {
	fmt.Println("=== error es una interfaz de un método ===")

	err := comprar("Notebook", 5, 2)
	if err != nil {
		fmt.Println("Error:", err) // usa Error() automáticamente, como Stringer
	}

	// ─────────────────────────────────────────────────────────
	// POR QUÉ err == nil FUNCIONA COMO "SIN ERROR"
	// ─────────────────────────────────────────────────────────
	// Cuando una función retorna "nil" como error, en realidad
	// retorna una interfaz vacía: ningún tipo concreto adentro.
	// Por eso "if err != nil" es el patrón universal de Go.

	fmt.Println("\n=== nil como ausencia de error ===")
	err = comprar("Mouse", 1, 10)
	fmt.Println("¿Hay error?", err != nil) // false

	// ─────────────────────────────────────────────────────────
	// RECUPERAR EL TIPO CONCRETO DEL ERROR (type assertion)
	// ─────────────────────────────────────────────────────────
	// Como error es una interfaz, podés recuperar su tipo concreto con
	// type assertion (tema 05) o, mejor aún en código real, con
	// errors.AsType[T] / errors.As (que vemos a fondo en el tema 16).

	fmt.Println("\n=== Recuperar detalles con type assertion ===")
	err = comprar("Teclado", 20, 3)
	if errStock, ok := err.(*ErrorStock); ok {
		fmt.Printf("  Detalle: producto=%s, faltante=%d\n",
			errStock.Producto, errStock.Pedido-errStock.Stock)
	}

	// La forma moderna con errors.AsType[T] (Go 1.26+) hace lo
	// mismo pero también revisa errores ENVUELTOS con %w:
	if errStock, ok := errors.AsType[*ErrorStock](err); ok {
		fmt.Printf("  Con AsType: %s necesita %d más\n",
			errStock.Producto, errStock.Pedido-errStock.Stock)
	}

	// ─────────────────────────────────────────────────────────
	// POR QUÉ ESTO IMPORTA: podés diseñar TUS PROPIOS errores
	// ─────────────────────────────────────────────────────────
	// Como error es solo un contrato, tus errores personalizados
	// pueden llevar TODA la información que necesites (no solo un
	// texto), y quien los reciba puede decidir si le interesa esa
	// info extra (con type assertion) o solo tratarlo como error
	// genérico (con err.Error() o fmt.Println(err)).

	fmt.Println("\n=== Caso real: distintos consumidores del mismo error ===")

	err = comprar("Monitor", 10, 4)

	// Consumidor 1: solo le importa el mensaje (nivel API pública)
	fmt.Println("  Para el cliente:", err.Error())

	// Consumidor 2: le importa el detalle (nivel interno/logs)
	if es, ok := err.(*ErrorStock); ok {
		fmt.Printf("  Para el equipo de depósito: reponer %d unidades de %s\n",
			es.Pedido-es.Stock, es.Producto)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  error            → interface { Error() string }, nada más`)
	fmt.Println(`  errors.New/       → devuelven valores que cumplen error`)
	fmt.Println(`  fmt.Errorf`)
	fmt.Println(`  Struct + Error()  → tu propio tipo YA es un error válido`)
	fmt.Println(`  Type assertion    → recuperás el detalle cuando lo necesitás`)
	fmt.Println(`  err == nil        → "sin error" (interfaz sin tipo concreto adentro)`)
}
