# 02 — Control de flujo

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Controlar qué instrucciones se ejecutan y cuántas veces. Requiere funciones,
variables y operadores de `01_fundamentos`.

## Recorrido de lecciones

1. [`01_if`](./01_if/main.go) — `if`, `else if`, `else`, `if` con inicialización (`if v, err := ...; err != nil`), *guard clauses* y por qué Go no tiene operador ternario `?:`.
2. [`02_switch`](./02_switch/main.go) — `switch` con valor, múltiples valores por `case`, expresiones dinámicas, `fallthrough` explícito y *type switch*.
3. [`03_switch_sin_expresion`](./03_switch_sin_expresion/main.go) — `switch` en blanco (`switch { ... }`, equivalente a `switch true`) para reemplazar cadenas largas de `if/else if`.
4. [`04_for_clasico`](./04_for_clasico/main.go) — `for init; cond; post` clásico, pasos personalizados, múltiples variables y acceso a `i` e `i+1`.
5. [`05_for_con_condicion`](./05_for_con_condicion/main.go) — `for cond` (el `while` de Go): colas, búsquedas, reintentos y algoritmos de convergencia.
6. [`06_for_infinito`](./06_for_infinito/main.go) — `for { }` sin condición: menús, workers, monitoreo y cómo salir desde un `switch` interior.
7. [`07_break_y_continue`](./07_break_y_continue/main.go) — Diferencias entre `continue`, `break` y `return`, filtrado en pipelines y cuidado con `continue` en bucles sin `post`.
8. [`08_labels`](./08_labels/main.go) — Etiquetas (`break ETIQUETA` y `continue ETIQUETA`) para controlar bucles anidados y cuándo preferir extraer una función.
9. [`09_range`](./09_range/main.go) — `for range` sobre slices (copia vs índice), arrays, strings (UTF-8 runes), maps, channels y enteros (`for i := range N`, Go 1.22+).

El `for` clásico se conserva porque aparece en mucho código. Desde Go 1.22,
`range` también puede recorrer enteros (`for i := range 3`).

## Prerrequisito y reto

Requiere `01_fundamentos`. Resolvé `retos/` (`go test -tags reto ./02_control_de_flujo/retos`)
y después compará con `soluciones/`.
