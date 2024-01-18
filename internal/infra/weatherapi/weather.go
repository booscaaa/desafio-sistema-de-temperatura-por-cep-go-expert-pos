package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/dto"
	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/entity"
	"github.com/spf13/viper"
)

var (
	BASE_URL = "https://api.weatherapi.com/v1/current.json"
)

type httpclient struct {
	client *http.Client
}

func NewWeatherHTTPClient(client *http.Client) entity.WeatherHTTPClient {
	return &httpclient{
		client: client,
	}
}

func (httpclient httpclient) Get(ctx context.Context, cityName string) (*dto.WeatherOutput, error) {
	weatherApiKey := viper.GetString("WEATHER_API_KEY")
	var weatherOutput dto.WeatherOutput

	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", BASE_URL, weatherApiKey, cityName)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, bytes.NewReader(nil))

	if err != nil {
		return nil, err
	}

	defer request.Body.Close()

	response, err := httpclient.client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&weatherOutput)

	if err != nil {
		return nil, err
	}

	return &weatherOutput, nil
}
