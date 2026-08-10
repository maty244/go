package main

import "fmt"

// =========================================================
// PUNTEROS: OPERADORES * Y &
// =========================================================
// Un puntero es una variable que GUARDA UNA DIRECCIÓN de memoria.
// En vez de guardar un valor (42, "hola"), guarda una dirección
// que APUNTA al lugar donde vive ese valor.
//
// Dos operadores fundamentales:
//
//   &  (ampersand / "address-of")
//      → obtiene la DIRECCIÓN de una variable
//      → "dame la dirección de memoria de esta variable"
//      → resultado: un puntero (*T)
//
//   *  (asterisco / "dereference")
//      → accede al VALOR al que apunta un puntero
//      → "ir a esa dirección y dame el valor que hay ahí"
//      → resultado: el valor del tipo T
//
// TAMBIÉN: * en la declaración del TIPO significa "puntero a"
//   var p *int  →  p es un puntero a int (guarda una dirección de int)

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       OPERADORES * Y &            ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// OPERADOR & : obtener la dirección
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Operador & (address-of) ===")

	edad := 25

	// &edad → "la dirección de memoria de la variable 'edad'"
	// El resultado es de tipo *int (puntero a int)
	punteroEdad := &edad

	fmt.Printf("edad          = %d\n", edad)
	fmt.Printf("&edad         = %p  (dirección de memoria)\n", &edad)
	fmt.Printf("punteroEdad   = %p  (guarda la misma dirección)\n", punteroEdad)
	fmt.Printf("tipo:           %T\n", punteroEdad)

	// ─────────────────────────────────────────────────────────
	// OPERADOR * : desreferenciar (acceder al valor)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Operador * (dereference) ===")

	// *punteroEdad → "ir a la dirección y dame el valor"
	valorApuntado := *punteroEdad

	fmt.Printf("punteroEdad   = %p  (es una dirección)\n", punteroEdad)
	fmt.Printf("*punteroEdad  = %d  (el valor EN esa dirección)\n", valorApuntado)

	// Diagrama mental:
	// edad          → [ 0xc000014080 ] → [ 25 ]
	// punteroEdad   → [ 0xc000014090 ] → [ 0xc000014080 ]
	//                                         ↓
	//                                    [ 25 ]  (es lo que da *punteroEdad)

	// ─────────────────────────────────────────────────────────
	// MODIFICAR EL ORIGINAL A TRAVÉS DEL PUNTERO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Modificar a través del puntero ===")

	n := 10
	p := &n

	fmt.Printf("n antes:   %d\n", n)
	fmt.Printf("*p antes:  %d\n", *p)

	*p = 99 // asignación MEDIANTE el puntero → modifica 'n'

	fmt.Printf("n después: %d  (modificado a través del puntero!)\n", n)
	fmt.Printf("*p después: %d\n", *p) // apunta al mismo lugar, ve el nuevo valor

	// ─────────────────────────────────────────────────────────
	// DECLARAR UN PUNTERO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Formas de declarar punteros ===")

	// Forma 1: var (inicializa en nil)
	var pNil *int
	fmt.Printf("var pNil *int → %v (es nil, no apunta a nada)\n", pNil)

	// Forma 2: & a una variable existente
	x := 42
	pX := &x
	fmt.Printf("pX := &x  → %p (apunta a x)\n", pX)

	// Forma 3: new() crea una variable y retorna un puntero. El compilador
	// decide si su almacenamiento queda en stack o heap.
	// Tiene dos variantes:
	//   new(int) → clásica: recibe un TIPO, el valor arranca en su zero value (0)
	//   new(42)  → Go 1.26+: recibe un VALOR, el puntero ya nace inicializado
	pNew := new(int) // crea un int en 0 y retorna *int
	*pNew = 77
	fmt.Printf("pNew := new(int) → %p → valor: %d\n", pNew, *pNew)

	pDirecto := new(42) // Go 1.26+: equivale a las dos líneas de arriba en una
	fmt.Printf("pDirecto := new(42) → %p → valor: %d\n", pDirecto, *pDirecto)

	// Forma 4: & a un literal de struct (muy común en Go)
	type Punto struct{ X, Y int }
	pPunto := &Punto{3, 4} // crea el valor y retorna *Punto
	fmt.Printf("pPunto := &Punto{3,4} → %p → valor: %v\n", pPunto, *pPunto)

	// ─────────────────────────────────────────────────────────
	// * EN EL TIPO vs * EN LA EXPRESIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== * en el tipo vs * en la expresión ===")

	var s *string // * en el TIPO: "s es un puntero a string"
	texto := "hola"
	s = &texto // & en expresión: "dame la dirección de texto"
	fmt.Printf("*s = '%s'  (usamos * para obtener el valor)\n", *s)

	*s = "modificado" // * en expresión: "escribir en la dirección"
	fmt.Printf("texto = '%s'  (el original cambió!)\n", texto)

	// ─────────────────────────────────────────────────────────
	// PUNTEROS EN FUNCIONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Punteros en funciones ===")

	saldo := 1000.0
	fmt.Printf("Saldo antes:  $%.2f\n", saldo)

	depositar(&saldo, 500)
	fmt.Printf("Saldo después de depositar $500:  $%.2f\n", saldo)

	if err := retirar(&saldo, 200); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Saldo después de retirar $200:   $%.2f\n", saldo)
	}

	if err := retirar(&saldo, 5000); err != nil {
		fmt.Println("Error al retirar $5000:", err)
	}

	// ─────────────────────────────────────────────────────────
	// PUNTERO A PUNTERO (raro pero existe)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Puntero a puntero (**T) ===")

	valor := 42
	pv := &valor // *int: apunta a valor
	ppv := &pv   // **int: apunta al puntero

	fmt.Printf("valor     = %d\n", valor)
	fmt.Printf("*pv       = %d  (un nivel de indirección)\n", *pv)
	fmt.Printf("**ppv     = %d  (dos niveles de indirección)\n", **ppv)

	**ppv = 999 // modifica a través de dos niveles
	fmt.Printf("valor después: %d\n", valor)
	// En la práctica, **T se usa muy poco. Puede confundir.

	// ─────────────────────────────────────────────────────────
	// ACCESO A CAMPOS DE STRUCT A TRAVÉS DE PUNTERO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Acceder a campos de struct via puntero ===")

	type Persona struct {
		Nombre string
		Edad   int
	}

	persona := Persona{"Ana", 28}
	pp := &persona

	// Go permite acceder a campos directamente sin desreferenciar manualmente
	fmt.Printf("pp.Nombre = '%s'  (Go hace (*pp).Nombre automáticamente)\n", pp.Nombre)
	fmt.Printf("(*pp).Edad = %d  (equivalente explícito)\n", (*pp).Edad)

	// Modificar campo a través del puntero
	pp.Edad = 29 // equivale a: (*pp).Edad = 29
	fmt.Printf("persona.Edad después: %d\n", persona.Edad)

	// ─────────────────────────────────────────────────────────
	// RESUMEN VISUAL
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen de operadores ===")
	fmt.Println("  &variable    → dirección de memoria (tipo: *T)")
	fmt.Println("  *puntero     → valor en esa dirección (tipo: T)")
	fmt.Println("  var p *T     → declara p como puntero a T (p = nil)")
	fmt.Println("  new(T)       → crea T en zero value y retorna *T")
	fmt.Println("  new(expr)    → Go 1.26: crea el valor inicial y retorna su puntero")
	fmt.Println("  &T{...}      → crea el literal y retorna *T")
	fmt.Println("  stack/heap   → lo decide el compilador mediante escape analysis")
	fmt.Println("  p.Campo      → Go lo transforma en (*p).Campo")
}

func depositar(saldo *float64, monto float64) {
	*saldo += monto // modifica el saldo original
}

func retirar(saldo *float64, monto float64) error {
	if *saldo < monto {
		return fmt.Errorf("saldo insuficiente: tenés $%.2f, querés retirar $%.2f", *saldo, monto)
	}
	*saldo -= monto
	return nil
}
