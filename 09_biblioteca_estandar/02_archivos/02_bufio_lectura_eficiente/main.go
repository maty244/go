package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// =========================================================
// bufio: LECTURA EFICIENTE, LÍNEA POR LÍNEA
// =========================================================
// os.ReadFile (tema 01) carga TODO el archivo en memoria de una
// vez. Para archivos grandes (logs de gigabytes, por ejemplo), eso
// puede ser un problema. bufio permite leer de a PEDAZOS (o línea
// por línea), sin necesitar todo el archivo en memoria al mismo
// tiempo.
//
// Ya usaste bufio.Scanner en 01_fundamentos/12_entrada_y_salida_fmt (entrada por consola,
// leyendo de os.Stdin). Es EXACTAMENTE el mismo mecanismo, pero
// leyendo de un archivo en vez del teclado: bufio no le importa
// DE DÓNDE viene el texto, mientras sea un io.Reader (la interfaz
// que representa "algo de donde se puede leer", que vas a ver
// mejor cuando lleguemos a HTTP/).

func main() {
	ruta := filepath.Join(os.TempDir(), "kiosco_ventas.txt")
	defer os.Remove(ruta)

	// Preparamos un archivo de varias líneas para leer después.
	contenido := "Alfajor,800,15\nGaseosa,1200,5\nNotebook,450000,3\nMouse,8500,20\n"
	os.WriteFile(ruta, []byte(contenido), 0644)

	// ─────────────────────────────────────────────────────────
	// bufio.Scanner: LEER LÍNEA POR LÍNEA
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== bufio.Scanner: leer un archivo línea por línea ===")

	archivo, err := os.Open(ruta) // os.Open abre el archivo (no lo carga entero)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer archivo.Close() // SIEMPRE cerrar lo que abrís (05_funciones_y_closures/04_defer: defer)

	scanner := bufio.NewScanner(archivo)
	numeroLinea := 1
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Printf("  Línea %d: %s\n", numeroLinea, linea)
		numeroLinea++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error leyendo:", err)
	}

	// ─────────────────────────────────────────────────────────
	// PROCESAR MIENTRAS SE LEE (sin cargar todo en memoria)
	// ─────────────────────────────────────────────────────────
	// La ventaja real de bufio: podés procesar cada línea A MEDIDA
	// que se lee, sin esperar a tener el archivo completo.

	fmt.Println("\n=== Procesando línea por línea (suma de ventas) ===")

	archivo2, _ := os.Open(ruta)
	defer archivo2.Close()

	scanner2 := bufio.NewScanner(archivo2)
	totalStock := 0
	for scanner2.Scan() {
		partes := strings.Split(scanner2.Text(), ",")
		if len(partes) == 3 {
			var stock int
			fmt.Sscanf(partes[2], "%d", &stock)
			totalStock += stock
		}
	}
	fmt.Println("Stock total en el archivo:", totalStock)

	// ─────────────────────────────────────────────────────────
	// bufio.NewWriter: ESCRITURA EFICIENTE (buffer en memoria)
	// ─────────────────────────────────────────────────────────
	// Simétrico a Scanner: en vez de escribir al disco en CADA
	// línea (lento), acumula en un buffer y escribe en bloques.
	// Hay que llamar Flush() al final para asegurar que todo se
	// grabó realmente.

	fmt.Println("\n=== bufio.Writer: escritura eficiente ===")

	rutaSalida := filepath.Join(os.TempDir(), "kiosco_reporte.txt")
	defer os.Remove(rutaSalida)

	salida, _ := os.Create(rutaSalida)
	defer salida.Close()

	writer := bufio.NewWriter(salida)
	for i := 1; i <= 3; i++ {
		fmt.Fprintf(writer, "Línea de reporte %d\n", i)
	}
	writer.Flush() // sin esto, podría no llegar a escribirse todo al disco

	datosReporte, _ := os.ReadFile(rutaSalida)
	fmt.Println(string(datosReporte))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.Open(ruta)          → abre el archivo (no lo carga entero)")
	fmt.Println("  bufio.NewScanner(f)    → lee línea por línea, eficiente en memoria")
	fmt.Println("  scanner.Scan()/.Text() → avanza una línea / obtiene su texto")
	fmt.Println("  bufio.NewWriter(f)     → escribe en bloques, más rápido que línea a línea")
	fmt.Println("  writer.Flush()         → OBLIGATORIO al final, para grabar lo pendiente")
	fmt.Println("  defer archivo.Close()  → siempre cerrar lo que abrís")
}
