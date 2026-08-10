// Package tareas contiene la lógica de negocio del gestor. No sabe nada de
// archivos, flags ni terminal: por eso se puede probar sin efectos externos.
package tareas

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrTituloVacio  = errors.New("el título no puede estar vacío")
	ErrNoEncontrada = errors.New("tarea no encontrada")
)

type Tarea struct {
	ID           int64     `json:"id"`
	Titulo       string    `json:"titulo"`
	Completada   bool      `json:"completada"`
	CreadaEn     time.Time `json:"creada_en"`
	CompletadaEn time.Time `json:"completada_en,omitzero"`
}

// Estado es el formato estable que cruza el límite de persistencia.
type Estado struct {
	SiguienteID int64   `json:"siguiente_id"`
	Tareas      []Tarea `json:"tareas"`
}

type Gestor struct{ estado Estado }

func Nuevo(estado Estado) *Gestor {
	estado.Tareas = append([]Tarea(nil), estado.Tareas...)
	if estado.SiguienteID < 1 {
		estado.SiguienteID = 1
	}
	for _, tarea := range estado.Tareas {
		if tarea.ID >= estado.SiguienteID {
			estado.SiguienteID = tarea.ID + 1
		}
	}
	return &Gestor{estado: estado}
}

func (g *Gestor) Agregar(titulo string, ahora time.Time) (Tarea, error) {
	titulo = strings.TrimSpace(titulo)
	if titulo == "" {
		return Tarea{}, ErrTituloVacio
	}

	tarea := Tarea{ID: g.estado.SiguienteID, Titulo: titulo, CreadaEn: ahora}
	g.estado.SiguienteID++
	g.estado.Tareas = append(g.estado.Tareas, tarea)
	return tarea, nil
}

func (g *Gestor) Completar(id int64, ahora time.Time) (Tarea, error) {
	for i := range g.estado.Tareas {
		if g.estado.Tareas[i].ID != id {
			continue
		}
		g.estado.Tareas[i].Completada = true
		g.estado.Tareas[i].CompletadaEn = ahora
		return g.estado.Tareas[i], nil
	}
	return Tarea{}, ErrNoEncontrada
}

func (g *Gestor) Listar() []Tarea {
	return append([]Tarea(nil), g.estado.Tareas...)
}

func (g *Gestor) Estado() Estado {
	estado := g.estado
	estado.Tareas = append([]Tarea(nil), g.estado.Tareas...)
	return estado
}
