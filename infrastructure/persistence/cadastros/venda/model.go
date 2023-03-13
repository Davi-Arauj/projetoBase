package venda

import "time"

// Venda modela uma resposta para listagem e busca de vendas
type Venda struct {
	ID              *string     `sql:"id" codinome:"id"`
	UsuarioID       *string    `sql:"usuario_id" codinome:"usuario_id"`
	ClienteID       *string    `sql:"cliente_id" codinome:"cliente_id"`
	DataCriacao     *time.Time `sql:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao" codinome:"data_atualizacao"`
}

// VendaPag modela uma lista de respostas com suporte para paginação dos vendas na listagem
type VendaPag struct {
	Dados []Venda
	Prox  *bool
	Total *int64
}
