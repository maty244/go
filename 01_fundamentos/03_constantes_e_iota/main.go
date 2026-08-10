package main

import "fmt"

// =========================================================
// CONSTANTES EN GO
// =========================================================
// Una constante es como una variable PERO su valor no puede
// cambiar nunca durante la ejecución del programa.
// Se usa para valores que son fijos por naturaleza:
//   - El número Pi (3.14159...)
//   - La velocidad de la luz
//   - El máximo de intentos de login
//   - Versión de la app
//
// En Go las constantes se declaran con la palabra "const".
// A diferencia de las variables, las constantes NO pueden usar :=

// ─────────────────────────────────────────────────────────
// CONSTANTES A NIVEL DE PAQUETE (fuera de funciones)
// ─────────────────────────────────────────────────────────
// Estas constantes son accesibles desde cualquier función del paquete.
const PI = 3.14159265358979
const VELOCIDAD_LUZ = 299_792_458 // metros por segundo
const VERSION_APP = "1.0.0"
const MAX_INTENTOS = 3
const APP_NOMBRE = "MiApp"

// Constantes en bloque: forma más limpia de declarar varias
const (
	LUNES     = "Lunes"
	MARTES    = "Martes"
	MIERCOLES = "Miércoles"
	JUEVES    = "Jueves"
	VIERNES   = "Viernes"
)

// ─────────────────────────────────────────────────────────
// iota: el generador automático de constantes
// ─────────────────────────────────────────────────────────
// iota es una herramienta especial de Go que genera números
// automáticamente dentro de un bloque const.
// Empieza en 0 y se incrementa en 1 por cada constante.
// Es IDEAL para crear enumeraciones (como en otros lenguajes).
const (
	DOMINGO    = iota // 0
	LUNES2            // 1
	MARTES2           // 2
	MIERCOLES2        // 3
	JUEVES2           // 4
	VIERNES2          // 5
	SABADO            // 6
)

// iota con operaciones matemáticas
// Podemos usarlo con fórmulas para generar secuencias útiles.
const (
	_  = iota             // ignoramos el 0 con blank identifier
	KB = 1 << (10 * iota) // 1 << 10 = 1024 bytes
	MB                    // 1 << 20 = 1048576 bytes
	GB                    // 1 << 30 = 1073741824 bytes
	TB                    // 1 << 40 = 1099511627776 bytes
)

// Ejemplo práctico: estados de un pedido
const (
	PEDIDO_PENDIENTE  = iota // 0
	PEDIDO_CONFIRMADO        // 1
	PEDIDO_EN_CAMINO         // 2
	PEDIDO_ENTREGADO         // 3
	PEDIDO_CANCELADO         // 4
)

func main() {
	fmt.Println("=== Constantes simples ===")
	fmt.Println("PI:", PI)
	fmt.Println("Velocidad de la luz:", VELOCIDAD_LUZ, "m/s")
	fmt.Println("Versión:", VERSION_APP)
	fmt.Println("App:", APP_NOMBRE)
	fmt.Println("Máx intentos:", MAX_INTENTOS)

	// ─────────────────────────────────────────────────────────
	// CONSTANTES LOCALES (dentro de funciones)
	// ─────────────────────────────────────────────────────────
	// También podemos declarar constantes dentro de funciones.
	// Solo existen en el scope de esa función.
	const IMPUESTO = 0.21 // 21% de IVA
	const DESCUENTO = 0.10

	precio := 100.0
	precioConIVA := precio * (1 + IMPUESTO)
	precioFinal := precioConIVA * (1 - DESCUENTO)

	fmt.Println("\n=== Constantes locales (cálculo de precio) ===")
	fmt.Printf("Precio base:    $%.2f\n", precio)
	fmt.Printf("Con IVA (21%%): $%.2f\n", precioConIVA)
	fmt.Printf("Con descuento:  $%.2f\n", precioFinal)

	// ─────────────────────────────────────────────────────────
	// DÍAS DE LA SEMANA con iota
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Días con iota ===")
	fmt.Println("Domingo =", DOMINGO)
	fmt.Println("Lunes =", LUNES2)
	fmt.Println("Sábado =", SABADO)

	// ─────────────────────────────────────────────────────────
	// TAMAÑOS DE MEMORIA con iota y bit shifting
	// ─────────────────────────────────────────────────────────
	// << es el operador de desplazamiento de bits (shift left)
	// 1 << 10 significa: toma el 1 en binario y desplázalo 10 lugares
	// eso es igual a 2^10 = 1024
	fmt.Println("\n=== Tamaños de memoria ===")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)

	// Aplicación real: mostrar tamaño de archivo en la unidad correcta
	archivoBytes := 2_500_000 // 2.5 MB en bytes
	fmt.Printf("\nArchivo de %d bytes = %.2f MB\n", archivoBytes, float64(archivoBytes)/float64(MB))

	// ─────────────────────────────────────────────────────────
	// ESTADOS DE PEDIDO con iota
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Estados de pedido ===")
	estadoActual := PEDIDO_EN_CAMINO
	fmt.Println("Estado del pedido #1234:", estadoActual) // imprime 2

	// Función para describir el estado
	fmt.Println("Descripción:", describirEstado(estadoActual))

	// ─────────────────────────────────────────────────────────
	// DIFERENCIA CLAVE: constante vs variable
	// ─────────────────────────────────────────────────────────
	// Esta línea daría ERROR de compilación si se descomenta:
	// PI = 3.14 // ERROR: cannot assign to PI (untyped float constant)

	// Las constantes son evaluadas en TIEMPO DE COMPILACIÓN,
	// no en tiempo de ejecución. Esto las hace más eficientes.
	fmt.Println("\n=== Constantes vs Variables ===")
	fmt.Println("Las constantes no pueden cambiar su valor.")
	fmt.Println("Son ideales para configuración, límites y enumeraciones.")
}

// Función que usa los estados definidos con iota
func describirEstado(estado int) string {
	switch estado {
	case PEDIDO_PENDIENTE:
		return "Esperando confirmación"
	case PEDIDO_CONFIRMADO:
		return "Pedido confirmado"
	case PEDIDO_EN_CAMINO:
		return "En camino hacia vos"
	case PEDIDO_ENTREGADO:
		return "Entregado con éxito"
	case PEDIDO_CANCELADO:
		return "Cancelado"
	default:
		return "Estado desconocido"
	}
}
