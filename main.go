package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tfrancar/api-rest-go/database"
	"github.com/tfrancar/api-rest-go/models"
	"github.com/tfrancar/api-rest-go/routes"
)

func main() {

	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env", err)
	}
	
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	database.ConectaBancoDeDados()

	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleRequest()
}
