package main

// reporteGeneral returns a map with totals for provided slices.
func reporteGeneral(usuarios, contenidos, suscripciones, reproducciones []interface{}) map[string]int {
	return map[string]int{
		"total_usuarios":       len(usuarios),
		"total_contenidos":     len(contenidos),
		"total_suscripciones":  len(suscripciones),
		"total_reproducciones": len(reproducciones),
	}
}
