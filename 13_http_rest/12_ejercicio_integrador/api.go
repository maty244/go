package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

const maximoBody = 64 << 10 // 64 KiB

type api struct {
	repo  RepositorioProductos
	token string
}

type respuestaError struct {
	Error string `json:"error"`
}

func NuevaAPI(repo RepositorioProductos, token string) http.Handler {
	a := &api{repo: repo, token: token}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /productos", a.listar)
	mux.HandleFunc("GET /productos/{id}", a.buscar)
	mux.Handle("POST /productos", a.conAuth(http.HandlerFunc(a.crear)))
	mux.Handle("PUT /productos/{id}", a.conAuth(http.HandlerFunc(a.actualizar)))
	mux.Handle("DELETE /productos/{id}", a.conAuth(http.HandlerFunc(a.borrar)))
	return a.conRecover(mux)
}

func (a *api) listar(w http.ResponseWriter, r *http.Request) {
	productos, err := a.repo.Listar(r.Context())
	if err != nil {
		a.errorInterno(w, r, err)
		return
	}
	responderJSON(w, http.StatusOK, productos)
}

func (a *api) buscar(w http.ResponseWriter, r *http.Request) {
	id, ok := leerID(w, r)
	if !ok {
		return
	}
	producto, err := a.repo.Buscar(r.Context(), id)
	if errors.Is(err, ErrProductoNoEncontrado) {
		responderError(w, http.StatusNotFound, "producto no encontrado")
		return
	}
	if err != nil {
		a.errorInterno(w, r, err)
		return
	}
	responderJSON(w, http.StatusOK, producto)
}

func (a *api) crear(w http.ResponseWriter, r *http.Request) {
	var entrada CrearProducto
	if !decodificarEstricto(w, r, &entrada) {
		return
	}
	if err := validarProducto(entrada.Nombre, entrada.PrecioCentavos, entrada.Stock); err != nil {
		responderError(w, http.StatusBadRequest, "datos de producto inválidos")
		return
	}
	producto, err := a.repo.Crear(r.Context(), entrada)
	if err != nil {
		a.errorInterno(w, r, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/productos/%d", producto.ID))
	responderJSON(w, http.StatusCreated, producto)
}

func (a *api) actualizar(w http.ResponseWriter, r *http.Request) {
	id, ok := leerID(w, r)
	if !ok {
		return
	}
	var entrada ActualizarProducto
	if !decodificarEstricto(w, r, &entrada) {
		return
	}
	if err := validarProducto(entrada.Nombre, entrada.PrecioCentavos, entrada.Stock); err != nil {
		responderError(w, http.StatusBadRequest, "datos de producto inválidos")
		return
	}
	producto, err := a.repo.Actualizar(r.Context(), id, entrada)
	if errors.Is(err, ErrProductoNoEncontrado) {
		responderError(w, http.StatusNotFound, "producto no encontrado")
		return
	}
	if err != nil {
		a.errorInterno(w, r, err)
		return
	}
	responderJSON(w, http.StatusOK, producto)
}

func (a *api) borrar(w http.ResponseWriter, r *http.Request) {
	id, ok := leerID(w, r)
	if !ok {
		return
	}
	err := a.repo.Borrar(r.Context(), id)
	if errors.Is(err, ErrProductoNoEncontrado) {
		responderError(w, http.StatusNotFound, "producto no encontrado")
		return
	}
	if err != nil {
		a.errorInterno(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodificarEstricto(w http.ResponseWriter, r *http.Request, destino any) bool {
	tipo, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || tipo != "application/json" {
		responderError(w, http.StatusUnsupportedMediaType, "Content-Type debe ser application/json")
		return false
	}

	r.Body = http.MaxBytesReader(w, r.Body, maximoBody)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(destino); err != nil {
		var grande *http.MaxBytesError
		if errors.As(err, &grande) {
			responderError(w, http.StatusRequestEntityTooLarge, "body demasiado grande")
			return false
		}
		responderError(w, http.StatusBadRequest, "JSON inválido")
		return false
	}

	// Un segundo valor también es JSON válido, pero no es un body válido para
	// esta API. Solo whitespace seguido de EOF está permitido.
	var extra any
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		responderError(w, http.StatusBadRequest, "el body debe contener un único objeto JSON")
		return false
	}
	return true
}

func leerID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		responderError(w, http.StatusBadRequest, "id inválido")
		return 0, false
	}
	return id, true
}

func responderJSON(w http.ResponseWriter, status int, dato any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(dato); err != nil {
		slog.Error("codificar respuesta", "error", err)
	}
}

func responderError(w http.ResponseWriter, status int, mensaje string) {
	responderJSON(w, status, respuestaError{Error: mensaje})
}

func (a *api) errorInterno(w http.ResponseWriter, r *http.Request, err error) {
	// El detalle va al log del servidor. El cliente recibe un mensaje fijo.
	slog.ErrorContext(r.Context(), "falló el repositorio", "ruta", r.URL.Path, "error", err)
	responderError(w, http.StatusInternalServerError, "error interno")
}

func (a *api) conAuth(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !ok || token == "" || token != a.token {
			responderError(w, http.StatusUnauthorized, "autenticación requerida")
			return
		}
		siguiente.ServeHTTP(w, r)
	})
}

func (a *api) conRecover(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if valor := recover(); valor != nil {
				slog.ErrorContext(r.Context(), "panic en handler", "panic", valor)
				responderError(w, http.StatusInternalServerError, "error interno")
			}
		}()
		siguiente.ServeHTTP(w, r)
	})
}
