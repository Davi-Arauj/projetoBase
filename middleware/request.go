package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// IdentificadorRequisicao adiciona um UUIDv4 em todos os contextos para melhor
// visualização do fluxo
func IdentificadorRequisicao() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("RID", uuid.New().String())
	}
}
