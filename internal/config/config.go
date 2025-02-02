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

type config struct {
	Database database
	Parser   parser `mapstructure:"parser"`
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

func MustInitForApp() {
	// viper.SetDefault()

}

func MustInitForParser() {
	// viper.SetDefault()
	LoadConfig("parserconfig")
}

func LoadConfig(defaultpath string) error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	viper.SetConfigFile("json")

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
