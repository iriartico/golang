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
	graboArchivo()
	graboArchivo2()
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
	archivo.Close()
}

func graboArchivo()  {
	archivo, err := os.Create("./newFile.txt")    //archivo funciona como un buffer && os.Create borra si existe el archvo con el mismo nombre 
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta linea es nueva XD")
	archivo.Close()
}

func graboArchivo2()  {
	fileName := "./newFile.txt"
	if Append(fileName, "\nEsta es una segunda linea1") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}
	return true
}