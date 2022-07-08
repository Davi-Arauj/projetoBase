package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var config *Config

// CarregarConfiguracao carrega as configurações através do arquivo definido
func CarregarConfiguracao() {
	path := "/etc/academia/config.json"
	if val, set := os.LookupEnv("PLAYGROUND_CONFIG"); set && val != "" {
		path = val
	} else {
		log.Println("variável de ambiente `PLAYGROUND_CONFIG` não está definida, usado diretorio: ", path)
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal(err)
	}

	if err := validarConfig(); err != nil {
		log.Fatal(err)
	}
}

// ObterConfiguracao retorna um ponteiro para uma estrutura de configuração
// que contém todos os dados da configuração
func ObterConfiguracao() *Config {
	if config == nil {
		log.Fatal("a configuração não pode ser carregada")
	}
	return config
}

func validarConfig() error {
	if config == nil {
		return errors.New("a configuração não pode ser carregada")
	}

	if err := validarBancoDadosPrincipal(); err != nil {
		return err
	}

	if config.EnderecoExterno == "" {
		config.EnderecoExterno = ":9876"
	}

	if config.EnderecoInterno == "" {
		config.EnderecoInterno = ":8765"
	}

	if config.CookieNome == "" {
		config.CookieNome = "NovoRevan"
	}

	if len(config.Secrets) == 0 {
		return errors.New("O secret da aplicação não foi definido")
	}

	if config.DiretorioLogAcesso == "" {
		config.DiretorioLogAcesso = "/var/log/academia/access.log"
	}

	if config.DiretorioLogErro == "" {
		config.DiretorioLogErro = "/var/log/academia/error.log"
	}

	if config.ConsultaTempoLimite == 0.0 {
		config.ConsultaTempoLimite = 2.0
	}

	if config.NotificacaoSistema == "" {
		config.NotificacaoSistema = "academia"
	}

	if config.NotificacaoURL == "" {
		config.NotificacaoURL = "http://localhost:6060/v1"
	}

	if config.AutenticacaoIndex == "" {
		config.AutenticacaoIndex = "auth"
	}

	if config.IntervaloNoficacao == 0 {
		config.IntervaloNoficacao = 5
	}

	return nil
}

func validarBancoDadosPrincipal() error {
	if len(config.BancosDados) == 0 {
		return errors.New("Configuração de banco de dados não definida")
	}

	main := false
	for i := 0; i < len(config.BancosDados); i++ {
		if config.BancosDados[i].Main {
			if main {
				return errors.New("Mais de um banco de dados foi definido como principal")
			}
			main = true
		}

		if config.BancosDados[i].TransactionTimeout == 0 {
			config.BancosDados[i].TransactionTimeout = 1
		}
	}

	if !main {
		return errors.New("É necessário definir um banco de dados principal")
	}

	return nil
}
