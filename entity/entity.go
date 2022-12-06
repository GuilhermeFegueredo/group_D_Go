package entity

import (
	"encoding/json"
	"log"
)

type ProdutoInterface interface {
	toString() string
}

type Produto struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Codigo      string  `json:"codigo"`
	Preco       float64 `json:"preco"`
	CriadoEm    string  `json:"criado_em,omitempty"`
	AtulizadoEm string  `json:"atulizado_em,omitempty"`
}

func (p Produto) toString() string {
	data, err := json.Marshal(p)
	if err != nil {
		log.Printf("Erro ao converter Produto para String: \n--> %v", err)
	}

	return string(data)
}

func NovoProduto(nome, codigo string, preco float64) *Produto {
	return &Produto{
		Nome:   nome,
		Codigo: codigo,
		Preco:  preco,
	}
}

type ListaDeProduto struct {
	List []*Produto `json:"lista"`
}

func (prodLista ListaDeProduto) toString() string {
	data, err := json.Marshal(prodLista)
	if err != nil {
		log.Printf("Erro ao converter Produto para String: \n--> %v", err)
	}

	return string(data)
}

type User struct {
	Username string `json:"username"`
	Senha    string `json:"senha"`
}

func NovoAdmin() *User {
	return &User{
		Username: "admin",
		Senha:    "megasenha",
	}
}
