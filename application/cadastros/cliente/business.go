package cliente

import (
	"context"

	"github.com/projetoBase/config/database"
	"github.com/projetoBase/domain/cadastros/cliente"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

func Listar(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um cliente"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := cliente.ObterRepo(tx)

	listacliente, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Dados = make([]Res, len(listacliente.Dados))
	for i := 0; i < len(listacliente.Dados); i++ {
		if err = util.ConvertStruct(&listacliente.Dados[i], &res.Dados[i]); err != nil {
			return res, oops.Wrap(err, msgErrPadrao)
		}
	}

	res.Total, res.Prox = listacliente.Total, listacliente.Prox

	return
}

// Adicionar contém a lógica de negócio para adicionar um novo cliente
func Adicionar(ctx context.Context, req *Req) (id *string, err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao cadastrar novo cliente"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := cliente.ObterRepo(tx)
	dados, err := repo.ConverterParaCliente(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	// devemos verificar se já existe um registro com o mesmo email
	p.Filtros = make(map[string][]string)
	p.Filtros["email"] = []string{*req.Email}
	p.Total = true

	lista, err := repo.Listar(&p)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if lista.Total != nil && *lista.Total > 0 {
		return id, oops.NovoErr("Já existe um cliente com esse email!")
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

// Alterar contém a lógica de negócio para alterar um novo cliente
func Alterar(ctx context.Context, clienteID string, req *Req) (err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao alterar cliente"
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := cliente.ObterRepo(tx)

	dados, err := repo.ConverterParaCliente(req)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	dados.ID = &clienteID
	// devemos verificar se já existe um registro com o mesmo email
	p.Filtros = make(map[string][]string)
	p.Filtros["email"] = []string{*req.Email}

	lista, err := repo.Listar(&p)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if len(lista.Dados) > 0 {
		if lista.Dados[0].ID != &clienteID {
			return oops.NovoErr("Já existe um cliente com esse email!")
		}
	}

	if err = repo.Alterar(dados); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Remover contém a lógica de negócio para remover um novo cliente
func Remover(ctx context.Context, clienteID string) (err error) {
	msgErrPadrao := "Erro ao remover cliente"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := cliente.ObterRepo(tx)

	if err = repo.Remover(clienteID); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}
