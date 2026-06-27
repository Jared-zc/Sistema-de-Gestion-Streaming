package main

import (
	"fmt"
	"net/http"
)

// Función principal del sistema
func main() {

	// ==========================
	// SISTEMA EN CONSOLA
	// ==========================

	usuarios := []Usuario{}
	contenidos := []Contenido{}
	suscripciones := []Suscripcion{}
	reproducciones := []Reproduccion{}

	usuario1, err := NuevoUsuario(1, "Ana López", "ana@gmail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	contenido1, err := NuevoContenido(1, "La aventura digital", Categorias["P"], "Aventura")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	suscripcion1, err := NuevaSuscripcion(1, usuario1.ObtenerID(), "Premium")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	reproduccion1, err := NuevaReproduccion(usuario1.ObtenerID(), contenido1.ObtenerID(), "2026-06-14")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usuarios = append(usuarios, usuario1)
	contenidos = append(contenidos, contenido1)
	suscripciones = append(suscripciones, suscripcion1)
	reproducciones = append(reproducciones, reproduccion1)

	fmt.Println("USUARIOS")
	ImprimirReporte([]Reportable{usuario1})

	fmt.Println("\nCONTENIDO")
	ImprimirReporte([]Reportable{contenido1})

	fmt.Println("\nSUSCRIPCIONES")
	ImprimirReporte([]Reportable{suscripcion1})

	fmt.Println("\nREPRODUCCIONES")
	ImprimirReporte([]Reportable{reproduccion1})

	fmt.Println()
	ReporteGeneral(usuarios, contenidos, suscripciones, reproducciones)

	// ==========================
	// SERVICIOS WEB
	// ==========================

	http.HandleFunc("/usuarios", ObtenerUsuarios)
	http.HandleFunc("/contenidos", ObtenerContenidos)
	http.HandleFunc("/suscripciones", ObtenerSuscripciones)
	http.HandleFunc("/reproducciones", ObtenerReproducciones)
	http.HandleFunc("/reportes", ObtenerReporteGeneral)

	// Nuevos servicios web
	http.HandleFunc("/estado", ObtenerEstadoSistema)
	http.HandleFunc("/planes", ObtenerPlanes)
	http.HandleFunc("/categorias", ObtenerCategorias)

	fmt.Println("\nServidor Web iniciado en http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
