def crear_suscripcion(lista_suscripciones, nueva_suscripcion):
    return lista_suscripciones + [nueva_suscripcion]


def listar_suscripciones(lista_suscripciones):
    return lista_suscripciones


def suscripciones_activas(lista_suscripciones):
    return list(filter(lambda suscripcion: suscripcion["estado"] == "Activa", lista_suscripciones))


def cancelar_suscripcion(lista_suscripciones, usuario_id):
    return list(map(
        lambda suscripcion: 
        {**suscripcion, "estado": "Cancelada"} 
        if suscripcion["usuario_id"] == usuario_id else suscripcion,
        lista_suscripciones
    ))