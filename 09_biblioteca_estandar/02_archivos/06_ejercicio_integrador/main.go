package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// =========================================================
// EJERCICIO INTEGRADOR: GUARDAR Y CARGAR UN CARRITO EN DISCO
// =========================================================
// Combinamos 09_biblioteca_estandar/03_json/ y 09_biblioteca_estandar/02_archivos/: un carrito de compras que se
// guarda en un archivo (persistencia simple, sin base de datos)
// y se puede volver a cargar más tarde. Es el mismo patrón que
// vas a usar para guardar config, cachés simples, o el estado de
// una CLI entre ejecuciones.

type ItemCarrito struct {
	Producto string  `json:"producto"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

type Carrito struct {
	Cliente string        `json:"cliente"`
	Items   []ItemCarrito `json:"items"`
}

func (c Carrito) Total() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Precio * float64(item.Cantidad)
	}
	return total
}

// guardarCarrito serializa el carrito a JSON y lo escribe a disco.
func guardarCarrito(ruta string, c Carrito) error {
	datos, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("guardarCarrito: %w", err)
	}
	if err := os.WriteFile(ruta, datos, 0644); err != nil {
		return fmt.Errorf("guardarCarrito: %w", err)
	}
	return nil
}

// cargarCarrito lee el archivo y lo convierte de vuelta a Carrito.
func cargarCarrito(ruta string) (Carrito, error) {
	var c Carrito

	datos, err := os.ReadFile(ruta)
	if err != nil {
		return c, fmt.Errorf("cargarCarrito: %w", err)
	}

	if err := json.Unmarshal(datos, &c); err != nil {
		return c, fmt.Errorf("cargarCarrito: %w", err)
	}
	return c, nil
}

func main() {
	fmt.Println("=== KIOSCO DIGITAL: carrito persistente en disco ===")

	ruta := filepath.Join(os.TempDir(), "kiosco_carrito.json")
	defer os.Remove(ruta)

	// ─────────────────────────────────────────────────────────
	// "SESIÓN 1": el cliente arma su carrito y cierra el programa
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Sesión 1: armar el carrito ---")

	carrito := Carrito{
		Cliente: "Matias",
		Items: []ItemCarrito{
			{Producto: "Alfajor", Cantidad: 3, Precio: 800},
			{Producto: "Gaseosa", Cantidad: 1, Precio: 1200},
		},
	}

	if err := guardarCarrito(ruta, carrito); err != nil {
		fmt.Println("Error guardando:", err)
		return
	}
	fmt.Printf("Carrito guardado en %s (total: $%.2f)\n", ruta, carrito.Total())

	// Mostramos el archivo tal cual quedó en disco
	contenido, _ := os.ReadFile(ruta)
	fmt.Println("\nContenido del archivo:")
	fmt.Println(string(contenido))

	// ─────────────────────────────────────────────────────────
	// "SESIÓN 2": el cliente vuelve más tarde, recupera su carrito
	// ─────────────────────────────────────────────────────────
	fmt.Println("--- Sesión 2: el cliente vuelve y recupera el carrito ---")

	carritoRecuperado, err := cargarCarrito(ruta)
	if err != nil {
		fmt.Println("Error cargando:", err)
		return
	}

	fmt.Printf("Bienvenido de nuevo, %s\n", carritoRecuperado.Cliente)
	for _, item := range carritoRecuperado.Items {
		fmt.Printf("  %d x %s ($%.2f c/u)\n", item.Cantidad, item.Producto, item.Precio)
	}
	fmt.Printf("Total: $%.2f\n", carritoRecuperado.Total())

	// ─────────────────────────────────────────────────────────
	// AGREGAR UN ITEM Y VOLVER A GUARDAR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Agregando un item más y guardando de nuevo ---")

	carritoRecuperado.Items = append(carritoRecuperado.Items,
		ItemCarrito{Producto: "Mouse", Cantidad: 1, Precio: 8500})

	guardarCarrito(ruta, carritoRecuperado)
	fmt.Printf("Nuevo total: $%.2f\n", carritoRecuperado.Total())

	// ─────────────────────────────────────────────────────────
	// QUÉ PASA SI EL ARCHIVO NO EXISTE (primera vez del cliente)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n--- Intentar cargar un carrito que no existe ---")

	_, err = cargarCarrito(filepath.Join(os.TempDir(), "no_existe.json"))
	if err != nil {
		fmt.Println("Error esperado:", err)
		fmt.Println("(en un programa real, acá se crearía un carrito vacío nuevo)")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 09_biblioteca_estandar/02_archivos/ ===")
	fmt.Println("  os.WriteFile/ReadFile   → guardar y cargar el carrito completo")
	fmt.Println("  json.MarshalIndent      → convertir el carrito a JSON legible")
	fmt.Println("  json.Unmarshal          → reconstruir el struct al cargar")
	fmt.Println("  os.IsNotExist (implícito en el error) → manejar 'primera vez'")
	fmt.Println("  Patrón general          → así persiste config/estado sin una base de datos")
}
