package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3
	fmt.Println("O Numero ", numero, " se escreve assim:")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("tres.")
	}

	fmt.Println("Eu sou :", runtime.GOOS)
	fmt.Println("Voce é da Familia do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("SIMM!")
	default:
		fmt.Println("Nao, eu sou :", runtime.GOOS)

	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hoje é final de semana ", time.Now().Weekday())
	default:
		fmt.Println("Hoje é ", time.Now().Weekday(), " dia de Semana: ")

	}

	//numero = 10
	fmt.Println("Este numero cabe um digito")
	switch {
	case numero < 10:
		fmt.Println("SIM")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve 2 digitos?")
	case numero >= 100:
		fmt.Println("Nao da o numero é muito grande")

	}
}
