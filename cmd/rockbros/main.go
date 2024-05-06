package main

import (
	"log"
	"net/http"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/redaced/rockbros/pkg/api/controllers"
	"github.com/redaced/rockbros/pkg/api/handlers"
	"github.com/redaced/rockbros/pkg/api/middleware"
	"github.com/redaced/rockbros/pkg/models"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginController).Methods("POST")

	s := r.PathPrefix("/").Subrouter()
	s.Use(middleware.AuthMiddleware)
	s.HandleFunc("/home", handlers.Hello).Methods("GET")
	s.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	err := models.ConnectDB("root:password@/rockbros?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Apply the CORS middleware to your router with the default options
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8080"}))(r)

	http.ListenAndServe(":3000", corsHandler)
}

//.j6NQU5tDK$zbgC
//root0011))