package solucion

import "unicode"

func EsPalindromo(texto string) bool {
	var letras []rune
	for _, r := range texto {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			letras = append(letras, unicode.ToLower(r))
		}
	}
	for izquierda, derecha := 0, len(letras)-1; izquierda < derecha; izquierda, derecha = izquierda+1, derecha-1 {
		if letras[izquierda] != letras[derecha] {
			return false
		}
	}
	return true
}
