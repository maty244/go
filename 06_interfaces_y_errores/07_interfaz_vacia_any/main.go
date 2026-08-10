package main

import "fmt"

// =========================================================
// LA INTERFAZ VACÍA: any (antes interface{})
// =========================================================
// interface{} es una interfaz SIN métodos. Como no pide nada,
// CUALQUIER valor de CUALQUIER tipo la cumple automáticamente.
//
// Desde Go 1.18 existe el alias "any", que es EXACTAMENTE lo
// mismo que interface{} pero más corto de leer. Hoy se usa "any"
// en código nuevo; vas a ver interface{} en código más viejo.
//
//   var x any = 42        // equivalente a: var x interface{} = 42
//
// ¿Para qué sirve algo que acepta "cualquier cosa"? Para casos
// donde el tipo REALMENTE no se puede saber de antemano: un
// valor que viene de un archivo JSON, una fila de base de datos
// genérica, o una función que loggea "lo que sea".

func main() {
	// ─────────────────────────────────────────────────────────
	// any ACEPTA CUALQUIER TIPO
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== any acepta cualquier valor ===")

	var cualquiera any

	cualquiera = 42
	fmt.Println("int:", cualquiera)

	cualquiera = "Matias"
	fmt.Println("string:", cualquiera)

	cualquiera = []int{1, 2, 3}
	fmt.Println("slice:", cualquiera)

	cualquiera = Producto{Nombre: "Mouse", Precio: 999.99}
	fmt.Println("struct:", cualquiera)

	// ─────────────────────────────────────────────────────────
	// []any: una lista que mezcla tipos distintos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== []any: mezclar tipos en un slice ===")

	mezcla := []any{1, "dos", 3.0, true, []int{4, 5}}
	for _, v := range mezcla {
		fmt.Printf("  valor=%v  tipo=%T\n", v, v)
	}

	// ─────────────────────────────────────────────────────────
	// PARA USAR EL VALOR REAL, NECESITÁS TYPE ASSERTION
	// ─────────────────────────────────────────────────────────
	// any no tiene métodos (es un contrato vacío), así que no
	// podés hacer nada específico con el valor hasta que
	// recuperás su tipo concreto (visto en los temas 05 y 06).

	fmt.Println("\n=== Recuperar el tipo real ===")

	var x any = 100
	// x + 1 // ERROR: any no tiene operador +, no sabe que es un número

	if n, ok := x.(int); ok {
		fmt.Println("Ahora sí puedo sumar:", n+1)
	}

	// ─────────────────────────────────────────────────────────
	// fmt.Println YA RECIBE any
	// ─────────────────────────────────────────────────────────
	// Este es el motivo por el que fmt.Println(loQueSea) siempre
	// compiló, sin importar qué le pasaste: su firma real es
	// func Println(a ...any) (n int, err error)

	fmt.Println("\n=== Por qué fmt.Println acepta todo ===")
	fmt.Println("func Println(a ...any) → por eso acepta cualquier cosa")

	// ─────────────────────────────────────────────────────────
	// CUIDADO: any NO ES "tipado dinámico mágico"
	// ─────────────────────────────────────────────────────────
	// Go sigue siendo fuertemente tipado. any solo pospone la
	// decisión del tipo hasta el momento de usar el valor.
	// Abusar de any hace que pierdas las ventajas del compilador
	// (errores que antes se detectaban en compilación, ahora
	// solo aparecen en tiempo de ejecución). Usalo con moderación;
	// en el próximo tema (Generics) vas a ver una alternativa
	// más segura para muchos casos.

	// ─────────────────────────────────────────────────────────
	// CASO REAL: un log genérico
	// ─────────────────────────────────────────────────────────
	// Una función de logging necesita aceptar CUALQUIER dato
	// para registrarlo, sin importar su tipo.

	fmt.Println("\n=== Caso real: función de log genérica ===")
	registrarEvento("usuario_login", "matias")
	registrarEvento("carrito_total", 15499.90)
	registrarEvento("stock_bajo", []string{"Mouse", "Teclado"})

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  any            → interface{} sin métodos, acepta CUALQUIER valor`)
	fmt.Println(`  []any          → slice con tipos mezclados`)
	fmt.Println(`  Para USAR el   → hace falta type assertion o type switch`)
	fmt.Println(`  valor real`)
	fmt.Println(`  Usalo poco     → perdés chequeos del compilador; ver Generics`)
}

type Producto struct {
	Nombre string
	Precio float64
}

func registrarEvento(clave string, valor any) {
	fmt.Printf("  [LOG] %s = %v (%T)\n", clave, valor, valor)
}
