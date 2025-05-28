package models

type Estudiante struct {
    Nombre   string
    Apellido string
    Legajo   string
    Materias []Materia
}

func (e *Estudiante) ObtenerNotas() map[string]float64 {
    notas := make(map[string]float64)
    for _, materia := range e.Materias {
        notas[materia.Nombre] = materia.Nota
    }
    return notas
}

func (e *Estudiante) ObtenerListaMaterias() []string {
    lista := make([]string, len(e.Materias))
    for i, materia := range e.Materias {
        lista[i] = materia.Nombre
    }
    return lista
}