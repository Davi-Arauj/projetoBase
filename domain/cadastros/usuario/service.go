package usuario

import "github.com/projetoBase/config/database"

// Servico define a estrutura base para uso dos métodos do serviço
type Servico struct {
	repo IUsuario
}

// ObterServico retorna um servico para acesso a funções de auxilio
// a lógica de negócio
func ObterServico(r IUsuario) *Servico {
	return &Servico{repo: r}
}

// ObterRepo retorna um repositório para acesso à camada de dados
func ObterRepo(tx *database.DBTransacao) IUsuario {
	return novoRepo(tx)
}