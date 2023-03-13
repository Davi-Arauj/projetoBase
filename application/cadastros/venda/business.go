package venda

import (
	"context"

	"github.com/projetoBase/config/database"
	"github.com/projetoBase/domain/cadastros/venda"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)
// Listar contém a lógica de negócio para buscar ou listar vendas
func Listar(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um venda"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := venda.ObterRepo(tx)

	listavenda, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Dados = make([]Res, len(listavenda.Dados))
	for i := 0; i < len(listavenda.Dados); i++ {
		if err = util.ConvertStruct(&listavenda.Dados[i], &res.Dados[i]); err != nil {
			return res, oops.Wrap(err, msgErrPadrao)
		}
	}

	res.Total, res.Prox = listavenda.Total, listavenda.Prox

	return
}

// Adicionar contém a lógica de negócio para adicionar uma nova venda
func Adicionar(ctx context.Context, req *Req) (id *int64, err error) {
	var (
		msgErrPadrao = "Erro ao cadastrar nova venda"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := venda.ObterRepo(tx)
	dados, err := repo.ConverterParaVenda(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = repo.Adicionar(dados); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Alterar contém a lógica de negócio para alterar uma nova venda
func Alterar(ctx context.Context, req *VendaUpdate) (err error) {
	var (
		msgErrPadrao = "Erro ao alterar venda"
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := venda.ObterRepo(tx)

	dados, err := repo.ConverterParaVenda(req)
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

// Remover contém a lógica de negócio para remover uma nova venda
func Remover(ctx context.Context, vendaID string) (err error) {
	msgErrPadrao := "Erro ao remover venda"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := venda.ObterRepo(tx)

	if err = repo.Remover(vendaID); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}
