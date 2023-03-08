package cliente

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projetoBase/application/cadastros/cliente"
	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
)

// listar - listagem de clientes
func listar(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := cliente.Listar(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusOK, res)
}

// buscar - busca um cliente usando como parametro o codigo de barras
func buscar(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))

	res, err := cliente.Buscar(c, int64(codigoBarras))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	c.JSON(http.StatusOK, res)
}

// adicionar - adiciona um cliente
func adicionar(c *gin.Context) {
	var req cliente.Req
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := cliente.Adicionar(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusCreated, id)
}

// alterar - altera um cliente
func alterar(c *gin.Context) {
	var req cliente.Req

	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := cliente.Alterar(c, int64(codigoBarras), &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// remover - remove um prodtuo
func remover(c *gin.Context) {
	codigoBarras, err := strconv.Atoi(c.Param("codigo_barras"))
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	if err := cliente.Remover(c, int64(codigoBarras)); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// total - busca o total de clientes em uma listagem
func total(c *gin.Context) {

	p, err := util.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}
	res, err := cliente.Total(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusOK, res)
}