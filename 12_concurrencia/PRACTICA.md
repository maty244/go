# Práctica — coordinación, bloqueos y cancelación

Comandos desde la raíz. El orden de los mensajes concurrentes puede cambiar.

## 1. Distinguir un channel nil de uno cerrado

    go run ./12_concurrencia/13_channels_nil_y_cerrados

Esperá un caso nil deshabilitado, luego "recibir: 10 true", "recibir: 20 true",
"recibir: 0 false" y finalmente otro caso deshabilitado.
Quitar cola = nil hace que el último select reciba el zero value del channel
cerrado. No convierte ese channel en uno nuevo.

**Comprobación:** ¿qué ocurre al enviar a nil? ¿Y al enviar a uno cerrado?
¿Quién debería cerrar la salida cuando hay varios productores?

## 2. Procesar más trabajos que el buffer

    go test -v ./12_concurrencia/08_worker_pool
    go run ./12_concurrencia/08_worker_pool

Los tests procesan 0, 1 y 20 trabajos y comprueban el cierre. La demo produce
ocho resultados, con IDs únicos; el orden y el worker que atiende cada ID
pueden variar. El productor y el consumidor trabajan simultáneamente.

**Experimento:** aumentá los productos y reducí el buffer a cero.
Debe seguir terminando. ¿Por qué enviar toda la entrada antes de recibir
resultados puede bloquear incluso con un número fijo de workers?

## 3. Detectar carreras y evitar fugas

    go run -race ./12_concurrencia/07_race_detector

Esta lección tiene una carrera INTENCIONAL: es esperable un reporte DATA RACE
y un código de salida de error. No es un fallo accidental del curso.
Después resolvé Cuadrados, que debe respetar cancelación al recibir y enviar:

    go test -race -tags reto ./12_concurrencia/retos

**Comprobación:** si el consumidor deja de leer, ¿quién libera al productor?
¿time.Sleep sincroniza acceso a memoria? ¿go test -race ejecuta el main de
una lección sin tests? No; una corrida limpia solo cubre lo que se ejecutó.
