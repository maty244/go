package solucion

import (
	"strings"

	"gocito/11_proyecto_cli/internal/tareas"
)

func Buscar(lista []tareas.Tarea, texto string) []tareas.Tarea {
	texto = strings.ToLower(strings.TrimSpace(texto))
	resultado := make([]tareas.Tarea, 0)
	for _, tarea := range lista {
		if strings.Contains(strings.ToLower(tarea.Titulo), texto) {
			resultado = append(resultado, tarea)
		}
	}
	return resultado
}
