package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// json.Unmarshal: DE JSON A STRUCT
// =========================================================
// Es el camino inverso a Marshal (tema 01): tomás texto JSON
// (por ejemplo, el body de un request HTTP) y lo convertís en un
// valor de Go con el que podés trabajar normalmente.
//
//   json.Unmarshal(datos []byte, destino *T) error
//
// OJO con la firma: recibe un PUNTERO al destino (por eso el &),
// porque necesita MODIFICAR la variable que le pasás para llenarla
// con los datos. Esto debería sonarte conocido de 04_structs_metodos_y_punteros/10_punteros_a_structs
// (punteros a structs) y 06_interfaces_y_errores/12_patron_de_errores (patrón valor, error).

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func main() {
	fmt.Println("=== json.Unmarshal: JSON → struct ===")

	// Este texto podría venir de un archivo, de una request HTTP,
	// o (como acá) de un literal para practicar.
	textoJSON := `{"Nombre":"Mouse","Precio":8500,"Stock":20}`

	var p Producto
	err := json.Unmarshal([]byte(textoJSON), &p) // ← puntero: llena "p"
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Producto: %+v\n", p)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Println("Precio:", p.Precio)

	// ─────────────────────────────────────────────────────────
	// SI EL JSON TIENE UN ERROR DE SINTAXIS, Unmarshal LO REPORTA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== JSON inválido: el error avisa ===")

	jsonRoto := `{"Nombre":"Mouse", "Precio": }` // falta el valor de Precio
	var p2 Producto
	err = json.Unmarshal([]byte(jsonRoto), &p2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────────────────
	// UNMARSHAL DE UN ARRAY JSON A []T
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Unmarshal de un array JSON ===")

	textoCatalogo := `[
		{"Nombre":"Alfajor","Precio":800,"Stock":15},
		{"Nombre":"Gaseosa","Precio":1200,"Stock":5}
	]`

	var catalogo []Producto
	if err := json.Unmarshal([]byte(textoCatalogo), &catalogo); err != nil {
		fmt.Println("Error:", err)
	}
	for _, prod := range catalogo {
		fmt.Printf("  %s: $%.2f (stock: %d)\n", prod.Nombre, prod.Precio, prod.Stock)
	}

	// ─────────────────────────────────────────────────────────
	// CAMPOS DEL JSON QUE NO EXISTEN EN EL STRUCT: SE IGNORAN
	// ─────────────────────────────────────────────────────────
	// Unmarshal es tolerante: si el JSON trae MÁS campos de los
	// que tiene tu struct, simplemente los descarta (a menos que
	// uses un decoder estricto con json.NewDecoder(r).DisallowUnknownFields()).

	fmt.Println("\n=== Campos extra en el JSON se ignoran ===")

	jsonConExtra := `{"Nombre":"Teclado","Precio":15000,"Stock":10,"Marca":"Logitech"}`
	var p3 Producto
	json.Unmarshal([]byte(jsonConExtra), &p3)
	fmt.Printf("  %+v (\"Marca\" no está en el struct, se ignoró)\n", p3)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  json.Unmarshal(datos, &destino)  → JSON → struct de Go")
	fmt.Println("  SIEMPRE con &                     → necesita modificar el destino")
	fmt.Println("  Devuelve error                    → si el JSON es inválido")
	fmt.Println("  Campos JSON de más                → se ignoran silenciosamente")
	fmt.Println("  []T destino                       → sirve para arrays JSON")
}
