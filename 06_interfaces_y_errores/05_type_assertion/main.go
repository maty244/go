package main

import "fmt"

// =========================================================
// TYPE ASSERTION: recuperar el tipo concreto
// =========================================================
// Cuando tenés un valor guardado en una interfaz, a veces
// necesitás volver a saber QUÉ TIPO CONCRETO es en realidad,
// para acceder a algo que no está en el contrato de la interfaz.
//
// La sintaxis es: valor.(TipoConcreto)
//
// Tiene DOS formas:
//   1. v := i.(T)       → PANIC si i no es de tipo T (insegura)
//   2. v, ok := i.(T)   → ok=false si no es T, sin panic (segura)
//
// Casi siempre usá la forma con "ok". La forma sin "ok" solo
// tiene sentido cuando estás 100% seguro del tipo (y aun así,
// un error de programación puede tirar abajo el programa).

type Animal interface {
	Sonido() string
}

type Perro struct{ Nombre string }

func (p Perro) Sonido() string { return "Guau" }
func (p Perro) Buscar()        { fmt.Println(p.Nombre, "trae la pelota") } // NO está en Animal

type Gato struct{ Nombre string }

func (g Gato) Sonido() string { return "Miau" }

func main() {
	// ─────────────────────────────────────────────────────────
	// FORMA INSEGURA: v := i.(T)  → panic si falla
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Type assertion (forma insegura) ===")

	var a Animal = Perro{Nombre: "Rex"}

	p := a.(Perro) // sabemos que es Perro, no hay riesgo acá
	fmt.Println("Recuperamos el tipo concreto:", p.Nombre)
	p.Buscar() // ahora sí podemos usar Buscar(), que no está en Animal

	// Esto SÍ paniquearía porque 'a' es un Perro, no un Gato:
	// g := a.(Gato) // panic: interface conversion

	// ─────────────────────────────────────────────────────────
	// FORMA SEGURA: v, ok := i.(T)  → recomendada
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Type assertion (forma segura) ===")

	animales := []Animal{Perro{Nombre: "Rex"}, Gato{Nombre: "Michi"}}

	for _, animal := range animales {
		if perro, ok := animal.(Perro); ok {
			fmt.Printf("  %s es un Perro → ", perro.Nombre)
			perro.Buscar()
		} else {
			fmt.Printf("  No es un Perro (sonido: %s)\n", animal.Sonido())
		}
	}

	// ─────────────────────────────────────────────────────────
	// CUÁNDO USAR TYPE ASSERTION
	// ─────────────────────────────────────────────────────────
	// Señal de alerta: si estás usando type assertion todo el
	// tiempo para "deshacer" una interfaz, probablemente la
	// interfaz está mal diseñada, o te falta un método en el
	// contrato. Usalo para casos puntuales, no como regla general.

	// ─────────────────────────────────────────────────────────
	// CASO REAL: manejar un tipo especial dentro de un conjunto genérico
	// ─────────────────────────────────────────────────────────
	// Un sistema de notificaciones donde la mayoría de los canales
	// se tratan igual, pero el canal "Prioritario" necesita un
	// tratamiento especial (por ejemplo, reintentar si falla).

	fmt.Println("\n=== Caso real: canal prioritario ===")

	canales := []Canal{CanalEmail{}, CanalPrioritario{Reintentos: 3}, CanalEmail{}}

	for _, c := range canales {
		if prioritario, ok := c.(CanalPrioritario); ok {
			fmt.Printf("  Canal prioritario: reintenta %d veces si falla\n", prioritario.Reintentos)
		} else {
			fmt.Println("  Canal normal: un solo intento")
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  v := i.(T)       → recupera el tipo concreto, PANIC si falla`)
	fmt.Println(`  v, ok := i.(T)   → misma idea, pero ok=false en vez de panic`)
	fmt.Println(`  Usalo cuando     → necesitás un método que NO está en la interfaz`)
	fmt.Println(`  Preferí siempre  → la forma con "ok", salvo certeza absoluta`)
}

type Canal interface {
	Enviar(msg string)
}

type CanalEmail struct{}

func (CanalEmail) Enviar(msg string) { fmt.Println("  [Email]", msg) }

type CanalPrioritario struct {
	Reintentos int
}

func (CanalPrioritario) Enviar(msg string) { fmt.Println("  [Prioritario]", msg) }
