# Reto: streaming

Copiá texto desde un `io.Reader` hacia un `io.Writer` en mayúsculas, sin asumir
que la entrada es un archivo ni cargarla completa con `io.ReadAll`.

Usá `bufio.Reader.ReadRune` para soportar UTF-8 partido entre lecturas.
Transformá cada rune con `unicode.ToUpper`, propagá los errores con `%w` y
detectá una escritura corta. EOF termina normalmente; los bytes UTF-8 inválidos
se sustituyen por U+FFFD, como hace ReadRune. No cierres los recursos: los abrió
el llamador. Activá criterios con `go test -tags reto .`.
