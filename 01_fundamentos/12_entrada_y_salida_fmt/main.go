package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// =========================================================
	// ENTRADA Y SALIDA POR CONSOLA
	// =========================================================
	// Para interactuar con el usuario por consola, Go usa principalmente
	// el paquete "fmt" (format).
	// Para entrada de texto más compleja, usamos "bufio" + "os".

	// ─────────────────────────────────────────────────────────
	// SALIDA: fmt.Print, fmt.Println, fmt.Printf
	// ─────────────────────────────────────────────────────────

	fmt.Println("=== Funciones de salida ===")

	// Print: imprime sin salto de línea al final
	fmt.Print("Hola ")
	fmt.Print("Mundo")
	fmt.Print("\n") // salto de línea manual

	// Println: imprime con salto de línea automático al final
	// También agrega espacios entre argumentos
	fmt.Println("Buenos días!")
	fmt.Println("Suma:", 5+3, "| Producto:", 5*3)

	// Printf: imprime con FORMATO usando verbos (%)
	// Es la más poderosa y la más usada en código real.
	nombre := "Matias"
	edad := 25
	saldo := 1234.567
	activo := true

	fmt.Println("\n=== fmt.Printf con verbos ===")
	fmt.Printf("Nombre: %s\n", nombre)         // %s → string
	fmt.Printf("Edad: %d años\n", edad)        // %d → entero decimal
	fmt.Printf("Saldo: $%.2f\n", saldo)        // %f → float, .2 = 2 decimales
	fmt.Printf("Activo: %v\n", activo)         // %v → valor genérico (cualquier tipo)
	fmt.Printf("Activo: %t\n", activo)         // %t → bool específico
	fmt.Printf("Tipo: %T\n", saldo)            // %T → tipo de la variable
	fmt.Printf("Binario: %b\n", 42)            // %b → binario
	fmt.Printf("Octal: %o\n", 42)              // %o → octal
	fmt.Printf("Hexadecimal: %x\n", 255)       // %x → hex minúscula
	fmt.Printf("Hexadecimal: %X\n", 255)       // %X → hex mayúscula
	fmt.Printf("Científico: %e\n", 123456.789) // %e → notación científica
	fmt.Printf("Carácter: %c\n", 65)           // %c → caracter Unicode
	fmt.Printf("Con comillas: %q\n", nombre)   // %q → string con comillas

	// Alineación y relleno
	fmt.Println("\n=== Alineación con Printf ===")
	fmt.Printf("|%-10s|%10s|\n", "izquierda", "derecha") // - = alinear izquierda
	fmt.Printf("|%-10d|%10d|\n", 42, 42)
	fmt.Printf("|%010d|\n", 42)          // rellenar con ceros
	fmt.Printf("|%+d| |%+d|\n", 42, -42) // mostrar signo siempre

	// Tabla formateada (ejemplo real)
	fmt.Println("\n=== Tabla de productos ===")
	fmt.Printf("%-15s %8s %6s\n", "Producto", "Precio", "Stock")
	fmt.Println(strings.Repeat("-", 32))
	productos := []struct {
		nombre string
		precio float64
		stock  int
	}{
		{"Notebook", 1500.00, 5},
		{"Mouse", 25.99, 42},
		{"Teclado", 75.50, 18},
		{"Monitor", 450.00, 3},
	}
	for _, p := range productos {
		fmt.Printf("%-15s %8.2f %6d\n", p.nombre, p.precio, p.stock)
	}

	// ─────────────────────────────────────────────────────────
	// SPRINTF: formatear a string sin imprimir
	// ─────────────────────────────────────────────────────────
	// Mismo que Printf pero retorna el string en vez de imprimirlo.
	// Muy útil para construir mensajes, URLs, nombres de archivos, etc.

	fmt.Println("\n=== fmt.Sprintf ===")
	mensaje := fmt.Sprintf("Hola %s, tenés %d mensajes nuevos.", nombre, 3)
	fmt.Println(mensaje)

	url := fmt.Sprintf("https://api.ejemplo.com/usuarios/%d/perfil", 42)
	fmt.Println("URL generada:", url)

	archivoLog := fmt.Sprintf("log_%s_%d.txt", "2026-01-15", 001)
	fmt.Println("Archivo:", archivoLog)

	// ─────────────────────────────────────────────────────────
	// ERRORF: crear errores con formato
	// ─────────────────────────────────────────────────────────
	err := fmt.Errorf("usuario con ID %d no encontrado en la base de datos", 99)
	fmt.Println("\nError creado:", err)

	// ─────────────────────────────────────────────────────────
	// FPRINTLN / FPRINTF: escribir a otros destinos
	// ─────────────────────────────────────────────────────────
	// Por defecto fmt escribe a os.Stdout (la consola normal).
	// Podemos escribir a os.Stderr (salida de errores) o cualquier io.Writer.

	fmt.Fprintln(os.Stderr, "Este mensaje va al canal de errores (stderr)")
	fmt.Fprintf(os.Stdout, "Esto va explícitamente a stdout: %s\n", "hola")

	// ─────────────────────────────────────────────────────────
	// ENTRADA: fmt.Scan, fmt.Scanf, fmt.Scanln
	// ─────────────────────────────────────────────────────────
	// NOTA: En este archivo los ejemplos de Scan están comentados
	// porque necesitan input interactivo del usuario.
	// Descomentalos para probarlos.

	fmt.Println("\n=== Lectura de entrada (comentada para demo) ===")
	fmt.Println("Para leer un entero: fmt.Scan(&miVariable)")
	fmt.Printf("Para leer con formato: fmt.Scanf(%q, &num, &texto)\n", "%d %s")

	/*
		// EJEMPLO INTERACTIVO (descomentar para usar):

		var n int
		fmt.Print("Ingresá un número: ")
		fmt.Scan(&n) // el & es el operador de dirección (referencia)
		fmt.Printf("Ingresaste: %d\n", n)

		var nombreUsuario string
		fmt.Print("Ingresá tu nombre: ")
		fmt.Scan(&nombreUsuario) // Scan se detiene en espacios
		fmt.Printf("Hola, %s!\n", nombreUsuario)

		// Leer múltiples valores en una línea
		var x, y int
		fmt.Print("Ingresá dos números separados por espacio: ")
		fmt.Scan(&x, &y)
		fmt.Printf("Suma: %d\n", x+y)
	*/

	// ─────────────────────────────────────────────────────────
	// ENTRADA DE TEXTO CON ESPACIOS: bufio.Scanner
	// ─────────────────────────────────────────────────────────
	// fmt.Scan se detiene en el primer espacio.
	// Para leer líneas completas (con espacios), usamos bufio.Scanner.

	fmt.Println("\n=== bufio.Scanner para leer líneas ===")
	fmt.Println("(ejemplo con string simulado)")

	// Simulamos stdin con un string para el ejemplo
	input := strings.NewReader("Hola Mundo Go\nSegunda línea\n")
	scanner := bufio.NewScanner(input)

	lineaNum := 1
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Printf("Línea %d: '%s'\n", lineaNum, linea)
		lineaNum++
	}

	/*
		// VERSIÓN INTERACTIVA con os.Stdin (descomentar para usar):
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Escribí tu nombre completo: ")
		scanner.Scan()
		nombreCompleto := scanner.Text()
		fmt.Printf("Hola, %s!\n", nombreCompleto)
	*/

	// ─────────────────────────────────────────────────────────
	// MINI APLICACIÓN: Calculadora por consola
	// ─────────────────────────────────────────────────────────
	// Este es cómo se vería una mini app real de consola.
	// Está con valores hardcodeados para que puedas ejecutar sin interacción.

	fmt.Println("\n=== Mini Calculadora (valores simulados) ===")
	simularCalculadora(15.0, "+", 5.0)
	simularCalculadora(15.0, "-", 5.0)
	simularCalculadora(15.0, "*", 5.0)
	simularCalculadora(15.0, "/", 5.0)
	simularCalculadora(15.0, "/", 0.0) // división por cero
}

func simularCalculadora(a float64, operacion string, b float64) {
	fmt.Printf("%.0f %s %.0f = ", a, operacion, b)
	switch operacion {
	case "+":
		fmt.Printf("%.2f\n", a+b)
	case "-":
		fmt.Printf("%.2f\n", a-b)
	case "*":
		fmt.Printf("%.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: no se puede dividir por cero")
		} else {
			fmt.Printf("%.2f\n", a/b)
		}
	default:
		fmt.Println("Operación no soportada")
	}
}
