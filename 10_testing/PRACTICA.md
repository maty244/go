# Práctica — comprobar, diagnosticar y medir

Comandos desde la raíz. La corrección se comprueba antes de comparar rendimiento.

## 1. Proteger comportamiento

    go test -v ./10_testing/09_diagnostico_y_perfiles

El test verifica texto vacío, ASCII y Unicode contra strings.Join.
Elegí otro caso con espacios o partes vacías. Ambas funciones deben devolver
exactamente el mismo texto.

**Comprobación:** ¿un test que pasa demuestra corrección para toda entrada?
Para depurar, ejecutá un subtest con -run, usá t.Logf y leé la primera línea del
stack que pertenezca a tu código. Si tenés debugger en el editor, poné un
breakpoint antes del return e inspeccioná los valores y la pila de llamadas.

## 2. Medir una operación

    go test -run='^$' -bench=. -benchmem -count=3 ./10_testing/09_diagnostico_y_perfiles

Compará ns/op, B/op y allocs/op entre += y Builder con la misma entrada.
No hay cifras universales ni una proporción fija que deba aparecer.
B/op describe asignaciones acumuladas por operación, no memoria retenida.

**Experimento:** cambiá 100 repeticiones por 1 y después por 1000.
¿La diferencia crece igual que el tamaño de la entrada?

## 3. Buscar dónde se gasta el tiempo

Antes de usar el visor, comprobá si tu instalación lo incluye:

    go tool

La lista debe incluir pprof. Si aparece "no such tool pprof", falta el visor
en esa instalación: podés generar los archivos de perfil y practicar benchmarks,
pero para leerlos con estos comandos necesitás una distribución de Go que
incluya la herramienta. Esto no implica un error en el código del ejemplo.

    go test -run='^$' -bench=BenchmarkConcatenar -benchtime=1s -o /tmp/gocito-perfiles.test -cpuprofile=/tmp/gocito-cpu.out ./10_testing/09_diagnostico_y_perfiles
    go tool pprof -top /tmp/gocito-perfiles.test /tmp/gocito-cpu.out

El perfil CPU toma muestras de ejecución; no mide directamente tiempo esperando
IO. Mirá flat (en la función) y cum (incluyendo llamadas). El trabajo puede
aparecer en funciones del runtime por inlining o asignaciones.

Para memoria, generá un perfil separado:

    go test -run='^$' -bench=BenchmarkConcatenar -benchtime=1s -o /tmp/gocito-perfiles.test -memprofile=/tmp/gocito-mem.out ./10_testing/09_diagnostico_y_perfiles
    go tool pprof -top -alloc_space /tmp/gocito-perfiles.test /tmp/gocito-mem.out

alloc_space observa bytes asignados; inuse_space observa memoria retenida en
el perfil. No son la misma pregunta. Los perfiles quedan en /tmp.

Fuentes: [diagnóstico de Go](https://go.dev/doc/diagnostics) y
[comando pprof](https://pkg.go.dev/cmd/pprof).
