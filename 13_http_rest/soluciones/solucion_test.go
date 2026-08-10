package solucion

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSalud(t *testing.T) {
	respuesta := httptest.NewRecorder()
	Salud(respuesta, httptest.NewRequest(http.MethodGet, "/", nil))
	if respuesta.Code != http.StatusOK {
		t.Fatalf("status=%d", respuesta.Code)
	}
}
