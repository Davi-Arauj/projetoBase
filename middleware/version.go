package middleware

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/projetoBase/config/database"

	"github.com/gin-gonic/gin"
)

var (
	// Versao possui informação sobre a versão
	Versao = ""
)

// VersaoInfo adicionar informaçãoes sobre a versão
func VersaoInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Versao == "" {
			Versao = strconv.FormatInt(time.Now().Unix(), 10)
		}
		c.Writer.Header().Set("Applicacao-Versao", Versao)
	}
}

// InicializarVersionamento inicializa o versionamento do sistema e notifica o usuário sobre
// a nova versão
func InicializarVersionamento() (err error) {
	if Versao == "" {
		return nil
	}

	tx, err := database.NovaTransacao(context.Background(), false)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var (
		versaoAnterior = new(string)
	)

	if err = tx.Builder.
		Select("versao").
		From("t_versao").
		OrderBy("data_criacao DESC").
		Limit(1).
		QueryRow().
		Scan(versaoAnterior); err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows || *versaoAnterior == "" || *versaoAnterior != Versao {
		if err = tx.Builder.
			Insert("t_versao").
			Columns("versao").
			Values(Versao).
			Suffix("RETURNING id").
			Scan(new(int)); err != nil {
			return
		}

		// if err = notify.EnviarBroadcast(notify.Notificacao{
		// 	CreatedBy: utils.PonteiroInt64(1),
		// 	Title:     utils.PonteiroString("Nova versão do sistema"),
		// 	Body:      utils.PonteiroString("A versão " + Versao + " do sistema já está disponivel, clique aqui para atualizar"),
		// 	Parameters: &utils.JSONB{
		// 		"version": Versao,
		// 		"alert": utils.JSONB{
		// 			"type": "warning",
		// 			"text": "Nova versão do sistema disponível, para atualizar, clique no botão abaixo",
		// 			"options": []utils.JSONB{
		// 				{
		// 					"type":   "primary",
		// 					"name":   "Atualizar",
		// 					"action": "reload",
		// 				},
		// 				{
		// 					"type":   "danger",
		// 					"name":   "Cancelar",
		// 					"action": "cancel",
		// 				},
		// 			},
		// 		},
		// 	},
		// 	Notify: utils.PonteiroBool(true),
		// }); err != nil {
		// 	return
		// }

		if _, err = tx.Exec(`
			UPDATE t_notificacao
			SET data_leitura = now()
			WHERE data_leitura IS NULL
			  AND (parametros->>'version')::TEXT IS NOT NULL
			  AND (parametros->>'version')::TEXT <>$1::TEXT
		`, Versao); err != nil {
			return
		}
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
