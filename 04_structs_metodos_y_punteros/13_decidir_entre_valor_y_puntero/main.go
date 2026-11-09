package main

import (
	"fmt"
	"math"
)

// =========================================================
// ELEGIR ENTRE VALOR Y PUNTERO SEGÚN EL COMPORTAMIENTO
// =========================================================
// No hay una regla que prohíba punteros o los vuelva automáticamente mejores.
// Primero preguntá qué querés compartir o modificar. Después elegí la interfaz.
// El costo depende del tipo, el compilador y el uso: se mide, no se adivina.
//
// CUÁNDO SÍ usar punteros:
//   1. Necesitás modificar el original desde una función
//   2. El struct es grande y la copia es costosa
//   3. Campos opcionales (nil = ausente)
//   4. Estructuras recursivas (linked list, árbol)
//   5. Consistencia de receptores en métodos
//
// CUÁNDO NO usar punteros (la parte que se olvida):
//   1. Structs pequeños: bool, int, string, pocas palabras
//   2. Solo lectura: no hace falta *T para leer
//   3. Interfaces: casi nunca se usa un puntero a interfaz
//   4. Slices y maps: sus descriptores ya permiten compartir datos internos
//   5. Valores inmutables que no deberían cambiar
//   6. Para "optimizar" sin medir: puede ser PEOR por GC pressure

// ─────────────────────────────────────────────────────────
// CASO 1: STRUCT PEQUEÑO → pasar por valor es correcto
// ─────────────────────────────────────────────────────────

type Punto struct {
	X, Y float64 // 16 bytes — PEQUEÑO
}

// Bien: receptor valor para struct pequeño de solo lectura
func (p Punto) Distancia(otro Punto) float64 {
	dx, dy := p.X-otro.X, p.Y-otro.Y
	return math.Hypot(dx, dy)
}

func (p Punto) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

// Innecesario: usar *Punto solo para leer
// Esto agrega complejidad sin beneficio
func distanciaPuntero(p *Punto, otro *Punto) float64 {
	dx, dy := p.X-otro.X, p.Y-otro.Y
	return math.Hypot(dx, dy)
}

// ─────────────────────────────────────────────────────────
// CASO 2: SLICES Y MAPS — descriptores con datos compartidos
// ─────────────────────────────────────────────────────────

// Alternativa válida: acceder al descriptor del slice del llamador.
// No es un error del lenguaje; retornar el slice suele hacer más visible
// que cambió su longitud. Compará ambas opciones antes de elegir una.
func agregarMAL(s *[]int, v int) {
	*s = append(*s, v) // reemplaza el descriptor al que apunta s
}

// BIEN: retornar el slice actualizado (como hace append)
func agregarBIEN(s []int, v int) []int {
	return append(s, v)
}

// Para modificar entradas de un map no necesitás *map: la copia del descriptor
// permite acceder a la misma tabla.
func registrarMAL(m *map[string]int, k string, v int) {
	(*m)[k] = v // no necesitamos reemplazar el map para cambiar una entrada
}

func registrarBIEN(m map[string]int, k string, v int) {
	m[k] = v // directo, ambos descriptores observan la misma tabla
}

// ─────────────────────────────────────────────────────────
// CASO 3: PUNTERO INNECESARIO PARA "RETORNAR MÚLTIPLES"
// ─────────────────────────────────────────────────────────

// Go tiene retorno múltiple — no necesitás punteros para esto
// En C/C++ a veces se hacía: func(int *resultado1, int *resultado2)

// MAL (estilo C): usar punteros para "retornar" múltiples valores
func dividirMAL(a, b float64, cociente *float64, resto *float64) bool {
	if b == 0 {
		return false
	}
	*cociente = float64(int(a) / int(b))
	*resto = a - *cociente*b
	return true
}

// BIEN (idiomático Go): retorno múltiple + error
func dividirBIEN(a, b float64) (float64, float64, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("división por cero")
	}
	cociente := float64(int(a) / int(b))
	resto := a - cociente*b
	return cociente, resto, nil
}

// ─────────────────────────────────────────────────────────
// CASO 4: PUNTERO REAL Y NECESARIO — struct grande que muta
// ─────────────────────────────────────────────────────────

type Pedido struct {
	ID        int
	Cliente   string
	Items     []string
	Subtotal  float64
	Descuento float64
	Impuestos float64
	Total     float64
	Estado    string
	Notas     string
	Direccion string
	Metodo    string
	Tracking  string
}

// Aquí SÍ tiene sentido *Pedido:
//   - Struct grande (evitar copia)
//   - Modificamos el original
func aplicarDescuento(p *Pedido, pct float64) {
	p.Descuento = p.Subtotal * pct
	p.Total = p.Subtotal - p.Descuento + p.Impuestos
	p.Estado = "descuento aplicado"
}

func calcularImpuestos(p *Pedido, tasa float64) {
	p.Impuestos = p.Subtotal * tasa
	p.Total = p.Subtotal - p.Descuento + p.Impuestos
}

// Solo lectura: puede recibir por valor si el pedido fuera pequeño.
// Para structs grandes, usamos puntero por eficiencia.
func resumenPedido(p *Pedido) string {
	return fmt.Sprintf("Pedido #%d | %s | $%.2f | %s",
		p.ID, p.Cliente, p.Total, p.Estado)
}

// ─────────────────────────────────────────────────────────
// CASO 5: CONFUSIÓN COMÚN — pensar que puntero = más rápido
// ─────────────────────────────────────────────────────────

// Un puntero evita una copia, pero agrega indirección, posibilidad de nil y
// aliasing. El compilador usa escape analysis para decidir almacenamiento;
// escribir un puntero o retornarlo no garantiza por sí solo una asignación en
// heap. Si la motivación es rendimiento, medí con un benchmark.

// ─────────────────────────────────────────────────────────
// TABLA DE DECISIÓN FINAL
// ─────────────────────────────────────────────────────────

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║     ELEGIR VALOR O PUNTERO        ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// STRUCT PEQUEÑO: valor es más limpio
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct pequeño: valor > puntero ===")

	p1 := Punto{0, 0}
	p2 := Punto{3, 4}

	// Limpio: sin &, sin *, sin nil checks
	dist := p1.Distancia(p2)
	fmt.Printf("Distancia de %s a %s: %.1f\n", p1, p2, dist)

	// Comparar con la versión con punteros (más verbose, sin ventaja)
	dist2 := distanciaPuntero(&p1, &p2)
	fmt.Printf("Con punteros (mismo resultado, más verboso): %.1f\n", dist2)

	// ─────────────────────────────────────────────────────────
	// SLICES Y MAPS: descriptores con datos compartidos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slices y maps: descriptores copiados ===")

	nums := []int{1, 2, 3}
	nums = agregarBIEN(nums, 4)
	nums = agregarBIEN(nums, 5)
	fmt.Printf("Slice con agregarBIEN: %v\n", nums)

	conteos := make(map[string]int)
	registrarBIEN(conteos, "go", 100)
	registrarBIEN(conteos, "rust", 50)
	fmt.Printf("Map con registrarBIEN: %v\n", conteos)

	// ─────────────────────────────────────────────────────────
	// RETORNO MÚLTIPLE: no necesitás punteros para esto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Retorno múltiple: idiomático Go ===")

	// Mal (estilo C con punteros de salida)
	var cociente, resto float64
	ok := dividirMAL(17, 5, &cociente, &resto)
	if ok {
		fmt.Printf("dividirMAL: cociente=%.0f, resto=%.0f\n", cociente, resto)
	}

	// Bien (retorno múltiple de Go)
	c, r, err := dividirBIEN(17, 5)
	if err == nil {
		fmt.Printf("dividirBIEN: cociente=%.0f, resto=%.0f\n", c, r)
	}

	_, _, err = dividirBIEN(10, 0)
	if err != nil {
		fmt.Printf("dividirBIEN(10,0): %v\n", err)
	}

	// ─────────────────────────────────────────────────────────
	// STRUCT GRANDE: puntero SÍ tiene sentido
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct grande: puntero justificado ===")

	pedido := &Pedido{
		ID:       42,
		Cliente:  "María López",
		Items:    []string{"Notebook", "Mouse", "Teclado"},
		Subtotal: 1600.00,
		Estado:   "pendiente",
	}

	calcularImpuestos(pedido, 0.21)
	aplicarDescuento(pedido, 0.10)
	fmt.Println(resumenPedido(pedido))

	// ─────────────────────────────────────────────────────────
	// TABLA DE DECISIÓN FINAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla de decisión: ¿puntero o valor? ===")
	fmt.Println()
	fmt.Println("  SITUACIÓN                              USAR")
	fmt.Println("  ─────────────────────────────────────  ───────")
	fmt.Println("  Función necesita modificar el dato     *T")
	fmt.Println("  Copia medida como costosa              *T")
	fmt.Println("  Campo opcional (nil = ausente)         *T")
	fmt.Println("  Estructura recursiva                   *T")
	fmt.Println("  Consistencia con otros métodos *T      *T")
	fmt.Println()
	fmt.Println("  Struct pequeño, solo lectura           T")
	fmt.Println("  Slice (descriptor copiable)            []T")
	fmt.Println("  Map (descriptor copiable)              map[K]V")
	fmt.Println("  Interfaz (no usar *interface)          Interface")
	fmt.Println("  Retornar múltiples valores             (T1, T2, error)")
	fmt.Println("  Valor inmutable que no debe cambiar    T")

	fmt.Println()
	fmt.Println("=== Filosofía Go sobre punteros ===")
	fmt.Println("  Go te da punteros, pero no los impone.")
	fmt.Println("  La legibilidad y simplicidad son prioridad.")
	fmt.Println("  Antes de usar *T preguntate: ¿realmente lo necesito?")
	fmt.Println()
	fmt.Println("  Si usás un puntero, el lector del código asume que:")
	fmt.Println("    → El dato PUEDE ser nil")
	fmt.Println("    → El dato PUEDE ser modificado")
	fmt.Println("    → La identidad compartida forma parte del diseño")
	fmt.Println()
	fmt.Println("  Si ninguna de esas tres cosas aplica → usá valor.")
}
