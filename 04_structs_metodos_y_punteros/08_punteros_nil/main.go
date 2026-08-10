package main

import "fmt"

// =========================================================
// NIL Y ERRORES TÍPICOS CON PUNTEROS
// =========================================================
// nil es el zero value de los punteros.
// Un puntero nil NO apunta a ningún lugar de memoria válido.
//
// REGLA DE ORO:
//   NUNCA desreferencies (*p) un puntero nil.
//   → Causa un panic en tiempo de ejecución: "nil pointer dereference"
//   → Es como ir a la dirección 0x0 a buscar algo: no existe.
//
// ¿Cuándo un puntero puede ser nil?
//   - Variable declarada pero no inicializada:  var p *int
//   - Función que retorna (*T, error) y falla:  return nil, err
//   - Campo de struct de tipo puntero no asignado

type Producto struct {
	ID     int
	Nombre string
	Precio float64
}

type Carrito struct {
	Cliente  string
	Producto *Producto // puede ser nil si no se eligió producto
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║    NIL Y ERRORES TÍPICOS          ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// EL CRASH MÁS COMÚN: nil pointer dereference
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El crash clásico (comentado por seguridad) ===")

	var p *int
	fmt.Printf("p = %v (nil)\n", p)
	fmt.Printf("¿Es nil? %v\n", p == nil)

	// ESTO CAUSA PANIC (descomenta para ver el crash):
	// fmt.Println(*p)  →  panic: runtime error: invalid memory address or nil pointer dereference

	// ─────────────────────────────────────────────────────────
	// CÓMO EVITARLO: siempre verificar nil antes de desreferenciar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrón correcto: verificar nil ===")

	var puntero *Producto
	if puntero != nil {
		fmt.Println("Producto:", puntero.Nombre) // seguro
	} else {
		fmt.Println("El puntero es nil, no hay producto")
	}

	// Ahora inicializamos y volvemos a intentar
	puntero = &Producto{1, "Notebook", 1500.00}
	if puntero != nil {
		fmt.Println("Producto:", puntero.Nombre)
	}

	// ─────────────────────────────────────────────────────────
	// FUNCIONES QUE RETORNAN nil
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Funciones que retornan nil ===")

	catalogo := []*Producto{
		{1, "Notebook", 1500.00},
		{2, "Mouse", 25.99},
		{3, "Teclado", 75.50},
	}

	// BIEN: verificar antes de usar
	if prod := buscarProducto(catalogo, 2); prod != nil {
		fmt.Printf("Encontrado: %s — $%.2f\n", prod.Nombre, prod.Precio)
	}

	if prod := buscarProducto(catalogo, 99); prod == nil {
		fmt.Println("Producto ID 99: no encontrado (nil)")
	}

	// MAL PATRÓN (crash si busca ID inexistente):
	// prod := buscarProducto(catalogo, 99)
	// fmt.Println(prod.Nombre)  ← PANIC si prod es nil!

	// ─────────────────────────────────────────────────────────
	// CAMPO PUNTERO DENTRO DE STRUCT
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Puntero nil dentro de struct ===")

	carrito := Carrito{
		Cliente:  "Ana",
		Producto: nil, // todavía no eligió producto
	}

	// Acceder a carrito.Producto.Nombre sin verificar → panic!
	fmt.Printf("Carrito de %s\n", carrito.Cliente)
	if carrito.Producto != nil {
		fmt.Printf("Producto elegido: %s\n", carrito.Producto.Nombre)
	} else {
		fmt.Println("Aún no eligió producto")
	}

	// Asignar producto
	carrito.Producto = &Producto{2, "Mouse", 25.99}
	if carrito.Producto != nil {
		fmt.Printf("Ahora eligió: %s\n", carrito.Producto.Nombre)
	}

	// ─────────────────────────────────────────────────────────
	// ERROR TÍPICO 1: asignar a través de puntero nil
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Error típico 1: asignar a nil ===")

	var pp *Producto
	fmt.Println("pp es nil, no se puede asignar a través de él.")
	fmt.Println("pp.Nombre = 'algo'  →  causaría PANIC")
	fmt.Println("Primero hay que inicializar: pp = &Producto{}")

	pp = &Producto{} // inicializar primero
	pp.Nombre = "Monitor"
	pp.Precio = 450.00
	fmt.Printf("Ahora sí: %s — $%.2f\n", pp.Nombre, pp.Precio)

	// ─────────────────────────────────────────────────────────
	// ERROR TÍPICO 2: Comparar puntero con su valor (confusión)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Error típico 2: confundir el puntero con el valor ===")

	pa := new(42) // Go 1.26+: new(valor) crea el puntero ya inicializado en 42

	fmt.Printf("pa   = %p  (dirección)\n", pa)
	fmt.Printf("*pa  = %d  (valor)\n", *pa)

	// No podés comparar un puntero con su valor directamente:
	// if pa == 42 { }  → ERROR de compilación: mismatched types

	// Correcto:
	if *pa == 42 {
		fmt.Println("El valor apuntado es 42 ✓")
	}

	// ─────────────────────────────────────────────────────────
	// ERROR TÍPICO 3: dangling pointer via closure (captura de loop var)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Error típico 3: capturar puntero de loop ===")

	numeros := []int{1, 2, 3, 4, 5}

	// MAL: todos los punteros apuntan a la MISMA variable del loop
	var punterosMAL []*int
	for _, n := range numeros {
		punterosMAL = append(punterosMAL, &n) // &n es siempre la misma variable!
	}
	fmt.Print("MAL (todos igual, el valor final de n): ")
	for _, p := range punterosMAL {
		fmt.Print(*p, " ")
	}
	fmt.Println()

	// BIEN: copiar en variable local dentro del loop
	var punterosBIEN []*int
	for _, n := range numeros {
		copia := n // variable nueva en cada iteración
		punterosBIEN = append(punterosBIEN, &copia)
	}
	fmt.Print("BIEN (cada uno tiene su propio valor): ")
	for _, p := range punterosBIEN {
		fmt.Print(*p, " ")
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// RECOVER: capturar el panic de nil pointer (caso extremo)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== recover: capturar un panic ===")
	resultado := operacionSegura(nil)
	fmt.Printf("Resultado seguro: '%s'\n", resultado)

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE REGLAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reglas para evitar nil panic ===")
	fmt.Println("1. Antes de *p, verificar: if p != nil { ... }")
	fmt.Println("2. Funciones que retornan *T deben documentar si pueden retornar nil")
	fmt.Println("3. Campos *T en structs: verificar antes de acceder")
	fmt.Println("4. Evitar capturar &loopVar directamente, usar una copia")
	fmt.Println("5. new(T) o &T{} son formas seguras de crear punteros no-nil")
}

func buscarProducto(catalogo []*Producto, id int) *Producto {
	for _, p := range catalogo {
		if p.ID == id {
			return p
		}
	}
	return nil // explícito: "no encontrado"
}

// operacionSegura usa recover para atrapar panics
func operacionSegura(p *Producto) (resultado string) {
	defer func() {
		if r := recover(); r != nil {
			resultado = fmt.Sprintf("panic recuperado: %v", r)
		}
	}()
	return p.Nombre // si p es nil, esto genera un panic que recover captura
}
