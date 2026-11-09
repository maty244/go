package main

import "fmt"

// Con Go 1.22+ (este módulo usa 1.26), := crea una variable n por vuelta.
// Cada n recibe una COPIA del int del slice. &n apunta a esa copia, no al array.
func direccionesDeCopias(numeros []int) []*int {
	var salida []*int
	for _, n := range numeros {
		salida = append(salida, &n)
	}
	return salida
}

// El índice permite tomar la dirección de la variable que ocupa esa posición
// del array subyacente. No se crea una copia del int para tomar su dirección.
func direccionesDeElementos(numeros []int) []*int {
	var salida []*int
	for i := range numeros {
		salida = append(salida, &numeros[i])
	}
	return salida
}

// Con = no declaramos n dentro del range: reasignamos la MISMA n exterior.
func direccionesDeVariableReutilizada(numeros []int) []*int {
	var salida []*int
	var n int
	for _, n = range numeros {
		salida = append(salida, &n)
	}
	return salida
}

func main() {
	fmt.Println("=== 1. &n y &numeros[i] no son lo mismo ===")
	numeros := []int{10, 20, 30}
	copias := direccionesDeCopias(numeros)
	*copias[0] = 99
	fmt.Println("cambia la copia; slice:", numeros)                           // [10 20 30]
	fmt.Println("valores de las copias:", *copias[0], *copias[1], *copias[2]) // 99, 20, 30
	elementos := direccionesDeElementos(numeros)
	*elementos[0] = 77
	fmt.Println("cambia el elemento; slice:", numeros) // [77 20 30]

	fmt.Println("\n=== 2. Una variable exterior es un único destino ===")
	reutilizados := direccionesDeVariableReutilizada([]int{10, 20, 30})
	fmt.Println("valor final en todos:", *reutilizados[0], *reutilizados[1], *reutilizados[2]) // 30 30 30
	fmt.Println("mismo destino:", reutilizados[0] == reutilizados[1])                          // true
	*reutilizados[0] = 88
	fmt.Println("cambian todas las lecturas:", *reutilizados[1], *reutilizados[2]) // 88 88

	fmt.Println("\n=== 3. append puede cambiar el array de un slice ===")
	slice := make([]int, 1, 1) // len=1 y cap=1: no cabe un elemento adicional
	slice[0] = 10
	anterior := &slice[0]
	slice = append(slice, 20)                                                      // DEBE usar otro array por falta de capacidad
	*anterior = 99                                                                 // el viejo elemento sigue vivo porque conservamos su puntero
	fmt.Println("slice usa el array nuevo:", slice)                                // [10 20]
	fmt.Println("puntero al array anterior:", *anterior)                           // 99
	fmt.Println("¿Sigue siendo el primer elemento actual?", anterior == &slice[0]) // false
	// Un puntero no sigue automáticamente al slice ni a un índice lógico.
	// Si append cabe y reutiliza el array, el resultado es distinto. Probalo
	// con make([]int, 1, 2). No confundas una reubicación con un puntero inválido.

	fmt.Println("\n=== 4. Los elementos de un map no son direccionables ===")
	valores := map[string]int{"a": 10}
	// p := &valores["a"] // no compila: no se puede tomar esa dirección
	copia := valores["a"]
	p := &copia
	*p = 99
	fmt.Println("copiar un int del map no cambia la entrada:", valores["a"]) // 10
	valores["a"] = copia                                                     // escribir la entrada sí reemplaza su valor
	fmt.Println("después de asignar al map:", valores["a"])                  // 99

	// Alternativa: guardar punteros como VALORES del map. El map nos da una
	// copia del *int, suficiente para acceder al mismo int que había guardado.
	punteros := map[string]*int{"a": new(10)}
	*punteros["a"] = 50 // premisa: la clave existe y su puntero no es nil
	guardado := punteros["a"]
	delete(punteros, "a")                                          // quita una entrada, no invalida otros punteros al int
	fmt.Println("sin entrada, el int sigue accesible:", *guardado) // 50
	// Si la clave no existe, punteros[clave] devuelve nil: hay que validarlo
	// antes de desreferenciar. Un map de punteros no elimina la ausencia.
}
