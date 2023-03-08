package venda

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/util"
)

// Router é um router para as rotas de vendas que não utilizam ID
func Router(r *gin.RouterGroup) {

	r.GET("", util.AddRota("Lista vendas", "Lista vendas", listar)...)
	r.POST("", util.AddRota("Adicionar venda", "Adicionar venda", adicionar)...)
	r.GET("/total", util.AddRota("Total vendas", "Total vendas", total)...)

}

// RouterID é um router para as rotas de vendas que utilizam ID
func RouterID(r *gin.RouterGroup) {
	r.GET(":codigo_barras", util.AddRota("Busca um venda", "Busca um venda", buscar)...)
	r.PUT(":codigo_barras", util.AddRota("Altera um venda", "Altera um venda", alterar)...)
	r.DELETE(":codigo_barras", util.AddRota("Deleta um venda", "Deleta um venda", remover)...)
}