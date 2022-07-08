package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/produto"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// PGproduto é uma estrutura base para acesso aos metodos do banco postgres para manipulação de produtos
type PGProduto struct {
	DB *database.DBTransacao
}

// Listar lista de forma paginada os dados de produtos do banco postgres
func (pg *PGProduto) Listar(p *util.ParametrosRequisicao) (res *produto.ProdutoPag, err error) {
	var t produto.Produto

	res = new(produto.ProdutoPag)

	campos, _, err := p.ValidarCampos(&t)
	if err != nil {
		return res, oops.Err(err)
	}

	consultaPrevia := pg.DB.Builder.
		Select(campos...).
		From("t_produtos")

	clausulaWhere := p.CriarFiltros(consultaPrevia, map[string]util.Filtro{
		"codigo_barras": util.CriarFiltros("codigo_barras = ?::BIGINT", util.FlagFiltroEq),
		"nao_contem_id": util.CriarFiltros("id", util.FlagFiltroNotIn),
		"nome":          util.CriarFiltros("lower(public.unaccent(nome)) LIKE lower('%'||public.unaccent(?)||'%')", util.FlagFiltroEq),
		"nome_exato":    util.CriarFiltros("public.unaccent(nome) LIKE public.unaccent(?)", util.FlagFiltroEq),
	})

	dados, prox, total, err := util.ConfigurarPaginacao(p, &t, &clausulaWhere)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Dados, res.Prox, res.Total = dados.([]produto.Produto), prox, total

	return
}

// Buscar busca um novo produto no banco de dados do postgres
func (pg *PGProduto) Buscar(req *produto.Produto) (err error) {
	if err = pg.DB.Builder.
		Select(`id, codigo_barras, nome,descricao,endereco_foto,valor_pago,valor_venda,quantidade,unidade_id,categoria_id,subcategoria_id,data_criacao,data_atualizacao`).
		From(`t_produtos`).
		Where(squirrel.Eq{
			"codigo_barras": req.CodigoBarras,
		}).
		Scan(
			&req.ID, &req.CodigoBarras, &req.Nome, &req.Descricao, &req.Foto, &req.Valorpago, &req.Valorvenda, &req.Qtde, &req.UndCod, &req.CatCod, &req.ScatCod, &req.DataCriacao, &req.DataAtualizacao,
		); err != nil {
		return oops.Err(err)
	}
	return
}

// Adicionar adiciona um novo produto ao banco de dados do postgres
func (pg *PGProduto) Adicionar(req *produto.Produto) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_produtos").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}

	return
}

// Alterar altera um produto no banco de dados do postgres
func (pg *PGProduto) Alterar(req *produto.Produto) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_produtos").
		SetMap(valores).
		Where(squirrel.Eq{
			"codigo_barras": req.CodigoBarras,
		}).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}

// Remover remove um produto no banco de dados do postgres
func (pg *PGProduto) Remover(codigoBarras int64) (err error) {
	resultado, err := pg.DB.Builder.
		Delete("t_produtos").
		Where(squirrel.Eq{
			"codigo_barras": codigoBarras,
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
