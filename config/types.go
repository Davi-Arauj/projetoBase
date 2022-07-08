package config

// Config main configuration struct
type Config struct {
	BancosDados         []BancoDados `json:"bancos_dados"`
	CookieNome          string       `json:"cookie_nome"`
	Secrets             []string     `json:"secrets"`
	Secret              string       `json:"secret"`
	PermissaoBase       string       `json:"permissao_base"`
	EnderecoExterno     string       `json:"endereco_externo"`
	EnderecoInterno     string       `json:"endereco_interno"`
	DiretorioLogAcesso  string       `json:"diretorio_log_acesso"`
	DiretorioLogErro    string       `json:"diretorio_log_erro"`
	NotificacaoSistema  string       `json:"notificacao_sistema"`
	NotificacaoURL      string       `json:"notificacao_url"`
	DiretorioUpload     string       `json:"diretorio_upload"`
	AutenticacaoIndex   string       `json:"autenticacao_index"`
	UploadTamanhoMax    int64        `json:"upload_tamanho_max"`
	IntervaloNoficacao  int64        `json:"intervalo_noficacao"`
	ConsultaTempoLimite float32      `json:"consulta_tempo_limite"`
	Producao            bool         `json:"producao"`
	Metricas            bool         `json:"metricas"`
	HeartbeatIntervalo  int64        `json:"heartbeat_intervalo"`
	HeartbeatEndpoint   string       `json:"heartbeat_endpoint"`
	RedesPermitidas     []string     `json:"redes_permitidas"`
	DiretorioTex        string       `json:"diretorio_tex"`
}

// BancoDados contém dados necessários para manter uma conexão com
// o banco de dados
type BancoDados struct {
	Nick               string   `json:"nick"`
	Name               string   `json:"name"`
	Username           string   `json:"username"`
	Password           string   `json:"password"`
	Host               string   `json:"hostname"`
	Port               string   `json:"port"`
	MaxConn            int      `json:"max_conn"`
	MaxIdle            int      `json:"max_idle"`
	ReadOnly           bool     `json:"read_only"`
	Main               bool     `json:"main"`
	Addresses          []string `json:"addresses"`
	TransactionTimeout int      `json:"transaction_timeout"`
}
