package solucion

import "strconv"

func FizzBuzz(limite int) []string {
	resultado := make([]string, 0, limite)
	for numero := 1; numero <= limite; numero++ {
		switch {
		case numero%15 == 0:
			resultado = append(resultado, "FizzBuzz")
		case numero%3 == 0:
			resultado = append(resultado, "Fizz")
		case numero%5 == 0:
			resultado = append(resultado, "Buzz")
		default:
			resultado = append(resultado, strconv.Itoa(numero))
		}
	}
	return resultado
}
