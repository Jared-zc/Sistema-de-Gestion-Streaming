package main

import (
	"encoding/json"
	"os"
)

func guardarDatos(nombreArchivo string, datos interface{}) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	encoder := json.NewEncoder(archivo)
	encoder.SetIndent("", "    ")
	return encoder.Encode(datos)
}

func cargarDatos(nombreArchivo string, destino interface{}) error {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer archivo.Close()

	return json.NewDecoder(archivo).Decode(destino)
}
