package weatherapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

func (httpclient httpclient) Get(ctx context.Context, cityName string) (*entity.Weather, error) {
	weatherApiKey := viper.GetString("weather_api_key")
	var weatherOutput dto.WeatherOutput

	params := url.Values{}
	params.Add("key", weatherApiKey)
	params.Add("q", cityName)
	params.Add("aqi", "no")

	url := fmt.Sprintf("%s?%s", BASE_URL, params.Encode())

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

	return &entity.Weather{
		TempC: weatherOutput.Current.TempC,
	}, nil
}
