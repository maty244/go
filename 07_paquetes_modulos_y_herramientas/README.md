# 07 — Paquetes, módulos y herramientas

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Pasar de archivos aislados a un módulo mantenible.

## Recorrido

Paquetes → carpetas → visibilidad → `go.mod`/`go.sum` → varios archivos →
ejercicio integrador → herramientas de Go 1.26 → compilador y runtime
(`08_compilador_y_runtime`: inicialización, vida de variables y escape analysis).

El módulo se llama `gocito` y exige como mínimo Go 1.26.0. No tiene
dependencias externas, por eso no necesita `go.sum`.

## Mapa de temas (`01` a `08`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_que_es_un_paquete` | ¿Qué es un paquete? | `package main` (ejecutable) vs paquetes de librería e `import` |
| `02_organizar_en_carpetas` | Organizar en carpetas | Crear paquetes propios, rutas de import (`modulo/ruta`) y prohibición de ciclos |
| `03_visibilidad` | Visibilidad (exportado vs privado) | Regla de la primera letra (Mayúscula vs minúscula) y protección de invariantes |
| `04_go_mod_a_fondo` | `go.mod` y `go.sum` a fondo | Directivas `module`, `go`, `require`, `go mod tidy` y rol de `go.sum` |
| `05_multiples_archivos_mismo_paquete` | Múltiples archivos en un paquete | Dividir un mismo paquete en varios archivos `.go` sin `import` entre ellos |
| `06_ejercicio_integrador` | Integrador: Kiosco multi-paquete | Proyecto separado en `main`, `productos` y `clientes` con estado encapsulado |
| `07_herramientas_go_126` | Herramientas de Go 1.26 | `go fmt`, `go vet`, `go test -race`, modernizadores con `go fix -diff` y `go tool` |
| `08_compilador_y_runtime` | Compilador y runtime | Orden de inicialización (`var` → `init()` → `main()`), escape analysis, `GOMAXPROCS` y GC |

## Prerrequisito y reto

Requiere `06_interfaces_y_errores`. Resolvé `retos/` y compará `soluciones/`.
