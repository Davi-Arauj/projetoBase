package middleware

import (
	"regexp"

	"github.com/projetoBase/oops"
	"github.com/gin-gonic/gin"
)

const (
	// Desktop define a plataforma desktop
	Desktop = iota
	// Mobile define a plataforma mobile
	Mobile
)

const (
	// MobileRegex define o regex para validar um userAgent valido para plataforma mobile
	MobileRegex = `(Dalvik\/\d{1,}.\d{1,}.\d{1,} \(Linux; U; Android \d{1,}(.\d{1,}.\d{0,})?.*)`
)

// ValidarPlataforma é usado para validar plataformas
func ValidarPlataforma(c *gin.Context, plataforma int) error {
	s := c.Request.UserAgent()
	if s == "" {
		return oops.NovoErr("Não foi possível detectar o UserAgent na requisição. Entre em contato com o suporte")
	}

	match, err := regexp.Match(MobileRegex, []byte(s))
	if err != nil {
		return oops.Err(err)
	}

	if plataforma == Desktop && !match {
		return nil
	}

	if plataforma == Mobile && match {
		return nil
	}

	return &oops.ErrInvalidPlatform
}
