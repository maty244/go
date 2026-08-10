// Package precios agrupa la lógica de cálculo de precios del kiosco.
// Es un paquete de librería: no tiene func main(), otros paquetes lo importan.
package precios

// AplicarDescuento calcula el precio final tras aplicar un
// porcentaje de descuento (0.10 = 10%).
func AplicarDescuento(precio, porcentaje float64) float64 {
	return precio * (1 - porcentaje)
}

// AplicarIVA suma el IVA (21% en Argentina) a un precio.
func AplicarIVA(precio float64) float64 {
	return precio * 1.21
}
