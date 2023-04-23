package controllers

import (
	"github.com/JulioSoaresA/go-gin-api/database"
	"github.com/JulioSoaresA/go-gin-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func ExibeTodosAlunos(c *gin.Context) []models.Aluno {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	//c.JSON(200, alunos)
	return alunos
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Olá " + nome + ", bem vindo a API",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	aluno.Nome = c.PostForm("nome")
	aluno.CPF = c.PostForm("cpf")
	aluno.RG = c.PostForm("rg")

	database.DB.Create(&aluno)

	location := url.URL{Path: "/"}
	c.Redirect(http.StatusMovedPermanently, location.RequestURI())
}

func Index(c *gin.Context) {
	alunos := ExibeTodosAlunos

	c.HTML(http.StatusOK, "index.html", alunos)
}

func Cadastrar(c *gin.Context) {
	c.HTML(http.StatusOK, "cadastra_aluno.html", gin.H{
		"content": "This is an Cadastro page...",
	})
}
