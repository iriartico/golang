/*
Los Middlewares son interceptores que permiten ejecutar instrucciones comunes a varias funciones que 
reciben y devuelven los mismos tipos de variables
*/

package main

import "fmt"

var result int
func main()  {
	fmt.Println("Inicio")
	
	result = operationsMidd(add)(10,3)
	fmt.Println(result)

	result = operationsMidd(subtract)(10,3)
	fmt.Println(result)

	result = operationsMidd(multiply)(10,3)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func operationsMidd(f func(int, int) int) func(int, int) int {
	return func (a, b int) int {
		fmt.Println("inicio de operacion:")
		return f(a,b)
	}
}