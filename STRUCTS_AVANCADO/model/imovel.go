package model

//Imovel é uma struct do tipo imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor atualiza o campo valor na struct
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do campo valor na struct
func (i *Imovel) GetValor() int {
	return i.valor
}
