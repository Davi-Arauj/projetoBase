package produto_venda

import (
	"context"

	"github.com/projetoBase/config/database"
	"github.com/projetoBase/domain/cadastros/produto_venda"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

func Listar(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um produto_venda"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := produto_venda.ObterRepo(tx)

	listaproduto_venda, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Dados = make([]Res, len(listaproduto_venda.Dados))
	for i := 0; i < len(listaproduto_venda.Dados); i++ {
		if err = util.ConvertStruct(&listaproduto_venda.Dados[i], &res.Dados[i]); err != nil {
			return res, oops.Wrap(err, msgErrPadrao)
		}
	}

	res.Total, res.Prox = listaproduto_venda.Total, listaproduto_venda.Prox

	return
}

// Adicionar contém a lógica de negócio para adicionar um novo produto_venda
func Adicionar(ctx context.Context, req *Req) (id *string, err error) {
	var (
		msgErrPadrao = "Erro ao cadastrar novo produto_venda"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := produto_venda.ObterRepo(tx)
	dados, err := repo.ConverterParaProdutoVenda(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = repo.Adicionar(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	id = dados.ID

	return
}

// Alterar contém a lógica de negócio para alterar um novo produto_venda
func Alterar(ctx context.Context, req *Req) (err error) {
	var (
		msgErrPadrao = "Erro ao alterar produto_venda"
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := produto_venda.ObterRepo(tx)

	dados, err := repo.ConverterParaProdutoVenda(req)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = repo.Alterar(dados); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Remover contém a lógica de negócio para remover um novo produto_venda
func Remover(ctx context.Context, codigoBarras int64) (err error) {
	msgErrPadrao := "Erro ao remover produto_venda"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := produto_venda.ObterRepo(tx)

	if err = repo.Remover(codigoBarras); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}
