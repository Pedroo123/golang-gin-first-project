package database

import (
	"github.com/Pedroo123/golang-gin-first-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDb() {
	stringConexao := "host=localhost user=postgres password=postgres dbname=alunos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		panic("Erro ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
