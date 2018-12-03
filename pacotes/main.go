package main

import (
	"fmt"

	"github.com/marcovargas74/cursoDeGo/pacotes/operadora"
	"github.com/marcovargas74/cursoDeGo/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuario do sistema
var NomeDoUsuario = "Marco"

func main() {
	fmt.Printf("Nome do usuario: %s\n\r", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Prefixo da Operadora: %s\r\n", operadora.PrefixoDaCapitalOperadora)
	fmt.Printf("Teste de Prefixo da Operadora: %s\r\n", prefixo.TesteComPrefixo)
}
