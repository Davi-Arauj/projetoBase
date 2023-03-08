package usuario

import "time"

// Usuario modela uma resposta para listagem e busca de usuarios
type Usuario struct {
	ID               *int64     `sql:"id" codinome:"id"`
	Nome             *string    `sql:"nome" codinome:"nome"`
	Senha            *string    `sql:"senha" codinome:"senha"`
	Cpf              *string    `sql:"cpf" codinome:"cpf"`
	Hash             *string    `sql:"hash" codinome:"hash"`
	Fone             *int64     `sql:"fone" codinome:"fone"`
	Email            *string    `sql:"email" codinome:"email"`
	Data_atualizacao *time.Time `sql:"data_atualizacao" codinome:"data_atualizacao"`
	Data_criacao     *time.Time `sql:"data_criacao" codinome:"data_criacao"`
}

// UsuarioPag modela uma lista de respostas com suporte para paginação dos usuarios na listagem
type UsuarioPag struct {
	Dados []Usuario
	Prox  *bool
	Total *int64
}
