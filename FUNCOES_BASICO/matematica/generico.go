package matematica

/*
Calculo executa qualquer tipo de calculo
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x vezes o Y
func Multiplicador(x, y int) int {
	return x * y
}

//Divisor divide x por Y
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto divide x por Y e retorna o Resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
