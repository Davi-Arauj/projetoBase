package produto_venda

import "time"

// Req modela uma requisição para a criação ou atualização de um produto_venda
type Req struct {
	VendaID    *string  `json:"venda_id" codinome:"venda_id"`
	ProdutoID  *string  `json:"produto_id" codinome:"produto_id"`
	ValorUnt   *float64 `json:"valor_unt" codinome:"valor_unt"`
	ValorTotal *float64 `json:"valor_total" codinome:"valor_total"`
	Quantidade *int64   `json:"quantidade" minLength:"1" codinome:"quantidade"`
}

// Res modela uma resposta para listagem e busca de produto_venda
type Res struct {
	ID              *string    `json:"id,omitempty" codinome:"id"`
	VendaID         *string    `json:"venda_id" codinome:"venda_id"`
	ProdutoID       *string    `json:"produto_id" codinome:"produto_id"`
	ValorUnt        *float64   `json:"valor_unt" codinome:"valor_unt"`
	ValorTotal      *float64   `json:"valor_total" codinome:"valor_total"`
	Quantidade      *int64     `json:"quantidade" minLength:"1" codinome:"quantidade"`
	DataCriacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos produtos na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
