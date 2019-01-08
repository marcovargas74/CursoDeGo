package manipulador

import "html/template"

//ModeloOla armazenam o modelo de paginas Ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazenam o modelos de pagina Local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
