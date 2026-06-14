package main

import "errors"

// Reproduccion registra el historial de visualización de un usuario.
type Reproduccion struct {
	usuarioID   int
	contenidoID int
	fecha       string
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
