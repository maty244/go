package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ContarBytes depende de io.Reader, no de archivos, HTTP o memoria. Cualquier
// tipo que implemente Read puede usarse sin adaptar la función.
func ContarBytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}

func main() {
	fuente := strings.NewReader("Go conecta componentes con Reader y Writer")
	var destino bytes.Buffer

	n, err := io.Copy(&destino, fuente)
	if err != nil {
		fmt.Println("copiar:", err)
		return
	}
	fmt.Printf("Copiados %d bytes: %q\n", n, destino.String())

	// LimitReader evita consumir más de lo permitido. Es útil para archivos y
	// cuerpos HTTP cuyo tamaño controla otra persona.
	limitado := io.LimitReader(strings.NewReader("0123456789"), 4)
	datos, err := io.ReadAll(limitado)
	if err != nil {
		fmt.Println("leer:", err)
		return
	}
	fmt.Println("LimitReader:", string(datos))

	total, err := ContarBytes(strings.NewReader("hola"))
	if err != nil {
		fmt.Println("contar:", err)
		return
	}
	fmt.Println("ContarBytes:", total)
}
