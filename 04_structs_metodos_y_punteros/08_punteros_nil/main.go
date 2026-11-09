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
//   → nil significa ausencia de destino. No lo uses como una dirección
//     numérica ni deduzcas dónde empieza la memoria física del proceso.
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
	// CÓMO EVITARLO: comprobar nil cuando el contrato admite ausencia
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
	// ERROR TÍPICO 3: confundir una copia de range con el elemento original
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Error típico 3: capturar puntero de loop ===")

	numeros := []int{1, 2, 3, 4, 5}

	// Go 1.22+: := declara una n distinta por iteración. Sus direcciones
	// siguen siendo válidas después del loop, pero apuntan a COPIAS de ints,
	// no a numeros[i]. No hay punteros colgantes en este ejemplo.
	var copias []*int
	for _, n := range numeros {
		copias = append(copias, &n)
	}
	fmt.Print("Copias por iteración: ")
	for _, p := range copias {
		fmt.Print(*p, " ")
	}
	fmt.Println()
	*copias[0] = 99
	fmt.Println("Cambiar la copia no cambia numeros:", numeros) // [1 2 3 4 5]

	// Si queremos los elementos reales, tomamos la dirección por índice.
	var elementos []*int
	for i := range numeros {
		elementos = append(elementos, &numeros[i])
	}
	*elementos[0] = 100
	fmt.Println("Cambiar el elemento sí cambia numeros:", numeros) // [100 2 3 4 5]

	// Con = reasignamos una variable exterior. En ese caso sí se comparte.
	var reutilizada int
	var compartidos []*int
	for _, reutilizada = range numeros {
		compartidos = append(compartidos, &reutilizada)
	}
	fmt.Print("Variable exterior reutilizada: ")
	for _, p := range compartidos {
		fmt.Print(*p, " ") // 5, 5, 5, 5, 5
	}
	fmt.Println()

	// ─────────────────────────────────────────────────────────
	// RECOVER: capturar el panic de nil pointer (caso extremo)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== recover: capturar un panic ===")
	// Experimento de panic/recover, no una forma de validar datos habituales.
	// Si nil es una entrada esperada, comprobalo antes de acceder al campo.
	resultado := demostrarRecover(nil)
	fmt.Printf("Resultado seguro: '%s'\n", resultado)

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE REGLAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reglas para evitar nil panic ===")
	fmt.Println("1. Si nil es posible, validalo antes de *p; documentá esa posibilidad")
	fmt.Println("2. Funciones que retornan *T deben documentar si pueden retornar nil")
	fmt.Println("3. Campos *T en structs: verificar antes de acceder")
	fmt.Println("4. &n con := apunta a una copia por iteración; &slice[i] al elemento real")
	fmt.Println("5. new(T) o &T{} son formas seguras de crear punteros no-nil")
}

func buscarProducto(catalogo []*Producto, id int) *Producto {
	for _, p := range catalogo {
		if p != nil && p.ID == id { // un slice de punteros también puede contener nil
			return p
		}
	}
	return nil // explícito: "no encontrado"
}

// demostrarRecover muestra el mecanismo; el manejo normal de ausencia usa if.
func demostrarRecover(p *Producto) (resultado string) {
	defer func() {
		if r := recover(); r != nil {
			resultado = fmt.Sprintf("panic recuperado: %v", r)
		}
	}()
	return p.Nombre // si p es nil, esto genera un panic que recover captura
}
