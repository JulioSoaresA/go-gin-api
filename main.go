package main

import (
	"github.com/JulioSoaresA/go-gin-api/database"
	"github.com/JulioSoaresA/go-gin-api/models"
	"github.com/JulioSoaresA/go-gin-api/routes"
)

func main() {
	database.ConectaComDB()
	models.Alunos = []models.Aluno{
		{Nome: "Julio Soares", CPF: "00000000000", RG: "001000000"},
		{Nome: "Monalisa Maria", CPF: "11111111111", RG: "002000000"},
	}
	routes.HandleRequests()
}
