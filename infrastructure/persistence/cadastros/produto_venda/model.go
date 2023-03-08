package produto_venda

import "time"

// ProdutoVenda modela uma resposta para listagem e busca de produto_vendas
type ProdutoVenda struct {
	ID              *string    `sql:"id" codinome:"id"`
	VendaID         *string    `sql:"venda_id" codinome:"venda_id"`
	ProdutoID       *string    `sql:"produto_id" codinome:"produto_id"`
	ValorUnt        *float64   `sql:"valor_unt" codinome:"valor_unt"`
	ValorTotal      *float64   `sql:"valor_total" codinome:"valor_total"`
	Quantidade      *int64     `sql:"quantidade" codinome:"quantidade"`
	DataCriacao     *time.Time `sql:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao" codinome:"data_atualizacao"`
}

// ProdutoVendaPag modela uma lista de respostas com suporte para paginação dos produto_vendas na listagem
type ProdutoVendaPag struct {
	Dados []ProdutoVenda
	Prox  *bool
	Total *int64
}
