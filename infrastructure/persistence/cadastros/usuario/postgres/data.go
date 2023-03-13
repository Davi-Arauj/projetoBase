package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/projetoBase/config/database"
	"github.com/projetoBase/infrastructure/persistence/cadastros/usuario"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// PGUsuario é uma estrutura base para acesso aos metodos do banco postgres para manipulação de usuarios
type PGUsuario struct {
	DB *database.DBTransacao
}

// Listar lista de forma paginada os dados de usuarios do banco postgres
func (pg *PGUsuario) Listar(p *util.ParametrosRequisicao) (res *usuario.UsuarioPag, err error) {
	var t usuario.Usuario

	res = new(usuario.UsuarioPag)

	campos, _, err := p.ValidarCampos(&t)
	if err != nil {
		return res, oops.Err(err)
	}

	consultaPrevia := pg.DB.Builder.
		Select(campos...).
		From("t_usuario")

	clausulaWhere := p.CriarFiltros(consultaPrevia, map[string]util.Filtro{
		"nao_contem_id": util.CriarFiltros("id", util.FlagFiltroNotIn),
		"nome":          util.CriarFiltros("lower(public.unaccent(nome)) LIKE lower('%'||public.unaccent(?)||'%')", util.FlagFiltroEq),
		"nome_exato":    util.CriarFiltros("public.unaccent(nome) LIKE public.unaccent(?)", util.FlagFiltroEq),
		"email":         util.CriarFiltros("email =?", util.FlagFiltroEq),
		"id":            util.CriarFiltros("id =?", util.FlagFiltroEq),
	})

	dados, prox, total, err := util.ConfigurarPaginacao(p, &t, &clausulaWhere)
	if err != nil {
		return res, oops.Err(err)
	}

	res.Dados, res.Prox, res.Total = dados.([]usuario.Usuario), prox, total

	return
}

// Adicionar adiciona um novo usuario ao banco de dados do postgres
func (pg *PGUsuario) Adicionar(req *usuario.Usuario) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("t_usuario").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		Scan(&req.ID); err != nil {
		return oops.Err(err)
	}
	return
}

// Alterar altera um usuario no banco de dados do postgres
func (pg *PGUsuario) Alterar(req *usuario.Usuario) (err error) {
	cols, vals, err := util.FormatarInsertUpdate(req)
	if err != nil {
		return oops.Err(err)
	}

	valores := make(map[string]interface{})
	for indice, elemento := range cols {
		valores[elemento] = vals[indice]
	}

	if err = pg.DB.Builder.
		Update("t_usuario").
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

// Remover remove um usuario no banco de dados do postgres
func (pg *PGUsuario) Remover(email string) (err error) {
	resultado, err := pg.DB.Builder.
		Delete("t_usuario").
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
