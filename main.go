package main

import (
	"fmt"

	"github.com/luismescobarf/claseModuloGo/estudiantes"
)

func main() {
	fmt.Println("Hola desde el paquete repo")

	var estudiante1 estudiantes.Estudiante
	estudiante1.Nombre = "Juan"
	estudiante1.Edad = 20
	estudiante1.Promedio = 4.8

	fmt.Println(estudiante1.GenerarNombre("$"))

}
