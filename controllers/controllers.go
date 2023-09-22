package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tfrancar/api-rest-go/database"
	"github.com/tfrancar/api-rest-go/models"
)

// underscore usado para suprimir o r
func Home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, _ *http.Request) {
	var array_personalidades []models.Personalidade
	database.DB.Find(&array_personalidades)
	json.NewEncoder(w).Encode(array_personalidades)
}

func RetornaUmaPersonaldiade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

}

func CriarUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)

}
