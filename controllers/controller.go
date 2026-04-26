package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Pedroo123/golang-gin-first-project/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H {
		"message": "Eae " + nome,
	})
}