package main

import (
	"fmt"
	"time"
)

// =========================================================
// SUMAR Y RESTAR TIEMPO
// =========================================================
// Con time.Time (un instante) y time.Duration (una cantidad de
// tiempo), Go te da métodos para moverte en el tiempo: sumar
// duraciones, calcular cuánto pasó entre dos instantes, y sumar
// años/meses/días de calendario.

func main() {
	ahora := time.Date(2026, time.July, 5, 10, 0, 0, 0, time.UTC)
	fmt.Println("=== Instante de partida ===")
	fmt.Println(ahora.Format("2006-01-02 15:04"))

	// ─────────────────────────────────────────────────────────
	// Add: SUMAR (O RESTAR) UNA Duration
	// ─────────────────────────────────────────────────────────
	// time.Time + Duration → otro time.Time

	fmt.Println("\n=== Add: sumar una Duration ===")

	en2Horas := ahora.Add(2 * time.Hour)
	fmt.Println("En 2 horas:", en2Horas.Format("2006-01-02 15:04"))

	hace30Min := ahora.Add(-30 * time.Minute) // Duration negativa = restar
	fmt.Println("Hace 30 minutos:", hace30Min.Format("2006-01-02 15:04"))

	// ─────────────────────────────────────────────────────────
	// Sub: LA DIFERENCIA ENTRE DOS time.Time
	// ─────────────────────────────────────────────────────────
	// time.Time - time.Time → Duration (cuánto tiempo hay entre ambos)

	fmt.Println("\n=== Sub: diferencia entre dos instantes ===")

	inicio := time.Date(2026, time.July, 5, 9, 0, 0, 0, time.UTC)
	fin := time.Date(2026, time.July, 5, 17, 30, 0, 0, time.UTC)

	duracionJornada := fin.Sub(inicio)
	fmt.Println("Duración de la jornada:", duracionJornada)
	fmt.Printf("En horas: %.1f\n", duracionJornada.Hours())

	// ─────────────────────────────────────────────────────────
	// time.Since: ATAJO PARA "cuánto pasó desde X hasta ahora"
	// ─────────────────────────────────────────────────────────
	// Equivale a time.Now().Sub(instante), muy usado para medir
	// cuánto tardó algo.

	fmt.Println("\n=== time.Since: medir cuánto pasó ===")

	inicioProceso := time.Now()
	simularTrabajo()
	tiempoTranscurrido := time.Since(inicioProceso)
	fmt.Println("El proceso tardó:", tiempoTranscurrido)

	// ─────────────────────────────────────────────────────────
	// AddDate: SUMAR AÑOS, MESES Y DÍAS DE CALENDARIO
	// ─────────────────────────────────────────────────────────
	// Distinto de Add: acá sumás unidades de CALENDARIO (respeta
	// meses de distinta longitud, años bisiestos, etc.), no una
	// cantidad fija de tiempo.

	fmt.Println("\n=== AddDate: sumar años/meses/días de calendario ===")

	fechaBase := time.Date(2026, time.January, 31, 0, 0, 0, 0, time.UTC)
	enUnMes := fechaBase.AddDate(0, 1, 0)
	fmt.Println("31 de enero + 1 mes:", enUnMes.Format("2006-01-02"))
	// Ojo: febrero no tiene 31 días, Go "desborda" al mes siguiente

	vencimiento := time.Now().AddDate(0, 0, 30) // hoy + 30 días
	fmt.Println("Vencimiento (hoy + 30 días):", vencimiento.Format("2006-01-02"))

	// ─────────────────────────────────────────────────────────
	// COMPARAR: Before, After, Equal
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Comparar fechas ===")

	fmt.Println("¿inicio es antes que fin?", inicio.Before(fin))
	fmt.Println("¿fin es después que inicio?", fin.After(inicio))

	// ─────────────────────────────────────────────────────────
	// CASO REAL: ¿ya venció algo?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: chequear vencimiento ===")

	vencimientoCupon := time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC)
	haVencido := time.Now().After(vencimientoCupon)
	fmt.Println("¿El cupón venció?", haVencido)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  t.Add(duracion)       → suma/resta una CANTIDAD de tiempo exacta")
	fmt.Println("  t2.Sub(t1)            → Duration entre dos instantes")
	fmt.Println("  time.Since(t)         → atajo: cuánto pasó desde t hasta ahora")
	fmt.Println("  t.AddDate(a, m, d)    → suma años/meses/días de CALENDARIO")
	fmt.Println("  Before/After/Equal    → comparar instantes")
}

func simularTrabajo() {
	suma := 0
	for i := 0; i < 1_000_000; i++ {
		suma += i
	}
	_ = suma
}
