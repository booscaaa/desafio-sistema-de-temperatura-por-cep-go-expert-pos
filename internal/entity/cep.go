package entity

import (
	"context"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/dto"
)

type CepHTTPClient interface {
	Get(context.Context, string) (*dto.CepOutput, error)
}
