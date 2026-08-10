package main

import (
	"errors"
	"fmt"
	"log/slog"
)

// =========================================================
// log/slog: LOGGING ESTRUCTURADO
// =========================================================
// Hasta ahora usaste fmt.Println para "avisar" qué está pasando
// en el programa. Para un programa real (sobre todo un servidor
// que corre 24/7), eso no alcanza: necesitás logs que se puedan
// FILTRAR, BUSCAR y PROCESAR automáticamente (por herramientas de
// monitoreo), no solo leer con los ojos.
//
// slog (parte de la librería estándar desde Go 1.21) resuelve
// esto con LOGGING ESTRUCTURADO: en vez de un mensaje de texto
// libre, cada log es un mensaje + una serie de pares clave-valor.
//
//   fmt.Println("Usuario", nombre, "compró", producto)       // MAL: texto libre
//   slog.Info("compra realizada", "usuario", nombre, "producto", producto) // BIEN

func main() {
	// ─────────────────────────────────────────────────────────
	// LOS 4 NIVELES BÁSICOS
	// ─────────────────────────────────────────────────────────
	// Cada nivel indica QUÉ TAN IMPORTANTE es el mensaje. De menor
	// a mayor severidad:
	//
	//   Debug → detalle fino, solo útil mientras desarrollás
	//   Info  → algo normal que pasó (un pedido se creó)
	//   Warn  → algo raro, pero el programa sigue funcionando
	//   Error → algo falló

	slog.Debug("iniciando validación de stock") // por defecto, Debug NO se imprime (ver tema 02)
	slog.Info("pedido creado", "id", 4521, "cliente", "Matias")
	slog.Warn("stock bajo", "producto", "Notebook", "restante", 2)
	slog.Error("pago rechazado", "pedido", 4521, "motivo", "fondos insuficientes")

	// ─────────────────────────────────────────────────────────
	// PARES CLAVE-VALOR: LA PARTE "ESTRUCTURADA"
	// ─────────────────────────────────────────────────────────
	// Después del mensaje, pasás pares "clave", valor, "clave",
	// valor... Cada par se muestra (y se puede buscar/filtrar)
	// como un campo separado, no mezclado en el texto.

	slog.Info("producto vendido",
		"nombre", "Mouse",
		"cantidad", 3,
		"precio", 8500.0,
		"total", 25500.0,
	)

	// ─────────────────────────────────────────────────────────
	// LOGGEAR UN ERROR: patrón típico
	// ─────────────────────────────────────────────────────────
	// Es MUY común pasar un valor error como uno de los pares.
	// slog sabe mostrarlo bien (usa Error() por dentro, como
	// vimos con fmt.Stringer en 06_interfaces_y_errores/04_fmt_stringer).

	err := errors.New("conexión a la base de datos perdida")
	slog.Error("no se pudo procesar el pedido", "pedido_id", 4521, "err", err)

	// ─────────────────────────────────────────────────────────
	// POR QUÉ ESTO ES MEJOR QUE fmt.Println EN UN PROGRAMA REAL
	// ─────────────────────────────────────────────────────────
	// Un sistema de monitoreo (o vos mismo, buscando en los logs)
	// puede filtrar "todos los Error de las últimas 2 horas donde
	// pedido_id=4521", algo IMPOSIBLE de hacer confiablemente con
	// texto libre armado a mano.

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  slog.Debug/Info/Warn/Error(msg, clave, valor, ...)")
	fmt.Println("  Cada par clave-valor queda como un CAMPO separado, no texto libre")
	fmt.Println("  Mucho más fácil de filtrar y buscar que fmt.Println")
	fmt.Println("  Por defecto imprime a stderr, en formato texto (tema 02: cómo cambiarlo)")
}
