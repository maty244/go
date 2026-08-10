//go:build reto

package reto

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopiarMayusculas(t *testing.T) {
	var salida bytes.Buffer
	if err := CopiarMayusculas(&salida, strings.NewReader("Go y ñ")); err != nil {
		t.Fatal(err)
	}
	if salida.String() != "GO Y Ñ" {
		t.Fatalf("salida=%q", salida.String())
	}
}
