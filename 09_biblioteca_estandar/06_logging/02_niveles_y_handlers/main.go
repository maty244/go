package main

import (
	"fmt"
	"log/slog"
	"os"
)

// =========================================================
// NIVELES Y HANDLERS: CONFIGURAR CÓMO SE VEN LOS LOGS
// =========================================================
// En el tema 01 usamos slog.Info/Warn/Error "a secas" (el logger
// por defecto). Acá vemos cómo CONFIGURAR el logging: en qué
// formato se muestra, y a partir de qué nivel se imprime.
//
// Un "Handler" es el encargado de decidir el FORMATO final del
// log. La librería estándar trae dos:
//
//   slog.NewTextHandler → texto plano, legible por humanos
//   slog.NewJSONHandler → JSON, ideal para que lo lea OTRO programa
//                          (herramientas de monitoreo, como Grafana/Loki)

func main() {
	fmt.Println("=== Handler de texto (el que ya viste, por defecto) ===")

	loggerTexto := slog.New(slog.NewTextHandler(os.Stdout, nil))
	loggerTexto.Info("pedido creado", "id", 4521, "total", 15000.50)
	loggerTexto.Warn("stock bajo", "producto", "Mouse")

	// ─────────────────────────────────────────────────────────
	// HANDLER JSON: ideal para producción
	// ─────────────────────────────────────────────────────────
	// En un servidor real, casi siempre vas a usar JSON: las
	// herramientas de monitoreo lo parsean automáticamente, cosa
	// que con texto libre es mucho más frágil.

	fmt.Println("\n=== Handler JSON ===")

	loggerJSON := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	loggerJSON.Info("pedido creado", "id", 4521, "total", 15000.50)
	loggerJSON.Error("pago rechazado", "pedido", 4521, "motivo", "fondos insuficientes")

	// ─────────────────────────────────────────────────────────
	// CONFIGURAR EL NIVEL MÍNIMO
	// ─────────────────────────────────────────────────────────
	// Por defecto, Debug NO se imprime (nivel mínimo: Info). Para
	// verlo (típicamente solo en desarrollo), bajás el nivel
	// mínimo con HandlerOptions.

	fmt.Println("\n=== Bajar el nivel mínimo a Debug ===")

	opciones := &slog.HandlerOptions{Level: slog.LevelDebug}
	loggerDebug := slog.New(slog.NewTextHandler(os.Stdout, opciones))

	loggerDebug.Debug("ahora SÍ se ve", "paso", "validación interna")
	loggerDebug.Info("esto también se ve, como siempre")

	// ─────────────────────────────────────────────────────────
	// SUBIR EL NIVEL: solo Warn y Error (menos ruido en producción)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Subir el nivel mínimo a Warn (menos ruido) ===")

	opcionesEstricto := &slog.HandlerOptions{Level: slog.LevelWarn}
	loggerEstricto := slog.New(slog.NewTextHandler(os.Stdout, opcionesEstricto))

	loggerEstricto.Info("este Info NO se va a imprimir")
	loggerEstricto.Warn("este Warn SÍ se imprime")
	loggerEstricto.Error("este Error también")

	// ─────────────────────────────────────────────────────────
	// slog.SetDefault: cambiar el logger por defecto de TODO el programa
	// ─────────────────────────────────────────────────────────
	// Si querés que TODOS los slog.Info/Warn/Error del programa
	// (sin usar una variable logger explícita) usen tu configuración,
	// se la asignás como default.

	fmt.Println("\n=== slog.SetDefault: cambiar el logger global ===")

	slog.SetDefault(loggerJSON)
	slog.Info("ahora el logger global también imprime en JSON")

	// ─────────────────────────────────────────────────────────
	// With: agregar contexto FIJO a un logger (sin repetirlo)
	// ─────────────────────────────────────────────────────────
	// Útil para, por ejemplo, que TODOS los logs de un módulo
	// lleven siempre el mismo campo (como "servicio": "pagos").

	fmt.Println("\n=== logger.With: contexto fijo reusable ===")

	loggerPagos := loggerTexto.With("servicio", "pagos")
	loggerPagos.Info("procesando cobro", "monto", 5000)
	loggerPagos.Error("timeout con la pasarela de pago")
	// Ambos logs incluyen "servicio=pagos" automáticamente

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  slog.NewTextHandler(w, opts)  → logs en texto plano")
	fmt.Println("  slog.NewJSONHandler(w, opts)  → logs en JSON (ideal para producción)")
	fmt.Println("  HandlerOptions{Level: X}      → nivel mínimo a partir del cual se imprime")
	fmt.Println("  slog.SetDefault(logger)       → cambia el logger global del programa")
	fmt.Println("  logger.With(clave, valor)     → logger reusable con contexto fijo")
}
