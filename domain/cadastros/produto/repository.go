package produto

import (
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto/postgres"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

type repositorio struct {
	pg *postgres.PGProduto
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGProduto{DB: db},
	}
}

// Listar é um gerenciador de fluxo de dados para listar um produto no banco de dados
func (r *repositorio) Listar(p *util.ParametrosRequisicao) (*produto.ProdutoPag, error) {
	return r.pg.Listar(p)
}

// Adicionar é um gerenciador de fluxo de dados para adicionar um produto no banco de dados
func (r *repositorio) Adicionar(req *produto.Produto) error {
	return r.pg.Adicionar(req)
}

// Alterar é um gerenciador de fluxo de dados para alterar um produto no banco de dados
func (r *repositorio) Alterar(req *produto.Produto) error {
	return r.pg.Alterar(req)
}

// Remover é um gerenciador de fluxo de dados para remover um produto no banco de dados
func (r *repositorio) Remover(codigoBarras int64) error {
	return r.pg.Remover(codigoBarras)
}

// ConverterParaProduto converte uma requisição em uma estrutura de produto para acesso à camada de dados
func (r *repositorio) ConverterParaProduto(dados interface{}) (*produto.Produto, error) {
	res := &produto.Produto{}

	if err := util.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}

	return res, nil
}
