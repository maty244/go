package solucion

import "testing"

func TestEsPalindromo(t *testing.T) {
	casos := []struct {
		nombre   string
		texto    string
		esperado bool
	}{
		{"frase con espacios y mayúsculas", "Anita lava la tina", true},
		{"UTF-8 con eñe", "Aña", true},
		{"UTF-8 con tilde y mayúscula", "Áñá", true},
		{"palabra capicúa", "Reconocer", true},
		{"texto vacío", "", true},
		{"no es palíndromo", "Golang", false},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			if obtenido := EsPalindromo(c.texto); obtenido != c.esperado {
				t.Errorf("EsPalindromo(%q) = %v; esperado %v", c.texto, obtenido, c.esperado)
			}
		})
	}
}
