package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Es una autenticación fija y deliberadamente simulada. JWT real, rotación
// de claves, scopes y gestión de identidades quedan fuera de esta etapa.
const tokenDemo = "token-educativo-no-usar-en-produccion"

func main() {
	repo := NuevaMemoria()
	if err := sembrar(repo); err != nil {
		fmt.Fprintln(os.Stderr, "sembrar:", err)
		os.Exit(1)
	}
	handler := NuevaAPI(repo, tokenDemo)

	if os.Getenv("SERVIDOR") == "1" {
		if err := servirManual(handler); err != nil {
			fmt.Fprintln(os.Stderr, "servidor:", err)
			os.Exit(1)
		}
		return
	}
	if err := demoAutomatica(handler); err != nil {
		fmt.Fprintln(os.Stderr, "demo:", err)
		os.Exit(1)
	}
}

func sembrar(repo RepositorioProductos) error {
	if _, err := repo.Crear(context.Background(), CrearProducto{Nombre: "Alfajor", PrecioCentavos: 80_000, Stock: 15}); err != nil {
		return err
	}
	if _, err := repo.Crear(context.Background(), CrearProducto{Nombre: "Gaseosa", PrecioCentavos: 120_000, Stock: 3}); err != nil {
		return err
	}
	return nil
}

func demoAutomatica(handler http.Handler) error {
	servidor := httptest.NewServer(handler) // puerto dinámico, sin colisiones
	defer servidor.Close()

	casos := []struct {
		metodo, ruta, body, auth string
	}{
		{http.MethodGet, "/productos", "", ""},
		{http.MethodGet, "/productos/1", "", ""},
		{http.MethodPost, "/productos", `{"nombre":"Mouse","precio_centavos":850000,"stock":20}`, ""},
		{http.MethodPost, "/productos", `{"nombre":"Mouse","precio_centavos":850000,"stock":20}`, "Bearer " + tokenDemo},
	}

	cliente := &http.Client{Timeout: 2 * time.Second}
	for _, caso := range casos {
		req, err := http.NewRequestWithContext(context.Background(), caso.metodo, servidor.URL+caso.ruta, bytes.NewBufferString(caso.body))
		if err != nil {
			return err
		}
		if caso.body != "" {
			req.Header.Set("Content-Type", "application/json")
		}
		if caso.auth != "" {
			req.Header.Set("Authorization", caso.auth)
		}
		resp, err := cliente.Do(req)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(resp.Body)
		errCerrar := resp.Body.Close()
		if err != nil || errCerrar != nil {
			return errors.Join(err, errCerrar)
		}
		fmt.Printf("%-6s %-14s → %s | %s\n", caso.metodo, caso.ruta, resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

func servirManual(handler http.Handler) error {
	direccion := os.Getenv("ADDR")
	if direccion == "" {
		direccion = ":8080"
	}

	servidor := &http.Server{
		Addr: direccion, Handler: handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	ctxSenal, detener := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer detener()
	errores := make(chan error, 1)
	go func() { errores <- servidor.ListenAndServe() }()
	fmt.Println("API educativa escuchando en", direccion)

	select {
	case err := <-errores:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctxSenal.Done():
		ctxCierre, cancelar := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelar()
		return servidor.Shutdown(ctxCierre)
	}
}
