# 10 — Testing

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Usar tests para expresar comportamiento, proteger refactors y encontrar casos
que un ejemplo manual no revela.

## Recorrido

Tests básicos → tablas → subtests → cobertura de statements → errores y panics
→ integrador → helpers/cleanup → fuzzing y benchmarks → diagnóstico y perfiles
(`09_diagnostico_y_perfiles`, sin necesitar HTTP).

El nombre del archivo puede ser cualquiera siempre que termine en `_test.go`.
La cobertura mide statements ejecutados; 100 % no demuestra ausencia de bugs.

## Mapa de temas (`01` a `09`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_go_test_basico` | `go test` básico | Convención `*_test.go`, `func TestX(t *testing.T)`, `t.Errorf` vs `t.Fatalf` |
| `02_table_driven_tests` | Table-driven tests | Slice de structs `{nombre, entrada, esperado}` recorrido con un único loop |
| `03_subtests_t_run` | Subtests con `t.Run` | Identificar cada fila con `t.Run`, correr un caso puntual con `-run` y nota sobre `t.Parallel()` |
| `04_coverage` | Cobertura (`-cover`) | Medir statements ejecutados con `go test -cover` y reporte HTML con `go tool cover` |
| `05_testing_errores` | Testear errores y panics | Verificar el caso feliz, errores con `errors.Is` y panics esperados con `defer` + `recover()` |
| `06_ejercicio_integrador` | Integrador: suite completa | Tabla + `t.Run` + `errors.Is` + orden de validación sobre reglas de negocio |
| `07_helpers_cleanup_tempdir` | Helpers, `TempDir` y entorno | `t.Helper()`, `t.TempDir()`, `t.Setenv()`, `t.Cleanup()`, `t.Context()` y `t.ArtifactDir()` (Go 1.26) |
| `08_fuzzing_benchmarks_ejemplos` | Fuzzing, Benchmarks y Examples | `FuzzX` (`*testing.F`), `BenchmarkX` con `for b.Loop()` y `ExampleX` ejecutable |
| `09_diagnostico_y_perfiles` | Diagnóstico y perfiles (`pprof`) | `b.ReportAllocs()`, comparar `+=` vs `strings.Builder` y perfiles de CPU/memoria con `go tool pprof` |

## Prerrequisito y reto

Requiere `09_biblioteca_estandar`. Resolvé `retos/` y después compará
`soluciones/`.
