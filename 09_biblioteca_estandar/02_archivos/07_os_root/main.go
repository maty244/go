package main

import (
	"fmt"
	"os"
)

func main() {
	directorio, err := os.MkdirTemp("", "gocito-root-")
	if err != nil {
		fmt.Println("crear temporal:", err)
		return
	}
	defer func() {
		if err := os.RemoveAll(directorio); err != nil {
			fmt.Println("limpiar temporal:", err)
		}
	}()

	raiz, err := os.OpenRoot(directorio)
	if err != nil {
		fmt.Println("abrir raíz:", err)
		return
	}
	defer func() {
		if err := raiz.Close(); err != nil {
			fmt.Println("cerrar raíz:", err)
		}
	}()

	if err := raiz.MkdirAll("datos", 0o755); err != nil {
		fmt.Println("crear carpeta:", err)
		return
	}
	if err := raiz.WriteFile("datos/mensaje.txt", []byte("contenido seguro\n"), 0o600); err != nil {
		fmt.Println("escribir:", err)
		return
	}
	contenido, err := raiz.ReadFile("datos/mensaje.txt")
	if err != nil {
		fmt.Println("leer:", err)
		return
	}
	fmt.Print(string(contenido))

	_, err = raiz.ReadFile("../fuera.txt")
	fmt.Println("¿os.Root rechazó escapar con ../?", err != nil)

	fmt.Println("\nos.Root limita operaciones a un árbol de directorios.")
	fmt.Println("No reemplaza permisos del sistema ni vuelve confiable cualquier archivo.")
}
