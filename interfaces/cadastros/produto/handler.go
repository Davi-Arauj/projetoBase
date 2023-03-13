package produto

import (
	"net/http"
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

	c.JSON(http.StatusOK, res)
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

	c.JSON(http.StatusCreated, id)
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

	req.CodigoBarras = util.PonteiroInt64(int64(codigoBarras))
	if err := produto.Alterar(c, &req); err != nil {
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
	if err := produto.Remover(c, int64(codigoBarras)); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
