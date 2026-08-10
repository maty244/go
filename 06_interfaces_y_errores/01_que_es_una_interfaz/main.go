package main

import "fmt"

// =========================================================
// ¿QUÉ ES UNA INTERFAZ?
// =========================================================
// Hasta ahora, tus funciones reciben un TIPO CONCRETO:
//   func cobrar(t Tarjeta) { ... }
// Esa función SOLO acepta Tarjeta. Si mañana agregás "Efectivo",
// tenés que escribir otra función casi idéntica.
//
// Una interfaz es un CONTRATO: en vez de decir "necesito ESTE
// tipo", decís "necesito CUALQUIER COSA que sepa hacer ESTO".
//
// No importa QUÉ ES el dato, importa QUÉ SABE HACER.
// Es como un enchufe: a la pared no le importa si conectás una
// lámpara o un cargador, solo que la ficha entre.

// ─────────────────────────────────────────────────────────
// DEFINIR UNA INTERFAZ
// ─────────────────────────────────────────────────────────
// Una interfaz es una lista de métodos. Cualquier tipo que
// tenga TODOS esos métodos, automáticamente "cumple" la interfaz.

type Saludador interface {
	Saludar() string
}

// ─────────────────────────────────────────────────────────
// TIPOS QUE CUMPLEN LA INTERFAZ
// ─────────────────────────────────────────────────────────
// Ninguno de estos tipos "declara" que implementa Saludador.
// Simplemente TIENEN el método Saludar() string, y con eso alcanza.

type Persona struct {
	Nombre string
}

func (p Persona) Saludar() string {
	return "Hola, soy " + p.Nombre
}

type Robot struct {
	ID int
}

func (r Robot) Saludar() string {
	return fmt.Sprintf("BEEP BOOP, unidad #%d reportándose", r.ID)
}

// ─────────────────────────────────────────────────────────
// UNA FUNCIÓN QUE ACEPTA LA INTERFAZ, NO EL TIPO CONCRETO
// ─────────────────────────────────────────────────────────
// saludarA acepta CUALQUIER cosa que tenga Saludar() string.
// No sabe (ni le importa) si es una Persona, un Robot, o
// cualquier otro tipo que definamos mañana.

func saludarA(s Saludador) {
	fmt.Println(s.Saludar())
}

func main() {
	// ─────────────────────────────────────────────────────────
	// USANDO LA INTERFAZ
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== Interfaces: el mismo código para tipos distintos ===")

	mati := Persona{Nombre: "Matias"}
	r2d2 := Robot{ID: 2}

	// La MISMA función funciona para los dos tipos:
	saludarA(mati)
	saludarA(r2d2)

	// ─────────────────────────────────────────────────────────
	// UNA VARIABLE DE TIPO INTERFAZ
	// ─────────────────────────────────────────────────────────
	// Una variable declarada con el tipo interfaz puede guardar
	// CUALQUIER valor que cumpla el contrato.

	fmt.Println("\n=== Variable de tipo interfaz ===")

	var quienSea Saludador
	quienSea = mati
	fmt.Println("quienSea =", quienSea.Saludar())

	quienSea = r2d2 // reasignamos con OTRO tipo, sin problema
	fmt.Println("quienSea =", quienSea.Saludar())

	// ─────────────────────────────────────────────────────────
	// SLICE DE INTERFACES: mezclar tipos distintos
	// ─────────────────────────────────────────────────────────
	// Esto es imposible con tipos concretos: no podés tener un
	// []Persona que también guarde Robots. Pero con la interfaz sí.

	fmt.Println("\n=== Slice de Saludador (tipos mezclados) ===")

	saludadores := []Saludador{
		Persona{Nombre: "Ana"},
		Robot{ID: 7},
		Persona{Nombre: "Carlos"},
	}

	for _, s := range saludadores {
		saludarA(s)
	}

	// ─────────────────────────────────────────────────────────
	// CASO REAL: notificaciones
	// ─────────────────────────────────────────────────────────
	// En un kiosco digital, querés avisar al cliente de distintas
	// formas (SMS, email, WhatsApp) sin que el código que envía
	// el aviso sepa los detalles de cada canal.

	fmt.Println("\n=== Caso real: notificaciones ===")

	type Notificador interface {
		Notificar(mensaje string)
	}

	notificar := func(n Notificador, msg string) {
		n.Notificar(msg)
	}

	notificar(canalSMS{}, "Tu pedido está en camino")
	notificar(canalEmail{"mati@kiosco.com"}, "Tu pedido está en camino")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  interface        → contrato: una lista de métodos")
	fmt.Println("  Cumplir el       → CUALQUIER tipo con esos métodos")
	fmt.Println("  contrato            cumple, sin declararlo explícitamente")
	fmt.Println("  Ventaja          → una función/slice acepta tipos distintos")
	fmt.Println("  Filosofía Go     → importa QUÉ SABE HACER, no QUÉ ES")
}

// canalSMS y canalEmail cumplen Notificador (definida dentro de main,
// pero los tipos que la implementan viven a nivel de paquete).
type canalSMS struct{}

func (canalSMS) Notificar(mensaje string) {
	fmt.Println("  [SMS] →", mensaje)
}

type canalEmail struct {
	direccion string
}

func (c canalEmail) Notificar(mensaje string) {
	fmt.Printf("  [Email a %s] → %s\n", c.direccion, mensaje)
}
