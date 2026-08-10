package solucion

import (
	"errors"
	"strings"
)

func RutaImport(modulo, paquete string) (string, error) {
	modulo = strings.Trim(modulo, "/ ")
	paquete = strings.Trim(paquete, "/ ")
	if modulo == "" || paquete == "" {
		return "", errors.New("módulo y paquete son obligatorios")
	}
	return modulo + "/" + paquete, nil
}
