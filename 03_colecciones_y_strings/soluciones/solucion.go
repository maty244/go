package solucion

// Clonar copia los elementos a otro array. Como string es inmutable, cambiar
// un elemento del clon no modifica la entrada.
// Nota: en la biblioteca estándar moderna (Go 1.21+), slices.Clone(entrada)
// hace exactamente esto mismo (la versión genérica la vemos en 08).
func Clonar(entrada []string) []string {
	if entrada == nil {
		return nil
	}
	return append([]string{}, entrada...)
}
