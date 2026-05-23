def registrar_usuario(lista_usuarios, nuevo_usuario):
    return lista_usuarios + [nuevo_usuario]


def listar_usuarios(lista_usuarios):
    return lista_usuarios


def buscar_usuario(lista_usuarios, usuario_id):
    return list(filter(lambda usuario: usuario["id"] == usuario_id, lista_usuarios))


def eliminar_usuario(lista_usuarios, usuario_id):
    return list(filter(lambda usuario: usuario["id"] != usuario_id, lista_usuarios))