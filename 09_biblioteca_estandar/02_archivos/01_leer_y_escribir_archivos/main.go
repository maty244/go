package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// //go:embed incrusta el contenido de un archivo (o carpeta con embed.FS)
// DENTRO del binario compilado en tiempo de compilación. Así tu ejecutable no
// depende de que el archivo externo exista en el disco cuando lo distribuís.
//
//go:embed main.go
var codigoFuenteIncrustado string

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

	// 0o644 (o el clásico 0644) son los permisos del archivo: lectura/escritura
	// para el dueño, solo lectura para el resto. El prefijo "0o" deja explícito
	// que es un número octal de Unix; en Windows este valor se ignora en la práctica.
	err := os.WriteFile(ruta, []byte(contenido), 0o644)
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
		// En Go moderno se prefiere errors.Is(err, os.ErrNotExist) porque
		// funciona incluso si el error fue envuelto con fmt.Errorf("...: %w", err).
		// También vas a ver la función clásica os.IsNotExist(err) en código existente.
		fmt.Println("¿Es porque no existe? (errors.Is):", errors.Is(err, os.ErrNotExist))
		fmt.Println("¿Es porque no existe? (os.IsNotExist):", os.IsNotExist(err))
	}

	// ─────────────────────────────────────────────────────────
	// SOBRESCRIBIR: WriteFile REEMPLAZA el contenido completo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== WriteFile sobrescribe el contenido anterior ===")

	os.WriteFile(ruta, []byte("Contenido nuevo, reemplaza todo"), 0o644)
	datosNuevos, _ := os.ReadFile(ruta)
	fmt.Println(string(datosNuevos))

	// ─────────────────────────────────────────────────────────
	// VERIFICAR SI UN ARCHIVO EXISTE, SIN LEERLO
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== os.Stat: chequear existencia sin leer ===")

	if _, err := os.Stat(ruta); err == nil {
		fmt.Println("El archivo existe")
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("El archivo NO existe")
	}

	// ─────────────────────────────────────────────────────────
	// INCRUSTAR ARCHIVOS EN EL BINARIO: //go:embed
	// ─────────────────────────────────────────────────────────
	// Si tenés plantillas, consultas SQL, configuraciones por defecto o
	// recursos estáticos que siempre viajan con tu programa, podés usar
	// la directiva //go:embed (mirá arriba de main) para meterlos directo
	// dentro del ejecutable al compilar (como string, []byte o embed.FS).
	fmt.Println("\n=== //go:embed: archivos incrustados en el binario ===")
	primeraLinea, _, _ := strings.Cut(codigoFuenteIncrustado, "\n")
	fmt.Printf("Bytes incrustados de main.go: %d | 1ra línea: %q\n", len(codigoFuenteIncrustado), primeraLinea)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.WriteFile(ruta, []byte(datos), 0o644)      → escribe todo de una vez")
	fmt.Println("  os.ReadFile(ruta)                             → lee todo de una vez")
	fmt.Println("  errors.Is(err, os.ErrNotExist)                → ¿el error es 'no existe'? (moderno)")
	fmt.Println("  os.Stat(ruta)                                 → info del archivo sin leerlo")
	fmt.Println("  //go:embed archivo                            → incrusta archivo en el binario al compilar")
	fmt.Println("  Bueno para                                    → archivos chicos/medianos")
}
