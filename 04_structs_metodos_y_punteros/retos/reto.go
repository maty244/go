package reto

import "errors"

var ErrRetiroInvalido = errors.New("retiro inválido")

type Cuenta struct{ SaldoCentavos int64 }

func (c *Cuenta) Retirar(montoCentavos int64) error {
	// TODO: validar antes de mutar el valor apuntado por c.
	return nil
}
