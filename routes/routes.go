package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vinidotruan/alura-go-rest-api/controllers"
	"github.com/vinidotruan/alura-go-rest-api/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.StorePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
