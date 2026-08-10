package main

import (
	"fmt"

	"gocito/07_paquetes_modulos_y_herramientas/06_ejercicio_integrador/clientes"
	"gocito/07_paquetes_modulos_y_herramientas/06_ejercicio_integrador/productos"
)

// =========================================================
// EJERCICIO INTEGRADOR: KIOSCO DIGITAL EN VARIOS PAQUETES
// =========================================================
// Un mini proyecto real, dividido como lo estaría en la vida
// real: cada RESPONSABILIDAD en su propio paquete.
//
//   06_ejercicio_integrador/
//   ├── main.go              ← package main: conecta todo
//   ├── productos/
//   │   └── productos.go     ← package productos: catálogo y stock
//   └── clientes/
//       └── clientes.go      ← package clientes: registro e historial
//
// main.go NO conoce los detalles internos de cada paquete (por
// ejemplo, no sabe que "stock" es un campo privado protegido por
// Vender()): solo usa la API pública que cada paquete expone.
// Esto es exactamente lo que veníamos armando en los temas
// anteriores, ahora combinado en un ejemplo más completo.

func main() {
	fmt.Println("=== KIOSCO DIGITAL: mini proyecto multi-paquete ===")

	catalogo := productos.NuevoCatalogo()
	catalogo.Agregar("Alfajor", 800, 20)
	catalogo.Agregar("Gaseosa", 1200, 5)

	registro := clientes.NuevoRegistro()
	mati := registro.Alta("Matias")
	ana := registro.Alta("Ana")

	fmt.Println("\n=== Procesando compras ===")

	comprar := func(nombreCliente, nombreProducto string, cantidad int) {
		prod, ok := catalogo.Buscar(nombreProducto)
		if !ok {
			fmt.Printf("  %s no existe en el catálogo\n", nombreProducto)
			return
		}
		if err := prod.Vender(cantidad); err != nil {
			fmt.Println("  Error:", err)
			return
		}
		cliente := registro.Alta(nombreCliente)
		cliente.RegistrarCompra(fmt.Sprintf("%d x %s", cantidad, nombreProducto))
		fmt.Printf("  %s compró %d x %s ($%.2f c/u)\n",
			nombreCliente, cantidad, nombreProducto, prod.Precio)
	}

	comprar("Matias", "Alfajor", 3)
	comprar("Ana", "Gaseosa", 2)
	comprar("Matias", "Gaseosa", 10) // esta falla: no hay stock

	fmt.Println("\n=== Historial de clientes ===")
	fmt.Println(" -", clientes.ResumenCliente(mati))
	fmt.Println(" -", clientes.ResumenCliente(ana))

	fmt.Println("\n=== Stock restante ===")
	for _, nombre := range []string{"Alfajor", "Gaseosa"} {
		if p, ok := catalogo.Buscar(nombre); ok {
			fmt.Printf("  %s: %d unidades\n", p.Nombre, p.Stock())
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué le tocó a cada paquete ===")
	fmt.Println("  productos  → catálogo, precios, stock protegido")
	fmt.Println("  clientes   → alta de clientes, historial de compras")
	fmt.Println("  main       → conecta ambos paquetes, sin conocer sus detalles internos")
}
