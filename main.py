from usuarios import registrar_usuario, listar_usuarios
from contenido import agregar_contenido, listar_contenido
from suscripciones import crear_suscripcion, listar_suscripciones
from reproducciones import registrar_reproduccion, listar_reproducciones
from reportes import reporte_general

usuarios = []
contenidos = []
suscripciones = []
reproducciones = []

usuarios = registrar_usuario(usuarios, {
    "id": 1,
    "nombre": "Ana López",
    "correo": "ana@gmail.com"
})

contenidos = agregar_contenido(contenidos, {
    "id": 1,
    "titulo": "La aventura digital",
    "tipo": "Película",
    "genero": "Aventura"
})

suscripciones = crear_suscripcion(suscripciones, {
    "id": 1,
    "usuario_id": 1,
    "plan": "Premium",
    "estado": "Activa"
})

reproducciones = registrar_reproduccion(reproducciones, {
    "usuario_id": 1,
    "contenido_id": 1,
    "fecha": "2026-05-23"
})

print("USUARIOS:")
print(listar_usuarios(usuarios))

print("CONTENIDO:")
print(listar_contenido(contenidos))

print("SUSCRIPCIONES:")
print(listar_suscripciones(suscripciones))

print("REPRODUCCIONES:")
print(listar_reproducciones(reproducciones))

print("REPORTE GENERAL:")
print(reporte_general(usuarios, contenidos, suscripciones, reproducciones))