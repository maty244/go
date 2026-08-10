package main

import (
	"fmt"
	"time"
)

// =========================================================
// time.Duration: EL TIPO PARA "CUÁNTO TIEMPO"
// =========================================================
// Mientras time.Time representa UN INSTANTE ("las 15:04 del
// 2026-07-05"), time.Duration representa una CANTIDAD de tiempo
// ("2 horas", "500 milisegundos", "3 días").
//
// Por dentro, Duration es un int64 que cuenta NANOSEGUNDOS, pero
// casi nunca trabajás con ese número directo: usás las constantes
// que trae el paquete.

func main() {
	fmt.Println("=== Constantes de Duration ===")

	fmt.Println("time.Second:", time.Second)
	fmt.Println("time.Minute:", time.Minute)
	fmt.Println("time.Hour:", time.Hour)

	// ─────────────────────────────────────────────────────────
	// CONSTRUIR DURACIONES COMBINANDO CONSTANTES
	// ─────────────────────────────────────────────────────────
	// Multiplicás la constante por un número para armar la
	// duración que necesites.

	fmt.Println("\n=== Construir duraciones ===")

	cincoMinutos := 5 * time.Minute
	dosHorasYMedia := 2*time.Hour + 30*time.Minute
	trescientosMilisegundos := 300 * time.Millisecond

	fmt.Println("5 minutos:          ", cincoMinutos)
	fmt.Println("2 horas y media:    ", dosHorasYMedia)
	fmt.Println("300 milisegundos:   ", trescientosMilisegundos)

	// ─────────────────────────────────────────────────────────
	// CONVERTIR UNA Duration A NÚMERO
	// ─────────────────────────────────────────────────────────
	// Cuando necesitás el valor como número (para mostrarlo,
	// guardarlo, etc.), usás estos métodos.

	fmt.Println("\n=== Convertir Duration a número ===")

	fmt.Printf("dosHorasYMedia en minutos: %.0f\n", dosHorasYMedia.Minutes())
	fmt.Printf("dosHorasYMedia en horas:   %.2f\n", dosHorasYMedia.Hours())
	fmt.Printf("cincoMinutos en segundos:  %.0f\n", cincoMinutos.Seconds())

	// ─────────────────────────────────────────────────────────
	// time.ParseDuration: DESDE UN TEXTO
	// ─────────────────────────────────────────────────────────
	// Útil para leer duraciones desde config, argumentos de línea
	// de comandos, o variables de entorno (09_biblioteca_estandar/02_archivos/, más adelante).

	fmt.Println("\n=== ParseDuration: texto → Duration ===")

	d, err := time.ParseDuration("1h30m")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(`"1h30m" parseado:`, d)
	}

	d2, _ := time.ParseDuration("500ms")
	fmt.Println(`"500ms" parseado:`, d2)

	// ─────────────────────────────────────────────────────────
	// OPERAR CON DURACIONES: sumar, comparar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Operar con duraciones ===")

	tiempoTranscurrido := 45 * time.Minute
	limiteMaximo := 1 * time.Hour

	if tiempoTranscurrido > limiteMaximo {
		fmt.Println("Se superó el límite")
	} else {
		restante := limiteMaximo - tiempoTranscurrido
		fmt.Println("Todavía quedan:", restante)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: calcular un timeout según el tamaño de algo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: timeout proporcional ===")

	itemsAProcesar := 250
	timeoutPorItem := 20 * time.Millisecond
	timeoutTotal := time.Duration(itemsAProcesar) * timeoutPorItem

	fmt.Printf("Timeout calculado para %d items: %v\n", itemsAProcesar, timeoutTotal)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  time.Duration      → 'cuánto tiempo' (por dentro, nanosegundos)")
	fmt.Println("  time.Second/Minute/Hour → constantes para construir duraciones")
	fmt.Println("  d.Minutes()/.Hours() → convertir Duration a número")
	fmt.Println("  time.ParseDuration(\"1h30m\") → texto → Duration")
	fmt.Println("  Se puede sumar, restar y comparar como cualquier número")
}
