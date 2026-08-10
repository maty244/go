package main

import "fmt"

// =========================================================
// fmt.Stringer: LA INTERFAZ MÁS USADA DE GO
// =========================================================
// La librería estándar define esta interfaz en el paquete fmt:
//
//   type Stringer interface {
//       String() string
//   }
//
// Cualquier tipo que tenga un método String() string cumple esta
// interfaz. Y fmt.Println, fmt.Printf con %v, fmt.Sprintf, etc.
// la buscan automáticamente: si tu tipo la tiene, la usan en vez
// de imprimir los campos "en crudo".
//
// Esto es EXACTAMENTE lo que viste en Punteros 12 sin que lo
// llamáramos por su nombre (el método String() en el struct Punto).

// ─────────────────────────────────────────────────────────
// SIN Stringer: Go imprime los campos tal cual
// ─────────────────────────────────────────────────────────

type ProductoSinFormato struct {
	Nombre string
	Precio float64
}

// ─────────────────────────────────────────────────────────
// CON Stringer: controlás exactamente cómo se ve
// ─────────────────────────────────────────────────────────

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

// Al definir String(), Producto cumple fmt.Stringer.
// A partir de acá, CUALQUIER fmt.Println(p) usa esto.
func (p Producto) String() string {
	estado := "disponible"
	if p.Stock == 0 {
		estado = "agotado"
	}
	return fmt.Sprintf("%s ($%.2f) — %s", p.Nombre, p.Precio, estado)
}

func main() {
	fmt.Println("=== Sin Stringer: se ven los campos crudos ===")
	sinFormato := ProductoSinFormato{Nombre: "Mouse", Precio: 999.99}
	fmt.Println(sinFormato) // {Mouse 999.99}

	fmt.Println("\n=== Con Stringer: formato controlado ===")
	notebook := Producto{Nombre: "Notebook", Precio: 450000, Stock: 3}
	mouse := Producto{Nombre: "Mouse", Precio: 8500, Stock: 0}

	fmt.Println(notebook) // usa String() automáticamente
	fmt.Println(mouse)

	// %v y %s también disparan Stringer
	fmt.Printf("Con %%v: %v\n", notebook)
	fmt.Printf("Con %%s: %s\n", notebook)

	// %+v y %#v SÍ muestran la estructura interna (para debug)
	fmt.Printf("Con %%+v (debug): %+v\n", notebook)

	// ─────────────────────────────────────────────────────────
	// Stringer EN UN SLICE
	// ─────────────────────────────────────────────────────────
	// Cada elemento se imprime con su propio String().

	fmt.Println("\n=== Slice de productos ===")
	catalogo := []Producto{notebook, mouse, {Nombre: "Teclado", Precio: 12000, Stock: 15}}
	for _, p := range catalogo {
		fmt.Println(" -", p)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: logs legibles
	// ─────────────────────────────────────────────────────────
	// Cuando tu programa crece, vas a loggear structs todo el
	// tiempo. Sin Stringer, tus logs son ilegibles ({...} crudo).
	// Con Stringer, cada log dice exactamente lo que importa.

	fmt.Println("\n=== Caso real: log de un pedido ===")
	pedido := Pedido{ID: 4521, Cliente: "Sofía", Total: 15400.50, Estado: "confirmado"}
	fmt.Println("[LOG]", pedido)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	pct := "%" // evita que go vet confunda estas líneas con un Printf mal usado
	fmt.Println(`  fmt.Stringer     → interface { String() string }`)
	fmt.Println("  Si tu tipo la    → fmt.Println/Printf(" + pct + "v/" + pct + "s) la usan sola")
	fmt.Println(`  tiene              (no hace falta avisarle a fmt)`)
	fmt.Println("  " + pct + "+v              → ignora Stringer, muestra campos internos")
	fmt.Println(`  Uso típico       → logs y mensajes legibles para humanos`)
}

type Pedido struct {
	ID      int
	Cliente string
	Total   float64
	Estado  string
}

func (p Pedido) String() string {
	return fmt.Sprintf("Pedido #%d de %s — $%.2f [%s]", p.ID, p.Cliente, p.Total, p.Estado)
}
