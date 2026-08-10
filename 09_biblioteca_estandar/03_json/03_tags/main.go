package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// =========================================================
// TAGS DE JSON: CONTROLAR CÓMO SE VE CADA CAMPO
// =========================================================
// Los "tags" son anotaciones que se escriben pegadas al tipo de
// cada campo del struct, entre backticks. Le dicen a encoding/json
// exactamente cómo tratar ese campo.
//
//   Campo string `json:"nombre_json"`
//
// Esto es MUY común porque las convenciones de nombres cambian
// entre lenguajes: Go usa PascalCase (Nombre, Precio) pero JSON
// (y JavaScript) suelen usar camelCase o snake_case (nombre,
// precio_final). Los tags traducen entre ambos mundos.

type ProductoAPI struct {
	Nombre         string    `json:"nombre"`
	PrecioCentavos int64     `json:"precio_centavos"`
	Stock          int       `json:"stock"`
	CodigoInterno  string    `json:"-"`                       // "-": NUNCA aparece
	DescuentoPct   int       `json:"descuento_pct,omitempty"` // omite zero value
	ActualizadoEn  time.Time `json:"actualizado_en,omitzero"` // Go 1.24+: usa IsZero
}

func main() {
	fmt.Println("=== Tags básicos: renombrar campos ===")

	p := ProductoAPI{
		Nombre:         "Notebook",
		PrecioCentavos: 45_000_000,
		Stock:          3,
		CodigoInterno:  "SKU-9921", // este campo NUNCA sale en el JSON
	}

	datos, err := json.Marshal(p)
	if err != nil {
		fmt.Println("No se pudo codificar:", err)
		return
	}
	fmt.Println(string(datos))
	// CodigoInterno no aparece. DescuentoPct y ActualizadoEn tampoco porque
	// tienen su valor cero y usan omitempty/omitzero.

	// ─────────────────────────────────────────────────────────
	// omitempty EN ACCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== omitempty: se omite si el valor es el zero value ===")

	sinDescuento := ProductoAPI{Nombre: "Mouse", PrecioCentavos: 850_000, Stock: 20}
	conDescuento := ProductoAPI{Nombre: "Mouse", PrecioCentavos: 850_000, Stock: 20, DescuentoPct: 10}

	d1, err := json.Marshal(sinDescuento)
	if err != nil {
		fmt.Println("No se pudo codificar:", err)
		return
	}
	d2, err := json.Marshal(conDescuento)
	if err != nil {
		fmt.Println("No se pudo codificar:", err)
		return
	}

	fmt.Println("Sin descuento:", string(d1)) // "descuento" NO aparece
	fmt.Println("Con descuento:", string(d2)) // "descuento" SÍ aparece

	// ─────────────────────────────────────────────────────────
	// UNMARSHAL TAMBIÉN RESPETA LOS TAGS
	// ─────────────────────────────────────────────────────────
	// La traducción funciona en los dos sentidos: si el JSON viene
	// con "nombre" (minúscula), Unmarshal lo mapea al campo Nombre
	// gracias al tag, aunque el campo de Go esté en mayúscula.

	fmt.Println("\n=== Unmarshal respeta los tags ===")

	textoJSON := `{"nombre":"Teclado","precio_centavos":1200000,"stock":15}`
	var nuevo ProductoAPI
	if err := json.Unmarshal([]byte(textoJSON), &nuevo); err != nil {
		fmt.Println("JSON inválido:", err)
		return
	}
	fmt.Printf("  %+v\n", nuevo)

	// ─────────────────────────────────────────────────────────
	// json:"-" : EXCLUIR UN CAMPO SIEMPRE
	// ─────────────────────────────────────────────────────────
	// Útil para datos sensibles (contraseñas, tokens, códigos
	// internos) que NUNCA deberían salir por una API.

	fmt.Println("\n=== json:\"-\" excluye el campo siempre ===")
	fmt.Println("CodigoInterno en Go:", p.CodigoInterno)
	fmt.Println("¿Aparece en el JSON?", "SKU-9921 no está en:", string(datos))

	// ─────────────────────────────────────────────────────────
	// SIN TAG: Go usa el nombre del campo tal cual
	// ─────────────────────────────────────────────────────────
	// Ya lo viste en el tema 01: si no ponés tag, la clave JSON
	// es EXACTAMENTE el nombre del campo de Go (con mayúscula).

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  `json:\"nombre\"`          → renombra la clave JSON")
	fmt.Println("  `json:\"x,omitempty\"`     → omite el campo si es su zero value")
	fmt.Println("  `json:\"x,omitzero\"`      → usa IsZero; útil con time.Time")
	fmt.Println("  `json:\"-\"`               → excluye el campo SIEMPRE")
	fmt.Println("  Sin tag                    → usa el nombre del campo de Go tal cual")
	fmt.Println("  Unmarshal también los usa  → traduce en ambos sentidos")
}
