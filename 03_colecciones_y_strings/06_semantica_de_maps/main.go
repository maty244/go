package main

import (
	"fmt"
	"maps"
	"slices"
)

// =========================================================
// MAPS: DESCRIPTORES QUE COMPARTEN UNA TABLA
// =========================================================
// Un valor map es un descriptor administrado por el runtime. Asignarlo o
// pasarlo a una función copia el descriptor, no todas sus entradas.
//
//   var m map[string]int   → m == nil (no inicializado)
//   m = make(map[string]int) → m apunta a la estructura interna
//
//   m2 := m   → se copia el descriptor; las entradas se comparten
//
// Consecuencias:
//   - Leer de nil map: OK (retorna zero value)
//   - Escribir en nil map: PANIC
//   - Pasar un map a una función copia el descriptor
//   - No necesitás *map[K]V para modificar sus entradas

// ─────────────────────────────────────────────────────────
// INVENTARIO: ejemplo de uso real
// ─────────────────────────────────────────────────────────

type Inventario struct {
	stock map[string]int
}

func NuevoInventario() *Inventario {
	return &Inventario{
		stock: make(map[string]int),
	}
}

func (inv *Inventario) Agregar(producto string, cantidad int) {
	inv.stock[producto] += cantidad
}

func (inv *Inventario) Retirar(producto string, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("cantidad debe ser positiva: %d", cantidad)
	}
	actual, existe := inv.stock[producto]
	if !existe {
		return fmt.Errorf("producto '%s' no existe", producto)
	}
	if actual < cantidad {
		return fmt.Errorf("stock insuficiente: %d < %d", actual, cantidad)
	}
	inv.stock[producto] -= cantidad
	if inv.stock[producto] == 0 {
		delete(inv.stock, producto)
	}
	return nil
}

func (inv *Inventario) Stock(producto string) (int, bool) {
	cantidad, existe := inv.stock[producto]
	return cantidad, existe
}

func (inv *Inventario) Imprimir() {
	// El orden del map no está especificado → ordenar para output consistente.
	keys := slices.Sorted(maps.Keys(inv.stock))

	for _, k := range keys {
		fmt.Printf("    %-15s %d unidades\n", k, inv.stock[k])
	}
}

// ─────────────────────────────────────────────────────────
// CACHÉ: mapa como caché simple
// ─────────────────────────────────────────────────────────

type Cache struct {
	datos map[string]string
}

func NuevaCache() Cache {
	return Cache{datos: make(map[string]string)}
}

// La función recibe Cache por valor. La copia de datos contiene un descriptor
// que accede a la misma tabla, así que las nuevas entradas se ven afuera.
// (Esto puede ser sorprendente si venís de otros lenguajes)
func setEnCache(c Cache, clave, valor string) {
	c.datos[clave] = valor // modifica el map compartido
}

// ─────────────────────────────────────────────────────────
// CONJUNTO (set pattern con map)
// ─────────────────────────────────────────────────────────

// Primero lo hacemos concreto para strings. La generalización con [T
// comparable] se estudia en 08_generics_e_iteradores, no es necesaria acá.
type Conjunto struct {
	elementos map[string]struct{} // struct{} es un valor sin campos
}

func NuevoConjunto() *Conjunto {
	return &Conjunto{elementos: make(map[string]struct{})}
}

func (c *Conjunto) Agregar(v string) {
	c.elementos[v] = struct{}{}
}

func (c *Conjunto) Contiene(v string) bool {
	_, existe := c.elementos[v]
	return existe
}

func (c *Conjunto) Largo() int {
	return len(c.elementos)
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║       MAPS: DATOS COMPARTIDOS     ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// nil MAP vs make MAP
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== nil map vs map inicializado ===")

	var nilMap map[string]int
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)
	fmt.Printf("Leer de nil map: nilMap[\"x\"] = %d  (OK, retorna zero value)\n", nilMap["x"])

	// Escribir en nil map → panic (comentado para que no crashee)
	// nilMap["x"] = 1  → panic: assignment to entry in nil map

	inicializado := make(map[string]int)
	inicializado["contador"] = 1
	fmt.Printf("Mapa inicializado: %v\n", inicializado)

	// ─────────────────────────────────────────────────────────
	// COPIAR EL DESCRIPTOR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Copiar el descriptor ===")

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := m1 // copia el descriptor, no las entradas

	m2["c"] = 3     // modifica el map compartido
	delete(m2, "a") // borra de ambos

	fmt.Printf("m1 = %v  (afectado por cambios en m2)\n", m1)
	fmt.Printf("m2 = %v\n", m2)
	fmt.Println("  → m1 y m2 son valores distintos que comparten la tabla")

	// ─────────────────────────────────────────────────────────
	// PASAR MAP A FUNCIÓN: el descriptor basta, no necesitás *map
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Map en funciones: no necesitás *map ===")

	contadores := map[string]int{}
	contarPalabras(contadores, []string{"go", "es", "go", "rápido", "go"})

	keys := []string{"go", "es", "rápido"}
	for _, k := range keys {
		fmt.Printf("  '%s': %d veces\n", k, contadores[k])
	}
	fmt.Println("  → La función modificó el map sin necesitar *map[string]int")

	// ─────────────────────────────────────────────────────────
	// CACHE: struct por valor pero map interno es referencia
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct por valor con map interno ===")

	cache := NuevaCache()
	// Pasamos cache por valor, pero el descriptor interno comparte la tabla
	setEnCache(cache, "usuario:1", "Ana García")
	setEnCache(cache, "usuario:2", "Carlos López")

	// Las claves se ven afuera porque ambos descriptores acceden a la tabla
	fmt.Printf("cache después de setEnCache: %v\n", cache.datos)
	fmt.Println("  → El struct se copió por valor pero el map interno se comparte")
	fmt.Println("  ⚠️  Esto puede ser confuso: cuidado al pasar structs con maps por valor")

	// ─────────────────────────────────────────────────────────
	// INVENTARIO: uso real
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Inventario con map ===")

	inv := NuevoInventario()
	inv.Agregar("Notebook", 10)
	inv.Agregar("Mouse", 50)
	inv.Agregar("Teclado", 30)
	inv.Agregar("Mouse", 20) // suma al stock existente

	fmt.Println("Inventario inicial:")
	inv.Imprimir()

	if err := inv.Retirar("Mouse", 15); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Retiro de 15 Mouse: OK")
	}

	if cant, existe := inv.Stock("Mouse"); existe {
		fmt.Printf("Stock de Mouse: %d\n", cant)
	}

	if err := inv.Retirar("Monitor", 1); err != nil {
		fmt.Printf("Error esperado: %v\n", err)
	}

	// ─────────────────────────────────────────────────────────
	// CONJUNTO (set) con map
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Conjunto (set) con map[string]struct{} ===")

	visitados := NuevoConjunto()
	urls := []string{"/home", "/about", "/home", "/contact", "/about", "/home"}

	for _, url := range urls {
		if !visitados.Contiene(url) {
			visitados.Agregar(url)
			fmt.Printf("  Primera visita a: %s\n", url)
		} else {
			fmt.Printf("  Ya visitado: %s\n", url)
		}
	}
	fmt.Printf("URLs únicas visitadas: %d\n", visitados.Largo())

	// ─────────────────────────────────────────────────────────
	// CLONAR UN MAP (copia real)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Clonar un map (copia independiente) ===")

	original := map[string]int{"a": 1, "b": 2, "c": 3}
	clon := clonarMap(original)

	clon["a"] = 999
	delete(clon, "b")

	fmt.Printf("original = %v  (intacto)\n", original)
	fmt.Printf("clon     = %v  (modificado independientemente)\n", clon)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen: descriptor de map ===")
	fmt.Println("  make(map[K]V)        → inicializa el map (necesario para escribir)")
	fmt.Println("  var m map[K]V        → nil, se puede leer pero NO escribir")
	fmt.Println("  m2 := m1             → copia descriptor; comparte entradas")
	fmt.Println("  func f(m map[K]V)    → puede modificar la tabla compartida")
	fmt.Println("  func f(m *map[K]V)   → solo si hay que reemplazar el descriptor")
	fmt.Println()
	fmt.Println("  Para copiar: maps.Clone crea un map independiente (copia superficial)")
	fmt.Println("  Para set:    map[string]struct{} (solo importa si la clave existe)")
}

func contarPalabras(m map[string]int, palabras []string) {
	for _, p := range palabras {
		m[p]++ // modifica la tabla compartida por los descriptores
	}
}

func clonarMap(m map[string]int) map[string]int {
	return maps.Clone(m)
}
