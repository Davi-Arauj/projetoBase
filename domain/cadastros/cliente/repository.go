package cliente

import (
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/cliente"
	"github.com/projetoBase/infrastructure/persistence/cadastros/cliente/postgres"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

type repositorio struct {
	pg *postgres.PGCliente
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGCliente{DB: db},
	}
}

// Listar é um gerenciador de fluxo de dados para listar um cliente no banco de dados
func (r *repositorio) Listar(p *util.ParametrosRequisicao) (*cliente.ClientePag, error) {
	return r.pg.Listar(p)
}

// Buscar é um gerenciador de fluxo de dados para buscar um cliente no banco de dados
func (r *repositorio) Buscar(req *cliente.Cliente) error {
	return r.pg.Buscar(req)
}

// Adicionar é um gerenciador de fluxo de dados para adicionar um cliente no banco de dados
func (r *repositorio) Adicionar(req *cliente.Cliente) error {
	return r.pg.Adicionar(req)
}

// Alterar é um gerenciador de fluxo de dados para alterar um cliente no banco de dados
func (r *repositorio) Alterar(req *cliente.Cliente) error {
	return r.pg.Alterar(req)
}

// Remover é um gerenciador de fluxo de dados para remover um cliente no banco de dados
func (r *repositorio) Remover(codigoBarras int64) error {
	return r.pg.Remover(codigoBarras)
}

// ConverterParaCliente converte uma requisição em uma estrutura de cliente para acesso à camada de dados
func (r *repositorio) ConverterParaCliente(dados interface{}) (*cliente.Cliente, error) {
	res := &cliente.Cliente{}

	if err := util.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}

	return res, nil
}
