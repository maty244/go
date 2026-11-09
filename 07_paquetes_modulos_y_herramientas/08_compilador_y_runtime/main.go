package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// =========================================================
// COMPILADOR, RUNTIME Y TÓPICOS AVANZADOS DEL LENGUAJE
// =========================================================
// Para dominar Go a fondo conviene distinguir tres capas:
//   1. Las reglas del lenguaje (la especificación).
//   2. El compilador (traduce código Go a código máquina, hace inlining y
//      decide mediante Escape Analysis si algo vive en el stack o en el heap).
//   3. El runtime (vive dentro de tu binario compilado y gestiona el scheduler
//      de goroutines, la memoria y el Garbage Collector concurrente).
//
// Además, Go incluye mecanismos de bajo nivel o metaprogramación para casos
// especiales: `reflect`, `unsafe` y `cgo`. Conocer qué son —y por qué casi
// nunca los necesitás en código de dominio— cierra el mapa del lenguaje.

// Antes de main se inicializan los paquetes importados y las variables del
// paquete. Dentro de este ejemplo inicializar se ejecuta para asignar etapa.
// init corre después de las variables. No lo llamamos nosotros ni tiene return.
// No conviene esconder configuración compleja en init: acá solo lo observamos.
var etapa = inicializar()

func inicializar() string {
	fmt.Println("1. inicialización de variable del paquete")
	return "lista"
}

func init() {
	fmt.Println("2. init ve la variable:", etapa)
}

// devolverPuntero es válido: una variable local puede seguir viva si alguien
// conserva acceso a ella. Go administra su vida; no queda una dirección inválida
// al retornar. El compilador decide almacenamiento mediante escape analysis:
//
//	go build -gcflags="-m" ./07_paquetes_modulos_y_herramientas/08_compilador_y_runtime
func devolverPuntero() *int {
	valor := 42
	return &valor
}

func main() {
	fmt.Println("3. main")
	fmt.Println("puntero a local sigue válido:", *devolverPuntero()) // 42

	// COMPILADOR: verifica tipos y genera código máquina; puede inlinear,
	// eliminar trabajo o guardar valores en registros/stack/heap. No cambia
	// los resultados garantizados por el lenguaje.
	// RUNTIME: acompaña al programa ejecutándose; administra goroutines,
	// stacks y memoria del heap, entre otras tareas.
	fmt.Println("versión del toolchain:", runtime.Version())
	fmt.Println("CPU lógicas:", runtime.NumCPU())
	// Con 0 solo consultamos el valor actual. No es un límite de goroutines:
	// limita cuántos threads ejecutan código Go simultáneamente. Los threads
	// bloqueados en llamadas al sistema no cuentan para ese límite.
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// El garbage collector recupera almacenamiento del heap que deja de ser
	// alcanzable. No implica liberar inmediatamente memoria al sistema, ni
	// cerrar archivos, ni detener goroutines: eso requiere diseño explícito.
	fmt.Println("GC administra memoria; defer/Close/cancel administran recursos.")

	// ─────────────────────────────────────────────────────────
	// TÓPICOS AVANZADOS: reflect, unsafe y cgo (cuándo existen)
	// ─────────────────────────────────────────────────────────
	// 1. Reflection (paquete "reflect"): permite inspeccionar tipos y valores
	//    en tiempo de ejecución (así funcionan encoding/json o fmt.Printf por
	//    dentro). En tu código preferí siempre interfaces o generics, porque
	//    reflect pierde chequeo estático de tipos y es más lento.
	// 2. Paquete "unsafe": permite saltear la seguridad de tipos de Go, medir
	//    el tamaño exacto en memoria (unsafe.Sizeof) o convertir punteros
	//    (unsafe.Pointer). Se usa solo en la stdlib o código de muy bajo nivel.
	// 3. CGO (`import "C"`): permite llamar bibliotecas en C desde Go. Tiene
	//    costo en cada llamada y complica la compilación cruzada; siempre que
	//    exista un paquete en Go puro, preferilo sobre CGO.
	fmt.Println("\n=== Vistazo a reflect y unsafe ===")
	x := int64(100)
	fmt.Printf("reflect.TypeOf(x) = %v | unsafe.Sizeof(x) = %d bytes\n", reflect.TypeOf(x), unsafe.Sizeof(x))

	fmt.Println("\nPara inspeccionar decisiones del compilador, abrí PRACTICA.md de esta etapa.")

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Orden de arranque    → variables de paquete → init() → main()")
	fmt.Println("  Escape Analysis      → el compilador decide stack vs heap (-gcflags=\"-m\")")
	fmt.Println("  Garbage Collector    → libera memoria inalcanzable, NO cierra archivos/goroutines")
	fmt.Println("  reflect / unsafe / C → herramientas de bajo nivel; preferí tipos, generics y Go puro")
}
