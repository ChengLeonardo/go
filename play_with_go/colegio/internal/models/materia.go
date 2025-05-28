package models

type Materia struct {
    Nombre string
    Nota   float64
}

func (m *Materia) ObtenerNota() float64 {
    return m.Nota
}