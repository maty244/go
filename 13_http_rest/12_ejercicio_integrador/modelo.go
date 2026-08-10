package main

import (
	"errors"
	"strings"
)

// Dinero siempre representa la unidad mínima: centavos. No usamos float64
// para importes de negocio porque sus redondeos binarios se acumulan.
type Dinero int64

type Producto struct {
	ID             int64  `json:"id"`
	Nombre         string `json:"nombre"`
	PrecioCentavos Dinero `json:"precio_centavos"`
	Stock          int    `json:"stock"`
}

type CrearProducto struct {
	Nombre         string `json:"nombre"`
	PrecioCentavos Dinero `json:"precio_centavos"`
	Stock          int    `json:"stock"`
}

type ActualizarProducto struct {
	Nombre         string `json:"nombre"`
	PrecioCentavos Dinero `json:"precio_centavos"`
	Stock          int    `json:"stock"`
}

var ErrDatosInvalidos = errors.New("datos de producto inválidos")

func validarProducto(nombre string, precio Dinero, stock int) error {
	if strings.TrimSpace(nombre) == "" || precio <= 0 || stock < 0 {
		return ErrDatosInvalidos
	}
	return nil
}
