package rest

import (
	"log"
	"net/http"

	"github.com/booscaaa/desafio-sistema-de-temperatura-por-cep-go-expert-pos/internal/di"
	"github.com/go-chi/chi"
)

func Initialize() {
	webController := di.ConfigWebController()

	r := chi.NewRouter()
	r.Get("/cep/{cep}", webController.Get)

	log.Println("Rodando na porta: 8080")
	http.ListenAndServe(":8080", r)
}
