package usuario

import (
	"github.com/projetoBase/infrastructure/persistence/cadastros/usuario"
	"github.com/projetoBase/util"
)

// IUsuario define uma interface para os metodos de acesso à camada de dados
type IUsuario interface {
	Listar(p *util.ParametrosRequisicao) (*usuario.UsuarioPag, error)
	Adicionar(req *usuario.Usuario) error
	Alterar(req *usuario.Usuario) error
	Remover(email string) error
	ConverterParaUsuario(data interface{}) (*usuario.Usuario, error)
}
