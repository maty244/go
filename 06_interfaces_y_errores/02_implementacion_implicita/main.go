package main

import "fmt"

// =========================================================
// IMPLEMENTACIÓN IMPLÍCITA
// =========================================================
// En otros lenguajes (Java, C#, TypeScript) para implementar una
// interfaz tenés que ESCRIBIRLO explícitamente:
//   class Tarjeta implements MetodoPago { ... }
//
// En Go NO existe la palabra "implements". Un tipo implementa una
// interfaz con solo TENER los métodos necesarios. Go lo detecta
// solo, en tiempo de compilación. A esto se le llama "duck typing
// estático": si camina como pato y grazna como pato... y además
// el compilador te avisa en el momento si te faltó un método.
//
// Ventaja enorme: podés definir una interfaz que describa un tipo
// que ni siquiera escribiste vos (por ejemplo, un tipo de una
// librería externa). Si ya tiene los métodos, ya la cumple.

// ─────────────────────────────────────────────────────────
// LA INTERFAZ SE PUEDE DEFINIR DESPUÉS DEL TIPO
// ─────────────────────────────────────────────────────────
// Esto es clave: la interfaz NO tiene que existir antes.
// Google podría publicar un tipo "Cuadrado", y VOS podés definir
// tu propia interfaz "Figura" años después. Si Cuadrado ya tenía
// el método Area(), automáticamente la cumple.

type Cuadrado struct {
	Lado float64
}

func (c Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return 3.14159 * c.Radio * c.Radio
}

// Definimos la interfaz DESPUÉS, y ambos tipos ya la cumplen
// sin que tengamos que tocarlos ni una línea.
type Figura interface {
	Area() float64
}

func mostrarArea(f Figura) {
	fmt.Printf("  Área: %.2f\n", f.Area())
}

// ─────────────────────────────────────────────────────────
// CÓMO EL COMPILADOR TE PROTEGE
// ─────────────────────────────────────────────────────────
// Si un tipo NO tiene todos los métodos, Go directamente no
// compila cuando intentás usarlo donde se espera la interfaz.
// No hay excepción en tiempo de ejecución tipo "método no
// encontrado": el error aparece ANTES de correr el programa.

type Triangulo struct {
	Base, Altura float64
}

// Si comentamos este método, Triangulo deja de cumplir Figura,
// y "var f Figura = Triangulo{...}" no compilaría más.
func (t Triangulo) Area() float64 {
	return t.Base * t.Altura / 2
}

// ─────────────────────────────────────────────────────────
// VERIFICACIÓN EXPLÍCITA (patrón común en librerías)
// ─────────────────────────────────────────────────────────
// Aunque Go no exige declarar "implements", es común agregar esta
// línea para que el compilador verifique la implementación en el
// momento, y falle temprano si alguien rompe el contrato sin querer.
// _ (blank identifier) descarta el valor: solo nos interesa que
// COMPILE.
var _ Figura = Cuadrado{}
var _ Figura = Circulo{}
var _ Figura = Triangulo{}

func main() {
	fmt.Println("=== Implementación implícita ===")

	figuras := []Figura{
		Cuadrado{Lado: 4},
		Circulo{Radio: 3},
		Triangulo{Base: 6, Altura: 5},
	}

	for _, f := range figuras {
		mostrarArea(f)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: repositorios intercambiables
	// ─────────────────────────────────────────────────────────
	// En un kiosco digital, guardar productos "en memoria" (para
	// testear) o "en una base de datos" (en producción) son cosas
	// MUY distintas por dentro, pero para el resto del programa
	// deberían verse igual. La interfaz lo permite.

	fmt.Println("\n=== Caso real: repositorio de productos ===")

	var repo RepositorioProductos = &repoMemoria{
		productos: map[string]float64{"Coca-Cola": 800, "Alfajor": 500},
	}

	precio, ok := repo.Buscar("Alfajor")
	fmt.Printf("Alfajor cuesta $%.2f (encontrado: %v)\n", precio, ok)

	_, ok = repo.Buscar("Producto inexistente")
	fmt.Println("Producto inexistente encontrado:", ok)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Go NO tiene 'implements': la implementación es implícita")
	fmt.Println("  Un tipo cumple una interfaz con solo TENER sus métodos")
	fmt.Println("  El compilador verifica esto ANTES de correr el programa")
	fmt.Println("  var _ Interfaz = Tipo{} → chequeo explícito, común en libs")
}

// RepositorioProductos: el contrato que cualquier "fuente de datos"
// de productos debe cumplir.
type RepositorioProductos interface {
	Buscar(nombre string) (precio float64, encontrado bool)
}

// repoMemoria implementa RepositorioProductos guardando todo en un map.
// Mañana podría existir un repoPostgres con el mismo contrato.
type repoMemoria struct {
	productos map[string]float64
}

func (r *repoMemoria) Buscar(nombre string) (float64, bool) {
	precio, ok := r.productos[nombre]
	return precio, ok
}
