package main

import (
	"encoding/json"
	"fmt"

	"github.com/marcovargas74/cursoDeGo/STRUCTS_AVANCADO/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("A Casa é: %p - %+v\r\n", &casa, casa)
	fmt.Printf("O Valor da Casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Printf("A Casa em JSON: %s\r\n", string(objJSON))
}
