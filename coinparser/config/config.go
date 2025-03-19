package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Server   server   `mapstructure:"server"`
	Database database `mapstructure:"database"`
	Parser   parser   `mapstructure:"parser"`
}

type server struct {
	Port              string        `mapstructure:"port"`
	ReadTimeout       time.Duration `mapstructure:"readTimeout"`
	WriteTimeout      time.Duration `mapstructure:"writeTimeout"`
	IdleTimeout       time.Duration `mapstructure:"idleTimeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"readHeaderTimeout"`
}

type database struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

// parser contains configuration for the CoinMarketCap API parser
// 
// ApiKeyCoinMarketCap is the API key for CoinMarketCap
// 
// Timestamp is the time in seconds between each full update of the database
// 
// TimeForReq is the time in seconds for the timeout when sending a request to
// CoinMarketCap API
type parser struct {
	ApiKeyCoinMarketCap string        `mapstructure:"cmc_apikey"`
	Timestamp           time.Duration `mapstructure:"timestamp"`
	TimeForReq          time.Duration `mapstructure:"timeout_for_req"`
}

var defaultConfig config
var defaultpath string = "parserconfig.json"

// init initializes the configuration by setting default values for server, database, 
// and parser settings using Viper. It also checks for a specified config file path 
// through command-line arguments and attempts to merge it with the defaults. If the 
// config file is not found or an error occurs, it logs a warning and continues with 
// default values. Finally, it unmarshals the configuration into the defaultConfig 
// struct and logs a fatal error if unmarshalling fails.
func init() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.readTimeout", time.Second*5)
	viper.SetDefault("server.writeTimeout", time.Second*5)
	viper.SetDefault("server.idleTimeout", time.Second*20)
	viper.SetDefault("server.readHeaderTimeout", time.Second*5)

	viper.SetDefault("database.name", "coinder")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", os.Getenv("POSTGRES_PASSWORD"))
	viper.SetDefault("database.host", "database")
	viper.SetDefault("database.port", "3306")

	viper.SetDefault("parser.cmc_apikey", os.Getenv("CMC_API_KEY"))
	viper.SetDefault("parser.timestamp", time.Hour*24)
	viper.SetDefault("parser.timeout_for_req", time.Second*10)

	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	viper.SetConfigType("json")
	if *configPath != "" {
		viper.SetConfigFile(*configPath)
	} else {
		viper.SetConfigName(defaultpath)
		viper.AddConfigPath(".")
	}

	if err := viper.MergeInConfig(); err != nil {
		log.Printf("Warning: Failed to read configpath: %s, error: %v. Using default values.", defaultpath, err)
	}

	if err := viper.Unmarshal(&defaultConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}

// Config returns the default configuration
func Config() config {
	return defaultConfig
}

// Server returns the default server configuration from the loaded config.
func Server() server {
	return defaultConfig.Server
}

// Database returns the default database configuration from the loaded config.
func Database() database {
	return defaultConfig.Database
}

// Parser returns the default parser configuration from the loaded config.
func Parser() parser {
	return defaultConfig.Parser
}