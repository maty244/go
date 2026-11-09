package main

import (
	"errors"
	"fmt"
)

// =========================================================
// INTERFACES Y NIL INTERFACE TRAP
// =========================================================
// Este es uno de los bugs MÁS FAMOSOS y confusos de Go.
//
// Una interfaz en Go tiene INTERNAMENTE dos componentes:
//
//   type iface struct {
//       type  *typeInfo  // puntero al tipo concreto
//       value unsafe.Pointer // puntero al valor concreto
//   }
//
// Una interfaz es nil SOLO cuando AMBOS son nil:
//   type == nil  AND  value == nil
//
// EL TRAP:
//   Si asignás un puntero nil de tipo concreto a una interfaz,
//   la interfaz NO es nil (tiene información de tipo).
//
//   var p *MiTipo = nil  → p es nil (puntero nil)
//   var i error = p      → i NO es nil (tiene type=*MiTipo, value=nil)
//   i == nil             → FALSE  ← esto sorprende a todo el mundo

// ─────────────────────────────────────────────────────────
// TIPOS PARA LOS EJEMPLOS
// ─────────────────────────────────────────────────────────

type ErrorPersonalizado struct {
	Codigo  int
	Mensaje string
}

func (e *ErrorPersonalizado) Error() string {
	return fmt.Sprintf("error %d: %s", e.Codigo, e.Mensaje)
}

// En Go, llamar un método con receptor puntero nil NO da panic por sí solo;
// solo da panic si el método desreferencia el puntero sin validar `e == nil`.
func (e *ErrorPersonalizado) DescribirSeguro() string {
	if e == nil {
		return "receptor *ErrorPersonalizado es nil (sin panic)"
	}
	return e.Error()
}

// ─────────────────────────────────────────────────────────
// EL BUG CLÁSICO: retornar (*T)(nil) como error
// ─────────────────────────────────────────────────────────

// MAL PATRÓN: esta función retorna un *ErrorPersonalizado nil
// envuelto en la interfaz error → la interfaz NO será nil
func operacionMAL(fallar bool) error {
	var err *ErrorPersonalizado // puntero nil de tipo concreto

	if fallar {
		err = &ErrorPersonalizado{500, "algo salió mal"}
	}

	// PROBLEMA: aunque err == nil, retornamos una interfaz
	// con type=*ErrorPersonalizado, value=nil
	// → el caller verá: err != nil  ← FALSO POSITIVO
	return err
}

// BIEN: retornar nil de tipo interfaz directamente
func operacionBIEN(fallar bool) error {
	if fallar {
		return &ErrorPersonalizado{500, "algo salió mal"}
	}
	return nil // nil de tipo interfaz → la interfaz ES nil
}

// ─────────────────────────────────────────────────────────
// INSPECCIONAR EL INTERIOR DE UNA INTERFAZ
// ─────────────────────────────────────────────────────────

func inspeccionarInterfaz(nombre string, i any) {
	if i == nil {
		fmt.Printf("  %s: interfaz nil (type=nil, value=nil)\n", nombre)
	} else {
		fmt.Printf("  %s: interfaz NO nil (type=%T, value=%v)\n", nombre, i, i)
	}
}

// ─────────────────────────────────────────────────────────
// INTERFAZ COMO CONTENEDOR
// ─────────────────────────────────────────────────────────

type Forma interface {
	Area() float64
	Nombre() string
}

type Circulo struct {
	Radio float64
}

func (c *Circulo) Area() float64  { return 3.14159 * c.Radio * c.Radio }
func (c *Circulo) Nombre() string { return "Círculo" }

type Rectangulo struct {
	Ancho, Alto float64
}

func (r *Rectangulo) Area() float64  { return r.Ancho * r.Alto }
func (r *Rectangulo) Nombre() string { return "Rectángulo" }

// ─────────────────────────────────────────────────────────
// TYPE ASSERTION Y TYPE SWITCH
// ─────────────────────────────────────────────────────────

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{ prefijo string }
type FileLogger struct{ archivo string }

func (l *ConsoleLogger) Log(msg string) {
	fmt.Printf("  [CONSOLE/%s] %s\n", l.prefijo, msg)
}

func (l *FileLogger) Log(msg string) {
	fmt.Printf("  [FILE/%s] %s\n", l.archivo, msg)
}

func usarLogger(l Logger) {
	l.Log("hola desde usarLogger")

	// Type assertion: verificar el tipo concreto
	switch concreto := l.(type) {
	case *ConsoleLogger:
		fmt.Printf("  (es ConsoleLogger con prefijo '%s')\n", concreto.prefijo)
	case *FileLogger:
		fmt.Printf("  (es FileLogger escribiendo en '%s')\n", concreto.archivo)
	default:
		fmt.Printf("  (tipo desconocido: %T)\n", l)
	}
}

// ─────────────────────────────────────────────────────────
// ERRORS.AS: trabajar con tipos de error concretos
// ─────────────────────────────────────────────────────────

type ErrorDB struct {
	Tabla string
	Op    string
}

func (e *ErrorDB) Error() string {
	return fmt.Sprintf("error DB en %s durante %s", e.Tabla, e.Op)
}

func consultarDB(tabla string) error {
	if tabla == "" {
		return &ErrorDB{Tabla: "(vacía)", Op: "SELECT"}
	}
	return nil
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║   INTERFACES Y NIL INTERFACE TRAP ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// EL TRAP: puntero nil envuelto en interfaz
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== El trap: *T nil dentro de una interfaz ===")

	var pNil *ErrorPersonalizado = nil
	var iNil error = nil
	var iTrap error = pNil // ← aquí ocurre el trap

	fmt.Printf("pNil == nil:  %v  (puntero nil, correcto)\n", pNil == nil)
	fmt.Printf("iNil == nil:  %v  (interfaz nil, correcto)\n", iNil == nil)
	fmt.Printf("iTrap == nil: %v  ← TRAMPA: tiene type=*ErrorPersonalizado aunque value=nil\n", iTrap == nil)
	fmt.Println()
	fmt.Println("  Internamente:")
	inspeccionarInterfaz("iNil ", iNil)
	inspeccionarInterfaz("iTrap", iTrap)
	// Nota clave: en Go, llamar un método sobre un receptor puntero nil
	// NO produce panic automáticamente al invocar el método: solo paniquea si el cuerpo
	// intenta desreferenciar el puntero nil (ej. leer e.Codigo) sin antes chequear `if e == nil`.
	fmt.Printf("  pNil.DescribirSeguro(): %s\n", pNil.DescribirSeguro())

	// ─────────────────────────────────────────────────────────
	// EFECTO EN FUNCIONES QUE RETORNAN error
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Bug en función que retorna error ===")

	err1 := operacionMAL(false)
	err2 := operacionBIEN(false)

	fmt.Printf("operacionMAL(false) == nil:  %v  ← FALSO POSITIVO (bug!)\n", err1 == nil)
	fmt.Printf("operacionBIEN(false) == nil: %v  ← correcto\n", err2 == nil)
	fmt.Println()
	fmt.Println("  operacionMAL retornó un *ErrorPersonalizado nil envuelto en error")
	fmt.Println("  → aunque no hay error real, if err != nil { } se activa")
	fmt.Println()
	fmt.Println("  Regla: NUNCA retornes un *T nil como interfaz.")
	fmt.Println("  Si no hay error → retorná nil directamente.")

	// Verificar con un error real
	err3 := operacionMAL(true)
	if err3 != nil {
		fmt.Printf("\n  Error real de operacionMAL: %v\n", err3)
	}

	// ─────────────────────────────────────────────────────────
	// FORMA CORRECTA: interfaz con valores concretos
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Uso correcto de interfaces ===")

	formas := []Forma{
		&Circulo{Radio: 5},
		&Rectangulo{Ancho: 4, Alto: 6},
		&Circulo{Radio: 3},
	}

	var totalArea float64
	for _, f := range formas {
		area := f.Area()
		totalArea += area
		fmt.Printf("  %s → área: %.2f\n", f.Nombre(), area)
	}
	fmt.Printf("  Total área: %.2f\n", totalArea)

	// ─────────────────────────────────────────────────────────
	// TYPE ASSERTION Y TYPE SWITCH
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Type assertion y type switch ===")

	loggers := []Logger{
		&ConsoleLogger{prefijo: "APP"},
		&FileLogger{archivo: "app.log"},
	}

	for _, l := range loggers {
		usarLogger(l)
	}

	// Type assertion segura (con ok)
	var l Logger = &ConsoleLogger{prefijo: "TEST"}
	if cl, ok := l.(*ConsoleLogger); ok {
		fmt.Printf("\n  Type assertion exitosa: prefijo = '%s'\n", cl.prefijo)
	}

	// Type assertion insegura (sin ok → panic si falla)
	// fl := l.(*FileLogger)  → panic si l no es *FileLogger

	// ─────────────────────────────────────────────────────────
	// ERRORS.ASTYPE: extraer tipo de error específico
	// ─────────────────────────────────────────────────────────
	// errors.AsType[T] (Go 1.26+) es la forma moderna de errors.As:
	// busca en la cadena de errores uno del tipo T y lo retorna con ok.
	// Maneja bien el nil trap: si no hay match, ok es false.
	fmt.Println("\n=== errors.AsType: extraer tipo de error ===")

	err := consultarDB("")
	if err != nil {
		if dbErr, ok := errors.AsType[*ErrorDB](err); ok {
			fmt.Printf("  Error de base de datos: tabla='%s', op='%s'\n",
				dbErr.Tabla, dbErr.Op)
		}
	}

	err = consultarDB("usuarios")
	fmt.Printf("  consultarDB('usuarios') == nil: %v\n", err == nil)

	// ─────────────────────────────────────────────────────────
	// RESUMEN DE LAS REGLAS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Reglas para evitar el nil interface trap ===")
	fmt.Println("1. NUNCA retornar (*T)(nil) como interfaz.")
	fmt.Println("   MAL: var e *MiError = nil; return e")
	fmt.Println("   BIEN: return nil")
	fmt.Println()
	fmt.Println("2. Si comparás una interfaz con nil, asegurate que realmente")
	fmt.Println("   sea nil de tipo interfaz, no un puntero nil envuelto.")
	fmt.Println()
	fmt.Println("3. Usar errors.AsType (o errors.As) en lugar de type assertions")
	fmt.Println("   para extraer tipos de error concretos.")
	fmt.Println()
	fmt.Println("4. Type assertions con la forma 'v, ok := i.(T)' (siempre)")
	fmt.Println("   para evitar panics en runtime.")
	fmt.Println()
	fmt.Println("Una interfaz es nil solo cuando type=nil AND value=nil.")
	fmt.Println("Un *T nil dentro de una interfaz tiene type=*T → interfaz != nil.")
}
