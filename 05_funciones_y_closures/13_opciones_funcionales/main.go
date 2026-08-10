package main

import "fmt"

// =========================================================
// MUTABILIDAD CON PUNTEROS
// =========================================================
// En Go, por defecto los datos son INMUTABLES desde el punto
// de vista del receptor: cada función recibe una copia y no
// puede cambiar el original.
//
// MUTABILIDAD: capacidad de cambiar el estado de un dato.
//
//   SIN puntero → inmutable desde la función (trabaja con copia)
//   CON puntero → mutable desde la función (trabaja con original)
//
// Esta distinción es fundamental para diseñar APIs en Go.
// Elegir conscientemente entre "quiero que este dato pueda
// cambiar" vs "quiero garantizar que nadie lo toque" es
// una decisión de diseño, no un detalle técnico.

// ─────────────────────────────────────────────────────────
// EJEMPLO: Configuración de servidor
// ─────────────────────────────────────────────────────────

type Config struct {
	Host     string
	Puerto   int
	MaxConns int
	Debug    bool
}

// configInmutable recibe por VALOR → trabaja con una copia.
// No importa lo que haga adentro, el original no cambia.
// Útil cuando querés GARANTIZAR que la función no modifica config.
func configInmutable(c Config) {
	c.Host = "hackeado" // modifica la COPIA, no el original
	c.Debug = true
	fmt.Printf("  Dentro (copia):   Host=%s, Debug=%v\n", c.Host, c.Debug)
}

// configMutable recibe por PUNTERO → trabaja con el original.
// La función PUEDE y probablemente VA a modificar el dato.
func configMutable(c *Config) {
	c.Host = "nuevo-host.com" // modifica el ORIGINAL
	c.Debug = true
	fmt.Printf("  Dentro (puntero): Host=%s, Debug=%v\n", c.Host, c.Debug)
}

// ─────────────────────────────────────────────────────────
// PATRÓN: BUILDER CON MUTABILIDAD CONTROLADA
// ─────────────────────────────────────────────────────────
// Los builders suelen retornar *T para encadenar métodos
// y mutarlo paso a paso.

type ServidorBuilder struct {
	host     string
	puerto   int
	maxConns int
	tls      bool
}

func NuevoBuilder() *ServidorBuilder {
	// Valores por defecto
	return &ServidorBuilder{
		host:     "localhost",
		puerto:   8080,
		maxConns: 100,
	}
}

// Cada setter muta el builder y retorna él mismo → encadenamiento
func (b *ServidorBuilder) ConHost(h string) *ServidorBuilder {
	b.host = h
	return b
}

func (b *ServidorBuilder) ConPuerto(p int) *ServidorBuilder {
	b.puerto = p
	return b
}

func (b *ServidorBuilder) ConTLS() *ServidorBuilder {
	b.tls = true
	return b
}

func (b *ServidorBuilder) ConMaxConns(n int) *ServidorBuilder {
	b.maxConns = n
	return b
}

func (b *ServidorBuilder) Build() Config {
	host := b.host
	if b.tls {
		host = "https://" + host
	}
	return Config{
		Host:     host,
		Puerto:   b.puerto,
		MaxConns: b.maxConns,
	}
}

// ─────────────────────────────────────────────────────────
// PATRÓN: FUNCTIONAL OPTIONS (configuración inmutable elegante)
// ─────────────────────────────────────────────────────────
// Técnica avanzada y muy usada en Go para configurar structs
// sin exponer campos ni romper compatibilidad.

type Servidor struct {
	host   string
	puerto int
	debug  bool
}

// OpcionServidor es una función que muta el servidor.
// Quien define las opciones controla la mutabilidad.
type OpcionServidor func(*Servidor)

func ConHost(h string) OpcionServidor {
	return func(s *Servidor) {
		s.host = h
	}
}

func ConPuerto(p int) OpcionServidor {
	return func(s *Servidor) {
		s.puerto = p
	}
}

func ConDebug() OpcionServidor {
	return func(s *Servidor) {
		s.debug = true
	}
}

// NuevoServidor crea un servidor con opciones.
// El caller decide qué mutar; el constructor controla el resto.
func NuevoServidor(opciones ...OpcionServidor) *Servidor {
	s := &Servidor{
		host:   "localhost", // default
		puerto: 8080,        // default
	}
	for _, op := range opciones {
		op(s) // cada opción muta el servidor controladamente
	}
	return s
}

func (s *Servidor) String() string {
	debug := ""
	if s.debug {
		debug = " [DEBUG]"
	}
	return fmt.Sprintf("%s:%d%s", s.host, s.puerto, debug)
}

// ─────────────────────────────────────────────────────────
// INMUTABILIDAD DEFENSIVA: clonar antes de modificar
// ─────────────────────────────────────────────────────────

type Pedido struct {
	ID       int
	Producto string
	Cantidad int
	Precio   float64
}

// clonar retorna una COPIA del pedido, no un puntero al mismo.
// Así podemos modificar la copia sin tocar el original.
func clonar(p Pedido) Pedido {
	return p // Go copia el struct entero por valor
}

func aplicarDescuento(p Pedido, pct float64) Pedido {
	// Trabajamos sobre una copia → el original queda intacto
	p.Precio *= (1 - pct)
	return p
}

// ─────────────────────────────────────────────────────────
// ACUMULADOR MUTABLE (counter, estadísticas)
// ─────────────────────────────────────────────────────────

type Estadisticas struct {
	Ventas   int
	Total    float64
	Promedio float64
}

// registrarVenta muta las estadísticas directamente.
// Usar puntero aquí es correcto: el propósito ES acumular estado.
func registrarVenta(e *Estadisticas, monto float64) {
	e.Ventas++
	e.Total += monto
	e.Promedio = e.Total / float64(e.Ventas)
}

func main() {
	fmt.Println("╔══════════════════════════════════╗")
	fmt.Println("║         MUTABILIDAD               ║")
	fmt.Println("╚══════════════════════════════════╝")

	// ─────────────────────────────────────────────────────────
	// VALOR (inmutable) vs PUNTERO (mutable)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Valor vs Puntero: efecto en el original ===")

	cfg := Config{Host: "mi-server.com", Puerto: 443, Debug: false}

	fmt.Printf("Antes de configInmutable: Host=%s, Debug=%v\n", cfg.Host, cfg.Debug)
	configInmutable(cfg)
	fmt.Printf("Después de configInmutable: Host=%s, Debug=%v\n", cfg.Host, cfg.Debug)
	fmt.Println("  → Paso por valor: el original NO cambió")

	fmt.Println()
	fmt.Printf("Antes de configMutable: Host=%s, Debug=%v\n", cfg.Host, cfg.Debug)
	configMutable(&cfg)
	fmt.Printf("Después de configMutable: Host=%s, Debug=%v\n", cfg.Host, cfg.Debug)
	fmt.Println("  → Paso por puntero: el original SÍ cambió")

	// ─────────────────────────────────────────────────────────
	// BUILDER PATTERN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Builder pattern (mutabilidad encadenada) ===")

	config := NuevoBuilder().
		ConHost("api.miapp.com").
		ConPuerto(443).
		ConTLS().
		ConMaxConns(500).
		Build()

	fmt.Printf("Config construida: Host=%s, Puerto=%d, MaxConns=%d\n",
		config.Host, config.Puerto, config.MaxConns)

	// ─────────────────────────────────────────────────────────
	// FUNCTIONAL OPTIONS PATTERN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Functional options (mutabilidad controlada) ===")

	// Servidor con opciones seleccionadas
	srv1 := NuevoServidor(ConHost("prod.miapp.com"), ConPuerto(443))
	srv2 := NuevoServidor(ConDebug()) // solo debug, resto por default
	srv3 := NuevoServidor()           // todo por default

	fmt.Printf("srv1: %s\n", srv1)
	fmt.Printf("srv2: %s\n", srv2)
	fmt.Printf("srv3: %s\n", srv3)

	// ─────────────────────────────────────────────────────────
	// INMUTABILIDAD DEFENSIVA
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Inmutabilidad defensiva: clonar y modificar ===")

	original := Pedido{ID: 1, Producto: "Notebook", Cantidad: 1, Precio: 1500.00}
	conDescuento := aplicarDescuento(original, 0.10) // 10% descuento

	fmt.Printf("Original:     %s → $%.2f\n", original.Producto, original.Precio)
	fmt.Printf("Con 10%% dto:  %s → $%.2f\n", conDescuento.Producto, conDescuento.Precio)
	fmt.Println("  → El original NO se tocó (recibimos/retornamos por valor)")

	// ─────────────────────────────────────────────────────────
	// ACUMULADOR MUTABLE
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Acumulador mutable (estadísticas) ===")

	ventas := []float64{150.00, 89.99, 320.00, 45.50, 210.00}
	stats := &Estadisticas{} // puntero → vamos a mutarlo

	for _, v := range ventas {
		registrarVenta(stats, v)
	}

	fmt.Printf("Ventas: %d | Total: $%.2f | Promedio: $%.2f\n",
		stats.Ventas, stats.Total, stats.Promedio)

	// ─────────────────────────────────────────────────────────
	// TABLA DE DECISIÓN: ¿valor o puntero?
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== ¿Cuándo usar puntero para mutabilidad? ===")
	fmt.Println("  ✓ La función NECESITA modificar el dato             → *T")
	fmt.Println("  ✓ El struct es grande y copiar es caro              → *T")
	fmt.Println("  ✓ Acumular estado (contador, estadísticas)          → *T")
	fmt.Println("  ✓ Builder o configuración mutable encadenada        → *T")
	fmt.Println()
	fmt.Println("  ✓ Solo leer, no modificar                           → T")
	fmt.Println("  ✓ Garantizar que nadie cambia el dato               → T")
	fmt.Println("  ✓ Retornar datos transformados sin tocar el origen  → T")
	fmt.Println("  ✓ El struct es pequeño (int, bool, struct chico)    → T")

	fmt.Println("\n=== Regla de oro ===")
	fmt.Println("Si una función dice 'voy a cambiar este dato',")
	fmt.Println("el tipo debe decirlo también: recibí *T, no T.")
	fmt.Println("La firma de la función ES documentación.")
}
