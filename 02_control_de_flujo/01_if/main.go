package main

import "fmt"

func main() {
	// =========================================================
	// CONDICIONAL IF EN GO
	// =========================================================
	// El if evalúa una expresión booleana y ejecuta un bloque
	// de código si el resultado es true.
	//
	// Diferencias importantes con otros lenguajes:
	//   ✓ Las condiciones NO llevan paréntesis (pero se permiten)
	//   ✓ Las llaves { } son OBLIGATORIAS siempre
	//   ✓ No hay versión de una sola línea sin llaves
	//   ✓ Go tiene una forma especial: if con inicialización

	// ─────────────────────────────────────────────────────────
	// IF BÁSICO
	// ─────────────────────────────────────────────────────────
	temperatura := 38.5

	fmt.Println("=== if básico ===")
	if temperatura > 37.5 {
		fmt.Println("Tenés fiebre, quedate en casa")
	}

	// Las llaves { y el } else { van en la MISMA línea (regla del lenguaje, no solo estilo).
	// ¿Por qué? Porque el compilador de Go inserta puntos y comas (;) automáticamente
	// al final de una línea si puede cerrar una sentencia. Si bajaras la llave:
	//   if temperatura > 37.5   ← Go inserta un ';' acá y rompe el if
	//   {
	//       fmt.Println("fiebre")
	//   }
	// Lo mismo pasa con el else: debe escribirse "} else {" en la misma línea.

	// ─────────────────────────────────────────────────────────
	// IF / ELSE (y por qué NO existe el operador ternario ?:)
	// ─────────────────────────────────────────────────────────
	// Nota: En Go NO existe el operador ternario `condicion ? valor1 : valor2`.
	// Los creadores de Go lo omitieron a propósito porque suele anidarse y
	// volver el código difícil de leer. En Go siempre se usa if / else:
	edad := 17

	fmt.Println("\n=== if / else ===")
	if edad >= 18 {
		fmt.Printf("%d años: sos mayor de edad\n", edad)
	} else {
		fmt.Printf("%d años: sos menor de edad\n", edad)
	}

	// ─────────────────────────────────────────────────────────
	// IF / ELSE IF / ELSE
	// ─────────────────────────────────────────────────────────
	// Podés encadenar tantos else if como necesités.
	nota := 72

	fmt.Println("\n=== if / else if / else ===")
	if nota >= 90 {
		fmt.Println("Nota A — Excelente")
	} else if nota >= 80 {
		fmt.Println("Nota B — Muy bien")
	} else if nota >= 70 {
		fmt.Println("Nota C — Bien")
	} else if nota >= 60 {
		fmt.Println("Nota D — Aprobado")
	} else {
		fmt.Println("Nota F — Desaprobado")
	}

	// ─────────────────────────────────────────────────────────
	// IF CON INICIALIZACIÓN (la joya de Go)
	// ─────────────────────────────────────────────────────────
	// Sintaxis: if inicialización; condición { }
	//
	// Podés declarar UNA variable justo antes de la condición.
	// Esa variable solo existe dentro del bloque if/else.
	// Esto mantiene el scope limpio y el código más legible.
	//
	// El caso de uso MÁS COMÚN en Go es capturar el error
	// de una función y verificarlo en la misma línea.

	fmt.Println("\n=== if con inicialización ===")

	// Ejemplo 1: capturar resultado y error en el mismo if
	if resultado, err := dividir(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", resultado)
	}
	// "resultado" y "err" NO existen aquí afuera del if

	if resultado, err := dividir(10, 0); err != nil {
		fmt.Println("Error capturado:", err)
	} else {
		fmt.Printf("resultado: %.1f\n", resultado) // no se ejecuta
	}

	// Ejemplo 2: buscar en un map
	stock := map[string]int{
		"manzanas": 50,
		"bananas":  0,
		"naranjas": 30,
	}

	fmt.Println()
	if cantidad, existe := stock["bananas"]; existe && cantidad > 0 {
		fmt.Printf("Hay %d bananas en stock\n", cantidad)
	} else if existe && cantidad == 0 {
		fmt.Println("Bananas agotadas")
	} else {
		fmt.Println("Producto no encontrado")
	}

	if cantidad, existe := stock["papas"]; existe {
		fmt.Printf("Papas: %d\n", cantidad)
	} else {
		fmt.Println("Papas no están en el catálogo")
	}

	// Ejemplo 3: verificar conversión de tipo
	valores := []any{42, "hola", 3.14, true}
	fmt.Println()
	for _, v := range valores {
		if num, ok := v.(int); ok {
			fmt.Printf("Es un entero: %d, el doble es %d\n", num, num*2)
		} else {
			fmt.Printf("%v no es un entero\n", v)
		}
	}

	// ─────────────────────────────────────────────────────────
	// CONDICIONES COMPUESTAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Condiciones compuestas ===")

	usuario := "admin"
	contraseña := "1234"
	intentosFallidos := 2
	const maxIntentos = 3

	// Condición compleja con && y ||
	if (usuario == "admin" || usuario == "root") &&
		contraseña != "" &&
		intentosFallidos < maxIntentos {
		fmt.Println("Acceso concedido")
	} else if intentosFallidos >= maxIntentos {
		fmt.Println("Cuenta bloqueada por demasiados intentos")
	} else {
		fmt.Println("Credenciales inválidas")
	}

	// ─────────────────────────────────────────────────────────
	// IF ANIDADOS (usar con moderación)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== if anidados ===")

	esMiembro := true
	saldoEnCuenta := 500.0
	montoCompra := 200.0

	if esMiembro {
		if saldoEnCuenta >= montoCompra {
			fmt.Println("Compra aprobada (miembro con saldo suficiente)")
		} else {
			fmt.Println("Saldo insuficiente para el miembro")
		}
	} else {
		fmt.Println("Necesitás ser miembro para comprar")
	}

	// MEJOR PRÁCTICA: aplanar if anidados con early return / guard clauses
	fmt.Println()
	validarCompra(esMiembro, saldoEnCuenta, montoCompra)

	// ─────────────────────────────────────────────────────────
	// NEGACIÓN EN CONDICIONES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Negación ===")

	var listaNegraUsuarios = map[string]bool{"baneado": true}

	nombre := "matias"
	if _, estaBaneado := listaNegraUsuarios[nombre]; !estaBaneado {
		fmt.Printf("'%s' puede acceder al sistema\n", nombre)
	} else {
		fmt.Printf("'%s' está bloqueado\n", nombre)
	}

	// ─────────────────────────────────────────────────────────
	// CASOS ESPECIALES CON BOOLEANOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Booleanos en if ===")

	activado := true

	// No necesitás comparar un bool con true/false
	if activado { // equivale a: if activado == true
		fmt.Println("Sistema activado")
	}

	if !activado { // equivale a: if activado == false
		fmt.Println("Sistema desactivado")
	} else {
		fmt.Println("Sistema activo (correcto)")
	}

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Resumen ===")
	fmt.Println("if cond { }                   → básico")
	fmt.Println("if cond { } else { }          → con alternativa")
	fmt.Println("if c1 { } else if c2 { }      → múltiples ramas")
	fmt.Println("if init; cond { }             → con inicialización (Go idiomático)")
	fmt.Println("if val, ok := m[k]; ok { }    → patrón map/type-assert")
	fmt.Println("if res, err := fn(); err==nil → patrón error handling")
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir %.0f por cero", a)
	}
	return a / b, nil
}

// Guard clauses: verificar condiciones negativas primero y salir (return early)
// Esto evita el "pyramid of doom" de ifs anidados
func validarCompra(esMiembro bool, saldo, monto float64) {
	if !esMiembro {
		fmt.Println("[guard] Rechazado: no es miembro")
		return
	}
	if saldo < monto {
		fmt.Printf("[guard] Rechazado: saldo %.0f insuficiente para compra de %.0f\n", saldo, monto)
		return
	}
	fmt.Printf("[guard] Compra aprobada: $%.0f debitado de saldo $%.0f\n", monto, saldo)
}
