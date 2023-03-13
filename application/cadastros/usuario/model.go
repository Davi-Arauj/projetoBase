package usuario

import "time"

// Req modela uma requisição para a criação ou atualização de um usuario
type Req struct {
	Nome             *string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Senha            *string    `json:"senha" codinome:"senha"`
	Cpf              *string    `json:"cpf" codinome:"cpf"`
	Fone             *int64     `json:"fone" minLength:"8" codinome:"fone"`
	Email            *string    `json:"email" codinome:"email"`
	Data_atualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
	Data_criacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
}

// Res modela uma resposta para listagem e busca de usuarios
type Res struct {
	ID               *int64     `json:"id,omitempty" codinome:"id"`
	Nome             *string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Senha            *string    `json:"senha" codinome:"senha"`
	Cpf              *string    `json:"cpf,omitempty" codinome:"cpf"`
	Fone             *int64     `json:"fone" minLength:"8" codinome:"fone"`
	Email            *string    `json:"email" codinome:"email"`
	Data_atualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
	Data_criacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos usuarios na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
