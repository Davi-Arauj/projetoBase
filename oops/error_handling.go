package oops

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/projetoBase/util"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
)

const (
	pgxCode         = 1000
	jsonCode        = 2000
	internalCode    = 3000
	defaultCode     = 4000
	validationCode  = 5000
	grpcCode        = 6000
	timeParseError  = 7000
	httpRequestCode = 8000
)

// Error define um tipo erro para tratamento
type Error struct {
	Msg        string   `json:"msg"`
	Code       int      `json:"code"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
	StatusCode int      `json:"-"`
}

// Error implementa a interface do tipo error
func (e *Error) Error() string {
	return e.Msg
}

// Unwrap retorna a causa especifica para um erro
func (e *Error) Unwrap() error {
	return e.Err
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type wrappedError interface {
	Unwrap() error
}

// Err constroi um instancia de erro a partir de um erro
func Err(err error) error {
	var e *Error
	if !errors.As(err, &e) {
		// trate o erro caso ele não esteja tratado
		err = tratarErro(err)
	} else if err == e {
		err = tratarErro(err)
	}
	return errors.WithStack(err)
}

// DefinirErro adicona uma mensagem e um status code para o erro
func DefinirErro(err error, c *gin.Context) {
	var e *Error

	if !errors.As(err, &e) {
		DefinirErro(Err(err), c)
		return
	}
	e.Msg = err.Error()
	e.Trace, _ = reconstruirStackTrace(err, e)

	c.JSON(e.StatusCode, e)
	c.Set("error", err)
	c.Abort()
}

func reconstruirStackTrace(err error, bound error) (output []string, traced bool) {
	var (
		wrapped wrappedError
		tracer  stackTracer
	)
	if errors.As(err, &wrapped) {
		internal := wrapped.Unwrap()
		if internal != bound {
			output, traced = reconstruirStackTrace(internal, bound)
		}
		if !traced && errors.As(err, &tracer) {
			stack := tracer.StackTrace()
			for _, frame := range stack {
				output = append(output, fmt.Sprintf("%+v", frame))
			}
			traced = true
		}
	}
	return
}

// tratarErro trata o erro para prover um mesagem para o usuário
func tratarErro(rawError error) error {
	msg, code, responseStatus := "Erro desconhecido", 0, 400
	switch err := rawError.(type) {
	// erros de entrada
	case *json.UnmarshalTypeError:
		msg, code = fmt.Sprintf("Tipo de valor %v não suportado no campo %v. Esperado tipo %v", err.Value, err.Field, err.Type.String()), jsonCode+1

	case validator.ValidationErrors:
		msg, code = parseValidationError(err)

	// erros internos
	case *reflect.ValueError:
		msg, code = fmt.Sprintf("Não é possível acessar o valor do tipo %v", err.Kind.String()), internalCode+1

	case *strconv.NumError:
		msg, code = fmt.Sprintf("Não é possível converter valor %v", err.Num), internalCode+2

	// erros de dados
	case pgx.PgError:
		msg, code = handlePgxError(&err)
		rawError = errors.Errorf("%s: %s", err.Error(), err.Hint)

	case *url.Error:
		msg, code = fmt.Sprintf("Falha no acesso à serviço. Operação: %v", err.Op), internalCode+3

	case *time.ParseError:
		msg, code = fmt.Sprintf("Impossível converter %v", err.Value), timeParseError+1

	case *Error:
		rawError, msg, code, responseStatus = err, err.Msg, err.Code, err.StatusCode

	case *util.HTTPErro:
		msg, code = handleHTTPRequestError(err, httpRequestCode)

	case error:
		// Erros padrões
		switch err {
		case sql.ErrNoRows:
			msg, code = "Referência inválida", defaultCode+1
			responseStatus = http.StatusNotFound

		case io.EOF:
			msg, code = "Nenhum dado disponível para leitura", defaultCode+2
		}

		// Erros externos de grpc
		if s, ok := grpcStatus.FromError(err); ok {
			msg, code = s.Message(), grpcCode+int(s.Code())
			rawError = fmt.Errorf(fmt.Sprintf("%v", s.Details()))
			if s.Code() == grpcCodes.DeadlineExceeded {
				msg = "A consulta demorou mais do que o esperado, tente novamente."
			}
		}
	case nil:
		return nil
	}

	return &Error{
		Msg:        msg,
		Err:        rawError,
		Code:       code,
		StatusCode: responseStatus,
	}
}

// Wrap encapsula o erro adicionando um mensagem
func Wrap(err error, mensagem string) error {
	return errors.Wrap(Err(err), mensagem)
}


// NovoErr cria uma nova instância de erro padrão
func NovoErr(mensagem string) error {
	return Err(&Error{
		Msg:        mensagem,
		Err:        errors.Errorf("Mensagem de erro interna: '%s'. Veja a stack para esse erro para ter informações adicionais.", mensagem),
		Code:       defaultCode,
		StatusCode: http.StatusBadRequest,
	})
}

func parseValidationError(err validator.ValidationErrors) (msg string, code int) {
	msg, code = "Não foi possível definir o erro de validação", validationCode

	if len(err) == 0 {
		return
	}

	switch err[0].ActualTag() {
	case "required":
		msg, code = "Campo "+err[0].Field()+" é obrigatorio", validationCode+1
	case "gt":
		msg, code = "Campo "+err[0].Field()+" deve ser maior que "+err[0].Param(), validationCode+2
	case "lt":
		msg, code = "Campo "+err[0].Field()+" deve ser menor que "+err[0].Param(), validationCode+2
	case "customerDocument":
		msg, code = "Documento inválido", validationCode+3
	case "gte":
		msg, code = "Campo "+err[0].Field()+" deve ser maior ou igual a "+err[0].Param(), validationCode+4
	case "lte":
		msg, code = "Campo "+err[0].Field()+" deve ser menor ou igual a "+err[0].Param(), validationCode+4
	case "stringField":
		msg, code = "Campo "+err[0].Field()+" não é uma string valida", validationCode+5
	case "required_with":
		msg, code = "Campo "+err[0].Field()+" é obrigatório quando campo "+err[0].Param()+" é enviado", validationCode+6
	case "required_without":
		msg, code = "Campo "+err[0].Field()+" é obrigatório se não for enviado o campo "+err[0].Param(), validationCode+7
	case "email":
		msg, code = "Campo "+err[0].Field()+" não contém email válido "+err[0].Param(), validationCode+8
	case "len":
		msg, code = "Campo "+err[0].Field()+" deve possuir tamanho igual a "+err[0].Param(), validationCode+9
	case "min":
		switch err[0].Kind() {
		case reflect.Int64, reflect.Int, reflect.Float64:
			msg, code = "Campo "+err[0].Field()+" deve possuir um valor de no mínimo "+err[0].Param(), validationCode+10
		case reflect.Array, reflect.Slice, reflect.String:
			msg, code = "Campo "+err[0].Field()+" deve possuir um tamanho de no mínimo "+err[0].Param(), validationCode+10
		default:
			msg, code = "Campo "+err[0].Field()+" deve possuir no mínimo "+err[0].Param(), validationCode+10
		}
	case "max":
		switch err[0].Kind() {
		case reflect.Int64, reflect.Int, reflect.Float64:
			msg, code = "Campo "+err[0].Field()+" deve possuir um valor de no máximo "+err[0].Param(), validationCode+11
		case reflect.Array, reflect.Slice, reflect.String:
			msg, code = "Campo "+err[0].Field()+" deve possuir um tamanho de no máximo "+err[0].Param(), validationCode+11
		default:
			msg, code = "Campo "+err[0].Field()+" deve possuir no máximo "+err[0].Param(), validationCode+11
		}
	}

	return
}

func handlePgxError(err *pgx.PgError) (string, int) {
	switch err.Code {
	case "23505":
		return "Registro duplicado", pgxCode + 1
	case "23502":
		return "Dado requerido não foi especificado", pgxCode + 2
	case "23503":
		return "Dado indicado não é uma referência válida", pgxCode + 3
	case "42P01", "42703":
		return "Acesso incorreto de elementos nos registros de dados: erro de síntaxe", pgxCode + 4
	case "42601", "42803", "42883":
		return "Uso incorreto de função ou operador durante acesso aos registros de dados: erro de sintax", pgxCode + 5
	case "22001":
		return "Dado excede capacidade do registro no banco de dados", pgxCode + 6
	case "42702":
		return "Referência ambigua: erro de sintax", pgxCode + 7
	}

	return "Erro de dados desconhecido", pgxCode
}

func handleHTTPRequestError(err *util.HTTPErro, baseCode int) (string, int) {
	servicoDescricao := "serviço externo"
	metodoDescricao := "executar ação solicitada"
	statusDescricao := "ocorreu um problema inesperado"
	usuarioInstrucao := "tente novamente mais tarde"

	switch err.Metodo {
	case "GET":
		metodoDescricao = "buscar dados"
	case "POST":
		metodoDescricao = "cadastrar novos dados"
	case "PUT":
		metodoDescricao = "atualizar informações"
	case "DELETE":
		metodoDescricao = "remover registro"
	}

	// verifica o serviço externo
	if err.ClienteNome != "" {
		servicoDescricao = "serviço de " + err.ClienteNome
	}

	// identifica a ação que causou o problema
	if err.TagRequisicao != "" {
		metodoDescricao = err.TagRequisicao
	}

	switch {
	case err.StatusCode >= 300 && err.StatusCode < 400:
		statusDescricao = "não está aceitando requisições"
	case err.StatusCode >= 400 && err.StatusCode < 500:
		statusDescricao = "foi passado parâmetros incorretos"
	case err.StatusCode >= 500:
		statusDescricao = "encontrou um problema interno"
	case err.StatusCode == -1:
		statusDescricao = "encontra-se indisponível"
	}

	if cause, ok := err.Causa.(net.Error); ok {
		if cause.Timeout() {
			statusDescricao = "demorou mais que o esperado"
			usuarioInstrucao = "tente novamente"
		}
	}

	return fmt.Sprintf("Não foi possível %s. O %s %s, %s.", metodoDescricao, servicoDescricao, statusDescricao, usuarioInstrucao), baseCode
}
