package routes

import (
	"github.com/gorilla/mux"
	"github.com/vinidotruan/alura-go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", r))
}
