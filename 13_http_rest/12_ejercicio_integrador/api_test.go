package main

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const tokenTest = "token-test"

func handlerConProducto(t *testing.T) http.Handler {
	t.Helper()
	repo := NuevaMemoria()
	if _, err := repo.Crear(context.Background(), CrearProducto{Nombre: "Mouse", PrecioCentavos: 850_000, Stock: 2}); err != nil {
		t.Fatal(err)
	}
	return NuevaAPI(repo, tokenTest)
}

func peticion(handler http.Handler, metodo, ruta, body, contentType, auth string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(metodo, ruta, strings.NewReader(body))
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if auth != "" {
		req.Header.Set("Authorization", auth)
	}
	respuesta := httptest.NewRecorder()
	handler.ServeHTTP(respuesta, req)
	return respuesta
}

func TestRutasYStatus(t *testing.T) {
	t.Parallel()
	handler := handlerConProducto(t)
	jsonValido := `{"nombre":"Teclado","precio_centavos":1200000,"stock":4}`

	casos := []struct {
		nombre, metodo, ruta, body, tipo, auth string
		status                                 int
	}{
		{"listar", http.MethodGet, "/productos", "", "", "", http.StatusOK},
		{"buscar", http.MethodGet, "/productos/1", "", "", "", http.StatusOK},
		{"buscar inexistente", http.MethodGet, "/productos/99", "", "", "", http.StatusNotFound},
		{"id inválido", http.MethodGet, "/productos/no", "", "", "", http.StatusBadRequest},
		{"crear sin auth", http.MethodPost, "/productos", jsonValido, "application/json", "", http.StatusUnauthorized},
		{"tipo incorrecto", http.MethodPost, "/productos", jsonValido, "text/plain", "Bearer " + tokenTest, http.StatusUnsupportedMediaType},
		{"JSON roto", http.MethodPost, "/productos", `{`, "application/json", "Bearer " + tokenTest, http.StatusBadRequest},
		{"campo desconocido", http.MethodPost, "/productos", `{"nombre":"x","precio_centavos":1,"stock":0,"admin":true}`, "application/json", "Bearer " + tokenTest, http.StatusBadRequest},
		{"contenido extra", http.MethodPost, "/productos", jsonValido + ` {}`, "application/json", "Bearer " + tokenTest, http.StatusBadRequest},
		{"datos inválidos", http.MethodPost, "/productos", `{"nombre":"","precio_centavos":0,"stock":-1}`, "application/json", "Bearer " + tokenTest, http.StatusBadRequest},
		{"crear", http.MethodPost, "/productos", jsonValido, "application/json; charset=utf-8", "Bearer " + tokenTest, http.StatusCreated},
		{"actualizar", http.MethodPut, "/productos/1", jsonValido, "application/json", "Bearer " + tokenTest, http.StatusOK},
		{"borrar", http.MethodDelete, "/productos/1", "", "", "Bearer " + tokenTest, http.StatusNoContent},
	}

	for _, caso := range casos {
		t.Run(caso.nombre, func(t *testing.T) {
			respuesta := peticion(handler, caso.metodo, caso.ruta, caso.body, caso.tipo, caso.auth)
			if respuesta.Code != caso.status {
				t.Fatalf("status = %d; quiero %d; body=%s", respuesta.Code, caso.status, respuesta.Body.String())
			}
		})
	}
}

func TestCrearDevuelveLocation(t *testing.T) {
	t.Parallel()
	respuesta := peticion(handlerConProducto(t), http.MethodPost, "/productos",
		`{"nombre":"Teclado","precio_centavos":1200000,"stock":4}`,
		"application/json", "Bearer "+tokenTest)

	if respuesta.Code != http.StatusCreated {
		t.Fatalf("status = %d; body=%s", respuesta.Code, respuesta.Body.String())
	}
	if location := respuesta.Header().Get("Location"); location != "/productos/2" {
		t.Fatalf("Location = %q; quiero /productos/2", location)
	}
	if tipo := respuesta.Header().Get("Content-Type"); !strings.HasPrefix(tipo, "application/json") {
		t.Fatalf("Content-Type = %q", tipo)
	}
}

func TestBodyDemasiadoGrande(t *testing.T) {
	t.Parallel()
	grande := `{"nombre":"` + strings.Repeat("x", maximoBody) + `","precio_centavos":1,"stock":0}`
	respuesta := peticion(handlerConProducto(t), http.MethodPost, "/productos", grande,
		"application/json", "Bearer "+tokenTest)
	if respuesta.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d; quiero 413; body=%s", respuesta.Code, respuesta.Body.String())
	}
}

type repoEspia struct {
	err           error
	valorContexto any
}

func (r *repoEspia) Listar(context.Context) ([]Producto, error) { return nil, r.err }
func (r *repoEspia) Buscar(context.Context, int64) (Producto, error) {
	return Producto{}, r.err
}
func (r *repoEspia) Crear(ctx context.Context, _ CrearProducto) (Producto, error) {
	r.valorContexto = ctx.Value(claveContexto{})
	return Producto{ID: 7}, r.err
}
func (r *repoEspia) Actualizar(context.Context, int64, ActualizarProducto) (Producto, error) {
	return Producto{}, r.err
}
func (r *repoEspia) Borrar(context.Context, int64) error { return r.err }

type claveContexto struct{}

func TestHandlerPasaElContextoAlRepositorio(t *testing.T) {
	t.Parallel()
	repo := &repoEspia{}
	handler := NuevaAPI(repo, tokenTest)
	req := httptest.NewRequest(http.MethodPost, "/productos", bytes.NewBufferString(
		`{"nombre":"Mouse","precio_centavos":10,"stock":1}`))
	req = req.WithContext(context.WithValue(req.Context(), claveContexto{}, "presente"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenTest)

	handler.ServeHTTP(httptest.NewRecorder(), req)
	if repo.valorContexto != "presente" {
		t.Fatalf("contexto recibido = %v", repo.valorContexto)
	}
}

func TestErrorInternoNoSeExpone(t *testing.T) {
	t.Parallel()
	secreto := errors.New("contraseña-db-super-secreta")
	respuesta := peticion(NuevaAPI(&repoEspia{err: secreto}, tokenTest), http.MethodGet,
		"/productos", "", "", "")
	if respuesta.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d", respuesta.Code)
	}
	if strings.Contains(respuesta.Body.String(), secreto.Error()) {
		t.Fatalf("se filtró el error interno: %s", respuesta.Body.String())
	}
}
