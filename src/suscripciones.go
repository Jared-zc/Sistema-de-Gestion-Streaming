package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Suscripcion representa el plan contratado por un usuario
type Suscripcion struct {
	id        int
	usuarioID int
	plan      string
	estado    string
}

// SuscripcionDTO se utiliza para enviar la información en formato JSON
type SuscripcionDTO struct {
	ID        int    `json:"id"`
	UsuarioID int    `json:"usuario_id"`
	Plan      string `json:"plan"`
	Estado    string `json:"estado"`
}

// NuevaSuscripcion crea una suscripción y valida que el plan no esté vacío.
func NuevaSuscripcion(id int, usuarioID int, plan string) (Suscripcion, error) {
	if plan == "" {
		return Suscripcion{}, errors.New("el plan de suscripción es obligatorio")
	}

	return Suscripcion{
		id:        id,
		usuarioID: usuarioID,
		plan:      plan,
		estado:    "Activa",
	}, nil
}

// MostrarInfo devuelve el plan y estado actual de la suscripción.
func (s Suscripcion) MostrarInfo() string {
	return s.plan + " - " + s.estado
}

// ConvertirSuscripcionDTO convierte una suscripción encapsulada a un objeto JSON
func ConvertirSuscripcionDTO(s Suscripcion) SuscripcionDTO {
	return SuscripcionDTO{
		ID:        s.id,
		UsuarioID: s.usuarioID,
		Plan:      s.plan,
		Estado:    s.estado,
	}
}

// ObtenerSuscripciones devuelve las suscripciones en formato JSON
func ObtenerSuscripciones(w http.ResponseWriter, r *http.Request) {

	suscripcion, err := NuevaSuscripcion(1, 1, "Premium")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode([]SuscripcionDTO{
		ConvertirSuscripcionDTO(suscripcion),
	})
}
