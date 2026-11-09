package main

import (
	"fmt"

	"gocito/07_paquetes_modulos_y_herramientas/03_visibilidad/catalogo"
)

// =========================================================
// VISIBILIDAD: MAYÚSCULA VS MINÚSCULA
// =========================================================
// Mirá catalogo/catalogo.go: tiene funciones y campos que
// empiezan con mayúscula (Producto, NuevoProducto, Vender, Stock)
// y otros con minúscula (stock, margenGanancia, catalogoInterno,
// necesitaReponer).
//
// Desde ESTE archivo (un paquete DISTINTO) solo podemos usar lo
// que empieza con mayúscula. Es la forma que tiene Go de decir
// "esto es parte de la API pública del paquete" vs "esto es un
// detalle de implementación, no lo toques desde afuera".

func main() {
	fmt.Println("=== Visibilidad entre paquetes ===")

	notebook := catalogo.NuevoProducto("Notebook", 450000, 5)
	fmt.Printf("Creado: %s a $%.2f\n", notebook.Nombre, notebook.Precio)

	// notebook.stock  ← ERROR de compilación si lo descomentás:
	// "stock" no existe fuera del paquete catalogo (es privado).
	// Por eso el paquete nos da Stock() para LEER el valor.
	fmt.Println("Stock actual:", notebook.Stock())

	if err := notebook.Vender(2); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stock después de vender 2:", notebook.Stock())

	if err := notebook.Vender(2); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Stock después de vender 2 más:", notebook.Stock())

	// notebook.necesitaReponer()  ← tampoco compilaría: es privada
	// catalogo.margenGanancia     ← tampoco: es privada del paquete

	// ─────────────────────────────────────────────────────────
	// POR QUÉ ESTO ES ÚTIL: proteger invariantes
	// ─────────────────────────────────────────────────────────
	// Si "stock" fuera público (Stock en vez de stock), cualquier
	// código externo podría hacer notebook.Stock = -500, rompiendo
	// la regla de que el stock nunca es negativo. Al ser privado
	// y exponer solo Vender() (que valida) y Stock() (solo lectura),
	// el paquete catalogo GARANTIZA que su propio estado siempre es
	// válido, sin importar quién lo use.

	fmt.Println("\n=== Intentando romper una regla (bloqueado por el compilador) ===")
	fmt.Println("  notebook.stock = -500  → no compila: 'stock' es privado")
	fmt.Println("  La única forma de bajar el stock es Vender(), que valida cantidad")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  Mayúscula inicial  → exportado: visible desde otros paquetes")
	fmt.Println("  minúscula inicial  → privado: solo visible DENTRO del paquete")
	fmt.Println("  Aplica a           → funciones, tipos, campos, vars, consts")
	fmt.Println("  Para qué sirve     → proteger el estado interno y su validez")
	fmt.Println("  Patrón común       → campo privado + función pública que valida")
}
