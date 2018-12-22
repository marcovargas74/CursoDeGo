package model

//Galinha é Galinha
type Galinha interface {
	Cacareja() string
}

//Pata é mulher do Pato
type Pata interface {
	Quack() string
}

//Ave representa um animal
type Ave struct {
	Nome string
}

//Cacareja blavlabla
func (a Ave) Cacareja() string {
	return "Cocoricó..."
}

//Quack e o som do PAto
func (a Ave) Quack() string {
	return "Quack, quack..."
}
