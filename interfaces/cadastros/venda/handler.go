package venda

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projetoBase/application/cadastros/venda"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// listar - listagem de vendas
func listar(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := venda.Listar(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

// buscar - busca um venda usando como parametro o codigo de barras
func buscar(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))

	res, err := venda.Buscar(c, int64(codigoBarras))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	c.JSON(200, res)
}

// adicionar - adiciona um venda
func adicionar(c *gin.Context) {
	var req venda.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := venda.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

// alterar - altera um venda
func alterar(c *gin.Context) {
	var req venda.Req

	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := venda.Alterar(c, int64(codigoBarras), &req); err != nil {
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
	if err := venda.Remover(c, int64(codigoBarras)); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

// total - busca o total de vendas em uma listagem
func total(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := venda.Total(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}