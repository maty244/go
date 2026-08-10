package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// STRUCTS ANIDADOS Y ARRAYS DENTRO DEL JSON
// =========================================================
// El JSON real casi nunca es plano: un pedido tiene un cliente,
// el cliente tiene una dirección, el pedido tiene una LISTA de
// items... encoding/json maneja todo esto de forma natural: un
// campo puede ser OTRO struct, o un slice de structs.

type Direccion struct {
	Calle  string `json:"calle"`
	Ciudad string `json:"ciudad"`
}

type Cliente struct {
	Nombre    string    `json:"nombre"`
	Direccion Direccion `json:"direccion"` // struct anidado
}

type Item struct {
	Producto string  `json:"producto"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

type Pedido struct {
	ID      int     `json:"id"`
	Cliente Cliente `json:"cliente"`
	Items   []Item  `json:"items"` // slice de structs anidado
}

func main() {
	fmt.Println("=== Marshal de structs anidados ===")

	pedido := Pedido{
		ID: 4521,
		Cliente: Cliente{
			Nombre: "Matias",
			Direccion: Direccion{
				Calle:  "Av. Siempre Viva 742",
				Ciudad: "Buenos Aires",
			},
		},
		Items: []Item{
			{Producto: "Notebook", Cantidad: 1, Precio: 450000},
			{Producto: "Mouse", Cantidad: 2, Precio: 8500},
		},
	}

	datos, _ := json.MarshalIndent(pedido, "", "  ")
	fmt.Println(string(datos))

	// ─────────────────────────────────────────────────────────
	// UNMARSHAL DE JSON ANIDADO: mismo patrón, más profundidad
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Unmarshal de JSON anidado ===")

	textoJSON := `{
		"id": 999,
		"cliente": {
			"nombre": "Ana",
			"direccion": {"calle": "San Martín 123", "ciudad": "Córdoba"}
		},
		"items": [
			{"producto": "Teclado", "cantidad": 1, "precio": 12000}
		]
	}`

	var nuevoPedido Pedido
	if err := json.Unmarshal([]byte(textoJSON), &nuevoPedido); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Cliente:", nuevoPedido.Cliente.Nombre)
	fmt.Println("Ciudad:", nuevoPedido.Cliente.Direccion.Ciudad)
	for _, item := range nuevoPedido.Items {
		fmt.Printf("  %d x %s ($%.2f)\n", item.Cantidad, item.Producto, item.Precio)
	}

	// ─────────────────────────────────────────────────────────
	// EMBEDDING (composición de structs) TAMBIÉN FUNCIONA
	// ─────────────────────────────────────────────────────────
	// Si "embebés" un struct (sin nombre de campo), sus campos se
	// "aplanan" en el JSON en vez de anidarse.

	fmt.Println("\n=== Embedding: los campos se aplanan ===")

	type Auditoria struct {
		CreadoPor string `json:"creado_por"`
	}

	type PedidoConAuditoria struct {
		Auditoria     // embebido: sin nombre de campo
		ID        int `json:"id"`
	}

	pa := PedidoConAuditoria{Auditoria: Auditoria{CreadoPor: "sistema"}, ID: 1}
	datosPA, _ := json.Marshal(pa)
	fmt.Println(string(datosPA))
	// {"creado_por":"sistema","id":1} → SIN anidar, todo al mismo nivel

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Un campo puede ser        → otro struct (JSON anidado)")
	fmt.Println("  o un []struct             → array JSON de objetos")
	fmt.Println("  Marshal/Unmarshal         → funcionan recursivamente, sin límite de profundidad")
	fmt.Println("  Embedding (sin nombre)    → los campos se APLANAN en el JSON")
}
