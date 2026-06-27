package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Reproduccion registra el historial de visualización de un usuario.
type Reproduccion struct {
	usuarioID   int
	contenidoID int
	fecha       string
}

// ReproduccionDTO se utiliza para mostrar la reproducción en formato JSON.
type ReproduccionDTO struct {
	UsuarioID   int    `json:"usuario_id"`
	ContenidoID int    `json:"contenido_id"`
	Fecha       string `json:"fecha"`
}

// NuevaReproduccion crea una reproducción y valida que la fecha no esté vacía.
func NuevaReproduccion(usuarioID int, contenidoID int, fecha string) (Reproduccion, error) {
	if fecha == "" {
		return Reproduccion{}, errors.New("la fecha de reproducción es obligatoria")
	}

	return Reproduccion{
		usuarioID:   usuarioID,
		contenidoID: contenidoID,
		fecha:       fecha,
	}, nil
}

// MostrarInfo devuelve la fecha de la reproducción.
func (r Reproduccion) MostrarInfo() string {
	return "Reproducción registrada en fecha: " + r.fecha
}

// ConvertirReproduccionDTO convierte una reproducción encapsulada a formato JSON.
func ConvertirReproduccionDTO(r Reproduccion) ReproduccionDTO {
	return ReproduccionDTO{
		UsuarioID:   r.usuarioID,
		ContenidoID: r.contenidoID,
		Fecha:       r.fecha,
	}
}

// ObtenerReproducciones devuelve el historial de reproducciones en formato JSON.
func ObtenerReproducciones(w http.ResponseWriter, r *http.Request) {
	reproduccion, err := NuevaReproduccion(1, 1, "2026-06-14")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode([]ReproduccionDTO{
		ConvertirReproduccionDTO(reproduccion),
	})
}
