package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gabriel-hawerroth/api-go-example/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter a lista de registros: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
