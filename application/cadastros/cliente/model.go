package cliente

import "time"

// Req modela uma requisição para a criação ou atualização de um produto
type Req struct {
	Nome  *string `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Email *string `json:"email" binding:"email" codinome:"email"`
	Cpf   *string `json:"cpf" binding:"required,gte=9" minLength:"9" codinome:"cpf"`
	Fone  *int64  `json:"fone" minLength:"8" codinome:"fone"`
}

// Res modela uma resposta para listagem e busca de produtos
type Res struct {
	ID              *int64     `json:"id,omitempty" codinome:"id"`
	Nome            *string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Email           *string    `json:"email" binding:"email" codinome:"email"`
	Cpf             *string    `json:"cpf" binding:"required,gte=9" minLength:"9" codinome:"cpf"`
	Fone            *int64     `json:"fone" minLength:"8" codinome:"fone"`
	DataCriacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos produtos na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
