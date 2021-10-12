package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
var RegistrarUsuarioRoute = func(router *mux.Router) {
	router.HandleFunc("/usuario/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/usuario/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/usuario/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuario/{id}", controllers.DeleteUser).Methods("DELETE")
	//	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
}
*/
func Manejadores() {

	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
