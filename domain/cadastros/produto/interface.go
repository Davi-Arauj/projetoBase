package produto

import (
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto"
	"github.com/projetoBase/util"
)

type IProduto interface {
	Listar(p *util.ParametrosRequisicao) (*produto.ProdutoPag, error)
	Buscar(req *produto.Produto) error
	Adicionar(req *produto.Produto) error
	Alterar(req *produto.Produto) error
	Remover(codigoBarras int64) error
	ConverterParaProduto(data interface{}) (*produto.Produto, error)
}
