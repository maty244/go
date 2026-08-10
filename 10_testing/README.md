# 10 — Testing

## Objetivo

Usar tests para expresar comportamiento, proteger refactors y encontrar casos
que un ejemplo manual no revela.

## Recorrido

Tests básicos → tablas → subtests → cobertura de statements → errores y panics
→ integrador → helpers/cleanup → fuzzing y benchmarks.

El nombre del archivo puede ser cualquiera siempre que termine en `_test.go`.
La cobertura mide statements ejecutados; 100 % no demuestra ausencia de bugs.

## Prerrequisito y reto

Requiere `09_biblioteca_estandar`. Resolvé `retos/` y después compará
`soluciones/`.
