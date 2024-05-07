package routes

import (
	"net/http"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/redaced/rockbros/pkg/api/controllers"
	"github.com/redaced/rockbros/pkg/api/handlers"
	"github.com/redaced/rockbros/pkg/api/middleware"
)

func SetupRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/homeUnauth", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controllers.LoginController).Methods("POST")

	s := r.PathPrefix("/").Subrouter()
	s.Use(middleware.AuthMiddleware)
	s.HandleFunc("/home", handlers.Hello).Methods("GET")
	s.HandleFunc("/logout", controllers.LogoutHandler).Methods("POST")

	// Apply the CORS middleware to your router with the default options
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8080"}))(r)

	http.ListenAndServe(":3000", corsHandler)

}
