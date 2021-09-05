package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func NewConfig() (Config, error) {
	var config Config
	err := envconfig.Process("OPA", &config)
	if err != nil {
		log.Fatalf("Fail to load config wiht env : %v", err)
		return config, err
	}
	return config, nil
}

type Config struct {
	ApiKey         string `envconfig:"APIKEY" split_words:"true"`
	ApiSecret      string `envconfig:"APISECRET" split_words:"true"`
	BASEURL        string `envconfig:"BASEURL" split_words:"true"`
	ASSUMEMERCHANT string `envconfig:"ASSUMEMERCHANT" split_words:"true"`
}
