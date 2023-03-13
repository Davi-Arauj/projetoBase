package venda

import (
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/venda"
	"github.com/projetoBase/infrastructure/persistence/cadastros/venda/postgres"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

type repositorio struct {
	pg *postgres.PGVenda
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGVenda{DB: db},
	}
}

// Listar é um gerenciador de fluxo de dados para listar um venda no banco de dados
func (r *repositorio) Listar(p *util.ParametrosRequisicao) (*venda.VendaPag, error) {
	return r.pg.Listar(p)
}

// Adicionar é um gerenciador de fluxo de dados para adicionar um venda no banco de dados
func (r *repositorio) Adicionar(req *venda.Venda) error {
	return r.pg.Adicionar(req)
}

// Alterar é um gerenciador de fluxo de dados para alterar um venda no banco de dados
func (r *repositorio) Alterar(req *venda.Venda) error {
	return r.pg.Alterar(req)
}

// Remover é um gerenciador de fluxo de dados para remover um venda no banco de dados
func (r *repositorio) Remover(vendaID string) error {
	return r.pg.Remover(vendaID)
}

// ConverterParaVenda converte uma requisição em uma estrutura de venda para acesso à camada de dados
func (r *repositorio) ConverterParaVenda(dados interface{}) (*venda.Venda, error) {
	res := &venda.Venda{}

	if err := util.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}

	return res, nil
}
