package main

import (
	"fmt"
	"time"
)

func main() {
	// =========================================================
	// SWITCH EN GO
	// =========================================================
	// Switch evalúa una expresión y ejecuta el "case" que coincide.
	// Es una alternativa más limpia que una cadena larga de if/else if.
	//
	// Diferencias con switch en C, Java, JavaScript:
	//   ✓ NO hay fallthrough automático (no necesitás break al final de cada case)
	//   ✓ Los cases pueden tener múltiples valores: case "a", "b":
	//   ✓ Funciona con strings, floats, structs y cualquier tipo comparable
	//   ✓ Los cases pueden ser variables o expresiones en ejecución, no solo constantes
	//   ✓ Existe el "switch sin expresión" (ver 03_switch_sin_expresion)
	//   ✓ Existe el "type switch" para descubrir el tipo concreto dentro de una interfaz

	// ─────────────────────────────────────────────────────────
	// SWITCH BÁSICO CON ENTERO
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Switch básico (int) ===")

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles") // ← este se ejecuta
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Número de día inválido")
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH CON STRING
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Switch con string ===")

	lenguaje := "go"

	switch lenguaje {
	case "go":
		fmt.Println("Go: rápido, compilado, con garbage collector")
	case "python":
		fmt.Println("Python: fácil de aprender, interpretado")
	case "rust":
		fmt.Println("Rust: seguro en memoria, sin GC")
	case "java":
		fmt.Println("Java: orientado a objetos, JVM")
	default:
		fmt.Printf("Lenguaje '%s' no reconocido\n", lenguaje)
	}

	// ─────────────────────────────────────────────────────────
	// MÚLTIPLES VALORES EN UN CASE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Múltiples valores en un case ===")

	nombreDia := "sábado"

	switch nombreDia {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Println("Día hábil → hay que trabajar")
	case "sábado", "domingo":
		fmt.Println("Fin de semana → a descansar!")
	default:
		fmt.Println("Día no reconocido")
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH CON INICIALIZACIÓN (igual que el if)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Switch con inicialización ===")

	switch hora := time.Now().Hour(); {
	case hora < 6:
		fmt.Println("Son las", hora, "h: madrugada")
	case hora < 12:
		fmt.Println("Son las", hora, "h: mañana")
	case hora < 18:
		fmt.Println("Son las", hora, "h: tarde")
	default:
		fmt.Println("Son las", hora, "h: noche")
	}

	// ─────────────────────────────────────────────────────────
	// DEFAULT: caso cuando ningún case coincide
	// ─────────────────────────────────────────────────────────
	// default puede ir en cualquier posición (inicio, medio, final).
	// Por convención suele ir al final.

	fmt.Println("\n=== default ===")

	codigoHTTP := 404

	switch codigoHTTP {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Creado")
	case 400:
		fmt.Println("Bad Request")
	case 401:
		fmt.Println("No autorizado")
	case 403:
		fmt.Println("Prohibido")
	case 404:
		fmt.Println("No encontrado") // ← este
	case 500:
		fmt.Println("Error interno del servidor")
	default:
		fmt.Printf("Código %d: sin descripción conocida\n", codigoHTTP)
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH CON EXPRESIONES DINÁMICAS EN LOS CASES
	// ─────────────────────────────────────────────────────────
	// A diferencia de C o Java (donde cada case debe ser una constante fija),
	// en Go los cases pueden usar variables y cálculos evaluados de arriba a abajo:
	fmt.Println("\n=== Cases con expresiones dinámicas ===")

	numeroSecreto := 7
	intento := 6

	switch intento {
	case numeroSecreto:
		fmt.Println("¡Adivinaste el número exacto!")
	case numeroSecreto - 1, numeroSecreto + 1:
		fmt.Println("¡Muy cerca! Estás a 1 de distancia") // ← intento=6 entra aquí
	default:
		fmt.Println("Frío, seguí intentando")
	}

	// Y si omitís el valor tras "switch", podés poner condiciones booleanas
	// completas en cada case ("switch en blanco", ver 03_switch_sin_expresion):
	nota := 85
	switch {
	case nota >= 90:
		fmt.Println("Nota: A - Excelente")
	case nota >= 80:
		fmt.Println("Nota: B - Muy bueno") // ← nota=85 entra aquí
	default:
		fmt.Println("Nota: C o menor")
	}

	// ─────────────────────────────────────────────────────────
	// FALLTHROUGH: pasar explícitamente al siguiente case
	// ─────────────────────────────────────────────────────────
	// En Go el fallthrough es EXPLÍCITO (al revés que en C/Java).
	// Se usa raramente. Ejecuta el CUERPO del siguiente case
	// SIN verificar su condición.

	fmt.Println("\n=== fallthrough (explícito en Go) ===")

	nivel := 2

	switch nivel {
	case 1:
		fmt.Println("Nivel 1: acceso a zona pública")
		fallthrough
	case 2:
		fmt.Println("Nivel 2: acceso a zona de empleados") // se ejecuta
		fallthrough
	case 3:
		fmt.Println("Nivel 3: acceso a zona restringida") // también se ejecuta (por fallthrough)
	case 4:
		fmt.Println("Nivel 4: acceso total")
	}
	// Con nivel=2: ejecuta case 2 y case 3 (por fallthrough de case 2)
	// NOTA: fallthrough no verifica la condición de case 3,
	// simplemente ejecuta su cuerpo.

	// ─────────────────────────────────────────────────────────
	// TYPE SWITCH: verificar el tipo dinámico de una interfaz
	// ─────────────────────────────────────────────────────────
	// Uno de los usos más potentes del switch en Go.
	// Cuando tenés un valor de tipo interface{} (o any),
	// el type switch te permite saber qué tipo real tiene.

	fmt.Println("\n=== Type switch ===")

	valores := []any{
		42,
		"hola mundo",
		3.14,
		true,
		[]int{1, 2, 3},
		nil,
	}

	for _, v := range valores {
		switch t := v.(type) {
		case int:
			fmt.Printf("int: %d (el doble: %d)\n", t, t*2)
		case string:
			fmt.Printf("string: '%s' (longitud: %d)\n", t, len(t))
		case float64:
			fmt.Printf("float64: %.4f\n", t)
		case bool:
			if t {
				fmt.Println("bool: verdadero")
			} else {
				fmt.Println("bool: falso")
			}
		case []int:
			fmt.Printf("[]int: %v (len=%d)\n", t, len(t))
		case nil:
			fmt.Println("nil: valor nulo")
		default:
			fmt.Printf("tipo desconocido: %T\n", t)
		}
	}

	// ─────────────────────────────────────────────────────────
	// SWITCH SIN DEFAULT (válido, pero pensar si tiene sentido)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Switch sin default ===")

	// Si ningún case coincide y no hay default, simplemente no hace nada.
	x := 99
	switch x {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	}
	fmt.Println("Después del switch (sin default, no hizo nada para x=99)")

	// ─────────────────────────────────────────────────────────
	// SWITCH COMO MÁQUINA DE ESTADOS (caso real avanzado)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Máquina de estados: semáforo ===")

	type Estado string
	const (
		Rojo     Estado = "rojo"
		Amarillo Estado = "amarillo"
		Verde    Estado = "verde"
	)

	semaforo := Rojo

	// Simulamos 5 cambios de estado
	for i := 0; i < 5; i++ {
		switch semaforo {
		case Rojo:
			fmt.Printf("🔴 ROJO: Stop!\n")
			semaforo = Verde
		case Verde:
			fmt.Printf("🟢 VERDE: Avanzá\n")
			semaforo = Amarillo
		case Amarillo:
			fmt.Printf("🟡 AMARILLO: Preparate\n")
			semaforo = Rojo
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("switch expr { case val: }         → básico")
	fmt.Println("case v1, v2, v3:                  → múltiples valores")
	fmt.Println("switch init; expr { }             → con inicialización")
	fmt.Println("fallthrough                       → pasar al siguiente case (raro)")
	fmt.Println("switch v := x.(type) { case T: }  → type switch (interfaces)")
	fmt.Println("default:                          → cuando ningún case coincide")
	fmt.Println()
	fmt.Println("NO necesita break → el case termina automáticamente")
	fmt.Println("Para fallthrough explícito → escribir 'fallthrough'")
}
