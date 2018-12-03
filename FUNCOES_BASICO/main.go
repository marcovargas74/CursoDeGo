package main

import (
	"fmt"

	"github.com/marcovargas74/cursoDeGo/FUNCOES_BASICO/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("Resultado Multiplicacao: %d \r\n", resultado)

	resultado = matematica.Soma(3, 3)
	fmt.Printf("Resultado da Soma: %d \r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("Resultado Divisor: %d \r\n", resultado)

	resultado, resto := matematica.DivisorComResto(14, 3)
	fmt.Printf("Resultado Divisor: %d  Resto da Divisao: %d \r\n", resultado, resto)
}
