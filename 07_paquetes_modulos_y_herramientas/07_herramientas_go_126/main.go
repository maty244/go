package main

import "fmt"

// Esta lección imprime comandos en vez de ejecutarlos. Los comandos que
// modifican código, como gofmt o go fix, deben correrse conscientemente desde
// la terminal y revisando su diff.
func main() {
	fmt.Println("=== Ciclo diario ===")
	fmt.Println("go fmt ./...       → formatea paquetes")
	fmt.Println("go vet ./...       → análisis estático")
	fmt.Println("go test ./...      → tests del módulo")
	fmt.Println("go test -race ./...→ tests instrumentados para detectar carreras")
	fmt.Println("go build ./...     → compila todos los paquetes")
	fmt.Println("go doc io.Reader   → documentación local")

	fmt.Println("\n=== Go 1.26: modernizadores ===")
	fmt.Println("go fix -diff ./... → muestra sugerencias sin editar")
	fmt.Println("go fix ./...       → aplica sugerencias; revisá el diff después")

	fmt.Println("\n=== Módulos ===")
	fmt.Println("go mod tidy        → sincroniza imports con go.mod/go.sum")
	fmt.Println("go list ./...      → enumera paquetes del módulo")
	fmt.Println("go env             → inspecciona el entorno de Go")
}
