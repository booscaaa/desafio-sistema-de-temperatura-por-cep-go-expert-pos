package entity

import (
	"context"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/dto"
)

type Weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (weather *Weather) CalculateFarenheit() {
	weather.TempF = weather.TempC*1.8 + 32
}

func (weather *Weather) CalculateKelvin() {
	weather.TempK = weather.TempC + 273
}

type WeatherHTTPClient interface {
	Get(context.Context, string) (*dto.WeatherOutput, error)
}

type WeatherUseCase interface {
	Get(context.Context, string) (*Weather, error)
}

type WeatherController interface {
	Get(http.ResponseWriter, *http.Request)
}
