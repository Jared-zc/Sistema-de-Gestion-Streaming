package main

func reporteGeneral(usuarios []Usuario, contenidos []Contenido, suscripciones []Suscripcion, reproducciones []Reproduccion) map[string]int {
	return map[string]int{
		"total_usuarios":       len(usuarios),
		"total_contenidos":     len(contenidos),
		"total_suscripciones":  len(suscripciones),
		"total_reproducciones": len(reproducciones),
	}
}
