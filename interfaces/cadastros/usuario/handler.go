package usuario

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projetoBase/application/cadastros/usuario"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// listar - listagem de usuarios
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

	c.JSON(200, res)
}

// buscar - busca um usuario usando como parametro o codigo de barras
func buscar(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))

	res, err := usuario.Buscar(c, int64(codigoBarras))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	c.JSON(200, res)
}

// adicionar - adiciona um usuario
func adicionar(c *gin.Context) {
	var req usuario.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := usuario.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

// alterar - altera um usuario
func alterar(c *gin.Context) {
	var req usuario.Req

	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := usuario.Alterar(c, int64(codigoBarras), &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

// remover - remove um prodtuo
func remover(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := usuario.Remover(c, int64(codigoBarras)); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

// total - busca o total de usuarios em uma listagem
func total(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := usuario.Total(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}