# Desafio do Sistema de Temperatura com CEP
## Objetivo
Este projeto em Go tem como objetivo desenvolver um sistema que, ao receber um CEP válido, identifica a cidade correspondente e retorna o clima atual em diferentes unidades de temperatura: Celsius, Fahrenheit e Kelvin. O sistema é hospedado no Google Cloud Run.

## Requisitos do Sistema
- Entrada de CEP: O sistema aceita um CEP válido de 8 dígitos.
- Localização e Clima: Realiza a pesquisa do CEP, identifica a localização e retorna a temperatura atual nas unidades Celsius, Fahrenheit e Kelvin.
- Respostas HTTP:
  - Em caso de sucesso: HTTP 200 com corpo { "temp_C": [valor], "temp_F": [valor], "temp_K": [valor] }.
  - Falha por CEP inválido: HTTP 422 com mensagem "invalid zipcode".
  - Falha por CEP não encontrado: HTTP 404 com mensagem "can not found zipcode".

## Serviços Utilizados
- APIs Externas:
  - ViaCEP para localização.
  - WeatherAPI para informações climáticas.

## Instruções de Execução
### Ambiente de Desenvolvimento
- Pré-requisitos: Certifique-se de ter o Go, Docker e Docker Compose instalados.
- Clonar o repositório: 
  ```go
  git clone https://github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos.git
  ```
- Acessar a pasta: 
  ```go
  cd desafio-sistema-de-temperatura-por-cep-go-expert-pos.git
  ```
- Configure a variavel com a chave da api do seu WeatherAPI:
  ```
  cp config.example.json config.json
  nano config.json

  {
    "weather_api_key": "<API_KEY>"
  }
  ```
- Execute:
  ```go
  docker compose up --build -d
  ```
- Server: http://localhost:8080
- Swagger: http://localhost:8080/swagger/index.html
- Endpoint para validação: http://localhost:8080/cep/{CEP}

### Ambiente de Produção
  - Server: https://desafio-sistema-de-temperatura-por-cep-go-expert--i7bjj44hla-rj.a.run.app
  - Swagger: https://desafio-sistema-de-temperatura-por-cep-go-expert--i7bjj44hla-rj.a.run.app/swagger/index.html
  - Endpoint para validação: https://desafio-sistema-de-temperatura-por-cep-go-expert--i7bjj44hla-rj.a.run.app/cep/{CEP}


# Testes Automatizados
Os testes cobrem as funcionalidades chave, incluindo a validação de CEPs e a conversão de temperaturas.
### Conversões de Temperatura
- Fahrenheit: F = C * 1.8 + 32
- Kelvin: K = C + 273

### Execute os testes automatizados com os comandos: 
  - Teste e2e e lógica de conversão de temperatura:
    ```bash
    go test ./tests/e2e/... ./internal/entity/...
    ```

  - Teste e2e:
    ```bash
    go test ./tests/e2e/...
    ```

  - Teste da lógica de conversão de temperatura:
    ```bash
    go test ./internal/entity/...
    ```


