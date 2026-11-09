# Práctica — contratos y errores

Comandos desde la raíz. Pensá qué sabe el consumidor, no solo qué tipo creaste.

## 1. Distinguir nil de un puntero nil dentro de una interfaz

    go run ./06_interfaces_y_errores/11_interfaz_con_puntero_nil

Buscá el caso que asigna un puntero nil a una interfaz. La comparación de esa
interfaz con nil es false: conserva un tipo dinámico aunque su valor sea nil.

**Comprobación:** ¿por qué una función que devuelve error debería devolver nil
directamente para indicar éxito? ¿Llamar un método con receptor nil siempre
provoca panic, o depende de lo que haga ese método?

## 2. Preservar identidad además del mensaje

    go run ./06_interfaces_y_errores/14_wrapping_con_porcentaje_w

errors.Is puede encontrar el error envuelto con %w. Cambiá %w por %v: el texto
puede parecer el mismo, pero se pierde la conexión con la causa.

**Comprobación:** ¿por qué comparar err.Error() con un texto resulta frágil?

## 3. Recuperar información de un error

    go run ./06_interfaces_y_errores/16_errors_as_astype
    go test -tags reto ./06_interfaces_y_errores/retos

Primero observá que AsType devuelve el error concreto y un bool. Después
implementá el reto: debe añadir el ID sin perder ErrNoEncontrado.

**Comprobación:** ¿cuándo querés errors.Is y cuándo AsType?
¿Dónde se decide cómo reaccionar al error y dónde solo se propaga?
