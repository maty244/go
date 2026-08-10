# 06 — Interfaces y errores

## Objetivo

Desacoplar comportamiento con contratos pequeños y conservar información útil
cuando algo falla.

## Recorrido

Interfaces implícitas → composición → `Stringer` → assertions y type switch →
`any` → nil en interfaces → `error` → sentinelas → wrapping → `errors.Is` →
`errors.AsType` → errores con campos → `errors.Join` → `panic/recover`.

Preferí aceptar interfaces y devolver tipos concretos. Definí la interfaz cerca
del código que la consume, no automáticamente junto a la implementación.

## Prerrequisito y reto

Requiere `05_funciones_y_closures`. Resolvé `retos/` y compará `soluciones/`.
