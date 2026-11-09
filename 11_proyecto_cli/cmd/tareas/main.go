package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"gocito/11_proyecto_cli/internal/almacenjson"
	"gocito/11_proyecto_cli/internal/tareas"
)

func main() {
	// Args[0] es el nombre del programa; [1:] conserva solo lo escrito por
	// el usuario. Stdout es un archivo que implementa io.Writer.
	if err := ejecutar(os.Args[1:], os.Stdout); err != nil {
		// Separamos errores de salida normal para poder redirigirlos. os.Exit
		// no ejecuta defers; la persistencia ya cerró sus recursos al retornar.
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func ejecutar(argumentos []string, salida io.Writer) error {
	// Pasar argumentos y Writer permite usar []string y bytes.Buffer en tests,
	// sin depender de una terminal ni lanzar otro proceso.
	flags := flag.NewFlagSet("tareas", flag.ContinueOnError)
	// ContinueOnError devuelve errores de Parse en vez de terminar el proceso.
	flags.SetOutput(salida)

	// String retorna *string. Parse escribe allí el flag o deja el default.
	// El *ruta de abajo desreferencia ese puntero.
	ruta := flags.String("archivo", "tareas.json", "ruta del archivo JSON")
	if err := flags.Parse(argumentos); err != nil {
		return err
	}
	args := flags.Args() // argumentos que quedaron después de los flags
	// Parse deja de buscar flags en el primer argumento no-flag: escribí
	// -archivo ANTES de agregar/listar/completar.
	if len(args) == 0 {
		return errors.New("uso: tareas [-archivo ruta] agregar|listar|completar")
	}

	almacen := almacenjson.Almacen{Ruta: *ruta}
	estado, err := almacen.Cargar()
	if err != nil {
		return err
	}
	gestor := tareas.Nuevo(estado)
	// Almacen transforma JSON ↔ Estado; Gestor opera con valores Go. Esa
	// separación permite estudiar y probar cada responsabilidad por separado.

	switch args[0] {
	case "agregar":
		// Las comillas de la terminal mantienen el título como un argumento.
		// No forman parte del string que recibe Go.
		if len(args) != 2 {
			return errors.New("uso: tareas agregar \"título\"")
		}
		tarea, err := gestor.Agregar(args[1], time.Now())
		if err != nil {
			return err
		}
		if err := almacen.Guardar(gestor.Estado()); err != nil {
			return err
		}
		fmt.Fprintf(salida, "tarea %d agregada\n", tarea.ID)
	case "listar":
		if len(args) != 1 {
			return errors.New("uso: tareas listar")
		}
		for _, tarea := range gestor.Listar() {
			marca := "[ ]"
			if tarea.Completada {
				marca = "[x]"
			}
			fmt.Fprintf(salida, "%s %d: %s\n", marca, tarea.ID, tarea.Titulo)
		}
	case "completar":
		if len(args) != 2 {
			return errors.New("uso: tareas completar ID")
		}
		id, err := strconv.ParseInt(args[1], 10, 64) // texto decimal → int64
		if err != nil || id < 1 {
			return fmt.Errorf("ID inválido: %q", args[1])
		}
		if _, err := gestor.Completar(id, time.Now()); err != nil {
			return err
		}
		if err := almacen.Guardar(gestor.Estado()); err != nil {
			return err
		}
		fmt.Fprintf(salida, "tarea %d completada\n", id)
	default:
		return fmt.Errorf("comando desconocido: %q", args[0])
	}
	return nil
}
