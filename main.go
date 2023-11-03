package main

import (
	"github.com/vinidotruan/alura-go-rest-api/database"
	"github.com/vinidotruan/alura-go-rest-api/models"
	"github.com/vinidotruan/alura-go-rest-api/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Teste", History: "Teste", Id: 1},
	}
	database.Connect()
	routes.HandleRequest()
}
