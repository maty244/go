# 11 — Proyecto CLI: gestor de tareas

## Objetivo

Juntar lenguaje, paquetes, IO, JSON y tests en un programa pequeño pero real.

Prerrequisito: completar `10_testing`.

## Uso

```bash
go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/tareas.json agregar "Estudiar slices"
go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/tareas.json listar
go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/tareas.json completar 1
```

La escritura usa un archivo temporal y `os.Rename` para no dejar el JSON a
medias si la operación falla.

## Reto

Implementá la búsqueda de `retos/`, corré `go test -tags reto .` allí y después
compará con `soluciones/`.
