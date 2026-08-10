package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// CAMPOS OPCIONALES CON PUNTEROS (conecta con 04_structs_metodos_y_punteros/09_valores_opcionales_y_new)
// =========================================================
// Este es uno de los usos MÁS importantes de los punteros en
// programación real, y la razón por la que 04_structs_metodos_y_punteros/09_valores_opcionales_y_new (zero
// value vs "no hay valor") importa tanto.
//
// Problema: en un PATCH (actualización parcial), el cliente puede
// mandar SOLO los campos que quiere cambiar. Si usás int para el
// stock, no podés distinguir:
//   a) "no mandaron el campo stock" (no tocar el stock)
//   b) "mandaron stock: 0" (poner el stock en 0 a propósito)
// Con int, ambos casos dan 0. La solución: usar *int.
//   nil     → el campo NO vino en el JSON, no tocar nada
//   *0      → el campo SÍ vino, y su valor es 0

type ActualizarProducto struct {
	Nombre         *string `json:"nombre,omitempty"`
	PrecioCentavos *int64  `json:"precio_centavos,omitempty"`
	Stock          *int    `json:"stock,omitempty"`
}

func main() {
	fmt.Println("=== El problema: distinguir 'no vino' de 'vino en 0' ===")

	// Actualización 1: el cliente SOLO quiere cambiar el precio
	json1 := `{"precio_centavos": 39999900}`

	// Actualización 2: el cliente SÍ quiere poner el stock en 0
	// (por ejemplo, para marcarlo como agotado a propósito)
	json2 := `{"stock": 0}`

	var act1, act2 ActualizarProducto
	if err := json.Unmarshal([]byte(json1), &act1); err != nil {
		fmt.Println("JSON 1 inválido:", err)
		return
	}
	if err := json.Unmarshal([]byte(json2), &act2); err != nil {
		fmt.Println("JSON 2 inválido:", err)
		return
	}

	fmt.Println("\n=== Actualización 1: solo precio ===")
	mostrarCampos(act1)

	fmt.Println("\n=== Actualización 2: stock puesto en 0 a propósito ===")
	mostrarCampos(act2)

	// ─────────────────────────────────────────────────────────
	// APLICANDO LA ACTUALIZACIÓN: solo tocar lo que vino
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Aplicar la actualización sobre un producto existente ===")

	producto := Producto{Nombre: "Notebook", PrecioCentavos: 45_000_000, Stock: 5}
	fmt.Printf("Antes:  %+v\n", producto)

	aplicarActualizacion(&producto, act1)
	fmt.Printf("Después de act1 (solo precio): %+v\n", producto)

	aplicarActualizacion(&producto, act2)
	fmt.Printf("Después de act2 (stock=0):     %+v\n", producto)

	// ─────────────────────────────────────────────────────────
	// SIN PUNTEROS, ESTO SERÍA IMPOSIBLE DE DISTINGUIR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== MAL: con int en vez de *int ===")
	fmt.Println("  type ActualizarMAL struct { Stock int }")
	fmt.Println(`  json.Unmarshal([]byte("{}"), &act)        → act.Stock = 0`)
	fmt.Println(`  json.Unmarshal([]byte(` + "`" + `{"stock":0}` + "`" + `), &act) → act.Stock = 0`)
	fmt.Println("  Los dos casos dan LO MISMO: no hay forma de saber si vino o no")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Campo *T + omitempty  → nil = 'no vino', valor = 'vino con ese dato'")
	fmt.Println("  Uso típico            → PATCH/actualizaciones parciales en APIs")
	fmt.Println("  Antes de aplicar      → chequear != nil, después usar *puntero")
	fmt.Println("  Sin esto              → no se puede distinguir 0/\"\"/false de 'no vino'")
}

func mostrarCampos(a ActualizarProducto) {
	if a.Nombre != nil {
		fmt.Println("  Nombre:", *a.Nombre)
	} else {
		fmt.Println("  Nombre: (no vino)")
	}
	if a.PrecioCentavos != nil {
		fmt.Println("  Precio en centavos:", *a.PrecioCentavos)
	} else {
		fmt.Println("  Precio en centavos: (no vino)")
	}
	if a.Stock != nil {
		fmt.Println("  Stock:", *a.Stock)
	} else {
		fmt.Println("  Stock: (no vino)")
	}
}

type Producto struct {
	Nombre         string
	PrecioCentavos int64
	Stock          int
}

func aplicarActualizacion(p *Producto, a ActualizarProducto) {
	if a.Nombre != nil {
		p.Nombre = *a.Nombre
	}
	if a.PrecioCentavos != nil {
		p.PrecioCentavos = *a.PrecioCentavos
	}
	if a.Stock != nil {
		p.Stock = *a.Stock
	}
}
