package main

import (
	"cmp"
	"context"
	"errors"
	"slices"
	"strings"
	"sync"
)

var ErrProductoNoEncontrado = errors.New("producto no encontrado")

// RepositorioProductos es el contrato que consumen los handlers. Una futura
// implementación PostgreSQL puede reemplazar a Memoria sin cambiar HTTP.
type RepositorioProductos interface {
	Listar(context.Context) ([]Producto, error)
	Buscar(context.Context, int64) (Producto, error)
	Crear(context.Context, CrearProducto) (Producto, error)
	Actualizar(context.Context, int64, ActualizarProducto) (Producto, error)
	Borrar(context.Context, int64) error
}

type Memoria struct {
	mu        sync.RWMutex
	productos map[int64]Producto
	proximoID int64
}

func NuevaMemoria() *Memoria {
	return &Memoria{productos: make(map[int64]Producto), proximoID: 1}
}

func contextoVigente(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}

func (m *Memoria) Listar(ctx context.Context) ([]Producto, error) {
	if err := contextoVigente(ctx); err != nil {
		return nil, err
	}
	m.mu.RLock()
	defer m.mu.RUnlock()

	lista := make([]Producto, 0, len(m.productos))
	for _, producto := range m.productos {
		lista = append(lista, producto)
	}
	slices.SortFunc(lista, func(a, b Producto) int { return cmp.Compare(a.ID, b.ID) })
	return lista, nil
}

func (m *Memoria) Buscar(ctx context.Context, id int64) (Producto, error) {
	if err := contextoVigente(ctx); err != nil {
		return Producto{}, err
	}
	m.mu.RLock()
	defer m.mu.RUnlock()

	producto, existe := m.productos[id]
	if !existe {
		return Producto{}, ErrProductoNoEncontrado
	}
	return producto, nil
}

func (m *Memoria) Crear(ctx context.Context, entrada CrearProducto) (Producto, error) {
	if err := contextoVigente(ctx); err != nil {
		return Producto{}, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	producto := Producto{
		ID: m.proximoID, Nombre: strings.TrimSpace(entrada.Nombre),
		PrecioCentavos: entrada.PrecioCentavos, Stock: entrada.Stock,
	}
	m.productos[producto.ID] = producto
	m.proximoID++
	return producto, nil
}

func (m *Memoria) Actualizar(ctx context.Context, id int64, entrada ActualizarProducto) (Producto, error) {
	if err := contextoVigente(ctx); err != nil {
		return Producto{}, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, existe := m.productos[id]; !existe {
		return Producto{}, ErrProductoNoEncontrado
	}
	producto := Producto{
		ID: id, Nombre: strings.TrimSpace(entrada.Nombre),
		PrecioCentavos: entrada.PrecioCentavos, Stock: entrada.Stock,
	}
	m.productos[id] = producto
	return producto, nil
}

func (m *Memoria) Borrar(ctx context.Context, id int64) error {
	if err := contextoVigente(ctx); err != nil {
		return err
	}
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, existe := m.productos[id]; !existe {
		return ErrProductoNoEncontrado
	}
	delete(m.productos, id)
	return nil
}
