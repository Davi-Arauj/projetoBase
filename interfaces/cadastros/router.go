package cadastros

import (
	"github.com/gin-gonic/gin"
	"github.com/projetoBase/interfaces/cadastros/cliente"
	"github.com/projetoBase/interfaces/cadastros/produto"
	"github.com/projetoBase/interfaces/cadastros/produto_venda"
	"github.com/projetoBase/interfaces/cadastros/usuario"
	"github.com/projetoBase/interfaces/cadastros/venda"
)

// Router é um agregador de todos os routers de
// cadastros
func Router(r *gin.RouterGroup) {
	// Grupamento das rotas de clientes em cadastros
	cliente.Router(r.Group("clientes"))
	cliente.RouterID(r.Group("cliente"))

	// Grupamento das rotas de produtos em cadastros
	produto.Router(r.Group("produtos"))
	produto.RouterID(r.Group("produto"))

	// Grupamento das rotas de produto_vendas em cadastros
	produto_venda.Router(r.Group("produto_vendas"))
	produto_venda.RouterID(r.Group("produto_venda"))

	// Grupamento das rotas de usuarios em cadastros
	usuario.Router(r.Group("usuarios"))
	usuario.RouterID(r.Group("usuario"))

	// Grupamento das rotas de vendas em cadastros
	venda.Router(r.Group("vendas"))
	venda.RouterID(r.Group("venda"))
}
