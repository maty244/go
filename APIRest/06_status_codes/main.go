package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// =========================================================
// STATUS CODES: EL VOCABULARIO DE TU API
// =========================================================
// El status code es la PRIMERA información que recibe el cliente:
// le dice cómo salió todo ANTES de leer el body. Usarlos bien es
// lo que distingue una API profesional.
//
// Las familias:
//   2xx → salió bien
//   4xx → EL CLIENTE hizo algo mal (puede corregirse y reintentar)
//   5xx → EL SERVIDOR falló (el cliente no tiene la culpa)
//
// Los que una API REST usa el 95% del tiempo:
//
//   200 OK          → GET/PUT exitoso (con body)
//   201 Created     → POST exitoso: se creó el recurso
//   204 No Content  → DELETE exitoso: sin body que devolver
//   400 Bad Request → el request está mal armado (JSON roto,
//                      datos inválidos)
//   401 Unauthorized→ no sé quién sos (falta token / credencial)
//   403 Forbidden   → sé quién sos, pero no tenés permiso
//   404 Not Found   → el recurso no existe
//   405 Method Not..→ la ruta existe, ese método no (lo da el mux)
//   409 Conflict    → choque de estado (ej: crear algo que ya existe)
//   500 Internal... → explotó algo del lado del servidor
//
// MAL (lo hacen muchas APIs malas): devolver TODO con 200 y poner
// el error adentro del body: {"ok": false, "error": "..."} con 200.
// Rompe clientes, caches, monitoreo y a cualquier programador
// que consuma tu API.

type respuestaError struct {
	Error string `json:"error"`
}

func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func main() {
	fmt.Println("=== Status codes en acción ===")

	// Un mini recurso en memoria para demostrar cada código
	usuarios := map[string]string{"1": "Matias"}

	mux := http.NewServeMux()

	// 200: lectura exitosa
	mux.HandleFunc("GET /usuarios/{id}", func(w http.ResponseWriter, r *http.Request) {
		nombre, existe := usuarios[r.PathValue("id")]
		if !existe {
			responderJSON(w, http.StatusNotFound, respuestaError{Error: "usuario no encontrado"}) // 404
			return
		}
		responderJSON(w, http.StatusOK, map[string]string{"nombre": nombre}) // 200
	})

	// 201: creación exitosa | 409: ya existía
	mux.HandleFunc("POST /usuarios/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if _, existe := usuarios[id]; existe {
			responderJSON(w, http.StatusConflict, respuestaError{Error: "ese usuario ya existe"}) // 409
			return
		}
		usuarios[id] = "Nuevo"
		responderJSON(w, http.StatusCreated, map[string]string{"id": id}) // 201
	})

	// 204: borrado exitoso, sin body
	mux.HandleFunc("DELETE /usuarios/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if _, existe := usuarios[id]; !existe {
			responderJSON(w, http.StatusNotFound, respuestaError{Error: "usuario no encontrado"})
			return
		}
		delete(usuarios, id)
		w.WriteHeader(http.StatusNoContent) // 204: éxito SIN body
	})

	// 500: algo explotó del lado del servidor
	mux.HandleFunc("GET /explota", func(w http.ResponseWriter, r *http.Request) {
		// Simulamos un fallo interno (DB caída, bug, etc.)
		// OJO: al cliente NUNCA le mandes el detalle interno del
		// error (puede filtrar información sensible). Se loguea
		// adentro, y afuera va un mensaje genérico.
		fmt.Println("  [SERVIDOR] error interno real: conexión a DB perdida (esto va al log)")
		responderJSON(w, http.StatusInternalServerError, respuestaError{Error: "error interno"})
	})

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor()

	// ─────────────────────────────────────────────────────────
	// LA PELÍCULA COMPLETA
	// ─────────────────────────────────────────────────────────
	fmt.Println()
	pedir("GET", "http://localhost:8080/usuarios/1", "leer algo que existe")
	pedir("GET", "http://localhost:8080/usuarios/99", "leer algo que NO existe")
	pedir("POST", "http://localhost:8080/usuarios/2", "crear algo nuevo")
	pedir("POST", "http://localhost:8080/usuarios/1", "crear algo que YA existe")
	pedir("DELETE", "http://localhost:8080/usuarios/2", "borrar (éxito sin body)")
	pedir("PUT", "http://localhost:8080/usuarios/1", "método no soportado en la ruta")
	pedir("GET", "http://localhost:8080/explota", "error del servidor")

	// ─────────────────────────────────────────────────────────
	// TABLA DE DECISIÓN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Tabla de decisión rápida ===")
	fmt.Println("  ¿Salió bien un GET/PUT?        → 200 + body")
	fmt.Println("  ¿Se creó algo (POST)?          → 201 + el recurso creado")
	fmt.Println("  ¿Se borró algo (DELETE)?       → 204 sin body")
	fmt.Println("  ¿El cliente mandó fruta?       → 400 + qué corregir")
	fmt.Println("  ¿No está identificado?         → 401")
	fmt.Println("  ¿Identificado pero sin permiso?→ 403")
	fmt.Println("  ¿El recurso no existe?         → 404")
	fmt.Println("  ¿Choque de estado?             → 409")
	fmt.Println("  ¿Falló el servidor?            → 500 + mensaje GENÉRICO (detalle al log)")

	fmt.Println("\n=== RESUMEN ===")
	fmt.Println("  2xx éxito | 4xx culpa del cliente | 5xx culpa del servidor")
	fmt.Println("  NUNCA todo-200: el status ES parte de la respuesta")
	fmt.Println("  500 → detalle al log, mensaje genérico al cliente")
}

func pedir(metodo, url, descripcion string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	req, _ := http.NewRequest(metodo, url, nil)
	resp, err := cliente.Do(req)
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	texto := string(body)
	if texto == "" {
		texto = "(sin body)"
	}
	fmt.Printf("  %-32s %-6s → %s | %s", descripcion, metodo, resp.Status, texto)
	if texto == "(sin body)" {
		fmt.Println()
	}
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
