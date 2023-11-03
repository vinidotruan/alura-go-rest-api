package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinidotruan/alura-go-rest-api/models"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}

func GetPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errId := strconv.Atoi(vars["id"])

	if errId != nil {
		panic(errId.Error())
	}

	for _, personality := range models.Personalities {
		if personality.Id == id {
			json.NewEncoder(w).Encode(personality)
		}
	}
}
