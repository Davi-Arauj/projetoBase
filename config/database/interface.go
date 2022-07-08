package database

import (
	"context"
	"database/sql"
	"github.com/projetoBase/config"
)

// BancoDados é a interface para multiplos banco de dados
type BancoDados interface {
	AbrirConexao(*config.BancoDados) error
	FecharConexao()
	NewTx(context.Context, *sql.TxOptions) (interface{}, error)
}

// Transacao é a interface para transacao dos bancos de dados
type Transacao interface {
	Commit() error
	Rollback() error
}
