package reto

import "errors"

type Dinero int64

var ErrValorNegativo = errors.New("valor negativo")

func CalcularTotal(precio Dinero, cantidad int) (Dinero, error) {
	// TODO: validar y multiplicar.
	return 0, nil
}
