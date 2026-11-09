package main

import "testing"

func TestBusquedaAdmiteElementosNilYDevuelveElObjetoReal(t *testing.T) {
	producto := &Producto{ID: 2, Nombre: "Mouse"}
	catalogo := []*Producto{nil, producto, nil}
	encontrado := buscarProducto(catalogo, 2)
	if encontrado != producto {
		t.Fatal("buscar debe devolver el objeto original, no una copia")
	}
	encontrado.Nombre = "Teclado"
	if producto.Nombre != "Teclado" {
		t.Fatal("la escritura mediante el puntero debe modificar el producto real")
	}
	if buscarProducto(catalogo, 99) != nil || buscarProducto(nil, 2) != nil {
		t.Fatal("una búsqueda sin resultado debe devolver nil")
	}
}
