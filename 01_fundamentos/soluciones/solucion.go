package solucion

import "errors"

type Dinero int64

var ErrValorNegativo = errors.New("valor negativo")

func CalcularTotal(precio Dinero, cantidad int) (Dinero, error) {
	if precio < 0 || cantidad < 0 {
		return 0, ErrValorNegativo
	}
	return precio * Dinero(cantidad), nil
}
