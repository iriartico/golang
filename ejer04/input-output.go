package main

import (
	"fmt"
	"os" //para errores con el Sistema Operativo
	"bufio" //Relacionado con la entrada y salida de datos
)

var num1, num2 int
var leyenda string
func main()  {
	fmt.Println("Ingrese un numero entero: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Ingrese un numero entero: ")
	fmt.Scanln(&num2)

	fmt.Println("Descipcion : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
	leyenda = scanner.Text()
	fmt.Println(leyenda, num1+num2)
	}
}