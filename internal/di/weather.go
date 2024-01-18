package di

import (
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/entity"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/infra/viacep"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/infra/weatherapi"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/infra/web"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/usecase"
)

func ConfigWebController() entity.WeatherController {
	httpClient := http.DefaultClient

	cepHttpClient := viacep.NewCepHTTPClient(httpClient)
	weatherHttpClient := weatherapi.NewWeatherHTTPClient(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(cepHttpClient, weatherHttpClient)
	weatherController := web.NewWebController(weatherUseCase)

	return weatherController
}
