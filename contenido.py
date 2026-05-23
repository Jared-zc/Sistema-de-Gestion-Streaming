def agregar_contenido(lista_contenido, nuevo_contenido):
    return lista_contenido + [nuevo_contenido]


def listar_contenido(lista_contenido):
    return lista_contenido


def buscar_por_genero(lista_contenido, genero):
    return list(filter(lambda contenido: contenido["genero"] == genero, lista_contenido))


def eliminar_contenido(lista_contenido, contenido_id):
    return list(filter(lambda contenido: contenido["id"] != contenido_id, lista_contenido))