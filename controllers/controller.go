package controllers

import (
	"github.com/Pedroo123/golang-gin-first-project/database"
	"github.com/Pedroo123/golang-gin-first-project/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func ExibeAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"message": "Eae " + nome,
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o aluno",
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(200, aluno)
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível atualizar o aluno",
		})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(&aluno)
	c.JSON(200, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(200, gin.H{
		"message": "Aluno deletado com sucesso",
	})
}

func ExibeAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}
