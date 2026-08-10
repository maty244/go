package solucion

import "context"

func Cuadrados(ctx context.Context, entrada <-chan int) <-chan int {
	salida := make(chan int)
	go func() {
		defer close(salida)
		for {
			select {
			case <-ctx.Done():
				return
			case numero, abierto := <-entrada:
				if !abierto {
					return
				}
				select {
				case salida <- numero * numero:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return salida
}
