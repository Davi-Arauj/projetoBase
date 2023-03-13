package usuario

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/util"
)

// Router é um router para as rotas de usuarios que não utilizam ID
func Router(r *gin.RouterGroup) {

	r.GET("", util.AddRota("Lista usuarios", "Lista usuarios", listar)...)
	r.POST("", util.AddRota("Adicionar usuario", "Adicionar usuario", adicionar)...)

}

// RouterID é um router para as rotas de usuarios que utilizam ID
func RouterID(r *gin.RouterGroup) {
	r.PUT(":codigo_barras", util.AddRota("Altera um usuario", "Altera um usuario", alterar)...)
	r.DELETE(":codigo_barras", util.AddRota("Deleta um usuario", "Deleta um usuario", remover)...)
}
