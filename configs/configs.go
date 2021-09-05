package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func NewConfig() (ApiConfig, error) {
	var config ApiConfig
	err := envconfig.Process("OPA", &config)
	if err != nil {
		log.Fatalf("Fail to load config wiht env : %v", err)
		return config, err
	}
	return config, nil
}

type ApiConfig struct {
	ApiKey         string `envconfig:"APIKEY" split_words:"true"`
	ApiSecret      string `envconfig:"APISECRET" split_words:"true"`
	BASEURL        string `envconfig:"BASEURL" split_words:"true"`
	ASSUMEMERCHANT string `envconfig:"ASSUMEMERCHANT" split_words:"true" default:""`
}
