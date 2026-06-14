package main

import "errors"

// Contenido representa una película, serie o documental dentro del catálogo
type Contenido struct {
	id     int
	titulo string
	tipo   string
	genero string
}

// NuevoContenido funciona como constructor
// Valida que el título, tipo y género sean obligatorios
func NuevoContenido(id int, titulo, tipo, genero string) (Contenido, error) {
	if titulo == "" || tipo == "" || genero == "" {
		return Contenido{}, errors.New("todos los campos del contenido son obligatorios")
	}

	return Contenido{id: id, titulo: titulo, tipo: tipo, genero: genero}, nil
}

// ObtenerID permite acceder al identificador del contenido
func (c Contenido) ObtenerID() int {
	return c.id
}

// MostrarInfo devuelve los datos principales del contenido
func (c Contenido) MostrarInfo() string {
	return c.titulo + " - " + c.tipo + " - " + c.genero
}
