package routes

import (
	"github.com/Pedroo123/golang-gin-first-project/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCPF)
	r.Run()
}
