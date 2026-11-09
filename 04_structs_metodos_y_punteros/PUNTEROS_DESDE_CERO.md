# Punteros desde cero: una operación por vez

El objetivo es poder leer una línea y explicar **qué variable cambia**, sin
adivinar por los símbolos. Primero estudiá los ejemplos con ints; después pasá
a structs, slices y métodos. Todos los comandos se ejecutan desde la raíz.

Los fragmentos cortos con asignaciones se leen dentro del cuerpo de una función
main. Cuando aparece un tipo como Persona, se refiere al struct definido en las
lecciones de esta etapa. Las carpetas ejecutables ya incluyen package, imports,
tipos y main; podés correrlas completas sin reconstruir esas piezas.

## 1. Una variable tiene un tipo, un valor y una dirección

```go
n := 10
```

- `n` es el nombre de una variable.
- Su tipo es `int`, inferido del inicializador.
- Su valor actual es el número `10`.
- Podemos obtener un puntero a esa variable escribiendo `&n`.

Pensá en una caja llamada `n` que contiene `10`. La dirección identifica la
caja; no es el número que contiene. Usaremos las etiquetas A y P para dibujar
direcciones, sin pretender que sean las direcciones reales del proceso.

## 2. `&n` obtiene un puntero; no copia el int

```go
p := &n
```

Se evalúa la derecha primero: `&n` obtiene un puntero a `n`. Después se crea
**otra variable**, `p`, que guarda ese puntero. Su tipo es `*int`.

```text
variable p, ubicada en P          variable n, ubicada en A
┌───────────────────────┐        ┌────────────────────────┐
│ contenido: dirección A│ ─────→ │ contenido: 10          │
└───────────────────────┘        └────────────────────────┘
```

`p` y `n` son variables distintas. Cambiar el int de `n` no cambia el destino
de `p`: el puntero sigue permitiendo acceder a esa misma variable.

## 3. El asterisco tiene sentidos distintos según dónde aparece

| Código | Cómo leerlo | Qué hace |
| --- | --- | --- |
| `var p *int` | p es un puntero a int | Declara una variable, inicialmente nil |
| `p = &n` | guardar en p un puntero a n | Asigna el puntero, no el int |
| `copia := *p` | leer el int al que apunta p | Copia ese int a otra variable |
| `*p = 20` | escribir en el int al que apunta p | Modifica n si p apunta a n |
| `2 * 3` | dos por tres | Multiplica; no usa punteros |

En un **tipo**, `*int` significa «puntero a int». En una **expresión**, `*p`
significa «acceder a la variable apuntada». La asignación determina si leemos
esa variable a la derecha o escribimos en ella a la izquierda.

## 4. Distinguir `n`, `&n`, `p`, `*p` y `&p`

Con `n := 10` y `p := &n`:

| Expresión | Tipo | Resultado |
| --- | --- | --- |
| `n` | `int` | 10 |
| `&n` | `*int` | Puntero a n |
| `p` | `*int` | El mismo destino que &n |
| `*p` | `int` | 10, leído desde n |
| `&p` | `**int` | Puntero a la variable p |

`p` identifica a `n`; `&p` identifica a `p`. No son intercambiables: sus tipos
ya expresan la diferencia. `%d` muestra un int; `%p` muestra un puntero;
`%T` muestra el tipo. Las direcciones concretas pueden variar entre ejecuciones.

Ejecutá el primer bloque de la lección:

```bash
go run ./04_structs_metodos_y_punteros/06_operadores_de_punteros
```

## 5. Leer mediante un puntero no vuelve puntero al resultado

```go
n := 10
p := &n
copia := *p
*p = 20
```

`*p` se evalúa al ejecutar la declaración de `copia`: da el int `10`.
`copia` recibe ese int. **No queda vinculada a cambios futuros de n.**
Después de escribir `20` mediante p, `n` vale 20 y `copia` sigue valiendo 10.
En cambio, declarar `q := p` copia el puntero y permite llegar al mismo n.

## 6. Copiar un puntero y reapuntarlo son operaciones distintas

Este es el recorrido de `copiarValorYCompartirDestino` en la lección 06:

```go
n := 10
p := &n
copia := *p
q := p
*p = 20
copia = 30
otro := 50
p = &otro
*p = 60
```

| Después de ejecutar | n | copia | otro | Destino de p | Destino de q |
| --- | --- | --- | --- | --- | --- |
| `q := p` | 10 | 10 | No declarado | n | n |
| `*p = 20` | 20 | 10 | No declarado | n | n |
| `copia = 30` | 20 | 30 | No declarado | n | n |
| `otro := 50` | 20 | 30 | 50 | n | n |
| `p = &otro` | 20 | 30 | 50 | otro | n |
| `*p = 60` | 20 | 30 | 60 | otro | n |

La diferencia esencial es **`p = ...` cambia el puntero; `*p = ...` cambia su
destino actual**. Reasignar p no reasigna q: son variables distintas aunque
antes guardaran punteros iguales.

## 7. Igual valor no significa igual destino

```go
a, b := 7, 7
pa, pb := &a, &b
```

- `a == b` es true: ambos ints valen 7.
- `pa == pb` es false: son dos variables int diferentes.
- `*pa == *pb` es true: leemos sus contenidos y comparamos los ints.

Para saber si dos punteros `*int` identifican a la misma variable, compará los
punteros. Para comparar los ints, desreferenciá ambos, siempre que no sean nil.
Este ejemplo no pretende cubrir identidad de objetos de tamaño cero.

## 8. `nil`, cero y `new` son tres cosas diferentes

```go
var ausente *int // nil: no hay int al que acceder mediante este puntero
cero := new(int) // puntero no-nil a un int cuyo valor inicial es 0
valor := new(42) // Go 1.26+: puntero no-nil a un int inicializado en 42
```

`ausente == nil` es true. `cero == nil` es false y `*cero` es 0.
`nil` no significa «un int con cero»: `*ausente` produce panic.
Asignar `ausente = cero` es válido; asignar `*ausente = 0` cuando sigue nil no.

También distinguí estas dos construcciones:

```go
x := 42
a := &x
b := new(x)
x = 100
```

`*a` da 100 porque a identifica a x. `*b` da 42 porque new(x) creó **otro int**
con el valor que tenía x al evaluar el inicializador. No es un alias de x.
`&42` no es válido: el literal entero no es una variable direccionable.
Para structs sí existe el caso especial `&Persona{Nombre: "Ana"}`.

## 9. Una función también recibe el puntero por valor

```go
package main

import "fmt"

func incrementar(p *int) {
    *p = *p + 1
}

func main() {
    n := 10
    incrementar(&n)
    fmt.Println(n) // 11
}
```

1. Se evalúa `&n`: obtenemos un puntero a n.
2. El parámetro local p recibe una copia de ese puntero.
3. `*p` lee el 10 que hay en n.
4. Se suma 1 y se escribe 11 en la variable apuntada.
5. Termina la llamada. El parámetro p ya no se usa; n sigue valiendo 11.

Si el parámetro fuera `n int`, cambiar su int local no cambiaría el argumento.
No hay un modo secreto de «paso por referencia»: cambia **qué valor copiamos**.

```bash
go run ./04_structs_metodos_y_punteros/15_punteros_en_funciones
```

En su segundo ejemplo la función reasigna un p local a otro destino. El p del
llamador conserva su destino original. Leé los diagramas antes de ejecutar.

## 10. ¿Para qué sirve `**int`?

Con `p := &n`, `&p` tiene tipo `**int`: identifica la variable que guarda el
puntero. Dentro de una función que recibe `pp **int`:

- `pp` identifica la variable p del llamador.
- `*pp` accede a esa p; puede usarse para reemplazar su puntero.
- `**pp` accede al int final; requiere que ni pp ni *pp sean nil.

`*pp = &otro` cambia el destino de p. `**pp = 99` cambia el int de su destino
actual. No necesitás dos niveles solo para incrementar un int: basta `*int`.
Retornar el nuevo puntero y escribir `p = crearContador(50)` suele ser más claro.

## 11. Punteros a structs: el punto permite omitir un `*`

```go
persona := Persona{Nombre: "Ana", Edad: 28}
p := &persona
p.Edad = 29
```

La última línea equivale a `(*p).Edad = 29`. Los paréntesis agrupan la
desreferenciación antes del campo. **El acceso automático no evita nil**:
si p no tiene destino, acceder al campo igualmente produce panic.

`copia := *p` copia el struct. Si solo contiene ints y strings, reemplazar un
campo de copia no cambia persona. Si contiene slices, maps o punteros, sus
datos internos pueden seguir compartidos: estudiá la lección 14.

## 12. El receptor de un método también es un parámetro

Con `func (p Persona) CumplirAnios()`, p recibe una copia del struct. Cambiar
`p.Edad` cambia el campo de esa copia. Con `func (p *Persona) CumplirAnios()`,
p recibe una copia del puntero, que permite modificar a la persona original.

Para una variable direccionable, `persona.CumplirAnios()` puede abreviar
`(&persona).CumplirAnios()` si el método tiene receptor puntero. Esa comodidad
de llamada no cambia el tipo de persona ni sus method sets. En la lección 12
separa primero el ejemplo del contador; dejá las interfaces para cuando conozcas
la etapa 06 de interfaces y errores.

## 13. `range`: una copia no es el elemento del slice

```go
for _, n := range numeros {
    punteros = append(punteros, &n)
}
```

En Go 1.22+, cada n declarada por el loop es distinta, pero contiene una copia
del elemento. Sus punteros siguen válidos; escribir mediante ellos no cambia
el slice de ints. Para acceder a los elementos reales usá `&numeros[i]`.
Con una n declarada afuera y `for _, n = range numeros`, se reutiliza una
misma variable y los punteros sí comparten el destino.

```bash
go run ./04_structs_metodos_y_punteros/16_punteros_a_elementos_y_range
```

Ahí también vas a ver un puntero a un elemento anterior a append. Si append
crea otro array, el puntero sigue siendo válido, pero ya no identifica al
elemento del array nuevo. Un puntero no sigue a una posición lógica del slice.

## 14. Límites que evitan confusiones

- Go conserva los datos alcanzables: retornar `&variableLocal` es válido.
- Asignar nil a un puntero no pone en cero ni invalida su antiguo destino.
- No hay aritmética de punteros ordinarios: `p + 1` no compila. Se omite unsafe
  en esta ruta para aprender primero las reglas habituales del lenguaje.
- `&mapa[clave]` no compila. Podés copiar y reasignar un valor, o guardar `*T`
  como valor del map si querés compartir un objeto mutable.
- Un puntero no garantiza inmutabilidad, concurrencia segura ni mejor velocidad.
- El uso de `&` o new no obliga por sí solo a alojar un valor en el heap.
- Validá nil si es una entrada posible. También podés exigir un puntero no-nil
  y documentarlo; no todas las funciones necesitan comprobaciones defensivas.

## 15. Comprobar que entendiste

Sin ejecutar, escribí el resultado y justificá qué variable cambió:

```go
a, b := 1, 2
p := &a
q := p
p = &b
*q = 9
*p = 8
```

<details>
<summary>Ver explicación después de resolverlo</summary>

a termina en 9 y b en 8. q conserva el destino a. Reasignar p no modifica q;
la última escritura mediante p afecta a b. p == q es false.

</details>

Seguí con los ejercicios de [PRACTICA.md](./PRACTICA.md). Las salidas estables
están documentadas; no compares los números de direcciones entre ejecuciones.

Para consultar la regla exacta: [tipos puntero](https://go.dev/ref/spec#Pointer_types),
[operadores de dirección](https://go.dev/ref/spec#Address_operators) y
[range](https://go.dev/ref/spec#For_range).
