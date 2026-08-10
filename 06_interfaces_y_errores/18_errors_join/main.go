package main

import (
	"errors"
	"fmt"
)

// =========================================================
// errors.Join: COMBINAR VARIOS ERRORES EN UNO SOLO
// =========================================================
// A veces necesitás juntar VARIOS errores independientes en un
// solo valor error (por ejemplo: validar un formulario entero y
// reportar TODOS los campos inválidos, no solo el primero).
//
// errors.Join(err1, err2, err3) devuelve un único error que:
//   - Al imprimirlo, muestra todos los mensajes (uno por línea)
//   - Con errors.Is / errors.As, "contiene" a CADA UNO de los
//     errores originales (no solo al primero)
//   - Ignora los nil que le pases (útil para joinear resultados
//     de validaciones que pueden o no fallar)

var (
	ErrEmailInvalido = errors.New("email inválido")
	ErrPasswordCorta = errors.New("la contraseña es muy corta")
	ErrEdadInvalida  = errors.New("la edad debe ser mayor a 0")
)

func validarEmail(email string) error {
	if !contieneArroba(email) {
		return ErrEmailInvalido
	}
	return nil
}

func validarPassword(pass string) error {
	if len(pass) < 8 {
		return ErrPasswordCorta
	}
	return nil
}

func validarEdad(edad int) error {
	if edad <= 0 {
		return ErrEdadInvalida
	}
	return nil
}

// validarRegistro junta los errores de las TRES validaciones.
// Si alguna devuelve nil (pasó), errors.Join simplemente la ignora.
func validarRegistro(email, password string, edad int) error {
	return errors.Join(
		validarEmail(email),
		validarPassword(password),
		validarEdad(edad),
	)
}

func main() {
	fmt.Println("=== errors.Join: reportar TODOS los errores juntos ===")

	err := validarRegistro("sin-arroba", "123", -5)
	if err != nil {
		fmt.Println("Errores encontrados:")
		fmt.Println(err)
	}

	// ─────────────────────────────────────────────────────────
	// errors.Is FUNCIONA CON CADA ERROR DEL JOIN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== errors.Is detecta cada error individualmente ===")
	fmt.Println("¿Incluye ErrEmailInvalido?", errors.Is(err, ErrEmailInvalido))
	fmt.Println("¿Incluye ErrPasswordCorta?", errors.Is(err, ErrPasswordCorta))
	fmt.Println("¿Incluye ErrEdadInvalida?", errors.Is(err, ErrEdadInvalida))

	// ─────────────────────────────────────────────────────────
	// SI TODO ESTÁ BIEN, errors.Join(nil, nil, nil) DEVUELVE nil
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Un registro válido no genera error ===")
	errOK := validarRegistro("mati@mail.com", "unaClaveSegura", 25)
	fmt.Println("¿Hay error?", errOK != nil)

	// ─────────────────────────────────────────────────────────
	// SOLO ALGUNOS CAMPOS FALLAN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Solo el email falla ===")
	errParcial := validarRegistro("sin-arroba", "unaClaveSegura", 25)
	fmt.Println(errParcial)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  errors.Join(e1, e2, e3)  → un solo error que combina varios")
	fmt.Println("  Ignora los nil           → útil con varias validaciones")
	fmt.Println("  errors.Is/As             → siguen funcionando sobre CADA error del join")
	fmt.Println("  Usalo para               → validar formularios completos, no cortar al primer error")
}

func contieneArroba(s string) bool {
	for _, r := range s {
		if r == '@' {
			return true
		}
	}
	return false
}
