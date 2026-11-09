# Práctica — valores, tipos y scope

Ejecutá los comandos desde la raíz. Primero escribí tu predicción.

## 1. Distinguir declaración y asignación

    go run ./01_fundamentos/16_scope_y_shadowing

En el primer ejemplo, x vale "exterior", dentro del bloque "interior" y afuera
otra vez "exterior". La asignación con = del segundo ejemplo sí cambia el valor
exterior a 99. Cambiá un := interior por = y explicá el efecto.

**Comprobación:** := requiere al menos un nombre nuevo en el mismo bloque.
En "valor, nuevo := 20, true", si valor ya existe ahí, ¿se declara otra valor
o se reasigna la existente? ¿Qué ocurre si ponés esa línea en otro bloque?

## 2. Probar los límites de un tipo

    go run ./01_fundamentos/18_errores_comunes

El ejemplo de int8 muestra que incrementar 127 produce -128. Los enteros
tienen límites aunque no tengan redondeo binario. La conversión de una variable
300 a int8 produce 44; convertir la constante 300 directamente no compila.

**Experimento:** escribí ambos casos por separado. ¿Por qué uno falla al
compilar y el otro produce un resultado en ejecución?

## 3. Explicar la unidad de un valor

    go run ./01_fundamentos/19_dinero_con_enteros

El total de 1299 centavos por 3 unidades se muestra como $38.97. Cambiá la
cantidad a -1: se devuelve un error y no se imprime un total válido.
Esta función básica no detecta overflow: no la interpretes como cálculo
financiero completo.

**Comprobación:** ¿qué aporta type Dinero int64? ¿Qué operación necesita una
conversión explícita? ¿Puede un nombre de tipo evitar por sí solo el overflow?
