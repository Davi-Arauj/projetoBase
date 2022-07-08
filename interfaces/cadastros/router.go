package cadastros

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/interfaces/cadastros/produto"
)

func Router(r *gin.RouterGroup){
	produto.Router(r.Group("produtos"))
	produto.RouterID(r.Group("produto"))
}