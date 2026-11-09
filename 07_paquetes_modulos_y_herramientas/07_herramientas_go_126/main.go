package main

import "fmt"

// =========================================================
// TOOLCHAIN DE GO: COMANDOS, CALIDAD, SEGURIDAD Y DESPLIEGUE
// =========================================================
// Uno de los mayores diferenciales de Go es que su toolchain oficial
// incluye casi todo lo que en otros lenguajes requiere instalar diez
// herramientas distintas. Además, el ecosistema suma linters y escáneres
// de vulnerabilidades que se integran directo con `go tool`.
//
// Esta lección imprime los comandos de referencia en vez de ejecutarlos.
// Los comandos que modifican código (como gofmt o go fix) deben correrse
// conscientemente desde la terminal y revisando su diff.
func main() {
	fmt.Println("=== 1. Comandos principales (Core Go Commands) ===")
	fmt.Println("go run ./paquete   → compila y ejecuta en el acto")
	fmt.Println("go build ./...     → compila paquetes o genera el binario ejecutable")
	fmt.Println("go install pkg@ver → compila e instala un binario en $GOBIN (~/go/bin)")
	fmt.Println("go fmt ./...       → formatea todo el código con el estilo oficial")
	fmt.Println("go test ./...      → corre tests unitarios del módulo")
	fmt.Println("go clean -cache    → limpia archivos objeto y caché de compilación")
	fmt.Println("go doc io.Reader   → consulta documentación desde la terminal")
	fmt.Println("go version -m bin  → muestra con qué versión y dependencias se compiló un binario")
	fmt.Println("go env             → muestra las variables del entorno de Go")

	fmt.Println("\n=== 2. Módulos y dependencias (go mod) ===")
	fmt.Println("go mod init nombre → crea un nuevo módulo (archivo go.mod)")
	fmt.Println("go mod tidy        → agrega dependencias faltantes y borra las no usadas")
	fmt.Println("go mod vendor      → copia las dependencias a una carpeta local vendor/")
	fmt.Println("git tag v1.0.0     → publicar un módulo es simplemente etiquetar con SemVer en Git")

	fmt.Println("\n=== 3. Calidad de código, Linters y Seguridad ===")
	fmt.Println("go vet ./...       → analizador estático oficial para detectar errores sospechosos")
	fmt.Println("goimports -w .     → formatea código y además ordena/agrega/quita imports")
	fmt.Println("staticcheck ./...  → linter avanzado de análisis estático")
	fmt.Println("revive ./...       → linter rápido y configurable de estilo")
	fmt.Println("golangci-lint run  → ejecutor que agrupa decenas de linters en paralelo")
	fmt.Println("govulncheck ./...  → escáner oficial de seguridad: avisa solo si LLAMÁS código vulnerable")

	fmt.Println("\n=== 4. Go 1.26: modernizadores y directiva tool ===")
	fmt.Println("go fix -diff ./... → muestra sugerencias de código moderno sin editar archivos")
	fmt.Println("go fix ./...       → aplica modernizadores oficiales; revisá el diff después")
	fmt.Println("go get -tool pkg   → registra una herramienta (ej: stringer, govulncheck) en go.mod")
	fmt.Println("go tool <nombre>   → ejecuta una herramienta declarada en go.mod")

	fmt.Println("\n=== 5. Generación de código, Build Tags y Cross-Compilation ===")
	fmt.Println("go generate ./...  → ejecuta directivas //go:generate (ej: generar mocks o métodos String)")
	fmt.Println("//go:build linux   → (arriba del archivo) compila ese archivo solo en Linux")
	fmt.Println("GOOS=linux GOARCH=amd64 go build -o app-linux ./cmd/app")
	fmt.Println("                   → compilación cruzada a otro SO/arquitectura sin instalar nada extra")
	fmt.Println("go build -ldflags=\"-s -w -X main.version=1.2.3\" ./cmd/app")
	fmt.Println("                   → bandera del linker para achicar el binario e inyectar variables")

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Ciclo diario   → go fmt, go vet, go test -race, go fix -diff")
	fmt.Println("  Seguridad      → govulncheck verifica vulnerabilidades reales en tus llamadas")
	fmt.Println("  Cross-compile  → definí GOOS y GOARCH antes de go build")
}
