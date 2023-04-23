package database

import (
	"github.com/JulioSoaresA/go-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	conexao := "user=postgres dbname=go-gin-api password=postgres host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
