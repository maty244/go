# Práctica — leer, compilar y ejecutar

Todos los comandos de esta guía se ejecutan desde la raíz del repositorio.
Predecí el resultado primero; después ejecutá y explicá cualquier diferencia.

## 1. Seguir el punto de entrada

    go run ./00_inicio/01_primer_programa

La primera línea es "Hola, Go 1.26". Después main imprime comandos: mostrarlos
en pantalla no los ejecuta. Cambiá el saludo y repetí el comando.

**Comprobación:** ¿para qué sirven package main, import y func main?
¿Quién llama a main? ¿Qué cambia si el paquete deja de llamarse main?

## 2. Separar compilación de ejecución

    go build -o /tmp/gocito-inicio ./00_inicio/01_primer_programa
    /tmp/gocito-inicio

El build crea un ejecutable pero no imprime el saludo; al ejecutar el binario
sí aparece. Go run compila y ejecuta en un paso. El archivo está en /tmp para
evitar confundir un binario con código del curso.

**Comprobación:** ¿cambiar el fuente cambia un binario ya compilado?

## 3. Consultar un contrato

    go doc fmt.Println

Leé qué recibe y qué devuelve. Es legal llamar a una función ignorando sus
retornos, como hacemos aquí. Resolvé el saludo de retos/ y ejecutá:

    go test -tags reto ./00_inicio/retos

Antes de resolverlo falla; después debe pasar. El test verifica un resultado,
no ejecuta el main de la lección.
