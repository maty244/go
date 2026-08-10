package main

import (
	"context"
	"errors"
	"sync"
	"testing"
)

func TestMemoriaEsSeguraParaConcurrencia(t *testing.T) {
	t.Parallel()
	repo := NuevaMemoria()
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Go(func() {
			_, err := repo.Crear(context.Background(), CrearProducto{
				Nombre: "producto", PrecioCentavos: Dinero(i + 1), Stock: i,
			})
			if err != nil {
				t.Errorf("Crear: %v", err)
			}
		})
	}
	wg.Wait()

	lista, err := repo.Listar(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(lista) != 100 {
		t.Fatalf("len(Listar) = %d; quiero 100", len(lista))
	}
}

func TestMemoriaRespetaContextoCancelado(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := NuevaMemoria().Listar(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Listar = %v; quiero context.Canceled", err)
	}
}
