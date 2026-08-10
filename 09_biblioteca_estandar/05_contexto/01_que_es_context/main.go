package main

import (
	"context"
	"fmt"
)

// =========================================================
// ¿QUÉ ES context.Context?
// =========================================================
// Problema que resuelve: en un programa real, una función A llama
// a B, que llama a C, que llama a D... y en algún momento el
// PEDIDO ORIGINAL se cancela (el usuario cerró la página, se
// venció un timeout, el servidor se está apagando). ¿Cómo le avisás
// a D, que está varias capas adentro, que ya no hace falta seguir
// trabajando?
//
// La respuesta de Go es context.Context: un valor que se pasa
// como PRIMER parámetro de una función en función, llevando dos
// cosas:
//   1. Una señal de cancelación/timeout (¿debería seguir trabajando?)
//   2. Valores de request (poco usado, no lo cubrimos en detalle)
//
// Vas a ver esta firma CONSTANTEMENTE en código Go real:
//
//   func hacerAlgo(ctx context.Context, otrosParams ...) error
//
// Por convención, ctx SIEMPRE es el primer parámetro.

// ─────────────────────────────────────────────────────────
// context.Background(): EL PUNTO DE PARTIDA
// ─────────────────────────────────────────────────────────
// Es un Context "vacío", sin cancelación ni timeout: el contexto
// raíz del que parten todos los demás. Se usa en main(), en tests,
// o como base para crear contextos más específicos (temas 02 y 03).

func main() {
	fmt.Println("=== context.Background(): el contexto raíz ===")

	ctx := context.Background()
	fmt.Println("Context:", ctx)
	fmt.Println("¿Tiene error?", ctx.Err())

	// ─────────────────────────────────────────────────────────
	// PASAR EL CONTEXT A TRAVÉS DE VARIAS FUNCIONES
	// ─────────────────────────────────────────────────────────
	// El patrón: cada función que podría necesitar cancelar su
	// trabajo recibe el ctx y se lo pasa a quien llama después.

	fmt.Println("\n=== Propagar el context entre funciones ===")
	procesarPedido(ctx, 4521)

	// ─────────────────────────────────────────────────────────
	// context.Context ES UNA INTERFAZ (conecta con 06_interfaces_y_errores/)
	// ─────────────────────────────────────────────────────────
	// Como error, context.Context es una interfaz de la librería
	// estándar. Tiene 4 métodos (Done, Err, Deadline, Value), pero
	// casi siempre solo vas a usar Done() y Err() (temas 02 y 03).

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  context.Context      → lleva cancelación/timeout a través de llamadas")
	fmt.Println("  context.Background() → el contexto raíz, sin cancelación")
	fmt.Println("  Convención           → ctx SIEMPRE es el primer parámetro")
	fmt.Println("  Se propaga           → pasándolo de función en función")
	fmt.Println("  Es una interfaz      → como error, con métodos (Done, Err, ...)")
}

func procesarPedido(ctx context.Context, id int) {
	fmt.Printf("  procesarPedido(%d): recibiendo el ctx\n", id)
	validarStock(ctx, id)
}

func validarStock(ctx context.Context, id int) {
	fmt.Printf("  validarStock(%d): el MISMO ctx sigue viajando\n", id)
	// En una función real, acá se chequearía ctx.Err() o se pasaría
	// a una llamada que sí lo use (una consulta a base de datos,
	// un request HTTP a otro servicio, etc.)
}
