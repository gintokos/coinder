package config

import (
	"flag"
	"time"

	"github.com/spf13/viper"
)

var Config = config{}

func GetConfig() *config {
	return &Config
}

const (
	LOCAL            = "local"
	LOCAL_WITH_NGROK = "local_with_ngrok"
	DEV              = "dev"
	PROD             = "prod"
)

type config struct {
	Env         string   `mapstructure:"env"`
	NgrokDomain string   `mapstructure:"ngrok_domain"`
	BotToken    string   `mapstructure:"botToken"`
	Domain      string   `mapstructure:"domain"`
	Server      server   `mapstructure:"server"`
	Database    database `mapstructure:"database"`
	Parser      parser   `mapstructure:"parser"`
}

type server struct {
	Host              string        `mapstructure:"host"`
	Port              string        `mapstructure:"port"`
	ReadTimeout       time.Duration `mapstructure:"readTimeout"`
	WriteTimeout      time.Duration `mapstructure:"writeTimeout"`
	IdleTimeout       time.Duration `mapstructure:"idleTimeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"readHeaderTimeout"`
}

type parser struct {
	ApiKeyCoinMarketCap string        `mapstructure:"cmc_apikey"`
	Timestamp           time.Duration `mapstructure:"timestamp"`
	TimeForReq          time.Duration `mapstructure:"timeout_for_req"`
}

type database struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

var defaultServer = server{
	Host:              "127.0.0.1",
	Port:              "8080",
	ReadTimeout:       time.Second * 5,
	WriteTimeout:      time.Second * 5,
	IdleTimeout:       time.Second * 20,
	ReadHeaderTimeout: time.Second * 5,
}

func MustInitForApp() {
	viper.SetDefault("server", defaultServer)

	LoadConfig("appconfig.json")
}

func MustInitForParser() {
	LoadConfig("parserconfig")
}

func LoadConfig(defaultpath string) error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	viper.SetConfigType("json")

	if *configPath != "" {
		viper.SetConfigFile(*configPath)
	} else {
		viper.SetConfigName(defaultpath)
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
