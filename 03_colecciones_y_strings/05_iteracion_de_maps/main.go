package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// =========================================================
	// ITERANDO MAPS CON FOR RANGE
	// =========================================================
	// Los maps son colecciones de pares clave→valor.
	// Para recorrerlos, Go usa for range igual que con slices,
	// pero con diferencias MUY importantes.
	//
	// REGLA DE ORO: el orden de iteración de un map NO ESTÁ
	// ESPECIFICADO. Puede cambiar entre pasadas o ejecuciones, pero
	// no hay que describirlo como una fuente de azar con garantías.

	// ─────────────────────────────────────────────────────────
	// ITERACIÓN BÁSICA: clave y valor
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Iteración básica: clave y valor ===")

	precios := map[string]float64{
		"manzana": 1.50,
		"banana":  0.75,
		"naranja": 2.00,
		"uva":     3.25,
		"mango":   4.50,
	}

	for producto, precio := range precios {
		fmt.Printf("%-10s → $%.2f\n", producto, precio)
	}
	// El orden puede cambiar. El lenguaje no promete uno concreto.

	// ─────────────────────────────────────────────────────────
	// ITERACIÓN SOLO DE CLAVES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Solo claves ===")

	for producto := range precios {
		fmt.Println(producto)
	}

	// ─────────────────────────────────────────────────────────
	// ITERACIÓN SOLO DE VALORES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Solo valores ===")

	total := 0.0
	for _, precio := range precios {
		total += precio
	}
	fmt.Printf("Precio total del catálogo: $%.2f\n", total)

	// ─────────────────────────────────────────────────────────
	// ORDEN NO ESPECIFICADO EN ACCIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Demostrando orden no especificado ===")

	colores := map[int]string{
		1: "rojo",
		2: "verde",
		3: "azul",
		4: "amarillo",
		5: "naranja",
	}

	fmt.Println("Primera pasada:")
	for k, v := range colores {
		fmt.Printf("  %d: %s\n", k, v)
	}
	fmt.Println("Segunda pasada (puede ser diferente):")
	for k, v := range colores {
		fmt.Printf("  %d: %s\n", k, v)
	}

	// ─────────────────────────────────────────────────────────
	// ITERAR EN ORDEN DETERMINÍSTICO
	// ─────────────────────────────────────────────────────────
	// Si necesitás orden, debés: extraer las claves → ordenarlas → iterar.
	// Podés hacerlo paso a paso con un slice + slices.Sort, o directo con
	// slices.Sorted(maps.Keys(m)) (Go 1.23+).

	fmt.Println("\n=== Iterar map en orden alfabético ===")

	claves := make([]string, 0, len(precios))
	for k := range precios {
		claves = append(claves, k)
	}
	slices.Sort(claves)

	fmt.Println("Lista de precios (ordenada):")
	for _, k := range claves {
		fmt.Printf("  %-10s $%.2f\n", k, precios[k])
	}

	// Orden numérico en una línea con slices.Sorted + maps.Keys (Go 1.23+)
	fmt.Println("\n=== Iterar map en orden numérico (slices.Sorted + maps.Keys) ===")

	for _, k := range slices.Sorted(maps.Keys(colores)) {
		fmt.Printf("  %d: %s\n", k, colores[k])
	}

	// ─────────────────────────────────────────────────────────
	// MODIFICAR UN MAP DURANTE LA ITERACIÓN
	// ─────────────────────────────────────────────────────────
	// Go PERMITE agregar o eliminar claves durante el range,
	// con reglas claras en la especificación del lenguaje:
	//
	// - Eliminar con delete(m, k) durante el range ES SEGURO: si la clave
	//   aún no fue visitada, ya no se visitará.
	// - Si agregás una clave nueva durante el range, puede o no aparecer
	//   en esa misma pasada (por eso evitá insertar mientras iterás).
	// - En Go moderno (1.21+), para borrar entradas según una condición,
	//   la forma más limpia es maps.DeleteFunc.

	fmt.Println("\n=== Eliminar entradas según condición (maps.DeleteFunc) ===")

	stock := map[string]int{
		"notebook":    3,
		"mouse":       0,
		"teclado":     5,
		"monitor":     0,
		"auriculares": 2,
	}

	maps.DeleteFunc(stock, func(producto string, cantidad int) bool {
		if cantidad == 0 {
			fmt.Printf("  Eliminado '%s' (sin stock)\n", producto)
			return true
		}
		return false
	})
	fmt.Println("Stock activo:", stock)

	// ─────────────────────────────────────────────────────────
	// CASOS PRÁCTICOS REALES
	// ─────────────────────────────────────────────────────────

	// Caso 1: Contar frecuencia de palabras
	fmt.Println("\n=== Caso: frecuencia de palabras ===")

	texto := "go es genial go me encanta go es rápido y go es simple"
	frecuencia := make(map[string]int)

	// Parsing simple del texto
	palabra := ""
	for _, r := range texto + " " {
		if r == ' ' {
			if palabra != "" {
				frecuencia[palabra]++
				palabra = ""
			}
		} else {
			palabra += string(r)
		}
	}

	// Ordenar por clave para output predecible
	for _, p := range slices.Sorted(maps.Keys(frecuencia)) {
		fmt.Printf("  '%-8s': %d veces\n", p, frecuencia[p])
	}

	// Caso 2: Invertir un map (valor → clave)
	fmt.Println("\n=== Caso: invertir map ===")

	codigosPais := map[string]string{
		"AR": "Argentina",
		"BR": "Brasil",
		"CL": "Chile",
		"UY": "Uruguay",
	}

	paisCodigos := make(map[string]string, len(codigosPais))
	for codigo, pais := range codigosPais {
		paisCodigos[pais] = codigo
	}

	fmt.Println("Original (código → país):")
	for _, k := range slices.Sorted(maps.Keys(codigosPais)) {
		fmt.Printf("  %s → %s\n", k, codigosPais[k])
	}
	fmt.Println("Invertido (país → código):")
	for _, k := range slices.Sorted(maps.Keys(paisCodigos)) {
		fmt.Printf("  %s → %s\n", k, paisCodigos[k])
	}

	// Caso 3: Agrupar elementos (map de slices)
	fmt.Println("\n=== Caso: agrupar por categoría ===")

	type Producto struct {
		Nombre    string
		Categoria string
		Precio    float64
	}

	productos := []Producto{
		{"Mouse", "periférico", 25.99},
		{"Teclado", "periférico", 75.50},
		{"Monitor", "pantalla", 450.00},
		{"Notebook", "computadora", 1500.00},
		{"Auriculares", "periférico", 80.00},
		{"Tablet", "computadora", 350.00},
		{"TV 4K", "pantalla", 800.00},
	}

	porCategoria := make(map[string][]Producto)
	for _, p := range productos {
		porCategoria[p.Categoria] = append(porCategoria[p.Categoria], p)
	}

	for _, cat := range slices.Sorted(maps.Keys(porCategoria)) {
		fmt.Printf("\n[%s]\n", cat)
		for _, p := range porCategoria[cat] {
			fmt.Printf("  - %-15s $%.2f\n", p.Nombre, p.Precio)
		}
	}

	// Caso 4: Estadísticas sobre un map
	fmt.Println("\n=== Caso: estadísticas de ventas ===")

	ventas := map[string]float64{
		"enero":   15000.50,
		"febrero": 18200.00,
		"marzo":   12000.75,
		"abril":   22000.00,
		"mayo":    19500.25,
		"junio":   25000.00,
	}

	var totalVentas, maxVenta, minVenta float64
	mesMayor, mesMenor := "", ""
	primero := true

	for mes, venta := range ventas {
		totalVentas += venta
		if primero || venta > maxVenta {
			maxVenta = venta
			mesMayor = mes
		}
		if primero || venta < minVenta {
			minVenta = venta
			mesMenor = mes
		}
		primero = false
	}

	promedio := totalVentas / float64(len(ventas))
	fmt.Printf("Total ventas 1er semestre: $%.2f\n", totalVentas)
	fmt.Printf("Promedio mensual:          $%.2f\n", promedio)
	fmt.Printf("Mejor mes: %s ($%.2f)\n", mesMayor, maxVenta)
	fmt.Printf("Peor mes:  %s ($%.2f)\n", mesMenor, minVenta)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("for k, v := range m   → clave y valor")
	fmt.Println("for k := range m      → solo claves")
	fmt.Println("for _, v := range m   → solo valores")
	fmt.Println()
	fmt.Println("⚠️ El orden NO está especificado")
	fmt.Println("   Para ordenar: slices.Sorted(maps.Keys(m)) o slices.Sort(claves)")
	fmt.Println("   Para filtrar/borrar: maps.DeleteFunc(m, func(k, v) bool { ... })")
}
