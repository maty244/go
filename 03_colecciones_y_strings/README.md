# 03 — Colecciones y strings

## Objetivo

Dominar cómo Go representa secuencias, tablas y texto UTF-8.

## Recorrido

Arrays y slices → operaciones y aliasing → maps y su semántica → bytes, runes y
UTF-8 → búsqueda, transformación, conversión y ordenamiento de strings.

Regla central: todos los valores se asignan y pasan por valor. Copiar un slice o
un map puede hacer que dos valores compartan almacenamiento interno; eso no crea
un mecanismo distinto llamado “paso por referencia”.

## Prerrequisito y reto

Requiere `02_control_de_flujo`. Resolvé `retos/` y después compará `soluciones/`.
