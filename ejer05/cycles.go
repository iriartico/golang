package main

import "fmt"

func main()  {
	// i := 1 
	// for i<10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }	

	// num := 0
	// for {
	// 	fmt.Println("Ingrese un numero secreto")
	// 	fmt.Scanln(&num)
	// 	if num == 100 {
	// 		fmt.Println("Welcome")
	// 		break
	// 	}
	// }

	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n Valor de i: %d", i) // Printf no contiene \n y println si
	// 	if i == 5 {
	// 		fmt.Printf(" multiplicamos por 2 \n")
	// 		i=i*2
	// 		continue
	// 	}
	// 	fmt.Printf("	Paso por aqui \n")
	// 	i++
	// }

	var i int = 0
	RUTINA:
	fmt.Println("Hello world")
	fmt.Println("Mask 255.255.255.0")
		for i < 10 {
			if i == 4 {
				i=i+2
				fmt.Println("voy a RUTINA sumando 2 a i")
				goto RUTINA
			}
			fmt.Println("Valor de i = ", i)
			i++
		}
}