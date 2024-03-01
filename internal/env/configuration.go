package env

import (
	"backend-comee/internal/ciphers"
	"encoding/json"
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

		stringConfig := `{
			"app":{
				"service_name": "backend-cff",
					"port": 6032,
					"allowed_domains": "*",
					"path_log": "./log",
					"log_review_interval":60,
					"register_log": true,
					"rsa_private_key": "rsa/app.rsa",
					"rsa_public_key": "rsa/app.rsa.pub",
					"logger_http": false,
					"is_cipher": false,
					"validate_ip": "U2FsdGVkX19v3BWlEyIr/BvH568E5r+dq/03m9qNZBQ=",
					"path_directory": "excel",
					"index_separator": 10
			},
			"db": {
				"engine": "postgres",
					"server": "roundhouse.proxy.rlwy.net",
					"port": 38816,
					"name": "railway",
					"user": "postgres",
					"password": "eD5c5G4cE--cFfE43eb6A213aa-CCbbg",
					"instance": "",
					"is_secure": false
			},
			"smtp" : {
				"port": 587,
					"host": "in-v3.mailjet.com",
					"email": "07366c0397ed0e45ed555e75c48247c3",
					"password": "2565011db3640655f1c16e9ad2669561"
			},
			"external": {
				"reniec": "https://dniruc.apisperu.com/api/v1/dni/",
					"credential": "?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6InlvbmlsLnJvamFzQGUtY2FwdHVyZS5jbyJ9.Rcu71yck4VAjnNAdlrNpXl-IzfKFicaGv0zpBJstMQQ"
			}
		}`

		/*b, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}*/

		/*err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}*/

		err := json.Unmarshal([]byte(stringConfig), &config)
		if err != nil {
			log.Fatalf("Error al deserializar el JSON: %v", err)
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
