package main

import (
	"fmt"
	"net/http"

	"github.com/marcovargas74/cursoDeGo/BANCO_SQL/manipulador"
	"github.com/marcovargas74/cursoDeGo/BANCO_SQL/repo"
)

func init() {
	fmt.Println("Iniciando o servidor...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Erro ao Abrir Banco de Dados", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola bem vindo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando...")
	http.ListenAndServe(":8181", nil)
}
