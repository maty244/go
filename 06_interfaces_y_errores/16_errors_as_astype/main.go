package main

import (
	"errors"
	"fmt"
)

// =========================================================
// errors.As / errors.AsType[T]: RECUPERAR EL TIPO CONCRETO
// =========================================================
// Ya viste la idea en 06_interfaces_y_errores/09_error_interface_a_fondo y 06_interfaces_y_errores/12_patron_de_errores. Acá la
// profundizamos: qué pasa cuando hay VARIOS tipos de error
// personalizados posibles, y wrapping de por medio.
//
//   errors.As(err, &destino)      → forma clásica (antes de 1.26)
//   errors.AsType[T](err)         → forma genérica (Go 1.26+),
//                                    devuelve (valor, ok)
//
// Ambas hacen lo mismo: recorren la cadena de errores (atravesando
// %w, como errors.Is) buscando uno que sea del tipo pedido.

type ErrorValidacion struct {
	Campo  string
	Motivo string
}

func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("campo %q inválido: %s", e.Campo, e.Motivo)
}

type ErrorPermisos struct {
	Usuario string
	Accion  string
}

func (e *ErrorPermisos) Error() string {
	return fmt.Sprintf("%s no puede hacer %q", e.Usuario, e.Accion)
}

func validarRegistro(email, usuario, accion string) error {
	if email == "" {
		return &ErrorValidacion{Campo: "email", Motivo: "no puede estar vacío"}
	}
	if usuario != "admin" && accion == "eliminar_producto" {
		return fmt.Errorf("validarRegistro: %w", &ErrorPermisos{Usuario: usuario, Accion: accion})
	}
	return nil
}

func main() {
	fmt.Println("=== errors.AsType[T] (Go 1.26+) ===")

	err := validarRegistro("", "mati", "crear_producto")
	if errValid, ok := errors.AsType[*ErrorValidacion](err); ok {
		fmt.Printf("  Validación falló en campo %q: %s\n", errValid.Campo, errValid.Motivo)
	}

	// ─────────────────────────────────────────────────────────
	// FUNCIONA A TRAVÉS DE WRAPPING
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== AsType atraviesa el wrapping ===")

	err = validarRegistro("mati@mail.com", "mati", "eliminar_producto")
	fmt.Println("Error:", err)

	if errPermisos, ok := errors.AsType[*ErrorPermisos](err); ok {
		fmt.Printf("  Permisos denegados: %s intentó %q\n", errPermisos.Usuario, errPermisos.Accion)
	}

	// ─────────────────────────────────────────────────────────
	// DISTINGUIR ENTRE VARIOS TIPOS POSIBLES
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Manejar varios tipos de error personalizado ===")

	casos := []struct{ email, usuario, accion string }{
		{"", "ana", "crear_producto"},
		{"ana@mail.com", "ana", "eliminar_producto"},
		{"ana@mail.com", "admin", "eliminar_producto"},
	}

	for _, c := range casos {
		manejar(validarRegistro(c.email, c.usuario, c.accion))
	}

	// ─────────────────────────────────────────────────────────
	// LA FORMA CLÁSICA: errors.As (por si ves código anterior a 1.26)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Forma clásica: errors.As ===")

	err = validarRegistro("", "mati", "crear_producto")
	var errValid *ErrorValidacion
	if errors.As(err, &errValid) {
		fmt.Printf("  (con errors.As) campo=%s motivo=%s\n", errValid.Campo, errValid.Motivo)
	}

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  errors.AsType[T](err) (ok)  → Go 1.26+, la forma más directa")
	fmt.Println("  errors.As(err, &destino)    → forma clásica, sigue funcionando")
	fmt.Println("  Ambas atraviesan wrapping   → igual que errors.Is")
	fmt.Println("  Usalas para                 → extraer DATOS del error, no solo saber que pasó")
}

func manejar(err error) {
	if err == nil {
		fmt.Println("  OK")
		return
	}
	if e, ok := errors.AsType[*ErrorValidacion](err); ok {
		fmt.Printf("  Validación: corregí el campo %q\n", e.Campo)
		return
	}
	if e, ok := errors.AsType[*ErrorPermisos](err); ok {
		fmt.Printf("  Permisos: %s no puede %q\n", e.Usuario, e.Accion)
		return
	}
	fmt.Println("  Error desconocido:", err)
}
