package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var (
	DB      *gorm.DB
	err     error
	USER    = "POSTGRES_USER"
	PASS    = "POSTGRES_PASSWORD"
	HOST    = "POSTGRES_HOST"
	PORT    = "POSTGRES_PORT"
	DB_NAME = "POSTGRES_DB_NAME"
)

func ConectaBancoDeDados() {
	db_user := os.Getenv(USER)
	db_pass := os.Getenv(PASS)
	db_host := os.Getenv(HOST)
	db_port := os.Getenv(PORT)
	db_name := os.Getenv(DB_NAME)
	stringDeConexao := fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%s sslmode=disable",
		db_host, db_user, db_pass, db_name, db_port)
	//"host=localhost user=root  password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
