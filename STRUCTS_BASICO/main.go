package main

import (
	"fmt"
)

//Imovel é uma struct do tipo imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A Casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 760000}
	fmt.Printf("O Apartamento é: %+v\r\n", apartamento)

	sitio := Imovel{
		Y:     85,
		Nome:  "Sitio",
		X:     22,
		valor: 55,
	}
	fmt.Printf("O sitio é: %+v\r\n", sitio)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 450000
	casa.X = 18
	casa.Y = 31
	casa.X = 19
	fmt.Printf("A Casa Com Valor é: %+v\r\n", casa)

}
