package solucion

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	esperado := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if resultado := FizzBuzz(15); !slices.Equal(resultado, esperado) {
		t.Fatalf("resultado=%v", resultado)
	}
	if vacio := FizzBuzz(0); len(vacio) != 0 {
		t.Fatalf("para límite 0 se esperaba vacío, obtenido=%v", vacio)
	}
	if negativo := FizzBuzz(-3); len(negativo) != 0 {
		t.Fatalf("para límite negativo se esperaba vacío, obtenido=%v", negativo)
	}
}
