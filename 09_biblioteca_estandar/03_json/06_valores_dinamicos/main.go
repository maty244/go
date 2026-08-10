package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// JSON CON ESTRUCTURA DESCONOCIDA (map[string]any)
// =========================================================
// A veces no sabés de antemano la forma exacta del JSON que vas
// a recibir (o te llega un JSON con estructura variable). Para
// esos casos, en vez de un struct, usás:
//
//   map[string]any
//
// (recordá "any" de 06_interfaces_y_errores/07_interfaz_vacia_any: acepta cualquier tipo).
// encoding/json llena ese map con lo que encuentre, usando estos
// tipos de Go para cada tipo de JSON:
//
//   JSON            → Go
//   objeto {}       → map[string]any
//   array []        → []any
//   string          → string
//   número          → float64  (por defecto, incluso "10")
//   true/false      → bool
//   null            → nil

func main() {
	fmt.Println("=== Unmarshal a map[string]any: estructura desconocida ===")

	textoJSON := `{
		"nombre": "Notebook",
		"precio": 450000,
		"disponible": true,
		"tags": ["oferta", "envio-gratis"],
		"detalles": {"marca": "Acer", "año": 2024}
	}`

	var datos map[string]any
	if err := json.Unmarshal([]byte(textoJSON), &datos); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("nombre:", datos["nombre"])
	fmt.Println("precio:", datos["precio"])
	fmt.Println("disponible:", datos["disponible"])

	// ─────────────────────────────────────────────────────────
	// LOS NÚMEROS LLEGAN COMO float64 POR DEFECTO
	// ─────────────────────────────────────────────────────────
	// Esta es LA trampa más común de map[string]any: aunque el
	// JSON tenga un entero (450000, sin punto decimal), Unmarshal lo decodifica
	// como float64. Con Decoder.UseNumber podés conservarlo como json.Number.

	fmt.Println("\n=== Trampa: por defecto los números son float64 ===")

	precio, ok := datos["precio"].(float64)
	fmt.Println("¿Es float64?", ok, "| valor:", precio)

	precioComoInt := int(precio) // conversión manual si necesitás un int
	fmt.Println("Convertido a int:", precioComoInt)

	// ─────────────────────────────────────────────────────────
	// NAVEGAR ESTRUCTURAS ANIDADAS CON TYPE ASSERTION
	// ─────────────────────────────────────────────────────────
	// Cada nivel de anidamiento requiere su propia type assertion,
	// porque "any" no sabe de antemano qué hay adentro.

	fmt.Println("\n=== Navegar JSON anidado sin struct ===")

	tags, ok := datos["tags"].([]any)
	if ok {
		fmt.Println("Tags:")
		for _, tag := range tags {
			fmt.Println("  -", tag)
		}
	}

	detalles, ok := datos["detalles"].(map[string]any)
	if ok {
		fmt.Println("Marca:", detalles["marca"])
	}

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR map[string]any VS UN STRUCT
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Struct vs map[string]any ===")
	fmt.Println("  Struct           → conocés la forma del JSON de antemano (LA MAYORÍA de los casos)")
	fmt.Println("  map[string]any   → estructura variable/desconocida, o solo necesitás UN campo")
	fmt.Println("  Preferí siempre  → un struct cuando puedas: es más seguro y más fácil de leer")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  map[string]any   → para JSON de forma desconocida o variable")
	fmt.Println("  Números          → float64 por defecto; Decoder.UseNumber da json.Number")
	fmt.Println("  Objetos/arrays   → se convierten en map[string]any / []any anidados")
	fmt.Println("  Requiere         → type assertion en cada nivel para usar los datos")
}
