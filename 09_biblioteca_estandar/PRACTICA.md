# Práctica — componer IO y cancelar trabajo

Comandos desde la raíz. Estas piezas también sirven para programas de terminal.

## 1. Consumir una fuente sin depender de archivos

    go run ./09_biblioteca_estandar/01_io

LimitReader devuelve "0123" y ContarBytes devuelve 4 para "hola".
Sustituí el Reader de la demo por otro texto; la función no cambia.

**Comprobación:** ¿qué método debe implementar io.Reader? ¿Qué significa
recibir datos y un error en la misma llamada? ¿LimitReader distingue una entrada
de longitud 4 de una entrada más larga que cortó al llegar al límite?

## 2. Transformar progresivamente

    go test ./09_biblioteca_estandar/soluciones
    go test -tags reto ./09_biblioteca_estandar/retos

La solución procesa UTF-8 incluso cuando cada Read entrega un solo byte.
También conserva la causa de errores y detecta escrituras cortas. El starter
debe implementarse antes de pasar sus pruebas. Una salida parcial es posible
si la lectura falla después de haber escrito texto.

**Comprobación:** ¿por qué io.ReadAll no resuelve el objetivo de streaming?
¿Quién debe cerrar el Reader o Writer si los recibió como argumentos?

## 3. Cancelar no equivale a interrumpir por fuerza

    go run ./09_biblioteca_estandar/05_contexto/02_with_timeout

El ejemplo muestra DeadlineExceeded al agotarse el plazo. Cambiá el tiempo de
trabajo y el timeout para observar éxito y cancelación.
El trabajo debe consultar Done o usar operaciones que respeten el contexto.

**Comprobación:** ¿cancelar un context mata automáticamente una goroutine?
¿Por qué se llama a cancel incluso si el deadline termina por sí solo?
