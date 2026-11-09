package main

import (
	"fmt"
	"maps"
	"slices"
)

// Ficha combina valores directos con valores que apuntan a otros datos.
// Copiar el struct copia TODOS los campos, pero no clona sus destinos.
type Ficha struct {
	Nombre    string
	Etiquetas []string
	Puntos    map[string]int
	Nivel     *int
}

// clonarFicha separa los datos mutables de ESTE tipo, no de cualquier struct.
// Si agregáramos slices dentro del map, también habría que clonar esos slices.
// "Copia profunda" depende de qué contiene el dato y qué queremos compartir.
func clonarFicha(original Ficha) Ficha {
	copia := original                                  // Nombre es string inmutable; basta copiar su valor
	copia.Etiquetas = slices.Clone(original.Etiquetas) // nuevo array subyacente (preserva nil vs vacío)
	copia.Puntos = maps.Clone(original.Puntos)         // nuevo map; sus valores son ints
	if original.Nivel != nil {
		copia.Nivel = new(*original.Nivel) // Go 1.26: int nuevo con el mismo valor
	}
	return copia
}

func main() {
	original := Ficha{
		Nombre: "Mati", Etiquetas: []string{"Go"},
		Puntos: map[string]int{"Go": 1}, Nivel: new(1),
	}
	superficial := original
	independiente := clonarFicha(original) // clonar ANTES de las mutaciones

	superficial.Nombre = "otra ficha"     // reemplaza solo el campo de la copia
	superficial.Etiquetas[0] = "cambiada" // modifica el array compartido
	superficial.Puntos["Go"] = 99         // modifica la tabla compartida
	*superficial.Nivel = 2                // modifica el int al que apuntan ambas fichas

	// Salida esperada: Mati [cambiada] 99 2; Mati [Go] 1 1.
	fmt.Println("original:", original.Nombre, original.Etiquetas, original.Puntos["Go"], *original.Nivel)
	fmt.Println("clon:", independiente.Nombre, independiente.Etiquetas, independiente.Puntos["Go"], *independiente.Nivel)
	fmt.Println("Compartir datos no cambia la regla: Go sigue pasando valores por valor.")

	// Experimento: sustituí maps.Clone por asignación. Predecí qué resultado
	// cambia. Después agregá un []int como valor del map y pensá qué clonarías.
}
