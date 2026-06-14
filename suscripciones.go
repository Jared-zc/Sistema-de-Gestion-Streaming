package main

import "errors"

// Suscripcion representa el plan contratado por un usuario
type Suscripcion struct {
	id        int
	usuarioID int
	plan      string
	estado    string
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
