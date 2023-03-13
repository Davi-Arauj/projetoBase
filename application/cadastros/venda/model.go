package venda

import "time"

// Req modela uma requisição para a criação de uma venda
type Req struct {
	UsuarioID *string `json:"usuario_id" codinome:"usuario_id"`
	ClienteID *string `json:"cliente_id" codinome:"cliente_id"`
}

// Res modela uma resposta para listagem e busca de vendas
type Res struct {
	ID              *string    `json:"id,omitempty" codinome:"id"`
	UsuarioID       *string    `json:"usuario_id" codinome:"usuario_id"`
	ClienteID       *string    `json:"cliente_id" codinome:"cliente_id"`
	DataCriacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
}

// VendaUpdate modela uma requisição para a atualização de uma venda
type VendaUpdate struct {
	ID        *string `json:"id,omitempty" codinome:"id"`
	UsuarioID *string `json:"usuario_id" codinome:"usuario_id"`
	ClienteID *string `json:"cliente_id" codinome:"cliente_id"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos vendas na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
