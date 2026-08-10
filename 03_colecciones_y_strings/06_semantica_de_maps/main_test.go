package main

import (
	"maps"
	"testing"
)

func TestMapAliasingCloneYClear(t *testing.T) {
	original := map[string]int{"mouse": 2}
	alias := original
	alias["mouse"] = 7
	if original["mouse"] != 7 {
		t.Fatal("la asignación de map debería compartir datos")
	}

	copia := maps.Clone(original)
	copia["mouse"] = 10
	if original["mouse"] == 10 {
		t.Fatal("maps.Clone mantuvo aliasing")
	}

	clear(copia)
	if len(copia) != 0 {
		t.Fatalf("clear dejó %d elementos", len(copia))
	}
	var nilMap map[string]int
	clear(nilMap) // es seguro; escribir nilMap["x"] no lo sería
}
