package solucion

func Clonar[T any](entrada []T) []T {
	if entrada == nil {
		return nil
	}
	return append([]T{}, entrada...)
}
