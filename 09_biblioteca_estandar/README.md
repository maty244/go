# 09 — Biblioteca estándar

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Aprender los contratos y paquetes que aparecen constantemente en programas
reales, antes de incorporar dependencias externas.

## Mapa de temas (`01` a `06`)

| Bloque / Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| **`01_io`** | Interfaces `io.Reader` e `io.Writer` | Contratos `Read`/`Write`, `io.Copy`, `io.LimitReader`, `io.Discard` y streaming sin acoplarse a archivos |
| **`02_archivos`** | **Archivos, entorno y sistema operativo** | |
| `02_archivos/01_leer_y_escribir_archivos` | `os.ReadFile` y `os.WriteFile` | Lectura y escritura completa, permisos octales (`0o644`), `os.Stat` y `errors.Is(err, os.ErrNotExist)` |
| `02_archivos/02_bufio_lectura_eficiente` | `bufio.Scanner` y `bufio.Writer` | Lectura línea por línea sin cargar todo en memoria y escritura bufferizada con `Flush()` |
| `02_archivos/03_leer_csv` | `encoding/csv` | Escribir y leer tablas CSV manejando comas y comillas internas (`csv.Reader` / `csv.Writer`) |
| `02_archivos/04_os_args` | `os.Args` | Argumentos de línea de comandos (`os.Args[0]` vs `os.Args[1:]`) y estructura de un mini-CLI |
| `02_archivos/05_variables_de_entorno` | Variables de entorno | `os.Getenv`, `os.LookupEnv`, `os.Setenv` y configuración sin hardcodear secretos |
| `02_archivos/06_ejercicio_integrador` | Integrador de archivos | Persistencia en disco combinando `os.WriteFile`/`os.ReadFile` y JSON |
| `02_archivos/07_os_root` | `os.Root` (`os.OpenRoot`) | Confinar operaciones de archivos a un directorio raíz previniendo escapes con `../` |
| **`03_json`** | **Codificación y decodificación JSON** | |
| `03_json/01_marshal_basico` | `json.Marshal` y `MarshalIndent` | Serializar structs, slices y maps a JSON; visibilidad de campos exportados |
| `03_json/02_unmarshal_basico` | `json.Unmarshal` | Deserializar JSON hacia structs y slices pasando puntero (`&destino`) |
| `03_json/03_tags` | Tags de struct para JSON | Renombrar claves, excluir con `json:"-"`, `omitempty` y `omitzero` (Go 1.24+) |
| `03_json/04_structs_anidados` | Structs anidados y embedding | Objetos y listas anidadas en JSON, y cómo el embedding aplana campos |
| `03_json/05_campos_opcionales_con_punteros` | Campos opcionales (`*T`) | Distinguir "no vino en el JSON" (`nil`) de "vino con valor cero" (`0`/`""`) y uso con `new(valor)` (Go 1.26) |
| `03_json/06_valores_dinamicos` | `map[string]any` | Leer JSON de esquema variable, números como `float64` y type assertions |
| `03_json/07_ejercicio_integrador` | Integrador de JSON | Validar un request JSON de pedido y construir una respuesta estructurada |
| **`04_tiempo`** | **Fechas, duraciones y timers (`time`)** | |
| `04_tiempo/01_time_now_y_time_time` | `time.Now()` y `time.Time` | Instantes, componentes de fecha, `time.Date` y comparación segura con `Before`/`After`/`Equal` |
| `04_tiempo/02_duration` | `time.Duration` | Cantidades de tiempo, constantes (`time.Second`, `time.Hour`) y `time.ParseDuration` |
| `04_tiempo/03_formatear_y_parsear` | `Format` y `Parse` | El layout de referencia `2006-01-02 15:04:05`, `time.RFC3339` y `time.DateOnly` |
| `04_tiempo/04_sumar_restar_tiempo` | Aritmética de tiempo | `Add`, `Sub`, `time.Since` y calendario real con `AddDate` |
| `04_tiempo/05_timers_y_sleep` | `time.Sleep`, `After` y `Timer` | Pausas, alarmas por canal, cancelación con `timer.Stop()` y reintentos con backoff |
| `04_tiempo/06_ejercicio_integrador` | Integrador de tiempo | Agenda de turnos con detección de solapamientos y vigencia de promociones |
| **`05_contexto`** | **Cancelación y plazos (`context`)** | |
| `05_contexto/01_que_es_context` | `context.Context` y `Background` | Propagación de contexto como primer parámetro (`ctx context.Context`) |
| `05_contexto/02_with_timeout` | `context.WithTimeout` | Plazos automáticos, `<-ctx.Done()`, `context.DeadlineExceeded` y por qué siempre llamar `defer cancel()` |
| `05_contexto/03_with_cancel` | `context.WithCancel` | Cancelación cooperativa manual ante eventos o errores en cadena |
| `05_contexto/04_ejercicio_integrador` | Integrador de contexto | Simulación de un request con timeout compartido entre múltiples consultas |
| **`06_logging`** | **Logging estructurado (`log/slog`)** | |
| `06_logging/01_slog_basico` | `slog` básico | Niveles (`Debug`, `Info`, `Warn`, `Error`) y pares clave-valor estructurados |
| `06_logging/02_niveles_y_handlers` | Niveles, Handlers y `MultiHandler` | `TextHandler`, `JSONHandler`, `HandlerOptions`, `With` y `slog.NewMultiHandler` (Go 1.26) |
| `06_logging/03_ejercicio_integrador` | Integrador de logging | Trazabilidad de un flujo de checkout con `logger.With` y errores centinela |

## Prerrequisito y reto

Requiere las etapas `01` a `08`. Resolvé `retos/` y compará `soluciones/`.
