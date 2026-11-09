# 04 — Structs, métodos y punteros

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Diseñar tipos propios y entender cuándo compartir o modificar un valor.

## Si los punteros todavía te cuestan

Empezá por [Punteros desde cero](./PUNTEROS_DESDE_CERO.md). Explica las
expresiones una por una, con diagramas de cajas y tablas del estado después
de cada asignación. Después leé los comentarios del código mientras lo ejecutás.

Para concentrarte en punteros, este orden evita mezclar demasiados temas:

| Lección | Pregunta que vas a responder |
| --- | --- |
| [06 — Operadores](./06_operadores_de_punteros/main.go) | ¿Qué son n, &n, p, *p y &p? ¿Qué cambia al asignar? |
| [08 — nil](./08_punteros_nil/main.go) | ¿Qué diferencia hay entre no tener destino y apuntar a cero? |
| [15 — Funciones](./15_punteros_en_funciones/main.go) | ¿Qué recibe la función? ¿Por qué cambiar p no cambia el p del llamador? |
| [09 — Opcionales y new](./09_valores_opcionales_y_new/main.go) | ¿Cómo distingo un dato ausente de un cero explícito? |
| [10 — Structs](./10_punteros_a_structs/main.go) | ¿Cómo comparten un destino dos structs? ¿Cómo se enlazan nodos? |
| [16 — Elementos y range](./16_punteros_a_elementos_y_range/main.go) | ¿Estoy apuntando a una copia o al elemento real? ¿Qué pasa con append? |
| [11 y 12 — Métodos](./12_receptores_y_method_sets/main.go) | ¿Qué copia recibe un receptor por valor o por puntero? |
| [13 — Elegir](./13_decidir_entre_valor_y_puntero/main.go) | ¿Qué comportamiento necesito y qué costo debo medir? |
| [14 — Copias](./14_copias_superficiales_y_profundas/main.go) | ¿Copié el struct o también sus datos mutables internos? |
| [07 — Integrador](./07_ejercicio_de_punteros/main.go) | ¿Puedo seguir dos cuentas compartidas entre varias funciones? |

Las carpetas mantienen su numeración para conservar las rutas existentes.
El integrador 07 puede retomarse después de las demostraciones pequeñas.
Las interfaces de la lección 12 son una ampliación para volver después de la
etapa `06_interfaces_y_errores`; primero entendé el receptor con el contador.

## Mapa de temas (`01` a `16`)

### Tipos y Structs (`01` a `05`)
- `01_tipos_definidos_y_aliases`: diferencia entre `type T subyacente` (tipo nuevo con sus propios métodos) y `type A = T` (alias intercambiable).
- `02_structs`: declaración, literales con nombres, métodos básicos, structs anidados, embedding (composición), comparación con `==` y `fmt.Stringer`.
- `03_structs_como_parametros`: funciones de solo lectura (`T`) vs funciones que mutan el struct original (`*T`).
- `04_todo_se_pasa_por_valor`: por qué en Go todo se pasa por valor (copia) y cómo se comportan `int`, `struct`, `slice`, `map` y `*T`.
- `05_modelo_de_memoria`: valores directos vs descriptores, direcciones con `%p`, stack/heap y escape analysis.

### Punteros, Métodos y Copias (`06` a `16`)
- `06_operadores_de_punteros`: `n`, `&n`, `p`, `*p`, `&p`, `new(T)` y `new(expr)` (Go 1.26+).
- `07_ejercicio_de_punteros`: sistema bancario con depósitos, retiros, transferencias con rollback y extractos.
- `08_punteros_nil`: zero value `nil`, cómo evitar `nil pointer dereference`, punteros en structs y variables de `range`.
- `09_valores_opcionales_y_new`: distinguir `zero value` de ausencia de dato (`nil`), APIs PATCH y `new(valor)` en Go 1.26+.
- `10_punteros_a_structs`: datos compartidos entre structs, estructuras recursivas (lista enlazada y árbol binario) y `[]T` vs `[]*T`.
- `11_metodos`: diferencia entre métodos y funciones, receptores por valor y por puntero en un carrito de compras.
- `12_receptores_y_method_sets`: cuándo elegir receptor `T` vs `*T`, encadenamiento inmutable, receptores `nil` y relación con interfaces (method sets).
- `13_decidir_entre_valor_y_puntero`: criterios reales de diseño y rendimiento para elegir `T` o `*T`.
- `14_copias_superficiales_y_profundas`: copia superficial vs clonado profundo (`slices.Clone`, `maps.Clone`, `new(*p)`).
- `15_punteros_en_funciones`: modificar `*p` vs reasignar `p` local, retornar `*T`, usar `**T` e intercambiar valores con aliasing.
- `16_punteros_a_elementos_y_range`: `&n` (copia del loop) vs `&slice[i]` (elemento real), realocación tras `append` y por qué `&mapa[k]` no compila.

## Recorrido

Tipos definidos y aliases → structs → parámetros → paso por valor → modelo de
memoria → operadores de punteros → nil → `new(expr)` → estructuras recursivas →
métodos → method sets → decisión entre valor y puntero → copias superficiales
y profundas → llamadas a funciones y reemplazo de punteros (`15`) → punteros
a elementos, variables de range y arrays después de append (`16`).

Stack y heap son decisiones de implementación del compilador. Un `new`, un `&`
o un retorno de puntero no son por sí solos una promesa de ubicación física.

## Prerrequisito y reto

Requiere `03_colecciones_y_strings`. Resolvé `retos/` y compará `soluciones/`.
