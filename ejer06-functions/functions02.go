// Funciones anonimas y closures
package main

import "fmt"

var Calculate func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main()  {
	// fmt.Printf("Resultado de la suma: %d \n", Calculate(5,7))
	
	//Modificando la funcion anonima
	// Calculate = func(num1 int, num2 int) int {
	// 	return num1 - num2
	// }
	// fmt.Printf("Resultado de la resta: %d \n", Calculate(5,7))

	//Modificando la funcion anonima
	// Calculate = func(num1 int, num2 int) int {
	// 	return num1 * num2
	// }
	// fmt.Printf("Resultado del producto: %d \n", Calculate(5,7))

	//nueva funcion
	// Operations()

	// Closures
	tablaDel2 := 2
	MiTabla := Table(tablaDel2)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
	// Table()
}

func Operations() {
	resultado := func() int {
		var a int = 23
		var b int = 45
		return a + b
	}
	fmt.Println(resultado())
}

//Closures
func Table(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}