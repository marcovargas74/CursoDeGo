package manipulador

import (
	"fmt"
	"net/http"
)

//Funcao é um manipulador de requisicao HTTP
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aqui é um manipulador usando funcao !")
}