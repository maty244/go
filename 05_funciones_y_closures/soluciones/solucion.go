package solucion

func Aplicar(entrada []int, transformaciones ...func(int) int) []int {
	resultado := append([]int(nil), entrada...)
	for _, transformar := range transformaciones {
		for i, valor := range resultado {
			resultado[i] = transformar(valor)
		}
	}
	return resultado
}
