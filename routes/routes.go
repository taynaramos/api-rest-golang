package routes

import (
	"log"
	"net/http"

	"github.com/taynaramos/api-rest-golang/controllers"
)

// criando rota para a função home
func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil)) // quando der algo errado mostrar log e ouvir e subir na porta 8000
}
