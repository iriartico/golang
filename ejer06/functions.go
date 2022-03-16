//Manejo de funciones
package main

import "fmt"

func main()  {
	// fmt.Println(one(5.2)) 

	// number, status := two(10/2)
	// fmt.Println(number)
	// fmt.Println(status)

	// fmt.Println(three("Saben...")) 

	fmt.Println("resultado:", calculo(1,2,3,4,5))
	fmt.Println("resultado:", calculo(1,2,3,4))
	fmt.Println("resultado:", calculo(3,41,3,5,8,6,5,1,2,5,7))
}


func one(num float64) float64 {
	return num*2
}

//Retornando n valores por funcion
func two(num int) (int, bool) {
	if num == 5 {
		return num, false
	} else {
		return 0, true
	}
}

//Devolviendo valor de funcion con un alias
func three(cadena string) (z string) {
	z = cadena + " me la pelan"
	return 
}

//Enviando n parametros a una funcion
func calculo(numero ...int) (total int)  {
	// total := 0
	for _, num := range numero { //declarando 2 variables, retorna => index && value 
		total = total + num
	}
	return 
}