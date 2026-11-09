# 05 — Funciones y closures

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Usar funciones como datos y construir comportamiento configurable sin perder
claridad.

## Recorrido

Variádicas → retornos nombrados → `defer` → funciones como valores y parámetros
→ anónimas → closures → captura de loops → fábricas → opciones funcionales →
recursividad.

Desde Go 1.22, las variables declaradas por el loop con `:=` son nuevas por
iteración. Con `=` sobre una variable exterior, la captura sigue compartida.
La lección 10 muestra ambos casos, además del patrón histórico.

## Mapa de temas (`01` a `15`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_variadicas` | Parámetros variádicos (`...T`) | Recibir 0 o más argumentos, expandir slices con `slice...` y diferencia con `max`/`min` built-in |
| `02_ejercicio_variadicas` | Ejercicio de variádicas | Estadísticas, filtros por rango y reportes flexibles |
| `03_retornos_nombrados` | Retornos nombrados y `_` | Variables de retorno pre-declaradas, *naked return* y cuándo usarlos o evitarlos |
| `04_defer` | `defer` a fondo | Orden LIFO, evaluación inmediata de argumentos, limpieza de recursos y modificar retornos nombrados |
| `05_funciones_como_valores` | Funciones de primera clase | Tipos de función (`type Op func(...)`), slices/maps de funciones, anónimas e IIFE |
| `06_funciones_como_parametros` | Higher-Order Functions (HOF) | Patrones `filtrar` (filter), `transformar` (map), `reducir` (reduce) y `forEach` |
| `07_funciones_anonimas` | Funciones anónimas | Funciones inline, IIFE, callbacks puntuales y `defer func() { ... }()` |
| `08_que_es_un_closure` | ¿Qué es un closure? | Captura de variables del entorno externo, estado propio y tiempo de vida extendido |
| `09_contador_con_estado` | Estado encapsulado con closures | Múltiples closures compartiendo la misma variable y structs con closures |
| `10_trampa_closure_en_for` | Closures dentro de `for` (Go 1.22+) | Cambio histórico en Go 1.22+ (variable nueva por iteración con `:=` vs compartida con `=`) |
| `11_funciones_que_devuelven_funciones` | Fábricas y decoradores | Aplicación parcial, funciones preconfiguradas, middlewares/decoradores y limitadores |
| `12_ejercicio_integrador` | Integrador: Kiosco configurable | Reglas de descuento con closures, caja registradora con estado y decorador de logging |
| `13_opciones_funcionales` | Mutabilidad, Builder y Functional Options | Paso por valor vs puntero, patrón Builder y patrón `OpcionServidor func(*Servidor)` |
| `14_recursividad` | Recursividad | Caso base, caso recursivo, stack de llamadas, acumuladores (sin TCO en Go) y árboles |
| `15_ejercicio_recursividad` | Ejercicio de recursividad | Formateo recursivo de moneda con separador de miles, descuentos y carritos |

## Prerrequisito y reto

Requiere `04_structs_metodos_y_punteros`. Resolvé `retos/` y luego mirá
`soluciones/`.
