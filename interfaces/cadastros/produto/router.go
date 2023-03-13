package produto

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/util"
)

// Router é um router para as rotas de produtos que não utilizam ID
func Router(r *gin.RouterGroup) {

	r.GET("", util.AddRota("Lista produtos", "Lista produtos", listar)...)
	r.POST("", util.AddRota("Adicionar produto", "Adicionar produto", adicionar)...)

}

// RouterID é um router para as rotas de produtos que utilizam ID
func RouterID(r *gin.RouterGroup) {
	r.PUT(":codigo_barras", util.AddRota("Altera um produto", "Altera um produto", alterar)...)
	r.DELETE(":codigo_barras", util.AddRota("Deleta um produto", "Deleta um produto", remover)...)
}
