package main

import (
	"fmt"
	"strconv" //libreria para conversiones y parseos de * a tipo string
)

var numero int //inicializa en 0
var cadena string //inicializa en cadena vacia ""
var status bool = true //incializa en false

func main()  {
	fmt.Println("Hello world")
	num2, num3, num4, letter, status := 2.55,3,4,"Aqui estoy", false
	// letter = "2"
	var newNumero float32 = 5.7
	fmt.Println(newNumero)

	//casteando el valor de newNumero a entero
	// num2 = int(newNumero) 

	//Convirtiendo el valor de num2 a variable con fmt
	// letter = fmt.Sprintf("%d", num2)

	//Convirtiendo el valor de num2 a variable con strconv; Itoa solo convierte numeros tipo int, no int32, no int64 ...
	letter = strconv.Itoa(int(num2))
	
	fmt.Println(status)
	// fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(letter)

	mostrarStatus()
}

func mostrarStatus()  {
	fmt.Println(status)
}