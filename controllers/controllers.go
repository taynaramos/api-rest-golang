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
	// para adicionar config no header da requisição e devolver um json
	// w.Header().Set("Content-type", "application/json")

	// lista de personalidades
	var p []models.Personalidade
	// usando o Find do GORM para encontrar todas essas listas de personalidades passando o endereço de memória de p
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

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	// pegando dado do request
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	// colocando na base de dados
	database.DB.Create(&novaPersonalidade)
	// exibindo o que foi criado
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	// exibindo a personalidade que foi deletada
	json.NewEncoder(w).Encode(personalidade)
}

func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)          // acha a personalidade que deve ser alterada
	json.NewDecoder(r.Body).Decode(&personalidade) // pega o body do request
	database.DB.Save(&personalidade)               // pede pro database salvar no endereço de memória de personalidade

	json.NewEncoder(w).Encode(personalidade)
}
