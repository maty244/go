package main

import "fmt"

// Este test ejecutable protege el caso que sigue existiendo en Go moderno:
// las funciones capturan una variable exterior, no una instantánea de su valor.
func Example_capturarExterior() {
	for _, f := range capturarExterior() {
		fmt.Println(f())
	}
	// Output:
	// 3
	// 3
	// 3
}
