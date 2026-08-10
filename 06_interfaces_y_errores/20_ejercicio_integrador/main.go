package main

import (
	"errors"
	"fmt"
)

// =========================================================
// EJERCICIO INTEGRADOR: REGISTRO DE USUARIO CON MANEJO DE
// ERRORES COMPLETO
// =========================================================
// Combinamos todo lo visto en ErroresAvanzados: errores centinela,
// wrapping con %w, un error personalizado con campos, errors.Join
// para juntar validaciones, y recover() como red de seguridad.

// ─────────────────────────────────────────────────────────
// ERRORES CENTINELA
// ─────────────────────────────────────────────────────────

var (
	ErrEmailVacio    = errors.New("el email no puede estar vacío")
	ErrEmailInvalido = errors.New("el email debe contener @")
	ErrEdadInvalida  = errors.New("la edad debe estar entre 0 y 120")
)

// ─────────────────────────────────────────────────────────
// ERROR PERSONALIZADO CON CAMPOS
// ─────────────────────────────────────────────────────────

type ErrorPasswordDebil struct {
	Longitud  int
	MinimoReq int
}

func (e *ErrorPasswordDebil) Error() string {
	return fmt.Sprintf("contraseña débil: tiene %d caracteres, mínimo %d", e.Longitud, e.MinimoReq)
}

// ─────────────────────────────────────────────────────────
// VALIDACIONES INDIVIDUALES
// ─────────────────────────────────────────────────────────

func validarEmail(email string) error {
	if email == "" {
		return ErrEmailVacio
	}
	tieneArroba := false
	for _, r := range email {
		if r == '@' {
			tieneArroba = true
			break
		}
	}
	if !tieneArroba {
		return fmt.Errorf("validarEmail(%q): %w", email, ErrEmailInvalido)
	}
	return nil
}

func validarPassword(password string) error {
	const minimo = 8
	if len(password) < minimo {
		return &ErrorPasswordDebil{Longitud: len(password), MinimoReq: minimo}
	}
	return nil
}

func validarEdad(edad int) error {
	if edad < 0 || edad > 120 {
		return fmt.Errorf("validarEdad(%d): %w", edad, ErrEdadInvalida)
	}
	return nil
}

// ─────────────────────────────────────────────────────────
// VALIDACIÓN COMPLETA: junta TODOS los errores con errors.Join
// ─────────────────────────────────────────────────────────

func validarRegistro(email, password string, edad int) error {
	return errors.Join(
		validarEmail(email),
		validarPassword(password),
		validarEdad(edad),
	)
}

// ─────────────────────────────────────────────────────────
// REGISTRAR: envuelve la validación con recover() como red de
// seguridad, por si algo revienta de forma inesperada
// ─────────────────────────────────────────────────────────

func registrarUsuario(email, password string, edad int) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("registrarUsuario: panic recuperado: %v", r)
		}
	}()

	if err := validarRegistro(email, password, edad); err != nil {
		return fmt.Errorf("registro rechazado: %w", err)
	}

	fmt.Printf("  Usuario registrado: %s\n", email)
	return nil
}

func main() {
	fmt.Println("=== EJERCICIO INTEGRADOR: registro de usuario ===")

	casos := []struct {
		email, password string
		edad            int
	}{
		{"mati@mail.com", "claveSegura123", 28},
		{"sin-arroba", "123", -5},
		{"", "unaClaveOK", 200},
		{"ana@mail.com", "corta", 30},
	}

	for i, c := range casos {
		fmt.Printf("\n--- Registro %d: email=%q, edad=%d ---\n", i+1, c.email, c.edad)
		err := registrarUsuario(c.email, c.password, c.edad)
		if err == nil {
			continue
		}

		fmt.Println("Falló:", err)

		// Reaccionar distinto según qué salió mal, usando errors.Is
		// y errors.AsType para atravesar TODO el wrapping + el join.
		if errors.Is(err, ErrEmailInvalido) || errors.Is(err, ErrEmailVacio) {
			fmt.Println("  → Sugerencia: revisá el formato del email")
		}
		if pw, ok := errors.AsType[*ErrorPasswordDebil](err); ok {
			fmt.Printf("  → Sugerencia: agregá %d caracteres más a la contraseña\n",
				pw.MinimoReq-pw.Longitud)
		}
		if errors.Is(err, ErrEdadInvalida) {
			fmt.Println("  → Sugerencia: la edad ingresada no es válida")
		}
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN: qué resolvió cada herramienta ===")
	fmt.Println("  Errores centinela (ErrX)   → identificar QUÉ pasó, sin importar el texto")
	fmt.Println("  fmt.Errorf con %w          → agregar contexto (qué función, qué dato)")
	fmt.Println("  ErrorPasswordDebil         → error con datos para dar una sugerencia concreta")
	fmt.Println("  errors.Join                → reportar TODOS los problemas de una vez")
	fmt.Println("  recover() en registrar     → red de seguridad ante un fallo inesperado")
}
