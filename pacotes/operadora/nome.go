package operadora

import (
	"strconv"

	"github.com/marcovargas74/cursoDeGo/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora O prefixo da Operadora na Capital
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
