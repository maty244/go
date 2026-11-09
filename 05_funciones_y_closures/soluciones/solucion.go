package solucion

import "slices"

func Aplicar(entrada []int, transformaciones ...func(int) int) []int {
	// En Go 1.21+, slices.Clone crea una copia superficial del slice sin mutar el original.
	resultado := slices.Clone(entrada)
	for _, transformar := range transformaciones {
		for i, valor := range resultado {
			resultado[i] = transformar(valor)
		}
	}
	return resultado
}
