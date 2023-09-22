package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tfrancar/api-rest-go/database"
	"github.com/tfrancar/api-rest-go/routes"
)

func main() {

	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env", err)
	}

	database.ConectaBancoDeDados()

	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleRequest()
}
