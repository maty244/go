package reto

import "errors"

var ErrNoEncontrado = errors.New("no encontrado")

func Cargar(id int64) error {
	// TODO: devolvé un error con el ID y %w.
	return ErrNoEncontrado
}
