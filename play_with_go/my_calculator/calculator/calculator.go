package calculator

import (
	"fmt"
	"my_calculator/operations"
)

type Calculator struct {
	result float64
}

func (c Calculator) Sumar(a, b float64) float64 {
	return operations.Sum(a, b)
}

func (c Calculator) Restar(a, b float64) float64 {
	return operations.Sub(a, b)
}

func (c Calculator) Multiplicar(a, b float64) float64 {
	return operations.Mult(a, b)
}

func (c Calculator) Dividir(a, b float64) float64 {
	return operations.Div(a, b)
}

func (c Calculator) recibirNumeros() []float64 {
	var ingreso float64
	var numeros []float64

	for{

		fmt.Print("Ingrese un numero, para parar, ingrese 0:")

		_, err := fmt.Scanln(&ingreso)
	
		if err != nil {
			fmt.Print("Entrada invalida.")
		}
		
		if ingreso == 0 {
			break
		} else {
			numeros = append(numeros, ingreso)
		}

		fmt.Println("Numeros ingresados:", numeros)
	}

	return numeros
}

func (c Calculator) recibirOperaciones(numerosIngresos []float64) []string {
	var ingreso string
	operacionesValidas := []string{"x", "*", "/", "-", "+"}
	var operacionesIngresos []string
	
	for i, valor := range numerosIngresos[:len(numerosIngresos)-1] {
		fmt.Print("Ingrese la operacion para estos dos numeros: ", valor, numerosIngresos[i+1])

		_, err := fmt.Scanln(&ingreso)

		valido := false

		for _, operacion := range operacionesValidas {
			if ingreso == operacion {
				valido = true
			}
		}

		if err != nil || !valido {
			fmt.Print("Entrada invalida.")
		}

		operacionesIngresos = append(operacionesIngresos, ingreso)
	}

	return operacionesIngresos
}

func (c Calculator) Calcular() float64 {
	numerosIngresos := c.recibirNumeros()

	if len(numerosIngresos) == 1{
		return numerosIngresos[0]
	} else if len(numerosIngresos) == 0{
		return 0
	}

	operacionesIngresos := c.recibirOperaciones(numerosIngresos)
	var respuesta float64

	for i, numero := range numerosIngresos[:len(numerosIngresos)-1] {
		if i == 0{
			respuesta = numero;
		}

		switch operacionesIngresos[i] {
		case "x", "*":
			respuesta = c.Multiplicar(respuesta, numerosIngresos[i+1])
		case "-":
			respuesta = c.Restar(respuesta, numerosIngresos[i+1])
		case "+":
			respuesta = c.Sumar(respuesta, numerosIngresos[i+1])
		default:
			respuesta = c.Dividir(respuesta, numerosIngresos[i+1])
		}
	}
	return respuesta

}