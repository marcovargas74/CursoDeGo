package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/marcovargas74/cursoDeGo/SERVIDOR_WEB/model"
)

//Ola  é um manipulador de requisicao OLA
func Ola(w http.ResponseWriter, r *http.Request) {
	hora := time.Now().Format("15:04:05")
	pagina := model.Pagina{}
	pagina.Hora = hora
	if err := ModeloOla.ExecuteTemplate(w, "ola.html", pagina); err != nil {
		http.Error(w, "Houve um erro na renderizacao da Pagina.", http.StatusInternalServerError)
		fmt.Println("[ola] Erro na execucao do modelo", err.Error())
	}

	//fmt.Fprintln(w, "Aqui é um manipulador usando funcao !")
}
