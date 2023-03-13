package usuario

import (
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/usuario"
	"github.com/projetoBase/infrastructure/persistence/cadastros/usuario/postgres"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

type repositorio struct {
	pg *postgres.PGUsuario
}

func novoRepo(db *database.DBTransacao) *repositorio {
	return &repositorio{
		pg: &postgres.PGUsuario{DB: db},
	}
}

// Listar é um gerenciador de fluxo de dados para listar um usuario no banco de dados
func (r *repositorio) Listar(p *util.ParametrosRequisicao) (*usuario.UsuarioPag, error) {
	return r.pg.Listar(p)
}

// Adicionar é um gerenciador de fluxo de dados para adicionar um usuario no banco de dados
func (r *repositorio) Adicionar(req *usuario.Usuario) error {
	return r.pg.Adicionar(req)
}

// Alterar é um gerenciador de fluxo de dados para alterar um usuario no banco de dados
func (r *repositorio) Alterar(req *usuario.Usuario) error {
	return r.pg.Alterar(req)
}

// Remover é um gerenciador de fluxo de dados para remover um usuario no banco de dados
func (r *repositorio) Remover(email string) error {
	return r.pg.Remover(email)
}

// ConverterParaUsuario converte uma requisição em uma estrutura de usuario para acesso à camada de dados
func (r *repositorio) ConverterParaUsuario(dados interface{}) (*usuario.Usuario, error) {
	res := &usuario.Usuario{}

	if err := util.ConvertStruct(dados, res); err != nil {
		return res, oops.Err(err)
	}

	return res, nil
}
