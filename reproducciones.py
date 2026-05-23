def registrar_reproduccion(lista_reproducciones, nueva_reproduccion):
    return lista_reproducciones + [nueva_reproduccion]


def listar_reproducciones(lista_reproducciones):
    return lista_reproducciones


def historial_usuario(lista_reproducciones, usuario_id):
    return list(filter(lambda reproduccion: reproduccion["usuario_id"] == usuario_id, lista_reproducciones))