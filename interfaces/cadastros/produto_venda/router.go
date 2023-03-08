package produto_venda

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/util"
)

// Router é um router para as rotas de produto_vendas que não utilizam ID
func Router(r *gin.RouterGroup) {

	r.GET("", util.AddRota("Lista produto_vendas", "Lista produto_vendas", listar)...)
	r.POST("", util.AddRota("Adicionar produto_venda", "Adicionar produto_venda", adicionar)...)
	r.GET("/total", util.AddRota("Total produto_vendas", "Total produto_vendas", total)...)
}

// RouterID é um router para as rotas de produto_vendas que utilizam ID
func RouterID(r *gin.RouterGroup) {
	r.GET(":codigo_barras", util.AddRota("Busca um produto_venda", "Busca um produto_venda", buscar)...)
	r.PUT(":codigo_barras", util.AddRota("Altera um produto_venda", "Altera um produto_venda", alterar)...)
	r.DELETE(":codigo_barras", util.AddRota("Deleta um produto_venda", "Deleta um produto_venda", remover)...)
}