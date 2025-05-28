package main

import (
    "fmt"
    "colegio/internal/models"
    "colegio/internal/services"
)


func main() {
    fmt.Println("Bienvenido al sistema de gestión del colegio")
    sistema := *services.NewSistema()
    for {
        var opcion int
        fmt.Println("Seleccione una opcion: ")
        fmt.Println("1. Agregar Colegio")
        fmt.Println("2. Agregar Curso")
        fmt.Println("3. Agregar Estudiante")
        fmt.Println("4. Asignar Materia a Estudiante")
        fmt.Println("5. Mostrar Colegios")
        fmt.Println("6. Mostrar Cursos")
        fmt.Println("7. Mostrar Estudiantes")
        fmt.Println("8. Salir")
        fmt.Scanln(&opcion)
        switch opcion {
        case 1:
            var nombre string
            fmt.Println("Ingrese el nombre del colegio:")
            fmt.Scanln(&nombre)
            colegio := models.Colegio{Nombre: nombre}
            sistema.Colegios[colegio.Nombre] = &colegio
        case 2:
            var nombreCurso, nombreColegio string
            fmt.Println("Ingrese el nombre del curso:")
            fmt.Scanln(&nombreCurso)
            fmt.Println("Ingrese el nombre del colegio:")
            fmt.Scanln(&nombreColegio)
            sistema.AgregarCurso(nombreCurso, nombreColegio)
        case 3:
            var nombreEstudiante, apellidoEstudiante, legajo, nombreCurso, nombreColegio string
            fmt.Println("Ingrese el nombre del estudiante:")
            fmt.Scanln(&nombreEstudiante)
            fmt.Println("Ingrese el apellido del estudiante:")
            fmt.Scanln(&apellidoEstudiante)
            fmt.Println("Ingrese el legajo del estudiante:")
            fmt.Scanln(&legajo)
            fmt.Println("Ingrese el nombre del curso:")
            fmt.Scanln(&nombreCurso)
            fmt.Println("Ingrese el nombre del colegio:")
            fmt.Scanln(&nombreColegio)
            estudiante := models.Estudiante{
                Nombre:  nombreEstudiante,
                Apellido: apellidoEstudiante,
                Legajo:  legajo,
            }
            err := sistema.AgregarEstudiante(nombreCurso, nombreColegio, estudiante)
            if err != nil {
                fmt.Println("Error al agregar estudiante:", err)
                continue
            }
        case 4:
            var nombreMateria, colegioNombre, legajoEstudiante string
            fmt.Println("Ingrese el nombre de la materia:")
            fmt.Scanln(&nombreMateria)
            fmt.Println("Ingrese el nombre del colegio:")
            fmt.Scanln(&colegioNombre)
            fmt.Println("Ingrese el legajo del estudiante:")
            fmt.Scanln(&legajoEstudiante)
            materia := models.Materia{Nombre: nombreMateria}
            err := sistema.AsignarMateria(colegioNombre, legajoEstudiante, materia)
            if err != nil {
                fmt.Println("Error al asignar materia:", err)
                continue
            }
        case 5:
            fmt.Println("Colegios:")
            for nombre := range sistema.Colegios {
                fmt.Printf("- %s\n", nombre)
            }
        case 6:
            fmt.Println("Ingrese el nombre del colegio para mostrar sus cursos:")
            var nombreColegio string
            fmt.Scanln(&nombreColegio)
            colegio, existe := sistema.Colegios[nombreColegio]
            if !existe {
                fmt.Println("Colegio no encontrado.")
                continue
            }
            fmt.Printf("Cursos del colegio %s:\n", colegio.Nombre)
            for _, curso := range colegio.ObtenerCursos() {
                fmt.Printf("- %s\n", curso.Nombre)
            }
        case 7:
            fmt.Println("Ingrese el nombre del colegio para mostrar sus estudiantes:")
            var nombreColegioEstudiantes string
            fmt.Scanln(&nombreColegioEstudiantes)
            colegio, existe := sistema.Colegios[nombreColegioEstudiantes]
            if !existe {
                fmt.Println("Colegio no encontrado.")
                continue
            }
            fmt.Printf("Estudiantes del colegio %s:\n", colegio.Nombre)
            for _, curso := range colegio.ObtenerCursos() {
                for _, estudiante := range curso.Estudiantes {
                    fmt.Printf("- %s %s (Legajo: %s)\n", estudiante.Nombre, estudiante.Apellido, estudiante.Legajo)
                }
            }
            continue
        case 8:
            fmt.Println("Saliendo del sistema...")
            return
        default:
            fmt.Println("Opción no válida, por favor intente de nuevo.")
            continue
            }   
    }
}