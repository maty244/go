# 12 — Concurrencia

## Objetivo

Coordinar trabajo simultáneo sin carreras, bloqueos ni goroutines abandonadas.

## Recorrido

Goroutines → `WaitGroup.Go` y compatibilidad `Add/Done` → channels → buffers →
`select` → Mutex/RWMutex/Once → race detector → worker pools → contexto →
integrador → `testing/synctest`.

Concurrencia no implica paralelismo. Una goroutine es relativamente liviana,
pero sigue consumiendo memoria, tiempo de scheduler y los recursos que abra.

## Prerrequisito y reto

Requiere `10_testing` y conviene haber terminado el CLI. Resolvé `retos/` y
compará `soluciones/`.
