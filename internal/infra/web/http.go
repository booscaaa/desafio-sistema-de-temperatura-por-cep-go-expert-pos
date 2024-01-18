package web

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/entity"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/pkg/adapter/errorhandle"
	"github.com/go-chi/chi"
)

type controller struct {
	usecase entity.WeatherUseCase
}

func NewWebController(usecase entity.WeatherUseCase) entity.WeatherController {
	return &controller{
		usecase: usecase,
	}
}

func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	cep := chi.URLParam(request, "cep")
	weather, err := controller.usecase.Get(request.Context(), cep)

	if err != nil {
		statusCode, message := errorhandle.Handle(err)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	json.NewEncoder(response).Encode(weather)
}
