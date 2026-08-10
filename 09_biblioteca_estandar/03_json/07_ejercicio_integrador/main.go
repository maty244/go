package main

import (
	"encoding/json"
	"fmt"
)

// =========================================================
// EJERCICIO INTEGRADOR: PROCESAR UN "REQUEST" JSON DE PEDIDO
// =========================================================
// Simulamos lo que va a pasar de verdad en HTTP/ (más adelante en
// la hoja de ruta): recibir un JSON "de afuera", parsearlo,
// validarlo, y devolver una RESPUESTA en JSON. Acá no hay servidor
// todavía (eso es el próximo bloque), pero la lógica de JSON es
// exactamente la misma.

// ─────────────────────────────────────────────────────────
// LO QUE LLEGA: un pedido en JSON
// ─────────────────────────────────────────────────────────

type ItemPedido struct {
	Producto string `json:"producto"`
	Cantidad int    `json:"cantidad"`
}

type PedidoRequest struct {
	Cliente string       `json:"cliente"`
	Items   []ItemPedido `json:"items"`
	Cupon   *string      `json:"cupon,omitempty"` // opcional: puede no venir
}

// ─────────────────────────────────────────────────────────
// LO QUE DEVOLVEMOS: una respuesta en JSON
// ─────────────────────────────────────────────────────────

type PedidoResponse struct {
	OK      bool     `json:"ok"`
	ID      int      `json:"id,omitempty"`
	Total   float64  `json:"total,omitempty"`
	Errores []string `json:"errores,omitempty"`
}

var precios = map[string]float64{
	"Alfajor":  800,
	"Gaseosa":  1200,
	"Notebook": 450000,
}

var siguienteID = 1000

// procesarPedido valida el request y arma la respuesta.
// Esta función NO sabe nada de HTTP: solo trabaja con los structs
// de Go. Cuando lleguemos a HTTP/, esta misma función se va a
// llamar desde un handler que solo se encarga de leer/escribir
// los bytes JSON del request/response.
func procesarPedido(req PedidoRequest) PedidoResponse {
	var errores []string

	if req.Cliente == "" {
		errores = append(errores, "el campo 'cliente' es obligatorio")
	}
	if len(req.Items) == 0 {
		errores = append(errores, "el pedido debe tener al menos un item")
	}

	total := 0.0
	for _, item := range req.Items {
		precio, existe := precios[item.Producto]
		if !existe {
			errores = append(errores, fmt.Sprintf("producto desconocido: %q", item.Producto))
			continue
		}
		if item.Cantidad <= 0 {
			errores = append(errores, fmt.Sprintf("cantidad inválida para %q", item.Producto))
			continue
		}
		total += precio * float64(item.Cantidad)
	}

	// El cupón es opcional: chequeamos != nil antes de usarlo
	// (mismo patrón visto en el tema 05).
	if req.Cupon != nil && *req.Cupon == "DESCUENTO10" {
		total *= 0.90
	}

	if len(errores) > 0 {
		return PedidoResponse{OK: false, Errores: errores}
	}

	siguienteID++
	return PedidoResponse{OK: true, ID: siguienteID, Total: total}
}

func main() {
	fmt.Println("=== EJERCICIO INTEGRADOR: procesar pedidos en JSON ===")

	requests := []string{
		// Pedido válido con cupón
		`{"cliente":"Matias","items":[{"producto":"Alfajor","cantidad":3},{"producto":"Gaseosa","cantidad":1}],"cupon":"DESCUENTO10"}`,

		// Pedido válido sin cupón
		`{"cliente":"Ana","items":[{"producto":"Notebook","cantidad":1}]}`,

		// Pedido con errores: sin cliente, producto inexistente
		`{"items":[{"producto":"Impresora","cantidad":1}]}`,

		// Pedido sin items
		`{"cliente":"Carlos","items":[]}`,
	}

	for i, textoJSON := range requests {
		fmt.Printf("\n--- Request %d ---\n", i+1)
		fmt.Println("Entrada: ", textoJSON)

		var req PedidoRequest
		if err := json.Unmarshal([]byte(textoJSON), &req); err != nil {
			fmt.Println("JSON inválido:", err)
			continue
		}

		resp := procesarPedido(req)

		salida, _ := json.Marshal(resp)
		fmt.Println("Respuesta:", string(salida))
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 09_biblioteca_estandar/03_json/ ===")
	fmt.Println("  Unmarshal          → convertir el request JSON a PedidoRequest")
	fmt.Println("  Tags (01, 03)      → nombres de campo en minúscula/snake_case")
	fmt.Println("  Structs anidados   → []ItemPedido dentro de PedidoRequest")
	fmt.Println("  *string opcional   → Cupon puede no venir en el JSON")
	fmt.Println("  omitempty          → la respuesta no muestra campos vacíos")
	fmt.Println("  Marshal            → convertir PedidoResponse de vuelta a JSON")
	fmt.Println("\nEsta MISMA lógica es la que va a vivir dentro de un handler HTTP")
	fmt.Println("cuando lleguemos a esa sección de la hoja de ruta.")
}
