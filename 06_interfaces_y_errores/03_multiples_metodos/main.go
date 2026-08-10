package main

import "fmt"

// =========================================================
// INTERFACES CON MÚLTIPLES MÉTODOS
// =========================================================
// Una interfaz puede pedir más de un método. Para cumplirla,
// el tipo necesita TODOS los métodos, no alcanza con algunos.
//
// Regla de oro en Go: las interfaces chicas son mejores.
// Es más fácil encontrar tipos que cumplan un contrato de 1-2
// métodos que uno de 10. La comunidad Go tiene un dicho:
// "cuanto más grande la interfaz, más débil la abstracción".

// ─────────────────────────────────────────────────────────
// INTERFAZ CON DOS MÉTODOS
// ─────────────────────────────────────────────────────────

type Empleado interface {
	Nombre() string
	Salario() float64
}

// ─────────────────────────────────────────────────────────
// UN TIPO QUE CUMPLE PARCIALMENTE NO CUENTA
// ─────────────────────────────────────────────────────────
// Si Cajero solo tuviera Nombre() pero no Salario(), NO cumpliría
// Empleado, y el compilador rechazaría cualquier intento de
// tratarlo como tal.

type Cajero struct {
	nombre      string
	salarioBase float64
	horasExtra  int
}

func (c Cajero) Nombre() string { return c.nombre }
func (c Cajero) Salario() float64 {
	return c.salarioBase + float64(c.horasExtra)*1500
}

type Gerente struct {
	nombre string
	sueldo float64
	bono   float64
}

func (g Gerente) Nombre() string   { return g.nombre }
func (g Gerente) Salario() float64 { return g.sueldo + g.bono }

// ─────────────────────────────────────────────────────────
// COMPONER INTERFACES: sumar contratos chicos
// ─────────────────────────────────────────────────────────
// Go permite armar una interfaz más grande combinando otras más
// chicas. Esto es preferible a escribir una interfaz gigante
// directamente: cada pieza se entiende y reutiliza por separado.

type Identificable interface {
	Nombre() string
}

type Remunerado interface {
	Salario() float64
}

// EmpleadoCompuesto exige TODOS los métodos de ambas interfaces.
// Es exactamente equivalente a Empleado, pero armado por partes.
type EmpleadoCompuesto interface {
	Identificable
	Remunerado
}

func main() {
	fmt.Println("=== Interfaz con múltiples métodos ===")

	empleados := []Empleado{
		Cajero{nombre: "Ana", salarioBase: 350000, horasExtra: 6},
		Gerente{nombre: "Carlos", sueldo: 700000, bono: 100000},
	}

	total := 0.0
	for _, e := range empleados {
		fmt.Printf("  %-8s → $%.2f\n", e.Nombre(), e.Salario())
		total += e.Salario()
	}
	fmt.Printf("Total nómina: $%.2f\n", total)

	// ─────────────────────────────────────────────────────────
	// USANDO LA INTERFAZ COMPUESTA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Interfaz compuesta (Identificable + Remunerado) ===")

	var ec EmpleadoCompuesto = Cajero{nombre: "Lucía", salarioBase: 400000}
	fmt.Printf("  %s gana $%.2f\n", ec.Nombre(), ec.Salario())

	// ─────────────────────────────────────────────────────────
	// CASO REAL: reportes que solo necesitan una parte del contrato
	// ─────────────────────────────────────────────────────────
	// Una función que solo arma una lista de nombres NO necesita
	// saber nada de salarios: le alcanza con Identificable.
	// Esto es "aceptar la interfaz más chica que te sirve".

	fmt.Println("\n=== Caso real: función que pide solo lo que necesita ===")
	fmt.Println("Nombres:", listarNombres(empleados[0], empleados[1]))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Interfaz con N métodos → el tipo necesita TODOS para cumplirla")
	fmt.Println("  Preferí interfaces CHICAS (1-2 métodos) sobre una gigante")
	fmt.Println("  Podés componer: type Grande interface { A; B }")
	fmt.Println("  Una función debería pedir la interfaz MÁS CHICA que le sirva")
}

// listarNombres solo necesita Identificable, aunque le pasemos
// valores que además cumplen Remunerado (Empleado completo).
func listarNombres(ids ...Identificable) []string {
	nombres := make([]string, 0, len(ids))
	for _, i := range ids {
		nombres = append(nombres, i.Nombre())
	}
	return nombres
}
