package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/taynaramos/api-rest-golang/models"
)

// criando home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

// verifica todas as personalidades para saber qual dá match com a que é passada por parametro na url
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// se assemelha com match.params
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		// strconv serve para converter um int para string
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
