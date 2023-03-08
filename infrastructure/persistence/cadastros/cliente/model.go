package cliente

import "time"

// Cliente modela uma resposta para listagem e busca de clientes
type Cliente struct {
	ID              *int64     `sql:"id" codinome:"id"`
	Nome            *string    `sql:"nome" codinome:"nome"`
	Email           *string    `sql:"email" codinome:"email"`
	Cpf             *string    `sql:"cpf" codinome:"cpf"`
	Fone            *int64     `sql:"fone" codinome:"fone"`
	DataCriacao     *time.Time `sql:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao" codinome:"data_atualizacao"`
}

// ClientePag modela uma lista de respostas com suporte para paginação dos clientes na listagem
type ClientePag struct {
	Dados []Cliente
	Prox  *bool
	Total *int64
}
