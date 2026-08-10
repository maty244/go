package main

import (
	"errors"
	"fmt"
)

// =========================================================
// DISEÑAR ERRORES PERSONALIZADOS CON CAMPOS ÚTILES
// =========================================================
// Ya sabés CREAR un error personalizado (struct + Error() string).
// Este tema se enfoca en CÓMO DISEÑARLO BIEN: qué campos incluir,
// y cómo hacer que sea fácil de envolver y de comparar.

// ─────────────────────────────────────────────────────────
// MAL: un error personalizado que no aporta nada más que texto
// ─────────────────────────────────────────────────────────
// Si tu error solo tiene un mensaje, ni siquiera necesitás un
// struct: alcanza con errors.New o fmt.Errorf.

type errorGenericoMAL struct {
	mensaje string
}

func (e *errorGenericoMAL) Error() string { return e.mensaje }

// ─────────────────────────────────────────────────────────
// BIEN: un error personalizado que aporta DATOS estructurados
// ─────────────────────────────────────────────────────────
// La razón de ser de un error personalizado es cargar información
// que el código que lo recibe pueda USAR programáticamente (no
// solo mostrar como texto).

type ErrorLimiteExcedido struct {
	Recurso string
	Limite  int
	Pedido  int
}

func (e *ErrorLimiteExcedido) Error() string {
	return fmt.Sprintf("%s: límite de %d excedido (pediste %d)", e.Recurso, e.Limite, e.Pedido)
}

// Exceso es un MÉTODO extra, más allá de Error(): cualquier
// consumidor que recupere el tipo concreto puede usarlo.
func (e *ErrorLimiteExcedido) Exceso() int {
	return e.Pedido - e.Limite
}

// ─────────────────────────────────────────────────────────
// IMPLEMENTAR Unwrap() PARA ENVOLVER OTRO ERROR MANUALMENTE
// ─────────────────────────────────────────────────────────
// fmt.Errorf con %w es la forma más común de envolver, pero un
// struct de error también puede implementar su PROPIO Unwrap()
// para participar del mismo mecanismo (errors.Is/As lo detectan
// automáticamente).

var ErrCuotaAgotada = errors.New("cuota agotada")

type ErrorAPI struct {
	Endpoint string
	interno  error // el error real, guardado adentro
}

func (e *ErrorAPI) Error() string {
	return fmt.Sprintf("llamada a %s falló: %v", e.Endpoint, e.interno)
}

// Unwrap le dice a errors.Is/As/AsType "el error de adentro es este".
func (e *ErrorAPI) Unwrap() error {
	return e.interno
}

func llamarAPI(endpoint string) error {
	return &ErrorAPI{Endpoint: endpoint, interno: ErrCuotaAgotada}
}

func main() {
	fmt.Println("=== Error personalizado con datos útiles ===")

	err := &ErrorLimiteExcedido{Recurso: "requests/minuto", Limite: 100, Pedido: 137}
	fmt.Println("Error:", err)
	fmt.Println("Exceso:", err.Exceso())

	// ─────────────────────────────────────────────────────────
	// Unwrap PROPIO: errors.Is atraviesa structs personalizados
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Unwrap() propio en un struct ===")

	errAPI := llamarAPI("/productos")
	fmt.Println("Error:", errAPI)
	fmt.Println("¿En el fondo es ErrCuotaAgotada?", errors.Is(errAPI, ErrCuotaAgotada))

	// ─────────────────────────────────────────────────────────
	// CUÁNDO CREAR UN STRUCT vs USAR errors.New
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cuándo usar cada uno ===")
	fmt.Println("  errors.New/fmt.Errorf → mensaje simple, sin datos extra que extraer")
	fmt.Println("  struct + Error()      → necesitás CAMPOS o MÉTODOS que el código")
	fmt.Println("                           consumidor pueda usar (no solo mostrar texto)")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Error personalizado    → struct + Error() string")
	fmt.Println("  Diseñalo con campos    → que aporten info útil, no solo texto")
	fmt.Println("  Método Unwrap() error  → hace que errors.Is/As atraviesen tu struct")
	fmt.Println("  Métodos extra          → disponibles tras recuperar el tipo concreto")
}
