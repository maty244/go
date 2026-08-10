# 13 — HTTP REST

## Objetivo

Construir servidores y clientes HTTP correctos usando la biblioteca estándar.

## Recorrido

HTTP → servidor → routing → JSON → body estricto → status → query/headers →
middleware → auth simulada → cliente → graceful shutdown → API integradora →
tests con `httptest`.

La autenticación fija sirve para estudiar middleware; no es seguridad de
producción y no pretende reemplazar gestión de identidades, claves o JWT bien
validado.

## Prerrequisito y reto

Requiere `12_concurrencia`. Resolvé `retos/` y después compará `soluciones/`.
