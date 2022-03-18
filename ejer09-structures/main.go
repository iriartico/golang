package main

import (
	"fmt"
	"time"

	us "structures/user"
)

type pepe struct {
	us.Usuario
}
func main()  {
	u := new(pepe)
	u.AltaUsuario(1, "Jose Iriarte", time.Now(), true)
	fmt.Println(u.Usuario)
}