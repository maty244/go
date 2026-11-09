package main

import "fmt"

// Estas funciones separan tres operaciones que suelen confundirse:
//   n = ...  cambia un int local;
//   *p = ... cambia el int del destino;
//   p = ...  cambia a dónde apunta el puntero local.
// El tipo del parámetro NO cambia la regla: Go siempre copia el argumento.

func incrementarCopia(n int) {
	n = n + 1 // n pertenece a esta llamada, no es la variable del llamador
}

// Premisa: p debe ser no-nil. El llamador decide pasar &n, que sí tiene destino.
func incrementarDestino(p *int) {
	*p = *p + 1 // se copia la dirección al llamar, no el int al que apunta
}

// p y destino son DOS copias locales de punteros recibidos como argumentos.
// Reasignar p no cambia la variable puntero que existe en el llamador.
func cambiarPunteroLocal(p, destino *int) {
	p = destino // desde esta línea, solo el p LOCAL identifica el otro int
	*p = 99     // esta escritura sí afecta al int que recibimos como destino
}

// Retornar el nuevo puntero permite que el llamador decida si quiere guardarlo.
// El int creado aquí sigue siendo válido después del retorno: Go administra
// su vida mientras sea alcanzable. No hay que llamar a free ni delete.
func crearContador(inicial int) *int {
	contador := inicial
	return &contador
}

// **int se lee "puntero a una variable que guarda un *int".
// Premisa: pp no es nil. *pp puede ser nil: lo vamos a REEMPLAZAR, no a leer
// como un int. destino también puede ser nil si queremos dejar *pp sin destino.
func reemplazarPuntero(pp **int, destino *int) {
	*pp = destino // accede a la variable puntero del llamador y cambia su contenido
}

// intercambiar documenta otra política: acepta nil y devuelve false sin mutar.
// Validamos ambas entradas ANTES de escribir; nunca hay un intercambio parcial.
func intercambiar(a, b *int) bool {
	if a == nil || b == nil {
		return false
	}
	*a, *b = *b, *a // ambos valores de la derecha se evalúan antes de asignar
	return true
}

func main() {
	fmt.Println("=== 1. Una copia de int y una copia de *int ===")
	n := 10
	incrementarCopia(n)
	fmt.Println("tras incrementarCopia(n):", n) // 10
	incrementarDestino(&n)
	fmt.Println("tras incrementarDestino(&n):", n) // 11

	fmt.Println("\n=== 2. Cambiar p local no cambia el puntero del llamador ===")
	a, b := 10, 20
	p := &a
	cambiarPunteroLocal(p, &b)
	fmt.Printf("a=%d b=%d *p=%d p==&a=%v\n", a, b, *p, p == &a) // 10, 99, 10, true
	// Dibujo ANTES: p ──→ a=10; b=20
	// DENTRO:      p local ──→ b=99; el p del llamador sigue ──→ a=10
	// DESPUÉS:     desaparece el parámetro local; a=10, b=99, p ──→ a

	fmt.Println("\n=== 3. Retornar un puntero y asignarlo explícitamente ===")
	p = crearContador(50) // ahora SÍ escribimos la variable p del llamador
	*p = 51
	fmt.Printf("a=%d *p=%d\n", a, *p) // 10, 51: p ya no identifica a

	fmt.Println("\n=== 4. Otra opción: pasar la dirección de p ===")
	reemplazarPuntero(&p, &b)
	fmt.Printf("*p=%d p==&b=%v\n", *p, p == &b) // 99, true
	// &p es **int: una forma de darle a la función acceso a la propia p.
	// No necesitás **int solo para cambiar un int: ahí basta *int.
	reemplazarPuntero(&p, nil)
	fmt.Println("p quedó nil:", p == nil)                       // true; no debemos evaluar *p ahora
	fmt.Printf("los ints siguen existiendo: a=%d b=%d\n", a, b) // 10, 99

	fmt.Println("\n=== 5. Un ejercicio pequeño: intercambiar dos ints ===")
	x, y := 3, 8
	fmt.Println("intercambio válido:", intercambiar(&x, &y))                // true
	fmt.Printf("x=%d y=%d\n", x, y)                                         // 8, 3
	fmt.Println("intercambio con nil:", intercambiar(&x, nil))              // false
	fmt.Printf("sin cambios: x=%d y=%d\n", x, y)                            // 8, 3
	fmt.Println("Mismo destino en ambos parámetros:", intercambiar(&x, &x)) // true; x sigue 8
	// Tener dos punteros no implica tener dos destinos diferentes. Antes de
	// diseñar una operación, preguntate si los argumentos podrían ser aliases.
}
