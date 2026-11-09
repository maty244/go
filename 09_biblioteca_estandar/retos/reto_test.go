//go:build reto

package reto

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func TestCopiarMayusculas(t *testing.T) {
	// OneByteReader parte incluso los bytes de ñ o 🙂 entre lecturas.
	for _, entrada := range []string{"", "Go ñ 🙂", "e\u0301", "\ufffd", string([]byte{0xff})} {
		var salida bytes.Buffer
		err := CopiarMayusculas(&salida, iotest.OneByteReader(strings.NewReader(entrada)))
		if esperado := strings.ToUpper(entrada); err != nil || salida.String() != esperado {
			t.Fatalf("entrada=%q salida=%q quiero=%q err=%v", entrada, salida.String(), esperado, err)
		}
	}
}

type lectorFallido struct{ err error }

func (l lectorFallido) Read([]byte) (int, error) { return 0, l.err }

type escritorFallido struct{ err error }

func (e escritorFallido) Write([]byte) (int, error) { return 0, e.err }

func TestEscribeAntesDelFinalYPropagaError(t *testing.T) {
	fallo := errors.New("lectura interrumpida")
	entrada := io.MultiReader(strings.NewReader("ñ"), lectorFallido{fallo})
	var salida bytes.Buffer
	err := CopiarMayusculas(&salida, entrada)
	// ReadAll esperaría al final y no escribiría esta salida parcial.
	if !errors.Is(err, fallo) || salida.String() != "Ñ" {
		t.Fatalf("salida=%q err=%v; quiero salida parcial y error original", salida.String(), err)
	}
}

func TestErroresDeEscritura(t *testing.T) {
	fallo := errors.New("destino cerrado")
	if err := CopiarMayusculas(escritorFallido{fallo}, strings.NewReader("Go")); !errors.Is(err, fallo) {
		t.Fatalf("se perdió el error de escritura: %v", err)
	}
	if err := CopiarMayusculas(escritorFallido{}, strings.NewReader("Go")); !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("no se detectó escritura corta: %v", err)
	}
}
