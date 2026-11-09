# Práctica — reutilizar sin borrar los tipos

Comandos desde la raíz. Identificá qué operaciones permite la restricción.

## 1. Generalizar un algoritmo conocido

    go run ./08_generics_e_iteradores/07_clonar_slices

Esperá original [1 2], clon [99 2], nil conservado true, vacío no-nil true
e interior compartido 7. T cambia el tipo de elemento, no la semántica de copiar.

**Comprobación:** ¿por qué `Clonar(numeros)` infiere T, pero `Clonar[int](nil)`
necesita escribirlo? ¿any garantiza que los elementos no compartan datos?

## 2. Elegir una restricción suficiente

En 03_constraints, intentá sustituir comparable por any en Contiene.
La operación == ya no estará garantizada para T, por lo que no compila.

**Comprobación:** ¿qué exige comparable? ¿Por qué una restricción numérica puede
usar ~int para admitir tipos definidos cuyo tipo subyacente sea int?

## 3. Detener un iterador

    go run ./08_generics_e_iteradores/06_iteradores_y_go_126

La demo entrega pares [2 4 6] y claves ordenadas [mouse teclado]. Sustituí la
colección completa de pares por un range que haga break al encontrar el primero.
Debe terminar con 2 sin seguir evaluando los candidatos siguientes; podés
imprimir dentro del predicado para observarlo.

**Comprobación:** ¿qué informa el bool devuelto por yield?
¿Quién produce los valores y quién decide cuándo dejar de consumirlos?
