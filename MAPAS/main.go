package main

import (
	"fmt"

	"github.com/marcovargas74/cursoDeGo/MAPAS/model"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, o que achei foi : %+v \r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	casa.X = 19
	casa.Y = 26
	casa.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos item tenho no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)

	}

	fmt.Println("Quantos item tenho no cache?", len(cache))

}
