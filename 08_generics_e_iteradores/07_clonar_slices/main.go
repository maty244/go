package main

import "fmt"

// En la etapa 03 clonaste []string. Ahora T permite reutilizar el mismo
// algoritmo con otros elementos. any no exige operaciones sobre T: solo copiar.
// Esto crea un array distinto, pero cada elemento se copia superficialmente.
func Clonar[T any](entrada []T) []T {
	if entrada == nil {
		return nil // preservar ausencia, distinta de un slice vacío inicializado
	}
	copia := make([]T, len(entrada))
	copy(copia, entrada)
	return copia
}

func main() {
	numeros := []int{1, 2}
	copia := Clonar(numeros) // el compilador infiere T = int desde el argumento
	copia[0] = 99
	fmt.Println("original:", numeros, "clon:", copia)          // [1 2], [99 2]
	fmt.Println("nil conservado:", Clonar[int](nil) == nil)    // T explícito: nil no dice qué elemento hay
	fmt.Println("vacío sigue no-nil:", Clonar([]int{}) != nil) // true

	// Clonar [][]int separa el array EXTERIOR, no los arrays interiores.
	anidado := [][]int{{1}}
	copiaAnidada := Clonar(anidado)
	copiaAnidada[0][0] = 7
	fmt.Println("interior compartido:", anidado[0][0]) // 7
	fmt.Println("Generics reutiliza el algoritmo; no agrega una copia profunda automática.")
}
