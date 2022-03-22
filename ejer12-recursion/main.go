package main

import "fmt"

func main()  {
	exponencial(2)	
}

func exponencial(num int)  {
	if num > 1000000000 {
		return 
	}
	fmt.Println(num)
	exponencial(num * 2)
}