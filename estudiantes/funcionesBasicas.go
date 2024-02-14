package estudiantes

// TAD -> Clase en Go
type Estudiante struct {
	Nombre   string
	Edad     int8
	Promedio float32
}

//Métodos o funcionalidades

// 1) No requiere actualización ninguno de los att
func (e Estudiante) GenerarNombre(caracter string) string {
	var mensajeNombre string

	mensajeNombre += caracter + " " + e.Nombre + " " + caracter

	return mensajeNombre

}

//2) Implican una actualización en sus atributos
