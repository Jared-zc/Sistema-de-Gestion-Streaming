package main

import "fmt"

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
