package main

import (
	"fmt"
	"time"
)

// =========================================================
// EJERCICIO INTEGRADOR: TURNOS DE UNA PELUQUERÍA
// =========================================================
// Combinamos todo lo visto en 09_biblioteca_estandar/04_tiempo/: time.Date para crear
// turnos, Duration para la duración de cada servicio, Add/Sub
// para calcular horarios y solapamientos, Format para mostrar
// todo lindo, y AddDate para vencimientos de promociones.

type Servicio struct {
	Nombre   string
	Duracion time.Duration
}

type Turno struct {
	Cliente  string
	Servicio Servicio
	Inicio   time.Time
}

// Fin calcula cuándo termina el turno.
func (t Turno) Fin() time.Time {
	return t.Inicio.Add(t.Servicio.Duracion)
}

// SeSuperponeCon indica si dos turnos comparten horario.
func (t Turno) SeSuperponeCon(otro Turno) bool {
	return t.Inicio.Before(otro.Fin()) && otro.Inicio.Before(t.Fin())
}

func (t Turno) String() string {
	return fmt.Sprintf("%s: %s (%s - %s, %v)",
		t.Cliente, t.Servicio.Nombre,
		t.Inicio.Format("15:04"), t.Fin().Format("15:04"),
		t.Servicio.Duracion)
}

func main() {
	fmt.Println("=== AGENDA DE LA PELUQUERÍA ===")

	corte := Servicio{Nombre: "Corte", Duracion: 30 * time.Minute}
	color := Servicio{Nombre: "Color", Duracion: 90 * time.Minute}
	barba := Servicio{Nombre: "Barba", Duracion: 15 * time.Minute}

	dia := time.Date(2026, time.July, 6, 9, 0, 0, 0, time.UTC)

	agenda := []Turno{
		{Cliente: "Matias", Servicio: corte, Inicio: dia.Add(0 * time.Hour)},
		{Cliente: "Ana", Servicio: color, Inicio: dia.Add(1 * time.Hour)},
		{Cliente: "Carlos", Servicio: barba, Inicio: dia.Add(2*time.Hour + 15*time.Minute)},
	}

	fmt.Println("\n=== Turnos del día ===")
	for _, t := range agenda {
		fmt.Println(" -", t)
	}

	// ─────────────────────────────────────────────────────────
	// ¿HAY LUGAR PARA UN TURNO NUEVO?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== ¿Hay lugar para un turno nuevo? ===")

	nuevoTurno := Turno{
		Cliente:  "Sofía",
		Servicio: corte,
		Inicio:   dia.Add(1 * time.Hour), // mismo horario que el turno de Ana
	}

	haySolapamiento := false
	for _, t := range agenda {
		if nuevoTurno.SeSuperponeCon(t) {
			fmt.Printf("  Conflicto con %s (%s - %s)\n",
				t.Cliente, t.Inicio.Format("15:04"), t.Fin().Format("15:04"))
			haySolapamiento = true
		}
	}
	if !haySolapamiento {
		fmt.Println("  Sin conflictos, se puede agendar")
	}

	// ─────────────────────────────────────────────────────────
	// DURACIÓN TOTAL DE LA JORNADA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Duración total ocupada en el día ===")

	var totalOcupado time.Duration
	for _, t := range agenda {
		totalOcupado += t.Servicio.Duracion
	}
	fmt.Printf("  Total: %v (%.1f horas)\n", totalOcupado, totalOcupado.Hours())

	// ─────────────────────────────────────────────────────────
	// PROMOCIÓN CON VENCIMIENTO (AddDate)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Promoción con vencimiento ===")

	promoInicio := time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC)
	promoVencimiento := promoInicio.AddDate(0, 0, 15) // válida 15 días

	fmt.Printf("  Promo activa desde %s hasta %s\n",
		promoInicio.Format("02/01/2006"), promoVencimiento.Format("02/01/2006"))

	hoy := time.Date(2026, time.July, 6, 0, 0, 0, 0, time.UTC)
	if hoy.Before(promoVencimiento) {
		diasRestantes := int(promoVencimiento.Sub(hoy).Hours() / 24)
		fmt.Printf("  Vigente: quedan %d días\n", diasRestantes)
	} else {
		fmt.Println("  La promoción ya venció")
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué se usó de todo 09_biblioteca_estandar/04_tiempo/ ===")
	fmt.Println("  time.Date               → crear los turnos en horarios exactos")
	fmt.Println("  time.Duration           → duración de cada servicio")
	fmt.Println("  Add / Sub               → calcular fin de turno y detectar solapamientos")
	fmt.Println("  Format                  → mostrar horarios legibles (HH:MM)")
	fmt.Println("  AddDate + Before/After  → vigencia de una promoción")
}
