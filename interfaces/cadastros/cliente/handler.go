package cliente

import (
	"net/http"

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

	email := (c.Param("email"))

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := cliente.Alterar(c, email, &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// remover - remove um prodtuo
func remover(c *gin.Context) {
	email := c.Param("email")

	if err := cliente.Remover(c, email); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
