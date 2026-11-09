# Práctica — seguir las ramas y los loops

Ejecutá desde la raíz y anotá qué instrucciones se alcanzan.

## 1. Entender qué recibe range

    go run ./02_control_de_flujo/09_range

En el ejemplo de precios, modificar p deja [10 20 30]; modificar precios[i]
produce [20 40 60]. p es una copia del elemento, no una forma de asignarlo.

**Experimento:** cambiá la operación del segundo loop por += 1. Esperá
[11 21 31]. ¿Por qué usar el índice cambia el resultado?

## 2. Cambiar la forma de iterar

En una copia de trabajo de la lección 04, reemplazá un loop de tres vueltas
por "for i := range 3". Ambos recorren los índices 0, 1 y 2.
Con range 0 no se ejecuta el cuerpo.

**Comprobación:** ¿dónde está la condición de salida de un for infinito?
¿Qué diferencia hay entre continue, break y return dentro de un loop?

## 3. Resolver FizzBuzz por reglas

    go test -tags reto ./02_control_de_flujo/retos

Completá el starter antes de esperar que pase. Para límite 15, las posiciones
3, 5 y 15 deben contener Fizz, Buzz y FizzBuzz. Para límite 0 no hay elementos.
Comprobá el caso divisible por ambos antes de las reglas individuales.

**Comprobación:** ¿qué falla si revisás divisibilidad por 3 antes que por 15?
