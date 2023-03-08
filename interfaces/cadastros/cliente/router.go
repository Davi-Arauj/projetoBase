package cliente

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/util"
)

// Router é um router para as rotas de clientes que não utilizam ID
func Router(r *gin.RouterGroup) {

	r.GET("", util.AddRota("Lista clientes", "Lista clientes", listar)...)
	r.POST("", util.AddRota("Adicionar cliente", "Adicionar cliente", adicionar)...)
	r.GET("/total", util.AddRota("Total clientes", "Total clientes", total)...)

}

// RouterID é um router para as rotas de clientes que utilizam ID
func RouterID(r *gin.RouterGroup) {
	r.GET(":codigo_barras", util.AddRota("Busca um cliente", "Busca um cliente", buscar)...)
	r.PUT(":codigo_barras", util.AddRota("Altera um cliente", "Altera um cliente", alterar)...)
	r.DELETE(":codigo_barras", util.AddRota("Deleta um cliente", "Deleta um cliente", remover)...)
}
