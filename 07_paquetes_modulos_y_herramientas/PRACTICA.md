# Práctica — del fuente al programa

Comandos desde la raíz. Separá reglas del lenguaje de decisiones del compilador.

## 1. Observar la inicialización

    go run ./07_paquetes_modulos_y_herramientas/08_compilador_y_runtime

Las primeras líneas son inicialización de variable, init y main, en ese orden.
El puntero a una variable local sigue dando 42 después del retorno.
La versión, las CPU y GOMAXPROCS dependen del entorno.

**Comprobación:** ¿quién ejecuta init? ¿Por qué devolver &valor no deja un
puntero colgante? ¿GC cierra un archivo que olvidaste cerrar?

## 2. Inspeccionar escape analysis

    go build -o /tmp/gocito-runtime -gcflags='-m=2' ./07_paquetes_modulos_y_herramientas/08_compilador_y_runtime

Buscá devolverPuntero, inlining, escapes to heap y moved to heap. Los mensajes
exactos pueden variar: inlinear una llamada puede cambiar la decisión en ese
sitio. No deduzcas "todo puntero va al heap" a partir de un ejemplo.

**Comprobación:** ¿el compilador cambió la semántica o solo cómo la ejecuta?

## 3. Leer dependencias de paquetes

    go list -f '{{.ImportPath}} -> {{join .Imports ", "}}' ./07_paquetes_modulos_y_herramientas/02_organizar_en_carpetas
    go doc gocito/07_paquetes_modulos_y_herramientas/02_organizar_en_carpetas/precios

La lista muestra la ruta del paquete y sus imports. El segundo comando consulta
la documentación del paquete precios. Un módulo reúne paquetes; no todo el
módulo es un único package.

**Comprobación:** ¿por qué se ejecuta una carpeta y no solo main.go cuando el
paquete contiene varios archivos? ¿Qué nombres pueden usar otros paquetes?
