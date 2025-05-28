# Proyecto Colegio

Este proyecto es un sistema de gestión para un colegio, diseñado para manejar la información de cursos, estudiantes y materias. Permite gestionar la cantidad de aprobados y desaprobados al final del año.

## Estructura del Proyecto

El proyecto está organizado en los siguientes directorios y archivos:

- **cmd/main.go**: Punto de entrada de la aplicación. Inicializa el sistema y gestiona la interacción con el usuario.
  
- **internal/models**: Contiene las definiciones de las estructuras principales del sistema.
  - **curso.go**: Define la estructura `Curso`, que incluye el nombre del curso y una lista de estudiantes.
  - **estudiante.go**: Define la estructura `Estudiante`, que incluye el nombre, apellido, legajo y una lista de materias.
  - **materia.go**: Define la estructura `Materia`, que contiene el nombre de la materia y la nota obtenida.

- **internal/services**: Contiene la lógica del sistema.
  - **gestion.go**: Funciones para gestionar la lógica del sistema, como agregar estudiantes y calcular aprobados y desaprobados.
  - **reporte.go**: Funciones para generar reportes sobre los estudiantes y sus notas.

- **internal/utils**: Funciones auxiliares que pueden ser utilizadas en diferentes partes del proyecto.
  - **helpers.go**: Incluye funciones de validación y formateo de datos.

- **go.mod**: Archivo de gestión de dependencias de Go.

## Instalación

Para instalar el proyecto, clona el repositorio y navega al directorio del proyecto:

```
git clone <URL_DEL_REPOSITORIO>
cd colegio
```

Luego, ejecuta el siguiente comando para descargar las dependencias:

```
go mod tidy
```

## Ejecución

Para ejecutar el sistema, utiliza el siguiente comando:

```
go run cmd/main.go
```

## Contribuciones

Las contribuciones son bienvenidas. Si deseas contribuir, por favor abre un issue o envía un pull request.