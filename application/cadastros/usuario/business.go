package usuario

import (
	"context"

	"github.com/projetoBase/config/database"
	"github.com/projetoBase/domain/cadastros/usuario"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

func Listar(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um usuario"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := usuario.ObterRepo(tx)

	listausuario, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Dados = make([]Res, len(listausuario.Dados))
	for i := 0; i < len(listausuario.Dados); i++ {
		if err = util.ConvertStruct(&listausuario.Dados[i], &res.Dados[i]); err != nil {
			return res, oops.Wrap(err, msgErrPadrao)
		}
	}

	res.Total, res.Prox = listausuario.Total, listausuario.Prox

	return
}

// Adicionar contém a lógica de negócio para adicionar um novo usuario
func Adicionar(ctx context.Context, req *Req) (res *Res, err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao cadastrar novo usuario"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := usuario.ObterRepo(tx)
	dados, err := repo.ConverterParaUsuario(req)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}
	// devemos verificar se já existe um registro com o mesmo email
	p.Filtros = make(map[string][]string)
	p.Filtros["email"] = []string{*req.Email}
	p.Total = true

	lista, err := repo.Listar(&p)
	if err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	if lista.Total != nil && *lista.Total > 0 {
		return nil, oops.NovoErr("Já existe um usuario com esse email!")
	}

	if err = repo.Adicionar(dados); err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return nil, oops.Wrap(err, msgErrPadrao)
	}

	res.ID = dados.ID
	return
}

// Alterar contém a lógica de negócio para alterar um novo usuario
func Alterar(ctx context.Context, usuario_id string, req *Req) (err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao alterar usuario"
	)
	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := usuario.ObterRepo(tx)

	dados, err := repo.ConverterParaUsuario(req)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	// no caso de alteração de nome
	//
	// verificar o id na url com o email passado

	// devemos verificar se já existe um registro com o mesmo email
	p.Filtros = make(map[string][]string)
	p.Filtros["email"] = []string{*req.Email}

	lista, err := repo.Listar(&p)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if len(lista.Dados) > 0 {
		// comparar o email passado no body com o email original do usuario que vai ser retornado pelo id
		return oops.NovoErr("Já existe um usuario com esse email!")
	}

	if err = repo.Alterar(dados); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Remover contém a lógica de negócio para remover um novo usuario
func Remover(ctx context.Context, email string) (err error) {
	msgErrPadrao := "Erro ao remover usuario"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := usuario.ObterRepo(tx)

	if err = repo.Remover(email); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}
