package main

import (
	"fmt"
	"sort"
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

	fmt.Println("\n=== Iterar map en orden alfabético ===")

	claves := make([]string, 0, len(precios))
	for k := range precios {
		claves = append(claves, k)
	}
	sort.Strings(claves)

	fmt.Println("Lista de precios (ordenada):")
	for _, k := range claves {
		fmt.Printf("  %-10s $%.2f\n", k, precios[k])
	}

	// Orden numérico
	fmt.Println("\n=== Iterar map en orden numérico ===")

	clavesInt := make([]int, 0, len(colores))
	for k := range colores {
		clavesInt = append(clavesInt, k)
	}
	sort.Ints(clavesInt)
	for _, k := range clavesInt {
		fmt.Printf("  %d: %s\n", k, colores[k])
	}

	// ─────────────────────────────────────────────────────────
	// MODIFICAR UN MAP DURANTE LA ITERACIÓN
	// ─────────────────────────────────────────────────────────
	// Go PERMITE agregar o eliminar claves durante el range,
	// pero el comportamiento tiene reglas:
	//
	// - Si eliminás una clave que AÚN NO fue visitada, no se visita.
	// - Si agregás una clave nueva, puede o no aparecer en el mismo range.
	// - Lo más seguro: no modificar el map mientras lo recorrés.

	fmt.Println("\n=== Modificar durante iteración (eliminar) ===")

	stock := map[string]int{
		"notebook":    3,
		"mouse":       0,
		"teclado":     5,
		"monitor":     0,
		"auriculares": 2,
	}

	// BIEN: recolectar claves a eliminar primero, después eliminar
	var sinStock []string
	for producto, cantidad := range stock {
		if cantidad == 0 {
			sinStock = append(sinStock, producto)
		}
	}
	for _, p := range sinStock {
		delete(stock, p)
		fmt.Printf("  Eliminado '%s' (sin stock)\n", p)
	}
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
	claves2 := make([]string, 0, len(frecuencia))
	for k := range frecuencia {
		claves2 = append(claves2, k)
	}
	sort.Strings(claves2)
	for _, p := range claves2 {
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
	for k, v := range codigosPais {
		fmt.Printf("  %s → %s\n", k, v)
	}
	fmt.Println("Invertido (país → código):")
	for k, v := range paisCodigos {
		fmt.Printf("  %s → %s\n", k, v)
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

	cats := make([]string, 0, len(porCategoria))
	for k := range porCategoria {
		cats = append(cats, k)
	}
	sort.Strings(cats)

	for _, cat := range cats {
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
	fmt.Println("   Para ordenar: extraé claves → sort → iterá las claves")
	fmt.Println("   Para eliminar durante iteración: guardá claves → eliminá después")
}
