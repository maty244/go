package main

import "fmt"

// =========================================================
// PUNTEROS: OPERADORES * Y &
// =========================================================
// Un valor de tipo puntero identifica una variable a la que podemos acceder.
// Una variable p de tipo *int guarda ese puntero, no el int de su destino.
// Hay que distinguir DOS variables: la que guarda el puntero y la apuntada.
//
// Antes de ejecutar: abrí PUNTEROS_DESDE_CERO.md en la carpeta de esta etapa.
// Tiene diagramas y tablas del estado de cada variable. No memorices 0xc...:
// cambian entre ejecuciones y no indican qué valor hay guardado en el destino.
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

	// Cada ejemplo usa sus propias variables: no necesitás recordar el estado
	// de un ejemplo para entender el siguiente.
	variablesYTipos()

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

	// Diagrama: [punteroEdad: dirección de edad] ──→ [edad: 25]
	// La caja edad CONTIENE 25, no una dirección hacia otra caja con el 25.
	// &punteroEdad es la dirección de la propia caja que guarda el puntero.

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

	copiarValorYCompartirDestino()
	compararPunterosYValores()

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

	pDirecto := new(42) // Go 1.26+: crea OTRO int, inicialmente 42 (no 77)
	fmt.Printf("pDirecto := new(42) → %p → valor: %d\n", pDirecto, *pDirecto)
	// new(x) usa el valor actual de x como inicializador, no apunta a x.
	// &x sí apunta a x. La diferencia se ve al cambiar x después.
	porDireccion := &x
	porCopia := new(x)
	x = 100
	fmt.Printf("x=%d | *&x=%d | *new(x anterior)=%d\n", x, *porDireccion, *porCopia) // 100, 100, 42

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

	contador := 10
	fmt.Println("contador antes:", contador)
	incrementar(&contador)                     // el argumento es *int; el parámetro recibe su COPIA
	fmt.Println("contador después:", contador) // 11
	// La dirección copiada permite llegar al mismo contador. No existe otra
	// modalidad de paso de argumentos: incluso el puntero se pasa por valor.
	fmt.Println("Para distinguir cambiar *p de cambiar p, seguí 15_punteros_en_funciones.")

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
	// ppv apunta a la VARIABLE pv. *ppv es pv; **ppv es valor.
	// Si asignamos *ppv = &otro, cambiamos pv, no el int que apuntaba antes.
	// La lección 15 compara esto con retornar el nuevo puntero, que suele
	// ser una interfaz más simple para quien llama a la función.

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

// Premisa: p no es nil. El llamador de este ejemplo pasa &contador.
func incrementar(p *int) {
	*p = *p + 1 // leer el int del destino, sumar uno y escribir en ese destino
}

func variablesYTipos() {
	fmt.Println("\n=== 0. Leer cada expresión antes de combinar operadores ===")
	n := 10 // n es una variable int; su contenido es el número 10
	p := &n // p es OTRA variable, *int; su contenido identifica a n
	fmt.Printf("n:  tipo=%T valor=%d\n", n, n)
	fmt.Printf("&n: tipo=%T dirección=%p\n", &n, &n)
	fmt.Printf("p:  tipo=%T destino=%p\n", p, p)
	fmt.Printf("*p: tipo=%T valor=%d\n", *p, *p)
	fmt.Printf("&p: tipo=%T dirección de la variable p=%p\n", &p, &p)
	fmt.Println("p == &n:", p == &n) // true: identifican a la misma variable
	// * en "var p *int" forma un TIPO. * en "*p" es una EXPRESIÓN.
	// En "2 * 3" significa multiplicación; el contexto determina el sentido.
}

func copiarValorYCompartirDestino() {
	fmt.Println("\n=== Copiar un int, copiar un puntero y reapuntar ===")
	n := 10
	p := &n
	copia := *p // leer *p da 10; asignarlo copia el int a otra variable
	q := p      // copiar p copia el puntero: ambos pueden acceder a n
	*p = 20
	fmt.Printf("tras *p=20: n=%d copia=%d *q=%d\n", n, copia, *q)    // 20, 10, 20
	copia = 30                                                       // no cambia n, p ni q
	fmt.Printf("tras copia=30: n=%d copia=%d *p=%d\n", n, copia, *p) // 20, 30, 20

	otro := 50
	p = &otro                                                                 // cambiar EL PUNTERO p no cambia q ni modifica n
	*p = 60                                                                   // ahora el destino de p es otro; q sigue apuntando a n
	fmt.Printf("tras reapuntar: n=%d otro=%d *p=%d *q=%d\n", n, otro, *p, *q) // 20, 60, 60, 20
	fmt.Println("¿p y q siguen apuntando a la misma variable?", p == q)       // false
}

func compararPunterosYValores() {
	fmt.Println("\n=== Igual contenido no implica igual variable ===")
	a, b := 7, 7 // dos variables int distintas con el mismo contenido
	pa, pb := &a, &b
	fmt.Println("a == b:", a == b)         // true: comparamos los ints
	fmt.Println("pa == pb:", pa == pb)     // false: comparamos sus destinos
	fmt.Println("*pa == *pb:", *pa == *pb) // true: leemos y comparamos los ints
	// No se puede escribir pa == 7: *int e int son tipos distintos.
	// Esta comparación usa ints: los objetos de tamaño cero tienen reglas
	// especiales de identidad. No generalices este ejemplo a &struct{}{}.
}
