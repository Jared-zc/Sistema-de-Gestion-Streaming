package main

import (
	"fmt"
	"log"
)

type Usuario struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

type Contenido struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Tipo   string `json:"tipo"`
	Genero string `json:"genero"`
}

type Suscripcion struct {
	Id        int    `json:"id"`
	UsuarioID int    `json:"usuario_id"`
	Plan      string `json:"plan"`
	Estado    string `json:"estado"`
}

type Reproduccion struct {
	UsuarioID   int    `json:"usuario_id"`
	ContenidoID int    `json:"contenido_id"`
	Fecha       string `json:"fecha"`
}

func main() {
	usuarios := []Usuario{}
	contenidos := []Contenido{}
	suscripciones := []Suscripcion{}
	reproducciones := []Reproduccion{}

	if err := cargarDatos("usuarios.json", &usuarios); err != nil {
		log.Fatalf("Error cargando usuarios: %v", err)
	}
	if err := cargarDatos("contenidos.json", &contenidos); err != nil {
		log.Fatalf("Error cargando contenidos: %v", err)
	}
	if err := cargarDatos("suscripciones.json", &suscripciones); err != nil {
		log.Fatalf("Error cargando suscripciones: %v", err)
	}
	if err := cargarDatos("reproducciones.json", &reproducciones); err != nil {
		log.Fatalf("Error cargando reproducciones: %v", err)
	}

	if len(usuarios) == 0 {
		usuarios = []Usuario{{Id: 1, Nombre: "Ana López", Correo: "ana@gmail.com"}}
	}
	if len(contenidos) == 0 {
		contenidos = []Contenido{{Id: 1, Titulo: "La aventura digital", Tipo: "Película", Genero: "Aventura"}}
	}
	if len(suscripciones) == 0 {
		suscripciones = []Suscripcion{{Id: 1, UsuarioID: 1, Plan: "Premium", Estado: "Activa"}}
	}
	if len(reproducciones) == 0 {
		reproducciones = []Reproduccion{{UsuarioID: 1, ContenidoID: 1, Fecha: "2026-05-23"}}
	}

	fmt.Println("USUARIOS:")
	for _, u := range usuarios {
		fmt.Printf("%+v\n", u)
	}

	fmt.Println("\nCONTENIDO:")
	for _, c := range contenidos {
		fmt.Printf("%+v\n", c)
	}

	fmt.Println("\nSUSCRIPCIONES:")
	for _, s := range suscripciones {
		fmt.Printf("%+v\n", s)
	}

	fmt.Println("\nREPRODUCCIONES:")
	for _, r := range reproducciones {
		fmt.Printf("%+v\n", r)
	}

	fmt.Println("\nREPORTE GENERAL:")
	reporte := reporteGeneral(usuarios, contenidos, suscripciones, reproducciones)
	for clave, valor := range reporte {
		fmt.Printf("%s: %d\n", clave, valor)
	}

	if err := guardarDatos("usuarios.json", usuarios); err != nil {
		log.Printf("No se pudo guardar usuarios: %v", err)
	}
	if err := guardarDatos("contenidos.json", contenidos); err != nil {
		log.Printf("No se pudo guardar contenidos: %v", err)
	}
	if err := guardarDatos("suscripciones.json", suscripciones); err != nil {
		log.Printf("No se pudo guardar suscripciones: %v", err)
	}
	if err := guardarDatos("reproducciones.json", reproducciones); err != nil {
		log.Printf("No se pudo guardar reproducciones: %v", err)
	}
}
