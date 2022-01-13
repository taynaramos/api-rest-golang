package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taynaramos/api-rest-golang/database"
	"github.com/taynaramos/api-rest-golang/models"
)

// criando home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

// verifica todas as personalidades para saber qual dá match com a que é passada por parametro na url
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// se assemelha com match.params
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
