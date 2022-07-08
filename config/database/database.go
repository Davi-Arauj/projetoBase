package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/projetoBase/config"
	"github.com/projetoBase/config/database/postgres"
)

const (
	// LimiteAvisoConsulta define a mensagem padrão para um aviso
	// de que o tempo limite de uma consulta foi atingido
	LimiteAvisoConsulta = "Tempo esperado para consulta foi excedido"
)

var (
	conexoes map[string]BancoDados
	roSuffix = "-ro"
)

// DBTransacao é usado para agregar transações para todos os
// banco de dados disponíveis
type DBTransacao struct {
	postgres *sql.Tx
	ctx      context.Context
	Builder  sq.StatementBuilderType
}

// AbrirConexao itera sobre os banco de dados indicados na configuração
// e tenta abrir uma conexão com eles
func AbrirConexao() error {
	inicializarMapaConexoes()

	conf := config.ObterConfiguracao()

	for d := 0; d < len(conf.BancosDados); d++ {
		connNome := conf.BancosDados[d].Nick
		if conf.BancosDados[d].ReadOnly {
			connNome += roSuffix
		}
		if _, set := conexoes[connNome]; set {
			if err := conexoes[connNome].AbrirConexao(&conf.BancosDados[d]); err != nil {
				return err
			}
		}
	}

	return nil
}

// FecharConexoes intera sobre o mapa de conexões abertas e
// as fecha
func FecharConexoes() {
	for _, v := range conexoes {
		v.FecharConexao()
	}
}

func inicializarMapaConexoes() {
	if conexoes != nil {
		return
	}
	conexoes = make(map[string]BancoDados)
	conexoes["academia"] = &postgres.Postgres{}
	conexoes["academia"+roSuffix] = &postgres.Postgres{}
}

// NovaTransacao tenta abrir uma nova transacao para a conexão
// com o banco de dados do academia
func NovaTransacao(ctx context.Context, apenasLeitura bool) (*DBTransacao, error) {
	t := &DBTransacao{}
	db := "academia"
	if apenasLeitura {
		db += "-ro"
	}

	pgsql, err := conexoes[db].NewTx(ctx, &sql.TxOptions{
		ReadOnly:  apenasLeitura,
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, err
	}

	t.postgres = pgsql.(*sql.Tx)
	t.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(t)
	t.ctx = ctx

	return t, nil
}

// Commit aplica as alterações pendentes em uma transacao
func (t *DBTransacao) Commit() (err error) {
	err = t.postgres.Commit()
	return
}

// Rollback desfaz todas as alterações pendentes em uma transacao
func (t *DBTransacao) Rollback() {
	_ = t.postgres.Rollback()
}

// Exec implementa a interface do método Exec
func (t *DBTransacao) Exec(query string, args ...interface{}) (sql.Result, error) {
	return t.postgres.Exec(obterCaller()+query, args...)
}

// QueryRow implementa a interface do método QueryRow
func (t *DBTransacao) QueryRow(query string, args ...interface{}) sq.RowScanner {
	return t.postgres.QueryRow(obterCaller()+query, args...)
}

// Query implementa a interface do método Query
func (t *DBTransacao) Query(query string, args ...interface{}) (*sql.Rows, error) {
	ch := make(chan bool, 1)
	conf := config.ObterConfiguracao()

	go func() {
		select {
		case <-time.After(time.Duration(conf.ConsultaTempoLimite * float32(time.Second))):
			logMsg := map[string]interface{}{
				"aviso":    LimiteAvisoConsulta,
				"consulta": query,
				"valores":  args,
			}

			msg, err := json.Marshal(logMsg)
			if err != nil {
				log.Printf("%+v\n", logMsg)
			} else {
				log.Println(string(msg))
			}

			return
		case <-ch:
			return
		}
	}()

	r, err := t.postgres.Query(obterCaller()+query, args...)
	ch <- true
	return r, err
}

func obterCaller() (saida string) {
	var (
		arquivo string
		linha   int
	)

	saida = "-- \r\n"

	for i := 3; i < 8; i++ {
		_, arquivo, linha, _ = runtime.Caller(i)
		if strings.Contains(arquivo, "/services/") {
			saida += "-- " + strconv.FormatInt(int64(i), 10) + " :: " + arquivo + ":" + strconv.FormatInt(int64(linha), 10) + "\r\n"
		}
	}

	return
}
