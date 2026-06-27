package main

import (
	"encoding/json"
	"net/http"
)

// ObtenerEstadoSistema indica si el sistema está funcionando correctamente.
func ObtenerEstadoSistema(w http.ResponseWriter, r *http.Request) {
	estado := map[string]string{
		"sistema": "Sistema de Gestión de Streaming",
		"estado":  "Activo",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estado)
}

// ObtenerPlanes devuelve los planes disponibles de suscripción.
func ObtenerPlanes(w http.ResponseWriter, r *http.Request) {
	planes := []map[string]string{
		{"plan": "Básico", "precio": "$5.00"},
		{"plan": "Estándar", "precio": "$8.00"},
		{"plan": "Premium", "precio": "$12.00"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planes)
}

// ObtenerCategorias devuelve las categorías de contenido disponibles.
func ObtenerCategorias(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categorias)
}
