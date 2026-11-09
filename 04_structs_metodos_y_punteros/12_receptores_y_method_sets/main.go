package main

import (
	"fmt"
	"math"
)

// =========================================================
// RECEPTORES DE MÉTODOS: VALOR vs PUNTERO
// =========================================================
// En Go, los métodos se definen sobre tipos. El "receptor"
// es el parámetro especial que aparece antes del nombre:
//
//   func (t Tipo)  Metodo() { }   → receptor por VALOR
//   func (t *Tipo) Metodo() { }   → receptor por PUNTERO
//
// La diferencia es fundamental:
//
//   RECEPTOR VALOR   → recibe una COPIA del struct.
//                       No puede modificar el original.
//                       Se puede llamar sobre valor O puntero.
//
//   RECEPTOR PUNTERO → recibe el puntero al struct.
//                       SÍ puede modificar el destino.
//                       Puede llamarse sobre un valor DIRECCIONABLE;
//                       Go toma su dirección automáticamente en esa llamada.
//
// REGLA DE ORO (consistencia):
//   Si un método necesita *T → todos los métodos del tipo usan *T.
//   Mezclar sin razón genera confusión y bugs sutiles.

// ─────────────────────────────────────────────────────────
// EJEMPLO BASE: Contador
// ─────────────────────────────────────────────────────────

type Contador struct {
	valor int
	paso  int
}

// MAL: receptor por valor → NO modifica el original
func (c Contador) IncrementarMAL() {
	c.valor += c.paso // modifica la COPIA, no el original
}

// BIEN: receptor por puntero → SÍ modifica el original
func (c *Contador) Incrementar() {
	c.valor += c.paso
}

func (c *Contador) Resetear() {
	c.valor = 0
}

// Lectura pura: receptor valor está bien aquí
// PERO por consistencia usamos puntero igual (ver regla de oro)
func (c *Contador) Valor() int {
	return c.valor
}

func (c *Contador) String() string {
	return fmt.Sprintf("Contador{valor: %d, paso: %d}", c.valor, c.paso)
}

// ─────────────────────────────────────────────────────────
// CUÁNDO RECEPTOR VALOR TIENE SENTIDO: tipos pequeños/inmutables
// ─────────────────────────────────────────────────────────

type Punto struct {
	X, Y float64
}

// Operaciones que retornan un NUEVO punto (inmutabilidad funcional)
// → receptor valor es correcto: no modificamos el original
func (p Punto) Trasladar(dx, dy float64) Punto {
	return Punto{p.X + dx, p.Y + dy} // retorna nuevo punto
}

func (p Punto) Distancia(otro Punto) float64 {
	dx := p.X - otro.X
	dy := p.Y - otro.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Punto) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

// Operación que SÍ modifica → receptor puntero
func (p *Punto) Escalar(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// ─────────────────────────────────────────────────────────
// EJEMPLO REAL: Cuenta bancaria
// ─────────────────────────────────────────────────────────

type Cuenta struct {
	titular string
	saldo   float64
	activa  bool
}

func NuevaCuenta(titular string, deposito float64) *Cuenta {
	return &Cuenta{
		titular: titular,
		saldo:   deposito,
		activa:  true,
	}
}

// Métodos que MODIFICAN → *Cuenta
func (c *Cuenta) Depositar(monto float64) error {
	if !c.activa {
		return fmt.Errorf("cuenta inactiva")
	}
	if monto <= 0 {
		return fmt.Errorf("monto inválido: %.2f", monto)
	}
	c.saldo += monto
	return nil
}

func (c *Cuenta) Retirar(monto float64) error {
	if !c.activa {
		return fmt.Errorf("cuenta inactiva")
	}
	if monto > c.saldo {
		return fmt.Errorf("saldo insuficiente: %.2f < %.2f", c.saldo, monto)
	}
	c.saldo -= monto
	return nil
}

func (c *Cuenta) Cerrar() {
	c.activa = false
}

// Métodos de solo lectura → usamos *Cuenta por consistencia
func (c *Cuenta) Saldo() float64 {
	return c.saldo
}

func (c *Cuenta) EstaActiva() bool {
	return c.activa
}

func (c *Cuenta) String() string {
	estado := "activa"
	if !c.activa {
		estado = "inactiva"
	}
	return fmt.Sprintf("[%s | $%.2f | %s]", c.titular, c.saldo, estado)
}

// ─────────────────────────────────────────────────────────
// INTERFACES: receptor puntero y la interfaz
// ─────────────────────────────────────────────────────────
// Si un método tiene receptor *T, solo *T implementa la interfaz.
// Si el método tiene receptor T, tanto T como *T implementan la interfaz.

type Describible interface {
	Describir() string
}

type Animal struct {
	Nombre string
	Tipo   string
}

// Receptor valor → Animal Y *Animal implementan Describible
func (a Animal) Describir() string {
	return fmt.Sprintf("%s es un %s", a.Nombre, a.Tipo)
}

type Robot struct {
	Modelo  string
	Version int
}

// Receptor PUNTERO → solo *Robot implementa Describible
func (r *Robot) Describir() string {
	return fmt.Sprintf("Robot %s v%d", r.Modelo, r.Version)
}

func mostrarDescripcion(d Describible) {
	fmt.Println(" ", d.Describir())
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║    RECEPTORES DE MÉTODOS          ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// VALOR vs PUNTERO: contador
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Receptor valor vs puntero ===")

	c := Contador{valor: 0, paso: 5}
	fmt.Printf("Inicial: %s\n", &c)

	c.IncrementarMAL() // receptor valor → no cambia nada
	fmt.Printf("Después de IncrementarMAL: %s  (sin cambio)\n", &c)

	c.Incrementar() // receptor puntero → sí cambia
	c.Incrementar()
	fmt.Printf("Después de 2x Incrementar: %s\n", &c)

	c.Resetear()
	fmt.Printf("Después de Resetear: %s\n", &c)

	// ─────────────────────────────────────────────────────────
	// CONVERSIÓN AUTOMÁTICA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Conversión automática valor ↔ puntero ===")

	p1 := Punto{3.0, 4.0} // variable de tipo Punto (valor)

	// p1 es un valor, pero Escalar tiene receptor *Punto
	// Go automáticamente convierte: p1.Escalar(2) → (&p1).Escalar(2)
	p1.Escalar(2.0) // Go hace la conversión automáticamente
	fmt.Printf("p1 después de Escalar(2): %s\n", p1)

	p2 := &Punto{1.0, 2.0} // puntero
	// Trasladar tiene receptor Punto (valor), p2 es *Punto
	// Go desreferencia: p2.Trasladar(...) → (*p2).Trasladar(...)
	p3 := p2.Trasladar(5.0, 5.0)
	fmt.Printf("p2 original: %s | p3 trasladado: %s\n", p2, p3)
	fmt.Println("  → En llamadas, Go aplica & o * cuando el valor lo permite")
	fmt.Println("  → No ocurre al satisfacer interfaces ni con valores no direccionables")

	// ─────────────────────────────────────────────────────────
	// RECEPTOR VALOR PARA OPERACIONES FUNCIONALES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Receptor valor: operaciones funcionales (sin mutación) ===")

	origen := Punto{0, 0}
	destino := Punto{3, 4}

	dist := origen.Distancia(destino)
	fmt.Printf("Distancia de %s a %s: %.2f\n", origen, destino, dist)

	// La cadena de traslados NO modifica origen en ningún momento
	nuevo := origen.
		Trasladar(1, 1).
		Trasladar(2, 2).
		Trasladar(3, 3)
	fmt.Printf("origen intacto: %s | nuevo: %s\n", origen, nuevo)
	fmt.Println("  → Receptor valor permite cadenas funcionales sin mutación")

	// ─────────────────────────────────────────────────────────
	// CUENTA BANCARIA: todos receptores puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Cuenta bancaria: receptores puntero por consistencia ===")

	cuenta := NuevaCuenta("Lucía", 500.00)
	fmt.Printf("Inicial: %s\n", cuenta)

	operaciones := []struct {
		nombre string
		fn     func() error
	}{
		{"Depositar $300", func() error { return cuenta.Depositar(300) }},
		{"Retirar $100", func() error { return cuenta.Retirar(100) }},
		{"Retirar $1000 (fallo)", func() error { return cuenta.Retirar(1000) }},
	}

	for _, op := range operaciones {
		if err := op.fn(); err != nil {
			fmt.Printf("  ✗ %s: %v\n", op.nombre, err)
		} else {
			fmt.Printf("  ✓ %s: OK → %s\n", op.nombre, cuenta)
		}
	}

	cuenta.Cerrar()
	fmt.Printf("Cuenta cerrada: %s\n", cuenta)
	if err := cuenta.Depositar(50); err != nil {
		fmt.Printf("  ✗ Depositar en cuenta cerrada: %v\n", err)
	}

	// ─────────────────────────────────────────────────────────
	// INTERFACES Y RECEPTORES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Interfaces: receptor valor vs puntero ===")

	gato := Animal{Nombre: "Michi", Tipo: "gato"}
	perro := &Animal{Nombre: "Rex", Tipo: "perro"}
	robot := &Robot{Modelo: "X-9", Version: 3}

	// Animal (valor) también implementa Describible aunque Describir tenga receptor valor
	// *Animal también lo implementa
	// *Robot implementa porque Describir tiene receptor *Robot
	var describibles []Describible
	describibles = append(describibles, gato)  // Animal (valor) → OK
	describibles = append(describibles, perro) // *Animal → OK
	describibles = append(describibles, robot) // *Robot → OK
	// describibles = append(describibles, Robot{}) → ERROR: Robot no implementa Describible

	for _, d := range describibles {
		mostrarDescripcion(d)
	}
	fmt.Println("  → Animal (valor receptor): Animal y *Animal cumplen la interfaz")
	fmt.Println("  → Robot (puntero receptor): solo *Robot cumple la interfaz")

	// ─────────────────────────────────────────────────────────
	// REGLAS DE ORO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reglas de oro para receptores ===")
	fmt.Println("1. Si algún método necesita modificar el struct → usa *T en TODOS")
	fmt.Println("   (consistencia: no mezcles sin motivo)")
	fmt.Println()
	fmt.Println("2. Si el tipo es pequeño e inmutable (Punto, Color, par de ints)")
	fmt.Println("   y NINGÚN método modifica → receptor valor está bien")
	fmt.Println()
	fmt.Println("3. Con interfaz: receptor puntero → solo *T cumple la interfaz")
	fmt.Println("   receptor valor → T y *T cumplen la interfaz")
	fmt.Println()
	fmt.Println("4. Si el tipo tiene sync.Mutex u otro campo que no debe copiarse")
	fmt.Println("   → siempre receptor puntero (copiar un mutex es un bug)")
}
