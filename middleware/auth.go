package middleware

// // Autenticacao é um midleware que checa se a requisição contém um cookie valido
// func Autenticacao() gin.HandlerFunc {
// 	return brisauth.AuthorizationGin(config.ObterConfiguracao().Secrets)
// }

// // BuscarSessao retorna a sessão para um determinado contexto
// func BuscarSessao(c *gin.Context) (sess *Sessao, err error) {
// 	var bsess *brisauth.Session
// 	if bsess, err = brisauth.GetGinSession(c); err != nil {
// 		return nil, err
// 	}

// 	sess = &Sessao{
// 		UsuarioNome:   *bsess.Username,
// 		Name:          *bsess.Name,
// 		FuncionarioID: *bsess.EmployeeID,
// 		Token:         *bsess.Token,
// 		UsuarioID:     *bsess.ID,
// 	}

// 	if bsess.Email == nil {
// 		bsess.Email = utils.PonteiroString("")
// 	}
// 	sess.Email = *bsess.Email

// 	if bsess.Phone == nil {
// 		bsess.Phone = utils.PonteiroInt64(0)
// 	}

// 	sess.Administrador = PerlBool(*bsess.Administrator)
// 	sess.Telefone = strconv.FormatInt(*bsess.Phone, 10)
// 	sess.CidadeID = int(*bsess.CityID)

// 	return
// }
