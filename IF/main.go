package main

import (
	"fmt"
)

func main() {

	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve a pouco tempo.")
	}

	if situacao {
		fmt.Println("Esse credor está Devendo.")
	}

	if !situacao {
		fmt.Println("Esse credor está adimplente.")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive na estado Teste.")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situacao do cliente", descricao)
	}
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente esta devendo"
		return
	}
	descricao = "O Cliente esta em dia."
	return
}
