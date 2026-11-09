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
//
// CONVENCIÓN DE NOMBRES EN GO:
// A diferencia de otros lenguajes (C, Java, Python, JS) que usan MAYUSCULAS_CON_GUION,
// en Go las constantes usan MixedCaps (PascalCase o camelCase), porque en Go la
// primera letra define la visibilidad:
//   - Mayúscula inicial (Pi, VelocidadLuz) → exportada (visible desde otros paquetes)
//   - Minúscula inicial (maxIntentos)      → privada (solo visible en este paquete)
const Pi = 3.14159265358979
const VelocidadLuz = 299_792_458 // metros por segundo (el _ es separador visual)
const VersionApp = "1.0.0"
const maxIntentos = 3
const appNombre = "MiApp"

// Constantes en bloque: forma más limpia de declarar varias juntas
const (
	Lunes     = "Lunes"
	Martes    = "Martes"
	Miercoles = "Miércoles"
	Jueves    = "Jueves"
	Viernes   = "Viernes"
)

// ─────────────────────────────────────────────────────────
// iota: el generador automático de constantes
// ─────────────────────────────────────────────────────────
// iota es un identificador predeclarado de Go que representa números
// enteros sucesivos sin tipo (0, 1, 2...) dentro de un bloque const.
// Reglas clave de un bloque const con iota:
//  1. iota se reinicia a 0 cada vez que aparece la palabra clave "const".
//  2. Incrementa en 1 por cada línea (ConstSpec) dentro del bloque.
//  3. Si omitís el "= expresión" en las líneas siguientes, Go REPITE
//     automáticamente la expresión de la línea anterior con el nuevo iota.
const (
	Domingo    = iota // iota = 0
	DiaLunes          // repite "= iota" → 1
	DiaMartes         // repite "= iota" → 2
	DiaMierc          // repite "= iota" → 3
	DiaJueves         // repite "= iota" → 4
	DiaViernes        // repite "= iota" → 5
	Sabado            // repite "= iota" → 6
)

// iota con operaciones matemáticas (desplazamiento de bits)
// Como Go repite la fórmula "1 << (10 * iota)" en cada línea sin "=",
// cada constante multiplica el desplazamiento por el nuevo valor de iota:
const (
	_  = iota             // iota = 0 → ignoramos el 0 con blank identifier _
	KB = 1 << (10 * iota) // iota = 1 → 1 << 10 = 1024 bytes
	MB                    // iota = 2 → 1 << 20 = 1_048_576 bytes
	GB                    // iota = 3 → 1 << 30 = 1_073_741_824 bytes
	TB                    // iota = 4 → 1 << 40 = 1_099_511_627_776 bytes
)

// Ejemplo práctico: estados de un pedido
const (
	PedidoPendiente  = iota // 0
	PedidoConfirmado        // 1
	PedidoEnCamino          // 2
	PedidoEntregado         // 3
	PedidoCancelado         // 4
)

func main() {
	fmt.Println("=== Constantes simples ===")
	fmt.Println("Pi:", Pi)
	fmt.Println("Velocidad de la luz:", VelocidadLuz, "m/s")
	fmt.Println("Versión:", VersionApp)
	fmt.Println("App:", appNombre)
	fmt.Println("Máx intentos:", maxIntentos)

	// ─────────────────────────────────────────────────────────
	// CONSTANTES LOCALES (dentro de funciones)
	// ─────────────────────────────────────────────────────────
	// También podemos declarar constantes dentro de funciones.
	// Solo existen en el scope de esa función.
	const impuestoIVA = 0.21 // 21% de IVA
	const descuento = 0.10

	precio := 100.0
	precioConIVA := precio * (1 + impuestoIVA)
	precioFinal := precioConIVA * (1 - descuento)

	fmt.Println("\n=== Constantes locales (cálculo de precio) ===")
	fmt.Printf("Precio base:    $%.2f\n", precio)
	fmt.Printf("Con IVA (21%%): $%.2f\n", precioConIVA)
	fmt.Printf("Con descuento:  $%.2f\n", precioFinal)

	// ─────────────────────────────────────────────────────────
	// DÍAS DE LA SEMANA con iota
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Días con iota ===")
	fmt.Println("Domingo =", Domingo)
	fmt.Println("Lunes =", DiaLunes)
	fmt.Println("Sábado =", Sabado)

	// ─────────────────────────────────────────────────────────
	// TAMAÑOS DE MEMORIA con iota y bit shifting
	// ─────────────────────────────────────────────────────────
	// << es el operador de desplazamiento de bits a la izquierda (shift left)
	// 1 << 10 significa: toma el 1 en binario y desplázalo 10 lugares a la izquierda.
	// En binario, desplazar N lugares a la izquierda equivale a multiplicar por 2^N:
	// 1 << 10 = 2^10 = 1024
	fmt.Println("\n=== Tamaños de memoria ===")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)

	// Aplicación real: mostrar tamaño de archivo en la unidad correcta
	archivoBytes := 2_500_000 // ~2.38 MB en bytes
	fmt.Printf("\nArchivo de %d bytes = %.2f MB\n", archivoBytes, float64(archivoBytes)/float64(MB))

	// ─────────────────────────────────────────────────────────
	// ESTADOS DE PEDIDO con iota
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Estados de pedido ===")
	estadoActual := PedidoEnCamino
	fmt.Println("Estado del pedido #1234:", estadoActual) // imprime 2

	// Función para describir el estado
	fmt.Println("Descripción:", describirEstado(estadoActual))

	// ─────────────────────────────────────────────────────────
	// DIFERENCIA CLAVE: constante vs variable
	// ─────────────────────────────────────────────────────────
	// Esta línea daría ERROR de compilación si se descomenta:
	// Pi = 3.14 // ERROR: cannot assign to Pi (neither addressable nor a map index expression)

	// Las constantes se evalúan en TIEMPO DE COMPILACIÓN (compile time),
	// no en tiempo de ejecución (runtime). Por eso no podés inicializar
	// una constante con el resultado de una función común:
	// const azar = rand.Int() // ERROR: rand.Int() is not constant
	fmt.Println("\n=== Constantes vs Variables ===")
	fmt.Println("Las constantes se conocen al compilar y jamás cambian su valor.")
	fmt.Println("Son ideales para configuración fija, límites y enumeraciones.")
}

// Función que usa los estados definidos con iota
func describirEstado(estado int) string {
	switch estado {
	case PedidoPendiente:
		return "Esperando confirmación"
	case PedidoConfirmado:
		return "Pedido confirmado"
	case PedidoEnCamino:
		return "En camino hacia vos"
	case PedidoEntregado:
		return "Entregado con éxito"
	case PedidoCancelado:
		return "Cancelado"
	default:
		return "Estado desconocido"
	}
}
