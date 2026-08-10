package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// JSON.Marshal: DE STRUCT A JSON
// =========================================================
// JSON (JavaScript Object Notation) es el formato universal para
// intercambiar datos entre programas, sobre todo en APIs web: casi
// cualquier API REST que consumas o construyas habla en JSON.
//
// El paquete "encoding/json" de la librería estándar convierte
// entre structs de Go y texto JSON. Es EL paquete que vas a usar
// constantemente cuando construyas tu primera API (HTTP/, más
// adelante en la hoja de ruta).
//
//   json.Marshal(valor)   → convierte un valor de Go A JSON (bytes)
//   json.Unmarshal(datos) → convierte JSON A un valor de Go (tema 02)
//
// "Marshal" es el término que usa Go para "serializar" (convertir
// una estructura de datos a una secuencia de bytes/texto).

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func main() {
	fmt.Println("=== json.Marshal: struct → JSON ===")

	p := Producto{Nombre: "Notebook", Precio: 450000, Stock: 3}

	datos, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// datos es []byte. Para verlo como texto, lo convertimos a string.
	fmt.Println(string(datos))
	// {"Nombre":"Notebook","Precio":450000,"Stock":3}

	// ─────────────────────────────────────────────────────────
	// LOS NOMBRES DE LOS CAMPOS SE MANTIENEN TAL CUAL
	// ─────────────────────────────────────────────────────────
	// Por defecto, JSON usa el MISMO nombre que el campo de Go.
	// Como en Go los campos exportados empiezan con mayúscula
	// (06_interfaces_y_errores/Paquetes ya vieron esto), el JSON generado
	// también tiene esas claves en mayúscula. En el tema 02 vemos
	// cómo controlar esto con "tags".

	// ─────────────────────────────────────────────────────────
	// SOLO SE SERIALIZAN LOS CAMPOS EXPORTADOS
	// ─────────────────────────────────────────────────────────
	// Recordá 07_paquetes_modulos_y_herramientas/03_visibilidad: mayúscula = exportado. json.Marshal NO
	// PUEDE ver campos privados (minúscula), así que los ignora
	// completamente.

	fmt.Println("\n=== Campos privados se ignoran ===")

	c := carritoConDescuentoInterno{Total: 1000, descuentoSecreto: 0.15}
	datosCarrito, _ := json.Marshal(c)
	fmt.Println(string(datosCarrito))
	// {"Total":1000} → descuentoSecreto ni aparece

	// ─────────────────────────────────────────────────────────
	// MARSHAL SOBRE SLICES Y MAPS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Marshal de un slice ===")

	catalogo := []Producto{
		{Nombre: "Mouse", Precio: 8500, Stock: 20},
		{Nombre: "Teclado", Precio: 12000, Stock: 15},
	}
	datosCatalogo, _ := json.Marshal(catalogo)
	fmt.Println(string(datosCatalogo))

	fmt.Println("\n=== Marshal de un map ===")

	precios := map[string]float64{"Mouse": 8500, "Teclado": 12000}
	datosPrecios, _ := json.Marshal(precios)
	fmt.Println(string(datosPrecios))

	// ─────────────────────────────────────────────────────────
	// MarshalIndent: JSON "lindo" con indentación, para debug/logs
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== MarshalIndent: JSON con formato legible ===")

	datosLindos, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(datosLindos))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  json.Marshal(v)              → struct/slice/map de Go → JSON ([]byte)")
	fmt.Println("  json.MarshalIndent(v, \"\", \"  \") → igual, pero con indentación")
	fmt.Println("  Solo campos EXPORTADOS        → los privados se ignoran siempre")
	fmt.Println("  Por defecto                   → la clave JSON = nombre del campo de Go")
}

type carritoConDescuentoInterno struct {
	Total            float64
	descuentoSecreto float64 // privado: json.Marshal no lo ve
}
