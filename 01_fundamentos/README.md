# 01 — Fundamentos

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Aprender a representar valores, operar con ellos y dividir un programa en
funciones pequeñas. Solo requiere haber completado `00_inicio`.

## Recorrido de lecciones

> **Nota sobre la numeración:** La lección `01_primer_programa` se encuentra en
> [`00_inicio`](../00_inicio/01_primer_programa/main.go); por eso esta carpeta
> arranca en `02_variables`.

1. [`02_variables`](./02_variables/main.go) — Declarar con `var`, inferencia de tipos, operador corto `:=`, reasignación y `_`.
2. [`03_constantes_e_iota`](./03_constantes_e_iota/main.go) — Constantes con `const`, bloques de constantes y enumeraciones con `iota`.
3. [`04_tipos_numericos`](./04_tipos_numericos/main.go) — Enteros con/sin signo (`int`, `uint8`, `int64`), `float32`/`float64`, complejos y overflow.
4. [`05_bool_string_rune`](./05_bool_string_rune/main.go) — Booleanos, strings inmutables, diferencia entre `byte` (`uint8`) y `rune` (`int32`, UTF-8).
5. [`06_zero_values`](./06_zero_values/main.go) — Valores por defecto de cada tipo en Go (`0`, `false`, `""`, `nil`) y por qué importan.
6. [`07_conversiones_de_tipo`](./07_conversiones_de_tipo/main.go) — Conversiones explícitas numéricas, paquete `strconv`, `[]byte` y `[]rune`.
7. [`08_constantes_tipadas`](./08_constantes_tipadas/main.go) — Diferencia entre constantes *untyped* (precisión arbitraria y flexibilidad) y *typed*.
8. [`09_var_vs_declaracion_corta`](./09_var_vs_declaracion_corta/main.go) — Cuándo elegir `var` y cuándo `:=`, redeclaración parcial y shadowing básico.
9. [`10_asignacion_multiple`](./10_asignacion_multiple/main.go) — Asignación en dos fases (swap), múltiples retornos y variables en `if`.
10. [`11_operadores`](./11_operadores/main.go) — Aritméticos, comparación, lógicos (cortocircuito), bit a bit (`&`, `|`, `^`, `&^`, `<<`, `>>`) y precedencia.
11. [`12_entrada_y_salida_fmt`](./12_entrada_y_salida_fmt/main.go) — Verbos de `fmt.Printf`, `Sprintf`, `Errorf`, `Fprintln` y lectura con `Scan` / `bufio.Scanner`.
12. [`13_primera_funcion`](./13_primera_funcion/main.go) — Anatomía de `func`, parámetros (paso por valor/copia) y retorno simple.
13. [`14_declarar_y_llamar_funciones`](./14_declarar_y_llamar_funciones/main.go) — Visibilidad (`PascalCase` exportado vs `camelCase` privado), parámetros abreviados y composición.
14. [`15_retornos_multiples`](./15_retornos_multiples/main.go) — Devolver varios valores, el patrón idiomático `(resultado, error)` y retornos nombrados.
15. [`16_scope_y_shadowing`](./16_scope_y_shadowing/main.go) — Ámbito de bloques `{}`, trampas clásicas de *shadowing* con `:=` vs `=` y cómo detectarlas.
16. [`17_comentarios`](./17_comentarios/main.go) — Comentarios de línea/bloque, convenciones de `godoc` y buenas prácticas.
17. [`18_errores_comunes`](./18_errores_comunes/main.go) — Los 12 errores típicos al empezar en Go (compilación y runtime) y cómo solucionarlos.
18. [`19_dinero_con_enteros`](./19_dinero_con_enteros/main.go) — Por qué usar enteros (`type Dinero int64` en centavos) en lugar de `float64` para dinero.

La lección `19_dinero_con_enteros` compara dos formas de representar importes:
`type Dinero int64` representa centavos; `float64` permite estudiar aritmética
de punto flotante, pero no es un modelo de cobro exacto. Algunos ejemplos usan
precios flotantes solo para practicar el lenguaje. Los enteros también tienen
límites: una multiplicación muy grande puede desbordar.

## Reto

Resolvé `retos/`, corré allí `go test -tags reto .` y recién después compará
con `soluciones/`.
