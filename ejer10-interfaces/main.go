package main

import "fmt"

type SerVivo interface{
	estaVivo() bool
}

type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface{
	respirar()
	comer()
	esCarnivoro() bool
	estaVivo() bool

}

/* Genero humano */
type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}

type mujer struct{
	hombre			// Heredando todo de la estructura hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer() { h.comiendo = true }
func (h *hombre) pensar() { h.pensando = true }
func (h *hombre) sexo() string { 
	if h.esHombre == true {
		return "Hombre"
	}
	return "Mujer" }
func (h *hombre) estaVivo() bool { return h.vivo }

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un(a) %s, y ya estoy respirando \n", hu.sexo())
}

/*---------------------------------------------------------------------------*/
/* Genero Animal*/
type perro struct {
	respirando bool
	comiendo bool
	carnivoro bool
	vivo bool
}

func (p *perro) respirar() { p.respirando = true }
func (p *perro) comer() { p.comiendo = true }
func (p *perro) esCarnivoro() bool { return p.carnivoro }
func (p *perro) estaVivo() bool { return p.vivo }



func AnimalRespirando(ani animal) {
	ani.respirar()
	fmt.Printf("Soy un animal y estoy respirando \n")
}

func AnimalCarnivoro(ani animal) int {
	if ani.esCarnivoro() == true {
		return 1
	}
	return 0
}

/* SER VIVO*/
func EstoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main()  {
	// Pedro := new (hombre)
	// Pedro.esHombre = true
	// HumanoRespirando(Pedro)

	// Maria := new (mujer)
	// HumanoRespirando(Maria)

	/* Animales */
	totalCarnivoros := 0
	Dog := new(perro)
	Dog.carnivoro = true
	Dog.vivo = true
	AnimalRespirando(Dog)
	totalCarnivoros =+ AnimalCarnivoro(Dog)
	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
	
	fmt.Printf("Estoy vivo = %t \n", EstoyVivo(Dog))
}