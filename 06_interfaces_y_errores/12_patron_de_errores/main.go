package main

import (
	"errors"
	"fmt"
)

// =========================================================
// MANEJO DE ERRORES EN GO
// =========================================================
// Go maneja errores de forma EXPLÍCITA, no con excepciones.
// No hay try/catch/finally. En cambio, las funciones retornan
// el error como un valor normal que el llamador DEBE verificar.
//
// Filosofía de Go: "Los errores son valores, no excepciones."
// Esto hace que el código sea más predecible y los errores
// más visibles.
//
// La interfaz error es simplemente:
//   type error interface {
//       Error() string
//   }

// ─────────────────────────────────────────────────────────
// CREAR ERRORES
// ─────────────────────────────────────────────────────────

// Forma 1: fmt.Errorf — crear un error con formato (la más común)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir %.0f por cero", a)
	}
	return a / b, nil
}

// Forma 2: errors.New — error sin formato
var ErrSinStock = errors.New("producto sin stock")
var ErrProductoInvalido = errors.New("producto inválido")

// ─────────────────────────────────────────────────────────
// ERRORES PERSONALIZADOS (tipos propios)
// ─────────────────────────────────────────────────────────
// Creamos un tipo de error propio con más información.
// Solo necesita implementar el método Error() string.

type ErrorCompra struct {
	Operacion string
	Motivo    string
	Codigo    int
}

// Implementamos la interfaz error
func (e ErrorCompra) Error() string {
	return fmt.Sprintf("[Código %d] Error en '%s': %s",
		e.Codigo, e.Operacion, e.Motivo)
}

// ─────────────────────────────────────────────────────────
// FUNCIONES DE LA TIENDA QUE PUEDEN FALLAR
// ─────────────────────────────────────────────────────────
type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

func buscarProducto(catalogo []Producto, id int) (Producto, error) {
	for _, p := range catalogo {
		if p.ID == id {
			return p, nil
		}
	}
	// Error descriptivo con fmt.Errorf
	return Producto{}, fmt.Errorf("buscarProducto: ID %d no existe en el catálogo", id)
}

func comprar(p Producto, cantidad int) (float64, error) {
	if cantidad <= 0 {
		return 0, ErrorCompra{
			Operacion: "comprar",
			Motivo:    fmt.Sprintf("cantidad %d no es válida, debe ser > 0", cantidad),
			Codigo:    400,
		}
	}
	if p.Stock == 0 {
		return 0, ErrSinStock // error centinela (variable pre-definida)
	}
	if p.Stock < cantidad {
		return 0, ErrorCompra{
			Operacion: "comprar",
			Motivo:    fmt.Sprintf("stock insuficiente: pedido %d, disponible %d", cantidad, p.Stock),
			Codigo:    409,
		}
	}
	return p.Precio * float64(cantidad), nil
}

func main() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║       MANEJO DE ERRORES       ║")
	fmt.Println("╚══════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// PATRÓN BÁSICO: if err != nil
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrón básico: if err != nil ===")

	if resultado, err := dividir(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", resultado)
	}

	if resultado, err := dividir(10, 0); err != nil {
		fmt.Println("Error capturado:", err)
	} else {
		fmt.Printf("resultado: %.1f\n", resultado)
	}

	// ─────────────────────────────────────────────────────────
	// ERRORES EN LA TIENDA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Errores en la tienda ===")

	catalogo := []Producto{
		{1, "Notebook", 1500.00, 3},
		{2, "Mouse", 25.99, 0}, // sin stock
		{3, "Teclado", 75.50, 5},
	}

	// Compra exitosa
	if p, err := buscarProducto(catalogo, 3); err != nil {
		fmt.Println("Error:", err)
	} else if total, err := comprar(p, 2); err != nil {
		fmt.Println("Error comprando:", err)
	} else {
		fmt.Printf("Compra exitosa: 2x %s = $%.2f\n", p.Nombre, total)
	}

	// Producto inexistente
	if _, err := buscarProducto(catalogo, 99); err != nil {
		fmt.Println("Error:", err)
	}

	// Sin stock
	if p, err := buscarProducto(catalogo, 2); err != nil {
		fmt.Println("Error:", err)
	} else if _, err := comprar(p, 1); err != nil {
		fmt.Println("Error:", err)
	}

	// Cantidad inválida
	if p, err := buscarProducto(catalogo, 1); err != nil {
		fmt.Println("Error:", err)
	} else if _, err := comprar(p, -1); err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────────────────
	// IDENTIFICAR EL TIPO DE ERROR (errors.Is / errors.As)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Identificar tipo de error ===")

	p, _ := buscarProducto(catalogo, 2) // Mouse sin stock
	_, err := comprar(p, 1)

	// errors.Is: compara con un error centinela específico
	if errors.Is(err, ErrSinStock) {
		fmt.Println("  → El producto está agotado, notificamos al depósito")
	}

	// errors.AsType (Go 1.26+): extrae el error a un tipo específico.
	// Es la versión genérica de errors.As. En código anterior a 1.26 vas a ver:
	//   var errCompra ErrorCompra
	//   if errors.As(errStock, &errCompra) { ... }
	// AsType hace lo mismo pero más directo: retorna (valor, ok).
	p2, _ := buscarProducto(catalogo, 1)
	_, errStock := comprar(p2, 100) // pedir más de lo disponible
	if errCompra, ok := errors.AsType[ErrorCompra](errStock); ok {
		fmt.Printf("  → Error personalizado capturado:\n")
		fmt.Printf("     Operación: %s\n", errCompra.Operacion)
		fmt.Printf("     Código:    %d\n", errCompra.Codigo)
		fmt.Printf("     Motivo:    %s\n", errCompra.Motivo)
	}

	// ─────────────────────────────────────────────────────────
	// ENVOLVER ERRORES CON %w (error wrapping)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Envolver errores con %w ===")

	// Cuando queremos agregar contexto a un error sin perder el original
	_, errOriginal := dividir(5, 0)
	errEnvuelto := fmt.Errorf("procesarPago: fallo al calcular la cuota: %w", errOriginal)
	fmt.Println("Error envuelto:", errEnvuelto)
	fmt.Println("¿Contiene el original?", errors.Is(errEnvuelto, errOriginal))

	// ─────────────────────────────────────────────────────────
	// ANTI-PATRÓN: NUNCA IGNORAR LOS ERRORES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Anti-patrones ===")
	fmt.Println("MAL:")
	fmt.Println("  resultado, _ := comprar(p, 100)  // ← ignorar el error es peligroso")
	fmt.Println("  // si hay error, resultado es 0 o el zero value")
	fmt.Println()
	fmt.Println("BIEN:")
	fmt.Println("  resultado, err := comprar(p, 100)")
	fmt.Println("  if err != nil { ... }")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("1. Las funciones que pueden fallar retornan (T, error)")
	fmt.Println("2. nil = sin error; non-nil = hay error")
	fmt.Println("3. Crear errores: fmt.Errorf() o errors.New()")
	fmt.Println("4. Error personalizado: tipo con método Error() string")
	fmt.Println("5. errors.Is() → comparar con error centinela")
	fmt.Println("6. errors.AsType[T]() → extraer tipo de error específico (antes: errors.As)")
	fmt.Println("7. %w en fmt.Errorf → envolver/preservar el error original")
	fmt.Println("8. NUNCA ignorar errores con _ en código de producción")
}
