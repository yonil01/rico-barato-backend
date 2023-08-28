package env

import (
	"backend-ccff/internal/ciphers"
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	once   sync.Once
	config = &configuration{}
)

type configuration struct {
	App      App      `json:"app"`
	DB       DB       `json:"db"`
	External External `json:"external"`
}

type App struct {
	ServiceName       string `json:"service_name"`
	Port              int    `json:"port"`
	AllowedDomains    string `json:"allowed_domains"`
	PathLog           string `json:"path_log"`
	LogReviewInterval int    `json:"log_review_interval"`
	RegisterLog       bool   `json:"register_log"`
	RSAPrivateKey     string `json:"rsa_private_key"`
	RSAPublicKey      string `json:"rsa_public_key"`
	LoggerHttp        bool   `json:"logger_http"`
	IsCipher          bool   `json:"is_cipher"`
	ValidateIp        string `json:"validate_ip"`
	TLS               bool   `json:"tls"`
	Cert              string `json:"cert"`
	Key               string `json:"key"`
	PathDirectory     string `json:"path_directory"`
	IndexSeparator    int    `json:"index_separator"`
}

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

type External struct {
	Reniec     string `json:"reniec"`
	Credential string `json:"credential"`
}

func NewConfiguration() *configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}
		if config.App.IsCipher {
			if config.DB.Password = ciphers.Decrypt(config.DB.Password); config.DB.Password == "" {
				log.Fatal("no se pudo obtener config.DB.Password Decrypt")
			}

			if config.App.ValidateIp = ciphers.Decrypt(config.App.ValidateIp); config.App.ValidateIp == "" {
				log.Fatal("no se pudo obtener config.DB.Password Decrypt")
			}
		}
	})
}
