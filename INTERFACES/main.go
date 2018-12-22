package main

import (
	"fmt"

	"github.com/marcovargas74/cursoDeGo/INTERFACES/model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo Da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmaPataNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
