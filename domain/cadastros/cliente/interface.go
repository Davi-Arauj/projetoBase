package cliente

import (
	"github.com/projetoBase/infrastructure/persistence/cadastros/cliente"
	"github.com/projetoBase/util"
)

// ICliente define uma interface para os metodos de acesso à camada de dados
type ICliente interface {
	Listar(p *util.ParametrosRequisicao) (*cliente.ClientePag, error)
	Adicionar(req *cliente.Cliente) error
	Alterar(req *cliente.Cliente) error
	Remover(email string) error
	ConverterParaCliente(data interface{}) (*cliente.Cliente, error)
}
