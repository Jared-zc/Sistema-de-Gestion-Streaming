package main

import "fmt"

type Usuario struct {
	Id     int
	Nombre string
	Correo string
}

type Contenido struct {
	Id     int
	Titulo string
	Tipo   string
	Genero string
}

type Suscripcion struct {
	Id        int
	UsuarioID int
	Plan      string
	Estado    string
}

type Reproduccion struct {
	UsuarioID   int
	ContenidoID int
	Fecha       string
}

func main() {
	usuarios := []Usuario{
		{Id: 1, Nombre: "Ana López", Correo: "ana@gmail.com"},
	}

	contenidos := []Contenido{
		{Id: 1, Titulo: "La aventura digital", Tipo: "Película", Genero: "Aventura"},
	}

	suscripciones := []Suscripcion{
		{Id: 1, UsuarioID: 1, Plan: "Premium", Estado: "Activa"},
	}

	reproducciones := []Reproduccion{
		{UsuarioID: 1, ContenidoID: 1, Fecha: "2026-05-23"},
	}

	fmt.Println("USUARIOS:")
	for _, u := range usuarios {
		fmt.Printf("%+v\n", u)
	}

	fmt.Println("CONTENIDO:")
	for _, c := range contenidos {
		fmt.Printf("%+v\n", c)
	}

	fmt.Println("SUSCRIPCIONES:")
	for _, s := range suscripciones {
		fmt.Printf("%+v\n", s)
	}

	fmt.Println("REPRODUCCIONES:")
	for _, r := range reproducciones {
		fmt.Printf("%+v\n", r)
	}

	fmt.Println("REPORTE GENERAL:")
	fmt.Printf("Total usuarios: %d\nTotal contenidos: %d\nTotal suscripciones: %d\nTotal reproducciones: %d\n",
		len(usuarios), len(contenidos), len(suscripciones), len(reproducciones))
}
