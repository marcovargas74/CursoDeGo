package main

import "fmt"

var (
	//Endereco é um valor importante e tem de ser publico
	Endereco string
	//Telefone é um valor importante e tem de ser publico
	Telefone            string
	quantidade, estoque int     //quantidade = 0
	comprou             bool    //comprou = false
	valor               float64 //valor = 0.00
	palavras            rune
)

func main() {
	teste := "Valor de Teste"
	fmt.Printf("Endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("valor: %f\r\n", valor)
	fmt.Printf("palavras: %c\r\n", palavras)
	fmt.Printf("teste: %s\r\n", teste)
}
