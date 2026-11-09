# 03 — Colecciones y strings

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Dominar cómo Go representa secuencias, tablas y texto UTF-8.

## Recorrido

Arrays y slices → operaciones y aliasing → maps y su semántica → bytes, runes y
UTF-8 → búsqueda, transformación, conversión y ordenamiento de strings.

Regla central: todos los valores se asignan y pasan por valor. Copiar un slice o
un map puede hacer que dos valores compartan almacenamiento interno; eso no crea
un mecanismo distinto llamado “paso por referencia”.

Los ejemplos de inventario/carrito usan structs y métodos, y algunas llamadas
de ordenamiento reciben funciones anónimas. Primero estudiá el comportamiento
de las colecciones; volvé al diseño de esos helpers después de `04` y `05`.
Las APIs `slices.Sorted` y `maps.Keys` usan iteradores que se explican en `08`.
El reto inicial usa `[]string`, sin exigir generics antes de esa etapa.

## Mapa de temas

### Arrays y Slices (`01` a `03`)
- `01_arrays_y_slices`: tamaño fijo vs dinámico, `len` y `cap`, `make`, `append`, `copy`, slicing y paquete `slices`.
- `02_operaciones_con_slices`: paso de descriptores de slice a funciones, mutación de elementos vs `append`, y patrones comunes (filtrar, insertar, eliminar, invertir).
- `03_internals_y_aliasing_de_slices`: header `(ptr, len, cap)`, array subyacente compartido, expresión de 3 índices (`s[a:b:c]`), `slices.Clone` y `clear`.

### Maps (`04` a `06`)
- `04_maps`: creación con `make` y literales, lectura, escritura, comma-ok idiom (`valor, existe := m[k]`), `delete`, `clear` y `maps.Equal`.
- `05_iteracion_de_maps`: recorrido con `for range`, orden no especificado, orden determinístico con `slices.Sort` / `slices.Sorted(maps.Keys(m))`, `maps.DeleteFunc` y casos prácticos.
- `06_semantica_de_maps`: `nil` map vs map inicializado, descriptores que comparten tabla, patrón conjunto (`map[T]struct{}`) y `maps.Clone`.

### Strings, UTF-8 y Conversiones (`07` a `16`)
- `07_que_es_un_string`: secuencia inmutable de bytes, `len()` en bytes, escapes, raw strings (backticks) y `[]byte`.
- `08_utf8_y_runes`: Unicode, UTF-8 de ancho variable (1 a 4 bytes), `rune` (`int32`), `utf8.RuneCountInString`, `utf8.ValidString` y trampa de `string(int)`.
- `09_indexar_e_iterar`: indexación por byte (`s[i]`) vs iteración por rune (`for range`), acceso por posición con `[]rune(s)` y slicing seguro.
- `10_inmutabilidad`: por qué los strings son inmutables, cómo modificar vía `[]rune` o `[]byte`, y seguridad al compartir memoria.
- `11_concatenar_y_builder`: `+`, `fmt.Sprintf`, `strings.Join`, `strings.Builder` para loops y `strings.Repeat`.
- `12_paquete_strings_buscar`: `Contains`, `HasPrefix`, `HasSuffix`, `Index`, `LastIndex`, `Cut`, `Count` y `EqualFold`.
- `13_paquete_strings_transformar`: `ToUpper`, `ToLower`, `Split`, `Fields`, `SplitSeq`/`FieldsSeq`, `Join`, `TrimSpace`/`TrimPrefix`/`TrimSuffix` y `ReplaceAll`.
- `14_strconv_conversiones`: `Itoa`, `Atoi`, `ParseFloat`, `ParseBool`, `ParseInt` y `FormatFloat` con manejo de errores.
- `15_comparar_y_ordenar`: orden lexicográfico por bytes, trampas con mayúsculas y tildes, `slices.Sort`, `slices.SortFunc` y `slices.BinarySearch`.
- `16_ejercicio_integrador`: analizador de frecuencia de palabras y parser de fichas CSV integrando colecciones, strings y manejo de errores.

## Prerrequisito y reto

Requiere `02_control_de_flujo`. Resolvé `retos/` y después compará `soluciones/`.
