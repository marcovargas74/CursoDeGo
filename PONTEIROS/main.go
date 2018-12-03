package main

import "fmt"

//Imovel é uma struct do tipo imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("A Casa é: %p - %+v\r\n", &casa, casa)

	sitio := Imovel{17, 28, "O SITIO", 280000}
	fmt.Printf("O Sitio é: %p - %+v\r\n", &sitio, sitio)

	mudaImovel(&sitio)
	fmt.Printf("O Sitio é: %p - %+v\r\n", &sitio, sitio)

}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
