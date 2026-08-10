package solucion

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopiarMayusculas(t *testing.T) {
	var salida bytes.Buffer
	if err := CopiarMayusculas(&salida, strings.NewReader("ñ")); err != nil || salida.String() != "Ñ" {
		t.Fatalf("salida=%q err=%v", salida.String(), err)
	}
}
