package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/projetoBase/config"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// Postgres representa um tipo de banco de dados que implementa a interface
// do banco de dados
type Postgres struct {
	db               *sql.DB
	transacaoTimeout int
}

// AbrirConexao abre conexões com os bancos presentes na configuração
func (p *Postgres) AbrirConexao(c *config.BancoDados) (err error) {
	driverConfig := stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			RuntimeParams: map[string]string{
				"application_name": "chamados2",
				"DateStyle":        "ISO",
				"IntervalStyle":    "iso_8601",
				"search_path":      "vader, public",
			},
		},
	}
	stdlib.RegisterDriverConfig(&driverConfig)

	db, err := sql.Open("pgx", driverConfig.ConnectionString("postgres://"+c.Username+":"+c.Password+"@"+c.Host+":"+c.Port+"/"+c.Name))
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(c.MaxIdle)
	db.SetMaxOpenConns(c.MaxConn)
	db.SetConnMaxLifetime(time.Second * 60)

	p.db = db
	p.transacaoTimeout = c.TransactionTimeout

	return nil
}

// FecharConexao fecha a conexão
func (p *Postgres) FecharConexao() {
	_ = p.db.Close()
}

// NewTx abre uma nova transação em uma conexão aberta
func (p *Postgres) NewTx(ctx context.Context, opcoes *sql.TxOptions) (interface{}, error) {
	var (
		tx  *sql.Tx
		err error
	)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-time.After(time.Duration(p.transacaoTimeout+1) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	tx, err = p.db.BeginTx(ctx, opcoes)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
