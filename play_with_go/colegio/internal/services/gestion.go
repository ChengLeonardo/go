package services

import (
	"errors"
	"colegio/internal/models"
)

type Sistema struct {
	Colegios map[string]*models.Colegio
}

func NewSistema() *Sistema {
	return &Sistema{
		Colegios: make(map[string]*models.Colegio),
	}
}

func (s *Sistema) AgregarCurso(nombre string, colegioNombre string) {
	if _, existe := s.Colegios[colegioNombre]; existe{
		colegio := s.Colegios[colegioNombre]
		colegio.Cursos = append(colegio.Cursos, models.Curso{
			Nombre: nombre,
			Estudiantes: make([]models.Estudiante, 0),
		})
	}
}

func (s *Sistema) AgregarEstudiante(cursoNombre string, colegioNombre string, estudiante models.Estudiante) error {
	colegio := s.Colegios[colegioNombre]
	if colegio == nil {
		return errors.New("colegio no encontrado")
	}
	var curso *models.Curso
	for i := range colegio.Cursos {
		if colegio.Cursos[i].Nombre == cursoNombre {
			curso = &colegio.Cursos[i]
			break
		}
	}
	if curso == nil {
		return errors.New("curso no encontrado")
	}
	curso.Estudiantes = append(curso.Estudiantes, estudiante)
	return nil
}

func (s *Sistema) AsignarMateria(colegioNombre string, estudianteLegajo string, materia models.Materia) error {
	colegio := s.Colegios[colegioNombre]
	if colegio == nil {
		return errors.New("colegio no encontrado")
	}
	for _, curso := range colegio.Cursos {
		for _, estudiante := range curso.Estudiantes {
			if estudiante.Legajo == estudianteLegajo {
				estudiante.Materias = append(estudiante.Materias, materia)
				return nil
			}
		}
	}
	return errors.New("estudiante no encontrado")
}
