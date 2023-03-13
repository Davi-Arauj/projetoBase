package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/venda"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// PGVenda é uma estrutura base para acesso aos metodos do banco postgres para manipulação de vendas
type PGVenda struct {
	DB *database.DBTransacao
}

// Listar lista de forma paginada os dados de vendas do banco postgres
func (pg *PGVenda) Listar(p *util.ParametrosRequisicao) (res *venda.VendaPag, err error) {
	var t venda.Venda

	res = new(venda.VendaPag)

	campos, _, err := p.ValidarCampos(&t)
	if err != nil {
		return res, oops.Err(err)
	}

	consultaPrevia := pg.DB.Builder.
		Select(campos...).
		From("t_venda")

	clausulaWhere := p.CriarFiltros(consultaPrevia, map[string]util.Filtro{

		"nao_contem_id": util.CriarFiltros("id", util.FlagFiltroNotIn),
	})

	dados, prox, total, err := util.ConfigurarPaginacao(p, &t, &clausulaWhere)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Dados, res.Prox, res.Total = dados.([]venda.Venda), prox, total

	return
}

// Adicionar adiciona um novo venda ao banco de dados do postgres
func (pg *PGVenda) Adicionar(req *venda.Venda) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_venda").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Alterar altera um venda no banco de dados do postgres
func (pg *PGVenda) Alterar(req *venda.Venda) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_venda").
		SetMap(valores).
		Where(squirrel.Eq{
			"id": req.ID,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}
	return
}

// Remover remove um venda no banco de dados do postgres
func (pg *PGVenda) Remover(vendaID string) (err error) {
	resultado, err := pg.DB.Builder.
		Delete("t_venda").
		Where(squirrel.Eq{
			"id": vendaID,
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
