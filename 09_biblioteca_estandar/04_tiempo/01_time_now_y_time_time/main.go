package main

import (
	"fmt"
	"time"
)

// =========================================================
// time.Now() Y EL TIPO time.Time
// =========================================================
// El paquete "time" de la librería estándar maneja fechas, horas
// y duraciones. Todo gira alrededor de UN tipo central:
//
//   time.Time → representa un INSTANTE específico (fecha + hora
//               + zona horaria), con precisión de nanosegundos.
//
// time.Now() devuelve el instante actual, según el reloj de la
// máquina donde corre el programa.

func main() {
	fmt.Println("=== time.Now(): el instante actual ===")

	ahora := time.Now()
	fmt.Println("Ahora:", ahora)

	// ─────────────────────────────────────────────────────────
	// EXTRAER COMPONENTES DE UN time.Time
	// ─────────────────────────────────────────────────────────
	// time.Time tiene MÉTODOS para cada componente (no son campos
	// públicos, siguiendo 07_paquetes_modulos_y_herramientas/03_visibilidad: se accede a través de la
	// API que el paquete expone).

	fmt.Println("\n=== Componentes de la fecha ===")
	fmt.Println("Año:      ", ahora.Year())
	fmt.Println("Mes:      ", ahora.Month()) // devuelve un time.Month (se imprime como texto)
	fmt.Println("Día:      ", ahora.Day())
	fmt.Println("Hora:     ", ahora.Hour())
	fmt.Println("Minuto:   ", ahora.Minute())
	fmt.Println("Segundo:  ", ahora.Second())
	fmt.Println("Día de la semana:", ahora.Weekday())

	// ─────────────────────────────────────────────────────────
	// time.Month ES UN TIPO PROPIO, NO UN int SUELTO
	// ─────────────────────────────────────────────────────────
	// Esto evita el clásico bug de otros lenguajes donde "mes"
	// puede empezar en 0 o en 1 según quién lo mire. En Go, podés
	// comparar directamente con las constantes con nombre.

	fmt.Println("\n=== time.Month es un tipo con nombre, no un int suelto ===")
	if ahora.Month() == time.December {
		fmt.Println("Es diciembre")
	} else {
		fmt.Println("No es diciembre (mes actual:", ahora.Month(), ")")
	}

	// Pero SÍ se puede convertir a int si hace falta un número:
	mesComoNumero := int(ahora.Month())
	fmt.Println("Mes como número:", mesComoNumero)

	// ─────────────────────────────────────────────────────────
	// CREAR UN time.Time ESPECÍFICO (no "ahora")
	// ─────────────────────────────────────────────────────────
	// time.Date construye un instante puntual. Muy útil para
	// pruebas, fechas fijas de negocio (vencimientos, feriados).

	fmt.Println("\n=== Crear una fecha específica con time.Date ===")

	navidad := time.Date(2026, time.December, 25, 0, 0, 0, 0, time.UTC)
	fmt.Println("Navidad 2026:", navidad)

	// ─────────────────────────────────────────────────────────
	// COMPARAR DOS time.Time
	// ─────────────────────────────────────────────────────────
	// NUNCA uses == para comparar time.Time con la intención de
	// "son la misma fecha/hora" de forma confiable (por temas de
	// zona horaria internos). Usá los métodos Before/After/Equal.

	fmt.Println("\n=== Comparar fechas ===")
	fmt.Println("¿Ahora es antes de Navidad?", ahora.Before(navidad))
	fmt.Println("¿Ahora es después de Navidad?", ahora.After(navidad))

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  time.Now()         → el instante actual (time.Time)")
	fmt.Println("  time.Date(...)     → construir un instante específico")
	fmt.Println("  .Year()/.Month()/  → extraer componentes con métodos")
	fmt.Println("  .Day()/.Hour()...")
	fmt.Println("  time.Month         → tipo con nombre, no un int suelto")
	fmt.Println("  Before/After/Equal → comparar dos time.Time (no usar ==)")
}
