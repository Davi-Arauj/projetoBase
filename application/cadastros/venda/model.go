package venda

import "time"

// Req modela uma requisição para a criação ou atualização de um produto
type Req struct {
	UsuarioID *string `json:"usuario_id" codinome:"usuario_id"`
	ClienteID *string `json:"cliente_id" codinome:"cliente_id"`
}

// Res modela uma resposta para listagem e busca de produtos
type Res struct {
	ID              *int64     `json:"id,omitempty" codinome:"id"`
	UsuarioID       *string    `json:"usuario_id" codinome:"usuario_id"`
	ClienteID       *string    `json:"cliente_id" codinome:"cliente_id"`
	DataCriacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos produtos na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
