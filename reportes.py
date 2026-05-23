def reporte_general(usuarios, contenidos, suscripciones, reproducciones):
    return {
        "total_usuarios": len(usuarios),
        "total_contenidos": len(contenidos),
        "total_suscripciones": len(suscripciones),
        "total_reproducciones": len(reproducciones)
    }