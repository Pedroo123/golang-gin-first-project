package main

import (
	"github.com/Pedroo123/golang-gin-first-project/database"
	"github.com/Pedroo123/golang-gin-first-project/models"
	"github.com/Pedroo123/golang-gin-first-project/routes"
)

func main() {
	database.ConnectDb()
	models.Alunos = []models.Aluno{
		{Nome: "Pedro", CPF: "12345678910", RG: "123456789"},
		{Nome: "Maria", CPF: "10987654321", RG: "987654321"},
	}
	routes.HandleRequests()
}
