package main

import (
	"cmp"
	"fmt"
)

// =========================================================
// CONSTRAINTS (RESTRICCIONES): DEFINIR QUÉ TIPOS ACEPTAR
// =========================================================
// En el tema anterior escribimos "int | float64" directo en la
// función. Cuando esa lista de tipos se repite en varias funciones,
// conviene darle un NOMBRE: eso es un "constraint" (restricción),
// y se define como una interfaz especial.

// ─────────────────────────────────────────────────────────
// UN CONSTRAINT PROPIO: una interfaz que lista tipos permitidos
// ─────────────────────────────────────────────────────────

type Numero interface {
	int | int64 | float32 | float64
}

func Promedio[T Numero](valores []T) float64 {
	var total T
	for _, v := range valores {
		total += v
	}
	return float64(total) / float64(len(valores))
}

// ─────────────────────────────────────────────────────────
// comparable: EL CONSTRAINT PREDEFINIDO PARA == Y !=
// ─────────────────────────────────────────────────────────
// "comparable" es un constraint que ya viene en Go: acepta
// cualquier tipo que soporte == y != (números, strings, bools,
// structs sin slices/maps adentro, etc.)

func Contiene[T comparable](lista []T, buscado T) bool {
	for _, v := range lista {
		if v == buscado {
			return true
		}
	}
	return false
}

// ─────────────────────────────────────────────────────────
// cmp.Ordered: EL CONSTRAINT DE LA LIBRERÍA ESTÁNDAR PARA <, >
// ─────────────────────────────────────────────────────────
// El paquete "cmp" (parte de la librería estándar) define
// Ordered: cualquier tipo que soporte <, <=, >, >= (todos los
// números y strings). Evita reinventar la rueda cada vez que
// necesitás una función que compara con mayor/menor.

func Maximo[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("=== Constraint propio: Numero ===")

	notas := []int{8, 9, 7, 10, 6}
	fmt.Printf("Promedio de notas: %.2f\n", Promedio(notas))

	precios := []float64{1500.50, 2300.00, 899.99}
	fmt.Printf("Promedio de precios: %.2f\n", Promedio(precios))

	// ─────────────────────────────────────────────────────────
	// comparable
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== comparable ===")

	productos := []string{"Mouse", "Teclado", "Monitor"}
	fmt.Println(`Contiene("Teclado"):`, Contiene(productos, "Teclado"))
	fmt.Println(`Contiene("Notebook"):`, Contiene(productos, "Notebook"))

	// ─────────────────────────────────────────────────────────
	// cmp.Ordered
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== cmp.Ordered (paquete estándar \"cmp\") ===")

	fmt.Println("Maximo(3, 9)          =", Maximo(3, 9))
	fmt.Println("Maximo(2.5, 1.1)      =", Maximo(2.5, 1.1))
	fmt.Println(`Maximo("pera","uva")  =`, Maximo("pera", "uva"))

	// cmp también trae funciones ya hechas, para no reinventarlas:
	fmt.Println("cmp.Compare(3, 9) =", cmp.Compare(3, 9)) // -1: 3 < 9

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Constraint            → interfaz que lista qué tipos son válidos para T")
	fmt.Println("  type X interface{...} → constraint propio, reusable en varias funciones")
	fmt.Println("  comparable            → predefinido, para tipos que soportan == / !=")
	fmt.Println("  cmp.Ordered           → del paquete \"cmp\", para tipos con < > <= >=")
}
