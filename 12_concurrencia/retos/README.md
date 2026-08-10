# Reto: pipeline cancelable

Devolvé los cuadrados desde un channel. Cerrá la salida y respetá cancelación
en cada recepción y envío para no filtrar goroutines. Usá `go test -tags reto .`.
