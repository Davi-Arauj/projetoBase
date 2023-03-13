package produto_venda

import (
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto_venda"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto_venda/postgres"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

type repositorio struct {
	pg *postgres.PGProdutoVenda
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGProdutoVenda{DB: db},
	}
}

// Listar é um gerenciador de fluxo de dados para listar um produto_venda no banco de dados
func (r *repositorio) Listar(p *util.ParametrosRequisicao) (*produto_venda.ProdutoVendaPag, error) {
	return r.pg.Listar(p)
}

// Adicionar é um gerenciador de fluxo de dados para adicionar um produto_venda no banco de dados
func (r *repositorio) Adicionar(req *produto_venda.ProdutoVenda) error {
	return r.pg.Adicionar(req)
}

// Alterar é um gerenciador de fluxo de dados para alterar um produto_venda no banco de dados
func (r *repositorio) Alterar(req *produto_venda.ProdutoVenda) error {
	return r.pg.Alterar(req)
}

// Remover é um gerenciador de fluxo de dados para remover um produto_venda no banco de dados
func (r *repositorio) Remover(codigoBarras int64) error {
	return r.pg.Remover(codigoBarras)
}

// ConverterParaProdutoVenda converte uma requisição em uma estrutura de produto_venda para acesso à camada de dados
func (r *repositorio) ConverterParaProdutoVenda(dados interface{}) (*produto_venda.ProdutoVenda, error) {
	res := &produto_venda.ProdutoVenda{}

	if err := util.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}

	return res, nil
}
