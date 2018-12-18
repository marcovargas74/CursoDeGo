package model

import "errors"

//Imovel é uma struct do tipo imovel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

/*
ErrValorDeImovelInvalido é um erro paresentado quando valor invalido
*/
var ErrValorDeImovelInvalido = errors.New("O valor informando nao é valido")

/*
ErrValorDeImovelMuitoAlto é um erro apresentado quando valor muito ALTO
*/
var ErrValorDeImovelMuitoAlto = errors.New("O valor informando nao é MUITO ALTO")

//SetValor atualiza o campo valor na struct
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do campo valor na struct
func (i *Imovel) GetValor() int {
	return i.valor
}
