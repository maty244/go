package solucion

import "errors"

var ErrRetiroInvalido = errors.New("retiro inválido")

type Cuenta struct{ SaldoCentavos int64 }

func (c *Cuenta) Retirar(montoCentavos int64) error {
	if montoCentavos < 0 || montoCentavos > c.SaldoCentavos {
		return ErrRetiroInvalido
	}
	c.SaldoCentavos -= montoCentavos
	return nil
}
