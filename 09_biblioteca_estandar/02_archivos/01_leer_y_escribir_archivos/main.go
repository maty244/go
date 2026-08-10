package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// =========================================================
// LEER Y ESCRIBIR ARCHIVOS: os.WriteFile / os.ReadFile
// =========================================================
// El paquete "os" trae las dos funciones más simples para trabajar
// con archivos completos (no gigantes): leer todo el contenido de
// una vez, o escribir todo el contenido de una vez.
//
//   os.WriteFile(ruta, datos []byte, permisos) error
//   os.ReadFile(ruta) ([]byte, error)
//
// Para archivos MUY grandes (que no entran cómodos en memoria)
// existe una forma más eficiente con bufio (tema 02). Para la
// mayoría de los casos (configs, logs chicos, exportar un reporte)
// estas dos funciones alcanzan y sobran.

func main() {
	// Usamos una carpeta temporal para no ensuciar el repo con
	// archivos de prueba. os.TempDir() da la carpeta temporal del
	// sistema operativo (funciona igual en Mac, Linux y Windows).
	ruta := filepath.Join(os.TempDir(), "kiosco_notas.txt")
	defer os.Remove(ruta) // limpiamos al terminar, buena práctica en ejemplos

	// ─────────────────────────────────────────────────────────
	// ESCRIBIR: os.WriteFile
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== os.WriteFile: escribir un archivo ===")

	contenido := "Pedido #4521\nCliente: Matias\nTotal: $15000.50\n"

	// 0644 son los permisos del archivo (lectura/escritura para el
	// dueño, solo lectura para el resto). Es la notación octal
	// estándar de Unix; en Windows este valor se ignora en la práctica.
	err := os.WriteFile(ruta, []byte(contenido), 0644)
	if err != nil {
		fmt.Println("Error escribiendo:", err)
		return
	}
	fmt.Println("Archivo escrito en:", ruta)

	// ─────────────────────────────────────────────────────────
	// LEER: os.ReadFile
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== os.ReadFile: leer un archivo completo ===")

	datos, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error leyendo:", err)
		return
	}
	fmt.Println(string(datos))

	// ─────────────────────────────────────────────────────────
	// QUÉ PASA SI EL ARCHIVO NO EXISTE
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Leer un archivo que no existe ===")

	_, err = os.ReadFile(filepath.Join(os.TempDir(), "no_existe_seguro.txt"))
	if err != nil {
		fmt.Println("Error:", err)
		// os.IsNotExist es la forma clásica de chequear ESTE error
		// específico (también se puede con errors.Is y os.ErrNotExist)
		fmt.Println("¿Es porque no existe?", os.IsNotExist(err))
	}

	// ─────────────────────────────────────────────────────────
	// SOBRESCRIBIR: WriteFile REEMPLAZA el contenido completo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== WriteFile sobrescribe el contenido anterior ===")

	os.WriteFile(ruta, []byte("Contenido nuevo, reemplaza todo"), 0644)
	datosNuevos, _ := os.ReadFile(ruta)
	fmt.Println(string(datosNuevos))

	// ─────────────────────────────────────────────────────────
	// VERIFICAR SI UN ARCHIVO EXISTE, SIN LEERLO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== os.Stat: chequear existencia sin leer ===")

	if _, err := os.Stat(ruta); err == nil {
		fmt.Println("El archivo existe")
	} else if os.IsNotExist(err) {
		fmt.Println("El archivo NO existe")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.WriteFile(ruta, []byte(datos), permisos)  → escribe todo de una vez")
	fmt.Println("  os.ReadFile(ruta)                             → lee todo de una vez")
	fmt.Println("  os.IsNotExist(err)                            → ¿el error es 'no existe'?")
	fmt.Println("  os.Stat(ruta)                                 → info del archivo sin leerlo")
	fmt.Println("  Bueno para                                    → archivos chicos/medianos")
}
