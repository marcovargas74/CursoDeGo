package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina. Erro:", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler conteudo da pagina. Erro:", err.Error())
			return
		}
		fmt.Println(string(corpo))

	}

	fmt.Println(" ")
	fmt.Println("------------------------------------------------------------- ")

	request, err := http.NewRequest("GET", "http://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para a pagina. Erro:", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina. Erro:", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler conteudo da pagina. Erro:", err.Error())
			return
		}
		fmt.Println(string(corpo))

	}

}
