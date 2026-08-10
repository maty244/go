package main

import (
	"fmt"
	"time"
)

func main() {
	// =========================================================
	// SWITCH EN BLANCO (Blank Switch / Switch sin expresión)
	// =========================================================
	// Un switch "en blanco" es un switch SIN expresión.
	// En vez de comparar un valor contra los cases,
	// cada case tiene su PROPIA condición booleana.
	//
	// Sintaxis:
	//   switch {
	//   case condicion1:
	//       ...
	//   case condicion2:
	//       ...
	//   }
	//
	// Es EXACTAMENTE equivalente a una cadena de if/else if/else,
	// pero más limpio y legible cuando hay muchas ramas.
	//
	// Se ejecuta el PRIMER case cuya condición sea true.
	// Si ninguno es true, se ejecuta default (si existe).

	// ─────────────────────────────────────────────────────────
	// EJEMPLO BÁSICO: comparación con if/else if
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Comparación: if/else if vs switch blanco ===")

	nota := 85

	// Con if/else if (funciona, pero switch es más legible con muchas ramas)
	fmt.Print("if/else if: ")
	if nota >= 90 {
		fmt.Println("A")
	} else if nota >= 80 {
		fmt.Println("B")
	} else if nota >= 70 {
		fmt.Println("C")
	} else if nota >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}

	// Con switch en blanco (equivalente, más limpio)
	fmt.Print("switch blanco: ")
	switch {
	case nota >= 90:
		fmt.Println("A")
	case nota >= 80:
		fmt.Println("B") // nota=85 entra aquí
	case nota >= 70:
		fmt.Println("C")
	case nota >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH BLANCO CON MÚLTIPLES CONDICIONES POR CASE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Condiciones complejas en cada case ===")

	edad := 25
	tieneLicencia := true
	tieneAuto := false
	usoTransportePublico := true

	switch {
	case edad < 16:
		fmt.Println("Muy joven para conducir")
	case edad >= 16 && !tieneLicencia:
		fmt.Println("Tiene edad pero no tiene licencia")
	case tieneLicencia && tieneAuto:
		fmt.Println("Puede conducir su propio auto")
	case tieneLicencia && !tieneAuto && usoTransportePublico:
		fmt.Println("Tiene licencia pero usa transporte público")
	default:
		fmt.Println("Situación de transporte no determinada")
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 1: Clasificar por hora del día
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Franja horaria ===")

	hora := time.Now().Hour()
	minuto := time.Now().Minute()

	switch {
	case hora >= 0 && hora < 6:
		fmt.Printf("Son las %02d:%02d — Madrugada\n", hora, minuto)
	case hora >= 6 && hora < 12:
		fmt.Printf("Son las %02d:%02d — Mañana\n", hora, minuto)
	case hora >= 12 && hora < 14:
		fmt.Printf("Son las %02d:%02d — Mediodía\n", hora, minuto)
	case hora >= 14 && hora < 20:
		fmt.Printf("Son las %02d:%02d — Tarde\n", hora, minuto)
	default:
		fmt.Printf("Son las %02d:%02d — Noche\n", hora, minuto)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 2: Calcular categoría de IMC (índice de masa corporal)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Categoría IMC ===")

	peso := 75.0   // kg
	altura := 1.75 // metros
	imc := peso / (altura * altura)

	fmt.Printf("Peso: %.1f kg | Altura: %.2f m | IMC: %.2f\n", peso, altura, imc)

	switch {
	case imc < 18.5:
		fmt.Println("Categoría: Bajo peso")
	case imc < 25.0:
		fmt.Println("Categoría: Peso normal ✓")
	case imc < 30.0:
		fmt.Println("Categoría: Sobrepeso")
	case imc < 35.0:
		fmt.Println("Categoría: Obesidad grado I")
	case imc < 40.0:
		fmt.Println("Categoría: Obesidad grado II")
	default:
		fmt.Println("Categoría: Obesidad grado III")
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 3: Categorizar una venta
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Categorías de venta ===")

	ventas := []float64{150.0, 1500.0, 7500.0, 25000.0, 60000.0}

	for _, venta := range ventas {
		var categoria string
		switch {
		case venta < 500:
			categoria = "Micro"
		case venta < 5000:
			categoria = "Pequeña"
		case venta < 20000:
			categoria = "Mediana"
		case venta < 50000:
			categoria = "Grande"
		default:
			categoria = "Enterprise"
		}
		fmt.Printf("  $%8.2f → %s\n", venta, categoria)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 4: Determinar tipo de triángulo
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tipo de triángulo ===")

	triangulos := [][3]float64{
		{5, 5, 5},  // equilátero
		{5, 5, 8},  // isósceles
		{3, 4, 5},  // escaleno
		{3, 4, 10}, // inválido
	}

	for _, lados := range triangulos {
		a, b, c := lados[0], lados[1], lados[2]
		fmt.Printf("Lados (%.0f, %.0f, %.0f): ", a, b, c)

		switch {
		case a+b <= c || a+c <= b || b+c <= a:
			fmt.Println("No es triángulo válido")
		case a == b && b == c:
			fmt.Println("Equilátero (3 lados iguales)")
		case a == b || b == c || a == c:
			fmt.Println("Isósceles (2 lados iguales)")
		default:
			fmt.Println("Escaleno (3 lados diferentes)")
		}
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL 5: Validación de datos con múltiples reglas
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Validación de formulario ===")

	type FormRegistro struct {
		Usuario    string
		Email      string
		Contraseña string
		Edad       int
	}

	formularios := []FormRegistro{
		{"", "user@mail.com", "pass123", 25},         // sin usuario
		{"juan", "", "pass123", 25},                  // sin email
		{"ana", "ana@mail.com", "123", 25},           // contraseña corta
		{"carlos", "carlos@mail.com", "pass123", 15}, // menor de edad
		{"mia", "mia@mail.com", "secreto99", 28},     // válido
	}

	for _, f := range formularios {
		fmt.Printf("Usuario '%s': ", f.Usuario)
		switch {
		case f.Usuario == "":
			fmt.Println("Error: usuario vacío")
		case f.Email == "":
			fmt.Println("Error: email vacío")
		case len(f.Contraseña) < 6:
			fmt.Println("Error: contraseña muy corta (mínimo 6 caracteres)")
		case f.Edad < 18:
			fmt.Println("Error: debe ser mayor de edad")
		default:
			fmt.Println("✓ Registro válido!")
		}
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH BLANCO CON INICIALIZACIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Switch blanco con inicialización ===")

	switch ahora := time.Now(); {
	case ahora.Weekday() == time.Saturday || ahora.Weekday() == time.Sunday:
		fmt.Printf("Hoy es %s → fin de semana\n", ahora.Weekday())
	case ahora.Hour() < 9 || ahora.Hour() > 18:
		fmt.Printf("Hoy es %s %02d:00 → fuera de horario laboral\n",
			ahora.Weekday(), ahora.Hour())
	default:
		fmt.Printf("Hoy es %s %02d:00 → horario laboral\n",
			ahora.Weekday(), ahora.Hour())
	}

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR switch BLANCO vs if/else if
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== ¿switch blanco o if/else if? ===")
	fmt.Println()
	fmt.Println("Usá switch blanco cuando:")
	fmt.Println("  ✓ Tenés 3 o más ramas condicionales")
	fmt.Println("  ✓ Las condiciones son del mismo 'tipo conceptual'")
	fmt.Println("    (todas sobre rangos de nota, de temperatura, de precio, etc.)")
	fmt.Println("  ✓ Querés que el lector vea claramente que son ramas mutuamente")
	fmt.Println("    excluyentes (solo una se ejecuta)")
	fmt.Println()
	fmt.Println("Usá if/else si cuando:")
	fmt.Println("  ✓ Solo tenés 1-2 ramas (un simple if/else alcanza)")
	fmt.Println("  ✓ Las condiciones son muy distintas entre sí")
	fmt.Println("  ✓ Necesitás early return dentro de las ramas")
}
