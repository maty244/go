//go:build reto

package reto

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	esperado := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if obtenido := FizzBuzz(15); !slices.Equal(obtenido, esperado) {
		t.Fatalf("obtenido=%v", obtenido)
	}
	if vacio := FizzBuzz(0); len(vacio) != 0 {
		t.Fatalf("para límite 0 se esperaba vacío, obtenido=%v", vacio)
	}
}
