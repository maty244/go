# 📖 Gocito — ruta completa para dominar Go 1.26

Gocito es una biblioteca personal de estudio en español. La meta no es
memorizar todo el lenguaje: es poder leer, diseñar, probar y mantener programas
Go con confianza, y llegar a servidores HTTP sin arrastrar dudas de la base.

El repositorio usa **Go 1.26** y solamente la biblioteca estándar. Cada carpeta
es una etapa y cada subcarpeta numerada es una lección ejecutable.

## Requisitos

```bash
go version
```

Necesitás Go 1.26 o una versión posterior compatible. El módulo declara
`go 1.26.0`: esa es la versión mínima del lenguaje, no una obligación de usar
exactamente el parche `.0`.

## Cómo estudiar una lección

Usá siempre este ciclo:

1. **Leé** el objetivo y el código de arriba abajo.
2. **Predecí** la salida antes de ejecutarlo.
3. **Ejecutá** la carpeta completa.
4. **Modificá** valores y provocá errores deliberadamente.
5. **Resolvé** el reto de la etapa sin abrir la solución.
6. **Compará** tu diseño con `soluciones/`, no solo el resultado final.

Ejemplos:

```bash
go run ./00_inicio/01_primer_programa
go run ./03_colecciones_y_strings/08_utf8_y_runes
go test ./10_testing/02_table_driven_tests
go test -race ./...
```

Los tests de un starter se activan a propósito con la etiqueta `reto`:

```bash
go test -tags reto ./01_fundamentos/retos
```

## Ruta de estudio

Seguí las etapas en orden. Dentro de cada etapa, seguí la numeración de sus
lecciones.

- [ ] [00 — Inicio](./00_inicio/README.md): entorno, anatomía y método de estudio.
- [ ] [01 — Fundamentos](./01_fundamentos/README.md): valores, tipos, funciones básicas y errores iniciales.
- [ ] [02 — Control de flujo](./02_control_de_flujo/README.md): decisiones y loops.
- [ ] [03 — Colecciones y strings](./03_colecciones_y_strings/README.md): arrays, slices, maps y UTF-8.
- [ ] [04 — Structs, métodos y punteros](./04_structs_metodos_y_punteros/README.md): modelo de valores y diseño de tipos.
- [ ] [05 — Funciones y closures](./05_funciones_y_closures/README.md): funciones de orden superior, estado y `defer`.
- [ ] [06 — Interfaces y errores](./06_interfaces_y_errores/README.md): contratos pequeños y manejo completo de errores.
- [ ] [07 — Paquetes, módulos y herramientas](./07_paquetes_modulos_y_herramientas/README.md): organizar proyectos y dominar el comando `go`.
- [ ] [08 — Generics e iteradores](./08_generics_e_iteradores/README.md): reutilización con tipos y APIs modernas.
- [ ] [09 — Biblioteca estándar](./09_biblioteca_estandar/README.md): IO, archivos, JSON, tiempo, contexto y logging.
- [ ] [10 — Testing](./10_testing/README.md): tests, fuzzing, benchmarks, coverage y race detector.
- [ ] [11 — Proyecto CLI](./11_proyecto_cli/README.md): gestor de tareas persistente y probado.
- [ ] [12 — Concurrencia](./12_concurrencia/README.md): goroutines, sincronización, cancelación y leaks.
- [ ] [13 — HTTP REST](./13_http_rest/README.md): servidores, clientes, middleware y una API completa.

## Qué significa “dominar” esta ruta

Al terminar deberías poder:

- explicar cuándo se copia un valor y cuándo dos valores comparten datos;
- modelar dominio con tipos, structs, métodos e interfaces pequeñas;
- propagar y examinar errores sin perder contexto;
- organizar un módulo con paquetes y límites claros;
- escribir tests útiles, benchmarks y fuzz tests;
- coordinar goroutines sin carreras ni fugas;
- construir y apagar correctamente un servidor HTTP con la biblioteca estándar;
- detectar qué parte todavía no sabés y encontrarla en la documentación oficial.

## Comandos de salud del repositorio

```bash
gofmt -w .
go vet ./...
go test ./...
go test -race ./...
go fix -diff ./...
```

`go fix -diff` puede sugerir modernizaciones dentro de lecciones que conservan
una forma histórica a propósito. En esos casos la lección explica ambas formas.

## Después de esta base

La siguiente ruta, deliberadamente fuera de esta revisión, es:

1. PostgreSQL con `database/sql` y un driver.
2. Autenticación real, autorización y gestión de secretos.
3. Estructura de servicios, observabilidad y despliegue.
4. Protocol Buffers y gRPC.

Primero dominá la base. Una librería o framework se aprende mucho más rápido
cuando el lenguaje debajo ya no es una incógnita.

## Fuentes oficiales

- [Especificación de Go](https://go.dev/ref/spec)
- [Documentación de la biblioteca estándar](https://pkg.go.dev/std)
- [Notas de Go 1.22](https://go.dev/doc/go1.22)
- [Notas de Go 1.23](https://go.dev/doc/go1.23)
- [Notas de Go 1.25](https://go.dev/doc/go1.25)
- [Notas de Go 1.26](https://go.dev/doc/go1.26)
