package solucion

import "testing"

func TestFizzBuzz(t *testing.T) {
	resultado := FizzBuzz(15)
	if len(resultado) != 15 || resultado[14] != "FizzBuzz" {
		t.Fatalf("resultado=%v", resultado)
	}
}
