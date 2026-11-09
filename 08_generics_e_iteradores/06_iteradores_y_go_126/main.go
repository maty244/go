package main

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"slices"
)

// Los aliases genéricos están disponibles de forma estable desde Go 1.24.
// Lista[T] y []T son exactamente el mismo tipo, no dos tipos definidos.
type Lista[T any] = []T

func Maximo[T cmp.Ordered](a, b T) T {
	return max(a, b)
}

// Filtrar devuelve un iterador lazy (iter.Seq[T], definido en Go 1.23+ como
// `type Seq[V any] func(yield func(V) bool)`). También existe `iter.Seq2[K, V]`
// para iterar de a pares clave/valor (como `maps.All` o `slices.All`).
// No crea el slice de salida hasta que el consumidor lo solicita con slices.Collect.
func Filtrar[T any](entrada iter.Seq[T], aceptar func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for valor := range entrada {
			if aceptar(valor) && !yield(valor) {
				return // el consumidor hizo break: hay que respetarlo
			}
		}
	}
}

// Go 1.26 permite que una restricción se refiera a la instanciación del mismo
// tipo genérico (polimorfismo F-acotado). Es potente cuando un tipo define su
// propio método de comparación consigo mismo.
type OrdenableConsigo[T OrdenableConsigo[T]] interface {
	Menor(T) bool
}

type Monto int

func (m Monto) Menor(otro Monto) bool { return m < otro }

func MenorDeDos[T OrdenableConsigo[T]](a, b T) T {
	if a.Menor(b) {
		return a
	}
	return b
}

func main() {
	// Range sobre enteros existe desde Go 1.22.
	for i := range 3 {
		fmt.Println("range entero:", i)
	}

	numeros := Lista[int]{1, 2, 3, 4, 5, 6}
	pares := Filtrar(slices.Values(numeros), func(n int) bool { return n%2 == 0 })
	fmt.Println("Pares:", slices.Collect(pares))

	stock := map[string]int{"mouse": 3, "teclado": 2}
	claves := slices.Sorted(maps.Keys(stock))
	fmt.Println("Claves ordenadas:", claves)
	fmt.Println("Máximo con cmp.Ordered:", Maximo(7, 11))
	fmt.Println("Menor con restricción autorreferencial (Go 1.26):", MenorDeDos(Monto(150), Monto(90)))

	fmt.Println("\nLa stdlib no trae slices.Map/Filter/Reduce.")
	fmt.Println("Un loop suele ser la solución más clara; iter sirve cuando compone mejor.")
}
