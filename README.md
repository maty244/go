# 📖 Gocito — mi biblia de Go

Biblioteca personal de estudio de Go, desde cero hasta APIs REST (con la
mira puesta en dominar el lenguaje para hacer APIs y programas reales).
Cada carpeta es **un tema**: un `main.go` ejecutable, con comentarios en español
que explican qué hace cada cosa, ejemplos prácticos y un resumen al final.
Los temas de `Testing/` además incluyen un `main_test.go` con tests reales.

## Cómo ejecutar cualquier tema

Desde la raíz del proyecto:

```bash
go run ./Fundamentos/01_variables
go run ./Strings/05_concatenar_y_builder
go run ./Punteros/08_receptores_metodos
```

La idea es **leer el código de arriba a abajo** como una lección, y después
ejecutarlo para ver la salida. Jugá: cambiá valores, descomentá los ejemplos
que rompen, mirá los errores del compilador.

## Orden de estudio sugerido

1. **Fundamentos** (01 → 45): la base completa del lenguaje.
2. **Strings** (01 → 10): se puede intercalar después de Fundamentos 17 (maps),
   porque usa slices, maps y manejo de errores.
3. **Punteros** (01 → 12): requiere structs y métodos.
4. **Interfaces** (01 → 10): el siguiente concepto grande después de punteros.
5. **Closures** (01 → 06) y **Paquetes** (01 → 06): se pueden estudiar en
   cualquier orden entre sí.
6. **ErroresAvanzados** (01 → 08) y **Generics** (01 → 05): cierran el lenguaje.
7. **JSON** (01 → 07), **Tiempo** (01 → 06), **Archivos** (01 → 06): la
   biblioteca estándar que usa cualquier programa real.
8. **Contexto** (01 → 04) y **Logging** (01 → 03): antes de concurrencia.
9. **Testing** (01 → 06): aprendé a testear antes de meterte en concurrencia.
10. **Concurrencia** (01 → 10): el bloque más avanzado del lenguaje, se apoya en TODO lo anterior.
11. **APIRest** (01 → 12): el objetivo final — acá se junta todo. Cada tema
    levanta su propio servidor y se hace requests a sí mismo: `go run` y ves
    la película completa.

---

## 🧱 Fundamentos (45 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_variables` | Qué es una variable, declaración con `var`, reasignación |
| 02 | `02_constantes` | `const`, qué se puede y qué no, `iota` |
| 03 | `03_tipos_primitivos_1` | Enteros: `int`, `int8/16/32/64`, `uint` y sus rangos |
| 04 | `04_tipos_primitivos_2` | `float32/64`, `bool`, `string`, `rune` |
| 05 | `05_zero_values` | El valor por defecto de cada tipo |
| 06 | `06_conversiones_tipo` | Conversión explícita entre tipos numéricos: `float64(x)` |
| 07 | `07_constantes_tipadas` | Constantes tipadas vs sin tipo |
| 08 | `08_var_vs_walrus` | `var` vs `:=` — cuándo usar cada una |
| 09 | `09_operadores` | Aritméticos, comparación, lógicos, bitwise |
| 10 | `10_entrada_salida` | `fmt.Print/Println/Printf/Sprintf`, verbos de formato, lectura de consola |
| 11 | `11_errores_comunes` | Los errores típicos del principiante y cómo evitarlos |
| 12 | `12_multiples_variables` | Declaración y asignación múltiple, intercambio de valores |
| 13 | `13_comentarios` | `//`, `/* */` y convenciones |
| 14 | `14_control_flujo` | Panorama general del control de flujo |
| 15 | `15_funciones` | Introducción a funciones |
| 16 | `16_arrays_slices` | Arrays vs slices, `make`, `append`, `copy`, slicing, ordenar |
| 17 | `17_maps` | Crear, leer, escribir, borrar, chequear existencia |
| 18 | `18_structs` | Definir structs, campos, structs anidados |
| 19 | `19_condicional_if` | `if`, `else`, `if` con inicialización |
| 20 | `20_variables_sombra` | Shadowing: cuando una variable "tapa" a otra |
| 21 | `21_for_completo` | El `for` clásico de 3 partes |
| 22 | `22_for_condicion` | `for` con solo condición (el "while" de Go) |
| 23 | `23_for_infinito` | `for` infinito y cómo salir |
| 24 | `24_break_continue` | Cortar o saltear iteraciones |
| 25 | `25_for_range` | `range` sobre slices, strings, maps |
| 26 | `26_iterando_maps` | El orden aleatorio de los maps y cómo ordenarlos |
| 27 | `27_for_labels` | Labels para salir de loops anidados |
| 28 | `28_switch` | `switch` clásico, múltiples casos, `fallthrough` |
| 29 | `29_switch_blanco` | `switch` sin expresión (reemplazo de if/else if) |
| 30 | `30_primer_funcion` | Tu primera función paso a paso |
| 31 | `31_declarando_llamando` | Declarar y llamar funciones, orden en el archivo |
| 32 | `32_agregar_eliminar` | Agregar y eliminar elementos de slices |
| 33 | `33_parametros_struct` | Pasar structs como parámetros |
| 34 | `34_valor_referencia` | Semántica de valor vs referencia (antesala de Punteros) |
| 35 | `35_variadicos` | Funciones variádicas: `func(nums ...int)` |
| 36 | `36_cupones_variadic` | Ejercicio práctico con variádicas |
| 37 | `37_multiples_valores` | Funciones que retornan varios valores |
| 38 | `38_manejo_errores` | El patrón `(valor, error)`, crear y chequear errores |
| 39 | `39_retorno_nombrado` | Retornos con nombre |
| 40 | `40_defer` | `defer`: posponer ejecución hasta el final |
| 41 | `41_funciones_valores` | Funciones como valores (asignarlas a variables) |
| 42 | `42_funciones_parametros` | Funciones que reciben funciones |
| 43 | `43_metodos` | Métodos sobre tipos propios |
| 44 | `44_recursividad` | Funciones que se llaman a sí mismas |
| 45 | `45_descuento_recursivo` | Ejercicio integrador con recursión |

## 🔤 Strings (10 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_un_string` | Bytes inmutables, `len()` cuenta bytes, escapes, raw strings, `[]byte` |
| 02 | `02_utf8_y_runes` | Unicode, UTF-8, el tipo `rune`, bytes vs letras |
| 03 | `03_indexar_e_iterar` | `s[i]` da bytes, `for range` da runes, `[]rune` para acceso por letra |
| 04 | `04_inmutabilidad` | Por qué no se pueden modificar y cómo "modificarlos" bien |
| 05 | `05_concatenar_y_builder` | `+`, `Sprintf`, `Join` y `strings.Builder` para loops |
| 06 | `06_paquete_strings_buscar` | `Contains`, `HasPrefix/Suffix`, `Index`, `Count`, `EqualFold` |
| 07 | `07_paquete_strings_transformar` | `ToUpper/ToLower`, `Split`, `Fields`, `Trim*`, `Replace*` |
| 08 | `08_strconv_conversiones` | `Atoi`, `Itoa`, `ParseFloat/Bool/Int` — texto ↔ número con errores |
| 09 | `09_comparar_y_ordenar` | `==`, `<`, las trampas de mayúsculas y tildes, `sort` |
| 10 | `10_ejercicio_integrador` | Analizador de texto + parser de fichas CSV (usa todo lo anterior) |

## 📍 Punteros (12 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_valores_referencias` | Modelo de memoria, direcciones, valor vs referencia |
| 02 | `02_operadores` | `&` (dirección) y `*` (desreferenciar) |
| 03 | `03_ejercicio` | Ejercicio práctico con punteros |
| 04 | `04_nil_errores` | `nil`, panics por desreferenciar nil y cómo protegerse |
| 05 | `05_zero_vs_novalue` | Zero value vs "sin valor": punteros para campos opcionales |
| 06 | `06_mutabilidad` | Modificar el original desde una función |
| 07 | `07_punteros_structs` | Punteros a structs, el atajo `p.Campo` |
| 08 | `08_receptores_metodos` | Receptores valor vs puntero en métodos |
| 09 | `09_slices_referencia` | Por qué los slices ya "actúan" como referencias |
| 10 | `10_maps_referencia` | Ídem para maps: no necesitás `*map` |
| 11 | `11_interfaces_nil_trap` | La trampa de la interfaz que contiene un puntero nil |
| 12 | `12_punteros_ultimo_recurso` | Cuándo SÍ y cuándo NO usar punteros (tabla de decisión) |

## 🧩 Interfaces (10 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_una_interfaz` | El contrato: "importa qué sabe hacer, no qué es" |
| 02 | `02_implementacion_implicita` | Sin `implements`: cumplir un contrato con solo tener los métodos |
| 03 | `03_multiples_metodos` | Interfaces con varios métodos, composición de interfaces |
| 04 | `04_fmt_stringer` | La interfaz más usada de Go: `String() string` |
| 05 | `05_type_assertion` | `v, ok := i.(T)` — recuperar el tipo concreto |
| 06 | `06_type_switch` | `switch v := i.(type)` — distinguir entre varios tipos |
| 07 | `07_interfaz_vacia_any` | `any`/`interface{}`, cuándo usarlo y cuándo no |
| 08 | `08_interfaces_como_parametros` | Polimorfismo real: un checkout, varios métodos de pago |
| 09 | `09_error_interface_a_fondo` | `error` es solo una interfaz de un método |
| 10 | `10_ejercicio_integrador` | Kiosco digital: pagos, notificaciones, eventos, todo con interfaces |

## 🔒 Closures (6 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_funciones_anonimas` | Funciones sin nombre, IIFE, como argumento |
| 02 | `02_que_es_un_closure` | Una función que "recuerda" variables de donde nació |
| 03 | `03_contador_con_estado` | Closures como objetos con estado privado |
| 04 | `04_trampa_closure_en_for` | El bug clásico del for — y por qué Go 1.22+ ya lo arregló |
| 05 | `05_funciones_que_devuelven_funciones` | Fábricas y decoradores (base de los middlewares) |
| 06 | `06_ejercicio_integrador` | Reglas de descuento configurables + caja con estado |

## 📦 Paquetes (6 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_un_paquete` | `package main` vs paquetes de librería |
| 02 | `02_organizar_en_carpetas` | Tu primer paquete propio, multi-archivo |
| 03 | `03_visibilidad` | Mayúscula = exportado, minúscula = privado |
| 04 | `04_go_mod_a_fondo` | `module`, `go X.Y.Z`, `require`, `go.sum`, comandos comunes |
| 05 | `05_multiples_archivos_mismo_paquete` | Un mismo paquete repartido en varios .go |
| 06 | `06_ejercicio_integrador` | Mini proyecto con paquetes `productos` + `clientes` |

## 🚨 ErroresAvanzados (8 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_errors_new_vs_fmt_errorf` | Diferencia y errores centinela |
| 02 | `02_wrapping_con_porcentaje_w` | `%w`, agregar contexto sin perder el error original |
| 03 | `03_errors_is_a_fondo` | `errors.Is` vs `==`, varios centinela |
| 04 | `04_errors_as_astype` | Recuperar el tipo concreto, con y sin wrapping |
| 05 | `05_errores_personalizados_con_campos` | Diseñar errores con datos útiles + `Unwrap()` propio |
| 06 | `06_errors_join` | Combinar varios errores en uno (validar formularios completos) |
| 07 | `07_panic_recover` | Cuándo sí usar panic, y cómo recuperarse |
| 08 | `08_ejercicio_integrador` | Registro de usuario con manejo de errores completo |

## 🧬 Generics (5 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_por_que_generics` | El problema que resuelven (ni copiar función, ni perder tipos con `any`) |
| 02 | `02_funciones_genericas` | Sintaxis `[T any]`, inferencia de tipos |
| 03 | `03_constraints` | Constraints propios, `comparable`, `cmp.Ordered` |
| 04 | `04_tipos_genericos` | Structs genéricos: una `Pila[T]` reusable |
| 05 | `05_ejercicio_integrador` | `Filtrar`/`Mapear`/`Reducir` genéricos sobre el catálogo |

## 🧾 JSON (7 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_marshal_basico` | `json.Marshal`: struct → JSON |
| 02 | `02_unmarshal_basico` | `json.Unmarshal`: JSON → struct |
| 03 | `03_tags` | `json:"nombre"`, `omitempty`, `-` |
| 04 | `04_structs_anidados` | JSON anidado, arrays de structs, embedding |
| 05 | `05_campos_opcionales_con_punteros` | `*T` + `omitempty` para PATCH parciales (conecta con Punteros/05) |
| 06 | `06_valores_dinamicos` | `map[string]any` para JSON de forma desconocida |
| 07 | `07_ejercicio_integrador` | Procesar un "request" de pedido, preview de un handler HTTP |

## ⏰ Tiempo (6 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_time_now_y_time_time` | `time.Now()`, `time.Time`, componentes de una fecha |
| 02 | `02_duration` | `time.Duration`, constantes, `ParseDuration` |
| 03 | `03_formatear_y_parsear` | El layout de referencia `2006-01-02 15:04:05` |
| 04 | `04_sumar_restar_tiempo` | `Add`, `Sub`, `AddDate`, `Before`/`After` |
| 05 | `05_timers_y_sleep` | `time.Sleep`, `time.After`, `time.Timer` |
| 06 | `06_ejercicio_integrador` | Agenda de turnos con detección de solapamientos |

## 📁 Archivos (6 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_leer_y_escribir_archivos` | `os.ReadFile`/`WriteFile`, `os.Stat` |
| 02 | `02_bufio_lectura_eficiente` | `bufio.Scanner`/`Writer` línea por línea |
| 03 | `03_leer_csv` | `encoding/csv`: leer y escribir tablas |
| 04 | `04_os_args` | Argumentos de línea de comandos, base de un CLI |
| 05 | `05_variables_de_entorno` | `os.Getenv`/`LookupEnv`, config y secretos |
| 06 | `06_ejercicio_integrador` | Carrito persistente en disco (JSON + Archivos) |

## ⏳ Contexto (4 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_context` | `context.Context`, propagación entre funciones |
| 02 | `02_with_timeout` | Cancelación automática por tiempo |
| 03 | `03_with_cancel` | Cancelación manual, frenar el resto del trabajo |
| 04 | `04_ejercicio_integrador` | Requests con timeout (preview de HTTP) |

## 📝 Logging (3 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_slog_basico` | `log/slog`: logging estructurado con clave-valor |
| 02 | `02_niveles_y_handlers` | Text vs JSON handler, niveles, `With` |
| 03 | `03_ejercicio_integrador` | Logging de un checkout completo, con contexto por pedido |

## 🧪 Testing (6 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_go_test_basico` | `go test`, `t.Errorf` vs `t.Fatalf` |
| 02 | `02_table_driven_tests` | El patrón de testing de Go: tabla de casos |
| 03 | `03_subtests_t_run` | `t.Run` para subtests con nombre |
| 04 | `04_coverage` | `go test -cover`, por qué `main()` diluye el porcentaje |
| 05 | `05_testing_errores` | Testear funciones con `error`, y panics esperados |
| 06 | `06_ejercicio_integrador` | Suite completa: tabla + subtests + `errors.Is` |

## 🔀 Concurrencia (10 temas)

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_goroutines` | `go func()`, por qué son baratas, el riesgo de no esperarlas |
| 02 | `02_waitgroup` | `sync.WaitGroup`: la forma correcta de esperar goroutines |
| 03 | `03_channels_basico` | Enviar, recibir, cerrar, `for range` sobre un channel |
| 04 | `04_channels_buffer` | Canales con buffer, `len`/`cap` |
| 05 | `05_select` | Esperar en varios channels, `default`, timeouts |
| 06 | `06_mutex` | `sync.Mutex`, condiciones de carrera explicadas con código |
| 07 | `07_race_detector` | `go run -race` / `go test -race` |
| 08 | `08_worker_pool` | N workers fijos procesando una cola de trabajos |
| 09 | `09_context_con_goroutines` | Cancelar varias goroutines a la vez |
| 10 | `10_ejercicio_integrador` | Procesar pedidos en paralelo: worker pool + Mutex + context |

---

## 🌐 APIRest (12 temas)

Cada tema levanta un servidor y se hace requests a sí mismo: con `go run`
ves el diálogo cliente ↔ servidor completo en la consola, sin curl ni Postman.
El `12_ejercicio_integrador` además acepta `SERVIDOR=1 go run .` para quedar
vivo y probarlo vos desde el browser o curl.

| # | Tema | Qué cubre |
|---|------|-----------|
| 01 | `01_que_es_http` | Anatomía de request/response, métodos, qué significa REST |
| 02 | `02_primer_servidor` | Handlers, `http.HandleFunc`, `ListenAndServe`, la interfaz `http.Handler` |
| 03 | `03_routing_moderno` | ServeMux de Go 1.22+: `GET /productos/{id}`, `PathValue`, 404/405 gratis |
| 04 | `04_responder_json` | Helper `responderJSON`, Content-Type, errores con forma consistente |
| 05 | `05_leer_json_del_body` | `Decode` del body, `MaxBytesReader`, `DisallowUnknownFields`, validación |
| 06 | `06_status_codes` | 200/201/204/400/401/403/404/405/409/500: cuándo va cada uno |
| 07 | `07_query_params_y_headers` | Filtros y paginación con `?limite=`, leer/escribir headers |
| 08 | `08_middleware` | `func(http.Handler) http.Handler`: logging con slog + recover de panics |
| 09 | `09_middleware_auth` | `Authorization: Bearer`, 401, rutas públicas vs protegidas |
| 10 | `10_http_client` | Consumir APIs de terceros: timeout SIEMPRE, chequear status, decodificar |
| 11 | `11_graceful_shutdown` | `http.Server`, `Shutdown(ctx)`: apagar sin cortar requests en curso |
| 12 | `12_ejercicio_integrador` | La API completa del Kiosco: CRUD + auth + middlewares + Mutex |

---

## 🗺️ Hoja de ruta completa: de acá a dominar Go

El objetivo NO es aprender todo (imposible), sino dominar lo necesario para
hacer **APIs REST bien hechas** y **programas reales**. Esta es la ruta, en
orden. Cada etapa se apoya en la anterior — no saltees.

### ✅ Etapa 1 — Completar el lenguaje 🧩 (hecho)

Lo que faltaba del lenguaje en sí, ya cubierto en `Interfaces/`, `Closures/`,
`Paquetes/`, `ErroresAvanzados/` y `Generics/`.

### ✅ Etapa 2 — Biblioteca estándar esencial 📚 (hecho)

Los paquetes que usa el 90% de los programas reales, cubiertos en `JSON/`,
`Tiempo/`, `Archivos/`, `Contexto/` y `Logging/`.

> Pendiente menor: `io`/`bufio` a fondo (ya viste `bufio` en `Archivos/02`) y
> `regexp` — se agregan si algún proyecto los necesita, no bloquean nada.

### ✅ Etapa 3 — Testing 🧪 (hecho)

Cubierto en `Testing/`: `go test`, table-driven tests, subtests con `t.Run`,
coverage, testing de errores y panics.

> Pendiente menor: `httptest` (testear handlers sin levantar el servidor) —
> con `APIRest/` ya hecho, es el candidato ideal para el primer proyecto real.

### ✅ Etapa 4 — Concurrencia 🔀 (hecho)

Cubierto en `Concurrencia/`: goroutines, `sync.WaitGroup`, channels (con y sin
buffer), `select`, `sync.Mutex`, el detector de *race conditions* (`-race`),
worker pools y `context` combinado con goroutines.

### ✅ Etapa 5 — APIs REST 🌐 (hecho, salvo DB)

Cubierto en `APIRest/`: servidor y handlers, routing moderno de Go 1.22+,
JSON in/out con validación, status codes, query params y headers, middlewares
(logging, recover, auth con Bearer token), `http.Client` con timeouts, y
graceful shutdown. Todo con la stdlib sola.

Lo que queda de esta etapa (necesita dependencias externas y va en secciones
propias, se estudia junto a los proyectos):

1. **Base de datos** → sección `BaseDatos/` futura: `database/sql` + driver de
   PostgreSQL (`pgx`); queries, transacciones; el `Almacen` del ejercicio
   integrador ya está diseñado para reemplazarlo por esto
2. **Auth con JWT real** — la idea ya está en `APIRest/09`; falta solo la
   librería (`github.com/golang-jwt/jwt`) y la firma criptográfica
3. **Estructura de proyecto** — layout `cmd/` + `internal/`: se aprende
   armando el primer proyecto real, no con un tema suelto
4. **Frameworks (opcional)** — `chi`, `Gin`, `Echo`: conocelos, pero la stdlib
   alcanza y los recruiters valoran que la domines

### Etapa 6 — Herramientas y producción 🚀

Para que tus programas salgan de tu máquina.

1. **Tooling** — `golangci-lint`, `go mod tidy`, `go build` para otros sistemas (cross-compile)
2. **Docker** — empaquetar tu API en una imagen (los binarios de Go son ideales para esto)
3. **Makefile** — automatizar build/test/run
4. **Deploy básico** — subir la API a algún servicio (Railway, Fly.io, un VPS)

### 💪 Proyectos para practicar (uno por etapa)

La teoría sin proyectos no se fija. Sugerencias en orden de dificultad:

1. **CLI de tareas** (etapas 1-2): agregar/listar/completar tareas, guardadas en un archivo JSON
2. **Conversor de monedas** (etapa 2): consume una API de cotizaciones con `http.Client`
3. **Acortador de URLs** (etapa 5): tu primera API REST, primero en memoria (map), después con DB
4. **API de tareas completa** (etapas 5-6): CRUD + PostgreSQL + auth JWT + tests + Docker — esto ya es nivel portfolio

> Regla de oro: cuando termines Punteros e Interfaces, empezá el proyecto 1
> AUNQUE sientas que te falta. Se aprende construyendo y volviendo a la
> teoría cuando algo no sale.
