package main

import (
	"fmt"
	"strings"
	"time"
)

func main()  {
	go slowName("Jose Iriarte")  // async await
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func slowName(name string)  {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}