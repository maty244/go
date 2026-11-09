package main

import (
	"fmt"
	"os"
)

// =========================================================
// os.Root: CONFINAR OPERACIONES A UN DIRECTORIO (Go 1.24+)
// =========================================================
// Un riesgo clásico de seguridad al trabajar con rutas que manda
// un usuario (por ejemplo, descargar o subir un archivo en una API)
// es el ataque de "path traversal": el usuario pide "../../etc/passwd"
// o usa un enlace simbólico (symlink) que apunta afuera de la
// carpeta permitida.
//
// os.OpenRoot(directorio) devuelve un *os.Root que confina todas
// las operaciones (Open, Create, ReadFile, WriteFile, MkdirAll,
// Remove, Stat...) estrictamente dentro de ese árbol de directorios.
// Si cualquier ruta o symlink intenta escapar hacia afuera, el
// sistema operativo y Go lo rechazan con error automáticamente.

func main() {
	fmt.Println("=== os.OpenRoot: confinar archivos a un directorio ===")

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

	// Intentar escapar del directorio raíz con "../" falla siempre:
	_, err = raiz.ReadFile("../fuera.txt")
	fmt.Println("¿os.Root rechazó escapar con ../?", err != nil)

	fmt.Println("\nos.Root limita operaciones a un árbol de directorios.")
	fmt.Println("No reemplaza permisos del sistema ni vuelve confiable cualquier archivo.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.OpenRoot(dir)       → abre un árbol confinado (*os.Root)")
	fmt.Println("  defer raiz.Close()     → liberar el descriptor del directorio")
	fmt.Println("  raiz.ReadFile/WriteFile → operan SOLO dentro de dir")
	fmt.Println("  Protección             → bloquea '../' y symlinks que salgan de dir")
}
