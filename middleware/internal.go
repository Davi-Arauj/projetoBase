package middleware

// // ModoAdministrador permite o acesso apenas para adminstradores
// func ModoAdministrador() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if sess, err := BuscarSessao(c); err != nil || !sess.Administrador {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 	}
// }
