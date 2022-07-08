package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/projetoBase/health"

	"github.com/projetoBase/config"

	"github.com/projetoBase/oops"
	"github.com/projetoBase/util"
	"github.com/gin-gonic/gin"
)

// Metricas é um middleware que disponibiliza os dados de uma requisição
// nos logs
func Metricas() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			output := map[string]interface{}{
				"data":               time.Now(),
				"status_code":        params.StatusCode,
				"cliente_ip":         params.ClientIP,
				"metodo":             params.Method,
				"path":               params.Path,
				"latencia":           float64(params.Latency) / float64(time.Millisecond),
				"cliente_user_agent": params.Request.UserAgent(),
				"log_tipo":           "acesso",
				"router":             "external",
			}

			exists := true
			if _, set := params.Keys["sessao"]; set {
				// if se, ok := s.(brisauth.Session); ok {
				// 	output["usuario_id"] = se.ID
				// 	output["usuario_nome"] = se.Name
				// 	output["usuario"] = se.Username
				// } else {
				// 	exists = false
				// }
			} else {
				exists = false
			}

			if !exists {
				output["usuario_id"] = 0
				output["usuario_nome"] = "Usuário não logado"
				output["usuario"] = "Indefinido"
			}

			if v, set := params.Keys["error"]; set {
				var (
					err = v.(error)
					e   *oops.Error
				)
				if errors.As(err, &e) {
					output["erro"] = e.Error()
					output["erro_codigo"] = e.Code
					output["trace"] = e.Trace
					output["causa"] = e.Err.Error()
					output["log_tipo"] = "erro"
				}
			}

			if params.StatusCode == 500 {
				output["erro_fatal"] = "Possible panic"
				output["log_tipo"] = "error"
			}

			if v, set := params.Keys["metrics"]; set {
				if v2, ok := v.(*util.JSONB); ok {
					output["metrica_falha"] = v2
				}
			}

			if strings.Split(params.Path, "/")[0] == "api" {
				output["router"] = "internal"
			}

			b, err := json.Marshal(output)
			if err != nil {
				log.Printf("%+v\n", output)
			}

			health.HandledRequest(
				time.Duration(float64(params.Latency)/float64(time.Millisecond)),
				output["log_tipo"] == "erro",
			)

			return string(b) + string('\n')
		},
		Output: func() io.Writer {
			accessFile, err := os.OpenFile(config.ObterConfiguracao().DiretorioLogAcesso, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				log.Fatal(err)
			}

			gin.DisableConsoleColor()
			return io.MultiWriter(accessFile)
		}(),
	})
}

// GinZap adiciona um middleware customizado do zap
func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Next()
		latency := time.Since(t1)

		fields := []zap.Field{
			zap.Time("date", time.Now()),
			zap.Int("status_code", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Float64("latency", float64(latency)/float64(time.Millisecond)),
			zap.String("client_user_agent", c.Request.UserAgent()),
			zap.String("log_tipo", "access"),
			zap.String("handler", strings.Join(strings.Split(strings.Replace(c.HandlerName(), config.ObterConfiguracao().PermissaoBase, "", 1), "/")[1:], "/")),
			zap.String("request_id", c.Value("RID").(string)),
		}

		var (
			usuarioID      = int64(0)
			uNome, usuario = "Usuário não logado", "Indefinido"
		)

		if strings.Split(c.Request.URL.Path, "/")[0] == "api" {
			zap.String("router", "internal")
		} else {
			zap.String("router", "external")
		}

		fields = append(fields, []zap.Field{
			zap.Int64("user_id", usuarioID),
			zap.String("username", uNome),
			zap.String("usuario", usuario),
		}...)

		Erro := false
		if v, set := c.Keys["error"]; set {
			var (
				err = v.(error)
				e   *oops.Error
			)
			if errors.As(err, &e) {
				fields = append(fields, []zap.Field{
					zap.Int("error_code", e.Code),
					zap.String("error", e.Error()),
					zap.String("cause", e.Err.Error()),
					zap.Strings("trace", e.Trace),
				}...)
				Erro = true
			}
		}

		health.HandledRequest(
			time.Duration(float64(latency)/float64(time.Millisecond)),
			Erro,
		)

		if Erro {
			logger.Error("tratamento da requisição falhou", fields...)
		} else {
			logger.Info("requisição tratada", fields...)
		}
	}
}
