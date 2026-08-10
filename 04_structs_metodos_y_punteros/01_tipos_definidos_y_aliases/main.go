package main

import "fmt"

// Un tipo DEFINIDO crea un tipo nuevo, aunque comparta representación con el
// tipo subyacente. Esto permite impedir mezclas accidentales y agregar métodos.
type Metros int64
type Segundos int64

func (m Metros) String() string { return fmt.Sprintf("%d m", m) }

// Un ALIAS es solamente otro nombre para el mismo tipo. No crea una frontera.
type Identificador = int64

// Desde Go 1.24 los aliases también pueden tener parámetros de tipo.
type Conjunto[T comparable] = map[T]struct{}

func main() {
	distancia := Metros(120)
	tiempo := Segundos(10)

	// distancia = tiempo // no compila: Metros y Segundos son tipos distintos
	velocidad := int64(distancia) / int64(tiempo)
	fmt.Println("Distancia:", distancia)
	fmt.Println("Velocidad:", velocidad, "m/s")

	var id Identificador = 42
	var numero int64 = id // son el mismo tipo: no necesita conversión
	fmt.Println("Alias:", numero)

	visitados := Conjunto[string]{"/inicio": {}}
	_, existe := visitados["/inicio"]
	fmt.Println("Conjunto genérico contiene /inicio:", existe)
}
