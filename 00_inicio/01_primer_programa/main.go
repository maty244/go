// =========================================================
// INTRODUCCIÓN A GO: HISTORIA, FILOSOFÍA Y PRIMER PROGRAMA
// =========================================================
// ¿Por qué Go y de dónde sale?
//   - Creado en Google en 2007 por Robert Griesemer, Rob Pike y Ken Thompson
//     (co-creadores de Unix, UTF-8 y Plan 9), y abierto al público en 2009.
//   - Desde Go 1.0 (2012) mantiene una estricta promesa de compatibilidad:
//     el código escrito hace años sigue compilando en versiones modernas.
//   - ¿Por qué se usa tanto?
//     1. Compilación ultrarrápida a un ÚNICO binario nativo (sin instalar VM).
//     2. Tipado estático fuerte con sintaxis limpia y minimalista (25 palabras clave).
//     3. Gestión automática de memoria (Garbage Collector de muy baja latencia).
//     4. Concurrencia nativa de primera clase (goroutines y channels).
//     5. Toolchain completo incluido de fábrica (formateador, tests, benchmarks,
//     detector de carreras, perfilador y gestor de paquetes).
//
// Anatomía mínima:
//
//	Un programa ejecutable de Go comienza siempre en el paquete `main` y
//	necesita una función `main()`.
package main

// import declara un paquete cuyo código queremos usar. "fmt" (format) pertenece
// a la biblioteca estándar y sirve para formatear e imprimir texto.
import "fmt"

// main es el punto de entrada del ejecutable. No recibe parámetros ni retorna
// valores (si necesitás devolver un código de salida al sistema operativo, usás os.Exit).
func main() {
	fmt.Println("Hola, Go 1.26")

	fmt.Println("\nProbá estos comandos desde la raíz:")
	fmt.Println("  go run ./00_inicio/01_primer_programa")
	fmt.Println("  go build ./00_inicio/01_primer_programa")
	fmt.Println("  go doc fmt.Println")
	fmt.Println("  gofmt -w ./00_inicio/01_primer_programa/main.go")

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  package main → indica que este paquete compila a un programa ejecutable")
	fmt.Println("  import       → trae paquetes de la biblioteca estándar o de terceros")
	fmt.Println("  func main()  → punto donde arranca la ejecución del programa")
	fmt.Println("  go run       → compila en una carpeta temporal y ejecuta al instante")
	fmt.Println("  go build     → compila y genera el archivo binario ejecutable")
}
