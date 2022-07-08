package produto

import "time"

// Res modela uma resposta para listagem e busca de produtos
type Produto struct {
	ID              *int64     `sql:"id" codinome:"id"`
	DataCriacao     *time.Time `sql:"data_criacao::TIMESTAMPTZ" codinome:"data_criacao"`
	DataAtualizacao *time.Time `sql:"data_atualizacao::TIMESTAMPTZ" codinome:"data_atualizacao"`
	CodigoBarras    *int64     `sql:"codigo_barras" codinome:"codigo_barras"`
	Nome            *string    `sql:"nome" codinome:"nome"`
	Descricao       *string    `sql:"descricao" codinome:"descricao"`
	Foto            *string    `sql:"endereco_foto" codinome:"endereco_foto"`
	Valorpago       *float64   `sql:"valor_pago" codinome:"valor_pago"`
	Valorvenda      *float64   `sql:"valor_venda" codinome:"valor_venda"`
	Qtde            *int64     `sql:"quantidade" codinome:"quantidade"`
	UndCod          *int64     `sql:"unidade_id" codinome:"unidade_id"`
	CatCod          *int64     `sql:"categoria_id" codinome:"categoria_id"`
	ScatCod         *int64     `sql:"subcategoria_id" codinome:"subcategoria_id"`
}

// ResPag modela uma lista de respostas com suporte para paginação dos produtos na listagem
type ProdutoPag struct {
	Dados []Produto
	Prox  *bool
	Total *int64
}
