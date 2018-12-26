package manipulador

import "html/template"

//Modelos armazenam os modelos de paginas que serao executados pleos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
