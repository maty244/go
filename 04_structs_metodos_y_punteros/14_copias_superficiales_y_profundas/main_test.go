package main

import "testing"

func TestClonSeparaLosDatosMutables(t *testing.T) {
	original := Ficha{Nombre: "Mati", Etiquetas: []string{"Go"}, Puntos: map[string]int{"Go": 1}, Nivel: new(1)}
	copia := clonarFicha(original)
	copia.Nombre = "otra"
	copia.Etiquetas[0] = "otra"
	copia.Puntos["Go"] = 99
	*copia.Nivel = 99
	if original.Nombre != "Mati" || original.Etiquetas[0] != "Go" || original.Puntos["Go"] != 1 || *original.Nivel != 1 {
		t.Fatalf("el clon modificó el original: %+v", original)
	}
}

func TestClonPreservaAusenciaYVacio(t *testing.T) {
	if copia := clonarFicha(Ficha{}); copia.Etiquetas != nil || copia.Puntos != nil || copia.Nivel != nil {
		t.Fatal("la ausencia de datos dejó de ser nil")
	}
	copia := clonarFicha(Ficha{Etiquetas: []string{}, Puntos: map[string]int{}})
	if copia.Etiquetas == nil || copia.Puntos == nil || len(copia.Etiquetas) != 0 || len(copia.Puntos) != 0 {
		t.Fatal("las colecciones vacías no-nil no se preservaron")
	}
}
