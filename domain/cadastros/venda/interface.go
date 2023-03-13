package venda

import (
	"github.com/projetoBase/infrastructure/persistence/cadastros/venda"
	"github.com/projetoBase/util"
)

// IVenda define uma interface para os metodos de acesso à camada de dados
type IVenda interface {
	Listar(p *util.ParametrosRequisicao) (*venda.VendaPag, error)
	Adicionar(req *venda.Venda) error
	Alterar(req *venda.Venda) error
	Remover(vendaID string) error
	ConverterParaVenda(data interface{}) (*venda.Venda, error)
}
