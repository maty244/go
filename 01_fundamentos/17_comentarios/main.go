package main

import "fmt"

// =========================================================
// COMENTARIOS EN GO
// =========================================================
// Los comentarios son texto que el compilador IGNORA completamente.
// Su propósito es explicar el código para humanos (vos en el futuro,
// o tus compañeros de equipo).
//
// Go tiene dos tipos de comentarios:
//   1. Comentario de línea: //
//   2. Comentario de bloque: /* */
//
// Buena práctica: comenta el POR QUÉ, no el QUÉ.
// El código ya dice QUÉ hace; el comentario debe explicar POR QUÉ.

// ─────────────────────────────────────────────────────────
// GODOC: La documentación oficial de Go usa comentarios especiales
// ─────────────────────────────────────────────────────────
// Los comentarios que empiezan justo antes de una función/tipo/variable
// sin línea en blanco son procesados por la herramienta "godoc"
// y se convierten en documentación oficial.
// Convención: el comentario empieza con el NOMBRE del símbolo.

// Saludar retorna un saludo personalizado para el nombre dado.
// Si nombre está vacío, retorna un saludo genérico.
func Saludar(nombre string) string {
	if nombre == "" {
		return "¡Hola, extraño!"
	}
	return "¡Hola, " + nombre + "!"
}

// MAX_RETRIES define el número máximo de reintentos para operaciones de red.
// Se eligió 3 porque las métricas de producción muestran que el 99.9% de
// los errores transitorios se resuelven en el primer o segundo reintento.
const MAX_RETRIES = 3

func main() {
	// ─────────────────────────────────────────────────────────
	// TIPO 1: Comentario de línea con //
	// ─────────────────────────────────────────────────────────

	// Esto es un comentario de una sola línea
	x := 42 // también puede ir al final de una línea de código

	fmt.Println("=== Comentarios de línea ===")
	fmt.Println("x:", x)

	// ─────────────────────────────────────────────────────────
	// TIPO 2: Comentario de bloque con /* */
	// ─────────────────────────────────────────────────────────
	/* Este es un comentario de bloque.
	   Puede ocupar múltiples líneas.
	   Se usa menos que // en el código Go moderno,
	   pero es útil para desactivar bloques de código temporalmente. */

	y := 100
	fmt.Println("y:", y)

	/*
		CÓDIGO TEMPORALMENTE DESACTIVADO:
		Esto puede ser útil durante el desarrollo para "comentar" código
		sin borrarlo, mientras experimentás con una alternativa.

		z := calcularAlgoComplejo()
		fmt.Println(z)
	*/

	// ─────────────────────────────────────────────────────────
	// BUENOS vs MALOS COMENTARIOS
	// ─────────────────────────────────────────────────────────

	// MAL COMENTARIO: dice lo mismo que el código (redundante)
	// suma a + b  ← el código ya dice eso!
	suma := x + y

	// BUEN COMENTARIO: explica el POR QUÉ o una decisión no obvia
	// Usamos suma entera porque los centavos ya están en la unidad mínima
	// y las operaciones con float causarían errores de redondeo en el total.
	fmt.Println("suma:", suma)

	// ─────────────────────────────────────────────────────────
	// COMENTARIO COMO SEPARADOR DE SECCIONES
	// ─────────────────────────────────────────────────────────
	// En funciones largas, los comentarios ayudan a separar secciones.
	// Es una práctica común en Go:

	// --- VALIDACIÓN ---
	nombre := "Matias"
	if nombre == "" {
		fmt.Println("Error: nombre vacío")
		return
	}

	// --- PROCESAMIENTO ---
	saludo := Saludar(nombre)

	// --- RESULTADO ---
	fmt.Println(saludo)

	// ─────────────────────────────────────────────────────────
	// COMENTARIOS TODO, FIXME, HACK, NOTE
	// ─────────────────────────────────────────────────────────
	// Es una convención (no oficial en Go pero universal) usar
	// estas palabras clave al inicio de comentarios para indicar
	// trabajo pendiente o advertencias. Los IDEs las resaltan en color.

	// TODO: agregar validación de email antes de procesar
	// FIXME: este cálculo falla cuando edad es 0, ver issue #42
	// HACK: workaround temporal hasta que el proveedor actualice su API
	// NOTE: esta función es llamada 1000 veces por segundo, no agregar logs aquí
	// OPTIMIZE: podría reemplazarse con un map para O(1) en lugar de O(n)

	email := "user@ejemplo.com"
	fmt.Println("Email (sin validar aún):", email) // TODO: validar

	// ─────────────────────────────────────────────────────────
	// DESACTIVAR CÓDIGO CON COMENTARIOS
	// ─────────────────────────────────────────────────────────
	// Útil en desarrollo para probar cosas sin borrar código.

	versiones := []string{"v1.0", "v1.1", "v2.0"}
	for _, v := range versiones {
		fmt.Println("Versión:", v)
		// fmt.Println("Debug info:", v) // comentado: solo para debugging
	}

	// ─────────────────────────────────────────────────────────
	// COMENTARIOS DE GODOC PARA TIPOS
	// ─────────────────────────────────────────────────────────
	// El comentario antes de un tipo aparece en la documentación.

	fmt.Println("\nUsando Saludar:", Saludar("Ana"))
	fmt.Println("MAX_RETRIES:", MAX_RETRIES)
}

// ─────────────────────────────────────────────────────────
// REGLAS DE ORO PARA COMENTARIOS EN GO:
// ─────────────────────────────────────────────────────────
//
// 1. Las funciones/tipos EXPORTADOS (con mayúscula) DEBEN tener comentario godoc.
//    Las no exportadas (minúscula) pueden o no tenerlo.
//
// 2. El comentario debe empezar con el nombre del símbolo.
//    BIEN:  // Saludar retorna un saludo...
//    MAL:   // Esta función retorna un saludo...
//
// 3. Preferí // sobre /* */ para la mayoría de los comentarios.
//
// 4. Comentá el POR QUÉ, no el QUÉ.
//    El código bien escrito ya es legible; el comentario agrega contexto.
//
// 5. Mantené los comentarios actualizados.
//    Un comentario desactualizado es PEOR que no tener comentario.
