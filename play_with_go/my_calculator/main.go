package main

import(
	"fmt"
	"my_calculator/calculator"
)

func main(){
	for {
		var ingreso string
		fmt.Println("Quiere calcular?s/n: ")
		_, err := fmt.Scanln(&ingreso)

		if err != nil {
			fmt.Println("Ingreso invalido. ")
			break
		}

		if ingreso != "s" {
			break
		}

		calculator := calculator.Calculator{}

		respuesta := calculator.Calcular()

		fmt.Println("Respuesta: ", respuesta)
	}
}