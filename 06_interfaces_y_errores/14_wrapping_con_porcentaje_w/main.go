package main

import (
	"errors"
	"fmt"
)

// =========================================================
// ENVOLVER ERRORES CON %w (error wrapping)
// =========================================================
// Cuando una función de bajo nivel falla, la función que la llamó
// muchas veces quiere agregar CONTEXTO sin perder el error
// original. Por ejemplo: "falló la conexión" es útil, pero
// "no se pudo procesar el pedido #4521: falló la conexión" dice
// mucho más.
//
// fmt.Errorf con el verbo %w hace exactamente eso: crea un error
// NUEVO que "envuelve" al original. El de afuera tiene su propio
// mensaje, pero por dentro sigue conectado al original, y podés
// recuperarlo con errors.Unwrap, errors.Is o errors.As/AsType.
//
//   fmt.Errorf("contexto extra: %w", errorOriginal)
//
// OJO: %w es distinto de %v. Con %v, el error original se
// convierte en TEXTO y se pierde la conexión. Con %w, el error
// original queda "adentro", recuperable.

var ErrConexion = errors.New("no se pudo conectar a la base de datos")

func consultarDB(query string) error {
	// Simulamos que la conexión falla
	return ErrConexion
}

func procesarPedido(id int) error {
	if err := consultarDB("SELECT * FROM pedidos"); err != nil {
		// Envolvemos el error de bajo nivel con contexto de ESTA capa
		return fmt.Errorf("procesarPedido(%d): %w", id, err)
	}
	return nil
}

func main() {
	pct := "%" // evita que go vet confunda estas líneas con un Printf mal usado
	fmt.Println("=== Envolver con " + pct + "w vs " + pct + "v ===")

	err := procesarPedido(4521)
	fmt.Println("Error final:", err)

	// ─────────────────────────────────────────────────────────
	// A PESAR DEL WRAPPING, errors.Is SIGUE ENCONTRANDO EL ORIGINAL
	// ─────────────────────────────────────────────────────────
	// Esto es LA razón de ser de %w: aunque el mensaje visible sea
	// otro, el programa puede seguir preguntando "¿en el fondo,
	// esto es un error de conexión?"

	fmt.Println("\n=== errors.Is atraviesa el wrapping ===")
	fmt.Println("¿Es ErrConexion?", errors.Is(err, ErrConexion))

	// Comparar con %v: acá se PIERDE la conexión
	errConV := fmt.Errorf("procesarPedido(%d): %v", 4521, ErrConexion)
	fmt.Println("Con "+pct+"v en vez de "+pct+"w, ¿es ErrConexion?", errors.Is(errConV, ErrConexion))

	// ─────────────────────────────────────────────────────────
	// errors.Unwrap: sacar UNA capa del envoltorio
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== errors.Unwrap ===")
	original := errors.Unwrap(err)
	fmt.Println("Error desenvuelto:", original)
	fmt.Println("¿Es el mismo que ErrConexion?", original == ErrConexion)

	// ─────────────────────────────────────────────────────────
	// SE PUEDEN ENCADENAR VARIAS CAPAS
	// ─────────────────────────────────────────────────────────
	// Cada capa de tu programa puede agregar su propio contexto,
	// y errors.Is/As atraviesan TODAS las capas hasta encontrar
	// lo que buscan (o no encontrarlo).

	fmt.Println("\n=== Varias capas de wrapping ===")
	errCapa1 := fmt.Errorf("capa 1: %w", ErrConexion)
	errCapa2 := fmt.Errorf("capa 2: %w", errCapa1)
	errCapa3 := fmt.Errorf("capa 3: %w", errCapa2)

	fmt.Println("Error final:", errCapa3)
	fmt.Println("¿Sigue siendo ErrConexion en el fondo?", errors.Is(errCapa3, ErrConexion))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  fmt.Errorf("...: ` + pct + `w", err)  → envuelve err, agregando contexto`)
	fmt.Println("  " + pct + "w                          → mantiene la conexión (errors.Is/As lo ven)")
	fmt.Println("  " + pct + "v                          → convierte a texto, se pierde la conexión")
	fmt.Println(`  errors.Unwrap(err)          → saca UNA capa del envoltorio`)
	fmt.Println(`  Se pueden encadenar         → múltiples capas, todas atravesables`)
}
