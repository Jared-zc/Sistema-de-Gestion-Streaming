package main

import "errors"

// Usuario representa a una persona registrada en el sistema de streaming
// Sus atributos están en minúscula para aplicar encapsulación
type Usuario struct {
	id     int
	nombre string
	correo string
}

// NuevoUsuario funciona como constructor
// Valida que el nombre y correo no estén vacíos antes de crear el usuario
func NuevoUsuario(id int, nombre, correo string) (Usuario, error) {
	if nombre == "" || correo == "" {
		return Usuario{}, errors.New("el nombre y correo son obligatorios")
	}

	return Usuario{id: id, nombre: nombre, correo: correo}, nil
}

// ObtenerID permite acceder al identificador del usuario sin modificarlo directamente
func (u Usuario) ObtenerID() int {
	return u.id
}

// MostrarInfo devuelve la información principal del usuario para reportes
func (u Usuario) MostrarInfo() string {
	return u.nombre + " - " + u.correo
}
