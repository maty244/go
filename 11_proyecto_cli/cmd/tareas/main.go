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
	if err := ejecutar(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func ejecutar(argumentos []string, salida io.Writer) error {
	flags := flag.NewFlagSet("tareas", flag.ContinueOnError)
	flags.SetOutput(salida)
	ruta := flags.String("archivo", "tareas.json", "ruta del archivo JSON")
	if err := flags.Parse(argumentos); err != nil {
		return err
	}
	args := flags.Args()
	if len(args) == 0 {
		return errors.New("uso: tareas [-archivo ruta] agregar|listar|completar")
	}

	almacen := almacenjson.Almacen{Ruta: *ruta}
	estado, err := almacen.Cargar()
	if err != nil {
		return err
	}
	gestor := tareas.Nuevo(estado)

	switch args[0] {
	case "agregar":
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
		id, err := strconv.ParseInt(args[1], 10, 64)
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
