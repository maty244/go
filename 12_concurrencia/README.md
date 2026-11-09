# 12 — Concurrencia

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Coordinar trabajo simultáneo sin carreras, bloqueos ni goroutines abandonadas.

## Recorrido

Goroutines → `WaitGroup.Go` y compatibilidad `Add/Done` → channels → buffers →
`select` → Mutex/RWMutex/Once → race detector → worker pools → contexto →
integrador → `testing/synctest` → diferencias entre channels nil y cerrados
(`13_channels_nil_y_cerrados`).

Concurrencia no implica paralelismo. Una goroutine es relativamente liviana,
pero sigue consumiendo memoria, tiempo de scheduler y los recursos que abra.

## Mapa de temas (`01` a `13`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_goroutines` | Goroutines (`go f()`) | Lanzar ejecución concurrente, diferencia entre concurrencia y paralelismo, y por qué `main()` no espera |
| `02_waitgroup` | `sync.WaitGroup` (`wg.Go`) | Esperar goroutines con `wg.Go` (Go 1.25+) y compatibilidad con el patrón clásico `Add`/`Done` |
| `03_channels_basico` | Channels sin buffer | Comunicación sincrónica (`ch <- v`, `<-ch`), canales direccionales (`chan<-`), `close` y `for range` |
| `04_channels_buffer` | Channels con buffer | Capacidad `make(chan T, n)`, envíos no bloqueantes hasta llenar el buffer, `len` y `cap` |
| `05_select` | `select` | Multiplexar varios canales, no bloqueo con `default`, drenaje con labels y timeouts |
| `06_mutex` | `sync.Mutex` | Condiciones de carrera sobre estado compartido y exclusión mutua con `Lock`/`Unlock` |
| `07_race_detector` | Race Detector (`-race`) | Detectar condiciones de carrera en tiempo de ejecución con `go run -race` y `go test -race` |
| `08_worker_pool` | Worker Pool | Limitar concurrencia con un número fijo de workers consumiendo de un canal de trabajos |
| `09_context_con_goroutines` | `context` + Goroutines | Cancelar múltiples goroutines por timeout o error fatal evitando *goroutine leaks* |
| `10_ejercicio_integrador` | Integrador de concurrencia | Worker pool + `context.WithTimeout` + `sync.Mutex` para estadísticas compartidas |
| `11_rwmutex_once_y_memory_model` | `RWMutex`, `OnceValue` y Memory Model | Lecturas concurrentes con `RLock`, inicialización única con `sync.OnceValue` y garantías de sincronización |
| `12_synctest` | `testing/synctest` (Go 1.25+) | Burbujas con reloj virtual e instantáneo (`synctest.Test` y `synctest.Wait`) para tests concurrentes |
| `13_channels_nil_y_cerrados` | Channels `nil` vs cerrados | Tabla de comportamiento al enviar/recibir/cerrar y cómo deshabilitar ramas de un `select` asignando `nil` |

## Prerrequisito y reto

Requiere `10_testing` y conviene haber terminado el CLI. Resolvé `retos/` y
compará `soluciones/`.
