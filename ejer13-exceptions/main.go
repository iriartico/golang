/*
DEFER --> Ejecuta una funcion o accion que se le indica si o si (salvaguarda)
PANIC --> Muestra un mensaje de error que se le manda
RECOVER --> Se ejecuta cuando un PANIC es detectado y toma el control del mensaje enviado

log.Fatalf --> Realiza ademas del control de un log un os.Exit(1)
*/
package main

import (
	// "fmt"
	// "io/ioutil"
	// "os"
	"log" //depuracion de sistemas 
)

func main()  {
/* Defer */
	// file := "prueba.txt"
	// f, err := os.Open(file)
	
	// defer f.Close()    

	// if err != nil {
	// 	fmt.Println("Error abriendo el archivo")
	// 	os.Exit(1)
	// }

/* Panic */
	examplePanic()
}


func examplePanic()  {
	defer func()  {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n%v", reco) // igual que Printf
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor de uno")
	}
}


