package main

import (
	"fmt"
	"os"
)

// =========================================================
// os.Args: ARGUMENTOS DE LÍNEA DE COMANDOS
// =========================================================
// Cuando ejecutás un programa Go compilado desde la terminal, le
// podés pasar argumentos extra:
//
//   ./miprograma agregar "Notebook" 450000
//
// os.Args es un []string con TODOS los argumentos, incluyendo el
// nombre del programa como PRIMER elemento (índice 0). Es la base
// para escribir CLIs (herramientas de línea de comandos), uno de
// los puntos fuertes de Go (mencionado en el README: Docker,
// Kubernetes y muchas herramientas están hechas así).

func main() {
	fmt.Println("=== os.Args: todos los argumentos recibidos ===")
	fmt.Println(os.Args)

	fmt.Println("\nos.Args[0] (nombre/ruta del programa):", os.Args[0])
	fmt.Println("Cantidad total de argumentos:", len(os.Args))

	// ─────────────────────────────────────────────────────────
	// LOS ARGUMENTOS "ÚTILES" EMPIEZAN EN EL ÍNDICE 1
	// ─────────────────────────────────────────────────────────
	// os.Args[0] es "quién soy" (el programa), no un dato que el
	// usuario haya escrito. Los argumentos reales están desde el 1.

	fmt.Println("\n=== Argumentos 'reales' (desde el índice 1) ===")

	argumentos := os.Args[1:]
	if len(argumentos) == 0 {
		fmt.Println("  No se pasó ningún argumento extra.")
		fmt.Println(`  Probá correr: go run . agregar Notebook 450000`)
	} else {
		for i, arg := range argumentos {
			fmt.Printf("  argumento[%d]: %s\n", i, arg)
		}
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: UN MINI-CLI CON COMANDOS
	// ─────────────────────────────────────────────────────────
	// El patrón típico de una herramienta de línea de comandos:
	// el primer argumento es el COMANDO, el resto son sus datos.
	// (git funciona así: "git commit -m mensaje" → comando="commit")

	fmt.Println("\n=== Caso real: mini-CLI de kiosco ===")

	if len(argumentos) == 0 {
		fmt.Println("  Uso: go run . <comando> [argumentos]")
		fmt.Println("  Comandos disponibles: agregar, listar, eliminar")
		return
	}

	comando := argumentos[0]
	switch comando {
	case "agregar":
		if len(argumentos) < 3 {
			fmt.Println("  Uso: agregar <nombre> <precio>")
			return
		}
		fmt.Printf("  Agregando producto: %s a $%s\n", argumentos[1], argumentos[2])
	case "listar":
		fmt.Println("  Listando productos... (simulado)")
	case "eliminar":
		if len(argumentos) < 2 {
			fmt.Println("  Uso: eliminar <nombre>")
			return
		}
		fmt.Printf("  Eliminando producto: %s\n", argumentos[1])
	default:
		fmt.Printf("  Comando desconocido: %q\n", comando)
	}

	// ─────────────────────────────────────────────────────────
	// PARA FLAGS CON NOMBRE (--precio=100, -v), EXISTE "flag"
	// ─────────────────────────────────────────────────────────
	// os.Args te da el texto CRUDO. Para algo más prolijo como
	// --nombre=Notebook o -v (banderas con nombre), la librería
	// estándar trae el paquete "flag", que parsea todo eso solo.
	// No lo cubrimos acá, pero ahora sabés que existe si tu CLI
	// crece en complejidad.

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.Args         → []string con TODOS los argumentos")
	fmt.Println("  os.Args[0]      → el programa mismo, no un dato del usuario")
	fmt.Println("  os.Args[1:]     → los argumentos 'reales'")
	fmt.Println("  Patrón CLI      → primer argumento = comando, resto = datos")
	fmt.Println("  Para flags -x   → ver el paquete \"flag\" de la librería estándar")
}
