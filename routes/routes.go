package routes

import (
	"log"
	"net/http"

	"github.com/tfrancar/api-rest-go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
