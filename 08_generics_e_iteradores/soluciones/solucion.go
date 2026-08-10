package solucion

import "iter"

func Filtrar[T any](entrada iter.Seq[T], aceptar func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for valor := range entrada {
			if aceptar(valor) && !yield(valor) {
				return
			}
		}
	}
}
