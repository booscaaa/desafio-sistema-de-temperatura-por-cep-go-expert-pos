// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Vinícius Boscardin",
            "email": "boscardinvinicius@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cep/{cep}": {
            "get": {
                "description": "Get temperature in celcius, kelvin and fahrenheit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cep"
                ],
                "summary": "Get temperature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cep",
                        "name": "cep",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Weather"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/errorhandle.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errorhandle.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Weather": {
            "type": "object",
            "properties": {
                "temp_C": {
                    "type": "number"
                },
                "temp_F": {
                    "type": "number"
                },
                "temp_K": {
                    "type": "number"
                }
            }
        },
        "errorhandle.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Desafio Sistema de Temperatura por Cep Go Expert API Docs",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
