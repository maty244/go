package main

import (
	"fmt"
	"sync"
)

// =========================================================
// EL RACE DETECTOR: go run -race / go test -race
// =========================================================
// En el tema 06 vimos QUÉ es una condición de carrera. El
// problema es que a veces el código CON el bug funciona bien en
// tus pruebas igual (por suerte, según cómo el sistema operativo
// decida intercalar las goroutines) y falla justo en producción.
//
// Go trae una herramienta que detecta carreras durante una ejecución
// instrumentada: la flag -race. Solo puede observar rutas y entrelazados que
// realmente se ejecutan; una corrida limpia no demuestra que sea imposible
// alcanzar otra carrera.
//
//   go run -race ./ruta/
//   go test -race ./ruta/
//   go build -race -o binario ./ruta/   (para producción, ocasional)
//
// -race hace que el programa corra más lento (por eso no se usa
// SIEMPRE en producción), pero para desarrollo y CI es
// prácticamente obligatorio activarlo.

// SaldoInseguro tiene la MISMA condición de carrera que el tema
// 06, a propósito, para que la detectes vos mismo con -race.
type Dinero int64 // centavos: los ejemplos de negocio no usan float64

type SaldoInseguro struct{ monto Dinero }

func (s *SaldoInseguro) Depositar(cantidad Dinero) {
	s.monto += cantidad // ← acá está la carrera: leer + sumar + guardar
}

func main() {
	fmt.Println("=== Código con una condición de carrera (a propósito) ===")

	saldo := &SaldoInseguro{}
	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			saldo.Depositar(10)
		})
	}
	wg.Wait()

	fmt.Println("Esperado: 1000 centavos | Obtenido:", saldo.monto)

	fmt.Println("\n=== Cómo detectar la carrera con -race ===")
	fmt.Println("  go run -race ./12_concurrencia/07_race_detector/")
	fmt.Println()
	fmt.Println("  El programa va a imprimir un reporte como:")
	fmt.Println("    WARNING: DATA RACE")
	fmt.Println("    Write at 0x... by goroutine N:")
	fmt.Println("      main.(*SaldoInseguro).Depositar()")
	fmt.Println("    Previous write at 0x... by goroutine M:")
	fmt.Println("      main.(*SaldoInseguro).Depositar()")
	fmt.Println()
	fmt.Println("  Te dice EXACTAMENTE en qué línea y con qué goroutines pasó,")
	fmt.Println("  aunque el resultado final (sin -race) hubiera parecido razonable.")

	// ─────────────────────────────────────────────────────────
	// -race EN TESTS: aún más importante
	// ─────────────────────────────────────────────────────────
	// La forma más común de usar -race en el día a día es junto
	// con go test, para que tu CI (integración continua) lo
	// chequee automáticamente en cada cambio.

	fmt.Println("\n=== -race también funciona con go test ===")
	fmt.Println("  go test -race ./...")
	fmt.Println("  (revisá 10_testing/ para repasar cómo se escriben los tests)")

	// ─────────────────────────────────────────────────────────
	// LA SOLUCIÓN: Mutex (o rediseñar para evitar el dato compartido)
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== La solución: Mutex (tema 06) ===")
	fmt.Println("  Agregar sync.Mutex + Lock()/Unlock() alrededor de 's.monto += cantidad'")
	fmt.Println("  elimina la carrera. Correlo con -race después de arreglarlo:")
	fmt.Println("  el warning debería desaparecer.")

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  go run -race / go test -race  → detecta carreras que la ejecución alcanza")
	fmt.Println("  Más lento                      → por eso no se usa siempre en producción")
	fmt.Println("  Sí siempre                     → en desarrollo y en CI (tests automáticos)")
	fmt.Println("  'Funcionó en mi máquina'       → no prueba que no haya una carrera")
	fmt.Println("  -race limpio                   → aumenta confianza, no es una prueba absoluta")
}
