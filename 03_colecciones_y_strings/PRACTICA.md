# Práctica — datos compartidos y texto

Comandos desde la raíz. El objetivo es distinguir valor, almacenamiento y texto.

## 1. Copiar el descriptor o los elementos

    go run ./03_colecciones_y_strings/03_internals_y_aliasing_de_slices

Con s2 := s1, modificar s2[0] también modifica s1[0]. Con make y copy,
modificar un int del destino deja intacta la fuente. Cambiá la capacidad
inicial en los ejemplos de append y comprobá cuándo se comparte el array.

**Comprobación:** aunque append reutilice el array, ¿cambia la longitud de la
variable del llamador si la función no devuelve el slice?

## 2. Contar bytes, runes y caracteres visibles

    go run ./03_colecciones_y_strings/08_utf8_y_runes

"é" tiene 2 bytes y 1 rune; "e\u0301" tiene 3 bytes y 2 runes. Se ven parecidas,
pero == devuelve false. El string con byte 0xff no es UTF-8 válido; range
produce U+FFFD y luego U+0041.

**Experimento:** probá un emoji compuesto como "👩‍💻". Una imagen visible no
implica una sola rune. ¿Cuál de estos conteos calcula len?

## 3. Clonar conservando nil y vacío

    go test -tags reto ./03_colecciones_y_strings/retos

Implementá Clonar para []string. nil debe seguir nil; []string{} debe seguir
vacío no-nil; cambiar copia[0] no debe cambiar original[0].
La versión genérica se estudia después en 08/07_clonar_slices.

**Comprobación:** ¿por qué append([]string(nil), entrada...) pierde la
diferencia entre nil y vacío cuando entrada tiene longitud cero?
