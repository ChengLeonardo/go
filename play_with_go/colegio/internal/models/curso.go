package models

type Curso struct {
    Nombre     string
    Estudiantes []Estudiante
}

func (c *Curso) ObtenerEstudiantes() []Estudiante {
    return c.Estudiantes
}

func (c *Curso) AgregarEstudiante(estudiante Estudiante) {
    c.Estudiantes = append(c.Estudiantes, estudiante)
}