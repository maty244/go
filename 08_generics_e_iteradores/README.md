# 08 — Generics e iteradores

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Reutilizar algoritmos conservando información de tipos.

## Recorrido

Motivación → funciones genéricas → constraints → tipos genéricos → ejercicio →
iteradores y novedades modernas → clonar slices con tipos genéricos
(`07_clonar_slices`: retoma el reto de colecciones y muestra el límite de la copia).

La biblioteca estándar `slices` no incluye `Map`, `Filter` ni `Reduce`. Esos
algoritmos se escriben con loops o helpers propios; `iter`, `slices` y `maps`
ofrecen piezas para componer secuencias cuando realmente mejora la claridad.

## Mapa de temas (`01` a `07`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_por_que_generics` | ¿Por qué Generics? | Evitar duplicar funciones por tipo sin perder seguridad estática como con `any` |
| `02_funciones_genericas` | Funciones genéricas | Parámetros de tipo `[T ...]`, inferencia automática de tipos y restricciones de unión |
| `03_constraints` | Constraints (`~T`, `comparable`, `cmp.Ordered`) | Interfaces de restricción, operador `~` para tipos subyacentes, `comparable` y `cmp.Ordered` |
| `04_tipos_genericos` | Tipos genéricos (`struct[T]`) | Estructuras de datos reutilizables (`Pila[T]`), métodos con receptor genérico y zero value de `T` |
| `05_ejercicio_integrador` | Integrador: `Filtrar`, `Mapear`, `Reducir` | Funciones con múltiples parámetros de tipo (`[T, R any]`, `[T any, K cmp.Ordered]`) |
| `06_iteradores_y_go_126` | Iteradores (`iter.Seq`) y Go 1.26 | Aliases genéricos (`type Lista[T any] = []T`), `iter.Seq`/`iter.Seq2`, `slices.Collect` y restricciones autorreferenciales (Go 1.26) |
| `07_clonar_slices` | Clonar slices genéricos | Preservar `nil` vs slice vacío `[]T{}` y límite de la copia superficial con tipos anidados |

## Prerrequisito y reto

Requiere `07_paquetes_modulos_y_herramientas`. Resolvé `retos/` y después
compará `soluciones/`.
