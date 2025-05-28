package models

type Colegio struct {
	Nombre string
	Cursos []Curso
}

func (c *Colegio) ObtenerCursos() []Curso {
	return c.Cursos
}