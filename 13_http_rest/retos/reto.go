package reto

import "net/http"

func Salud(w http.ResponseWriter, r *http.Request) {
	// TODO: GET => 200 {"estado":"ok"}; otros métodos => 405.
}
