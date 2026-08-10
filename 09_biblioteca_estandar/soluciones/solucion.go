package solucion

import (
	"fmt"
	"io"
	"strings"
)

func CopiarMayusculas(destino io.Writer, origen io.Reader) error {
	contenido, err := io.ReadAll(origen)
	if err != nil {
		return fmt.Errorf("leer: %w", err)
	}
	if _, err := io.WriteString(destino, strings.ToUpper(string(contenido))); err != nil {
		return fmt.Errorf("escribir: %w", err)
	}
	return nil
}
