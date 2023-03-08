package usuario

import (
	"context"
	"strconv"

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

// Buscar contém a lógica de negócio para buscar um usuario
func Buscar(ctx context.Context, codigoBarras int64) (res *Res, err error) {
	msgErrPadrao := "Erro ao buscar um usuario"
	res = new(Res)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := usuario.ObterRepo(tx)

	req, err := repo.ConverterParausuario(res)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	req.CodigoBarras = &codigoBarras

	if err = repo.Buscar(req); err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	if err = util.ConvertStruct(req, res); err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Adicionar contém a lógica de negócio para adicionar um novo usuario
func Adicionar(ctx context.Context, req *Req) (id *int64, err error) {
	var (
		p            util.ParametrosRequisicao
		msgErrPadrao = "Erro ao cadastrar novo usuario"
	)

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := usuario.ObterRepo(tx)
	dados, err := repo.ConverterParausuario(req)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}
	// devemos verificar se já existe um registro com mesmo codigo de barras
	p.Filtros = make(map[string][]string)
	p.Filtros["codigo_barras"] = []string{strconv.FormatInt(*req.CodigoBarras, 10)}
	p.Total = true

	lista, err := repo.Listar(&p)
	if err != nil {
		return id, oops.Wrap(err, msgErrPadrao)
	}

	if lista.Total != nil && *lista.Total > 0 {
		return id, oops.NovoErr("Já existe um usuario com esse codigo de barras!")
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

// Alterar contém a lógica de negócio para alterar um novo usuario
func Alterar(ctx context.Context, codigoBarras int64, req *Req) (err error) {
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

	dados, err := repo.ConverterParausuario(req)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	dados.CodigoBarras = &codigoBarras
	// devemos verificar se já existe um registro com o mesmo codigo de barras
	p.Filtros = make(map[string][]string)
	p.Filtros["codigo_barras"] = []string{strconv.FormatInt(*req.CodigoBarras, 10)}

	lista, err := repo.Listar(&p)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if len(lista.Dados) > 0 {
		if lista.Dados[0].CodigoBarras != &codigoBarras{
			return oops.NovoErr("Já existe um usuario com esse codigo de barras!")
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

// Remover contém a lógica de negócio para remover um novo usuario
func Remover(ctx context.Context, codigoBarras int64) (err error) {
	msgErrPadrao := "Erro ao remover usuario"

	tx, err := database.NovaTransacao(ctx, false)
	if err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()

	repo := usuario.ObterRepo(tx)

	if err = repo.Remover(codigoBarras); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	if err = tx.Commit(); err != nil {
		return oops.Wrap(err, msgErrPadrao)
	}

	return
}

// Total contém a lógica de negócio para buscar o total de uma listagem
func Total(ctx context.Context, p *util.ParametrosRequisicao) (res *ResPag, err error) {
	msgErrPadrao := "Erro ao listar um usuario"

	res = new(ResPag)

	tx, err := database.NovaTransacao(ctx, true)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}
	defer tx.Rollback()
	repo := usuario.ObterRepo(tx)

	p.Filtros = make(map[string][]string)
	p.Total = true

	listausuario, err := repo.Listar(p)
	if err != nil {
		return res, oops.Wrap(err, msgErrPadrao)
	}

	res.Total, res.Prox = listausuario.Total, listausuario.Prox

	return
}
