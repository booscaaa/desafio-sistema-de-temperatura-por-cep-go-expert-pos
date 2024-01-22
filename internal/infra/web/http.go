package web

import (
	"encoding/json"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/dto"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/entity"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/pkg/adapter/errorhandle"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	usecase   entity.WeatherUseCase
	validator *validator.Validate
}

func NewWebController(usecase entity.WeatherUseCase, validator *validator.Validate) entity.WeatherController {
	return &controller{
		usecase:   usecase,
		validator: validator,
	}
}

// Get goDoc
// @Summary Get temperature
// @Description Get temperature in celcius, kelvin and fahrenheit
// @Tags cep
// @Accept  json
// @Produce  json
// @Param cep path string true "cep"
// @Success 200 {object} entity.Weather
// @Failure 404 {object} errorhandle.Response
// @Failure 422 {object} errorhandle.Response
// @Router /cep/{cep} [get]
func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	cep, err := dto.FromQueryStringRequestToCep(chi.URLParam(request, "cep"), controller.validator)

	if err != nil {
		statusCode, message := errorhandle.Handle(errorhandle.ErrUnprocessableEntity)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weather, err := controller.usecase.Get(request.Context(), cep.Cep)

	if err != nil {
		statusCode, message := errorhandle.Handle(err)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	json.NewEncoder(response).Encode(weather)
}
