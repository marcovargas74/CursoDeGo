package main

import (
	"fmt"
	"net/http"

	"github.com/marcovargas74/cursoDeGo/SERVIDOR_WEB/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ola bem vindo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Estou escutando...")
	http.ListenAndServe(":8181", nil)
}
