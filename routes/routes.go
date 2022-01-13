package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taynaramos/api-rest-golang/controllers"
)

// criando rota para a função home
func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET") // por padrão o método é GET
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r)) // quando der algo errado mostrar log e ouvir e subir na porta 8000
}
