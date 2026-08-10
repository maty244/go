package main

import (
	"fmt"
	"math"
)

// =========================================================
// STRUCTS (Estructuras)
// =========================================================
// Un struct es un tipo de dato que agrupa campos relacionados.
// Es la forma de Go de crear tipos complejos y personalizados.
// Es comparable a una "clase" en otros lenguajes, pero sin herencia.

// ─────────────────────────────────────────────────────────
// DEFINICIÓN DE STRUCTS (a nivel de paquete)
// ─────────────────────────────────────────────────────────

// Persona representa a una persona en el sistema.
type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
	Email    string
	Activo   bool
}

// Punto representa una coordenada en el plano cartesiano.
type Punto struct {
	X, Y float64 // múltiples campos del mismo tipo en una línea
}

// Dirección demuestra structs anidados.
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
	Pais   string
}

// Empleado tiene una Persona embebida y una Dirección anidada.
type Empleado struct {
	Persona   // embedding: herencia de composición
	Legajo    int
	Salario   float64
	Cargo     string
	Domicilio Direccion
}

// ─────────────────────────────────────────────────────────
// MÉTODOS
// ─────────────────────────────────────────────────────────
// En Go no hay clases, pero podemos agregar métodos a los tipos.
// Un método es una función con un "receiver" (receptor).
// El receiver puede ser por valor o por puntero.

// Método por VALOR: recibe una copia de Persona.
// Úsalo cuando el método NO modifica el struct.
func (p Persona) NombreCompleto() string {
	return p.Nombre + " " + p.Apellido
}

// Método por VALOR en Punto.
func (p Punto) Distancia(otro Punto) float64 {
	dx := p.X - otro.X
	dy := p.Y - otro.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Punto) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

// Método por PUNTERO: recibe el puntero al struct original.
// Úsalo cuando el método MODIFICA el struct o el struct es grande.
func (p *Persona) Cumpleanios() {
	p.Edad++ // modifica el campo del struct original
}

func (p *Persona) Desactivar() {
	p.Activo = false
}

// Método en Empleado
func (e Empleado) Info() string {
	return fmt.Sprintf("[%d] %s - %s ($%.0f/mes)",
		e.Legajo, e.NombreCompleto(), e.Cargo, e.Salario)
}

func (e *Empleado) Aumento(porcentaje float64) {
	e.Salario *= (1 + porcentaje/100)
}

func main() {
	// ─────────────────────────────────────────────────────────
	// CREAR STRUCTS
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Crear structs ===")

	// Forma 1: zero value (todos los campos en su zero value)
	var p1 Persona
	fmt.Printf("Zero value: %+v\n", p1) // %+v muestra los nombres de campos

	// Forma 2: struct literal con nombres (recomendado)
	p2 := Persona{
		Nombre:   "Ana",
		Apellido: "García",
		Edad:     28,
		Email:    "ana@ejemplo.com",
		Activo:   true,
	}
	fmt.Printf("Con nombres: %+v\n", p2)

	// Forma 3: struct literal posicional (menos recomendado, frágil)
	p3 := Persona{"Carlos", "López", 35, "carlos@ejemplo.com", true}
	fmt.Printf("Posicional: %v\n", p3)

	// Acceso a campos
	fmt.Printf("\nNombre: %s | Edad: %d\n", p2.Nombre, p2.Edad)

	// ─────────────────────────────────────────────────────────
	// MÉTODOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Métodos ===")

	fmt.Println("Nombre completo:", p2.NombreCompleto())
	fmt.Printf("Edad antes de cumpleaños: %d\n", p2.Edad)
	p2.Cumpleanios()
	fmt.Printf("Edad después: %d\n", p2.Edad)

	// Método en Punto
	origen := Punto{0, 0}
	punto := Punto{3, 4}
	fmt.Printf("\n%s → %s = distancia %.2f\n",
		origen.String(), punto.String(), origen.Distancia(punto))

	// ─────────────────────────────────────────────────────────
	// STRUCTS ANIDADOS Y EMBEDDING
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Structs anidados ===")

	empleado := Empleado{
		Persona: Persona{
			Nombre:   "Matias",
			Apellido: "Rodríguez",
			Edad:     30,
			Email:    "matias@empresa.com",
			Activo:   true,
		},
		Legajo:  1042,
		Salario: 150000.00,
		Cargo:   "Desarrollador Senior",
		Domicilio: Direccion{
			Calle:  "Av. Rivadavia 1234",
			Ciudad: "Buenos Aires",
			CP:     "1033",
			Pais:   "Argentina",
		},
	}

	// Gracias al embedding, podemos acceder a los campos de Persona directamente
	fmt.Println(empleado.Info())
	fmt.Println("Ciudad:", empleado.Domicilio.Ciudad)
	fmt.Println("Nombre (desde embedding):", empleado.Nombre) // en vez de empleado.Persona.Nombre

	// Métodos también se promueven por el embedding
	fmt.Println("Nombre completo:", empleado.NombreCompleto())

	empleado.Aumento(15)
	fmt.Printf("Salario con aumento: $%.0f\n", empleado.Salario)

	// ─────────────────────────────────────────────────────────
	// PUNTEROS A STRUCTS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Punteros a structs ===")

	// Go permite acceder a campos de un puntero a struct sin desreferenciar
	pPersona := &Persona{Nombre: "Elena", Edad: 25, Activo: true}
	pPersona.Edad++ // Go automáticamente hace (*pPersona).Edad++
	fmt.Println("Edad vía puntero:", pPersona.Edad)

	// Función que modifica un struct (necesita puntero)
	duplicarEdad(pPersona)
	fmt.Println("Edad duplicada:", pPersona.Edad)

	// ─────────────────────────────────────────────────────────
	// COMPARACIÓN DE STRUCTS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Comparación ===")

	a := Punto{1.0, 2.0}
	b := Punto{1.0, 2.0}
	c := Punto{3.0, 4.0}

	fmt.Printf("a(%v) == b(%v): %v\n", a, b, a == b) // true
	fmt.Printf("a(%v) == c(%v): %v\n", a, c, a == c) // false

	// Solo funciona si todos los campos son comparables (no slices/maps)

	// ─────────────────────────────────────────────────────────
	// STRUCT ANÓNIMO (inline, sin nombre)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct anónimo ===")

	// Útil para datos temporales o retornos de funciones internas
	config := struct {
		Host    string
		Port    int
		Debug   bool
		MaxConn int
	}{
		Host:    "localhost",
		Port:    8080,
		Debug:   true,
		MaxConn: 100,
	}
	fmt.Printf("Config: %+v\n", config)

	// ─────────────────────────────────────────────────────────
	// STRINGER INTERFACE (fmt.Stringer)
	// ─────────────────────────────────────────────────────────
	// Si un tipo tiene un método String() string, fmt lo usa automáticamente.
	p4 := Punto{5, 12}
	fmt.Println("\nPunto con Stringer:", p4) // usa String() automáticamente

	// ─────────────────────────────────────────────────────────
	// SLICE DE STRUCTS (muy común en Go)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slice de structs ===")

	personas := []Persona{
		{Nombre: "Ana", Edad: 28},
		{Nombre: "Carlos", Edad: 35},
		{Nombre: "Mia", Edad: 22},
		{Nombre: "Juan", Edad: 31},
	}

	for _, p := range personas {
		fmt.Printf("  %s: %d años\n", p.Nombre, p.Edad)
	}

	// Filtrar personas mayores de 25
	fmt.Println("\nMayores de 25:")
	for _, p := range personas {
		if p.Edad > 25 {
			fmt.Printf("  %s (%d)\n", p.Nombre, p.Edad)
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("type X struct { Campo tipo }  → define un struct")
	fmt.Println("func (x X) Método() {}        → método por valor")
	fmt.Println("func (x *X) Método() {}       → método por puntero (modifica)")
	fmt.Println("X{Campo: val}                 → literal con nombres (recomendado)")
	fmt.Println("Embedding:                    → composición en vez de herencia")
	fmt.Println("&X{...}                       → puntero al struct")
	_ = p3 // suprimir warning de no usado
}

func duplicarEdad(p *Persona) {
	p.Edad *= 2
}
