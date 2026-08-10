package reto

import "iter"

func Filtrar[T any](entrada iter.Seq[T], aceptar func(T) bool) iter.Seq[T] {
	// TODO: devolvé una función iteradora lazy.
	return func(func(T) bool) {}
}
