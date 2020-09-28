package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"Inlog.Go.Service.Api/service"
	"github.com/go-chi/chi"
)

//BaiscController controller
type BaiscController struct{}

// Routes rotas api
func (mr BaiscController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/veiculos", mr.Get)
	return r
}

// Get de Veiculos
func (mr BaiscController) Get(w http.ResponseWriter, r *http.Request) {
	dto, err := service.Service.ConsultarDados()
	js, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	log.Print("Dados Retornados com sucesso!")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
