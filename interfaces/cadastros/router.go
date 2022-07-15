package cadastros

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/interfaces/cadastros/produto"
)

// Router é um agregador de todos os routers de
// cadastros
func Router(r *gin.RouterGroup){
	// Grupamento das rotas de produtos em cadastros
	produto.Router(r.Group("produtos"))
	produto.RouterID(r.Group("produto"))
}