package util

import (
	"fmt"
	"net/url"
)

// HTTPErro define um erro causado por uma requisição HTTP
type HTTPErro struct {
	// URL da requisição
	URL url.URL
	// Metodo da requisição
	Metodo string
	// StatusCode da requisição
	StatusCode int
	// Causa de um error HTTP
	Causa error
	// TagRequisicao identifica a tag que causou o erro
	TagRequisicao string
	// ClienteNome nome do cliente que tentou performar a requisição
	ClienteNome string
}

// Error implementa a interface de erro
func (e *HTTPErro) Error() string {
	return fmt.Sprintf("A requisição falhou para [%v] '%v' com status %v: %+v", e.Metodo, e.URL.String(), e.StatusCode, e.Causa)
}

// Unwrap implementa a interface par aunwrap
func (e *HTTPErro) Unwrap() error {
	return e.Causa
}
