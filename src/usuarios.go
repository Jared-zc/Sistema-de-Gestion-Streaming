package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Usuario representa a una persona registrada en el sistema de streaming
// Sus atributos están en minúscula para aplicar encapsulación
type Usuario struct {
	id     int
	nombre string
	correo string
}

// UsuarioDTO se usa para serializar y mostrar usuarios en formato JSON
type UsuarioDTO struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
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

// ConvertirUsuarioDTO transforma un Usuario encapsulado en un formato visible para JSON
func ConvertirUsuarioDTO(u Usuario) UsuarioDTO {
	return UsuarioDTO{
		ID:     u.id,
		Nombre: u.nombre,
		Correo: u.correo,
	}
}

// ObtenerUsuarios es un servicio web que devuelve usuarios en formato JSON
func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	usuario, err := NuevoUsuario(1, "Ana López", "ana@gmail.com")
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]UsuarioDTO{ConvertirUsuarioDTO(usuario)})
}
