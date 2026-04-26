package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Pedroo123/golang-gin-first-project/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET(":nome", controllers.Saudacao)

	r.Run()
}