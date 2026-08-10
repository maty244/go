//go:build reto

package reto

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSalud(t *testing.T) {
	respuesta := httptest.NewRecorder()
	Salud(respuesta, httptest.NewRequest(http.MethodGet, "/salud", nil))
	if respuesta.Code != http.StatusOK || !strings.HasPrefix(respuesta.Header().Get("Content-Type"), "application/json") {
		t.Fatalf("status=%d headers=%v", respuesta.Code, respuesta.Header())
	}
	if respuesta.Body.String() != "{\"estado\":\"ok\"}\n" {
		t.Fatalf("body=%q", respuesta.Body.String())
	}
}
