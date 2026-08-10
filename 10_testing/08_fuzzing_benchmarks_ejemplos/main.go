package main

import "fmt"

func Invertir(texto string) string {
	runas := []rune(texto)
	for izquierda, derecha := 0, len(runas)-1; izquierda < derecha; izquierda, derecha = izquierda+1, derecha-1 {
		runas[izquierda], runas[derecha] = runas[derecha], runas[izquierda]
	}
	return string(runas)
}

func main() {
	fmt.Println(Invertir("Go y UTF-8: ñ"))
	fmt.Println("Probá go test, go test -fuzz=Fuzz -fuzztime=10s y go test -bench=.")
}
