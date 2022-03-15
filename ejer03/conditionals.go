package main

import "fmt"

var status bool = true

func main()  {
	//Condicional if-else-else if
	if status=false; status==true {
		fmt.Println(status)
	}else if status==false{
		fmt.Println("The value is false")
	}else{
		fmt.Println("error")
	}

	//Condicional switch
	switch numero:=5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	fmt.Println(4)
	case 5:
		fmt.Println(55)
	fmt.Println(55)
	default:
		fmt.Println("Mayor a 5")
	}

}