package main

import (
	"github.com/psnavega/api-go-gin/database"
	"github.com/psnavega/api-go-gin/src/models"
	"github.com/psnavega/api-go-gin/src/routes"
)

func main() {
	database.ConnDatabase()
	models.Students = []models.Student{
		{
			Name: "Patrick Navega",
			CPF:  "7777777777",
			RG:   "99999",
		},
		{
			Name: "Renato Chagas",
			CPF:  "8888888",
			RG:   "1515151",
		},
	}
	routes.HandleRequests()
}
