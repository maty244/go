package main

import (
	"fmt"
	"time"
)

// =========================================================
// FORMATEAR Y PARSEAR FECHAS: EL LAYOUT DE REFERENCIA
// =========================================================
// Esta es LA parte más rara (y más preguntada en cualquier curso
// de Go) del paquete time: para decirle a Go CÓMO mostrar o leer
// una fecha, no usás letras tipo "YYYY-MM-DD" como en otros
// lenguajes. Usás una FECHA DE REFERENCIA específica, siempre la
// misma:
//
//   Lunes 2 de enero de 2006, a las 15:04:05 (hora -0700)
//
// Ese instante exacto (2006-01-02 15:04:05 -0700) es el "layout
// de referencia" de Go. Cada número en esa fecha representa una
// PARTE del tiempo:
//
//   2006 → año        01 → mes       02 → día
//   15   → hora (24h)  04 → minuto    05 → segundo
//   -0700 → zona horaria
//
// Para formatear, escribís el layout usando ESOS números en el
// orden y forma que querés mostrar. Parece raro, pero una vez que
// lo entendés, es más flexible que "YYYY-MM-DD".

func main() {
	fmt.Println("=== Format: time.Time → texto ===")

	ahora := time.Date(2026, time.July, 5, 14, 30, 0, 0, time.UTC)

	fmt.Println("Layout \"2006-01-02\":         ", ahora.Format("2006-01-02"))
	fmt.Println("Layout \"02/01/2006\":         ", ahora.Format("02/01/2006"))
	fmt.Println("Layout \"2006-01-02 15:04:05\":", ahora.Format("2006-01-02 15:04:05"))
	fmt.Println("Layout \"15:04\":              ", ahora.Format("15:04"))
	fmt.Println("Layout \"Mon, 02 Jan 2006\":   ", ahora.Format("Mon, 02 Jan 2006"))

	// ─────────────────────────────────────────────────────────
	// LAYOUTS PREDEFINIDOS: no hace falta escribirlos siempre
	// ─────────────────────────────────────────────────────────
	// El paquete time ya trae los formatos más comunes como
	// constantes, para no reescribir el layout cada vez.

	fmt.Println("\n=== Layouts predefinidos ===")
	fmt.Println("time.RFC3339:", ahora.Format(time.RFC3339))
	fmt.Println("time.Kitchen:", ahora.Format(time.Kitchen))
	fmt.Println("time.DateOnly:", ahora.Format(time.DateOnly))
	fmt.Println("time.TimeOnly:", ahora.Format(time.TimeOnly))

	// RFC3339 es EL formato más usado en APIs (JSON, HTTP): es el
	// que vas a ver constantemente cuando trabajes con 09_biblioteca_estandar/03_json/ y 13_http_rest/.

	// ─────────────────────────────────────────────────────────
	// Parse: texto → time.Time (el camino inverso)
	// ─────────────────────────────────────────────────────────
	// Le pasás el MISMO layout que tiene el texto que estás
	// leyendo, y Go lo interpreta según esa forma.
	// Nota: si el texto no incluye zona horaria, time.Parse asume
	// UTC; si querés parsearlo en una zona específica (ej. time.Local),
	// existe time.ParseInLocation(layout, texto, loc).

	fmt.Println("\n=== Parse: texto → time.Time ===")

	fechaTexto := "2026-12-25"
	fecha, err := time.Parse("2006-01-02", fechaTexto)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parseado:", fecha)
		fmt.Println("Día de la semana:", fecha.Weekday())
	}

	// ─────────────────────────────────────────────────────────
	// SI EL LAYOUT NO COINCIDE CON EL TEXTO, DA ERROR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Layout que no coincide con el texto ===")

	_, err = time.Parse("2006-01-02", "25/12/2026") // formato distinto al layout
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: leer una fecha desde un formulario
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: parsear fecha de nacimiento ===")

	fechaNacimiento := "1998-03-15"
	nacimiento, err := time.Parse(time.DateOnly, fechaNacimiento)
	if err == nil {
		edad := time.Now().Year() - nacimiento.Year()
		fmt.Printf("Nació el %s, tiene aproximadamente %d años\n",
			nacimiento.Format("02/01/2006"), edad)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  Fecha de referencia → 2006-01-02 15:04:05 (SIEMPRE esa)`)
	fmt.Println(`  t.Format(layout)    → time.Time → texto, según el layout`)
	fmt.Println(`  time.Parse(layout, texto) → texto → time.Time`)
	fmt.Println(`  time.RFC3339        → el formato más usado en APIs/JSON`)
	fmt.Println(`  El layout DEBE coincidir → con la forma real del texto`)
}
