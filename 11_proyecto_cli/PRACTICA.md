# Práctica — conectar las piezas de Go

Comandos desde la raíz. Usá un archivo nuevo para que los IDs empiecen en 1.

## 1. Seguir el viaje de un argumento

    go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/gocito-practica-tareas.json agregar "Estudiar Go"
    go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/gocito-practica-tareas.json listar
    go run ./11_proyecto_cli/cmd/tareas -archivo /tmp/gocito-practica-tareas.json completar 1

Con un archivo inicialmente inexistente: se agrega tarea 1, listar muestra
"[ ] 1: Estudiar Go" y completar marca la tarea.
Seguí Args → Parse → args → Cargar → Nuevo → Agregar/Completar → Guardar.

**Comprobación:** ¿por qué -archivo va antes del comando? ¿Qué recibe main
si el título se escribe sin comillas? ¿Por qué flags.String devuelve *string?

## 2. Separar lógica de efectos externos

    go test -v ./11_proyecto_cli/internal/tareas ./11_proyecto_cli/internal/almacenjson

Los tests del gestor usan una fecha fija; los de almacenamiento usan TempDir.
No necesitan una terminal. Seguí las copias de slices en Nuevo, Listar y Estado.

**Comprobación:** ¿por qué la lógica recibe ahora como argumento?
¿Qué podría cambiar el consumidor si Listar devolviera el slice interno?

## 3. Provocar un fallo y seguir su retorno

Intentá agregar un título con solo espacios y completar el ID "abc".
Ambos deben terminar con error sin guardar una nueva tarea. Explicá cada
return hasta llegar a main, donde el error se imprime en Stderr.

**Comprobación:** ¿por qué un error de negocio no llama a os.Exit?
¿Una escritura con temporal y Rename evita también conflictos entre dos
procesos que cargan el mismo archivo a la vez? No: ese caso requiere otra
forma de coordinación; no confundas reemplazo de archivo con exclusión mutua.
