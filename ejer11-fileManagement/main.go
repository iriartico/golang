package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main()  {
	leoArchivo()
	leoArchivo2()
}

func leoArchivo()  {
	archivo, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		fmt.Println(string(archivo))
	}
}

func leoArchivo2()  {
	archivo, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {  //En esta linea ya lee el archivo y lo guarda en el buffer (Read line-to-line)
			registro := scanner.Text()
			fmt.Printf("Linea > "+registro+"\n")
		}		
	}
}