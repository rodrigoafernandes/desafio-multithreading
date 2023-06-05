package main

import "github.com/kelseyhightower/envconfig"

type AppConfig struct {
	ApiCepUrl    string `envconfig:"APICEP_URL"`
	ViaCepUrl    string `envconfig:"VIACEP_URL"`
	ApiTimeoutMS int    `envconfig:"API_TIMEOUT_MS"`
}

var AppCFG AppConfig

func Setup() {
	var appConfig AppConfig
	if err := envconfig.Process("", &appConfig); err != nil {
		panic(err)
	}
	AppCFG = appConfig
}
