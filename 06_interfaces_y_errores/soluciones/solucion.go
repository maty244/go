package solucion

import (
	"errors"
	"fmt"
)

var ErrNoEncontrado = errors.New("no encontrado")

func Cargar(id int64) error {
	return fmt.Errorf("cargar producto %d: %w", id, ErrNoEncontrado)
}
