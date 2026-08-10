// Package almacenjson adapta el gestor de tareas a un archivo JSON.
package almacenjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gocito/11_proyecto_cli/internal/tareas"
)

type Almacen struct{ Ruta string }

func (a Almacen) Cargar() (tareas.Estado, error) {
	datos, err := os.ReadFile(a.Ruta)
	if errors.Is(err, os.ErrNotExist) {
		return tareas.Estado{}, nil
	}
	if err != nil {
		return tareas.Estado{}, fmt.Errorf("leer tareas: %w", err)
	}

	var estado tareas.Estado
	if err := json.Unmarshal(datos, &estado); err != nil {
		return tareas.Estado{}, fmt.Errorf("decodificar tareas: %w", err)
	}
	return estado, nil
}

// Guardar escribe en un temporal del mismo directorio y recién al final hace
// Rename. De ese modo, un fallo a mitad de escritura no corta el JSON anterior.
func (a Almacen) Guardar(estado tareas.Estado) (errRetorno error) {
	datos, err := json.MarshalIndent(estado, "", "  ")
	if err != nil {
		return fmt.Errorf("codificar tareas: %w", err)
	}
	datos = append(datos, '\n')

	directorio := filepath.Dir(a.Ruta)
	if err := os.MkdirAll(directorio, 0o755); err != nil {
		return fmt.Errorf("crear directorio: %w", err)
	}

	temporal, err := os.CreateTemp(directorio, ".tareas-*.tmp")
	if err != nil {
		return fmt.Errorf("crear temporal: %w", err)
	}
	nombreTemporal := temporal.Name()
	defer func() {
		// Es intencional ignorar ErrNotExist: Rename ya movió el temporal.
		_ = os.Remove(nombreTemporal)
	}()

	cerrarConError := func(causa error) error {
		if errCerrar := temporal.Close(); errCerrar != nil {
			return errors.Join(causa, errCerrar)
		}
		return causa
	}

	if err := temporal.Chmod(0o600); err != nil {
		return cerrarConError(fmt.Errorf("proteger temporal: %w", err))
	}
	if _, err := temporal.Write(datos); err != nil {
		return cerrarConError(fmt.Errorf("escribir temporal: %w", err))
	}
	if err := temporal.Sync(); err != nil {
		return cerrarConError(fmt.Errorf("sincronizar temporal: %w", err))
	}
	if err := temporal.Close(); err != nil {
		return fmt.Errorf("cerrar temporal: %w", err)
	}
	if err := os.Rename(nombreTemporal, a.Ruta); err != nil {
		return fmt.Errorf("reemplazar archivo: %w", err)
	}

	// Sincronizar el directorio hace durable el Rename en sistemas que lo
	// soportan. Si la plataforma no lo permite, el archivo ya quedó completo.
	dir, err := os.Open(directorio)
	if err != nil {
		return fmt.Errorf("abrir directorio: %w", err)
	}
	defer func() { errRetorno = errors.Join(errRetorno, dir.Close()) }()
	if err := dir.Sync(); err != nil {
		return fmt.Errorf("sincronizar directorio: %w", err)
	}
	return nil
}
