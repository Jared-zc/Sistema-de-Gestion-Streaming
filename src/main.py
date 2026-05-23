from usuarios import registrar_usuario

usuarios = []

usuarios = registrar_usuario(usuarios, {
    "id": 1,
    "nombre": "Ana"
})

print(usuarios)
