# 06 — Interfaces y errores

La [práctica guiada](./PRACTICA.md) propone experimentos, resultados esperados
y preguntas para comprobar que entendiste el comportamiento.

## Objetivo

Desacoplar comportamiento con contratos pequeños y conservar información útil
cuando algo falla.

## Recorrido

Interfaces implícitas → composición → `Stringer` → assertions y type switch →
`any` → nil en interfaces → `error` → sentinelas → wrapping → `errors.Is` →
`errors.AsType` → errores con campos → `errors.Join` → `panic/recover`.

Preferí aceptar interfaces y devolver tipos concretos. Definí la interfaz cerca
del código que la consume, no automáticamente junto a la implementación.

## Mapa de temas (`01` a `20`)

| Carpeta | Tema | Qué vas a aprender |
|---|---|---|
| `01_que_es_una_interfaz` | ¿Qué es una interfaz? | Contratos de comportamiento, slices de interfaces y desacoplamiento |
| `02_implementacion_implicita` | Implementación implícita | Duck typing estático, definir interfaces después del tipo y chequeo `var _ I = T{}` |
| `03_multiples_metodos` | Múltiples métodos y composición | Interfaces pequeñas (1-2 métodos) y composición de interfaces (`interface { A; B }`) |
| `04_fmt_stringer` | `fmt.Stringer` | Formateo automático en `fmt`, diferencia entre `%v`/`%s` y `%+v`, y trampa de recursión infinita |
| `05_type_assertion` | Type assertion (`i.(T)`) | Recuperar el tipo concreto de forma insegura vs segura (`v, ok := i.(T)`) |
| `06_type_switch` | Type switch (`i.(type)`) | Distinguir múltiples tipos concretos, casos agrupados y despacho de eventos |
| `07_interfaz_vacia_any` | Interfaz vacía (`any`) | `any` (`interface{}`), colecciones heterogéneas y cuándo preferir Generics |
| `08_interfaces_como_parametros` | Polimorfismo en funciones | Procesar distintos métodos de pago con una única función sin acoplamiento |
| `09_error_interface_a_fondo` | `error` es una interfaz | `type error interface { Error() string }` y tipos de error propios |
| `10_ejercicio_integrador` | Integrador Parte 1: Kiosco | Polimorfismo con `MetodoPago`, `Notificador`, `Evento`, `Stringer` y errores propios |
| `11_interfaz_con_puntero_nil` | Nil interface trap | Par `(type, value)`, por qué `(*T)(nil)` en `error` no es `nil`, y receptores `nil` |
| `12_patron_de_errores` | Patrón de errores en Go | Errores como valores, `if err != nil`, centinelas, `errors.Is` y `errors.AsType` |
| `13_errors_new_vs_fmt_errorf` | `errors.New` vs `fmt.Errorf` | Errores de texto fijo (identidad única) vs errores formateados con contexto |
| `14_wrapping_con_porcentaje_w` | Wrapping con `%w` | Envolver errores con `%w` vs `%v`, cadenas de errores y `errors.Unwrap` |
| `15_errors_is_a_fondo` | `errors.Is` a fondo | Comparar contra errores centinela atravesando capas de wrapping (`%w`) |
| `16_errors_as_astype` | `errors.AsType[T]` y `errors.As` | Extraer errores tipados en Go 1.26+ con `errors.AsType[T](err)` y forma clásica `errors.As` |
| `17_errores_personalizados_con_campos` | Diseño de errores con campos | Métodos extra en errores estructurados e implementación manual de `Unwrap() error` |
| `18_errors_join` | `errors.Join` | Combinar múltiples errores independientes (ej. validaciones) en un único `error` |
| `19_panic_recover` | `panic` y `recover` | Errores esperables vs situaciones excepcionales, desenrollado de stack y `recover()` en `defer` |
| `20_ejercicio_integrador` | Integrador Parte 2: Registro | Validación completa combinando centinelas, `%w`, struct de error, `errors.Join`, `AsType` y `recover` |

## Prerrequisito y reto

Requiere `05_funciones_y_closures`. Resolvé `retos/` y compará `soluciones/`.
