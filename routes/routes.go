package routes

import (
	"github.com/JulioSoaresA/go-gin-api/controllers"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

func HandleRequests() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", controllers.Index)
	r.GET("/novo_aluno", controllers.Cadastrar)
	//r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/cadastrar_aluno", controllers.CriaNovoAluno)
	r.Run(":8000")
}
