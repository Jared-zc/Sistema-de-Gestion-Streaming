package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Contenido representa una película, serie o documental dentro del catálogo
type Contenido struct {
	id     int
	titulo string
	tipo   string
	genero string
}

// ContenidoDTO se utiliza para serializar la información en formato JSON
type ContenidoDTO struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Tipo   string `json:"tipo"`
	Genero string `json:"genero"`
}

// NuevoContenido funciona como constructor
// Valida que el título, tipo y género sean obligatorios
func NuevoContenido(id int, titulo, tipo, genero string) (Contenido, error) {
	if titulo == "" || tipo == "" || genero == "" {
		return Contenido{}, errors.New("todos los campos del contenido son obligatorios")
	}

	return Contenido{
		id:     id,
		titulo: titulo,
		tipo:   tipo,
		genero: genero,
	}, nil
}

// ObtenerID permite acceder al identificador del contenido
func (c Contenido) ObtenerID() int {
	return c.id
}

// MostrarInfo devuelve los datos principales del contenido
func (c Contenido) MostrarInfo() string {
	return c.titulo + " - " + c.tipo + " - " + c.genero
}

// ConvertirContenidoDTO transforma un Contenido en un objeto listo para JSON
func ConvertirContenidoDTO(c Contenido) ContenidoDTO {
	return ContenidoDTO{
		ID:     c.id,
		Titulo: c.titulo,
		Tipo:   c.tipo,
		Genero: c.genero,
	}
}

// ObtenerContenidos es un servicio web que devuelve el catálogo en formato JSON
func ObtenerContenidos(w http.ResponseWriter, r *http.Request) {

	contenido, err := NuevoContenido(
		1,
		"La aventura digital",
		Categorias["P"],
		"Aventura",
	)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode([]ContenidoDTO{
		ConvertirContenidoDTO(contenido),
	})
}
