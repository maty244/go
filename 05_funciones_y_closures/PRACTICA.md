# Práctica — funciones como valores

Comandos desde la raíz. Distinguí cuándo se crea una función de cuándo se llama.

## 1. Capturar una variable o declarar una por vuelta

    go run ./05_funciones_y_closures/10_trampa_closure_en_for

El primer loop imprime 0, 1, 2. capturarExterior imprime 3, 3, 3 porque usa =
sobre una i declarada afuera. Los closures leen i cuando se ejecutan.

**Experimento:** en capturarExterior quitá var i int y declarala en el for con
:=. Esperá 0, 1, 2; el test existente fallará porque documentaba el otro caso.
Explicá el cambio antes de ajustar la expectativa.

## 2. Decidir cuándo se evalúa una expresión

    go run ./05_funciones_y_closures/04_defer

Los defers de demostrarLIFO corren 3, 2, 1. El argumento x de defer Println
guarda 10 aunque después x sea 99. Cambiá esa llamada por un closure que lea x:
el closure verá 99 cuando se ejecute.

**Comprobación:** ¿el defer de un loop se ejecuta al terminar la vuelta o al
salir de la función? ¿Qué parte de defer medirTiempo(...)() corre inmediatamente?

## 3. Componer transformaciones

    go test -tags reto ./05_funciones_y_closures/retos

Implementá Aplicar sin cambiar la entrada. Para [1, 2], sumar 1 y luego
multiplicar por 2 debe producir [4, 6]. Invertir el orden produce [3, 5].
Probá también sin transformaciones.

**Comprobación:** ¿qué valores tendría la entrada si trabajaras sobre su array?
