package produto

import (
	"context"
	"strconv"

	"github.com/projetoBase/config/database"
	"github.com/projetoBase/domain/cadastros/produto"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)
// Listar contém a lógica de negócio para buscar ou listar produtos
func Listar(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um produto"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := produto.ObterRepo(tx)

	listaproduto, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Dados = make([]Res, len(listaproduto.Dados))
	for i := 0; i < len(listaproduto.Dados); i++ {
		if err = util.ConvertStruct(&listaproduto.Dados[i], &res.Dados[i]); err != nil {
			return res, oops.Wrap(err, msgErrPadrao)
		}
	}

	res.Total, res.Prox = listaproduto.Total, listaproduto.Prox

	return
}

// Adicionar contém a lógica de negócio para adicionar um novo produto
func Adicionar(ctx context.Context, req *Req) (res *Res, err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao cadastrar novo produto"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := produto.ObterRepo(tx)
	dados, err := repo.ConverterParaProduto(req)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}
	// devemos verificar se já existe um registro com mesmo codigo de barras
	p.Filtros = make(map[string][]string)
	p.Filtros["codigo_barras"] = []string{strconv.FormatInt(*req.CodigoBarras, 10)}
	p.Total = true

	lista, err := repo.Listar(&p)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	if lista.Total != nil && *lista.Total > 0 {
		return nil, oops.NovoErr("Já existe um produto com esse codigo de barras!")
	}

	if err = repo.Adicionar(dados); err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Alterar contém a lógica de negócio para alterar um novo produto
func Alterar(ctx context.Context, req *Req) (err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao alterar produto"
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := produto.ObterRepo(tx)

	dados, err := repo.ConverterParaProduto(req)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	// devemos verificar se já existe um registro com o mesmo codigo de barras
	p.Filtros = make(map[string][]string)
	p.Filtros["codigo_barras"] = []string{strconv.FormatInt(*req.CodigoBarras, 10)}

	lista, err := repo.Listar(&p)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if len(lista.Dados) > 0 {
			return oops.NovoErr("Já existe um produto com esse codigo de barras!")
	}

	if err = repo.Alterar(dados); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Remover contém a lógica de negócio para remover um novo produto
func Remover(ctx context.Context, codigoBarras int64) (err error) {
	msgErrPadrao := "Erro ao remover produto"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := produto.ObterRepo(tx)

	if err = repo.Remover(codigoBarras); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

