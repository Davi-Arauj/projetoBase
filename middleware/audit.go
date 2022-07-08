package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/projetoBase/config/database"

	"github.com/projetoBase/config"

	"github.com/gin-gonic/gin"
)

// Auditoria é um middleware que registra as ações no banco de dados
func Auditoria() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			return
		}

		auditoriaHandler := getContentHandler(c.ContentType())
		if auditoriaHandler == nil {
			c.Abort()
			c.JSON(400, gin.H{
				"error": "O content-type da requisição não é válido ou não foi especificado",
			})
			return
		}

		// uma copia do corpo é necessária
		corpo, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			corpo = nil
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(corpo))

		// adicione a resposta atual um buffer para acessar os dados retornados
		c.Writer = &EmpacotadorEscrita{c.Writer, &bytes.Buffer{}}

		c.Next()

		var (
			ModuloAcao, Modulo string
			handler            []string
		)

		handler = strings.Split(strings.Replace(c.HandlerName(), config.ObterConfiguracao().PermissaoBase, "", 1), "/")

		if strings.Split(handler[0], ".")[0] == "middleware" {
			ModuloAcao = "Undefined"
			Modulo = "Undefined"
		} else {
			ModuloAcao = strings.Join(handler[2:], "/")
			Modulo = strings.Join(handler[:2], "/")
		}

		a := &auditoria{
			Acao:          c.Request.Method,
			URL:           c.Request.URL.Path,
			RetornoStatus: c.Writer.Status(),
			IP:            strings.ReplaceAll(strings.ReplaceAll(strings.Join(strings.Split(c.Request.RemoteAddr, ":")[:len(strings.Split(c.Request.RemoteAddr, ":"))-1], ":"), "[", ""), "]", ""),
			ModuloAcao:    ModuloAcao,
			Modulo:        Modulo,
		}

		p, _ := json.Marshal(c.Request.URL.Query())
		a.URLParametros = string(p)

		a.ConteudoRequisicao = "{}"
		if conteudo, err := auditoriaHandler(corpo); err == nil {
			a.ConteudoRequisicao = conteudo
		}

		saida := c.Writer.(*EmpacotadorEscrita)
		dados, _ := ioutil.ReadAll(saida.buff)
		a.RetornoConteudo = string(dados)
		if a.RetornoConteudo == "" {
			a.RetornoConteudo = "{}"
		}

		a.UsuarioCodigo = 0
		a.UsuarioNome = "Usuario-nao-logado"
		if err := a.Save(); err != nil {
			dados, _ := json.Marshal(a)
			log.Printf("[Erro-auditoria]: %s | %s", err.Error(), string(dados))
		}
	}
}

//
// EmpacotadorEscrita é um empacotador de escrita que intercepta o corpo da requisição
type EmpacotadorEscrita struct {
	gin.ResponseWriter
	buff *bytes.Buffer
}

// Escreva escreve dados no escritor interno e no buffer intermediario
func (w *EmpacotadorEscrita) Escreva(p []byte) (int, error) {
	w.buff.Write(p)
	return w.ResponseWriter.Write(p)
}

type auditoria struct {
	Modulo             string `sql:"modulo"`
	ModuloAcao         string `sql:"modulo_acao"`
	UsuarioCodigo      int    `sql:"usuario_codigo"`
	UsuarioNome        string `sql:"usuario_nome"`
	Acao               string `sql:"acao"`
	IP                 string `sql:"ip"`
	URL                string `sql:"url"`
	URLParametros      string `sql:"url_param"`
	ConteudoRequisicao string `sql:"conteudo"`
	RetornoStatus      int    `sql:"retorno_status"`
	RetornoConteudo    string `sql:"retorno_conteudo"`
}

// Save salva um registro de auditoria no banco de dados
func (a *auditoria) Save() error {
	tx, err := database.NovaTransacao(context.Background(), false)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	result, err := tx.Builder.
		Insert("t_auditoria").
		Columns("modulo", "modulo_acao", "usuario_codigo", "usuario_nome", "acao", "ip", "url", "url_param",
			"conteudo", "retorno_status", "retorno_conteudo").
		Values(a.Modulo, a.ModuloAcao, a.UsuarioCodigo, a.UsuarioNome, a.Acao, a.IP, a.URL, a.URLParametros, a.ConteudoRequisicao, a.RetornoStatus, a.RetornoConteudo).
		Exec()
	if err != nil {
		return err
	}

	if c, err := result.RowsAffected(); err != nil {
		return err
	} else if c == 0 {
		return errors.New("Nenhuma linha inserida")
	}

	return tx.Commit()
}

// auditeHandleJSON formata um payload para auditoria
func auditeHandlerJSON(corpo []byte) (string, error) {
	dados := make(map[string]interface{})
	if err := json.Unmarshal(corpo, &dados); err != nil {
		return "", err
	}
	return string(corpo), nil
}

// auditeHandlerEmpty usado quando a o conteúdo da requisição não pode
// ser formatado para auditoria
func auditeHandlerEmpty(corpo []byte) (string, error) {
	return "{}", nil
}

// getContentHandler seleciona um handler de auditoria para ser usado
// através do tipo que conteúdo da requisição
func getContentHandler(contentType string) (handler func([]byte) (string, error)) {
	switch contentType {
	case "application/json":
		return auditeHandlerJSON
	case "multipart/form-data":
		return auditeHandlerEmpty
	}
	return nil
}
