package usuario

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/projetoBase/application/cadastros/usuario"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// listar godoc
// @Summary Lista usuarios
// @Description Lista usuarios dados os seus devidos filtros
// @Tags cadastros/usuario
// @Param id query string false "ID do usuario"
// @Param nome query string false "Nome do usuarios"
// @Param email query string false "Email do usuarios"
// @Produce json
// @Success 200 {object} usuario.ResPag "OK. Prox define se existe ou não uma proxima pagina."
// @Router /v2/usuarios [get]
func listar(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := usuario.Listar(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusOK, res)
}

// adicionar godoc
// @Summary Adiciona um novo usuario
// @Description Adiciona um novo usuario dado o corpo deste novo usuario
// @Tags cadastros/usuario
// @Accept json
// @Produce json
// @Param tag body usuario.Req true "Dados do usuario"
// @Success 201 {string} id "id do novo usuario adicionado"
// @Router /v2/usuarios [post]
func adicionar(c *gin.Context) {
	var req usuario.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	novoUsuario, err := usuario.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusCreated, novoUsuario)
}

// alterar godoc
// @Summary Altera um usuario
// @Description Altera um usuario dado o seu id e o corpo do novo usuario
// @Tags cadastros/usuario
// @Accept json
// @Param usuario_id path integer true "ID do usuario a ser alterado"
// @Param tag body usuario.Req true "Dados do novo usuario"
// @Success 204 nil nil
// @Router /v2/usuario/{usuario_id} [put]
func alterar(c *gin.Context) {
	var req usuario.Req
	usuario_id := (c.Param("usuario_id"))

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := usuario.Alterar(c, usuario_id, &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// remover godoc
// @Summary Remove um usuario
// @Description Remove um usuario
// @Tags cadastros/usuario
// @Param usuario_id path integer true "ID do usuario a ser removido"
// @Success 204 nil nil
// @Router /v2/usuario/{usuario_id} [delete]
func remover(c *gin.Context) {
	email := (c.Param("email"))

	if err := usuario.Remover(c, email); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
