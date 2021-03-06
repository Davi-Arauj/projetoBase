package util

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// FlagFiltro é usada para definir um tipo filtro
type FlagFiltro int

// Usado para especificar o tipo de filtro que está sendo usado
const (
	FlagFiltroNenhum FlagFiltro = 1 << iota
	FlagFiltroEq
	FlagFiltroIn
	FlagFiltroNotIn
	FlagFiltroArray

	// MaxLimit define o valor máximo que uma listagem pode requerir
	MaxLimit int = 100000
)

// Filtro é a representação de um filtro disponível
type Filtro struct {
	Valor   string
	Flag    FlagFiltro
	Tamanho int
}

// ParametrosRequisicao é usado para requisições do tipo GET quando
// o parâmetros de query são necessários
type ParametrosRequisicao struct {
	Campos      []string
	OrderCampo  string
	CamposNome  map[string]string
	Limite      uint64
	Offset      uint64
	Filtros     map[string][]string
	OrderByNome bool
	Desc        bool
	Total       bool
	Aggregate   bool
	Chart       bool
}

// ParseParams recebe um gin.Context e prepara os parâmetros de query da requisição
func ParseParams(c *gin.Context) (parametros ParametrosRequisicao, err error) {
	lim, err := strconv.Atoi(c.DefaultQuery("limit", "15"))
	if err != nil {
		return
	}

	if lim <= 0 {
		lim = MaxLimit // maximum value
	}
	parametros.Limite = uint64(Min(lim, MaxLimit))

	off, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		return
	}
	parametros.Offset = uint64(off)

	parametros.Campos, _ = c.GetQueryArray("campo")

	parametros.OrderCampo = c.DefaultQuery("order", "")

	parametros.Desc, err = strconv.ParseBool(c.DefaultQuery("desc", "false"))
	if err != nil {
		return
	}

	if parametros.Total, err = strconv.ParseBool(c.DefaultQuery("total", "false")); err != nil {
		return
	}

	if parametros.Aggregate, err = strconv.ParseBool(c.DefaultQuery("aggregate", "false")); err != nil {
		return
	}

	if parametros.Chart, err = strconv.ParseBool(c.DefaultQuery("chart", "false")); err != nil {
		return
	}

	parametros.Filtros = map[string][]string{}
	for k, v := range c.Request.URL.Query() {
		if k == "limit" || k == "offset" || k == "order" || k == "campo" || k == "desc" {
			continue
		}

		if len(v) > 0 {
			parametros.Filtros[k] = append(parametros.Filtros[k], v...)
		}
	}

	return
}

// ValidarCampos retorna todos os campos válidos a partir de uma
// estrutura dada
func (p *ParametrosRequisicao) ValidarCampos(dst interface{}, opts ...map[string]string) (f []string, d []interface{}, err error) {
	const (
		srcTag      = "codinome"
		dstTag      = "sql"
		aliasTag    = "alias"
		distinctTag = "distinct"
	)
	var (
		enableDistinct bool
	)

	p.CamposNome = make(map[string]string)
	elemDst := reflect.ValueOf(dst).Elem()

	if elemDst.Kind() != reflect.Struct {
		err = errors.New("Não é uma estrutura")
		return
	}

	if elemDst.NumField() == 0 {
		err = errors.New("Nenhum campo disponível")
		return
	}

	if p.Total {
		count := "count(1)"
		if len(opts) > 0 && opts[0]["count"] != "" {
			count = opts[0]["count"]
		}
		f = append(f, count)
		return
	}

	if len(opts) > 0 && opts[0]["distinct"] != "" {
		enableDistinct = true
	}

	apply := func(v string) {
		for s := 0; s < elemDst.NumField(); s++ {
			field := elemDst.Type().Field(s)
			inTag := field.Tag.Get(srcTag)
			outTag := field.Tag.Get(dstTag)
			alias := field.Tag.Get(aliasTag)
			distinct := field.Tag.Get(distinctTag)
			if inTag == "" || outTag == "" {
				continue
			}
			internal := ""

			if v == inTag || v == "" {
				pt := reflect.New(reflect.PtrTo(elemDst.Field(s).Type()))
				pt.Elem().Set(elemDst.Field(s).Addr())

				if alias != "" {
					internal = alias + "." + outTag
					outTag = internal + " AS " + inTag
				} else {
					outTag = outTag + " AS " + inTag
				}

				if distinct != "" && enableDistinct {
					outTag = "DISTINCT ON (" + internal + ") " + outTag
				}
				f = append(f, outTag)
				p.CamposNome[outTag] = inTag
				d = append(d, pt.Elem().Interface())
			}
		}
	}

	if len(p.Campos) > 0 {
		for _, v := range p.Campos {
			apply(v)
		}
	} else {
		apply("")
	}

	return
}

// CriarFiltros retorna um squirrel.SelectBuilder com todos os filtros aplicados a ele
func (p *ParametrosRequisicao) CriarFiltros(builder sq.SelectBuilder, disponiveis map[string]Filtro) sq.SelectBuilder {
	for k := range disponiveis {
		var v = disponiveis[k]
		for k1, v1 := range p.Filtros {
			if k == k1 {
				v.Valor = "( " + v.Valor + " )"
				switch v.Flag {
				case FlagFiltroIn:
					builder = builder.Where(sq.Eq{
						v.Valor: v1,
					})
				case FlagFiltroNotIn:
					builder = builder.Where(sq.NotEq{
						v.Valor: v1,
					})
				case FlagFiltroEq:
					builder = builder.Where(v.Valor, func(xs []string) (v []interface{}) {
						for x := range xs {
							v = append(v, xs[x])
						}
						return
					}(v1[0:v.Tamanho])...)
				case FlagFiltroArray:
					builder = builder.Where(func(q string, qtd int) string {
						if strings.Contains(q, "$$$") {
							placeholders := make([]string, 0, qtd)
							for i := 0; i < qtd; i++ {
								placeholders = append(placeholders, "?")
							}

							tmp := strings.Join(placeholders, ",")

							q = strings.Replace(q, "$$$", "ARRAY["+tmp+"]", 1)
						}
						return q
					}(v.Valor, len(v1)), func(str []string) []interface{} {
						itf := make([]interface{}, 0, len(str))
						for i := range str {
							itf = append(itf, str[i])
						}
						return itf
					}(v1)...)
				}
			}
		}
	}

	return builder
}

// CriarFiltros cria os filtro
func CriarFiltros(v string, flag FlagFiltro, tamanho ...int) Filtro {
	if len(tamanho) == 0 {
		tamanho = []int{1}
	}

	return Filtro{
		Valor:   v,
		Flag:    flag,
		Tamanho: tamanho[0],
	}
}

// ValidarOrdenador checa se o ordenador definido é válido ou não e retorna
// a clausula de ordenação
func (p *ParametrosRequisicao) ValidarOrdenador(dst interface{}) string {
	const (
		srcTag   = "codinome"
		dstTag   = "sql"
		aliasTag = "alias"
	)

	elemDst := reflect.ValueOf(dst).Elem()

	if elemDst.Kind() != reflect.Struct {
		return ""
	}

	if elemDst.NumField() == 0 {
		return ""
	}

	fst := ""
	for s := 0; s < elemDst.NumField(); s++ {
		field := elemDst.Type().Field(s)
		inTag := field.Tag.Get(srcTag)
		outTag := field.Tag.Get(dstTag)
		alias := field.Tag.Get(aliasTag)
		if inTag == "" || outTag == "" {
			continue
		}

		if fst == "" {
			if p.OrderByNome {
				fst = inTag
			} else {
				fst = outTag
				if alias != "" {
					fst = alias + "." + fst
				}
			}
		}

		if p.OrderCampo == inTag {
			if p.OrderByNome {
				fst = inTag
			} else {
				if alias != "" {
					fst = alias + "." + outTag
				} else {
					fst = outTag
				}
			}
		}
	}

	if p.Desc {
		fst += " DESC"
	} else {
		fst += " ASC"
	}

	return fst
}
