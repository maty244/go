package solucion

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"unicode"
)

// CopiarMayusculas procesa runes sin cargar toda la entrada. ReadRune conserva
// bytes entre lecturas, así que una ñ partida en dos Read se decodifica bien.
// La memoria auxiliar no crece con el texto; el Writer elegido puede acumularlo.
// Como strings.ToUpper, transforma puntos de código, no grafemas. ReadRune
// reemplaza codificaciones inválidas por U+FFFD; esta solución no las rechaza.
func CopiarMayusculas(destino io.Writer, origen io.Reader) error {
	lector := bufio.NewReader(origen)
	for {
		r, _, err := lector.ReadRune()
		if errors.Is(err, io.EOF) {
			return nil // EOF es el final normal, no un fallo
		}
		if err != nil {
			return fmt.Errorf("leer: %w", err)
		}
		texto := string(unicode.ToUpper(r))
		n, err := io.WriteString(destino, texto)
		if err != nil {
			return fmt.Errorf("escribir: %w", err)
		}
		if n != len(texto) {
			return fmt.Errorf("escribir: %w", io.ErrShortWrite)
		}
	}
}
