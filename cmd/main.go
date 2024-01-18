/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/cmd/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("Erro ao ler o arquivo .env: %s", err))
	}
}

func main() {
	cmd.Execute()
}
