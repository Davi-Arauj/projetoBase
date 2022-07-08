package logger

import (
	"github.com/projetoBase/config"
	"github.com/projetoBase/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupLogger initialize zap logger
func SetupLogger() (*zap.Logger, error) {
	c := config.ObterConfiguracao()
	var cfg zap.Config

	if c.Producao {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = true
		cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	cfg.Encoding = "json"

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.MessageKey = "mensagem"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.NameKey = "nome"
	cfg.EncoderConfig.TimeKey = "tempo"
	cfg.EncoderConfig.StacktraceKey = "stack_trace"
	cfg.InitialFields = map[string]interface{}{
		"aplicacao": "academia",
		"versao":    middleware.Versao,
	}

	cfg.OutputPaths = []string{c.DiretorioLogAcesso, "stdout"}
	cfg.ErrorOutputPaths = []string{c.DiretorioLogErro, "stderr"}

	return cfg.Build()
}

// PanicRecovery handles recovered panics
func PanicRecovery(p interface{}) (err error) {
	zap.S().Error(
		"PANIC detected: ",
		p,
	)

	return
}
