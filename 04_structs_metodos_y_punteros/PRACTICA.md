# Práctica — copiar valores y acceder a sus destinos

Comandos desde la raíz. Dibujá qué variable contiene cada dato o dirección.

Antes de los ejercicios, podés seguir [Punteros desde cero](./PUNTEROS_DESDE_CERO.md).
En cada caso respondé: tipo de la expresión, variable que cambia y resultado.

## 1. Separar tipo definido de alias

    go run ./04_structs_metodos_y_punteros/01_tipos_definidos_y_aliases

Metros y Segundos son tipos distintos. Descomentá distancia = tiempo:
debe fallar al compilar. Identificador e int64 son el mismo tipo, así que
asignar id a numero no requiere conversión.

**Comprobación:** ¿qué información perdés al convertir ambos valores a int64?

## 2. Observar una copia superficial

    go run ./04_structs_metodos_y_punteros/14_copias_superficiales_y_profundas

Las líneas relevantes son:

    original: Mati [cambiada] 99 2
    clon: Mati [Go] 1 1

La copia superficial reemplaza su propio Nombre, pero modifica los destinos
compartidos por el slice, map y puntero. Sustituí maps.Clone por una asignación:
el original y el clon compartirán sus puntos.

**Comprobación:** ¿por qué copiar un struct no garantiza inmutabilidad?
Si el map guardara slices, ¿maps.Clone bastaría?

## 3. Elegir un receptor

    go run ./04_structs_metodos_y_punteros/12_receptores_y_method_sets

Compará qué métodos tiene T y cuáles tiene *T. Que Go permita escribir
valor.MetodoConPuntero() para una variable direccionable no implica que T
implemente una interfaz que requiere ese método.

**Experimento:** cambiá un receptor que modifica un campo simple de *T a T.
El cambio del campo queda en la copia. ¿Por qué un map dentro de T puede seguir
cambiando aunque el receptor sea por valor?

## 4. Leer antes de escribir

    go run ./04_structs_metodos_y_punteros/06_operadores_de_punteros

En el bloque de copias esperá:

    tras *p=20: n=20 copia=10 *q=20
    tras copia=30: n=20 copia=30 *p=20
    tras reapuntar: n=20 otro=60 *p=60 *q=20

**Experimento:** cambiá q := p por q := &copia. Ahora *q lee copia, no n.
Predecí los tres resultados nuevos antes de ejecutar. Si cambiás p = &otro
por *p = otro, ¿se cambia la dirección guardada en p o se escribe un int?

**Comprobación:** ¿por qué copia := *p no crea un puntero? ¿Qué tipos tienen
p, *p y &p? ¿Por qué el mismo símbolo * no siempre significa lo mismo?

## 5. Separar nil de un destino con cero

    go run ./04_structs_metodos_y_punteros/08_punteros_nil

Probá este fragmento en una función main de práctica:

```go
var p *int
fmt.Println(p == nil) // true
p = new(int)
fmt.Println(p == nil, *p) // false 0
q := p
p = nil
fmt.Println(p == nil, *q) // true 0
```

**Comprobación:** al asignar p = nil, ¿qué variable dejó de guardar un destino?
¿Por qué q sigue siendo válido? ¿Qué pasaría al evaluar *p en la última línea?
No hace falta recover para manejar una ausencia que podés comprobar con if.

## 6. Copiar argumentos, modificar destinos, reemplazar punteros

    go run ./04_structs_metodos_y_punteros/15_punteros_en_funciones
    go test -v ./04_structs_metodos_y_punteros/15_punteros_en_funciones

El primer bloque deja n en 10 al pasar int y lo cambia a 11 al pasar &n.
Después de cambiarPunteroLocal, esperá:

    a=10 b=99 *p=10 p==&a=true

**Experimento:** en cambiarPunteroLocal, quitá p = destino. ¿Se modifica a
o b al ejecutar *p = 99? Explicá por qué el llamador no necesita **int para
modificar a, pero sí otro nivel si quiere que la función reemplace su p.

En reemplazarPuntero, distinguí:

- pp == nil: no hay variable puntero a la que acceder.
- *pp == nil: esa variable sí existe, pero aún no identifica un int.
- **pp: solo se puede evaluar si ambos niveles tienen destino.

**Comprobación:** ¿por qué reemplazarPuntero(&p, &b) puede funcionar cuando
p empieza nil? ¿Qué aporta retornar un puntero y asignarlo con p = crearContador(50)?

## 7. Intercambiar y pensar en aliases

En la lección 15, el intercambio de 3 y 8 debe dejar x=8 e y=3.
Si una entrada es nil, el bool es false y ambos ints quedan intactos.
Pasar &x dos veces devuelve true y conserva x: ambos parámetros tienen el
mismo destino, no dos ints independientes.

**Experimento:** sustituí la asignación múltiple por dos instrucciones:

```go
*a = *b
*b = *a
```

Para a=3 y b=8, las dos variables terminan en 8: la primera escritura perdió
el 3 original. Corregilo con una variable temporal y compará los dos diseños.

## 8. Apuntar al elemento o a una copia del loop

    go run ./04_structs_metodos_y_punteros/16_punteros_a_elementos_y_range
    go test -v ./04_structs_metodos_y_punteros/16_punteros_a_elementos_y_range

Cambiar *copias[0] deja el slice en [10 20 30]. Cambiar *elementos[0] lo deja
en [77 20 30]. La variable exterior reutilizada se lee como 30 desde todos
sus punteros; cambiarla mediante uno a 88 se observa desde los otros.

**Comprobación:** ¿qué declara := en for _, n := range numeros?
¿Qué reasigna = en for _, n = range numeros? ¿Alguna de esas dos n es
necesariamente la variable que ocupa numeros[i]?

## 9. Seguir un puntero después de append

Con make([]int, 1, 1), append obliga a usar otro array. El slice resultante
es [10 20], pero escribir mediante el puntero antiguo lo cambia a 99 en el
array anterior. Ese puntero no queda inválido ni se mueve al array nuevo.

**Experimento:** cambiá la capacidad a 2. Ahora append cabe en el mismo
array: *anterior = 99 sí cambia el slice, que termina en [99 20].

**Comprobación:** ¿por qué len y cap son diferentes? ¿El puntero identifica
una variable concreta o la idea de «posición cero del slice que use después»?

## 10. Volver al integrador sin saltarse las escrituras

    go run ./04_structs_metodos_y_punteros/07_ejercicio_de_punteros

Seguí cuentaAna desde nuevaCuenta hasta transferir. Al llamar mostrarCuenta
se pasa *cuentaAna, un struct copiado. Al llamar retirar se pasa cuentaAna,
un puntero copiado que permite cambiar el saldo del mismo struct.

**Comprobación:** si cambiás c.Saldo dentro de mostrarCuenta, ¿se ve afuera?
Si el depósito de una transferencia falla después de retirar, ¿qué hace la
reversión? Este ejemplo secuencial no garantiza atomicidad entre goroutines.
