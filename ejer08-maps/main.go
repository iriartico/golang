/*
%d = enteros
%s = cadenas
%t = booleanos
*/
package main

import "fmt"

func main()  {
	paises := make(map[string]string)
	paises["Mexico"]="D.F."
	paises["Bolivia"]="L.P."
	// fmt.Println(paises)


	campeonato := map[string]int {
		"Barcelona": 34,
		"Madrid": 12,
		"Grecia": 22 }
	campeonato["Rive Plate"] = 99
	campeonato["Grecia"] = 76
	// fmt.Println(campeonato)
	delete(campeonato, "Barcelona")
	// fmt.Println(campeonato)


	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo: %s	Puntaje: %d \n", equipo, puntaje)
	}


	puntaje, existe := campeonato["Grecia"]
	fmt.Printf("Puntaje capturado: %d, Existencia del equipo: %t \n", puntaje, existe)
}