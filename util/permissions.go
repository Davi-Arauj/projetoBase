package util


import (
	"reflect"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	permissoes = make(map[string]Handler)
)

// Handler contem informação adicional sobre o handler
type Handler struct {
	Descricao    string
	NomeAmigavel string
}

// AddRota cria uma lista de funções para auxilio na criação das permissões
func AddRota(descricao, nomeAmigavel string, handlers ...gin.HandlerFunc) (out []gin.HandlerFunc) {
	if len(handlers) == 0 {
		return out
	}

	permissoes[runtime.FuncForPC(reflect.ValueOf(handlers[len(handlers)-1]).Pointer()).Name()] = Handler{
		Descricao:    descricao,
		NomeAmigavel: nomeAmigavel,
	}

	return handlers
}

// ObterNomePermissao retorna uma descricao e nome amigável para o handler
func ObterNomePermissao(handler string) Handler {
	return permissoes[handler]
}