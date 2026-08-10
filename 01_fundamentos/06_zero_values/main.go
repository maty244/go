package main

import "fmt"

// =========================================================
// ZERO VALUES (VALORES CERO)
// =========================================================
// En Go, cuando declarás una variable SIN asignarle un valor,
// Go automáticamente le asigna el "zero value" (valor cero).
// Esto es una gran ventaja sobre otros lenguajes donde
// las variables no inicializadas tienen "basura" en memoria.
//
// Regla: en Go, TODA variable tiene siempre un valor definido.
// Nunca hay variables con valores indefinidos o null (a menos que
// uses punteros o interfaces, que veremos más adelante).

func main() {
	// ─────────────────────────────────────────────────────────
	// ZERO VALUES DE LOS TIPOS BÁSICOS
	// ─────────────────────────────────────────────────────────

	// Enteros → 0
	var enteroSinValor int
	var int8SinValor int8
	var int64SinValor int64
	var uintSinValor uint

	// Flotantes → 0.0
	var floatSinValor float64
	var float32SinValor float32

	// Booleano → false
	var boolSinValor bool

	// String → "" (string vacío, NO es nil)
	var stringSinValor string

	// Byte y Rune → 0
	var byteSinValor byte
	var runeSinValor rune

	fmt.Println("=== Zero Values de tipos básicos ===")
	fmt.Printf("int     zero value: %d\n", enteroSinValor)
	fmt.Printf("int8    zero value: %d\n", int8SinValor)
	fmt.Printf("int64   zero value: %d\n", int64SinValor)
	fmt.Printf("uint    zero value: %d\n", uintSinValor)
	fmt.Printf("float64 zero value: %f\n", floatSinValor)
	fmt.Printf("float32 zero value: %f\n", float32SinValor)
	fmt.Printf("bool    zero value: %v\n", boolSinValor)
	fmt.Printf("string  zero value: '%s' (string vacío)\n", stringSinValor)
	fmt.Printf("byte    zero value: %d\n", byteSinValor)
	fmt.Printf("rune    zero value: %d\n", runeSinValor)

	// ─────────────────────────────────────────────────────────
	// VERIFICAR SI UN STRING ESTÁ VACÍO
	// ─────────────────────────────────────────────────────────
	// Como el zero value del string es "", podemos chequear fácilmente.
	fmt.Println("\n=== Verificar string vacío ===")
	var nombre string
	if nombre == "" {
		fmt.Println("El nombre no fue asignado (está en zero value)")
	}
	nombre = "Matias"
	if nombre != "" {
		fmt.Println("Ahora el nombre es:", nombre)
	}

	// ─────────────────────────────────────────────────────────
	// ZERO VALUES EN ESTRUCTURAS (STRUCTS)
	// ─────────────────────────────────────────────────────────
	// Cuando creás un struct sin inicializar, cada campo
	// tiene su propio zero value según su tipo.
	type Persona struct {
		Nombre string
		Edad   int
		Activo bool
		Saldo  float64
	}

	var p Persona // todos los campos en zero value
	fmt.Println("\n=== Zero Values en struct ===")
	fmt.Printf("Persona vacía: %+v\n", p)
	// Output: {Nombre: Edad:0 Activo:false Saldo:0}

	// Después la podemos poblar
	p.Nombre = "Ana"
	p.Edad = 30
	p.Activo = true
	p.Saldo = 1500.50
	fmt.Printf("Persona poblada: %+v\n", p)

	// ─────────────────────────────────────────────────────────
	// ZERO VALUES EN ARRAYS
	// ─────────────────────────────────────────────────────────
	// Todos los elementos del array tienen su zero value.
	var numeros [5]int
	var flags [3]bool
	var nombres [4]string

	fmt.Println("\n=== Zero Values en arrays ===")
	fmt.Println("Array de ints:", numeros)    // [0 0 0 0 0]
	fmt.Println("Array de bools:", flags)     // [false false false]
	fmt.Println("Array de strings:", nombres) // [   ] (4 strings vacíos)

	// ─────────────────────────────────────────────────────────
	// ZERO VALUES EN PUNTEROS, SLICES, MAPS, CHANNELS
	// ─────────────────────────────────────────────────────────
	// Para estos tipos de referencia, el zero value es nil.
	// nil significa "no apunta a nada" o "no está inicializado".

	var puntero *int        // nil
	var slice []int         // nil
	var mapa map[string]int // nil
	var canal chan int      // nil

	fmt.Println("\n=== Zero Values de tipos referencia (nil) ===")
	fmt.Printf("*int   zero value: %v\n", puntero) // <nil>
	fmt.Printf("[]int  zero value: %v\n", slice)   // []
	fmt.Printf("map    zero value: %v\n", mapa)    // map[]
	fmt.Printf("chan   zero value: %v\n", canal)   // <nil>

	// Verificar si un slice es nil
	if slice == nil {
		fmt.Println("\nEl slice no fue inicializado (es nil)")
	}

	// ─────────────────────────────────────────────────────────
	// POR QUÉ IMPORTAN LOS ZERO VALUES (CASO REAL)
	// ─────────────────────────────────────────────────────────
	// Imaginemos que estamos contando votos en una elección.
	// No necesitamos inicializar los contadores, ya empiezan en 0.

	type ResultadoEleccion struct {
		CandidatoA int
		CandidatoB int
		CandidatoC int
		VotosNulos int
	}

	var eleccion ResultadoEleccion // todos en 0 por zero value

	// Simulamos votos
	votos := []string{"A", "B", "A", "C", "A", "B", "nulo", "A", "C", "B"}
	for _, voto := range votos {
		switch voto {
		case "A":
			eleccion.CandidatoA++
		case "B":
			eleccion.CandidatoB++
		case "C":
			eleccion.CandidatoC++
		default:
			eleccion.VotosNulos++
		}
	}

	fmt.Println("\n=== Ejemplo real: conteo de votos ===")
	fmt.Printf("Candidato A: %d votos\n", eleccion.CandidatoA)
	fmt.Printf("Candidato B: %d votos\n", eleccion.CandidatoB)
	fmt.Printf("Candidato C: %d votos\n", eleccion.CandidatoC)
	fmt.Printf("Votos nulos: %d\n", eleccion.VotosNulos)

	// ─────────────────────────────────────────────────────────
	// ZERO VALUE VS nil EN STRINGS
	// ─────────────────────────────────────────────────────────
	// Los strings en Go NUNCA son nil (a diferencia de Python/Java).
	// Su zero value es "" (string vacío).
	// Si necesitás representar "string ausente", usá *string o una
	// librería de tipos opcionales.

	fmt.Println("\n=== Strings: nunca son nil ===")
	var s string
	fmt.Printf("string vacío: '%s'\n", s)
	fmt.Printf("¿Es igual a string vacío? %v\n", s == "")
	// s == nil daría ERROR de compilación porque string nunca es nil

	// ─────────────────────────────────────────────────────────
	// TABLA RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla de Zero Values ===")
	fmt.Println("Tipo           | Zero Value")
	fmt.Println("---------------|-----------")
	fmt.Println("int, uint, ... | 0")
	fmt.Println("float32/64     | 0.0")
	fmt.Println("bool           | false")
	fmt.Println("string         | \"\" (vacío)")
	fmt.Println("byte, rune     | 0")
	fmt.Println("pointer        | nil")
	fmt.Println("slice          | nil")
	fmt.Println("map            | nil")
	fmt.Println("channel        | nil")
	fmt.Println("interface      | nil")
	fmt.Println("struct         | cada campo con su zero value")
}
