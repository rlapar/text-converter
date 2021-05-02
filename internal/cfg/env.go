package cfg

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	Port		int		`env:"HTTP_PORT" envDefault:"8080"`
	Environment string 	`env:"ENVIRONMENT" envDefault:"sandbox"`
	Username	string	`env:"USERNAME" envDefault:"username"`
	Password	string	`env:"PASSWORD" envDefault:"password"`
	RequestId 	string
}

//singleton config
var config *Config

func init() {
	config = loadConfig()
}

func loadConfig() *Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &config
}

func GetConfig() *Config {
	return config
}