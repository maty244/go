// Un programa ejecutable de Go comienza en el paquete main y necesita una
// función main. Los comentarios de esta lección explican incluso las piezas
// que normalmente se dan por conocidas.
package main

// import declara un paquete cuyo código queremos usar. fmt pertenece a la
// biblioteca estándar y sirve, entre otras cosas, para imprimir texto.
import "fmt"

// main es el punto de entrada. No recibe parámetros ni retorna valores.
func main() {
	fmt.Println("Hola, Go 1.26")

	fmt.Println("\nProbá estos comandos desde la raíz:")
	fmt.Println("  go run ./00_inicio/01_primer_programa")
	fmt.Println("  go build ./00_inicio/01_primer_programa")
	fmt.Println("  go doc fmt.Println")
	fmt.Println("  gofmt -w ./00_inicio/01_primer_programa/main.go")
}
