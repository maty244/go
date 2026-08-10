package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// =========================================================
// LEER Y ESCRIBIR CSV: encoding/csv
// =========================================================
// CSV (Comma-Separated Values) es un formato de texto MUY común
// para datos tabulares: exportar/importar desde Excel, Google
// Sheets, reportes de otros sistemas. En el tema 02 leímos texto
// separado por comas "a mano" con strings.Split; el paquete
// encoding/csv hace eso mismo pero maneja bien los casos raros
// (comas dentro de un campo con comillas, saltos de línea dentro
// de un campo, etc.) que un Split simple rompería.

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func main() {
	ruta := filepath.Join(os.TempDir(), "kiosco_catalogo.csv")
	defer os.Remove(ruta)

	// ─────────────────────────────────────────────────────────
	// ESCRIBIR UN CSV: csv.Writer
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Escribir un CSV ===")

	catalogo := []Producto{
		{Nombre: "Alfajor", Precio: 800, Stock: 15},
		{Nombre: "Gaseosa", Precio: 1200, Stock: 5},
		{Nombre: "Notebook", Precio: 450000, Stock: 3},
	}

	archivo, err := os.Create(ruta)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	writer := csv.NewWriter(archivo)

	// Encabezado
	writer.Write([]string{"nombre", "precio", "stock"})

	// Una fila por producto: csv.Writer espera []string
	for _, p := range catalogo {
		fila := []string{
			p.Nombre,
			strconv.FormatFloat(p.Precio, 'f', 2, 64),
			strconv.Itoa(p.Stock),
		}
		writer.Write(fila)
	}

	writer.Flush() // igual que bufio.Writer (tema 02): obligatorio al final
	archivo.Close()

	datosCSV, _ := os.ReadFile(ruta)
	fmt.Println(string(datosCSV))

	// ─────────────────────────────────────────────────────────
	// LEER UN CSV: csv.Reader
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Leer un CSV ===")

	archivoLectura, err := os.Open(ruta)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer archivoLectura.Close()

	reader := csv.NewReader(archivoLectura)

	// ReadAll carga todas las filas de una vez, cada una como []string
	filas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error leyendo CSV:", err)
		return
	}

	fmt.Println("Encabezado:", filas[0])
	fmt.Println("Filas de datos:")
	for _, fila := range filas[1:] { // saltamos el encabezado
		fmt.Println(" ", fila)
	}

	// ─────────────────────────────────────────────────────────
	// CONVERTIR LAS FILAS DE VUELTA A []Producto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Convertir filas CSV a structs ===")

	var catalogoLeido []Producto
	for _, fila := range filas[1:] {
		precio, _ := strconv.ParseFloat(fila[1], 64)
		stock, _ := strconv.Atoi(fila[2])
		catalogoLeido = append(catalogoLeido, Producto{
			Nombre: fila[0],
			Precio: precio,
			Stock:  stock,
		})
	}

	for _, p := range catalogoLeido {
		fmt.Printf("  %s: $%.2f (stock: %d)\n", p.Nombre, p.Precio, p.Stock)
	}

	// ─────────────────────────────────────────────────────────
	// CSV CON COMAS DENTRO DE UN CAMPO: por qué NO usar strings.Split
	// ─────────────────────────────────────────────────────────
	// Si un campo necesita tener una coma, el CSV lo pone entre
	// comillas: "Notebook, 15 pulgadas",450000,3
	// strings.Split(",") rompería esto (partiría el campo en dos).
	// csv.Reader lo maneja correctamente.

	fmt.Println("\n=== csv.Reader maneja comas dentro de comillas ===")

	textoConComa := `nombre,precio
"Notebook, 15 pulgadas",450000
`
	readerEspecial := csv.NewReader(strings.NewReader(textoConComa))
	filasEspeciales, _ := readerEspecial.ReadAll()
	fmt.Println("Fila con coma interna:", filasEspeciales[1])

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  csv.NewWriter(archivo)  → escribir filas ([]string) a CSV")
	fmt.Println("  writer.Write(fila)      → una fila por llamada")
	fmt.Println("  writer.Flush()          → obligatorio al final")
	fmt.Println("  csv.NewReader(archivo)  → leer CSV")
	fmt.Println("  reader.ReadAll()        → todas las filas como [][]string")
	fmt.Println("  Ventaja sobre Split     → maneja comas y comillas dentro de un campo")
}
