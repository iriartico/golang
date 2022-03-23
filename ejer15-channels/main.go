/*
Los Channels permiten enviar informacion a otra funcion por medio de canales precisamente para que cada
ejecucuon paralela que se desarrolle pueda dialogar con otra 
*/

package main

import (
	"fmt"
	"time"
)

func main()  {
	channel01 := make(chan time.Duration) // data type = time.Duration
	go bucle(channel01)					// async
	fmt.Println("Llegue hasta aqui")

	msg := <- channel01					// await
	fmt.Println(msg)	
}

func bucle(chan01 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {
		
	}
	final := time.Now()
	chan01 <- final.Sub(inicio)
}