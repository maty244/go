package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// =========================================================
// LEER JSON DEL BODY (POST / PUT)
// =========================================================
// Cuando el cliente CREA o ACTUALIZA algo, manda los datos como
// JSON en el body del request. El servidor tiene que:
//
//   1. DECODIFICAR el JSON a un struct (json.NewDecoder(r.Body))
//   2. VALIDAR los datos (¡nunca confiar en lo que llega!)
//   3. Responder 400 Bad Request si el body es basura,
//      con un mensaje que le sirva al cliente para corregirse
//
// Dos protecciones extra que separan una API de juguete de una
// seria:
//   - http.MaxBytesReader: límite de tamaño del body (que no te
//     manden 10GB y te tiren el servidor)
//   - DisallowUnknownFields: rechazar campos que no existen
//     (típico typo del cliente: "precip" en vez de "precio")

// crearProductoDTO: el struct que representa LO QUE EL CLIENTE
// PUEDE MANDAR. Se le dice "DTO" (Data Transfer Object) y suele
// ser distinto del modelo interno (no dejás que el cliente mande
// el ID ni campos internos).
type crearProductoDTO struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
	Stock  int     `json:"stock"`
}

type respuestaError struct {
	Error string `json:"error"`
}

func responderJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	fmt.Println("=== Leer JSON del body ===")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /productos", func(w http.ResponseWriter, r *http.Request) {
		// ─────────────────────────────────────────────────────
		// PASO 1: limitar el tamaño del body (acá, 1MB)
		// ─────────────────────────────────────────────────────
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

		// ─────────────────────────────────────────────────────
		// PASO 2: decodificar, rechazando campos desconocidos
		// ─────────────────────────────────────────────────────
		var dto crearProductoDTO
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		if err := dec.Decode(&dto); err != nil {
			// El body no es JSON válido, o trae campos que no existen
			responderJSON(w, http.StatusBadRequest,
				respuestaError{Error: "JSON inválido: " + err.Error()})
			return
		}

		// ─────────────────────────────────────────────────────
		// PASO 3: validar los DATOS (que sea JSON válido no
		// significa que tenga sentido)
		// ─────────────────────────────────────────────────────
		var problemas []string
		if strings.TrimSpace(dto.Nombre) == "" {
			problemas = append(problemas, "nombre es obligatorio")
		}
		if dto.Precio <= 0 {
			problemas = append(problemas, "precio debe ser mayor a 0")
		}
		if dto.Stock < 0 {
			problemas = append(problemas, "stock no puede ser negativo")
		}
		if len(problemas) > 0 {
			responderJSON(w, http.StatusBadRequest,
				respuestaError{Error: strings.Join(problemas, "; ")})
			return
		}

		// Todo OK → 201 Created + el recurso creado (con su ID nuevo)
		responderJSON(w, http.StatusCreated, map[string]any{
			"id":     123, // en una API real, lo genera la base de datos
			"nombre": dto.Nombre,
			"precio": dto.Precio,
			"stock":  dto.Stock,
		})
	})

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != http.ErrServerClosed {
			fmt.Println("servidor:", err)
		}
	}()
	esperarServidor()

	// ─────────────────────────────────────────────────────────
	// PROBANDO: el caso feliz y TODOS los casos de error
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== Caso feliz: JSON válido y con sentido ===")
	postear(`{"nombre": "Alfajor", "precio": 800, "stock": 15}`)

	fmt.Println("\n=== Body que no es JSON ===")
	postear(`esto no es json`)

	fmt.Println("\n=== Campo desconocido (typo del cliente) ===")
	postear(`{"nombre": "Alfajor", "precip": 800}`)

	fmt.Println("\n=== JSON válido pero datos inválidos ===")
	postear(`{"nombre": "", "precio": -50, "stock": -1}`)

	// ─────────────────────────────────────────────────────────
	// RESUMEN
	// ─────────────────────────────────────────────────────────
	fmt.Println("\n=== RESUMEN ===")
	fmt.Println(`  json.NewDecoder(r.Body).Decode(&dto) → body JSON → struct`)
	fmt.Println(`  http.MaxBytesReader                  → límite de tamaño (1<<20 = 1MB)`)
	fmt.Println(`  dec.DisallowUnknownFields()          → rechaza campos con typo`)
	fmt.Println(`  Validar SIEMPRE                      → JSON válido ≠ datos válidos`)
	fmt.Println(`  Error → 400 con mensaje útil         → el cliente puede corregirse`)
	fmt.Println(`  Éxito en POST → 201 Created          → devolviendo el recurso creado`)
}

func postear(bodyJSON string) {
	cliente := &http.Client{Timeout: 3 * time.Second}
	resp, err := cliente.Post("http://localhost:8080/productos",
		"application/json", bytes.NewBufferString(bodyJSON))
	if err != nil {
		fmt.Println("  error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("  POST body: %s\n", bodyJSON)
	fmt.Printf("  → %s | %s", resp.Status, string(body))
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
