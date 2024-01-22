/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/pkg/adapter/http/rest"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/pkg/adapter/validator"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		validator := validator.Initialize()
		rest.Initialize(validator)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
