package main

import (
	"fmt"

	"github.com/tfrancar/api-rest-go/routes"
)

func main() {
	fmt.Println("Iniciando o servidor rest com go")
	routes.HandleRequest()
}
