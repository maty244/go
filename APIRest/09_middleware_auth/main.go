package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// =========================================================
// MIDDLEWARE DE AUTENTICACIÓN
// =========================================================
// La mayoría de las APIs tienen rutas públicas (ver el catálogo)
// y rutas protegidas (crear/borrar productos). La autenticación
// por token funciona así:
//
//   1. El cliente consigue un TOKEN (al hacer login, o una API key
//      que le diste de antemano)
//   2. En cada request protegido, lo manda en el header:
//        Authorization: Bearer eyJhbGc...
//   3. Un MIDDLEWARE verifica el token ANTES de llegar al handler:
//      - Sin token o token inválido → 401 Unauthorized, cortamos acá
//      - Token válido → el request sigue al handler
//
// En este ejemplo usamos un token FIJO simulado para enfocarnos en
// el FLUJO. En producción se usa JWT: la MISMA idea, pero el token
// lleva datos firmados criptográficamente (quién sos, hasta cuándo
// vale) y se valida con una librería externa (github.com/golang-jwt).
// Si entendés este archivo, JWT es un cambio de 10 líneas.

type respuestaError struct {
	Error string `json:"error"`
}

func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ─────────────────────────────────────────────────────────
// EL MIDDLEWARE DE AUTH
// ─────────────────────────────────────────────────────────

// En la vida real esto vendría de la config (Archivos/05: env vars),
// NUNCA hardcodeado en el código.
const tokenValido = "token-secreto-del-kiosco"

func conAuth(siguiente http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Leer el header Authorization
		auth := r.Header.Get("Authorization")
		if auth == "" {
			responderJSON(w, http.StatusUnauthorized,
				respuestaError{Error: "falta el header Authorization"})
			return // ¡CORTAR ACÁ! El handler real nunca se ejecuta
		}

		// 2. El formato esperado es "Bearer <token>"
		token, ok := strings.CutPrefix(auth, "Bearer ")
		if !ok {
			responderJSON(w, http.StatusUnauthorized,
				respuestaError{Error: "formato esperado: Bearer <token>"})
			return
		}

		// 3. Validar el token (acá: comparación simple; con JWT:
		//    verificar la firma y el vencimiento)
		if token != tokenValido {
			responderJSON(w, http.StatusUnauthorized,
				respuestaError{Error: "token inválido"})
			return
		}

		// 4. Autenticado → que siga al handler real
		siguiente.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("=== Middleware de autenticación ===")

	// ─────────────────────────────────────────────────────────
	// DOS MUX: RUTAS PÚBLICAS Y RUTAS PROTEGIDAS
	// ─────────────────────────────────────────────────────────
	// Patrón común: un mux para lo protegido, montado dentro del
	// mux principal CON el middleware puesto. Así el auth aplica
	// SOLO a las rutas que lo necesitan.

	protegido := http.NewServeMux()
	protegido.HandleFunc("POST /productos", func(w http.ResponseWriter, r *http.Request) {
		responderJSON(w, http.StatusCreated, map[string]string{"mensaje": "producto creado"})
	})
	protegido.HandleFunc("DELETE /productos/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	principal := http.NewServeMux()

	// Pública: cualquiera puede ver el catálogo
	principal.HandleFunc("GET /productos", func(w http.ResponseWriter, r *http.Request) {
		responderJSON(w, http.StatusOK, []string{"Alfajor", "Gaseosa", "Mouse"})
	})

	// Todo lo demás de /productos pasa por el middleware de auth.
	// "/" en el patrón delega el resto al mux protegido.
	principal.Handle("POST /productos", conAuth(protegido))
	principal.Handle("DELETE /productos/{id}", conAuth(protegido))

	go func() {
		if err := http.ListenAndServe(":8080", principal); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor()

	// ─────────────────────────────────────────────────────────
	// PROBANDO LOS CUATRO ESCENARIOS
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Ruta pública: sin token, funciona ===")
	pedir("GET", "http://localhost:8080/productos", "")

	fmt.Println("\n=== Ruta protegida SIN token → 401 ===")
	pedir("POST", "http://localhost:8080/productos", "")

	fmt.Println("\n=== Ruta protegida con token INVÁLIDO → 401 ===")
	pedir("POST", "http://localhost:8080/productos", "Bearer token-falso")

	fmt.Println("\n=== Ruta protegida con token VÁLIDO → pasa ===")
	pedir("POST", "http://localhost:8080/productos", "Bearer "+tokenValido)

	fmt.Println("\n=== DELETE protegido con token válido ===")
	pedir("DELETE", "http://localhost:8080/productos/42", "Bearer "+tokenValido)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  Authorization: Bearer <token>  → el header estándar de auth`)
	fmt.Println(`  El middleware corta con 401    → el handler nunca se entera`)
	fmt.Println(`  Rutas públicas vs protegidas   → el middleware se aplica por ruta`)
	fmt.Println(`  401 = no sé quién sos          → 403 = sé quién sos, no podés`)
	fmt.Println(`  JWT = misma idea + firma       → librería externa, mismo middleware`)
	fmt.Println(`  El token real va en config/env → NUNCA hardcodeado (Archivos/)`)
}

func pedir(metodo, url, authHeader string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	req, _ := http.NewRequest(metodo, url, nil)
	if authHeader != "" {
		req.Header.Set("Authorization", authHeader)
	}
	resp, err := cliente.Do(req)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	texto := strings.TrimSpace(string(body))
	if texto == "" {
		texto = "(sin body)"
	}
	fmt.Printf("  %-6s con auth=%q\n  → %s | %s\n", metodo, authHeader, resp.Status, texto)
}

func esperarServidor() {
	cliente := &http.Client{Timeout: 200 * time.Millisecond}
	for i := 0; i < 50; i++ {
		resp, err := cliente.Get("http://localhost:8080/")
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
}
