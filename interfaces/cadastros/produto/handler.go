package produto

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projetoBase/application/cadastros/produto"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// listar - listagem de produtos
func listar(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := produto.Listar(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

// buscar - busca um produto usando como parametro o codigo de barras
func buscar(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))

	res, err := produto.Buscar(c, int64(codigoBarras))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	c.JSON(200, res)
}

// adicionar - adiciona um produto
func adicionar(c *gin.Context) {
	var req produto.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := produto.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

// alterar - altera um produto
func alterar(c *gin.Context) {
	var req produto.Req

	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := produto.Alterar(c, int64(codigoBarras), &req); err != nil {
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
	if err := produto.Remover(c, int64(codigoBarras)); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

// total - busca o total de produtos em uma listagem
func total(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := produto.Total(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}