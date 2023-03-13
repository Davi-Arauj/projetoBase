package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto_venda"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// PGProdutoVenda é uma estrutura base para acesso aos metodos do banco postgres para manipulação de produto_vendas
type PGProdutoVenda struct {
	DB *database.DBTransacao
}

// Listar lista de forma paginada os dados de produto_vendas do banco postgres
func (pg *PGProdutoVenda) Listar(p *util.ParametrosRequisicao) (res *produto_venda.ProdutoVendaPag, err error) {
	var t produto_venda.ProdutoVenda

	res = new(produto_venda.ProdutoVendaPag)

	campos, _, err := p.ValidarCampos(&t)
	if err != nil {
		return res, oops.Err(err)
	}

	consultaPrevia := pg.DB.Builder.
		Select(campos...).
		From("t_produto_venda")

	clausulaWhere := p.CriarFiltros(consultaPrevia, map[string]util.Filtro{
		"nao_contem_id": util.CriarFiltros("id", util.FlagFiltroNotIn),
		"id":            util.CriarFiltros("id = ?", util.FlagFiltroEq),
		"venda_id":      util.CriarFiltros("venda_id = ?", util.FlagFiltroEq),
		"produto_id":    util.CriarFiltros("produto_id = ?", util.FlagFiltroEq),
	})

	dados, prox, total, err := util.ConfigurarPaginacao(p, &t, &clausulaWhere)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Dados, res.Prox, res.Total = dados.([]produto_venda.ProdutoVenda), prox, total

	return
}

// Adicionar adiciona um novo produto_venda ao banco de dados do postgres
func (pg *PGProdutoVenda) Adicionar(req *produto_venda.ProdutoVenda) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_produto_venda").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Alterar altera um produto_venda no banco de dados do postgres
func (pg *PGProdutoVenda) Alterar(req *produto_venda.ProdutoVenda) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_produto_venda").
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

// Remover remove um produto_venda no banco de dados do postgres
func (pg *PGProdutoVenda) Remover(produtoVendaID int64) (err error) {
	resultado, err := pg.DB.Builder.
		Delete("t_produto_venda").
		Where(squirrel.Eq{
			"id": produtoVendaID,
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
