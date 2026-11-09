# 13 — HTTP REST (opcional)

## Objetivo

Construir servidores y clientes HTTP correctos usando la biblioteca estándar.

Esta etapa es una aplicación de Go a la web. La ruta principal sobre cómo
funciona el lenguaje termina en `12_concurrencia`; podés dejar HTTP para más
adelante sin perder los fundamentos. No necesitás crear una API para practicar
tipos, memoria, errores, tests o concurrencia.

## Recorrido

HTTP → servidor → routing → JSON → body estricto → status → query/headers →
middleware → auth simulada → cliente → graceful shutdown → API integradora →
tests con `httptest`.

La autenticación fija sirve para estudiar middleware; no es seguridad de
producción y no pretende reemplazar gestión de identidades, claves o JWT bien
validado.

## Prerrequisito y reto

Requiere `12_concurrencia`. Resolvé `retos/` y después compará `soluciones/`.
