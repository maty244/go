package main

import "fmt"

// =========================================================
// TIPOS DEFINIDOS VS TYPE ALIASES
// =========================================================
// En Go hay dos formas de darle nombre a un tipo:
//
// 1. TIPO DEFINIDO (sin "="):
//      type Metros int64
//    Crea un tipo NUEVO y distinto de int64 (su "tipo subyacente").
//    - Impide mezclar unidades por accidente (Metros ≠ Segundos ≠ int64).
//    - Permite definirle MÉTODOS propios (como String()).
//    - No hereda métodos de otros tipos definidos.
//
// 2. TYPE ALIAS (con "="):
//      type Identificador = int64
//    NO crea un tipo nuevo: es exactamente otro nombre para el mismo tipo.
//    - No necesita conversión para asignarse a int64.
//    - Útil para migraciones graduales de paquetes o abreviar tipos compuestos.

// Un tipo DEFINIDO crea un tipo nuevo, aunque comparta representación con el
// tipo subyacente. Esto permite impedir mezclas accidentales y agregar métodos.
type Metros int64
type Segundos int64

func (m Metros) String() string { return fmt.Sprintf("%d m", m) }

// Un ALIAS es solamente otro nombre para el mismo tipo. No crea una frontera.
type Identificador = int64

// Un alias también puede nombrar un tipo compuesto. Los aliases genéricos
// (Go 1.24+) se explican después, en 08_generics_e_iteradores/06_iteradores_y_go_126.
type Visitados = map[string]struct{}

func main() {
	distancia := Metros(120)
	tiempo := Segundos(10)

	// distancia = tiempo // no compila: Metros y Segundos son tipos distintos
	// var crudo int64 = distancia // tampoco compila sin conversión explícita: int64(distancia)
	velocidad := int64(distancia) / int64(tiempo)
	fmt.Println("Distancia:", distancia) // usa el método String() de Metros
	fmt.Println("Velocidad:", velocidad, "m/s")

	var id Identificador = 42
	var numero int64 = id // son el mismo tipo: no necesita conversión
	fmt.Printf("Alias: %d (tipo real de id: %T)\n", numero, id)

	visitados := Visitados{"/inicio": {}}
	_, existe := visitados["/inicio"]
	fmt.Println("Alias de map contiene /inicio:", existe)
}
