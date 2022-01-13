package main

import (
	"fmt"

	"github.com/taynaramos/api-rest-golang/database"
	"github.com/taynaramos/api-rest-golang/models"
	"github.com/taynaramos/api-rest-golang/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
