package main

import "fmt"

// =========================================================
// PUNTEROS EN STRUCTS
// =========================================================
// Un struct puede contener campos de tipo puntero (*T).
// Esto permite:
//
//  1. Campos OPCIONALES (nil = ausente) — ya visto en 05
//  2. Campos que COMPARTEN el mismo dato entre structs
//  3. Estructuras RECURSIVAS (linked list, árbol, grafo)
//     → Un struct NO puede contenerse a sí mismo por valor
//       (tamaño infinito), pero SÍ puede tener un *propio_tipo.
//
// También veremos:
//  4. Slices de punteros vs slice de valores
//  5. El acceso automático a campos: s.Campo == (*s).Campo

// ─────────────────────────────────────────────────────────
// 1. CAMPOS COMPARTIDOS (dos structs apuntan al mismo dato)
// ─────────────────────────────────────────────────────────

type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

type Empleado struct {
	Nombre    string
	Direccion *Direccion // puntero → puede compartirse
}

// ─────────────────────────────────────────────────────────
// 2. LINKED LIST (lista enlazada)
// ─────────────────────────────────────────────────────────
// La estructura más simple que requiere punteros internos.
// Cada nodo tiene un valor y un puntero al siguiente nodo.
// El último nodo apunta a nil (fin de la lista).
//
//  [1|→] → [2|→] → [3|→] → [nil]

type Nodo struct {
	Valor     int
	Siguiente *Nodo // puntero a sí mismo: SOLO posible con *Nodo
}

type ListaEnlazada struct {
	Cabeza *Nodo
	Largo  int
}

func (l *ListaEnlazada) Agregar(v int) {
	nuevo := &Nodo{Valor: v}
	if l.Cabeza == nil {
		l.Cabeza = nuevo
	} else {
		// Recorrer hasta el último nodo
		actual := l.Cabeza
		for actual.Siguiente != nil {
			actual = actual.Siguiente
		}
		actual.Siguiente = nuevo
	}
	l.Largo++
}

func (l *ListaEnlazada) Imprimir() {
	actual := l.Cabeza
	fmt.Print("  [")
	for actual != nil {
		fmt.Print(actual.Valor)
		if actual.Siguiente != nil {
			fmt.Print(" → ")
		}
		actual = actual.Siguiente
	}
	fmt.Println("]")
}

func (l *ListaEnlazada) Buscar(v int) *Nodo {
	actual := l.Cabeza
	for actual != nil {
		if actual.Valor == v {
			return actual
		}
		actual = actual.Siguiente
	}
	return nil
}

// ─────────────────────────────────────────────────────────
// 3. ÁRBOL BINARIO
// ─────────────────────────────────────────────────────────
// Cada nodo tiene dos hijos opcionales: Izquierdo y Derecho.
// nil significa "no hay hijo".
//
//          10
//         /  \
//        5    15
//       / \     \
//      3   7    20

type NodoArbol struct {
	Valor     int
	Izquierdo *NodoArbol // nil = sin hijo izquierdo
	Derecho   *NodoArbol // nil = sin hijo derecho
}

func insertar(raiz *NodoArbol, v int) *NodoArbol {
	if raiz == nil {
		return &NodoArbol{Valor: v}
	}
	if v < raiz.Valor {
		raiz.Izquierdo = insertar(raiz.Izquierdo, v)
	} else if v > raiz.Valor {
		raiz.Derecho = insertar(raiz.Derecho, v)
	}
	return raiz
}

// recorrido in-order (izquierdo → raíz → derecho) da orden ascendente
func inOrder(n *NodoArbol, resultado *[]int) {
	if n == nil {
		return
	}
	inOrder(n.Izquierdo, resultado)
	*resultado = append(*resultado, n.Valor)
	inOrder(n.Derecho, resultado)
}

// ─────────────────────────────────────────────────────────
// 4. SLICE DE PUNTEROS vs SLICE DE VALORES
// ─────────────────────────────────────────────────────────

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

// Con slice de VALORES: modificar un elemento requiere índice
func aumentarPrecioValor(productos []Producto, pct float64) {
	for i := range productos {
		productos[i].Precio *= (1 + pct) // i es necesario para modificar
	}
}

// Con slice de PUNTEROS: podemos iterar con range directamente
func aumentarPrecioPuntero(productos []*Producto, pct float64) {
	for _, p := range productos { // p es *Producto
		p.Precio *= (1 + pct) // modificamos el original directamente
	}
}

// ─────────────────────────────────────────────────────────
// 5. ACCESO AUTOMÁTICO: s.Campo == (*s).Campo
// ─────────────────────────────────────────────────────────

type Punto struct {
	X, Y float64
}

func (p *Punto) Escalar(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Punto) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       PUNTEROS EN STRUCTS         ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// CAMPOS COMPARTIDOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Campos compartidos con punteros ===")

	dir := &Direccion{
		Calle:  "Av. Corrientes 1234",
		Ciudad: "Buenos Aires",
		CP:     "1043",
	}

	empleado1 := Empleado{Nombre: "Ana", Direccion: dir}
	empleado2 := Empleado{Nombre: "Carlos", Direccion: dir} // comparten la misma dirección

	fmt.Printf("Ana vive en: %s, %s\n", empleado1.Direccion.Calle, empleado1.Direccion.Ciudad)
	fmt.Printf("Carlos vive en: %s, %s\n", empleado2.Direccion.Calle, empleado2.Direccion.Ciudad)

	// Actualizar la dirección afecta a AMBOS
	dir.Ciudad = "Córdoba"
	fmt.Println("\nDespués de cambiar la ciudad a Córdoba:")
	fmt.Printf("Ana vive en: %s\n", empleado1.Direccion.Ciudad)
	fmt.Printf("Carlos vive en: %s\n", empleado2.Direccion.Ciudad)
	fmt.Println("  → Ambos ven el cambio porque comparten el mismo *Direccion")

	// ─────────────────────────────────────────────────────────
	// LINKED LIST
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Lista enlazada (Linked List) ===")

	lista := &ListaEnlazada{}
	for _, v := range []int{10, 20, 30, 40, 50} {
		lista.Agregar(v)
	}

	fmt.Printf("Lista (%d elementos):\n", lista.Largo)
	lista.Imprimir()

	if nodo := lista.Buscar(30); nodo != nil {
		fmt.Printf("Encontrado: %d (siguiente: %d)\n", nodo.Valor, nodo.Siguiente.Valor)
	}

	// ─────────────────────────────────────────────────────────
	// ÁRBOL BINARIO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Árbol binario de búsqueda ===")

	var raiz *NodoArbol
	for _, v := range []int{10, 5, 15, 3, 7, 20} {
		raiz = insertar(raiz, v)
	}

	var ordenado []int
	inOrder(raiz, &ordenado)
	fmt.Printf("Valores insertados: [10, 5, 15, 3, 7, 20]\n")
	fmt.Printf("In-order (ascendente): %v\n", ordenado)
	fmt.Println("  → El árbol organiza los datos automáticamente")

	// ─────────────────────────────────────────────────────────
	// SLICE DE VALORES vs SLICE DE PUNTEROS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Slice de valores vs slice de punteros ===")

	// Slice de valores
	productosV := []Producto{
		{"Notebook", 1000.00, 5},
		{"Mouse", 20.00, 50},
		{"Teclado", 60.00, 30},
	}

	// Slice de punteros (todos permiten acceder a los mismos structs).
	// La ubicación stack/heap concreta la decide el compilador.
	productosP := []*Producto{
		{"Notebook", 1000.00, 5},
		{"Mouse", 20.00, 50},
		{"Teclado", 60.00, 30},
	}

	fmt.Println("Antes del aumento de precio:")
	for _, p := range productosV {
		fmt.Printf("  %-10s $%.2f\n", p.Nombre, p.Precio)
	}

	aumentarPrecioValor(productosV, 0.10)
	aumentarPrecioPuntero(productosP, 0.10)

	fmt.Println("Después de aumentar 10%:")
	fmt.Println("  [valores]:")
	for _, p := range productosV {
		fmt.Printf("    %-10s $%.2f\n", p.Nombre, p.Precio)
	}
	fmt.Println("  [punteros]:")
	for _, p := range productosP {
		fmt.Printf("    %-10s $%.2f\n", p.Nombre, p.Precio)
	}

	// ─────────────────────────────────────────────────────────
	// ACCESO AUTOMÁTICO A CAMPOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Acceso automático: p.Campo == (*p).Campo ===")

	pt := &Punto{3.0, 4.0}

	fmt.Printf("Punto original:  %s\n", pt)
	fmt.Printf("pt.X            = %.1f  (Go lo transforma en (*pt).X)\n", pt.X)
	fmt.Printf("(*pt).Y         = %.1f  (equivalente explícito)\n", (*pt).Y)

	pt.Escalar(2.0) // método con receptor *Punto
	fmt.Printf("Después Escalar(2): %s\n", pt)

	// ─────────────────────────────────────────────────────────
	// CUÁNDO usar cada struct con campos puntero
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen: cuándo usar *T en campos de struct ===")
	fmt.Println("  Campo *T en struct:")
	fmt.Println("    → Estructura recursiva (linked list, árbol, grafo)")
	fmt.Println("    → Campo opcional (nil = ausente)")
	fmt.Println("    → Compartir el mismo dato entre múltiples structs")
	fmt.Println()
	fmt.Println("  Slice de *T:")
	fmt.Println("    → Modificar elementos con range (sin índice)")
	fmt.Println("    → Struct grande donde copiar sería caro")
	fmt.Println("    → Elementos compartidos entre múltiples slices")
	fmt.Println()
	fmt.Println("  Slice de T:")
	fmt.Println("    → Struct pequeño, acceso por índice para modificar")
	fmt.Println("    → Mejor localidad de caché (datos contiguos en memoria)")
}
