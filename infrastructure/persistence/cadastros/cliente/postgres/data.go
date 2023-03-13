package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/cliente"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// PGCliente é uma estrutura base para acesso aos metodos do banco postgres para manipulação de clientes
type PGCliente struct {
	DB *database.DBTransacao
}

// Listar lista de forma paginada os dados de clientes do banco postgres
func (pg *PGCliente) Listar(p *util.ParametrosRequisicao) (res *cliente.ClientePag, err error) {
	var t cliente.Cliente

	res = new(cliente.ClientePag)

	campos, _, err := p.ValidarCampos(&t)
	if err != nil {
		return res, oops.Err(err)
	}

	consultaPrevia := pg.DB.Builder.
		Select(campos...).
		From("t_cliente")

	clausulaWhere := p.CriarFiltros(consultaPrevia, map[string]util.Filtro{
		"nao_contem_id": util.CriarFiltros("id", util.FlagFiltroNotIn),
		"nome":          util.CriarFiltros("lower(public.unaccent(nome)) LIKE lower('%'||public.unaccent(?)||'%')", util.FlagFiltroEq),
		"nome_exato":    util.CriarFiltros("public.unaccent(nome) LIKE public.unaccent(?)", util.FlagFiltroEq),
	})

	dados, prox, total, err := util.ConfigurarPaginacao(p, &t, &clausulaWhere)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Dados, res.Prox, res.Total = dados.([]cliente.Cliente), prox, total

	return
}

// Adicionar adiciona um novo cliente ao banco de dados do postgres
func (pg *PGCliente) Adicionar(req *cliente.Cliente) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_cliente").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Alterar altera um cliente no banco de dados do postgres
func (pg *PGCliente) Alterar(req *cliente.Cliente) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_cliente").
		SetMap(valores).
		Where(squirrel.Eq{
			"email": req.Email,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}
	return
}

// Remover remove um cliente no banco de dados do postgres
func (pg *PGCliente) Remover(email string) (err error) {
	resultado, err := pg.DB.Builder.
		Delete("t_cliente").
		Where(squirrel.Eq{
			"email": email,
		}).Exec()

	if err != nil {
		return oops.Err(err)
	}

	linhas, err := resultado.RowsAffected()
	if err != nil {
		return oops.Err(err)
	} else if linhas == 0 {
		return oops.Err(sql.ErrNoRows)
	}

	return
}
