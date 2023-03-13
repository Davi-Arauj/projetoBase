package produto_venda

import (
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto_venda"
	"github.com/projetoBase/util"
)

// IProdutoVenda define uma interface para os metodos de acesso à camada de dados
type IProdutoVenda interface {
	Listar(p *util.ParametrosRequisicao) (*produto_venda.ProdutoVendaPag, error)
	Adicionar(req *produto_venda.ProdutoVenda) error
	Alterar(req *produto_venda.ProdutoVenda) error
	Remover(codigoBarras int64) error
	ConverterParaProdutoVenda(data interface{}) (*produto_venda.ProdutoVenda, error)
}
