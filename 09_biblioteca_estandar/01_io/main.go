package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// =========================================================
// EL CONTRATO MÁS IMPORTANTE DE GO: io.Reader E io.Writer
// =========================================================
// Gran parte del poder de la biblioteca estándar de Go viene de
// dos interfaces diminutas (de un solo método cada una):
//
//   type Reader interface {
//       Read(p []byte) (n int, err error)
//   }
//
//   type Writer interface {
//       Write(p []byte) (n int, err error)
//   }
//
// ¿Por qué son tan importantes? Porque un archivo (*os.File), el
// cuerpo de una petición HTTP (req.Body), una conexión de red
// (net.Conn), un archivo comprimido (gzip.Reader) o un simple
// string en memoria (*strings.Reader) implementan io.Reader.
// Si tu función recibe un io.Reader, funciona con TODOS ellos
// sin cambiar una sola línea.
//
// DETALLE CLAVE DEL CONTRATO DE Read:
// Read puede devolver datos (n > 0) Y un error (por ejemplo,
// io.EOF al llegar al final) EN LA MISMA LLAMADA. Por eso quien
// implementa un loop de lectura manual SIEMPRE debe procesar
// primero los p[:n] bytes recibidos antes de fijarse en err.

// ContarBytes depende de io.Reader, no de archivos, HTTP o memoria. Cualquier
// tipo que implemente Read puede usarse sin adaptar la función.
func ContarBytes(r io.Reader) (int64, error) {
	// io.Copy copia de a bloques (por defecto usa un buffer de 32 KB,
	// sin cargar todo en memoria) desde un Reader hacia un Writer.
	// io.Discard es un Writer especial que "tira a la basura" todo
	// lo que recibe sin gastar memoria, pero cuenta los bytes.
	return io.Copy(io.Discard, r)
}

func main() {
	fmt.Println("=== io.Copy: conectar un Reader con un Writer ===")

	fuente := strings.NewReader("Go conecta componentes con Reader y Writer")
	var destino bytes.Buffer

	// Pasamos &destino porque los métodos Write/Read de bytes.Buffer
	// tienen receptor de puntero (*bytes.Buffer) para poder crecer el buffer.
	n, err := io.Copy(&destino, fuente)
	if err != nil {
		fmt.Println("copiar:", err)
		return
	}
	fmt.Printf("Copiados %d bytes: %q\n", n, destino.String())

	// ─────────────────────────────────────────────────────────
	// io.LimitReader: PROTEGERSE DE ENTRADAS GIGANTES
	// ─────────────────────────────────────────────────────────
	// LimitReader envuelve un Reader y devuelve io.EOF apenas se
	// leyeron N bytes. Es útil para archivos y cuerpos HTTP cuyo
	// tamaño controla otra persona.
	//
	// OJO: por sí solo, LimitReader(r, 4) NO distingue entre una
	// entrada que medía exactamente 4 bytes y una más larga que se
	// cortó al llegar a 4 (en ambos casos termina con io.EOF).
	// Truco común: si el máximo válido es 4, leés con LimitReader(r, 5);
	// si recibís 5 bytes, sabés que el cliente se pasó del límite.

	fmt.Println("\n=== io.LimitReader: poner un tope de lectura ===")
	limitado := io.LimitReader(strings.NewReader("0123456789"), 4)
	datos, err := io.ReadAll(limitado)
	if err != nil {
		fmt.Println("leer:", err)
		return
	}
	fmt.Println("LimitReader:", string(datos))

	fmt.Println("\n=== ContarBytes sobre cualquier io.Reader ===")
	total, err := ContarBytes(strings.NewReader("hola"))
	if err != nil {
		fmt.Println("contar:", err)
		return
	}
	fmt.Println("ContarBytes:", total)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  io.Reader              → interfaz con Read(p []byte) (n int, err error)")
	fmt.Println("  io.Writer              → interfaz con Write(p []byte) (n int, err error)")
	fmt.Println("  Regla de Read          → procesar los n > 0 bytes ANTES de mirar err")
	fmt.Println("  io.Copy(dst, src)      → copia en streaming sin cargar todo en memoria")
	fmt.Println("  io.LimitReader(r, max) → corta con io.EOF al llegar a max bytes")
}
