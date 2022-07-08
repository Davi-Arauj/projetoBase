package middleware

import (
	"encoding/json"
)

// Sessao é usado para guarnecer os dados do usuário logado
type Sessao struct {
	Administrador PerlBool `json:"administrador"`
	UsuarioID     int64    `json:"id"`
	FuncionarioID int64    `json:"funcionario_id"`
	Email         string   `json:"email"`
	Name          string   `json:"nome"`
	UsuarioNome   string   `json:"login"`
	CidadeID      int      `json:"cidade_id"`
	Telefone      string   `json:"telefone"`
	Token         string   `json:"token"`
}

// PerlBool é um utilitário para interpretar booleanos da linguagem Perl
type PerlBool bool

// UnmarshalJSON implementa a interface para encoding/json
func (p *PerlBool) UnmarshalJSON(b []byte) error {
	var tmp interface{}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	if val, ok := tmp.(int); ok {
		if val > 0 {
			*p = true
		}
	}

	if val, ok := tmp.(int64); ok {
		if val > 0 {
			*p = true
		}
	}

	if val, ok := tmp.(float64); ok {
		if val > 0.0 {
			*p = true
		}
	}

	if val, ok := tmp.(bool); ok {
		*p = PerlBool(val)
	}

	return nil
}

// PermissaoDados é um estrutura base para permissões
type PermissaoDados struct {
	UsuarioID  int64    `json:"user_id"`
	Admin      bool     `json:"admin"`
	Permissoes []string `json:"permissions"`
}
