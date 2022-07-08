package produto

import "time"

// Req modela uma requisição para a criação ou atualização de um produto
type Req struct {
	CodigoBarras *int64   `json:"codigo_barras" codinome:"codigo_barras"`
	Nome         *string  `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao    *string  `json:"descricao" codinome:"descricao"`
	Foto         *string  `json:"endereco_foto" codinome:"endereco_foto"`
	Valorpago    *float64 `json:"valor_pago" codinome:"valor_pago"`
	Valorvenda   *float64 `json:"valor_venda" codinome:"valor_venda"`
	Qtde         *int64   `json:"quantidade" minLength:"1" codinome:"quantidade"`
	UndCod       *int64   `json:"unidade_id" codinome:"unidade_id"`
	CatCod       *int64   `json:"categoria_id" codinome:"categoria_id"`
	ScatCod      *int64   `json:"subcategoria_id" codinome:"subcategoria_id"`
}

// Res modela uma resposta para listagem e busca de produtos
type Res struct {
	ID              *int64     `json:"id,omitempty" codinome:"id"`
	CodigoBarras    *int64     `json:"codigo_barras" codinome:"codigo_barras"`
	Nome            *string    `json:"nome" binding:"required,gte=1" minLength:"1" codinome:"nome"`
	Descricao       *string    `json:"descricao" codinome:"descricao"`
	Foto            *string    `json:"endereco_foto" codinome:"endereco_foto"`
	Valorpago       *float64   `json:"valor_pago" codinome:"valor_pago"`
	Valorvenda      *float64   `json:"valor_venda" codinome:"valor_venda"`
	Qtde            *int64     `json:"quantidade" minLength:"1" codinome:"quantidade"`
	UndCod          *int64     `json:"unidade_id" codinome:"unidade_id"`
	CatCod          *int64     `json:"categoria_id" codinome:"categoria_id"`
	ScatCod         *int64     `json:"subcategoria_id" codinome:"subcategoria_id"`
	DataCriacao     *time.Time `json:"data_criacao" codinome:"data_criacao"`
	DataAtualizacao *time.Time `json:"data_atualizacao" codinome:"data_atualizacao"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos produtos na listagem
type ResPag struct {
	Dados []Res  `json:"dados,omitempty"`
	Prox  *bool  `json:"prox,omitempty"`
	Total *int64 `json:"total,omitempty"`
}
