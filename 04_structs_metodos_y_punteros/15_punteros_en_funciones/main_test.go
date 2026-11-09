package main

import "testing"

func TestCopiarArgumentosNoEsReapuntarAlLlamador(t *testing.T) {
	n := 10
	incrementarCopia(n)
	if n != 10 {
		t.Fatal("una modificación del parámetro int alteró el argumento")
	}
	incrementarDestino(&n)
	if n != 11 {
		t.Fatal("la escritura mediante *int no llegó a n")
	}
	a, b := 10, 20
	p := &a
	cambiarPunteroLocal(p, &b)
	if p != &a || a != 10 || b != 99 {
		t.Fatalf("p==&a=%v a=%d b=%d; quiero true, 10, 99", p == &a, a, b)
	}
}

func TestReemplazarPunteroAdmiteDestinoNil(t *testing.T) {
	var p *int // todavía no tiene destino, pero &p sí es una dirección válida
	valor := 7
	reemplazarPuntero(&p, &valor)
	if p != &valor {
		t.Fatal("no se reemplazó el puntero nil del llamador")
	}
	reemplazarPuntero(&p, nil)
	if p != nil || valor != 7 {
		t.Fatal("dejar p nil debe conservar intacto el antiguo destino")
	}
}

func TestIntercambioYNilSinMutacionParcial(t *testing.T) {
	a, b := 3, 8
	if !intercambiar(&a, &b) || a != 8 || b != 3 {
		t.Fatalf("intercambio incorrecto: a=%d b=%d", a, b)
	}
	if !intercambiar(&a, &a) || a != 8 {
		t.Fatal("dos parámetros con el mismo destino deben dejarlo intacto")
	}
	if intercambiar(&a, nil) || intercambiar(nil, &b) || intercambiar(nil, nil) || a != 8 || b != 3 {
		t.Fatal("nil debe rechazarse sin cambiar ningún int")
	}
}

func TestContadoresRetornadosSonIndependientes(t *testing.T) {
	a, b := crearContador(5), crearContador(5)
	*a = 9
	if *a != 9 || *b != 5 || a == b {
		t.Fatal("cada llamada debe crear un contador independiente y todavía válido")
	}
}
