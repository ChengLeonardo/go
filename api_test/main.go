package main

import (
	"bytes"
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
	fmt.Println("Get:")
    // Hacer una solicitud GET
    respGet, errGet := http.Get("https://jsonplaceholder.typicode.com/posts/1")
    if errGet != nil {
        fmt.Println("Error:", errGet)
        return
    }
    defer respGet.Body.Close()  // Cerrar el cuerpo de la respuesta al final

    // Leer el cuerpo de la respuesta
    bodyGet, errGet := ioutil.ReadAll(respGet.Body)
    if errGet != nil {
        fmt.Println("Error leyendo:", errGet)
        return
    }

    // Mostrar el resultado
    fmt.Println(string(bodyGet))

	fmt.Println("Post:")

	jsonData := []byte(`{"title": "Nuevo Post", "body": "Este es un post de prueba", "userId": 1}`)

    // Hacer una solicitud POST
    resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Leer la respuesta
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error leyendo:", err)
        return
    }

    // Mostrar la respuesta
    fmt.Println("Respuesta:", string(body))

}
