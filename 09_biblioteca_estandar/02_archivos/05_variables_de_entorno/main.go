package main

import (
	"fmt"
	"os"
)

// =========================================================
// VARIABLES DE ENTORNO: os.Getenv, os.LookupEnv, os.Setenv
// =========================================================
// Las variables de entorno son la forma ESTÁNDAR de configurar un
// programa sin tocar el código: la URL de una base de datos, una
// clave de API, si estamos en modo "desarrollo" o "producción".
// Cuando lleguemos a HTTP/ y BaseDatos/, la config de tu API
// (puerto, credenciales) va a venir de acá, NUNCA hardcodeada en
// el código (y mucho menos, un secreto commiteado en git).

func main() {
	// ─────────────────────────────────────────────────────────
	// os.Getenv: LEER UNA VARIABLE (string vacío si no existe)
	// ─────────────────────────────────────────────────────────
	fmt.Println("=== os.Getenv: leer una variable de entorno ===")

	// PATH casi seguro existe en cualquier sistema
	fmt.Println("PATH (primeros 50 caracteres):", primeros(os.Getenv("PATH"), 50))

	// Una variable que probablemente NO existe
	valor := os.Getenv("KIOSCO_API_KEY")
	fmt.Printf("KIOSCO_API_KEY: %q (vacío si no está definida)\n", valor)

	// ─────────────────────────────────────────────────────────
	// EL PROBLEMA DE Getenv: NO DISTINGUE "vacía" DE "no existe"
	// ─────────────────────────────────────────────────────────
	// Igual que vimos con 09_biblioteca_estandar/03_json/05_campos_opcionales_con_punteros (punteros) y maps: Getenv devuelve
	// "" tanto si la variable no existe, como si existe pero está
	// vacía. Si necesitás la diferencia, usá LookupEnv.

	fmt.Println("\n=== os.LookupEnv: distinguir 'no existe' de 'vacía' ===")

	valor, existe := os.LookupEnv("KIOSCO_API_KEY")
	fmt.Printf("Valor: %q | Existe: %v\n", valor, existe)

	// ─────────────────────────────────────────────────────────
	// os.Setenv: DEFINIR UNA VARIABLE (dura solo este proceso)
	// ─────────────────────────────────────────────────────────
	// Normalmente las variables de entorno se configuran DESDE
	// AFUERA del programa (en la terminal, en un archivo .env, en
	// la configuración del servidor). os.Setenv sirve sobre todo
	// para tests, o para dar un valor por defecto en desarrollo.

	fmt.Println("\n=== os.Setenv: definir una variable (solo en este proceso) ===")

	os.Setenv("KIOSCO_MODO", "desarrollo")
	fmt.Println("KIOSCO_MODO:", os.Getenv("KIOSCO_MODO"))

	// ─────────────────────────────────────────────────────────
	// PATRÓN COMÚN: config con valor por defecto
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Patrón: leer config con valor por defecto ===")

	puerto := obtenerConDefault("KIOSCO_PUERTO", "8080")
	fmt.Println("Puerto configurado:", puerto)

	// ─────────────────────────────────────────────────────────
	// CASO REAL: cargar config de una "API" al arrancar
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso real: config de arranque ===")

	os.Setenv("KIOSCO_DB_HOST", "localhost")
	os.Setenv("KIOSCO_DB_NOMBRE", "kiosco_db")

	config := cargarConfig()
	fmt.Printf("  %+v\n", config)

	// ─────────────────────────────────────────────────────────
	// SECRETOS: NUNCA HARDCODEAR NI COMMITEAR
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Regla de oro con secretos ===")
	fmt.Println("  MAL:  const APIKey = \"sk-abc123...\"  (queda en el código, en git)")
	fmt.Println("  BIEN: apiKey := os.Getenv(\"API_KEY\")  (vive fuera del código)")
	fmt.Println("  Los .env con secretos van en .gitignore, NUNCA se commitean")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  os.Getenv(\"X\")        → valor, o \"\" si no existe")
	fmt.Println("  os.LookupEnv(\"X\")     → (valor, existe) — distingue vacía de ausente")
	fmt.Println("  os.Setenv(\"X\", v)     → define una variable (solo este proceso)")
	fmt.Println("  Uso típico            → config de la app: puertos, DB, API keys")
	fmt.Println("  Secretos              → SIEMPRE por variable de entorno, nunca en el código")
}

func primeros(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func obtenerConDefault(clave, valorDefault string) string {
	if valor, existe := os.LookupEnv(clave); existe {
		return valor
	}
	return valorDefault
}

type Config struct {
	DBHost   string
	DBNombre string
	Modo     string
}

func cargarConfig() Config {
	return Config{
		DBHost:   obtenerConDefault("KIOSCO_DB_HOST", "localhost"),
		DBNombre: obtenerConDefault("KIOSCO_DB_NOMBRE", "default_db"),
		Modo:     obtenerConDefault("KIOSCO_MODO", "produccion"),
	}
}
