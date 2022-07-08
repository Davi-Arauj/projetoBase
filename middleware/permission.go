package middleware

// // ChecarPermissao é um middeleware para checar se o user-agent pode
// // executar uma ação
// func ChecarPermissao() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			bsess *brisauth.Session
// 			err   error
// 		)

// 		if bsess, err = brisauth.GetGinSession(c); err != nil {
// 			c.JSON(401, gin.H{
// 				"msg": "Sessão inválida",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		cfg := config.ObterConfiguracao()

// 		zap.L().Debug("sessao", zap.String("handler", c.HandlerName()), zap.String("base", cfg.PermissaoBase))

// 		if strings.Split(strings.Replace(c.HandlerName(), cfg.PermissaoBase, "", 1), ".")[0] == "middleware" {
// 			c.AbortWithStatus(404)
// 			return
// 		}

// 		if !*bsess.Administrator {
// 			perm := strings.Split(strings.Replace(c.HandlerName(), cfg.PermissaoBase, "", 1), "/")[1:]
// 			if len(perm) == 1 {
// 				perm = strings.Split(perm[0], ".")
// 			}
// 			if !brisauth.SessionHasPermission(
// 				perm[0]+"::"+strings.Join(perm[1:], "/"),
// 				*bsess.SessionID,
// 			) {
// 				c.JSON(403, gin.H{
// 					"msg": "Você não tem permissão para executar essa ação",
// 				})
// 				c.Abort()
// 				return
// 			}
// 		}

// 		c.Next()
// 	}
// }

// AutoPerm adiciona as permissões ao banco de dados através das rotas
// func AutoPerm(routes gin.RoutesInfo) {
// 	cfg := config.ObterConfiguracao()
// 	tx, err := database.NovaTransacaoRevan(context.Background(), false)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tx.Rollback()
// 	for i := 0; i < len(routes); i++ {
// 		perm := strings.Join(strings.Split(strings.Replace(routes[i].Handler, cfg.PermissaoBase, "", 1), "/")[1:], "/")
// 		dominio, nome := strings.Split(perm, "/")[0], strings.Join(strings.Split(perm, "/")[1:], "/")
// 		if dominio == "" {
// 			continue
// 		}

// 		var modulo int
// 		var dadosAmigavel = utils.ObterNomePermissao(routes[i].Handler)

// 		if err = tx.Builder.Select("id").
// 			From("public.t_autenticacao_modulo").
// 			Where(squirrel.Eq{"location": dominio, "autenticacao_sistema_id": 4}).
// 			Scan(&modulo); err == sql.ErrNoRows {
// 			if err = tx.Builder.Insert("public.t_autenticacao_modulo").
// 				Columns("nome", "location", "descricao", "autenticacao_sistema_id").
// 				Values(dominio, dominio, "", 4).
// 				Suffix("RETURNING id").
// 				Scan(&modulo); err != nil {
// 				log.Fatal(err)
// 			}
// 		} else if err != nil {
// 			log.Fatal(err)
// 		}

// 		_, err = tx.Builder.Insert("public.t_autenticacao_modulo_acao").
// 			Columns("autenticacao_modulo_id", "nome", "location", "descricao", "nome_amigavel").
// 			Values(
// 				modulo,
// 				routes[i].Method+"/"+routes[i].Path[1:],
// 				nome,
// 				dadosAmigavel.Descricao,
// 				dadosAmigavel.NomeAmigavel,
// 			).
// 			Suffix(
// 				"ON CONFLICT (autenticacao_modulo_id, nome) DO UPDATE SET descricao = ?, nome_amigavel = ?, location = ?",
// 				dadosAmigavel.Descricao, dadosAmigavel.NomeAmigavel, nome,
// 			).
// 			Exec()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	var (
// 		rows  *sql.Rows
// 		dados []utils.JSONBA
// 	)

// 	if rows, err = tx.Query(`
// 		SELECT
// 			TAM.location || '::' || TAMA.location AS location,
// 			ARRAY_TO_JSON(ARRAY_AGG(json_build_object('id', TAMA.id, 'nome', TAMA.nome, 'location', TAM.location || '::' || TAMA.location)))
// 		FROM t_autenticacao_modulo TAM
// 		INNER JOIN t_autenticacao_modulo_acao TAMA ON TAMA.autenticacao_modulo_id = TAM.id
// 		WHERE TAM.autenticacao_sistema_id = 4
// 		GROUP BY 1
// 		HAVING COUNT(DISTINCT TAMA.indice) > 1
// 	`); err != nil && err != sql.ErrNoRows {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		var r utils.JSONBA
// 		if err = rows.Scan(new(string), &r); err != nil {
// 			log.Fatal(err)
// 		}

// 		dados = append(dados, r)
// 	}
// 	_ = rows.Close()

// 	if err = tx.Commit(); err != nil {
// 		log.Fatal(err)
// 	}

// 	if len(dados) > 0 {
// 		for i := range dados {
// 			zap.L().Warn("colisão de permissões detectada",
// 				zap.Array("permissões", dados[i]),
// 			)
// 		}
// 	}
// }
