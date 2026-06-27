package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Reportable es una interfaz que permite mostrar información de diferentes módulos
type Reportable interface {
	MostrarInfo() string
}

// ImprimirReporte recibe elementos que cumplen la interfaz Reportable
// Esto permite aplicar polimorfismo usando usuarios, contenidos, suscripciones y reproducciones
func ImprimirReporte(items []Reportable) {
	for _, item := range items {
		fmt.Println(item.MostrarInfo())
	}
}

// ReporteGeneral muestra el total de elementos registrados en cada módulo
func ReporteGeneral(usuarios []Usuario, contenidos []Contenido, suscripciones []Suscripcion, reproducciones []Reproduccion) {
	fmt.Println("REPORTE GENERAL")
	fmt.Println("Total de usuarios:", len(usuarios))
	fmt.Println("Total de contenidos:", len(contenidos))
	fmt.Println("Total de suscripciones:", len(suscripciones))
	fmt.Println("Total de reproducciones:", len(reproducciones))
}

// ObtenerReporteGeneral devuelve un reporte general en formato JSON
func ObtenerReporteGeneral(w http.ResponseWriter, r *http.Request) {
	reporte := map[string]int{
		"total_usuarios":       1,
		"total_contenidos":     1,
		"total_suscripciones":  1,
		"total_reproducciones": 1,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reporte)
}
